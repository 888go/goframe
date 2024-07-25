// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package ghttp

// SetRewrite 为服务器设置静态URI的重写规则。 md5:de7f1cfb68c0440c
func (s *Server) SetRewrite(uri string, rewrite string) {
	s.config.Rewrites[uri] = rewrite
}

// SetRewriteMap 设置服务器的重写映射。 md5:61d51060723e66b8
func (s *Server) SetRewriteMap(rewrites map[string]string) {
	for k, v := range rewrites {
		s.config.Rewrites[k] = v
	}
}

// SetRouteOverWrite 设置服务器的路由覆盖。 md5:613439182bb14ec2
func (s *Server) SetRouteOverWrite(enabled bool) {
	s.config.RouteOverWrite = enabled
}
