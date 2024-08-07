// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gins

import (
	"context"
	"fmt"

	"github.com/888go/goframe/internal/consts"
	"github.com/888go/goframe/internal/instance"
	"github.com/888go/goframe/internal/intlog"
	ghttp "github.com/888go/goframe/net/ghttp"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

// Server 函数返回一个指定名称的 http 服务器实例。
// 注意，如果在创建实例过程中发生任何错误，它将引发 panic。
// md5:09f2ffb4e37b28a6
func Server(name ...interface{}) *ghttp.X服务 {
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
		server := ghttp.X取服务对象(instanceName)
		if Config().X是否可用(ctx) {
						// 从配置中初始化服务器。 md5:0d26e3a48d836ee7
			var (
				configMap             map[string]interface{}
				serverConfigMap       map[string]interface{}
				serverLoggerConfigMap map[string]interface{}
				configNodeName        string
			)
			if configMap, err = Config().X取Map(ctx); err != nil {
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
			serverConfigMap = Config().X取值PANI(
				ctx,
				fmt.Sprintf(`%s.%s`, configNodeName, instanceName),
			).X取Map()
			if len(serverConfigMap) == 0 {
				serverConfigMap = Config().X取值PANI(ctx, configNodeName).X取Map()
			}
			if len(serverConfigMap) > 0 {
				if err = server.X设置配置项Map(serverConfigMap); err != nil {
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
			serverLoggerConfigMap = Config().X取值PANI(
				ctx,
				fmt.Sprintf(`%s.%s.%s`, configNodeName, instanceName, consts.ConfigNodeNameLogger),
			).X取Map()
			if len(serverLoggerConfigMap) == 0 && len(serverConfigMap) > 0 {
				serverLoggerConfigMap = gconv.X取Map(serverConfigMap[consts.ConfigNodeNameLogger])
			}
			if len(serverLoggerConfigMap) > 0 {
				if err = server.Logger别名().X设置配置Map(serverLoggerConfigMap); err != nil {
					panic(err)
				}
			}
		}
				// 服务器名称是必需的。如果未配置，默认情况下它将设置一个服务器名称。 md5:db6207755fb7815b
		if server.X取服务名称() == "" || server.X取服务名称() == ghttp.DefaultServerName {
			server.X设置服务名称(instanceName)
		}
		// 由于可能使用了模板功能，
		// 它也会初始化视图实例。
		// md5:7c98f1273cb1ece0
		_ = getViewInstance()
		return server
	}).(*ghttp.X服务)
}
