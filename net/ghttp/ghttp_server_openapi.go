// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"context"
	
	"github.com/888go/goframe/net/goai"
	"github.com/888go/goframe/text/gstr"
)

// initOpenApi 使用 OpenApiV3 协议生成 API 规范。
func (s *Server) initOpenApi() {
	if s.config.OpenApiPath == "" {
		return
	}
	var (
		ctx     = context.TODO()
		err     error
		methods []string
	)
	for _, item := range s.X取路由数组() {
		switch item.Type {
		case HandlerTypeMiddleware, HandlerTypeHook:
			continue
		}
		if item.Handler.Info.IsStrictRoute {
			methods = []string{item.Method}
			if 文本类.X相等比较并忽略大小写(item.Method, defaultMethod) {
				methods = X取所支持的HTTP方法()
			}
			for _, method := range methods {
				err = s.openapi.Add(goai.AddInput{
					Path:   item.Route,
					Method: method,
					Object: item.Handler.Info.Value.Interface(),
				})
				if err != nil {
					s.Logger别名().X输出并格式化FATA(ctx, `%+v`, err)
				}
			}
		}
	}
}

// openapiSpec 是一个内置处理器，用于自动生成 openapi 规范的 JSON 文件。
func (s *Server) openapiSpec(r *Request) {
	if s.config.OpenApiPath == "" {
		r.Response.X写响应缓冲区(`OpenApi specification file producing is disabled`)
	} else {
		r.Response.X写响应缓冲区JSON(s.openapi)
	}
}
