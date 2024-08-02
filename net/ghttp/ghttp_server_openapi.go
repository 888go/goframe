// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package http类

import (
	"context"

	"github.com/888go/goframe/net/goai"
	gstr "github.com/888go/goframe/text/gstr"
)

// initOpenApi generates api specification using OpenApiV3 protocol.
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

// openapiSpec is a build-in handler automatic producing for openapi specification json file.
func (s *Server) openapiSpec(r *Request) {
	if s.config.OpenApiPath == "" {
		r.Response.Write(`OpenApi specification file producing is disabled`)
	} else {
		r.Response.WriteJson(s.openapi)
	}
}
