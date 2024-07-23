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

	"github.com/gogf/gf/v2/internal/consts"
	"github.com/gogf/gf/v2/internal/instance"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gutil"
)

// Log 返回一个 glog.Logger 的实例。
// 参数 `name` 是该实例的名称。
// 请注意，如果在创建实例过程中发生任何错误，它将引发 panic。
// md5:9578e0721b3e1c0b
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
		// 为了避免在不必要的时候出现文件未找到的错误。 md5:bb553aa936e788c9
		var (
			configMap      map[string]interface{}
			loggerNodeName = consts.ConfigNodeNameLogger
		)
		// 尝试以不区分大小写的方式查找可能的`loggerNodeName`。 md5:1c5803eaec1f4f06
		if configData, _ := Config().Data(ctx); len(configData) > 0 {
			if v, _ := gutil.MapPossibleItemByKey(configData, consts.ConfigNodeNameLogger); v != "" {
				loggerNodeName = v
			}
		}
		// 通过logger名称获取特定的日志配置。 md5:d07e016490e45c59
		certainLoggerNodeName := fmt.Sprintf(`%s.%s`, loggerNodeName, instanceName)
		if v, _ := Config().Get(ctx, certainLoggerNodeName); !v.IsEmpty() {
			configMap = v.Map()
		}
		// 如果没有为特定日志器名称获取到配置，则从全局日志配置中检索。 md5:40acd9c22e400784
		if len(configMap) == 0 {
			if v, _ := Config().Get(ctx, loggerNodeName); !v.IsEmpty() {
				configMap = v.Map()
			}
		}
		// 如果配置映射不为空，设置日志配置。 md5:7b4a24e8a6a74588
		if len(configMap) > 0 {
			if err := logger.SetConfigWithMap(configMap); err != nil {
				panic(err)
			}
		}
		return logger
	}).(*glog.Logger)
}
