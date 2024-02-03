// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

// Components是按照OpenAPI/Swagger 3.0标准指定的。
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
