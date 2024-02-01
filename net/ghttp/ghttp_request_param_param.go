// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp


import (
	"github.com/888go/goframe/container/gvar"
	)
// SetParam 设置自定义参数，使用键值对。
func (r *Request) SetParam(key string, value interface{}) {
	if r.paramsMap == nil {
		r.paramsMap = make(map[string]interface{})
	}
	r.paramsMap[key] = value
}

// SetParamMap 用于设置自定义参数，采用键值对形式的映射。
func (r *Request) SetParamMap(data map[string]interface{}) {
	if r.paramsMap == nil {
		r.paramsMap = make(map[string]interface{})
	}
	for k, v := range data {
		r.paramsMap[k] = v
	}
}

// GetParam 函数用于获取自定义参数，通过给定的名称 `key`。
// 若 `key` 不存在，则返回 `def`。
// 若未提供 `def`，则返回 nil。
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
