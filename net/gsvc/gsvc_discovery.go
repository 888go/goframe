// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gsvc

import (
	"context"
	"time"

	gmap "github.com/888go/goframe/container/gmap"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	gutil "github.com/888go/goframe/util/gutil"
)

// watchedMap 存储了发现对象及其关注的服务映射。 md5:131b7e9a4eb2667c
var watchedMap = gmap.X创建(true)

// ServiceWatch 用于观察服务状态。 md5:075e96d2f9f06fe0
type ServiceWatch func(service Service)

// Get通过服务名称检索并返回服务。 md5:74843a42d759b705
func Get(ctx context.Context, name string) (service Service, err error) {
	return GetAndWatchWithDiscovery(ctx, defaultRegistry, name, nil)
}

// GetWithDiscovery 通过`discovery`中的服务名称检索并返回服务。 md5:f1ca28780ddf8348
func GetWithDiscovery(ctx context.Context, discovery Discovery, name string) (service Service, err error) {
	return GetAndWatchWithDiscovery(ctx, discovery, name, nil)
}

// GetAndWatch 用于获取服务并使用自定义的监视回调函数进行监视。 md5:9fa8d7df3bbbbe6d
func GetAndWatch(ctx context.Context, name string, watch ServiceWatch) (service Service, err error) {
	return GetAndWatchWithDiscovery(ctx, defaultRegistry, name, watch)
}

// GetAndWatchWithDiscovery 用于在`discovery`中获取服务并使用自定义的观察回调函数。 md5:07dc90075ba8e7c6
func GetAndWatchWithDiscovery(ctx context.Context, discovery Discovery, name string, watch ServiceWatch) (service Service, err error) {
	if discovery == nil {
		return nil, gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `discovery cannot be nil`)
	}
		// 通过发现对象获取服务映射。 md5:a91dd67cf7ae237f
	watchedServiceMap := watchedMap.X取值或设置值_函数(discovery, func() interface{} {
		return gmap.X创建StrAny(true)
	}).(*gmap.StrAnyMap)
		// 通过名称获取服务。 md5:6c20d3bc7e9e9d09
	storedService := watchedServiceMap.X取值或设置值_函数带锁(name, func() interface{} {
		var (
			services []Service
			watcher  Watcher
		)
		services, err = discovery.Search(ctx, SearchInput{
			Name: name,
		})
		if err != nil {
			return nil
		}
		if len(services) == 0 {
			err = gerror.X创建错误码并格式化(gcode.CodeNotFound, `service not found with name "%s"`, name)
			return nil
		}

				// 如果有多个，只选择一个。 md5:9ea08b90e82cd566
		service = services[0]

				// 在goroutine中观察服务的变化。 md5:365a338a3be90ab6
		if watch != nil {
			if watcher, err = discovery.Watch(ctx, service.GetPrefix()); err != nil {
				return nil
			}
			go watchAndUpdateService(watchedServiceMap, watcher, service, watch)
		}
		return service
	})
	if storedService != nil {
		service = storedService.(Service)
	}
	return
}

// watchAndUpdateService 监控服务并在其发生更改时更新内存中的服务。 md5:08346bbe119e4400
func watchAndUpdateService(watchedServiceMap *gmap.StrAnyMap, watcher Watcher, service Service, watchFunc ServiceWatch) {
	var (
		ctx      = context.Background()
		err      error
		services []Service
	)
	for {
		time.Sleep(time.Second)
		if services, err = watcher.Proceed(); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
			continue
		}
		if len(services) > 0 {
			watchedServiceMap.X设置值(service.GetName(), services[0])
			if watchFunc != nil {
				gutil.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
					watchFunc(services[0])
				}, func(ctx context.Context, exception error) {
					intlog.Errorf(ctx, `%+v`, exception)
				})
			}
		}
	}
}

// Search 搜索并返回符合指定条件的服务。 md5:62e529e326dae7b7
func Search(ctx context.Context, in SearchInput) ([]Service, error) {
	if defaultRegistry == nil {
		return nil, gerror.X创建错误码并格式化(gcode.CodeNotImplemented, `no Registry is registered`)
	}
	ctx, _ = context.WithTimeout(ctx, defaultTimeout)
	return defaultRegistry.Search(ctx, in)
}

// Watch 监视指定条件的变化。 md5:9fb048527d2a1698
func Watch(ctx context.Context, key string) (Watcher, error) {
	if defaultRegistry == nil {
		return nil, gerror.X创建错误码并格式化(gcode.CodeNotImplemented, `no Registry is registered`)
	}
	return defaultRegistry.Watch(ctx, key)
}
