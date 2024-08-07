// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"context"
	"reflect"

	"github.com/888go/goframe/debug/gdebug"
)

const (
		// 全局中间件的默认路由模式。 md5:104618f487cee8ac
	defaultMiddlewarePattern = "/*"
)

// X绑定全局中间件 会在服务器上注册一个或多个全局中间件。全局中间件可以在没有服务处理器的情况下单独使用，它会在服务处理器之前或之后拦截所有动态请求。参数 `pattern` 指定了中间件拦截的路由模式，通常是一个模糊模式，如 "/:name"、"/*any" 或 "/{field}"。
// md5:a58488c3f3613ab4
func (s *X服务) X绑定全局中间件(路由规则 string, 处理函数 ...HandlerFunc) {
	var (
		ctx = context.TODO()
	)
	for _, handler := range 处理函数 {
		s.setHandler(ctx, setHandlerInput{
			Prefix:  "",
			Pattern: 路由规则,
			HandlerItem: &X路由处理函数{
				Type: HandlerTypeMiddleware,
				X处理器名称: gdebug.FuncPath(handler),
				X处理器函数信息: handlerFuncInfo{
					Func: handler,
					Type: reflect.TypeOf(handler),
				},
			},
		})
	}
}

// X绑定全局默认中间件 使用默认模式 "/*" 向服务器注册一个或多个全局中间件。
// 全局中间件可以独立使用，无需服务处理器，它能在所有动态请求的前后拦截处理。
// md5:fc212697fcedf39e
func (s *X服务) X绑定全局默认中间件(处理函数 ...HandlerFunc) {
	var (
		ctx = context.TODO()
	)
	for _, handler := range 处理函数 {
		s.setHandler(ctx, setHandlerInput{
			Prefix:  "",
			Pattern: defaultMiddlewarePattern,
			HandlerItem: &X路由处理函数{
				Type: HandlerTypeMiddleware,
				X处理器名称: gdebug.FuncPath(handler),
				X处理器函数信息: handlerFuncInfo{
					Func: handler,
					Type: reflect.TypeOf(handler),
				},
			},
		})
	}
}

// Use别名 是 BindMiddlewareDefault 的别名。
// 参见 BindMiddlewareDefault。
// md5:713ca39a398552e8
func (s *X服务) Use别名(处理函数 ...HandlerFunc) {
	s.X绑定全局默认中间件(处理函数...)
}
