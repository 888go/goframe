// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"github.com/888go/goframe/os/gview"
)

// SetView 设置模板视图引擎对象，用于当前请求。
func (r *X请求) X设置模板对象(模板对象 *模板类.View) {
	r.viewObject = 模板对象
}

// GetView 返回当前请求的模板视图引擎对象。
func (r *X请求) X取模板对象() *模板类.View {
	view := r.viewObject
	if view == nil {
		view = r.X服务.config.X模板默认
	}
	if view == nil {
		view = 模板类.Instance()
	}
	return view
}

// Assigns 将多个模板变量绑定到当前请求。
func (r *X请求) X绑定模板变量Map(Map值 模板类.Params) {
	if r.viewParams == nil {
		r.viewParams = make(模板类.Params, len(Map值))
	}
	for k, v := range Map值 {
		r.viewParams[k] = v
	}
}

// Assign 将模板变量绑定到当前请求。
func (r *X请求) X绑定模板变量(名称 string, 值 interface{}) {
	if r.viewParams == nil {
		r.viewParams = make(模板类.Params)
	}
	r.viewParams[名称] = 值
}
