// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"context"
	"net/http"
	"reflect"
	
	"github.com/888go/goframe/debug/gdebug"
)

// BindHookHandler 为指定的钩子注册处理器。
func (s *X服务) X绑定Hook(路由规则 string, 触发时机 Hook名称, 处理函数 HandlerFunc) {
	s.doBindHookHandler(context.TODO(), doBindHookHandlerInput{
		Prefix:   "",
		Pattern:  路由规则,
		HookName: 触发时机,
		Handler:  处理函数,
		Source:   "",
	})
}

// doBindHookHandlerInput 是 BindHookHandler 的输入参数。
type doBindHookHandlerInput struct {
	Prefix   string
	Pattern  string
	HookName Hook名称
	Handler  HandlerFunc
	Source   string
}

// doBindHookHandler 是 BindHookHandler 的内部处理程序。
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

// BindHookHandlerByMap 为指定的钩子注册处理器。
func (s *X服务) X绑定HookMap(路由规则 string, HookMap map[Hook名称]HandlerFunc) {
	for k, v := range HookMap {
		s.X绑定Hook(路由规则, k, v)
	}
}

// callHookHandler 按照已注册的顺序调用钩子处理器
func (s *X服务) callHookHandler(hook Hook名称, r *X请求) {
	if !r.hasHookHandler {
		return
	}
	hookItems := r.getHookHandlers(hook)
	if len(hookItems) > 0 {
		// 备份旧的路由器变量映射。
		oldRouterMap := r.routerMap
		for _, item := range hookItems {
			r.routerMap = item.X路由值
// **不要在钩子处理器中使用路由器**，
// 这可能会覆盖服务处理器的路由器。
// r.Router = item.handler.router
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
		// 恢复旧的路由器变量映射。
		r.routerMap = oldRouterMap
	}
}

// getHookHandlers 获取并返回指定钩子的钩子处理程序。
func (r *X请求) getHookHandlers(hook Hook名称) []*X路由解析 {
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

// niceCallHookHandler 优雅地调用钩子处理器函数，
// 这意味着它会自动捕获并返回可能的 panic 错误，以避免 goroutine 中止。
func (s *X服务) niceCallHookHandler(f HandlerFunc, r *X请求) (err interface{}) {
	defer func() {
		err = recover()
	}()
	f(r)
	return
}
