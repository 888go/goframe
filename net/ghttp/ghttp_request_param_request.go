// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/net/goai"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

// GetRequest 函数用于获取并返回客户端通过任意HTTP方法传递的名为`key`的参数以及作为interface{}类型的自定义参数。参数`def`用于指定当`key`不存在时的默认值。
//
// GetRequest 是用于检索参数的最常用函数之一。
//
// 注意，可获取客户端提交的所有参数，不区分提交方式。如果有多个同名参数，按照以下优先级顺序获取并覆盖：路由参数 < 查询参数 < 请求体参数 < 表单参数 < 自定义参数。
func (r *Request) X取参数(名称 string, 默认 ...interface{}) *泛型类.Var {
	value := r.X取自定义参数到泛型类(名称)
	if value.X是否为Nil() {
		value = r.X取表单值到泛型类(名称)
	}
	if value.X是否为Nil() {
		r.parseBody()
		if len(r.bodyMap) > 0 {
			if v := r.bodyMap[名称]; v != nil {
				value = 泛型类.X创建(v)
			}
		}
	}
	if value.X是否为Nil() {
		value = r.X取查询参数到泛型类(名称)
	}
	if value.X是否为Nil() {
		value = r.X取路由器值到泛型类(名称)
	}
	if !value.X是否为Nil() {
		return value
	}
	if len(默认) > 0 {
		return 泛型类.X创建(默认[0])
	}
	return nil
}

// GetRequestMap 从客户端获取并返回所有传递的参数以及自定义参数，无论客户端使用何种HTTP方法。参数`kvMap`指定了从客户端参数中检索的键，关联的值是如果客户端未传递相应键时的默认值。
//
// GetRequestMap 是用于检索参数的最常用函数之一。
//
// 注意，可获取客户端提交的所有参数，不区分提交方式。如果有多个同名参数，按照以下优先级顺序获取并覆盖：路由参数 < 查询参数 < 请求体参数 < 表单参数 < 自定义参数。
func (r *Request) X取参数到Map(kvMap ...map[string]interface{}) map[string]interface{} {
	r.parseQuery()
	r.parseForm()
	r.parseBody()
	var (
		ok, filter bool
	)
	if len(kvMap) > 0 && kvMap[0] != nil {
		filter = true
	}
	m := make(map[string]interface{})
	for k, v := range r.routerMap {
		if filter {
			if _, ok = kvMap[0][k]; !ok {
				continue
			}
		}
		m[k] = v
	}
	for k, v := range r.queryMap {
		if filter {
			if _, ok = kvMap[0][k]; !ok {
				continue
			}
		}
		m[k] = v
	}
	for k, v := range r.formMap {
		if filter {
			if _, ok = kvMap[0][k]; !ok {
				continue
			}
		}
		m[k] = v
	}
	for k, v := range r.bodyMap {
		if filter {
			if _, ok = kvMap[0][k]; !ok {
				continue
			}
		}
		m[k] = v
	}
	for k, v := range r.paramsMap {
		if filter {
			if _, ok = kvMap[0][k]; !ok {
				continue
			}
		}
		m[k] = v
	}
	// File uploading.
	if r.MultipartForm != nil {
		for name := range r.MultipartForm.File {
			if uploadFiles := r.X取上传文件数组对象(name); len(uploadFiles) == 1 {
				m[name] = uploadFiles[0]
			} else {
				m[name] = uploadFiles
			}
		}
	}
	// 检查不存在的参数，并赋予其默认值。
	if filter {
		for k, v := range kvMap[0] {
			if _, ok = m[k]; !ok {
				m[k] = v
			}
		}
	}
	return m
}

// GetRequestMapStrStr 从客户端获取并返回所有传递的参数以及自定义参数，无论客户端使用何种HTTP方法。
// 参数`kvMap`指定了从客户端参数中检索的键，如果客户端未传递，则关联的值是默认值。返回类型为map[string]string。
// 注意，可获取客户端提交的所有参数，不区分提交方式。如果有多个同名参数，按照以下优先级顺序获取并覆盖：路由参数 < 查询参数 < 请求体参数 < 表单参数 < 自定义参数。
func (r *Request) X取参数到MapStrStr(kvMap ...map[string]interface{}) map[string]string {
	requestMap := r.X取参数到Map(kvMap...)
	if len(requestMap) > 0 {
		m := make(map[string]string, len(requestMap))
		for k, v := range requestMap {
			m[k] = 转换类.String(v)
		}
		return m
	}
	return nil
}

