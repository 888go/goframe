// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

// SetNameToUriType 为服务器设置 NameToUriType。
func (s *X服务) SetNameToUriType(t int) {
	s.config.NameToUriType = t
}

// SetDumpRouterMap 为服务器设置 DumpRouterMap。
// 如果启用了 DumpRouterMap，服务器启动时会自动导出路由映射表。
func (s *X服务) SetDumpRouterMap(enabled bool) {
	s.config.DumpRouterMap = enabled
}

// SetClientMaxBodySize 为服务器设置客户端最大请求体大小。
func (s *X服务) X设置客户端请求最大长度(最大长度 int64) {
	s.config.X客户端请求最大长度 = 最大长度
}

// SetFormParsingMemory 用于设置服务器的表单解析内存。
func (s *X服务) X设置表单解析最大缓冲区长度(最大长度 int64) {
	s.config.X表单解析最大缓冲区长度 = 最大长度
}

// SetSwaggerPath 用于设置服务器的 SwaggerPath。
func (s *X服务) X设置APISwaggerUI路径(路径 string) {
	s.config.APISwaggerUI路径 = 路径
}

// SetOpenApiPath 为服务器设置 OpenApiPath。
func (s *X服务) X设置APIOpenApiUI路径(路径 string) {
	s.config.APIOpenApiUI路径 = 路径
}
