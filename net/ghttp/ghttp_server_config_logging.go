// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

import (
	"github.com/888go/goframe/os/glog"
)

// SetLogPath 设置服务器的日志路径。
// 只有当设置了日志路径时，才会将日志内容记录到文件中。
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

// SetLogger 设置日志记录器以承担日志记录职责。
// 注意，由于可能存在并发安全问题，因此不能在运行时设置。
func (s *Server) SetLogger(logger *glog.Logger) {
	s.config.Logger = logger
}

// Logger 是 GetLogger 的别名。
func (s *Server) Logger() *glog.Logger {
	return s.config.Logger
}

// SetLogLevel 通过级别字符串设置日志等级。
func (s *Server) SetLogLevel(level string) {
	s.config.LogLevel = level
}

// SetLogStdout 设置是否将日志内容输出到 stdout。
func (s *Server) SetLogStdout(enabled bool) {
	s.config.LogStdout = enabled
}

// SetAccessLogEnabled 开启/关闭访问日志功能。
func (s *Server) SetAccessLogEnabled(enabled bool) {
	s.config.AccessLogEnabled = enabled
}

// SetErrorLogEnabled 用于启用/禁用错误日志。
func (s *Server) SetErrorLogEnabled(enabled bool) {
	s.config.ErrorLogEnabled = enabled
}

// SetErrorStack 开启或关闭错误堆栈功能。
func (s *Server) SetErrorStack(enabled bool) {
	s.config.ErrorStack = enabled
}

// GetLogPath 返回日志路径。
func (s *Server) GetLogPath() string {
	return s.config.LogPath
}

// IsAccessLogEnabled 检查访问日志是否已启用。
func (s *Server) IsAccessLogEnabled() bool {
	return s.config.AccessLogEnabled
}

// IsErrorLogEnabled 检查错误日志是否已启用。
func (s *Server) IsErrorLogEnabled() bool {
	return s.config.ErrorLogEnabled
}
