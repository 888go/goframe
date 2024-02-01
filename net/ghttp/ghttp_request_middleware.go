// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
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
	request        *Request // 请求对象指针。
	handlerIndex   int      // Index 号用于执行顺序目的，针对处理项。
	handlerMDIndex int      // Index 数值用于执行顺序的目的，用于处理器项绑定的中间件。
}

// Next调用下一个工作流处理器。
// 这是一个重要的函数，用于控制服务器请求执行的工作流程。
func (m *middleware) Next() {
	var item *HandlerItemParsed
	var loop = true
	for loop {
		// 检查请求是否已激发。
		if m.request.IsExited() || m.handlerIndex >= len(m.request.handlers) {
			break
		}
		item = m.request.handlers[m.handlerIndex]
		// 过滤HOOK处理程序，它们设计为在另一个独立的进程中被调用。
		if item.Handler.Type == HandlerTypeHook {
			m.handlerIndex++
			continue
		}
		// 当前路由器切换中。
		m.request.Router = item.Handler.Router

		// 路由值切换。
		m.request.routerMap = item.Values

		var ctx = m.request.Context()
		gutil.TryCatch(ctx, func(ctx context.Context) {
			// 如果item的绑定中间件数组不为空，则执行该数组中的中间件。
			if m.handlerMDIndex < len(item.Handler.Middleware) {
				md := item.Handler.Middleware[m.handlerMDIndex]
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
				if m.request.IsExited() {
					break
				}
				if item.Handler.InitFunc != nil {
					niceCallFunc(func() {
						item.Handler.InitFunc(m.request)
					})
				}
				if !m.request.IsExited() {
					m.callHandlerFunc(item.Handler.Info)
				}
				if !m.request.IsExited() && item.Handler.ShutFunc != nil {
					niceCallFunc(func() {
						item.Handler.ShutFunc(m.request)
					})
				}

			// Service handler.
			case HandlerTypeHandler:
				m.served = true
				if m.request.IsExited() {
					break
				}
				niceCallFunc(func() {
					m.callHandlerFunc(item.Handler.Info)
				})

			// 全局中间件数组。
			case HandlerTypeMiddleware:
				niceCallFunc(func() {
					item.Handler.Info.Func(m.request)
				})
// 当某个中间件执行完毕后，它不会继续调用下一个中间件。
// 若要管理工作流程，应在中间件中调用“Next”函数。
				loop = false
			}
		}, func(ctx context.Context, exception error) {
			if gerror.HasStack(exception) {
				// 这已经是一个带有堆栈信息的错误。
				m.request.error = exception
			} else {
// 创建一个包含堆栈信息的新错误。
// 注意，这里有一个skip参数用于指向实际错误点的堆栈跟踪起始位置。
				m.request.error = gerror.WrapCodeSkip(gcode.CodeInternalError, 1, exception, "")
			}
			m.request.Response.WriteStatus(http.StatusInternalServerError, exception)
			loop = false
		})
	}
	// 在所有处理器和中间件执行完毕后，检查HTTP状态码。
	if m.request.IsExited() || m.handlerIndex >= len(m.request.handlers) {
		if m.request.Response.Status == 0 {
			if m.request.Middleware.served {
				m.request.Response.WriteHeader(http.StatusOK)
			} else {
				m.request.Response.WriteHeader(http.StatusNotFound)
			}
		}
	}
}

func (m *middleware) callHandlerFunc(funcInfo handlerFuncInfo) {
	niceCallFunc(func() {
		funcInfo.Func(m.request)
	})
}
