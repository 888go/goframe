// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"net/http"
	"time"
)

// SetCookieMaxAge sets the CookieMaxAge for server.

// ff:设置Cookie最大存活时长
// ttl:时长
func (s *Server) SetCookieMaxAge(ttl time.Duration) {
	s.config.CookieMaxAge = ttl
}

// SetCookiePath sets the CookiePath for server.

// ff:设置Cookie路径
// path:路径
func (s *Server) SetCookiePath(path string) {
	s.config.CookiePath = path
}

// SetCookieDomain sets the CookieDomain for server.

// ff:设置Cookie域名
// domain:域名
func (s *Server) SetCookieDomain(domain string) {
	s.config.CookieDomain = domain
}

// GetCookieMaxAge returns the CookieMaxAge of the server.

// ff:取Cookie最大存活时长
func (s *Server) GetCookieMaxAge() time.Duration {
	return s.config.CookieMaxAge
}

// GetCookiePath returns the CookiePath of server.

// ff:取Cookie路径
func (s *Server) GetCookiePath() string {
	return s.config.CookiePath
}

// GetCookieDomain returns CookieDomain of server.

// ff:取Cookie域名
func (s *Server) GetCookieDomain() string {
	return s.config.CookieDomain
}

// GetCookieSameSite return CookieSameSite of server.

// ff:取CookieSameSite
func (s *Server) GetCookieSameSite() http.SameSite {
	switch s.config.CookieSameSite {
	case "lax":
		return http.SameSiteLaxMode
	case "none":
		return http.SameSiteNoneMode
	case "strict":
		return http.SameSiteStrictMode
	default:
		return http.SameSiteDefaultMode
	}
}


// ff:取Cookie安全
func (s *Server) GetCookieSecure() bool {
	return s.config.CookieSecure
}


// ff:取CookieHttpOnly
func (s *Server) GetCookieHttpOnly() bool {
	return s.config.CookieHttpOnly
}
