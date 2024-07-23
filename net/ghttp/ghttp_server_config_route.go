// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

// SetRewrite sets rewrites for static URI for server.
// ff:设置路由URI重写规则
// s:
// uri:URI
// rewrite:新URI
func (s *Server) SetRewrite(uri string, rewrite string) {
	s.config.Rewrites[uri] = rewrite
}

// SetRewriteMap sets the rewritten map for server.
// ff:设置路由URI重写规则Map
// s:
// rewrites:重写规则Map
func (s *Server) SetRewriteMap(rewrites map[string]string) {
	for k, v := range rewrites {
		s.config.Rewrites[k] = v
	}
}

// SetRouteOverWrite sets the RouteOverWrite for server.
// ff:设置路由允许覆盖
// s:
// enabled:开启
func (s *Server) SetRouteOverWrite(enabled bool) {
	s.config.RouteOverWrite = enabled
}
