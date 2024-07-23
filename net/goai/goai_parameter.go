// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/util/gconv"
)

// 参数由 OpenAPI/Swagger 3.0 标准定义。
// 参见 https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterObject
// md5:c86f94d3392df58e
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
	if err := gconv.Struct(mergedTagMap, parameter); err != nil {
		return gerror.Wrap(err, `mapping struct tags to Parameter failed`)
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
	type tempParameter Parameter // 为了防止JSON序列化时的递归错误。 md5:add9f5a47e638cc5
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
