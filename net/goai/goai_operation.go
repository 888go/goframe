// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

import (
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/util/gconv"
)

// Operation 表示由 OpenAPI/Swagger 3.0 标准定义的 "操作"。
type Operation struct {
	Tags         []string              `json:"tags,omitempty"`
	Summary      string                `json:"summary,omitempty"`
	Description  string                `json:"description,omitempty"`
	OperationID  string                `json:"operationId,omitempty"`
	Parameters   Parameters            `json:"parameters,omitempty"`
	RequestBody  *RequestBodyRef       `json:"requestBody,omitempty"`
	Responses    Responses             `json:"responses"`
	Deprecated   bool                  `json:"deprecated,omitempty"`
	Callbacks    *Callbacks            `json:"callbacks,omitempty"`
	Security     *SecurityRequirements `json:"security,omitempty"`
	Servers      *Servers              `json:"servers,omitempty"`
	ExternalDocs *ExternalDocs         `json:"externalDocs,omitempty"`
	XExtensions  XExtensions           `json:"-"`
}

func (oai *OpenApiV3) tagMapToOperation(tagMap map[string]string, operation *Operation) error {
	var mergedTagMap = oai.fillMapWithShortTags(tagMap)
	if err := 转换类.Struct(mergedTagMap, operation); err != nil {
		return 错误类.X多层错误(err, `mapping struct tags to Operation failed`)
	}
	oai.tagMapToXExtensions(mergedTagMap, operation.XExtensions)
	return nil
}

func (o Operation) MarshalJSON() ([]byte, error) {
	var (
		b   []byte
		m   map[string]json.RawMessage
		err error
	)
	type tempOperation Operation // 为防止JSON序列化时出现递归错误
	if b, err = json.Marshal(tempOperation(o)); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for k, v := range o.XExtensions {
		if b, err = json.Marshal(v); err != nil {
			return nil, err
		}
		m[k] = b
	}
	return json.Marshal(m)
}
