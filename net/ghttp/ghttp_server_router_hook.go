// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

import (
	"context"
	"net/http"
	"reflect"
	
	"github.com/888go/goframe/debug/gdebug"
)

// BindHookHandler 为指定的钩子注册处理器。
func (s *Server) BindHookHandler(pattern string, hook HookName, handler HandlerFunc) {
	s.doBindHookHandler(context.TODO(), doBindHookHandlerInput{
		Prefix:   "",
		Pattern:  pattern,
		HookName: hook,
		Handler:  handler,
		Source:   "",
	})
}

// doBindHookHandlerInput 是 BindHookHandler 的输入参数。
type doBindHookHandlerInput struct {
	Prefix   string
	Pattern  string
	HookName HookName
	Handler  HandlerFunc
	Source   string
}

// doBindHookHandler 是 BindHookHandler 的内部处理程序。
func (s *Server) doBindHookHandler(ctx context.Context, in doBindHookHandlerInput) {
	s.setHandler(
		ctx,
		setHandlerInput{
			Prefix:  in.Prefix,
			Pattern: in.Pattern,
			HandlerItem: &HandlerItem{
				Type: HandlerTypeHook,
				Name: gdebug.FuncPath(in.Handler),
				Info: handlerFuncInfo{
					Func: in.Handler,
					Type: reflect.TypeOf(in.Handler),
				},
				HookName: in.HookName,
				Source:   in.Source,
			},
		},
	)
}

// BindHookHandlerByMap 为指定的钩子注册处理器。
func (s *Server) BindHookHandlerByMap(pattern string, hookMap map[HookName]HandlerFunc) {
	for k, v := range hookMap {
		s.BindHookHandler(pattern, k, v)
	}
}

// callHookHandler 按照已注册的顺序调用钩子处理器
func (s *Server) callHookHandler(hook HookName, r *Request) {
	if !r.hasHookHandler {
		return
	}
	hookItems := r.getHookHandlers(hook)
	if len(hookItems) > 0 {
		// 备份旧的路由器变量映射。
		oldRouterMap := r.routerMap
		for _, item := range hookItems {
			r.routerMap = item.Values
// **不要在钩子处理器中使用路由器**，
// 这可能会覆盖服务处理器的路由器。
// r.Router = item.handler.router
			if err := s.niceCallHookHandler(item.Handler.Info.Func, r); err != nil {
				switch err {
				case exceptionExit:
					break
				case exceptionExitAll:
					fallthrough
				case exceptionExitHook:
					return
				default:
					r.Response.WriteStatus(http.StatusInternalServerError, err)
					panic(err)
				}
			}
		}
		// 恢复旧的路由器变量映射。
		r.routerMap = oldRouterMap
	}
}

// getHookHandlers 获取并返回指定钩子的钩子处理程序。
func (r *Request) getHookHandlers(hook HookName) []*HandlerItemParsed {
	parsedItems := make([]*HandlerItemParsed, 0, 4)
	for _, v := range r.handlers {
		if v.Handler.HookName != hook {
			continue
		}
		item := v
		parsedItems = append(parsedItems, item)
	}
	return parsedItems
}

// niceCallHookHandler 优雅地调用钩子处理器函数，
// 这意味着它会自动捕获并返回可能的 panic 错误，以避免 goroutine 中止。
func (s *Server) niceCallHookHandler(f HandlerFunc, r *Request) (err interface{}) {
	defer func() {
		err = recover()
	}()
	f(r)
	return
}
