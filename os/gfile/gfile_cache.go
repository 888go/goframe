// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类

import (
	"context"
	"time"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/internal/intlog"
	gcache "github.com/888go/goframe/os/gcache"
	gfsnotify "github.com/888go/goframe/os/gfsnotify"
)

const (
	defaultCacheDuration  = "1m"             // defaultCacheExpire是文件内容缓存的过期时间（以秒为单位）。 md5:93f4150c6283fef8
	commandEnvKeyForCache = "gf.gfile.cache" // commandEnvKeyForCache 是用于配置缓存过期持续时间的命令行参数或环境变量的配置键。 md5:e8e411869780802b
)

var (
		// 默认的文件内容缓存过期时间。 md5:848c5089a9dc23eb
	cacheDuration = getCacheDuration()

		// internalCache是内部使用的内存缓存。 md5:5cd10c891525ec8d
	internalCache = gcache.X创建()
)

func getCacheDuration() time.Duration {
	cacheDurationConfigured := command.GetOptWithEnv(commandEnvKeyForCache, defaultCacheDuration)
	d, err := time.ParseDuration(cacheDurationConfigured)
	if err != nil {
		panic(gerror.X多层错误码并格式化(
			gcode.CodeInvalidConfiguration,
			err,
			`error parsing string "%s" to time duration`,
			cacheDurationConfigured,
		))
	}
	return d
}

// X缓存读文本 通过`path`从缓存中返回给定文件的字符串内容。如果缓存中没有内容，它将从指定的磁盘文件（由`path`提供）中读取。参数`expire`指定了该文件内容的缓存过期时间（以秒为单位）。
// md5:ee3ca4011fe59d23
func X缓存读文本(路径 string, 缓存时长 ...time.Duration) string {
	return string(X缓存读字节集(路径, 缓存时长...))
}

// X缓存读字节集 通过`path`从缓存中返回给定文件的[]byte内容。
// 如果缓存中没有内容，它将从由`path`指定的磁盘文件中读取。
// 参数`expire`以秒为单位指定该文件内容的缓存时间。
// md5:8b877378627c94a2
func X缓存读字节集(路径 string, 缓存时长 ...time.Duration) []byte {
	var (
		ctx      = context.Background()
		expire   = cacheDuration
		cacheKey = commandEnvKeyForCache + 路径
	)

	if len(缓存时长) > 0 {
		expire = 缓存时长[0]
	}
	r, _ := internalCache.X取值或设置值_并发安全函数(ctx, cacheKey, func(ctx context.Context) (interface{}, error) {
		b := X读字节集(路径)
		if b != nil {
			// 将此`path`添加到gfsnotify，
			// 如果文件有任何更改，它将清除其缓存。
			// md5:d6795c29773b5d37
			_, _ = gfsnotify.Add(路径, func(event *gfsnotify.Event) {
				_, err := internalCache.X删除并带返回值(ctx, cacheKey)
				if err != nil {
					intlog.Errorf(ctx, `%+v`, err)
				}
				gfsnotify.Exit()
			})
		}
		return b, nil
	}, expire)
	if r != nil {
		return r.X取字节集()
	}
	return nil
}
