// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package ghttp

import (
	"context"

	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/text/gstr"
)

// initOpenApi 使用OpenApiV3协议生成API规范。 md5:99e2f65c4ca9221a
func (s *Server) initOpenApi() {
	if s.config.OpenApiPath == "" {
		return
	}
	var (
		ctx     = context.TODO()
		err     error
		methods []string
	)
	for _, item := range s.GetRoutes() {
		switch item.Type {
		case HandlerTypeMiddleware, HandlerTypeHook:
			continue
		}
		if item.Handler.Info.IsStrictRoute {
			methods = []string{item.Method}
			if gstr.Equal(item.Method, defaultMethod) {
				methods = SupportedMethods()
			}
			for _, method := range methods {
				err = s.openapi.Add(goai.AddInput{
					Path:   item.Route,
					Method: method,
					Object: item.Handler.Info.Value.Interface(),
				})
				if err != nil {
					s.Logger().Fatalf(ctx, `%+v`, err)
				}
			}
		}
	}
}

// openapiSpec 是一个内置处理器，用于自动生成 OpenAPI 规范的 JSON 文件。 md5:6f3f66a4ccde5784
func (s *Server) openapiSpec(r *Request) {
	if s.config.OpenApiPath == "" {
		r.Response.Write(`OpenApi specification file producing is disabled`)
	} else {
		r.Response.WriteJson(s.openapi)
	}
}
