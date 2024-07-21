// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"time"

	"github.com/gogf/gf/v2/os/gsession"
)

// SetSessionMaxAge 用于设置服务器的会话最大存活时间。 md5:b9197f93d7981d91
// ff:设置Session最大存活时长
// s:
// ttl:时长
func (s *Server) SetSessionMaxAge(ttl time.Duration) {
	s.config.SessionMaxAge = ttl
}

// SetSessionIdName 为服务器设置 SessionIdName。 md5:d2b73f9cc3d5d247
// ff:设置SessionID名称
// s:
// name:名称
func (s *Server) SetSessionIdName(name string) {
	s.config.SessionIdName = name
}

// SetSessionStorage 为服务器设置SessionStorage。 md5:92b3b7fbc295084e
// ff:设置Session存储对象
// s:
// storage:Session存储对象
func (s *Server) SetSessionStorage(storage gsession.Storage) {
	s.config.SessionStorage = storage
}

// SetSessionCookieOutput 设置服务器的SetSessionCookieOutput。 md5:a75d67eb511dd5ab
// ff:设置SessionID输出到Cookie
// s:
// enabled:开启
func (s *Server) SetSessionCookieOutput(enabled bool) {
	s.config.SessionCookieOutput = enabled
}

// SetSessionCookieMaxAge 为服务器设置 SessionCookieMaxAge。 md5:977671b06e3e5674
// ff:设置SessionCookie存活时长
// s:
// maxAge:最大时长
func (s *Server) SetSessionCookieMaxAge(maxAge time.Duration) {
	s.config.SessionCookieMaxAge = maxAge
}

// GetSessionMaxAge 返回服务器的会话最大存活时间。 md5:3019bd154cdbe978
// ff:取Session最大存活时长
// s:
func (s *Server) GetSessionMaxAge() time.Duration {
	return s.config.SessionMaxAge
}

// GetSessionIdName 返回服务器的SessionIdName。 md5:b0dffea716eef695
// ff:取SessionID名称
// s:
func (s *Server) GetSessionIdName() string {
	return s.config.SessionIdName
}

// GetSessionCookieMaxAge 返回服务器的 SessionCookieMaxAge。 md5:37a1e72edd958c1f
// ff:取SessionCookie存活时长
// s:
func (s *Server) GetSessionCookieMaxAge() time.Duration {
	return s.config.SessionCookieMaxAge
}
