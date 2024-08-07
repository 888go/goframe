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

// Link 是由OpenAPI/Swagger标准版本3.0定义的。 md5:7e7a4f467c0cd7ed
type Link struct {
	OperationID  string                 `json:"operationId,omitempty"`
	OperationRef string                 `json:"operationRef,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
	Server       *Server                `json:"server,omitempty"`
	RequestBody  interface{}            `json:"requestBody,omitempty"`
}

type Links map[string]LinkRef

type LinkRef struct {
	Ref   string
	Value *Link
}

func (r LinkRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}