// GetRequestMapStrVar 获取并返回客户端通过任何HTTP方法传递的所有参数，以及自定义参数，
// 并以map[string]*gvar.Var的形式返回。参数`kvMap`指定了从客户端参数中获取的键，
// 相关联的值是当客户端未传递时的默认值。
// 注意，可获取客户端提交的所有参数，不区分提交方式。如果有多个同名参数，按照以下优先级顺序获取并覆盖：路由参数 < 查询参数 < 请求体参数 < 表单参数 < 自定义参数。
func (r *Request) X取参数到Map泛型类(kvMap ...map[string]interface{}) map[string]*泛型类.Var {
	requestMap := r.X取参数到Map(kvMap...)
	if len(requestMap) > 0 {
		m := make(map[string]*泛型类.Var, len(requestMap))
		for k, v := range requestMap {
			m[k] = 泛型类.X创建(v)
		}
		return m
	}
	return nil
}

// GetRequestStruct 从客户端获取所有传递的参数以及自定义参数，无论客户端使用何种HTTP方法，
// 并将它们转换为给定的结构体对象。注意，参数`pointer`是指向结构体对象的指针。
// 可选参数`mapping`用于指定键到属性的映射关系。
// 注意，可获取客户端提交的所有参数，不区分提交方式。如果有多个同名参数，按照以下优先级顺序获取并覆盖：路由参数 < 查询参数 < 请求体参数 < 表单参数 < 自定义参数。
func (r *Request) X取参数到结构体(结构体指针 interface{}, 名称映射 ...map[string]string) error {
	_, err := r.doGetRequestStruct(结构体指针, 名称映射...)
	return err
}

func (r *Request) doGetRequestStruct(pointer interface{}, mapping ...map[string]string) (data map[string]interface{}, err error) {
	data = r.X取参数到Map()
	if data == nil {
		data = map[string]interface{}{}
	}
	// 默认结构体值。
	if err = r.mergeDefaultStructValue(data, pointer); err != nil {
		return data, nil
	}
	// `in` 标签结构体值。
	if err = r.mergeInTagStructValue(data, pointer); err != nil {
		return data, nil
	}
	return data, 转换类.Struct(data, pointer, mapping...)
}

// mergeDefaultStructValue 将请求参数与来自结构体标签定义的默认值进行合并。
func (r *Request) mergeDefaultStructValue(data map[string]interface{}, pointer interface{}) error {
	fields := r.serveHandler.Handler.Info.ReqStructFields
	if len(fields) > 0 {
		var (
			foundKey   string
			foundValue interface{}
		)
		for _, field := range fields {
			if tagValue := field.TagDefault(); tagValue != "" {
				foundKey, foundValue = 工具类.MapPossibleItemByKey(data, field.Name())
				if foundKey == "" {
					data[field.Name()] = tagValue
				} else {
					if empty.IsEmpty(foundValue) {
						data[foundKey] = tagValue
					}
				}
			}
		}
		return nil
	}

	// 提供非严格路由
	tagFields, err := gstructs.TagFields(pointer, defaultValueTags)
	if err != nil {
		return err
	}
	if len(tagFields) > 0 {
		var (
			foundKey   string
			foundValue interface{}
		)
		for _, field := range tagFields {
			foundKey, foundValue = 工具类.MapPossibleItemByKey(data, field.Name())
			if foundKey == "" {
				data[field.Name()] = field.TagValue
			} else {
				if empty.IsEmpty(foundValue) {
					data[foundKey] = field.TagValue
				}
			}
		}
	}

	return nil
}

// mergeInTagStructValue 将请求参数与来自 `in` 标签定义的结构体中的头部或cookie值进行合并。
func (r *Request) mergeInTagStructValue(data map[string]interface{}, pointer interface{}) error {
	fields := r.serveHandler.Handler.Info.ReqStructFields
	if len(fields) > 0 {
		var (
			foundKey   string
			foundValue interface{}
			headerMap  = make(map[string]interface{})
			cookieMap  = make(map[string]interface{})
		)

		for k, v := range r.Header {
			if len(v) > 0 {
				headerMap[k] = v[0]
			}
		}

		for _, cookie := range r.Cookies() {
			cookieMap[cookie.Name] = cookie.Value
		}

		for _, field := range fields {
			if tagValue := field.TagIn(); tagValue != "" {
				switch tagValue {
				case goai.ParameterInHeader:
					foundHeaderKey, foundHeaderValue := 工具类.MapPossibleItemByKey(headerMap, field.Name())
					if foundHeaderKey != "" {
						foundKey, foundValue = 工具类.MapPossibleItemByKey(data, foundHeaderKey)
						if foundKey == "" {
							data[field.Name()] = foundHeaderValue
						} else {
							if empty.IsEmpty(foundValue) {
								data[foundKey] = foundHeaderValue
							}
						}
					}
				case goai.ParameterInCookie:
					foundCookieKey, foundCookieValue := 工具类.MapPossibleItemByKey(cookieMap, field.Name())
					if foundCookieKey != "" {
						foundKey, foundValue = 工具类.MapPossibleItemByKey(data, foundCookieKey)
						if foundKey == "" {
							data[field.Name()] = foundCookieValue
						} else {
							if empty.IsEmpty(foundValue) {
								data[foundKey] = foundCookieValue
							}
						}
					}
				}
			}
		}
	}
	return nil
}
