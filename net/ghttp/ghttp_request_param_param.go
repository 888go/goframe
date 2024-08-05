// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import "github.com/gogf/gf/v2/container/gvar"

// SetParam 设置自定义参数，使用键值对形式。 md5:d3e3851975cc7c6e
func (r *Request) SetParam(key string, value interface{}) {
	if r.paramsMap == nil {
		r.paramsMap = make(map[string]interface{})
	}
	r.paramsMap[key] = value
}

// SetParamMap 设置自定义参数，通过键值对映射。 md5:f6685d3faaf3bb78
func (r *Request) SetParamMap(data map[string]interface{}) {
	if r.paramsMap == nil {
		r.paramsMap = make(map[string]interface{})
	}
	for k, v := range data {
		r.paramsMap[k] = v
	}
}

// GetParam 通过给定的名称 `key` 返回自定义参数。
// 如果 `key` 不存在，它将返回 `def`。
// 如果没有传递 `def`，它将返回 nil。
// md5:4fe03a677e843703
func (r *Request) GetParam(key string, def ...interface{}) *gvar.Var {
	if len(r.paramsMap) > 0 {
		if value, ok := r.paramsMap[key]; ok {
			return gvar.New(value)
		}
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}
