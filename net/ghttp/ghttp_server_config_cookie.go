// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"net/http"
	"time"
)

// X设置Cookie最大存活时长 设置服务器的CookieMaxAge。 md5:1bd691f49b3f7d7e
func (s *X服务) X设置Cookie最大存活时长(时长 time.Duration) {
	s.config.CookieMaxAge = 时长
}

// X设置Cookie路径 为服务器设置Cookie路径。 md5:427faaa8ad8633d8
func (s *X服务) X设置Cookie路径(路径 string) {
	s.config.CookiePath = 路径
}

// X设置Cookie域名 设置服务器的Cookie域。 md5:0337df246d313fe6
func (s *X服务) X设置Cookie域名(域名 string) {
	s.config.CookieDomain = 域名
}

// X取Cookie最大存活时长 返回服务器的 CookieMaxAge。 md5:515f434ef606a7c5
func (s *X服务) X取Cookie最大存活时长() time.Duration {
	return s.config.CookieMaxAge
}

// X取Cookie路径 返回服务器的CookiePath。 md5:5879f664e379d096
func (s *X服务) X取Cookie路径() string {
	return s.config.CookiePath
}

// X取Cookie域名 返回服务器的Cookie域名。 md5:f44dbe67d6622660
func (s *X服务) X取Cookie域名() string {
	return s.config.CookieDomain
}

// X取CookieSameSite 返回服务器的CookieSameSite设置。 md5:e015f6293556b8a8
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
	return s.config.CookieSecure
}

func (s *X服务) X取CookieHttpOnly() bool {
	return s.config.CookieHttpOnly
}
