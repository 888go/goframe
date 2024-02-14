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
func (s *X服务) X设置Cookie最大存活时长(时长 time.Duration) {
	s.config.Cookie最大存活时长 = 时长
}

// SetCookiePath 为服务器设置 CookiePath。
func (s *X服务) X设置Cookie路径(路径 string) {
	s.config.Cookie路径 = 路径
}

// SetCookieDomain 为服务器设置 CookieDomain。
func (s *X服务) X设置Cookie域名(域名 string) {
	s.config.Cookie域名 = 域名
}

// GetCookieMaxAge 返回服务器的 CookieMaxAge 值。
func (s *X服务) X取Cookie最大存活时长() time.Duration {
	return s.config.Cookie最大存活时长
}

// GetCookiePath 返回服务器的 CookiePath。
func (s *X服务) X取Cookie路径() string {
	return s.config.Cookie路径
}

// GetCookieDomain 返回服务器的 CookieDomain。
func (s *X服务) X取Cookie域名() string {
	return s.config.Cookie域名
}

// GetCookieSameSite 返回服务器的 CookieSameSite 值。
func (s *X服务) X取CookieSameSite() http.SameSite {
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

func (s *X服务) X取Cookie安全() bool {
	return s.config.Cookie安全
}

func (s *X服务) X取CookieHttpOnly() bool {
	return s.config.Cookie跨站访问控制
}
