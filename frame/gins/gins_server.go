// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gins

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/internal/consts"
	"github.com/gogf/gf/v2/internal/instance"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// Server 函数返回一个指定名称的 http 服务器实例。
// 注意，如果在创建实例过程中发生任何错误，它将引发 panic。 md5:09f2ffb4e37b28a6
func Server(name ...interface{}) *ghttp.Server {
	var (
		err          error
		ctx          = context.Background()
		instanceName = ghttp.DefaultServerName
		instanceKey  = fmt.Sprintf("%s.%v", frameCoreComponentNameServer, name)
	)
	if len(name) > 0 && name[0] != "" {
		instanceName = gconv.String(name[0])
	}
	return instance.GetOrSetFuncLock(instanceKey, func() interface{} {
		server := ghttp.GetServer(instanceName)
		if Config().Available(ctx) {
			// 从配置中初始化服务器。 md5:0d26e3a48d836ee7
			var (
				configMap             map[string]interface{}
				serverConfigMap       map[string]interface{}
				serverLoggerConfigMap map[string]interface{}
				configNodeName        string
			)
			if configMap, err = Config().Data(ctx); err != nil {
				intlog.Errorf(ctx, `retrieve config data map failed: %+v`, err)
			}
			// 根据可能的名称查找服务器配置项。 md5:9a64296188cf11c8
			if len(configMap) > 0 {
				if v, _ := gutil.MapPossibleItemByKey(configMap, consts.ConfigNodeNameServer); v != "" {
					configNodeName = v
				}
				if configNodeName == "" {
					if v, _ := gutil.MapPossibleItemByKey(configMap, consts.ConfigNodeNameServerSecondary); v != "" {
						configNodeName = v
					}
				}
			}
			// 根据实例名称自动获取配置。 md5:7140fae5fa8c1aec
			serverConfigMap = Config().MustGet(
				ctx,
				fmt.Sprintf(`%s.%s`, configNodeName, instanceName),
			).Map()
			if len(serverConfigMap) == 0 {
				serverConfigMap = Config().MustGet(ctx, configNodeName).Map()
			}
			if len(serverConfigMap) > 0 {
				if err = server.SetConfigWithMap(serverConfigMap); err != nil {
					panic(err)
				}
			} else {
				// 配置不是必需的，所以它只会打印内部日志。 md5:4f36a44ab1690d54
				intlog.Printf(
					ctx,
					`missing configuration from configuration component for HTTP server "%s"`,
					instanceName,
				)
			}
			// 服务器日志配置检查。 md5:22846287f007266c
			serverLoggerConfigMap = Config().MustGet(
				ctx,
				fmt.Sprintf(`%s.%s.%s`, configNodeName, instanceName, consts.ConfigNodeNameLogger),
			).Map()
			if len(serverLoggerConfigMap) == 0 && len(serverConfigMap) > 0 {
				serverLoggerConfigMap = gconv.Map(serverConfigMap[consts.ConfigNodeNameLogger])
			}
			if len(serverLoggerConfigMap) > 0 {
				if err = server.Logger().SetConfigWithMap(serverLoggerConfigMap); err != nil {
					panic(err)
				}
			}
		}
		// 服务器名称是必需的。如果未配置，默认情况下它将设置一个服务器名称。 md5:db6207755fb7815b
		if server.GetName() == "" || server.GetName() == ghttp.DefaultServerName {
			server.SetName(instanceName)
		}
		// 由于可能使用了模板功能，
		// 它也会初始化视图实例。 md5:7c98f1273cb1ece0
		_ = getViewInstance()
		return server
	}).(*ghttp.Server)
}
