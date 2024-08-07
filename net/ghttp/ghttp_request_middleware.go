// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"context"
	"net/http"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	gutil "github.com/888go/goframe/util/gutil"
)

// middleware是用于请求工作流程管理的插件。 md5:49f6a68042876e4b
type middleware struct {
	served         bool     // 请求是否已处理，用于检查响应状态码404。 md5:d6a59089bf4bca74
	request        *Request // 请求对象的指针。 md5:d2b2b51d27020540
	handlerIndex   int      // 用于处理项执行顺序的索引号。 md5:98c1f7684660f958
	handlerMDIndex int      // 用于处理项绑定中间件的执行顺序编号。 md5:ea6ccd0cbc909909
}

// Next 调用下一个工作流处理器。
// 这是一个重要的函数，用于控制服务器请求执行的工作流程。
// md5:9993825368a59675
func (m *middleware) Next() {
	var item *X路由解析
	var loop = true
	for loop {
				// 检查请求是否已激活。 md5:87f631160593048d
		if m.request.X是否已退出() || m.handlerIndex >= len(m.request.handlers) {
			break
		}
		item = m.request.handlers[m.handlerIndex]
				// 过滤HOOK处理器，这些处理器设计用于在另一个独立的程序中被调用。 md5:f285ad394cb72a16
		if item.Handler.Type == HandlerTypeHook {
			m.handlerIndex++
			continue
		}
				// 当前路由器切换。 md5:6e5ea35720c091cd
		m.request.X路由 = item.Handler.X路由

				// 路由器值切换。 md5:8a03c212d83f07ec
		m.request.routerMap = item.X路由值

		var ctx = m.request.Context别名()
		gutil.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
						// 如果项的绑定中间件数组不为空，则执行它。 md5:c2326e1bb33e0423
			if m.handlerMDIndex < len(item.Handler.X中间件切片) {
				md := item.Handler.X中间件切片[m.handlerMDIndex]
				m.handlerMDIndex++
				niceCallFunc(func() {
					md(m.request)
				})
				loop = false
				return
			}
			m.handlerIndex++

			switch item.Handler.Type {
			// Service object.
			case HandlerTypeObject:
				m.served = true
				if m.request.X是否已退出() {
					break
				}
				if item.Handler.X初始化回调函数 != nil {
					niceCallFunc(func() {
						item.Handler.X初始化回调函数(m.request)
					})
				}
				if !m.request.X是否已退出() {
					m.callHandlerFunc(item.Handler.X处理器函数信息)
				}
				if !m.request.X是否已退出() && item.Handler.X关闭回调函数 != nil {
					niceCallFunc(func() {
						item.Handler.X关闭回调函数(m.request)
					})
				}

			// Service handler.
			case HandlerTypeHandler:
				m.served = true
				if m.request.X是否已退出() {
					break
				}
				niceCallFunc(func() {
					m.callHandlerFunc(item.Handler.X处理器函数信息)
				})

						// 全局中间件数组。 md5:1b4420db0d912753
			case HandlerTypeMiddleware:
				niceCallFunc(func() {
					item.Handler.X处理器函数信息.Func(m.request)
				})
				// 它不会在其他中间件完成之后继续调用下一个中间件。为了管理工作流程，中间件应该有一个名为 "Next" 的函数可供调用。
				// md5:0a1a7642101f1bb9
				loop = false
			}
		}, func(ctx context.Context, exception error) {
			if gerror.X判断是否带堆栈(exception) {
								// 它已经是一个带有堆栈信息的错误。 md5:ec045ebe21bca18d
				m.request.error = exception
			} else {
				// 创建一个带有堆栈信息的新错误。
				// 注意，skip 参数指定了从哪个调用栈开始追踪真正的错误点。
				// md5:e23da1f0a4a0c90f
				m.request.error = gerror.X多层错误码并跳过堆栈(gcode.CodeInternalError, 1, exception, "")
			}
			m.request.X响应.X写响应缓冲区与HTTP状态码(http.StatusInternalServerError, exception)
			loop = false
		})
	}
		// 在所有处理程序和中间件完成后检查HTTP状态码。 md5:6db5e4718ab69458
	if m.request.X是否已退出() || m.handlerIndex >= len(m.request.handlers) {
		if m.request.X响应.Status == 0 {
			if m.request.X中间件管理器.served {
				m.request.X响应.WriteHeader(http.StatusOK)
			} else {
				m.request.X响应.WriteHeader(http.StatusNotFound)
			}
		}
	}
}

func (m *middleware) callHandlerFunc(funcInfo handlerFuncInfo) {
	niceCallFunc(func() {
		funcInfo.Func(m.request)
	})
}
