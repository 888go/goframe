// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

// SetRewrite 为服务器设置静态URI的重写规则。
func (s *Server) X设置路由URI重写规则(URI string, 新URI string) {
	s.config.Rewrites[URI] = 新URI
}

// SetRewriteMap 为服务器设置重写映射。
func (s *Server) X设置路由URI重写规则Map(重写规则Map map[string]string) {
	for k, v := range 重写规则Map {
		s.config.Rewrites[k] = v
	}
}

// SetRouteOverWrite 为服务器设置 RouteOverWrite。
func (s *Server) X设置路由允许覆盖(开启 bool) {
	s.config.RouteOverWrite = 开启
}
