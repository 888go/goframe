// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

// MiddlewareNeverDoneCtx为当前进程设置永不完成的上下文。. md5:82de09cee1c428a4
func MiddlewareNeverDoneCtx(r *Request) {
	r.SetCtx(r.GetNeverDoneCtx())
	r.Middleware.Next()
}
