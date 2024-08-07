// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	gvar "github.com/888go/goframe/container/gvar"
)

// X设置自定义参数 设置自定义参数，使用键值对形式。 md5:d3e3851975cc7c6e
func (r *Request) X设置自定义参数(名称 string, 值 interface{}) {
	if r.paramsMap == nil {
		r.paramsMap = make(map[string]interface{})
	}
	r.paramsMap[名称] = 值
}

// X设置自定义参数Map 设置自定义参数，通过键值对映射。 md5:f6685d3faaf3bb78
func (r *Request) X设置自定义参数Map(参数 map[string]interface{}) {
	if r.paramsMap == nil {
		r.paramsMap = make(map[string]interface{})
	}
	for k, v := range 参数 {
		r.paramsMap[k] = v
	}
}

// X取自定义参数到泛型类 通过给定的名称 `key` 返回自定义参数。
// 如果 `key` 不存在，它将返回 `def`。
// 如果没有传递 `def`，它将返回 nil。
// md5:4fe03a677e843703
func (r *Request) X取自定义参数到泛型类(名称 string, 默认值 ...interface{}) *gvar.Var {
	if len(r.paramsMap) > 0 {
		if value, ok := r.paramsMap[名称]; ok {
			return gvar.X创建(value)
		}
	}
	if len(默认值) > 0 {
		return gvar.X创建(默认值[0])
	}
	return nil
}
