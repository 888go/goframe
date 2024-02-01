// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

// MiddlewareCORS 是一个用于 CORS（跨源资源共享）的中间件处理器，采用默认选项。
func MiddlewareCORS(r *Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
