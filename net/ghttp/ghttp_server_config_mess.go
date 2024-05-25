// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

// SetNameToUriType 为服务器设置NameToUriType。. md5:aa74886d065469f2
func (s *Server) SetNameToUriType(t int) {
	s.config.NameToUriType = t
}

// SetDumpRouterMap 为服务器设置DumpRouterMap。如果启用了DumpRouterMap，服务器启动时会自动dump路由映射。
// md5:1570556de016c76c
func (s *Server) SetDumpRouterMap(enabled bool) {
	s.config.DumpRouterMap = enabled
}

// SetClientMaxBodySize 为服务器设置客户端最大请求体大小。. md5:5aa4d07a0684f2f7
func (s *Server) SetClientMaxBodySize(maxSize int64) {
	s.config.ClientMaxBodySize = maxSize
}

// SetFormParsingMemory 设置服务器的表单解析内存大小。. md5:fad9dd48cd138769
func (s *Server) SetFormParsingMemory(maxMemory int64) {
	s.config.FormParsingMemory = maxMemory
}

// SetSwaggerPath 设置服务器的SwaggerPath。. md5:08e17ed1ebc823b0
func (s *Server) SetSwaggerPath(path string) {
	s.config.SwaggerPath = path
}

// SetSwaggerUITemplate 设置服务器的Swagger模板。. md5:c49e80113f25e335
func (s *Server) SetSwaggerUITemplate(swaggerUITemplate string) {
	s.config.SwaggerUITemplate = swaggerUITemplate
}

// SetOpenApiPath 为服务器设置OpenApiPath。. md5:c6ced131768ea48a
func (s *Server) SetOpenApiPath(path string) {
	s.config.OpenApiPath = path
}
