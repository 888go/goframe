// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"fmt"
	
	"github.com/888go/goframe/text/gstr"
)

const (
	swaggerUIDocName            = `redoc.standalone.js`
	swaggerUIDocNamePlaceHolder = `{SwaggerUIDocName}`
	swaggerUIDocURLPlaceHolder  = `{SwaggerUIDocUrl}`
	swaggerUITemplate           = `
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
		<script src="https://unpkg.com/redoc@2.0.0-rc.70/bundles/redoc.standalone.js"> </script>
	</body>
</html>
`
)

// swaggerUI 是一个内置钩子处理器，用于将默认的Swagger JSON URL替换为本地OpenAPI JSON文件路径。
// 只有在启用了OpenAPI规范自动生成功能的配置时，此处理器才有意义。
func (s *Server) swaggerUI(r *Request) {
	if s.config.OpenApiPath == "" {
		return
	}
	if r.StaticFile != nil && r.StaticFile.File != nil && r.StaticFile.IsDir {
		content := 文本类.Map替换(swaggerUITemplate, map[string]string{
			swaggerUIDocURLPlaceHolder:  s.config.OpenApiPath,
			swaggerUIDocNamePlaceHolder: 文本类.X过滤尾字符并含空白(fmt.Sprintf(`//%s%s`, r.Host, r.Server.config.SwaggerPath), "/") + "/" + swaggerUIDocName,
		})
		r.Response.X写响应缓冲区(content)
		r.X退出全部()
	}
}
