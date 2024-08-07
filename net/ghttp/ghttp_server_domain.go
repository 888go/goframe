// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"context"
	"strings"
)

// Domain 用于域名的路由注册。 md5:c83832d3612fc101
type Domain struct {
	server  *X服务             // Belonged server
	domains map[string]struct{} // 支持多个域名。 md5:8ddb7825d8136c0d
}

// X创建域名路由 为一个或多个域创建并返回用于管理的域对象。 md5:613c2d9b16e14c56
func (s *X服务) X创建域名路由(域名 string) *Domain {
	d := &Domain{
		server:  s,
		domains: make(map[string]struct{}),
	}
	for _, v := range strings.Split(域名, ",") {
		d.domains[strings.TrimSpace(v)] = struct{}{}
	}
	return d
}

// X绑定 为指定的模式绑定处理器。 md5:91d1c2e239074063
func (d *Domain) X绑定(路由规则 string, 处理函数 interface{}) {
	for domain := range d.domains {
		d.server.X绑定(路由规则+"@"+domain, 处理函数)
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

// X绑定对象 将对象绑定到指定的模式。 md5:3beffc65c22e5098
func (d *Domain) X绑定对象(路由规则 string, 处理对象 interface{}, 方法名 ...string) {
	for domain := range d.domains {
		d.server.X绑定对象(路由规则+"@"+domain, 处理对象, 方法名...)
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

// X绑定对象方法 为指定的模式绑定方法。 md5:44230959de888ace
func (d *Domain) X绑定对象方法(路由规则 string, 处理对象 interface{}, 方法 string) {
	for domain := range d.domains {
		d.server.X绑定对象方法(路由规则+"@"+domain, 处理对象, 方法)
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

// X绑定RESTfulAPI对象 为指定模式绑定RESTful API。 md5:c63a87d6bb5ebd6c
func (d *Domain) X绑定RESTfulAPI对象(路由规则 string, 处理对象 interface{}) {
	for domain := range d.domains {
		d.server.X绑定RESTfulAPI对象(路由规则+"@"+domain, 处理对象)
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

// X绑定Hook 为指定的模式绑定钩子处理器。 md5:c2455777873cd92a
func (d *Domain) X绑定Hook(路由规则 string, 触发时机 Hook名称, 处理函数 HandlerFunc) {
	for domain := range d.domains {
		d.server.X绑定Hook(路由规则+"@"+domain, 触发时机, 处理函数)
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

// X绑定HookMap 为指定的模式绑定钩子处理程序。 md5:39439ccca98ce817
func (d *Domain) X绑定HookMap(路由规则 string, HookMap map[Hook名称]HandlerFunc) {
	for domain := range d.domains {
		d.server.X绑定HookMap(路由规则+"@"+domain, HookMap)
	}
}

// X绑定状态码中间件 为指定的模式绑定状态处理器。 md5:5cb032dd618649e2
func (d *Domain) X绑定状态码中间件(状态码 int, 处理函数 HandlerFunc) {
	for domain := range d.domains {
		d.server.addStatusHandler(d.server.statusHandlerKey(状态码, domain), 处理函数)
	}
}

// X绑定状态码中间件Map 为指定的模式绑定状态处理器。 md5:04388d53c4410f82
func (d *Domain) X绑定状态码中间件Map(中间件Map map[int]HandlerFunc) {
	for k, v := range 中间件Map {
		d.X绑定状态码中间件(k, v)
	}
}

// X绑定中间件 为指定的模式绑定中间件。 md5:40c97b890ebb105a
func (d *Domain) X绑定中间件(路由规则 string, 处理函数 ...HandlerFunc) {
	for domain := range d.domains {
		d.server.X绑定全局中间件(路由规则+"@"+domain, 处理函数...)
	}
}

// X绑定默认中间件 为指定的模式绑定默认中间件。 md5:792e7f694ab5eeb3
func (d *Domain) X绑定默认中间件(处理函数 ...HandlerFunc) {
	for domain := range d.domains {
		d.server.X绑定全局中间件(defaultMiddlewarePattern+"@"+domain, 处理函数...)
	}
}

// Use别名向域添加中间件。 md5:4aeb37c314d609f3
func (d *Domain) Use别名(处理函数 ...HandlerFunc) {
	d.X绑定默认中间件(处理函数...)
}
