// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

// Tags是根据OpenAPI/Swagger 3.0标准指定的。 md5:dfa5f209aa36f9b2
type Tags []Tag

// Tag由OpenAPI/Swagger 3.0标准指定。 md5:94223e7c6dcb7c0f
type Tag struct {
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
}
