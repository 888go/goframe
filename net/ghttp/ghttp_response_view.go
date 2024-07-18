// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package ghttp

import (
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/gogf/gf/v2/util/gutil"
)

// WriteTpl解析并响应给定的模板文件。参数`params`指定了解析时的模板变量。
// md5:f7af01616060ef2a
// ff:输出到模板文件
// r:
// tpl:模板文件路径
// params:模板变量
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

// WriteTplDefault 函数用于解析并响应默认的模板文件。
// 参数 `params` 用于指定解析模板时所需的变量。
// md5:746b7bfd331d0eb8
// ff:输出到默认模板文件
// r:
// params:模板变量
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
// 参数 `params` 用于指定模板解析时的变量。
// md5:967e05a26da5c949
// ff:输出文本模板
// r:
// content:文本模板
// params:模板变量
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

// ParseTpl 使用给定的模板文件 `tpl` 和模板变量 `params` 进行解析，然后返回解析后的模板内容。
// md5:170f6327b48f33cd
// ff:解析模板文件
// r:
// tpl:模板文件路径
// params:模板变量
func (r *Response) ParseTpl(tpl string, params ...gview.Params) (string, error) {
	return r.Request.GetView().Parse(r.Request.Context(), tpl, r.buildInVars(params...))
}

// ParseTplDefault 使用参数解析默认模板文件。 md5:83eb637c4bdf2659
// ff:解析默认模板文件
// r:
// params:模板变量
func (r *Response) ParseTplDefault(params ...gview.Params) (string, error) {
	return r.Request.GetView().ParseDefault(r.Request.Context(), r.buildInVars(params...))
}

// ParseTplContent 使用给定的模板参数`params`解析指定的模板文件`file`，
// 并返回解析后的模板内容。
// md5:e91c27dd95553a3d
// ff:解析文本模板
// r:
// content:文本模板
// params:模板变量
func (r *Response) ParseTplContent(content string, params ...gview.Params) (string, error) {
	return r.Request.GetView().ParseContent(r.Request.Context(), content, r.buildInVars(params...))
}

// buildInVars 将内置变量合并到 `params` 中，并返回新的模板变量。
// TODO：优化性能。
// md5:c30048db610c3f6d
func (r *Response) buildInVars(params ...map[string]interface{}) map[string]interface{} {
	m := gutil.MapMergeCopy(r.Request.viewParams)
	if len(params) > 0 {
		gutil.MapMerge(m, params[0])
	}
	// 从请求对象中获取自定义模板变量。 md5:dd1c30411d4be18c
	sessionMap := gconv.MapDeep(r.Request.Session.MustData())
	gutil.MapMerge(m, map[string]interface{}{
		"Form":    r.Request.GetFormMap(),
		"Query":   r.Request.GetQueryMap(),
		"Request": r.Request.GetMap(),
		"Cookie":  r.Request.Cookie.Map(),
		"Session": sessionMap,
	})
// 注意，如果没有配置文件，不应将任何Config变量分配给模板。
// md5:83fc6f393ee5e27d
	if v, _ := gcfg.Instance().Data(r.Request.Context()); len(v) > 0 {
		m["Config"] = v
	}
	return m
}
