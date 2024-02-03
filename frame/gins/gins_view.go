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
func View(name ...string) *gview.View {
	instanceName := gview.DefaultName
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	instanceKey := fmt.Sprintf("%s.%s", frameCoreComponentNameViewer, instanceName)
	return instance.GetOrSetFuncLock(instanceKey, func() interface{} {
		return getViewInstance(instanceName)
	}).(*gview.View)
}

func getViewInstance(name ...string) *gview.View {
	var (
		err          error
		ctx          = context.Background()
		instanceName = gview.DefaultName
	)
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	view := gview.Instance(instanceName)
	if Config().Available(ctx) {
		var (
			configMap      map[string]interface{}
			configNodeName = consts.ConfigNodeNameViewer
		)
		if configMap, err = Config().Data(ctx); err != nil {
			intlog.Errorf(ctx, `retrieve config data map failed: %+v`, err)
		}
		if len(configMap) > 0 {
			if v, _ := gutil.MapPossibleItemByKey(configMap, consts.ConfigNodeNameViewer); v != "" {
				configNodeName = v
			}
		}
		configMap = Config().MustGet(ctx, fmt.Sprintf(`%s.%s`, configNodeName, instanceName)).Map()
		if len(configMap) == 0 {
			configMap = Config().MustGet(ctx, configNodeName).Map()
		}
		if len(configMap) > 0 {
			if err = view.SetConfigWithMap(configMap); err != nil {
				panic(err)
			}
		}
	}
	return view
}
