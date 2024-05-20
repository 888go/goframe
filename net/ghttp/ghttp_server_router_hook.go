// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"context"
	"net/http"
	"reflect"

	"github.com/gogf/gf/v2/debug/gdebug"
)

// BindHookHandler 为指定的钩子注册处理程序。. md5:325d65ceb75f1b33
func (s *Server) BindHookHandler(pattern string, hook HookName, handler HandlerFunc) {
	s.doBindHookHandler(context.TODO(), doBindHookHandlerInput{
		Prefix:   "",
		Pattern:  pattern,
		HookName: hook,
		Handler:  handler,
		Source:   "",
	})
}

// doBindHookHandlerInput是BindHookHandler的输入参数。. md5:1a2c964079a2227c
type doBindHookHandlerInput struct {
	Prefix   string
	Pattern  string
	HookName HookName
	Handler  HandlerFunc
	Source   string
}

// doBindHookHandler是BindHookHandler的内部处理程序。. md5:5393ffad5084d597
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

// BindHookHandlerByMap 为指定的钩子注册处理器。. md5:38d4b79317ac1b3f
func (s *Server) BindHookHandlerByMap(pattern string, hookMap map[HookName]HandlerFunc) {
	for k, v := range hookMap {
		s.BindHookHandler(pattern, k, v)
	}
}

// callHookHandler 按照注册的顺序调用钩子处理器。. md5:4e1a8b2998b73ddb
func (s *Server) callHookHandler(hook HookName, r *Request) {
	if !r.hasHookHandler {
		return
	}
	hookItems := r.getHookHandlers(hook)
	if len(hookItems) > 0 {
		// 备份旧的路由器变量映射。. md5:4a1427ee4ccef0a6
		oldRouterMap := r.routerMap
		for _, item := range hookItems {
			r.routerMap = item.Values
// 不要使用钩子处理器的路由器，
// 因为它可能会覆盖服务处理器的路由器。
// r.Router = item.handler.router
// md5:9c797403c522d44d
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
		// 恢复旧的路由器变量映射。. md5:6ae23d30567bb237
		r.routerMap = oldRouterMap
	}
}

// getHookHandlers 获取并返回指定钩子的处理函数。. md5:f19f77b15aa76d7a
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

// niceCallHookHandler 美好地调用钩子处理函数，
// 即它会自动捕获并返回可能的恐慌错误，以防止goroutine崩溃。
// md5:915bcff9c5f9cc4e
func (s *Server) niceCallHookHandler(f HandlerFunc, r *Request) (err interface{}) {
	defer func() {
		err = recover()
	}()
	f(r)
	return
}
