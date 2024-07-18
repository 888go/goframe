// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

// Info是根据OpenAPI/Swagger标准版本3.0定义的。 md5:5426a3817852daba
type Info struct {
	Title          string   `json:"title"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Version        string   `json:"version"`
}

// Contact 是根据 OpenAPI/Swagger 标准版本3.0规定的。 md5:604c32a512535948
type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

// License 是由OpenAPI/Swagger标准版本3.0指定的。 md5:6b613c07133dcaa7
type License struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}
