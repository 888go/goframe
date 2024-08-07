// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"net/http"

	gvar "github.com/888go/goframe/container/gvar"
	gconv "github.com/888go/goframe/util/gconv"
)

// X设置查询参数 使用键值对设置自定义查询值。 md5:464e6b634ef97c90
func (r *Request) X设置查询参数(名称 string, 值 interface{}) {
	r.parseQuery()
	if r.queryMap == nil {
		r.queryMap = make(map[string]interface{})
	}
	r.queryMap[名称] = 值
}

// X取查询参数到泛型类 从查询字符串和请求体中检索并返回给定名称 `key` 的参数。如果 `key` 不在查询中并且提供了 `def`，则返回 `def`；否则返回 nil。
// 
// 注意，如果有多个同名的参数，将按照优先级顺序进行检索和覆盖：查询参数 > 身体参数。
// md5:3948868b7e507e93
func (r *Request) X取查询参数到泛型类(名称 string, 默认值 ...interface{}) *gvar.Var {
	r.parseQuery()
	if len(r.queryMap) > 0 {
		if value, ok := r.queryMap[名称]; ok {
			return gvar.X创建(value)
		}
	}
	if r.Method == http.MethodGet {
		r.parseBody()
	}
	if len(r.bodyMap) > 0 {
		if v, ok := r.bodyMap[名称]; ok {
			return gvar.X创建(v)
		}
	}
	if len(默认值) > 0 {
		return gvar.X创建(默认值[0])
	}
	return nil
}

// X取查询参数到Map 从客户端通过HTTP GET方法传递的所有参数中检索并返回它们作为映射。参数 `kvMap` 指定了从客户端参数中获取的键，如果客户端未提供，则关联的值为默认值。
// 
// 注意，如果有多个具有相同名称的参数，将按照优先级顺序检索和覆盖：查询参数 > 身体（请求体）参数。
// md5:72471cd6457be5f2
func (r *Request) X取查询参数到Map(kvMap ...map[string]interface{}) map[string]interface{} {
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

// X取查询参数到MapStrStr 获取并返回所有通过HTTP GET方法从客户端传递过来的参数，作为一个
//
// map[string]string。参数 `kvMap` 指定了从客户端参数中提取的键
//
// ，关联的值是如果客户端没有传递时的默认值。
// md5:b1d5d46b8cc53f3a
func (r *Request) X取查询参数到MapStrStr(kvMap ...map[string]interface{}) map[string]string {
	queryMap := r.X取查询参数到Map(kvMap...)
	if len(queryMap) > 0 {
		m := make(map[string]string, len(queryMap))
		for k, v := range queryMap {
			m[k] = gconv.String(v)
		}
		return m
	}
	return nil
}

// X取查询参数到Map泛型类切片 从使用 HTTP GET 方法传递的客户端参数中获取并返回所有参数，形式为 map[string]*gvar.Var。参数 `kvMap` 指定了要从客户端参数中获取的键，对应的值是如果客户端未传递时的默认值。
// md5:3db7496b4b165e99
func (r *Request) X取查询参数到Map泛型类切片(kvMap ...map[string]interface{}) map[string]*gvar.Var {
	queryMap := r.X取查询参数到Map(kvMap...)
	if len(queryMap) > 0 {
		m := make(map[string]*gvar.Var, len(queryMap))
		for k, v := range queryMap {
			m[k] = gvar.X创建(v)
		}
		return m
	}
	return nil
}

// X取查询参数到结构体 从客户端通过HTTP GET方法获取所有传递的参数，并将它们转换为给定的结构体对象。请注意，参数`pointer`是指向结构体对象的指针。可选参数`mapping`用于指定键到属性的映射。
// md5:7061a83f935b7317
func (r *Request) X取查询参数到结构体(结构体指针 interface{}, mapping ...map[string]string) error {
	_, err := r.doGetQueryStruct(结构体指针, mapping...)
	return err
}

func (r *Request) doGetQueryStruct(pointer interface{}, mapping ...map[string]string) (data map[string]interface{}, err error) {
	r.parseQuery()
	data = r.X取查询参数到Map()
	if data == nil {
		data = map[string]interface{}{}
	}
	if err = r.mergeDefaultStructValue(data, pointer); err != nil {
		return data, nil
	}
	return data, gconv.Struct(data, pointer, mapping...)
}
