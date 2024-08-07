// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	glog "github.com/888go/goframe/os/glog"
)

// X设置日志存储目录 设置服务器的日志路径。
// 只有当设置了日志路径时，才会将内容写入文件。
// md5:a5143c3c45534bef
func (s *X服务) X设置日志存储目录(目录 string) error {
	if len(目录) == 0 {
		return nil
	}
	s.config.LogPath = 目录
	s.config.ErrorLogEnabled = true
	s.config.AccessLogEnabled = true
	if s.config.LogPath != "" && s.config.LogPath != s.config.Logger.X取文件路径() {
		if err := s.config.Logger.X设置文件路径(s.config.LogPath); err != nil {
			return err
		}
	}
	return nil
}

// X设置日志记录器 用于设置负责日志记录的logger。
// 注意，由于可能存在并发安全问题，因此无法在运行时进行设置。
// md5:560d266f79fb0915
func (s *X服务) X设置日志记录器(日志记录器 *glog.Logger) {
	s.config.Logger = 日志记录器
}

// Logger别名 是 GetLogger 的别名。 md5:56065d1ce5cca4c9
func (s *X服务) Logger别名() *glog.Logger {
	return s.config.Logger
}

// X设置日志开启记录等级 通过level字符串设置日志级别。 md5:c479bcdf03ab0fa2
func (s *X服务) X设置日志开启记录等级(等级 string) {
	s.config.LogLevel = 等级
}

// X设置日志开启输出到CMD 设置是否将日志内容输出到stdout。 md5:c93557220e40f70b
func (s *X服务) X设置日志开启输出到CMD(开启 bool) {
	s.config.LogStdout = 开启
}

// X设置日志开启访客记录 用于启用或禁用访问日志。 md5:a353da90e0c3de0d
func (s *X服务) X设置日志开启访客记录(开启 bool) {
	s.config.AccessLogEnabled = 开启
}

// X设置日志开启错误记录 开启或关闭错误日志。 md5:38a0655c0083b6d0
func (s *X服务) X设置日志开启错误记录(开启 bool) {
	s.config.ErrorLogEnabled = 开启
}

// X设置日志开启错误堆栈记录 启用/禁用错误堆栈功能。 md5:6c411a957b96e186
func (s *X服务) X设置日志开启错误堆栈记录(开启 bool) {
	s.config.ErrorStack = 开启
}

// X取日志存储目录 返回日志路径。 md5:0fe087f2f9b0a123
func (s *X服务) X取日志存储目录() string {
	return s.config.LogPath
}

// X日志访客记录是否已开启 检查访问日志是否已启用。 md5:b076cc230602118d
func (s *X服务) X日志访客记录是否已开启() bool {
	return s.config.AccessLogEnabled
}

// X日志错误记录是否已开启 检查错误日志是否启用。 md5:2231c72f34764e99
func (s *X服务) X日志错误记录是否已开启() bool {
	return s.config.ErrorLogEnabled
}
