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
	gview "github.com/888go/goframe/os/gview"
	gutil "github.com/888go/goframe/util/gutil"
)

// View 返回一个具有默认设置的View实例。
// 参数 `name` 是实例的名称。
// 请注意，如果在创建实例期间发生任何错误，它将引发 panic。
// md5:5f91ad34b8070939
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
	if Config().X是否可用(ctx) {
		var (
			configMap      map[string]interface{}
			configNodeName = consts.ConfigNodeNameViewer
		)
		if configMap, err = Config().X取Map(ctx); err != nil {
			intlog.Errorf(ctx, `retrieve config data map failed: %+v`, err)
		}
		if len(configMap) > 0 {
			if v, _ := gutil.MapPossibleItemByKey(configMap, consts.ConfigNodeNameViewer); v != "" {
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
