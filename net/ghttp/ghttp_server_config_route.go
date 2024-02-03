// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

// SetRewrite 为服务器设置静态URI的重写规则。
func (s *Server) SetRewrite(uri string, rewrite string) {
	s.config.Rewrites[uri] = rewrite
}

// SetRewriteMap 为服务器设置重写映射。
func (s *Server) SetRewriteMap(rewrites map[string]string) {
	for k, v := range rewrites {
		s.config.Rewrites[k] = v
	}
}

// SetRouteOverWrite 为服务器设置 RouteOverWrite。
func (s *Server) SetRouteOverWrite(enabled bool) {
	s.config.RouteOverWrite = enabled
}
