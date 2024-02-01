// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

// SetNameToUriType 为服务器设置 NameToUriType。
func (s *Server) SetNameToUriType(t int) {
	s.config.NameToUriType = t
}

// SetDumpRouterMap 为服务器设置 DumpRouterMap。
// 如果启用了 DumpRouterMap，服务器启动时会自动导出路由映射表。
func (s *Server) SetDumpRouterMap(enabled bool) {
	s.config.DumpRouterMap = enabled
}

// SetClientMaxBodySize 为服务器设置客户端最大请求体大小。
func (s *Server) SetClientMaxBodySize(maxSize int64) {
	s.config.ClientMaxBodySize = maxSize
}

// SetFormParsingMemory 用于设置服务器的表单解析内存。
func (s *Server) SetFormParsingMemory(maxMemory int64) {
	s.config.FormParsingMemory = maxMemory
}

// SetSwaggerPath 用于设置服务器的 SwaggerPath。
func (s *Server) SetSwaggerPath(path string) {
	s.config.SwaggerPath = path
}

// SetOpenApiPath 为服务器设置 OpenApiPath。
func (s *Server) SetOpenApiPath(path string) {
	s.config.OpenApiPath = path
}
