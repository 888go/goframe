// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"net/http"
	"time"
	)
// SetCookieMaxAge 设置服务器的 CookieMaxAge。
func (s *Server) SetCookieMaxAge(ttl time.Duration) {
	s.config.CookieMaxAge = ttl
}

// SetCookiePath 为服务器设置 CookiePath。
func (s *Server) SetCookiePath(path string) {
	s.config.CookiePath = path
}

// SetCookieDomain 为服务器设置 CookieDomain。
func (s *Server) SetCookieDomain(domain string) {
	s.config.CookieDomain = domain
}

// GetCookieMaxAge 返回服务器的 CookieMaxAge 值。
func (s *Server) GetCookieMaxAge() time.Duration {
	return s.config.CookieMaxAge
}

// GetCookiePath 返回服务器的 CookiePath。
func (s *Server) GetCookiePath() string {
	return s.config.CookiePath
}

// GetCookieDomain 返回服务器的 CookieDomain。
func (s *Server) GetCookieDomain() string {
	return s.config.CookieDomain
}

// GetCookieSameSite 返回服务器的 CookieSameSite 值。
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
