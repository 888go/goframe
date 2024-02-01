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
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/util/gutil"
	)
// Log 返回一个 glog.Logger 实例。
// 参数 `name` 是该实例的名称。
// 注意：如果在创建实例过程中发生任何错误，它会触发 panic。
func Log(name ...string) *glog.Logger {
	var (
		ctx          = context.Background()
		instanceName = glog.DefaultName
	)
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	instanceKey := fmt.Sprintf("%s.%s", frameCoreComponentNameLogger, instanceName)
	return instance.GetOrSetFuncLock(instanceKey, func() interface{} {
		logger := glog.Instance(instanceName)
		// 为了避免在不必要的时候出现文件未找到错误
		var (
			configMap      map[string]interface{}
			loggerNodeName = consts.ConfigNodeNameLogger
		)
		// 尝试以不区分大小写的方式查找可能的`loggerNodeName`。
		if configData, _ := Config().Data(ctx); len(configData) > 0 {
			if v, _ := gutil.MapPossibleItemByKey(configData, consts.ConfigNodeNameLogger); v != "" {
				loggerNodeName = v
			}
		}
		// 通过日志器名称获取特定的日志器配置。
		certainLoggerNodeName := fmt.Sprintf(`%s.%s`, loggerNodeName, instanceName)
		if v, _ := Config().Get(ctx, certainLoggerNodeName); !v.IsEmpty() {
			configMap = v.Map()
		}
		// 如果不存在特定日志器名称的配置，则获取全局日志器配置。
		if len(configMap) == 0 {
			if v, _ := Config().Get(ctx, loggerNodeName); !v.IsEmpty() {
				configMap = v.Map()
			}
		}
		// 如果配置映射不为空，则设置日志器配置。
		if len(configMap) > 0 {
			if err := logger.SetConfigWithMap(configMap); err != nil {
				panic(err)
			}
		}
		return logger
	}).(*glog.Logger)
}
