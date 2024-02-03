// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package ghttp

import (
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/os/gview"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gmode"
	"github.com/888go/goframe/util/gutil"
)

// WriteTpl 解析并响应给定的模板文件。
// 参数 `params` 指定了用于解析的模板变量。
func (r *Response) WriteTpl(tpl string, params ...gview.Params) error {
	r.Header().Set("Content-Type", contentTypeHtml)
	b, err := r.ParseTpl(tpl, params...)
	if err != nil {
		if !gmode.IsProduct() {
			r.Write("Template Parsing Error: " + err.Error())
		}
		return err
	}
	r.Write(b)
	return nil
}

// WriteTplDefault 解析并响应默认模板文件。
// 参数`params`用于指定模板解析所需的变量。
func (r *Response) WriteTplDefault(params ...gview.Params) error {
	r.Header().Set("Content-Type", contentTypeHtml)
	b, err := r.ParseTplDefault(params...)
	if err != nil {
		if !gmode.IsProduct() {
			r.Write("Template Parsing Error: " + err.Error())
		}
		return err
	}
	r.Write(b)
	return nil
}

// WriteTplContent 解析并响应模板内容。
// 参数 `params` 指定了用于解析的模板变量。
func (r *Response) WriteTplContent(content string, params ...gview.Params) error {
	r.Header().Set("Content-Type", contentTypeHtml)
	b, err := r.ParseTplContent(content, params...)
	if err != nil {
		if !gmode.IsProduct() {
			r.Write("Template Parsing Error: " + err.Error())
		}
		return err
	}
	r.Write(b)
	return nil
}

// ParseTpl 将给定的模板文件 `tpl` 与给定的模板变量 `params` 进行解析，
// 并返回解析后的模板内容。
func (r *Response) ParseTpl(tpl string, params ...gview.Params) (string, error) {
	return r.Request.GetView().Parse(r.Request.Context(), tpl, r.buildInVars(params...))
}

// ParseTplDefault 使用参数解析默认模板文件。
func (r *Response) ParseTplDefault(params ...gview.Params) (string, error) {
	return r.Request.GetView().ParseDefault(r.Request.Context(), r.buildInVars(params...))
}

// ParseTplContent 函数用于解析给定的模板文件 `file`，并使用给定的模板参数 `params` 进行解析，
// 然后返回解析后的模板内容。
func (r *Response) ParseTplContent(content string, params ...gview.Params) (string, error) {
	return r.Request.GetView().ParseContent(r.Request.Context(), content, r.buildInVars(params...))
}

// buildInVars 将内置变量合并到 `params` 中，并返回新的模板变量。
// TODO：提升性能。
func (r *Response) buildInVars(params ...map[string]interface{}) map[string]interface{} {
	m := gutil.MapMergeCopy(r.Request.viewParams)
	if len(params) > 0 {
		gutil.MapMerge(m, params[0])
	}
	// 从请求对象中获取自定义模板变量。
	sessionMap := gconv.MapDeep(r.Request.Session.MustData())
	gutil.MapMerge(m, map[string]interface{}{
		"Form":    r.Request.GetFormMap(),
		"Query":   r.Request.GetQueryMap(),
		"Request": r.Request.GetMap(),
		"Cookie":  r.Request.Cookie.Map(),
		"Session": sessionMap,
	})
// 注意，如果没有配置文件，则不应将任何Config变量赋值给模板。
	if v, _ := gcfg.Instance().Data(r.Request.Context()); len(v) > 0 {
		m["Config"] = v
	}
	return m
}
