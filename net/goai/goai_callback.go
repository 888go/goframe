// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

import (
	"github.com/gogf/gf/v2/internal/json"
)

// Callback 是由 OpenAPI/Swagger 标准版本 3.0 定义的。. md5:af37b49b06e0c2b0
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
