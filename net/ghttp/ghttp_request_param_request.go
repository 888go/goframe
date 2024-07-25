// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// GetRequest 从客户端获取并返回名为 `key` 的参数，以及作为接口传递的自定义参数。无论客户端使用何种HTTP方法。如果`key`不存在，`def`参数指定了默认值。
// 
// GetRequest 是最常用的用于检索参数的函数之一。
// 
// 注意，如果有多个同名参数，将按照以下优先级顺序进行获取和覆盖：路由器 < 查询参数 < 身份验证 < 表单数据 < 自定义参数。
// md5:a008e7f428967448
func (r *Request) GetRequest(key string, def ...interface{}) *gvar.Var {
	value := r.GetParam(key)
	if value.IsNil() {
		value = r.GetForm(key)
	}
	if value.IsNil() {
		r.parseBody()
		if len(r.bodyMap) > 0 {
			if v := r.bodyMap[key]; v != nil {
				value = gvar.New(v)
			}
		}
	}
	if value.IsNil() {
		value = r.GetQuery(key)
	}
	if value.IsNil() {
		value = r.GetRouter(key)
	}
	if !value.IsNil() {
		return value
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}

// GetRequestMap 从客户端获取并返回所有传递的参数及自定义参数，无论客户端使用的是哪种HTTP方法。参数 `kvMap` 指定了从客户端参数中提取的键，关联的值是在客户端未传递相应键时的默认值。
//
// GetRequestMap 是最常用于检索参数的函数之一。
//
// 注意，如果有多个同名参数，参数将按照优先级顺序被获取及覆盖：路由参数 < 查询参数 < 请求体参数 < 表单参数 < 自定义参数。
// md5:b01ba4caf2092f12
func (r *Request) GetRequestMap(kvMap ...map[string]interface{}) map[string]interface{} {
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
			if uploadFiles := r.GetUploadFiles(name); len(uploadFiles) == 1 {
				m[name] = uploadFiles[0]
			} else {
				m[name] = uploadFiles
			}
		}
	}
		// 检查不存在的参数并为其分配默认值。 md5:2c9c16dac85c432c
	if filter {
		for k, v := range kvMap[0] {
			if _, ok = m[k]; !ok {
				m[k] = v
			}
		}
	}
	return m
}

// GetRequestMapStrStr 从客户端和自定义参数中获取并返回所有传递的参数，无论客户端使用何种HTTP方法。参数`kvMap`指定了从客户端参数中检索的键，关联的值是客户端未传递时的默认值。
// md5:18e353330403d45b
func (r *Request) GetRequestMapStrStr(kvMap ...map[string]interface{}) map[string]string {
	requestMap := r.GetRequestMap(kvMap...)
	if len(requestMap) > 0 {
		m := make(map[string]string, len(requestMap))
		for k, v := range requestMap {
			m[k] = gconv.String(v)
		}
		return m
	}
	return nil
}

// GetRequestMapStrVar 从客户端和自定义参数中检索并返回所有传递的参数，作为map[string]*gvar.Var。无论客户端使用何种HTTP方法，都会进行检索。参数`kvMap`指定了从客户端参数中获取的键，关联的值是客户端未传递时的默认值。
// md5:1063c291381a5048
func (r *Request) GetRequestMapStrVar(kvMap ...map[string]interface{}) map[string]*gvar.Var {
	requestMap := r.GetRequestMap(kvMap...)
	if len(requestMap) > 0 {
		m := make(map[string]*gvar.Var, len(requestMap))
		for k, v := range requestMap {
			m[k] = gvar.New(v)
		}
		return m
	}
	return nil
}

// GetRequestStruct 无论客户端使用何种HTTP方法，都会获取客户端传递的所有参数和自定义参数，
// 并将它们转换为结构体对象。注意，参数`pointer`是一个指向结构体对象的指针。
// 可选参数`mapping`用于指定键到属性的映射。
// md5:a117b2c0722fc3fe
func (r *Request) GetRequestStruct(pointer interface{}, mapping ...map[string]string) error {
	_, err := r.doGetRequestStruct(pointer, mapping...)
	return err
}

func (r *Request) doGetRequestStruct(pointer interface{}, mapping ...map[string]string) (data map[string]interface{}, err error) {
	data = r.GetRequestMap()
	if data == nil {
		data = map[string]interface{}{}
	}
	// Default struct values.
	if err = r.mergeDefaultStructValue(data, pointer); err != nil {
		return data, nil
	}
		// `in` 标签结构体值。 md5:225b15f233b09df1
	if err = r.mergeInTagStructValue(data, pointer); err != nil {
		return data, nil
	}

	return data, gconv.Struct(data, pointer, mapping...)
}

// mergeDefaultStructValue 将请求参数与结构体标签定义中的默认值合并。 md5:0a73ebb7f647201a
func (r *Request) mergeDefaultStructValue(data map[string]interface{}, pointer interface{}) error {
	fields := r.serveHandler.Handler.Info.ReqStructFields
	if len(fields) > 0 {
		var (
			foundKey   string
			foundValue interface{}
		)
		for _, field := range fields {
			if tagValue := field.TagDefault(); tagValue != "" {
				foundKey, foundValue = gutil.MapPossibleItemByKey(data, field.Name())
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

			// 提供非严格的路由. md5:c3f73d5de1159867
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
			foundKey, foundValue = gutil.MapPossibleItemByKey(data, field.Name())
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

// mergeInTagStructValue 将请求参数与根据结构体`in`标签定义的头或cookie值合并。 md5:a6444655a59f403d
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
					foundHeaderKey, foundHeaderValue := gutil.MapPossibleItemByKey(headerMap, field.TagPriorityName())
					if foundHeaderKey != "" {
						foundKey, foundValue = gutil.MapPossibleItemByKey(data, foundHeaderKey)
						if foundKey == "" {
							data[field.Name()] = foundHeaderValue
						} else {
							if empty.IsEmpty(foundValue) {
								data[foundKey] = foundHeaderValue
							}
						}
					}
				case goai.ParameterInCookie:
					foundCookieKey, foundCookieValue := gutil.MapPossibleItemByKey(cookieMap, field.TagPriorityName())
					if foundCookieKey != "" {
						foundKey, foundValue = gutil.MapPossibleItemByKey(data, foundCookieKey)
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
