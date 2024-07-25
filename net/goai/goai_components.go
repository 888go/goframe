// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

// Components是根据OpenAPI/Swagger标准版本3.0指定的。 md5:ac796372f37158be
type Components struct {
	Schemas         Schemas         `json:"schemas,omitempty"`
	Parameters      ParametersMap   `json:"parameters,omitempty"`
	Headers         Headers         `json:"headers,omitempty"`
	RequestBodies   RequestBodies   `json:"requestBodies,omitempty"`
	Responses       Responses       `json:"responses,omitempty"`
	SecuritySchemes SecuritySchemes `json:"securitySchemes,omitempty"`
	Examples        Examples        `json:"examples,omitempty"`
	Links           Links           `json:"links,omitempty"`
	Callbacks       Callbacks       `json:"callbacks,omitempty"`
}

type ParametersMap map[string]*ParameterRef

type RequestBodies map[string]*RequestBodyRef
