// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gins

import (
	"context"
	"fmt"
	
	"github.com/888go/goframe/internal/consts"
	"github.com/888go/goframe/internal/instance"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

// Server 返回指定名称的 http 服务器实例。
// 注意，如果在创建实例期间发生任何错误，它会引发panic。
func Server(name ...interface{}) *http类.X服务 {
	var (
		err          error
		ctx          = context.Background()
		instanceName = http类.DefaultServerName
		instanceKey  = fmt.Sprintf("%s.%v", frameCoreComponentNameServer, name)
	)
	if len(name) > 0 && name[0] != "" {
		instanceName = 转换类.String(name[0])
	}
	return instance.GetOrSetFuncLock(instanceKey, func() interface{} {
		server := http类.X取服务对象(instanceName)
		if Config().X是否可用(ctx) {
			// 从配置中初始化服务器。
			var (
				configMap             map[string]interface{}
				serverConfigMap       map[string]interface{}
				serverLoggerConfigMap map[string]interface{}
				configNodeName        string
			)
			if configMap, err = Config().X取Map(ctx); err != nil {
				intlog.Errorf(ctx, `retrieve config data map failed: %+v`, err)
			}
			// 根据可能的名称查找可能的服务器配置项。
			if len(configMap) > 0 {
				if v, _ := 工具类.MapPossibleItemByKey(configMap, consts.ConfigNodeNameServer); v != "" {
					configNodeName = v
				}
				if configNodeName == "" {
					if v, _ := 工具类.MapPossibleItemByKey(configMap, consts.ConfigNodeNameServerSecondary); v != "" {
						configNodeName = v
					}
				}
			}
			// 自动通过实例名称获取配置。
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
				// 配置不是必需的，因此它仅打印内部日志。
				intlog.Printf(
					ctx,
					`missing configuration from configuration component for HTTP server "%s"`,
					instanceName,
				)
			}
			// 服务器日志记录器配置检查。
			serverLoggerConfigMap = Config().X取值PANI(
				ctx,
				fmt.Sprintf(`%s.%s.%s`, configNodeName, instanceName, consts.ConfigNodeNameLogger),
			).X取Map()
			if len(serverLoggerConfigMap) == 0 && len(serverConfigMap) > 0 {
				serverLoggerConfigMap = 转换类.X取Map(serverConfigMap[consts.ConfigNodeNameLogger])
			}
			if len(serverLoggerConfigMap) > 0 {
				if err = server.Logger别名().X设置配置Map(serverLoggerConfigMap); err != nil {
					panic(err)
				}
			}
		}
		// 服务器名称是必需的。如果未配置，它将设置一个默认服务器名称。
		if server.X取服务名称() == "" || server.X取服务名称() == http类.DefaultServerName {
			server.X设置服务名称(instanceName)
		}
// 由于可能会使用模板功能，
// 因此它也会初始化视图实例。
		_ = getViewInstance()
		return server
	}).(*http类.X服务)
}
