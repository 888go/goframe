// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile

import (
	"context"
	"time"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gcache"
	"github.com/888go/goframe/os/gfsnotify"
)

const (
	defaultCacheDuration  = "1m"             // defaultCacheExpire 是文件内容缓存的默认过期时间，单位为秒。
	commandEnvKeyForCache = "gf.gfile.cache" // commandEnvKeyForCache 是用于配置缓存过期时间的命令行参数或环境变量的配置键。
)

var (
	// 默认的文件内容缓存过期时间。
	cacheDuration = getCacheDuration()

	// internalCache 是内部使用的内存缓存。
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

// GetContentsWithCache 通过`path`从缓存返回指定文件的字符串内容。
// 如果缓存中没有内容，则会从由`path`指定的磁盘文件中读取内容。
// 参数`expire`指定了此文件内容在缓存中的有效期，单位为秒。
func GetContentsWithCache(path string, duration ...time.Duration) string {
	return string(GetBytesWithCache(path, duration...))
}

// GetBytesWithCache 函数通过 `path` 从缓存中获取指定文件的 []byte 内容。
// 如果缓存中没有内容，会从由 `path` 指定的磁盘文件中读取内容。
// 参数 `expire` 指定了此文件内容在缓存中的有效期，单位为秒。
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
// 若该文件有任何变化，将会清除其缓存。
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
