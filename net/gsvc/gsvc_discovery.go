// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gsvc

import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/util/gutil"
)

// watchedMap 存储发现对象及其被监视的服务映射。
var watchedMap = map类.X创建(true)

// ServiceWatch 用于监视服务状态。
type ServiceWatch func(service Service)

// Get通过服务名称获取并返回服务。
func Get(ctx context.Context, name string) (service Service, err error) {
	return GetAndWatchWithDiscovery(ctx, defaultRegistry, name, nil)
}

// GetWithDiscovery 通过 `discovery` 获取并返回指定服务名的服务。
func GetWithDiscovery(ctx context.Context, discovery Discovery, name string) (service Service, err error) {
	return GetAndWatchWithDiscovery(ctx, discovery, name, nil)
}

// GetAndWatch 用于获取服务，并使用自定义的 watch 回调函数进行监听。
func GetAndWatch(ctx context.Context, name string, watch ServiceWatch) (service Service, err error) {
	return GetAndWatchWithDiscovery(ctx, defaultRegistry, name, watch)
}

// GetAndWatchWithDiscovery 用于在`discovery`中使用自定义的观察回调函数获取服务。
func GetAndWatchWithDiscovery(ctx context.Context, discovery Discovery, name string, watch ServiceWatch) (service Service, err error) {
	if discovery == nil {
		return nil, 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `discovery cannot be nil`)
	}
	// 通过发现对象获取服务映射。
	watchedServiceMap := watchedMap.X取值或设置值_函数(discovery, func() interface{} {
		return map类.X创建StrAny(true)
	}).(*map类.StrAnyMap)
	// 通过名称获取服务。
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
			err = 错误类.X创建错误码并格式化(错误码类.CodeNotFound, `service not found with name "%s"`, name)
			return nil
		}

		// 如果有多个，仅选择其中一个。
		service = services[0]

		// 在goroutine中监视服务变化。
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

// watchAndUpdateService 监视并更新服务，如果服务发生更改，则在内存中进行更新。
func watchAndUpdateService(watchedServiceMap *map类.StrAnyMap, watcher Watcher, service Service, watchFunc ServiceWatch) {
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
				工具类.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
					watchFunc(services[0])
				}, func(ctx context.Context, exception error) {
					intlog.Errorf(ctx, `%+v`, exception)
				})
			}
		}
	}
}

// Search 按照指定条件搜索并返回服务。
func Search(ctx context.Context, in SearchInput) ([]Service, error) {
	if defaultRegistry == nil {
		return nil, 错误类.X创建错误码并格式化(错误码类.CodeNotImplemented, `no Registry is registered`)
	}
	ctx, _ = context.WithTimeout(ctx, defaultTimeout)
	return defaultRegistry.Search(ctx, in)
}

// Watch 监视指定条件的变化。
func Watch(ctx context.Context, key string) (Watcher, error) {
	if defaultRegistry == nil {
		return nil, 错误类.X创建错误码并格式化(错误码类.CodeNotImplemented, `no Registry is registered`)
	}
	return defaultRegistry.Watch(ctx, key)
}
