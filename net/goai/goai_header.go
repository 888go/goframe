// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

import (
	"github.com/888go/goframe/internal/json"
)

// Header遵循OpenAPI/Swagger 3.0标准。
// 参见：https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#headerObject
// md5:df7c9f506710dc39
type Header struct {
	Parameter
}

type Headers map[string]HeaderRef

type HeaderRef struct {
	Ref   string
	Value *Header
}

func (r HeaderRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}
