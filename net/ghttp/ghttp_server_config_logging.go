// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import "github.com/gogf/gf/v2/os/glog"

// SetLogPath 设置服务器的日志路径。
// 只有当设置了日志路径时，才会将内容写入文件。
// md5:a5143c3c45534bef
// ff:设置日志存储目录
// s:
// path:目录
func (s *Server) SetLogPath(path string) error {
	if len(path) == 0 {
		return nil
	}
	s.config.LogPath = path
	s.config.ErrorLogEnabled = true
	s.config.AccessLogEnabled = true
	if s.config.LogPath != "" && s.config.LogPath != s.config.Logger.GetPath() {
		if err := s.config.Logger.SetPath(s.config.LogPath); err != nil {
			return err
		}
	}
	return nil
}

// SetLogger 用于设置负责日志记录的logger。
// 注意，由于可能存在并发安全问题，因此无法在运行时进行设置。
// md5:560d266f79fb0915
// ff:设置日志记录器
// s:
// logger:日志记录器
func (s *Server) SetLogger(logger *glog.Logger) {
	s.config.Logger = logger
}

// Logger 是 GetLogger 的别名。 md5:56065d1ce5cca4c9
// ff:Logger别名
// s:
func (s *Server) Logger() *glog.Logger {
	return s.config.Logger
}

// SetLogLevel 通过level字符串设置日志级别。 md5:c479bcdf03ab0fa2
// ff:设置日志开启记录等级
// s:
// level:等级
func (s *Server) SetLogLevel(level string) {
	s.config.LogLevel = level
}

// SetLogStdout 设置是否将日志内容输出到stdout。 md5:c93557220e40f70b
// ff:设置日志开启输出到CMD
// s:
// enabled:开启
func (s *Server) SetLogStdout(enabled bool) {
	s.config.LogStdout = enabled
}

// SetAccessLogEnabled 用于启用或禁用访问日志。 md5:a353da90e0c3de0d
// ff:设置日志开启访客记录
// s:
// enabled:开启
func (s *Server) SetAccessLogEnabled(enabled bool) {
	s.config.AccessLogEnabled = enabled
}

// SetErrorLogEnabled 开启或关闭错误日志。 md5:38a0655c0083b6d0
// ff:设置日志开启错误记录
// s:
// enabled:开启
func (s *Server) SetErrorLogEnabled(enabled bool) {
	s.config.ErrorLogEnabled = enabled
}

// SetErrorStack 启用/禁用错误堆栈功能。 md5:6c411a957b96e186
// ff:设置日志开启错误堆栈记录
// s:
// enabled:开启
func (s *Server) SetErrorStack(enabled bool) {
	s.config.ErrorStack = enabled
}

// GetLogPath 返回日志路径。 md5:0fe087f2f9b0a123
// ff:取日志存储目录
// s:
func (s *Server) GetLogPath() string {
	return s.config.LogPath
}

// IsAccessLogEnabled 检查访问日志是否已启用。 md5:b076cc230602118d
// ff:日志访客记录是否已开启
// s:
func (s *Server) IsAccessLogEnabled() bool {
	return s.config.AccessLogEnabled
}

// IsErrorLogEnabled 检查错误日志是否启用。 md5:2231c72f34764e99
// ff:日志错误记录是否已开启
// s:
func (s *Server) IsErrorLogEnabled() bool {
	return s.config.ErrorLogEnabled
}
