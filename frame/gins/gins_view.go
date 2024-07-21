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
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/util/gutil"
)

		// View 返回一个具有默认设置的View实例。
		// 参数 `name` 是实例的名称。
		// 请注意，如果在创建实例期间发生任何错误，它将引发 panic。
		// md5:5f91ad34b8070939
// ff:
// name:
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
