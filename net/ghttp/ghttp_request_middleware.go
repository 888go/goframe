// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"context"
	"net/http"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/util/gutil"
)

// middleware 是用于请求工作流程管理的插件。
type middleware struct {
	served         bool     // Is the request served 是用于检查响应状态404的，即该请求是否已成功处理。
	request        *X请求 // 请求对象指针。
	handlerIndex   int      // Index 号用于执行顺序目的，针对处理项。
	handlerMDIndex int      // Index 数值用于执行顺序的目的，用于处理器项绑定的中间件。
}

// Next调用下一个工作流处理器。
// 这是一个重要的函数，用于控制服务器请求执行的工作流程。
func (m *middleware) Next() {
	var item *X路由解析
	var loop = true
	for loop {
		// 检查请求是否已激发。
		if m.request.X是否已退出() || m.handlerIndex >= len(m.request.handlers) {
			break
		}
		item = m.request.handlers[m.handlerIndex]
		// 过滤HOOK处理程序，它们设计为在另一个独立的进程中被调用。
		if item.Handler.Type == HandlerTypeHook {
			m.handlerIndex++
			continue
		}
		// 当前路由器切换中。
		m.request.X路由 = item.Handler.X路由

		// 路由值切换。
		m.request.routerMap = item.X路由值

		var ctx = m.request.Context别名()
		工具类.X异常捕捉并带异常处理(ctx, func(ctx context.Context) {
			// 如果item的绑定中间件数组不为空，则执行该数组中的中间件。
			if m.handlerMDIndex < len(item.Handler.X中间件数组) {
				md := item.Handler.X中间件数组[m.handlerMDIndex]
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

			// 全局中间件数组。
			case HandlerTypeMiddleware:
				niceCallFunc(func() {
					item.Handler.X处理器函数信息.Func(m.request)
				})
// 当某个中间件执行完毕后，它不会继续调用下一个中间件。
// 若要管理工作流程，应在中间件中调用“Next”函数。
				loop = false
			}
		}, func(ctx context.Context, exception error) {
			if 错误类.X判断是否带堆栈(exception) {
				// 这已经是一个带有堆栈信息的错误。
				m.request.error = exception
			} else {
// 创建一个包含堆栈信息的新错误。
// 注意，这里有一个skip参数用于指向实际错误点的堆栈跟踪起始位置。
				m.request.error = 错误类.X多层错误码并跳过堆栈(错误码类.CodeInternalError, 1, exception, "")
			}
			m.request.X响应.X写响应缓冲区与HTTP状态码(http.StatusInternalServerError, exception)
			loop = false
		})
	}
	// 在所有处理器和中间件执行完毕后，检查HTTP状态码。
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
