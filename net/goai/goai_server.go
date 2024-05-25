// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

// Server由OpenAPI/Swagger标准版本3.0定义。. md5:e48b24980b097496
type Server struct {
	URL         string                     `json:"url"`
	Description string                     `json:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty"`
}

// ServerVariable 是由OpenAPI/Swagger标准版本3.0定义的。. md5:7cbb925477ff19bd
type ServerVariable struct {
	Enum        []string `json:"enum,omitempty"`
	Default     string   `json:"default,omitempty"`
	Description string   `json:"description,omitempty"`
}

// Servers 是由 OpenAPI/Swagger 标准版本3.0规定的。. md5:8a433e291e8eabb8
type Servers []Server
