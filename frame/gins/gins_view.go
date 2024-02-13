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
	"github.com/888go/goframe/os/gview"
	"github.com/888go/goframe/util/gutil"
)

// View 返回一个使用默认设置的 View 实例。
// 参数 `name` 是该实例的名称。
// 注意：如果在创建实例过程中发生任何错误，将会导致程序 panic。
func View(name ...string) *模板类.View {
	instanceName := 模板类.DefaultName
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	instanceKey := fmt.Sprintf("%s.%s", frameCoreComponentNameViewer, instanceName)
	return instance.GetOrSetFuncLock(instanceKey, func() interface{} {
		return getViewInstance(instanceName)
	}).(*模板类.View)
}

func getViewInstance(name ...string) *模板类.View {
	var (
		err          error
		ctx          = context.Background()
		instanceName = 模板类.DefaultName
	)
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	view := 模板类.Instance(instanceName)
	if Config().X是否可用(ctx) {
		var (
			configMap      map[string]interface{}
			configNodeName = consts.ConfigNodeNameViewer
		)
		if configMap, err = Config().X取Map(ctx); err != nil {
			intlog.Errorf(ctx, `retrieve config data map failed: %+v`, err)
		}
		if len(configMap) > 0 {
			if v, _ := 工具类.MapPossibleItemByKey(configMap, consts.ConfigNodeNameViewer); v != "" {
				configNodeName = v
			}
		}
		configMap = Config().X取值PANI(ctx, fmt.Sprintf(`%s.%s`, configNodeName, instanceName)).X取Map()
		if len(configMap) == 0 {
			configMap = Config().X取值PANI(ctx, configNodeName).X取Map()
		}
		if len(configMap) > 0 {
			if err = view.SetConfigWithMap(configMap); err != nil {
				panic(err)
			}
		}
	}
	return view
}
