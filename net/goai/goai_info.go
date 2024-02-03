// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

// Info 由 OpenAPI/Swagger 标准版本 3.0 规定。
type Info struct {
	Title          string   `json:"title"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Version        string   `json:"version"`
}

// Contact 是由 OpenAPI/Swagger 标准版本 3.0 规定的联系信息结构体。
type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

// License 由 OpenAPI/Swagger 标准版本 3.0 规定。
type License struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}
