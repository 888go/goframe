// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"github.com/888go/goframe/container/gvar"
)

// SetParam 设置自定义参数，使用键值对。
func (r *Request) X设置自定义参数(名称 string, 值 interface{}) {
	if r.paramsMap == nil {
		r.paramsMap = make(map[string]interface{})
	}
	r.paramsMap[名称] = 值
}

// SetParamMap 用于设置自定义参数，采用键值对形式的映射。
func (r *Request) X设置自定义参数Map(参数 map[string]interface{}) {
	if r.paramsMap == nil {
		r.paramsMap = make(map[string]interface{})
	}
	for k, v := range 参数 {
		r.paramsMap[k] = v
	}
}

// GetParam 函数用于获取自定义参数，通过给定的名称 `key`。
// 若 `key` 不存在，则返回 `def`。
// 若未提供 `def`，则返回 nil。
func (r *Request) X取自定义参数到泛型类(名称 string, 默认值 ...interface{}) *泛型类.Var {
	if len(r.paramsMap) > 0 {
		if value, ok := r.paramsMap[名称]; ok {
			return 泛型类.X创建(value)
		}
	}
	if len(默认值) > 0 {
		return 泛型类.X创建(默认值[0])
	}
	return nil
}
