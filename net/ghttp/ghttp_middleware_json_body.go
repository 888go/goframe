// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"github.com/888go/goframe/internal/json"
)

// MiddlewareJsonBody 验证并返回请求体是否为 JSON 格式。
func X中间件函数_验证JSON格式请求体(r *Request) {
	requestBody := r.X取请求体字节集()
	if len(requestBody) > 0 {
		if !json.Valid(requestBody) {
			r.X设置错误信息(ErrNeedJsonBody)
			return
		}
	}
	r.Middleware.Next()
}
