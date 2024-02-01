// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"net/http"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/util/gconv"
	)
// SetQuery 用于设置自定义查询值，通过键值对的方式。
func (r *Request) SetQuery(key string, value interface{}) {
	r.parseQuery()
	if r.queryMap == nil {
		r.queryMap = make(map[string]interface{})
	}
	r.queryMap[key] = value
}

// GetQuery 从查询字符串和请求体中获取并返回指定名称`key`的参数。如果`key`在查询中不存在且提供了`def`，则返回`def`；否则返回nil。
//
// 注意，如果有多个同名参数，将以优先级顺序获取并覆盖：query > body。
// 从GET方式传递过来的参数，包括Query String及Body参数解析。
func (r *Request) GetQuery(key string, def ...interface{}) *gvar.Var {
	r.parseQuery()
	if len(r.queryMap) > 0 {
		if value, ok := r.queryMap[key]; ok {
			return gvar.New(value)
		}
	}
	if r.Method == http.MethodGet {
		r.parseBody()
	}
	if len(r.bodyMap) > 0 {
		if v, ok := r.bodyMap[key]; ok {
			return gvar.New(v)
		}
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}

// GetQueryMap 从客户端通过 HTTP GET 方法传递的所有参数中获取并以 map 形式返回。参数 `kvMap` 指定了要从客户端参数中检索的键，关联的值是如果客户端未传递时的默认值。
//
// 注意，如果有多个同名参数，则按照 query > body 的优先级顺序获取并覆盖这些参数。
// 从GET方式传递过来的参数，包括Query String及Body参数解析。
func (r *Request) GetQueryMap(kvMap ...map[string]interface{}) map[string]interface{} {
	r.parseQuery()
	if r.Method == http.MethodGet {
		r.parseBody()
	}
	var m map[string]interface{}
	if len(kvMap) > 0 && kvMap[0] != nil {
		if len(r.queryMap) == 0 && len(r.bodyMap) == 0 {
			return kvMap[0]
		}
		m = make(map[string]interface{}, len(kvMap[0]))
		if len(r.bodyMap) > 0 {
			for k, v := range kvMap[0] {
				if postValue, ok := r.bodyMap[k]; ok {
					m[k] = postValue
				} else {
					m[k] = v
				}
			}
		}
		if len(r.queryMap) > 0 {
			for k, v := range kvMap[0] {
				if postValue, ok := r.queryMap[k]; ok {
					m[k] = postValue
				} else {
					m[k] = v
				}
			}
		}
	} else {
		m = make(map[string]interface{}, len(r.queryMap)+len(r.bodyMap))
		for k, v := range r.bodyMap {
			m[k] = v
		}
		for k, v := range r.queryMap {
			m[k] = v
		}
	}
	return m
}

// GetQueryMapStrStr 从客户端通过HTTP GET方法获取并返回所有传递的参数，以
// map[string]string 的形式。参数 `kvMap` 指定了要从客户端参数中检索的键，
// 关联的值是如果客户端未传递时的默认值。
// 从GET方式传递过来的参数，包括Query String及Body参数解析。
func (r *Request) GetQueryMapStrStr(kvMap ...map[string]interface{}) map[string]string {
	queryMap := r.GetQueryMap(kvMap...)
	if len(queryMap) > 0 {
		m := make(map[string]string, len(queryMap))
		for k, v := range queryMap {
			m[k] = gconv.String(v)
		}
		return m
	}
	return nil
}

// GetQueryMapStrVar 从客户端通过HTTP GET方法获取并返回所有传递的参数，以map[string]*gvar.Var的形式。参数`kvMap`指定了从客户端参数中检索的键，如果客户端未传递，则关联值为默认值。
// 从GET方式传递过来的参数，包括Query String及Body参数解析。
func (r *Request) GetQueryMapStrVar(kvMap ...map[string]interface{}) map[string]*gvar.Var {
	queryMap := r.GetQueryMap(kvMap...)
	if len(queryMap) > 0 {
		m := make(map[string]*gvar.Var, len(queryMap))
		for k, v := range queryMap {
			m[k] = gvar.New(v)
		}
		return m
	}
	return nil
}

// GetQueryStruct 通过HTTP GET方法获取客户端传递的所有参数，并将它们转换为给定的结构体对象。注意，参数`pointer`是指向该结构体对象的指针。
// 可选参数`mapping`用于指定键到属性的映射关系。
// 从GET方式传递过来的参数，包括Query String及Body参数解析。
func (r *Request) GetQueryStruct(pointer interface{}, mapping ...map[string]string) error {
	_, err := r.doGetQueryStruct(pointer, mapping...)
	return err
}

func (r *Request) doGetQueryStruct(pointer interface{}, mapping ...map[string]string) (data map[string]interface{}, err error) {
	r.parseQuery()
	data = r.GetQueryMap()
	if data == nil {
		data = map[string]interface{}{}
	}
	if err = r.mergeDefaultStructValue(data, pointer); err != nil {
		return data, nil
	}
	return data, gconv.Struct(data, pointer, mapping...)
}
