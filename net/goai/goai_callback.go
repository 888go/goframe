// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

import (
	"github.com/888go/goframe/internal/json"
)

// Callback 根据 OpenAPI/Swagger 3.0 标准进行定义。
type Callback map[string]*Path

type Callbacks map[string]*CallbackRef

type CallbackRef struct {
	Ref   string
	Value *Callback
}

func (r CallbackRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}
