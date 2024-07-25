// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package ghttp

import "net/http"

// WrapF 是一个辅助函数，用于包装http.HandlerFunc并返回一个ghttp.HandlerFunc。 md5:f1b5a37e2bddfd19
func WrapF(f http.HandlerFunc) HandlerFunc {
	return func(r *Request) {
		f(r.Response.Writer, r.Request)
	}
}

// WrapH 是一个辅助函数，用于包装 http.Handler，并返回一个 ghttp.HandlerFunc。 md5:0d35a772811803c8
func WrapH(h http.Handler) HandlerFunc {
	return func(r *Request) {
		h.ServeHTTP(r.Response.Writer, r.Request)
	}
}
