// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"github.com/gogf/gf/v2/text/gstr"
)

const (
	swaggerUIDocURLPlaceHolder = `{SwaggerUIDocUrl}`
	swaggerUITemplate          = `
<!DOCTYPE html>
<html>
	<head>
	<title>API Reference</title>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		body {
			margin:  0;
			padding: 0;
		}
	</style>
	</head>
	<body>
		<redoc spec-url="{SwaggerUIDocUrl}" show-object-schema-examples="true"></redoc>
		<script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"> </script>
	</body>
</html>
`
)

	// swaggerUI 是一个内置的钩子处理器，用于将默认的 Swagger JSON URL 替换为本地 OpenAPI JSON 文件路径。
	// 该处理器仅在启用了 OpenAPI 规范自动生成配置的情况下才有意义。
	// md5:7110c29f8ea820f7
func (s *Server) swaggerUI(r *Request) {
	if s.config.OpenApiPath == "" {
		return
	}
	var templateContent = swaggerUITemplate
	if s.config.SwaggerUITemplate != "" {
		templateContent = s.config.SwaggerUITemplate
	}

	if r.StaticFile != nil && r.StaticFile.File != nil && r.StaticFile.IsDir {
		content := gstr.ReplaceByMap(templateContent, map[string]string{
			swaggerUIDocURLPlaceHolder: s.config.OpenApiPath,
		})
		r.Response.Write(content)
		r.ExitAll()
	}
}
