// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	gvar "github.com/888go/goframe/container/gvar"
	gconv "github.com/888go/goframe/util/gconv"
)

// X设置表单值 使用键值对设置自定义表单值。 md5:eca1a8c094c9ff19
func (r *Request) X设置表单值(名称 string, 值 interface{}) {
	r.parseForm()
	if r.formMap == nil {
		r.formMap = make(map[string]interface{})
	}
	r.formMap[名称] = 值
}

// X取表单值到泛型类 从表单中检索并返回键为 `key` 的参数。如果表单中不存在 `key`，并且提供了默认值 `def`，则返回 `def`；否则返回 `nil`。
// md5:f4a13744025f01b8
func (r *Request) X取表单值到泛型类(名称 string, 默认值 ...interface{}) *gvar.Var {
	r.parseForm()
	if len(r.formMap) > 0 {
		if value, ok := r.formMap[名称]; ok {
			return gvar.X创建(value)
		}
	}
	if len(默认值) > 0 {
		return gvar.X创建(默认值[0])
	}
	return nil
}

// X取表单值到Map 从客户端获取并返回所有的表单参数，以map形式。参数`kvMap`指定了从客户端参数中检索的键，如果客户端未传递，则关联的值为默认值。
// md5:bc80893a54c1e60c
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

// X取表单值到MapStrStr 获取并以map[string]string的形式返回客户端传递的所有表单参数。
// 参数 `kvMap` 指定了从客户端参数中提取的键，如果客户端未传递，则关联的值是默认值。
// md5:09a548d91ee42cff
func (r *Request) X取表单值到MapStrStr(kvMap ...map[string]interface{}) map[string]string {
	formMap := r.X取表单值到Map(kvMap...)
	if len(formMap) > 0 {
		m := make(map[string]string, len(formMap))
		for k, v := range formMap {
			m[k] = gconv.String(v)
		}
		return m
	}
	return nil
}

// X取表单值到Map泛型类 从客户端传递的所有表单参数中获取并返回一个 map[string]*gvar.Var。
// 参数 `kvMap` 指定了要从客户端参数中检索的键，对应的值是如果客户端未传递时的默认值。
// md5:0e9cf1899de0705b
func (r *Request) X取表单值到Map泛型类(kvMap ...map[string]interface{}) map[string]*gvar.Var {
	formMap := r.X取表单值到Map(kvMap...)
	if len(formMap) > 0 {
		m := make(map[string]*gvar.Var, len(formMap))
		for k, v := range formMap {
			m[k] = gvar.X创建(v)
		}
		return m
	}
	return nil
}

// X取表单值到结构 从客户端获取所有传递的表单参数，并将它们转换为给定的结构体对象。需要注意的是，参数 `pointer` 是指向结构体对象的指针。可选参数 `mapping` 用于指定键到属性的映射。
// md5:36ac7f24ad6e766e
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
	return data, gconv.Struct(data, pointer, mapping...)
}
