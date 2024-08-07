// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"context"

	"github.com/888go/goframe/net/goai"
	gstr "github.com/888go/goframe/text/gstr"
)

// initOpenApi 使用OpenApiV3协议生成API规范。 md5:99e2f65c4ca9221a
func (s *X服务) initOpenApi() {
	if s.config.OpenApiPath == "" {
		return
	}
	var (
		ctx     = context.TODO()
		err     error
		methods []string
	)
	for _, item := range s.X取路由切片() {
		switch item.Type {
		case HandlerTypeMiddleware, HandlerTypeHook:
			continue
		}
		if item.Handler.X处理器函数信息.IsStrictRoute {
			methods = []string{item.Method}
			if gstr.X相等比较并忽略大小写(item.Method, defaultMethod) {
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

// openapiSpec 是一个内置处理器，用于自动生成 OpenAPI 规范的 JSON 文件。 md5:6f3f66a4ccde5784
func (s *X服务) openapiSpec(r *Request) {
	if s.config.OpenApiPath == "" {
		r.X响应.X写响应缓冲区(`OpenApi specification file producing is disabled`)
	} else {
		r.X响应.X写响应缓冲区JSON(s.openapi)
	}
}
