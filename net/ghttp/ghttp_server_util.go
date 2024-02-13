// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"net/http"
)

// WrapF 是一个辅助函数，用于封装 http.HandlerFunc，并返回一个 ghttp.HandlerFunc 类型的处理函数。
func WrapF(f http.HandlerFunc) HandlerFunc {
	return func(r *Request) {
		f(r.Response.Writer, r.Request)
	}
}

// WrapH 是一个辅助函数，用于包装 http.Handler，并返回一个 ghttp.HandlerFunc 类型。
func WrapH(h http.Handler) HandlerFunc {
	return func(r *Request) {
		h.ServeHTTP(r.Response.Writer, r.Request)
	}
}
