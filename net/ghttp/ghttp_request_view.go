// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import "github.com/gogf/gf/v2/os/gview"

// SetView 为当前请求设置模板视图引擎对象。 md5:ec41ed91daaf7bd3
func (r *Request) SetView(view *gview.View) {
	r.viewObject = view
}

// GetView 返回此请求的模板视图引擎对象。 md5:1eb5934f5359a959
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

// Assigns 将多个模板变量绑定到当前请求。 md5:9676a02a50e91095
func (r *Request) Assigns(data gview.Params) {
	if r.viewParams == nil {
		r.viewParams = make(gview.Params, len(data))
	}
	for k, v := range data {
		r.viewParams[k] = v
	}
}

// Assign 将模板变量绑定到当前请求。 md5:0a82d7a20f0d7265
func (r *Request) Assign(key string, value interface{}) {
	if r.viewParams == nil {
		r.viewParams = make(gview.Params)
	}
	r.viewParams[key] = value
}
