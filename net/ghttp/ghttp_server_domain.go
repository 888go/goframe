// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"context"
	"strings"
)

// Domain is used for route register for domains.
type Domain struct {
	server  *Server             // Belonged server
	domains map[string]struct{} // Support multiple domains.
}

// Domain creates and returns a domain object for management for one or more domains.

// ff:创建域名路由
// domains:域名
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

// BindHandler binds the handler for the specified pattern.

// ff:X绑定
// handler:处理函数
// pattern:路由规则
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

// BindObject binds the object for the specified pattern.

// ff:绑定对象
// methods:方法名
// obj:处理对象
// pattern:路由规则
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

// BindObjectMethod binds the method for the specified pattern.

// ff:绑定对象方法
// method:方法
// obj:处理对象
// pattern:路由规则
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

// BindObjectRest binds the RESTful API for the specified pattern.

// ff:绑定RESTfulAPI对象
// obj:处理对象
// pattern:路由规则
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

// BindHookHandler binds the hook handler for the specified pattern.

// ff:绑定Hook
// handler:处理函数
// hook:触发时机
// pattern:路由规则
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

// BindHookHandlerByMap binds the hook handler for the specified pattern.

// ff:绑定HookMap
// hookMap:HookMap
// pattern:路由规则
func (d *Domain) BindHookHandlerByMap(pattern string, hookMap map[HookName]HandlerFunc) {
	for domain := range d.domains {
		d.server.BindHookHandlerByMap(pattern+"@"+domain, hookMap)
	}
}

// BindStatusHandler binds the status handler for the specified pattern.

// ff:绑定状态码中间件
// handler:处理函数
// status:状态码
func (d *Domain) BindStatusHandler(status int, handler HandlerFunc) {
	for domain := range d.domains {
		d.server.addStatusHandler(d.server.statusHandlerKey(status, domain), handler)
	}
}

// BindStatusHandlerByMap binds the status handler for the specified pattern.

// ff:绑定状态码中间件Map
// handlerMap:中间件Map
func (d *Domain) BindStatusHandlerByMap(handlerMap map[int]HandlerFunc) {
	for k, v := range handlerMap {
		d.BindStatusHandler(k, v)
	}
}

// BindMiddleware binds the middleware for the specified pattern.

// ff:绑定中间件
// handlers:处理函数
// pattern:路由规则
func (d *Domain) BindMiddleware(pattern string, handlers ...HandlerFunc) {
	for domain := range d.domains {
		d.server.BindMiddleware(pattern+"@"+domain, handlers...)
	}
}

// BindMiddlewareDefault binds the default middleware for the specified pattern.

// ff:绑定默认中间件
// handlers:处理函数
func (d *Domain) BindMiddlewareDefault(handlers ...HandlerFunc) {
	for domain := range d.domains {
		d.server.BindMiddleware(defaultMiddlewarePattern+"@"+domain, handlers...)
	}
}

// Use adds middleware to the domain.

// ff:Use别名
// handlers:处理函数
func (d *Domain) Use(handlers ...HandlerFunc) {
	d.BindMiddlewareDefault(handlers...)
}
