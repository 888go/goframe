// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"time"
	
	"github.com/888go/goframe/os/gsession"
)

// SetSessionMaxAge 设置服务器的 SessionMaxAge。
func (s *Server) X设置Session最大存活时长(时长 time.Duration) {
	s.config.SessionMaxAge = 时长
}

// SetSessionIdName 为服务器设置 SessionIdName。
func (s *Server) X设置SessionID名称(名称 string) {
	s.config.SessionIdName = 名称
}

// SetSessionStorage 为服务器设置 SessionStorage。
func (s *Server) X设置Session存储对象(Session存储对象 session类.Storage) {
	s.config.SessionStorage = Session存储对象
}

// SetSessionCookieOutput设置服务器的SetSessionCookieOutput。
func (s *Server) X设置SessionID输出到Cookie(开启 bool) {
	s.config.SessionCookieOutput = 开启
}

// SetSessionCookieMaxAge 为服务器设置 SessionCookieMaxAge。
func (s *Server) X设置SessionCookie存活时长(最大时长 time.Duration) {
	s.config.SessionCookieMaxAge = 最大时长
}

// GetSessionMaxAge 返回服务器的 SessionMaxAge 值。
func (s *Server) X取Session最大存活时长() time.Duration {
	return s.config.SessionMaxAge
}

// GetSessionIdName 返回服务器的 SessionIdName。
func (s *Server) X取SessionID名称() string {
	return s.config.SessionIdName
}

// GetSessionCookieMaxAge 返回服务器的 SessionCookieMaxAge 值。
func (s *Server) X取SessionCookie存活时长() time.Duration {
	return s.config.SessionCookieMaxAge
}
