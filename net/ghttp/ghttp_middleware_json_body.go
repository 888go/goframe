// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"github.com/888go/goframe/internal/json"
)

// X中间件函数_验证JSON格式请求体 验证并返回请求体是否为JSON格式。 md5:861efc2736ae2e13
func X中间件函数_验证JSON格式请求体(r *Request) {
	requestBody := r.X取请求体字节集()
	if len(requestBody) > 0 {
		if !json.Valid(requestBody) {
			r.X设置错误信息(ERR请求体必须json格式)
			return
		}
	}
	r.X中间件管理器.Next()
}
