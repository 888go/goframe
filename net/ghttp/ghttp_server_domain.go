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
	server  *Server             // Belonged server
	domains map[string]struct{} // 支持多个域名。 md5:8ddb7825d8136c0d
}

// Domain 为一个或多个域创建并返回用于管理的域对象。 md5:613c2d9b16e14c56
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

// BindHandler 为指定的模式绑定处理器。 md5:91d1c2e239074063
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

// BindObject 将对象绑定到指定的模式。 md5:3beffc65c22e5098
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

// BindObjectMethod 为指定的模式绑定方法。 md5:44230959de888ace
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

// BindObjectRest 为指定模式绑定RESTful API。 md5:c63a87d6bb5ebd6c
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

// BindHookHandler 为指定的模式绑定钩子处理器。 md5:c2455777873cd92a
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

// BindHookHandlerByMap 为指定的模式绑定钩子处理程序。 md5:39439ccca98ce817
func (d *Domain) BindHookHandlerByMap(pattern string, hookMap map[HookName]HandlerFunc) {
	for domain := range d.domains {
		d.server.BindHookHandlerByMap(pattern+"@"+domain, hookMap)
	}
}

// BindStatusHandler 为指定的模式绑定状态处理器。 md5:5cb032dd618649e2
func (d *Domain) BindStatusHandler(status int, handler HandlerFunc) {
	for domain := range d.domains {
		d.server.addStatusHandler(d.server.statusHandlerKey(status, domain), handler)
	}
}

// BindStatusHandlerByMap 为指定的模式绑定状态处理器。 md5:04388d53c4410f82
func (d *Domain) BindStatusHandlerByMap(handlerMap map[int]HandlerFunc) {
	for k, v := range handlerMap {
		d.BindStatusHandler(k, v)
	}
}

// BindMiddleware 为指定的模式绑定中间件。 md5:40c97b890ebb105a
func (d *Domain) BindMiddleware(pattern string, handlers ...HandlerFunc) {
	for domain := range d.domains {
		d.server.BindMiddleware(pattern+"@"+domain, handlers...)
	}
}

// BindMiddlewareDefault 为指定的模式绑定默认中间件。 md5:792e7f694ab5eeb3
func (d *Domain) BindMiddlewareDefault(handlers ...HandlerFunc) {
	for domain := range d.domains {
		d.server.BindMiddleware(defaultMiddlewarePattern+"@"+domain, handlers...)
	}
}

// Use向域添加中间件。 md5:4aeb37c314d609f3
func (d *Domain) Use(handlers ...HandlerFunc) {
	d.BindMiddlewareDefault(handlers...)
}
