// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gfile

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/command"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfsnotify"
)

const (
	defaultCacheDuration  = "1m"             // defaultCacheExpire是文件内容缓存的过期时间（以秒为单位）。 md5:93f4150c6283fef8
	commandEnvKeyForCache = "gf.gfile.cache" // commandEnvKeyForCache 是用于配置缓存过期持续时间的命令行参数或环境变量的配置键。 md5:e8e411869780802b
)

var (
	// 默认的文件内容缓存过期时间。 md5:848c5089a9dc23eb
	cacheDuration = getCacheDuration()

	// internalCache是内部使用的内存缓存。 md5:5cd10c891525ec8d
	internalCache = gcache.New()
)

func getCacheDuration() time.Duration {
	cacheDurationConfigured := command.GetOptWithEnv(commandEnvKeyForCache, defaultCacheDuration)
	d, err := time.ParseDuration(cacheDurationConfigured)
	if err != nil {
		panic(gerror.WrapCodef(
			gcode.CodeInvalidConfiguration,
			err,
			`error parsing string "%s" to time duration`,
			cacheDurationConfigured,
		))
	}
	return d
}

// GetContentsWithCache 通过`path`从缓存中返回给定文件的字符串内容。如果缓存中没有内容，它将从指定的磁盘文件（由`path`提供）中读取。参数`expire`指定了该文件内容的缓存过期时间（以秒为单位）。 md5:ee3ca4011fe59d23
func GetContentsWithCache(path string, duration ...time.Duration) string {
	return string(GetBytesWithCache(path, duration...))
}

// GetBytesWithCache 通过`path`从缓存中返回给定文件的[]byte内容。
// 如果缓存中没有内容，它将从由`path`指定的磁盘文件中读取。
// 参数`expire`以秒为单位指定该文件内容的缓存时间。 md5:8b877378627c94a2
func GetBytesWithCache(path string, duration ...time.Duration) []byte {
	var (
		ctx      = context.Background()
		expire   = cacheDuration
		cacheKey = commandEnvKeyForCache + path
	)

	if len(duration) > 0 {
		expire = duration[0]
	}
	r, _ := internalCache.GetOrSetFuncLock(ctx, cacheKey, func(ctx context.Context) (interface{}, error) {
		b := GetBytes(path)
		if b != nil {
			// 将此`path`添加到gfsnotify，
			// 如果文件有任何更改，它将清除其缓存。 md5:d6795c29773b5d37
			_, _ = gfsnotify.Add(path, func(event *gfsnotify.Event) {
				_, err := internalCache.Remove(ctx, cacheKey)
				if err != nil {
					intlog.Errorf(ctx, `%+v`, err)
				}
				gfsnotify.Exit()
			})
		}
		return b, nil
	}, expire)
	if r != nil {
		return r.Bytes()
	}
	return nil
}
