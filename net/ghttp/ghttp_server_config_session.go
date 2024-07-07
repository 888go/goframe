// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"time"

	"github.com/gogf/gf/v2/os/gsession"
)

// SetSessionMaxAge sets the SessionMaxAge for server.
// ff:设置Session最大存活时长
// s:
// ttl:时长
func (s *Server) SetSessionMaxAge(ttl time.Duration) {
	s.config.SessionMaxAge = ttl
}

// SetSessionIdName sets the SessionIdName for server.
// ff:设置SessionID名称
// s:
// name:名称
func (s *Server) SetSessionIdName(name string) {
	s.config.SessionIdName = name
}

// SetSessionStorage sets the SessionStorage for server.
// ff:设置Session存储对象
// s:
// storage:Session存储对象
func (s *Server) SetSessionStorage(storage gsession.Storage) {
	s.config.SessionStorage = storage
}

// SetSessionCookieOutput sets the SetSessionCookieOutput for server.
// ff:设置SessionID输出到Cookie
// s:
// enabled:开启
func (s *Server) SetSessionCookieOutput(enabled bool) {
	s.config.SessionCookieOutput = enabled
}

// SetSessionCookieMaxAge sets the SessionCookieMaxAge for server.
// ff:设置SessionCookie存活时长
// s:
// maxAge:最大时长
func (s *Server) SetSessionCookieMaxAge(maxAge time.Duration) {
	s.config.SessionCookieMaxAge = maxAge
}

// GetSessionMaxAge returns the SessionMaxAge of server.
// ff:取Session最大存活时长
// s:
func (s *Server) GetSessionMaxAge() time.Duration {
	return s.config.SessionMaxAge
}

// GetSessionIdName returns the SessionIdName of server.
// ff:取SessionID名称
// s:
func (s *Server) GetSessionIdName() string {
	return s.config.SessionIdName
}

// GetSessionCookieMaxAge returns the SessionCookieMaxAge of server.
// ff:取SessionCookie存活时长
// s:
func (s *Server) GetSessionCookieMaxAge() time.Duration {
	return s.config.SessionCookieMaxAge
}
