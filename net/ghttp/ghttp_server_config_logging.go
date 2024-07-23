// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import "github.com/gogf/gf/v2/os/glog"

// SetLogPath sets the log path for server.
// It logs content to file only if the log path is set.
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

// SetLogger sets the logger for logging responsibility.
// Note that it cannot be set in runtime as there may be concurrent safety issue.
// ff:设置日志记录器
// s:
// logger:日志记录器
func (s *Server) SetLogger(logger *glog.Logger) {
	s.config.Logger = logger
}

// Logger is alias of GetLogger.
// ff:Logger别名
// s:
func (s *Server) Logger() *glog.Logger {
	return s.config.Logger
}

// SetLogLevel sets logging level by level string.
// ff:设置日志开启记录等级
// s:
// level:等级
func (s *Server) SetLogLevel(level string) {
	s.config.LogLevel = level
}

// SetLogStdout sets whether output the logging content to stdout.
// ff:设置日志开启输出到CMD
// s:
// enabled:开启
func (s *Server) SetLogStdout(enabled bool) {
	s.config.LogStdout = enabled
}

// SetAccessLogEnabled enables/disables the access log.
// ff:设置日志开启访客记录
// s:
// enabled:开启
func (s *Server) SetAccessLogEnabled(enabled bool) {
	s.config.AccessLogEnabled = enabled
}

// SetErrorLogEnabled enables/disables the error log.
// ff:设置日志开启错误记录
// s:
// enabled:开启
func (s *Server) SetErrorLogEnabled(enabled bool) {
	s.config.ErrorLogEnabled = enabled
}

// SetErrorStack enables/disables the error stack feature.
// ff:设置日志开启错误堆栈记录
// s:
// enabled:开启
func (s *Server) SetErrorStack(enabled bool) {
	s.config.ErrorStack = enabled
}

// GetLogPath returns the log path.
// ff:取日志存储目录
// s:
func (s *Server) GetLogPath() string {
	return s.config.LogPath
}

// IsAccessLogEnabled checks whether the access log enabled.
// ff:日志访客记录是否已开启
// s:
func (s *Server) IsAccessLogEnabled() bool {
	return s.config.AccessLogEnabled
}

// IsErrorLogEnabled checks whether the error log enabled.
// ff:日志错误记录是否已开启
// s:
func (s *Server) IsErrorLogEnabled() bool {
	return s.config.ErrorLogEnabled
}
