// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"time"

	gsession "github.com/888go/goframe/os/gsession"
)

// X设置Session最大存活时长 用于设置服务器的会话最大存活时间。 md5:b9197f93d7981d91
func (s *X服务) X设置Session最大存活时长(时长 time.Duration) {
	s.config.SessionMaxAge = 时长
}

// X设置SessionID名称 为服务器设置 SessionIdName。 md5:d2b73f9cc3d5d247
func (s *X服务) X设置SessionID名称(名称 string) {
	s.config.SessionIdName = 名称
}

// X设置Session存储对象 为服务器设置SessionStorage。 md5:92b3b7fbc295084e
func (s *X服务) X设置Session存储对象(Session存储对象 gsession.Storage) {
	s.config.SessionStorage = Session存储对象
}

// X设置SessionID输出到Cookie 设置服务器的X设置SessionID输出到Cookie。 md5:a75d67eb511dd5ab
func (s *X服务) X设置SessionID输出到Cookie(开启 bool) {
	s.config.SessionCookieOutput = 开启
}

// X设置SessionCookie存活时长 为服务器设置 SessionCookieMaxAge。 md5:977671b06e3e5674
func (s *X服务) X设置SessionCookie存活时长(最大时长 time.Duration) {
	s.config.SessionCookieMaxAge = 最大时长
}

// X取Session最大存活时长 返回服务器的会话最大存活时间。 md5:3019bd154cdbe978
func (s *X服务) X取Session最大存活时长() time.Duration {
	return s.config.SessionMaxAge
}

// X取SessionID名称 返回服务器的SessionIdName。 md5:b0dffea716eef695
func (s *X服务) X取SessionID名称() string {
	return s.config.SessionIdName
}

// X取SessionCookie存活时长 返回服务器的 SessionCookieMaxAge。 md5:37a1e72edd958c1f
func (s *X服务) X取SessionCookie存活时长() time.Duration {
	return s.config.SessionCookieMaxAge
}
