// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import "github.com/gogf/gf/v2/os/gview"

// SetView sets template view engine object for this request.
// ff:设置模板对象
// r:
// view:模板对象
func (r *Request) SetView(view *gview.View) {
	r.viewObject = view
}

// GetView returns the template view engine object for this request.
// ff:取模板对象
// r:
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

// Assigns binds multiple template variables to current request.
// ff:绑定模板变量Map
// r:
// data:Map值
func (r *Request) Assigns(data gview.Params) {
	if r.viewParams == nil {
		r.viewParams = make(gview.Params, len(data))
	}
	for k, v := range data {
		r.viewParams[k] = v
	}
}

// Assign binds a template variable to current request.
// ff:绑定模板变量
// r:
// key:名称
// value:值
func (r *Request) Assign(key string, value interface{}) {
	if r.viewParams == nil {
		r.viewParams = make(gview.Params)
	}
	r.viewParams[key] = value
}
