// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/util/gconv"
)

// SetForm 设置自定义表单值，通过键值对的方式。
func (r *Request) X设置表单值(名称 string, 值 interface{}) {
	r.parseForm()
	if r.formMap == nil {
		r.formMap = make(map[string]interface{})
	}
	r.formMap[名称] = 值
}

// GetForm 从表单中检索并返回参数 `key`。
// 如果 `key` 在表单中不存在，且提供了 `def`，则返回 `def`，否则返回 nil。
func (r *Request) X取表单值到泛型类(名称 string, 默认值 ...interface{}) *泛型类.Var {
	r.parseForm()
	if len(r.formMap) > 0 {
		if value, ok := r.formMap[名称]; ok {
			return 泛型类.X创建(value)
		}
	}
	if len(默认值) > 0 {
		return 泛型类.X创建(默认值[0])
	}
	return nil
}

// GetFormMap 从客户端获取并返回所有表单参数，以map形式返回。
// 参数 `kvMap` 指定了从客户端参数中检索的键，如果客户端未传递，则关联的值是默认值。
func (r *Request) X取表单值到Map(kvMap ...map[string]interface{}) map[string]interface{} {
	r.parseForm()
	if len(kvMap) > 0 && kvMap[0] != nil {
		if len(r.formMap) == 0 {
			return kvMap[0]
		}
		m := make(map[string]interface{}, len(kvMap[0]))
		for k, defValue := range kvMap[0] {
			if postValue, ok := r.formMap[k]; ok {
				m[k] = postValue
			} else {
				m[k] = defValue
			}
		}
		return m
	} else {
		return r.formMap
	}
}

// GetFormMapStrStr 从客户端获取并返回所有以 map[string]string 形式传递的表单参数。
// 参数 `kvMap` 指定了从客户端参数中检索的键，如果客户端未传递，则关联的值是默认值。
func (r *Request) X取表单值到MapStrStr(kvMap ...map[string]interface{}) map[string]string {
	formMap := r.X取表单值到Map(kvMap...)
	if len(formMap) > 0 {
		m := make(map[string]string, len(formMap))
		for k, v := range formMap {
			m[k] = 转换类.String(v)
		}
		return m
	}
	return nil
}

// GetFormMapStrVar 从客户端获取并返回所有以map[string]*gvar.Var形式传递的表单参数。
// 参数`kvMap`指定了从客户端参数中检索的键，如果客户端未传递，则关联的值是默认值。
func (r *Request) X取表单值到Map泛型类(kvMap ...map[string]interface{}) map[string]*泛型类.Var {
	formMap := r.X取表单值到Map(kvMap...)
	if len(formMap) > 0 {
		m := make(map[string]*泛型类.Var, len(formMap))
		for k, v := range formMap {
			m[k] = 泛型类.X创建(v)
		}
		return m
	}
	return nil
}

// GetFormStruct 从客户端获取所有传递的表单参数，并将其转换为给定的结构体对象。
// 注意，参数 `pointer` 是指向结构体对象的指针。
// 可选参数 `mapping` 用于指定键到属性的映射关系。
func (r *Request) X取表单值到结构(结构指针 interface{}, mapping ...map[string]string) error {
	_, err := r.doGetFormStruct(结构指针, mapping...)
	return err
}

func (r *Request) doGetFormStruct(pointer interface{}, mapping ...map[string]string) (data map[string]interface{}, err error) {
	r.parseForm()
	data = r.formMap
	if data == nil {
		data = map[string]interface{}{}
	}
	if err = r.mergeDefaultStructValue(data, pointer); err != nil {
		return data, nil
	}
	return data, 转换类.Struct(data, pointer, mapping...)
}
