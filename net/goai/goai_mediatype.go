// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

// MediaType 是由 OpenAPI/Swagger 3.0 标准定义的。. md5:1bb86450c61271e0
type MediaType struct {
	Schema   *SchemaRef           `json:"schema,omitempty"`
	Example  interface{}          `json:"example,omitempty"`
	Examples Examples             `json:"examples,omitempty"`
	Encoding map[string]*Encoding `json:"encoding,omitempty"`
}

// 内容遵循 OpenAPI/Swagger 3.0 规范定义。. md5:c70257c0a9cf1576
type Content map[string]MediaType

// Encoding 按照 OpenAPI/Swagger 3.0 标准指定。. md5:f20471879db70e3a
type Encoding struct {
	ContentType   string  `json:"contentType,omitempty"`
	Headers       Headers `json:"headers,omitempty"`
	Style         string  `json:"style,omitempty"`
	Explode       *bool   `json:"explode,omitempty"`
	AllowReserved bool    `json:"allowReserved,omitempty"`
}
