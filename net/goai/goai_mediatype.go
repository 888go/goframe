// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

// MediaType 是由 OpenAPI/Swagger 3.0 标准指定的。
type MediaType struct {
	Schema   *SchemaRef           `json:"schema,omitempty"`
	Example  interface{}          `json:"example,omitempty"`
	Examples Examples             `json:"examples,omitempty"`
	Encoding map[string]*Encoding `json:"encoding,omitempty"`
}

// Content 是由 OpenAPI/Swagger 3.0 标准指定的。
type Content map[string]MediaType

// 编码由 OpenAPI/Swagger 3.0 标准指定。
type Encoding struct {
	ContentType   string  `json:"contentType,omitempty"`
	Headers       Headers `json:"headers,omitempty"`
	Style         string  `json:"style,omitempty"`
	Explode       *bool   `json:"explode,omitempty"`
	AllowReserved bool    `json:"allowReserved,omitempty"`
}
