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

// Response 由 OpenAPI/Swagger 3.0 标准指定。
type Response struct {
	Description string      `json:"description"`
	Headers     Headers     `json:"headers,omitempty"`
	Content     Content     `json:"content,omitempty"`
	Links       Links       `json:"links,omitempty"`
	XExtensions XExtensions `json:"-"`
}

func (oai *OpenApiV3) tagMapToResponse(tagMap map[string]string, response *Response) error {
	var mergedTagMap = oai.fillMapWithShortTags(tagMap)
	if err := 转换类.Struct(mergedTagMap, response); err != nil {
		return 错误类.X多层错误(err, `mapping struct tags to Response failed`)
	}
	oai.tagMapToXExtensions(mergedTagMap, response.XExtensions)
	return nil
}

func (r Response) MarshalJSON() ([]byte, error) {
	var (
		b   []byte
		m   map[string]json.RawMessage
		err error
	)
	type tempResponse Response // 为防止JSON序列化时出现递归错误
	if b, err = json.Marshal(tempResponse(r)); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for k, v := range r.XExtensions {
		if b, err = json.Marshal(v); err != nil {
			return nil, err
		}
		m[k] = b
	}
	return json.Marshal(m)
}
