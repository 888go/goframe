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
	"reflect"

	"github.com/888go/goframe/debug/gdebug"
)

// X绑定Hook 为指定的钩子注册处理程序。 md5:325d65ceb75f1b33
func (s *X服务) X绑定Hook(路由规则 string, 触发时机 Hook名称, 处理函数 HandlerFunc) {
	s.doBindHookHandler(context.TODO(), doBindHookHandlerInput{
		Prefix:   "",
		Pattern:  路由规则,
		HookName: 触发时机,
		Handler:  处理函数,
		Source:   "",
	})
}

// doBindHookHandlerInput是BindHookHandler的输入参数。 md5:1a2c964079a2227c
type doBindHookHandlerInput struct {
	Prefix   string
	Pattern  string
	HookName Hook名称
	Handler  HandlerFunc
	Source   string
}

// doBindHookHandler是BindHookHandler的内部处理程序。 md5:5393ffad5084d597
func (s *X服务) doBindHookHandler(ctx context.Context, in doBindHookHandlerInput) {
	s.setHandler(
		ctx,
		setHandlerInput{
			Prefix:  in.Prefix,
			Pattern: in.Pattern,
			HandlerItem: &X路由处理函数{
				Type: HandlerTypeHook,
				X处理器名称: gdebug.FuncPath(in.Handler),
				X处理器函数信息: handlerFuncInfo{
					Func: in.Handler,
					Type: reflect.TypeOf(in.Handler),
				},
				Hook名称: in.HookName,
				X注册来源:   in.Source,
			},
		},
	)
}

// X绑定HookMap 为指定的钩子注册处理器。 md5:38d4b79317ac1b3f
func (s *X服务) X绑定HookMap(路由规则 string, HookMap map[Hook名称]HandlerFunc) {
	for k, v := range HookMap {
		s.X绑定Hook(路由规则, k, v)
	}
}

// callHookHandler 按照注册的顺序调用钩子处理器。 md5:4e1a8b2998b73ddb
func (s *X服务) callHookHandler(hook Hook名称, r *Request) {
	if !r.hasHookHandler {
		return
	}
	hookItems := r.getHookHandlers(hook)
	if len(hookItems) > 0 {
				// 备份旧的路由器变量映射。 md5:4a1427ee4ccef0a6
		oldRouterMap := r.routerMap
		for _, item := range hookItems {
			r.routerMap = item.X路由值
			// 不要使用钩子处理器的路由器，
			// 因为它可能会覆盖服务处理器的路由器。
			// r.Router = item.handler.router
			// md5:9c797403c522d44d
			if err := s.niceCallHookHandler(item.Handler.X处理器函数信息.Func, r); err != nil {
				switch err {
				case exceptionExit:
					break
				case exceptionExitAll:
					fallthrough
				case exceptionExitHook:
					return
				default:
					r.X响应.X写响应缓冲区与HTTP状态码(http.StatusInternalServerError, err)
					panic(err)
				}
			}
		}
				// 恢复旧的路由器变量映射。 md5:6ae23d30567bb237
		r.routerMap = oldRouterMap
	}
}

// getHookHandlers 获取并返回指定钩子的处理函数。 md5:f19f77b15aa76d7a
func (r *Request) getHookHandlers(hook Hook名称) []*X路由解析 {
	parsedItems := make([]*X路由解析, 0, 4)
	for _, v := range r.handlers {
		if v.Handler.Hook名称 != hook {
			continue
		}
		item := v
		parsedItems = append(parsedItems, item)
	}
	return parsedItems
}

// niceCallHookHandler 美好地调用钩子处理函数，
// 即它会自动捕获并返回可能的恐慌错误，以防止goroutine崩溃。
// md5:915bcff9c5f9cc4e
func (s *X服务) niceCallHookHandler(f HandlerFunc, r *Request) (err interface{}) {
	defer func() {
		err = recover()
	}()
	f(r)
	return
}
