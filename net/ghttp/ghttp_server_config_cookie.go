// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"net/http"
	"time"
)

// SetCookieMaxAge 设置服务器的 CookieMaxAge。
func (s *Server) X设置Cookie最大存活时长(时长 time.Duration) {
	s.config.CookieMaxAge = 时长
}

// SetCookiePath 为服务器设置 CookiePath。
func (s *Server) X设置Cookie路径(路径 string) {
	s.config.CookiePath = 路径
}

// SetCookieDomain 为服务器设置 CookieDomain。
func (s *Server) X设置Cookie域名(域名 string) {
	s.config.CookieDomain = 域名
}

// GetCookieMaxAge 返回服务器的 CookieMaxAge 值。
func (s *Server) X取Cookie最大存活时长() time.Duration {
	return s.config.CookieMaxAge
}

// GetCookiePath 返回服务器的 CookiePath。
func (s *Server) X取Cookie路径() string {
	return s.config.CookiePath
}

// GetCookieDomain 返回服务器的 CookieDomain。
func (s *Server) X取Cookie域名() string {
	return s.config.CookieDomain
}

// GetCookieSameSite 返回服务器的 CookieSameSite 值。
func (s *Server) X取CookieSameSite() http.SameSite {
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

func (s *Server) X取Cookie安全() bool {
	return s.config.CookieSecure
}

func (s *Server) X取CookieHttpOnly() bool {
	return s.config.CookieHttpOnly
}
