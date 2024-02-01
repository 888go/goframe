// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai
import (
	"github.com/888go/goframe/internal/json"
	)
// Header 按照 OpenAPI/Swagger 3.0 标准进行指定。
// 详情请参阅 https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#headerObject
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
