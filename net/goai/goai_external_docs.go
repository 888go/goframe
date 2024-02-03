// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

import (
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

// ExternalDocs 是由 OpenAPI/Swagger 标准 3.0 版本定义的。
type ExternalDocs struct {
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

func (ed *ExternalDocs) UnmarshalValue(value interface{}) error {
	var valueBytes = gconv.Bytes(value)
	if json.Valid(valueBytes) {
		return json.UnmarshalUseNumber(valueBytes, ed)
	}
	var (
		valueString = string(valueBytes)
		valueArray  = gstr.Split(valueString, "|")
	)
	ed.URL = valueArray[0]
	if len(valueArray) > 1 {
		ed.Description = valueArray[1]
	}
	return nil
}
