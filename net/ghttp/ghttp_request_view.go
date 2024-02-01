// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"github.com/888go/goframe/os/gview"
	)
// SetView 设置模板视图引擎对象，用于当前请求。
func (r *Request) SetView(view *gview.View) {
	r.viewObject = view
}

// GetView 返回当前请求的模板视图引擎对象。
func (r *Request) GetView() *gview.View {
	view := r.viewObject
	if view == nil {
		view = r.Server.config.View
	}
	if view == nil {
		view = gview.Instance()
	}
	return view
}

// Assigns 将多个模板变量绑定到当前请求。
func (r *Request) Assigns(data gview.Params) {
	if r.viewParams == nil {
		r.viewParams = make(gview.Params, len(data))
	}
	for k, v := range data {
		r.viewParams[k] = v
	}
}

// Assign 将模板变量绑定到当前请求。
func (r *Request) Assign(key string, value interface{}) {
	if r.viewParams == nil {
		r.viewParams = make(gview.Params)
	}
	r.viewParams[key] = value
}
