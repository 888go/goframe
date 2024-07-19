// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

// SetNameToUriType sets the NameToUriType for server.
// ff:
// s:
// t:
func (s *Server) SetNameToUriType(t int) {
	s.config.NameToUriType = t
}

// SetDumpRouterMap sets the DumpRouterMap for server.
// If DumpRouterMap is enabled, it automatically dumps the route map when server starts.
// ff:
// s:
// enabled:
func (s *Server) SetDumpRouterMap(enabled bool) {
	s.config.DumpRouterMap = enabled
}

// SetClientMaxBodySize sets the ClientMaxBodySize for server.
// ff:设置客户端请求最大长度
// s:
// maxSize:最大长度
func (s *Server) SetClientMaxBodySize(maxSize int64) {
	s.config.ClientMaxBodySize = maxSize
}

// SetFormParsingMemory sets the FormParsingMemory for server.
// ff:设置表单解析最大缓冲区长度
// s:
// maxMemory:最大长度
func (s *Server) SetFormParsingMemory(maxMemory int64) {
	s.config.FormParsingMemory = maxMemory
}

// SetSwaggerPath sets the SwaggerPath for server.
// ff:设置APISwaggerUI路径
// s:
// path:路径
func (s *Server) SetSwaggerPath(path string) {
	s.config.SwaggerPath = path
}

// SetSwaggerUITemplate sets the Swagger template for server.
// ff:
// s:
// swaggerUITemplate:
func (s *Server) SetSwaggerUITemplate(swaggerUITemplate string) {
	s.config.SwaggerUITemplate = swaggerUITemplate
}

// SetOpenApiPath sets the OpenApiPath for server.
// ff:设置APIOpenApiUI路径
// s:
// path:路径
func (s *Server) SetOpenApiPath(path string) {
	s.config.OpenApiPath = path
}
