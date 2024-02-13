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

// 参数遵循OpenAPI/Swagger 3.0标准指定。
// 详情请参阅：https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterObject
type Parameter struct {
	Name            string      `json:"name,omitempty"`
	In              string      `json:"in,omitempty"`
	Description     string      `json:"description,omitempty"`
	Style           string      `json:"style,omitempty"`
	Explode         *bool       `json:"explode,omitempty"`
	AllowEmptyValue bool        `json:"allowEmptyValue,omitempty"`
	AllowReserved   bool        `json:"allowReserved,omitempty"`
	Deprecated      bool        `json:"deprecated,omitempty"`
	Required        bool        `json:"required,omitempty"`
	Schema          *SchemaRef  `json:"schema,omitempty"`
	Example         interface{} `json:"example,omitempty"`
	Examples        *Examples   `json:"examples,omitempty"`
	Content         *Content    `json:"content,omitempty"`
	XExtensions     XExtensions `json:"-"`
}

func (oai *OpenApiV3) tagMapToParameter(tagMap map[string]string, parameter *Parameter) error {
	var mergedTagMap = oai.fillMapWithShortTags(tagMap)
	if err := 转换类.Struct(mergedTagMap, parameter); err != nil {
		return 错误类.X多层错误(err, `mapping struct tags to Parameter failed`)
	}
	oai.tagMapToXExtensions(mergedTagMap, parameter.XExtensions)
	return nil
}

func (p Parameter) MarshalJSON() ([]byte, error) {
	var (
		b   []byte
		m   map[string]json.RawMessage
		err error
	)
	type tempParameter Parameter // 为防止JSON序列化时出现递归错误
	if b, err = json.Marshal(tempParameter(p)); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for k, v := range p.XExtensions {
		if b, err = json.Marshal(v); err != nil {
			return nil, err
		}
		m[k] = b
	}
	return json.Marshal(m)
}
