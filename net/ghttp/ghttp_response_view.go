// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package http类

import (
	gcfg "github.com/888go/goframe/os/gcfg"
	gview "github.com/888go/goframe/os/gview"
	gconv "github.com/888go/goframe/util/gconv"
	gmode "github.com/888go/goframe/util/gmode"
	gutil "github.com/888go/goframe/util/gutil"
)

// X输出到模板文件解析并响应给定的模板文件。参数`params`指定了解析时的模板变量。
// md5:f7af01616060ef2a
func (r *Response) X输出到模板文件(模板文件路径 string, 模板变量 ...gview.Params) error {
	r.Header().Set("Content-Type", contentTypeHtml)
	b, err := r.X解析模板文件(模板文件路径, 模板变量...)
	if err != nil {
		if !gmode.IsProduct() {
			r.X写响应缓冲区("Template Parsing Error: " + err.Error())
		}
		return err
	}
	r.X写响应缓冲区(b)
	return nil
}

// X输出到默认模板文件 函数用于解析并响应默认的模板文件。
// 参数 `params` 用于指定解析模板时所需的变量。
// md5:746b7bfd331d0eb8
func (r *Response) X输出到默认模板文件(模板变量 ...gview.Params) error {
	r.Header().Set("Content-Type", contentTypeHtml)
	b, err := r.X解析默认模板文件(模板变量...)
	if err != nil {
		if !gmode.IsProduct() {
			r.X写响应缓冲区("Template Parsing Error: " + err.Error())
		}
		return err
	}
	r.X写响应缓冲区(b)
	return nil
}

// X输出文本模板 解析并响应模板内容。
// 参数 `params` 用于指定模板解析时的变量。
// md5:967e05a26da5c949
func (r *Response) X输出文本模板(文本模板 string, 模板变量 ...gview.Params) error {
	r.Header().Set("Content-Type", contentTypeHtml)
	b, err := r.X解析文本模板(文本模板, 模板变量...)
	if err != nil {
		if !gmode.IsProduct() {
			r.X写响应缓冲区("Template Parsing Error: " + err.Error())
		}
		return err
	}
	r.X写响应缓冲区(b)
	return nil
}

// X解析模板文件 使用给定的模板文件 `tpl` 和模板变量 `params` 进行解析，然后返回解析后的模板内容。
// md5:170f6327b48f33cd
func (r *Response) X解析模板文件(模板文件路径 string, 模板变量 ...gview.Params) (string, error) {
	return r.Request.X取模板对象().Parse(r.Request.Context别名(), 模板文件路径, r.buildInVars(模板变量...))
}

// X解析默认模板文件 使用参数解析默认模板文件。 md5:83eb637c4bdf2659
func (r *Response) X解析默认模板文件(模板变量 ...gview.Params) (string, error) {
	return r.Request.X取模板对象().ParseDefault(r.Request.Context别名(), r.buildInVars(模板变量...))
}

// X解析文本模板 使用给定的模板参数`params`解析指定的模板文件`file`，
// 并返回解析后的模板内容。
// md5:e91c27dd95553a3d
func (r *Response) X解析文本模板(文本模板 string, 模板变量 ...gview.Params) (string, error) {
	return r.Request.X取模板对象().ParseContent(r.Request.Context别名(), 文本模板, r.buildInVars(模板变量...))
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
	sessionMap := gconv.X取Map_递归(r.Request.Session.MustData())
	gutil.MapMerge(m, map[string]interface{}{
		"Form":    r.Request.X取表单值到Map(),
		"Query":   r.Request.X取查询参数到Map(),
		"Request": r.Request.GetMap别名(),
		"Cookie":  r.Request.Cookie.X取Map(),
		"Session": sessionMap,
	})
	// 注意，如果没有配置文件，不应将任何Config变量分配给模板。
	// md5:83fc6f393ee5e27d
	if v, _ := gcfg.X取单例对象().X取Map(r.Request.Context别名()); len(v) > 0 {
		m["Config"] = v
	}
	return m
}
