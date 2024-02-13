// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package http类

import (
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/os/gview"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gmode"
	"github.com/888go/goframe/util/gutil"
)

// WriteTpl 解析并响应给定的模板文件。
// 参数 `params` 指定了用于解析的模板变量。
func (r *Response) X输出到模板文件(模板文件路径 string, 模板变量 ...模板类.Params) error {
	r.Header().Set("Content-Type", contentTypeHtml)
	b, err := r.X解析模板文件(模板文件路径, 模板变量...)
	if err != nil {
		if !环境类.IsProduct() {
			r.X写响应缓冲区("Template Parsing Error: " + err.Error())
		}
		return err
	}
	r.X写响应缓冲区(b)
	return nil
}

// WriteTplDefault 解析并响应默认模板文件。
// 参数`params`用于指定模板解析所需的变量。
func (r *Response) X输出到默认模板文件(模板变量 ...模板类.Params) error {
	r.Header().Set("Content-Type", contentTypeHtml)
	b, err := r.X解析默认模板文件(模板变量...)
	if err != nil {
		if !环境类.IsProduct() {
			r.X写响应缓冲区("Template Parsing Error: " + err.Error())
		}
		return err
	}
	r.X写响应缓冲区(b)
	return nil
}

// WriteTplContent 解析并响应模板内容。
// 参数 `params` 指定了用于解析的模板变量。
func (r *Response) X输出文本模板(文本模板 string, 模板变量 ...模板类.Params) error {
	r.Header().Set("Content-Type", contentTypeHtml)
	b, err := r.X解析文本模板(文本模板, 模板变量...)
	if err != nil {
		if !环境类.IsProduct() {
			r.X写响应缓冲区("Template Parsing Error: " + err.Error())
		}
		return err
	}
	r.X写响应缓冲区(b)
	return nil
}

// ParseTpl 将给定的模板文件 `tpl` 与给定的模板变量 `params` 进行解析，
// 并返回解析后的模板内容。
func (r *Response) X解析模板文件(模板文件路径 string, 模板变量 ...模板类.Params) (string, error) {
	return r.Request.X取模板对象().Parse(r.Request.Context别名(), 模板文件路径, r.buildInVars(模板变量...))
}

// ParseTplDefault 使用参数解析默认模板文件。
func (r *Response) X解析默认模板文件(模板变量 ...模板类.Params) (string, error) {
	return r.Request.X取模板对象().ParseDefault(r.Request.Context别名(), r.buildInVars(模板变量...))
}

// ParseTplContent 函数用于解析给定的模板文件 `file`，并使用给定的模板参数 `params` 进行解析，
// 然后返回解析后的模板内容。
func (r *Response) X解析文本模板(文本模板 string, 模板变量 ...模板类.Params) (string, error) {
	return r.Request.X取模板对象().ParseContent(r.Request.Context别名(), 文本模板, r.buildInVars(模板变量...))
}

// buildInVars 将内置变量合并到 `params` 中，并返回新的模板变量。
// TODO：提升性能。
func (r *Response) buildInVars(params ...map[string]interface{}) map[string]interface{} {
	m := 工具类.MapMergeCopy(r.Request.viewParams)
	if len(params) > 0 {
		工具类.MapMerge(m, params[0])
	}
	// 从请求对象中获取自定义模板变量。
	sessionMap := 转换类.X取Map_递归(r.Request.Session.MustData())
	工具类.MapMerge(m, map[string]interface{}{
		"Form":    r.Request.X取表单值到Map(),
		"Query":   r.Request.X取查询参数到Map(),
		"Request": r.Request.GetMap别名(),
		"Cookie":  r.Request.Cookie.X取Map(),
		"Session": sessionMap,
	})
// 注意，如果没有配置文件，则不应将任何Config变量赋值给模板。
	if v, _ := 配置类.X取单例对象().X取Map(r.Request.Context别名()); len(v) > 0 {
		m["Config"] = v
	}
	return m
}
