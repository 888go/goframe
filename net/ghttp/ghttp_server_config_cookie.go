// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"net/http"
	"time"
)

// SetCookieMaxAge 设置服务器的CookieMaxAge。 md5:1bd691f49b3f7d7e
func (s *Server) SetCookieMaxAge(ttl time.Duration) {
	s.config.CookieMaxAge = ttl
}

// SetCookiePath 为服务器设置Cookie路径。 md5:427faaa8ad8633d8
func (s *Server) SetCookiePath(path string) {
	s.config.CookiePath = path
}

// SetCookieDomain 设置服务器的Cookie域。 md5:0337df246d313fe6
func (s *Server) SetCookieDomain(domain string) {
	s.config.CookieDomain = domain
}

// GetCookieMaxAge 返回服务器的 CookieMaxAge。 md5:515f434ef606a7c5
func (s *Server) GetCookieMaxAge() time.Duration {
	return s.config.CookieMaxAge
}

// GetCookiePath 返回服务器的CookiePath。 md5:5879f664e379d096
func (s *Server) GetCookiePath() string {
	return s.config.CookiePath
}

// GetCookieDomain 返回服务器的Cookie域名。 md5:f44dbe67d6622660
func (s *Server) GetCookieDomain() string {
	return s.config.CookieDomain
}

// GetCookieSameSite 返回服务器的CookieSameSite设置。 md5:e015f6293556b8a8
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

func (s *Server) GetCookieSecure() bool {
	return s.config.CookieSecure
}

func (s *Server) GetCookieHttpOnly() bool {
	return s.config.CookieHttpOnly
}
