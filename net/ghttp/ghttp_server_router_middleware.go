// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/debug/gdebug"
)

const (
	// 全局中间件的默认路由模式。 md5:104618f487cee8ac
	defaultMiddlewarePattern = "/*"
)

// BindMiddleware 会在服务器上注册一个或多个全局中间件。全局中间件可以在没有服务处理器的情况下单独使用，它会在服务处理器之前或之后拦截所有动态请求。参数 `pattern` 指定了中间件拦截的路由模式，通常是一个模糊模式，如 "/:name"、"/*any" 或 "/{field}"。
// md5:a58488c3f3613ab4
// ff:绑定全局中间件
// s:
// pattern:路由规则
// handlers:处理函数
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

// BindMiddlewareDefault 使用默认模式 "/*" 向服务器注册一个或多个全局中间件。
// 全局中间件可以独立使用，无需服务处理器，它能在所有动态请求的前后拦截处理。
// md5:fc212697fcedf39e
// ff:绑定全局默认中间件
// s:
// handlers:处理函数
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
// 参见 BindMiddlewareDefault。
// md5:713ca39a398552e8
// ff:Use别名
// s:
// handlers:处理函数
func (s *Server) Use(handlers ...HandlerFunc) {
	s.BindMiddlewareDefault(handlers...)
}
