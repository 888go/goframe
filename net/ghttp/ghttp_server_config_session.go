// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"time"
	
	"github.com/888go/goframe/os/gsession"
	)
// SetSessionMaxAge 设置服务器的 SessionMaxAge。
func (s *Server) SetSessionMaxAge(ttl time.Duration) {
	s.config.SessionMaxAge = ttl
}

// SetSessionIdName 为服务器设置 SessionIdName。
func (s *Server) SetSessionIdName(name string) {
	s.config.SessionIdName = name
}

// SetSessionStorage 为服务器设置 SessionStorage。
func (s *Server) SetSessionStorage(storage gsession.Storage) {
	s.config.SessionStorage = storage
}

// SetSessionCookieOutput设置服务器的SetSessionCookieOutput。
func (s *Server) SetSessionCookieOutput(enabled bool) {
	s.config.SessionCookieOutput = enabled
}

// SetSessionCookieMaxAge 为服务器设置 SessionCookieMaxAge。
func (s *Server) SetSessionCookieMaxAge(maxAge time.Duration) {
	s.config.SessionCookieMaxAge = maxAge
}

// GetSessionMaxAge 返回服务器的 SessionMaxAge 值。
func (s *Server) GetSessionMaxAge() time.Duration {
	return s.config.SessionMaxAge
}

// GetSessionIdName 返回服务器的 SessionIdName。
func (s *Server) GetSessionIdName() string {
	return s.config.SessionIdName
}

// GetSessionCookieMaxAge 返回服务器的 SessionCookieMaxAge 值。
func (s *Server) GetSessionCookieMaxAge() time.Duration {
	return s.config.SessionCookieMaxAge
}
