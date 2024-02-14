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
func (s *X服务) initOpenApi() {
	if s.config.APIOpenApiUI路径 == "" {
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
		if item.Handler.X处理器函数信息.IsStrictRoute {
			methods = []string{item.Method}
			if 文本类.X相等比较并忽略大小写(item.Method, defaultMethod) {
				methods = X取所支持的HTTP方法()
			}
			for _, method := range methods {
				err = s.openapi.Add(goai.AddInput{
					Path:   item.X路由URI,
					Method: method,
					Object: item.Handler.X处理器函数信息.Value.Interface(),
				})
				if err != nil {
					s.Logger别名().X输出并格式化FATA(ctx, `%+v`, err)
				}
			}
		}
	}
}

// openapiSpec 是一个内置处理器，用于自动生成 openapi 规范的 JSON 文件。
func (s *X服务) openapiSpec(r *X请求) {
	if s.config.APIOpenApiUI路径 == "" {
		r.X响应.X写响应缓冲区(`OpenApi specification file producing is disabled`)
	} else {
		r.X响应.X写响应缓冲区JSON(s.openapi)
	}
}
