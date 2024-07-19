// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"github.com/gogf/gf/v2/internal/json"
)

// MiddlewareJsonBody validates and returns request body whether JSON format.
// ff:中间件函数_验证JSON格式请求体
// r:
func MiddlewareJsonBody(r *Request) {
	requestBody := r.GetBody()
	if len(requestBody) > 0 {
		if !json.Valid(requestBody) {
			r.SetError(ErrNeedJsonBody)
			return
		}
	}
	r.Middleware.Next()
}
