// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/debug/gdebug"
)

const (
	// The default route pattern for global middleware.
	defaultMiddlewarePattern = "/*"
)

// BindMiddleware registers one or more global middleware to the server.
// Global middleware can be used standalone without service handler, which intercepts all dynamic requests
// before or after service handler. The parameter `pattern` specifies what route pattern the middleware intercepts,
// which is usually a "fuzzy" pattern like "/:name", "/*any" or "/{field}".

// ff:绑定全局中间件
// handlers:处理函数
// pattern:路由规则
func (s *Server) BindMiddleware(pattern string, handlers ...HandlerFunc) {
	var (
		ctx = context.TODO()
	)
	for _, handler := range handlers {
		s.setHandler(ctx, setHandlerInput{
			Prefix:  "",
			Pattern: pattern,
			HandlerItem: &HandlerItem{
				Type: HandlerTypeMiddleware,
				Name: gdebug.FuncPath(handler),
				Info: handlerFuncInfo{
					Func: handler,
					Type: reflect.TypeOf(handler),
				},
			},
		})
	}
}

// BindMiddlewareDefault registers one or more global middleware to the server using default pattern "/*".
// Global middleware can be used standalone without service handler, which intercepts all dynamic requests
// before or after service handler.

// ff:绑定全局默认中间件
// handlers:处理函数
func (s *Server) BindMiddlewareDefault(handlers ...HandlerFunc) {
	var (
		ctx = context.TODO()
	)
	for _, handler := range handlers {
		s.setHandler(ctx, setHandlerInput{
			Prefix:  "",
			Pattern: defaultMiddlewarePattern,
			HandlerItem: &HandlerItem{
				Type: HandlerTypeMiddleware,
				Name: gdebug.FuncPath(handler),
				Info: handlerFuncInfo{
					Func: handler,
					Type: reflect.TypeOf(handler),
				},
			},
		})
	}
}

// Use is the alias of BindMiddlewareDefault.
// See BindMiddlewareDefault.

// ff:Use别名
// handlers:处理函数
func (s *Server) Use(handlers ...HandlerFunc) {
	s.BindMiddlewareDefault(handlers...)
}
