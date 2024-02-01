// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"context"
	"strings"
	)
// Domain 用于为域名进行路由注册。
type Domain struct {
	server  *Server             // Belonged server
	domains map[string]struct{} // 支持多个域名。
}

// Domain 创建并返回一个域名管理对象，用于管理一个或多个域名。
func (s *Server) Domain(domains string) *Domain {
	d := &Domain{
		server:  s,
		domains: make(map[string]struct{}),
	}
	for _, v := range strings.Split(domains, ",") {
		d.domains[strings.TrimSpace(v)] = struct{}{}
	}
	return d
}

// BindHandler 为指定模式绑定处理器。
func (d *Domain) BindHandler(pattern string, handler interface{}) {
	for domain := range d.domains {
		d.server.BindHandler(pattern+"@"+domain, handler)
	}
}

func (d *Domain) doBindHandler(ctx context.Context, in doBindHandlerInput) {
	for domain := range d.domains {
		d.server.doBindHandler(ctx, doBindHandlerInput{
			Prefix:     in.Prefix,
			Pattern:    in.Pattern + "@" + domain,
			FuncInfo:   in.FuncInfo,
			Middleware: in.Middleware,
			Source:     in.Source,
		})
	}
}

// BindObject 为指定的模式绑定对象。
//
// BindObjectMethod和BindObject的区别：
// BindObjectMethod将对象中的指定方法与指定路由规则进行绑定，第三个method参数只能指定一个方法名称；
// BindObject注册时，所有的路由都是对象方法名称按照规则生成的，第三个methods参数可以指定多个注册的方法名称。
func (d *Domain) BindObject(pattern string, obj interface{}, methods ...string) {
	for domain := range d.domains {
		d.server.BindObject(pattern+"@"+domain, obj, methods...)
	}
}

func (d *Domain) doBindObject(ctx context.Context, in doBindObjectInput) {
	for domain := range d.domains {
		d.server.doBindObject(ctx, doBindObjectInput{
			Prefix:     in.Prefix,
			Pattern:    in.Pattern + "@" + domain,
			Object:     in.Object,
			Method:     in.Method,
			Middleware: in.Middleware,
			Source:     in.Source,
		})
	}
}

// BindObjectMethod 将指定模式的方法绑定。
// 
// BindObjectMethod和BindObject的区别：
// BindObjectMethod将对象中的指定方法与指定路由规则进行绑定，第三个method参数只能指定一个方法名称；
// BindObject注册时，所有的路由都是对象方法名称按照规则生成的，第三个methods参数可以指定多个注册的方法名称。
func (d *Domain) BindObjectMethod(pattern string, obj interface{}, method string) {
	for domain := range d.domains {
		d.server.BindObjectMethod(pattern+"@"+domain, obj, method)
	}
}

func (d *Domain) doBindObjectMethod(ctx context.Context, in doBindObjectMethodInput) {
	for domain := range d.domains {
		d.server.doBindObjectMethod(ctx, doBindObjectMethodInput{
			Prefix:     in.Prefix,
			Pattern:    in.Pattern + "@" + domain,
			Object:     in.Object,
			Method:     in.Method,
			Middleware: in.Middleware,
			Source:     in.Source,
		})
	}
}

// BindObjectRest 为指定模式绑定RESTful API。
// RESTful设计方式的控制器，通常用于API服务。
// 在这种模式下，HTTP的Method将会映射到控制器对应的方法名称，
// 例如：POST方式将会映射到控制器的Post方法中(公开方法，首字母大写)，DELETE方式将会映射到控制器的Delete方法中，以此类推。
// 其他非HTTP Method命名的方法，即使是定义的包公开方法，将不会自动注册，对于应用端不可见。
// 当然，如果控制器并未定义对应HTTP Method的方法，该Method请求下将会返回 HTTP Status 404。
func (d *Domain) BindObjectRest(pattern string, obj interface{}) {
	for domain := range d.domains {
		d.server.BindObjectRest(pattern+"@"+domain, obj)
	}
}

func (d *Domain) doBindObjectRest(ctx context.Context, in doBindObjectInput) {
	for domain := range d.domains {
		d.server.doBindObjectRest(ctx, doBindObjectInput{
			Prefix:     in.Prefix,
			Pattern:    in.Pattern + "@" + domain,
			Object:     in.Object,
			Method:     in.Method,
			Middleware: in.Middleware,
			Source:     in.Source,
		})
	}
}

// BindHookHandler 为指定模式绑定钩子处理器。
func (d *Domain) BindHookHandler(pattern string, hook HookName, handler HandlerFunc) {
	for domain := range d.domains {
		d.server.BindHookHandler(pattern+"@"+domain, hook, handler)
	}
}

func (d *Domain) doBindHookHandler(ctx context.Context, in doBindHookHandlerInput) {
	for domain := range d.domains {
		d.server.doBindHookHandler(ctx, doBindHookHandlerInput{
			Prefix:   in.Prefix,
			Pattern:  in.Pattern + "@" + domain,
			HookName: in.HookName,
			Handler:  in.Handler,
			Source:   in.Source,
		})
	}
}

// BindHookHandlerByMap 通过映射为特定模式绑定钩子处理器。
func (d *Domain) BindHookHandlerByMap(pattern string, hookMap map[HookName]HandlerFunc) {
	for domain := range d.domains {
		d.server.BindHookHandlerByMap(pattern+"@"+domain, hookMap)
	}
}

// BindStatusHandler 为指定模式绑定状态处理器。
func (d *Domain) BindStatusHandler(status int, handler HandlerFunc) {
	for domain := range d.domains {
		d.server.addStatusHandler(d.server.statusHandlerKey(status, domain), handler)
	}
}

// BindStatusHandlerByMap 通过给定的模式绑定状态处理器。
func (d *Domain) BindStatusHandlerByMap(handlerMap map[int]HandlerFunc) {
	for k, v := range handlerMap {
		d.BindStatusHandler(k, v)
	}
}

// BindMiddleware 为指定模式绑定中间件。
func (d *Domain) BindMiddleware(pattern string, handlers ...HandlerFunc) {
	for domain := range d.domains {
		d.server.BindMiddleware(pattern+"@"+domain, handlers...)
	}
}

// BindMiddlewareDefault 为指定模式绑定默认中间件。
func (d *Domain) BindMiddlewareDefault(handlers ...HandlerFunc) {
	for domain := range d.domains {
		d.server.BindMiddleware(defaultMiddlewarePattern+"@"+domain, handlers...)
	}
}

// Use 向域名添加中间件。
// Use 是 BindMiddlewareDefault 的别名。
// 请参阅 BindMiddlewareDefault。
func (d *Domain) Use(handlers ...HandlerFunc) {
	d.BindMiddlewareDefault(handlers...)
}
