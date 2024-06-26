// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/util/gconv"
)

// SetForm sets custom form value with key-value pairs.

// ff:设置表单值
// value:值
// key:名称
func (r *Request) SetForm(key string, value interface{}) {
	r.parseForm()
	if r.formMap == nil {
		r.formMap = make(map[string]interface{})
	}
	r.formMap[key] = value
}

// GetForm retrieves and returns parameter `key` from form.
// It returns `def` if `key` does not exist in the form and `def` is given, or else it returns nil.

// ff:取表单值到泛型类
// def:默认值
// key:名称
func (r *Request) GetForm(key string, def ...interface{}) *gvar.Var {
	r.parseForm()
	if len(r.formMap) > 0 {
		if value, ok := r.formMap[key]; ok {
			return gvar.New(value)
		}
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}

// GetFormMap retrieves and returns all form parameters passed from client as map.
// The parameter `kvMap` specifies the keys retrieving from client parameters,
// the associated values are the default values if the client does not pass.

// ff:取表单值到Map
// kvMap:
func (r *Request) GetFormMap(kvMap ...map[string]interface{}) map[string]interface{} {
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

// GetFormMapStrStr retrieves and returns all form parameters passed from client as map[string]string.
// The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values
// are the default values if the client does not pass.

// ff:取表单值到MapStrStr
// kvMap:
func (r *Request) GetFormMapStrStr(kvMap ...map[string]interface{}) map[string]string {
	formMap := r.GetFormMap(kvMap...)
	if len(formMap) > 0 {
		m := make(map[string]string, len(formMap))
		for k, v := range formMap {
			m[k] = gconv.String(v)
		}
		return m
	}
	return nil
}

// GetFormMapStrVar retrieves and returns all form parameters passed from client as map[string]*gvar.Var.
// The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values
// are the default values if the client does not pass.

// ff:取表单值到Map泛型类
// kvMap:
func (r *Request) GetFormMapStrVar(kvMap ...map[string]interface{}) map[string]*gvar.Var {
	formMap := r.GetFormMap(kvMap...)
	if len(formMap) > 0 {
		m := make(map[string]*gvar.Var, len(formMap))
		for k, v := range formMap {
			m[k] = gvar.New(v)
		}
		return m
	}
	return nil
}

// GetFormStruct retrieves all form parameters passed from client and converts them to
// given struct object. Note that the parameter `pointer` is a pointer to the struct object.
// The optional parameter `mapping` is used to specify the key to attribute mapping.

// ff:取表单值到结构
// mapping:
// pointer:结构指针
func (r *Request) GetFormStruct(pointer interface{}, mapping ...map[string]string) error {
	_, err := r.doGetFormStruct(pointer, mapping...)
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
