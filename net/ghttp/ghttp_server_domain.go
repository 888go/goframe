// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"context"
	"strings"
)

// Domain 用于为域名进行路由注册。
type X域名路由 struct {
	server  *X服务             // Belonged server
	domains map[string]struct{} // 支持多个域名。
}

// Domain 创建并返回一个域名管理对象，用于管理一个或多个域名。
func (s *X服务) X创建域名路由(域名 string) *X域名路由 {
	d := &X域名路由{
		server:  s,
		domains: make(map[string]struct{}),
	}
	for _, v := range strings.Split(域名, ",") {
		d.domains[strings.TrimSpace(v)] = struct{}{}
	}
	return d
}

// BindHandler 为指定模式绑定处理器。
func (d *X域名路由) X绑定(路由规则 string, 处理函数 interface{}) {
	for domain := range d.domains {
		d.server.X绑定(路由规则+"@"+domain, 处理函数)
	}
}

func (d *X域名路由) doBindHandler(ctx context.Context, in doBindHandlerInput) {
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
func (d *X域名路由) X绑定对象(路由规则 string, 处理对象 interface{}, 方法名 ...string) {
	for domain := range d.domains {
		d.server.X绑定对象(路由规则+"@"+domain, 处理对象, 方法名...)
	}
}

func (d *X域名路由) doBindObject(ctx context.Context, in doBindObjectInput) {
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
func (d *X域名路由) X绑定对象方法(路由规则 string, 处理对象 interface{}, 方法 string) {
	for domain := range d.domains {
		d.server.X绑定对象方法(路由规则+"@"+domain, 处理对象, 方法)
	}
}

func (d *X域名路由) doBindObjectMethod(ctx context.Context, in doBindObjectMethodInput) {
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
func (d *X域名路由) X绑定RESTfulAPI对象(路由规则 string, 处理对象 interface{}) {
	for domain := range d.domains {
		d.server.X绑定RESTfulAPI对象(路由规则+"@"+domain, 处理对象)
	}
}

func (d *X域名路由) doBindObjectRest(ctx context.Context, in doBindObjectInput) {
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
func (d *X域名路由) X绑定Hook(路由规则 string, 触发时机 Hook名称, 处理函数 HandlerFunc) {
	for domain := range d.domains {
		d.server.X绑定Hook(路由规则+"@"+domain, 触发时机, 处理函数)
	}
}

func (d *X域名路由) doBindHookHandler(ctx context.Context, in doBindHookHandlerInput) {
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
func (d *X域名路由) X绑定HookMap(路由规则 string, HookMap map[Hook名称]HandlerFunc) {
	for domain := range d.domains {
		d.server.X绑定HookMap(路由规则+"@"+domain, HookMap)
	}
}

// BindStatusHandler 为指定模式绑定状态处理器。
func (d *X域名路由) X绑定状态码中间件(状态码 int, 处理函数 HandlerFunc) {
	for domain := range d.domains {
		d.server.addStatusHandler(d.server.statusHandlerKey(状态码, domain), 处理函数)
	}
}

// BindStatusHandlerByMap 通过给定的模式绑定状态处理器。
func (d *X域名路由) X绑定状态码中间件Map(中间件Map map[int]HandlerFunc) {
	for k, v := range 中间件Map {
		d.X绑定状态码中间件(k, v)
	}
}

// BindMiddleware 为指定模式绑定中间件。
func (d *X域名路由) X绑定中间件(路由规则 string, 处理函数 ...HandlerFunc) {
	for domain := range d.domains {
		d.server.X绑定全局中间件(路由规则+"@"+domain, 处理函数...)
	}
}

// BindMiddlewareDefault 为指定模式绑定默认中间件。
func (d *X域名路由) X绑定默认中间件(处理函数 ...HandlerFunc) {
	for domain := range d.domains {
		d.server.X绑定全局中间件(defaultMiddlewarePattern+"@"+domain, 处理函数...)
	}
}

// Use 向域名添加中间件。
// Use 是 BindMiddlewareDefault 的别名。
// 请参阅 BindMiddlewareDefault。
func (d *X域名路由) Use别名(处理函数 ...HandlerFunc) {
	d.X绑定默认中间件(处理函数...)
}
