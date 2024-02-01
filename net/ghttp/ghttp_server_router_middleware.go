// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"context"
	"reflect"
	
	"github.com/888go/goframe/debug/gdebug"
	)
const (
	// 全局中间件的默认路由模式。
	defaultMiddlewarePattern = "/*"
)

// BindMiddleware 注册一个或多个全局中间件到服务器。
// 全局中间件可以在没有服务处理器的情况下独立使用，它会在服务处理器执行前或执行后拦截所有的动态请求。
// 参数 `pattern` 指定了中间件要拦截的路由模式，通常是一个“模糊”模式，如 "/:name"、"/*any" 或 "/{field}"。
func (s *Server) BindMiddleware(pattern string, handlers ...HandlerFunc) {
	var (
		ctx = context.TODO()
	)
	for _, handler := range handlers {
		s.setHandler(ctx, setHandlerInput{
			Prefix:  "",
			Pattern: pattern,
			HandlerItem: &HandlerItem{
				Type: HandlerTypeMiddleware,
				Name: gdebug.FuncPath(handler),
				Info: handlerFuncInfo{
					Func: handler,
					Type: reflect.TypeOf(handler),
				},
			},
		})
	}
}

// BindMiddlewareDefault 使用默认模式"/*"将一个或多个全局中间件注册到服务器。
// 全局中间件可以在没有服务处理器的情况下独立使用，它会在服务处理器处理所有动态请求之前或之后进行拦截。
func (s *Server) BindMiddlewareDefault(handlers ...HandlerFunc) {
	var (
		ctx = context.TODO()
	)
	for _, handler := range handlers {
		s.setHandler(ctx, setHandlerInput{
			Prefix:  "",
			Pattern: defaultMiddlewarePattern,
			HandlerItem: &HandlerItem{
				Type: HandlerTypeMiddleware,
				Name: gdebug.FuncPath(handler),
				Info: handlerFuncInfo{
					Func: handler,
					Type: reflect.TypeOf(handler),
				},
			},
		})
	}
}

// Use 是 BindMiddlewareDefault 的别名。
// 请参阅 BindMiddlewareDefault。
func (s *Server) Use(handlers ...HandlerFunc) {
	s.BindMiddlewareDefault(handlers...)
}
