// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"github.com/gogf/gf/v2/internal/json"
)

// MiddlewareJsonBody 验证并返回请求体是否为JSON格式。 md5:861efc2736ae2e13
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
