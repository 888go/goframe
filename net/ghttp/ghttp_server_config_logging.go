// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"github.com/888go/goframe/os/glog"
)

// SetLogPath 设置服务器的日志路径。
// 只有当设置了日志路径时，才会将日志内容记录到文件中。
func (s *X服务) X设置日志存储目录(目录 string) error {
	if len(目录) == 0 {
		return nil
	}
	s.config.X日志存储目录 = 目录
	s.config.X日志开启错误记录 = true
	s.config.X日志开启访客记录 = true
	if s.config.X日志存储目录 != "" && s.config.X日志存储目录 != s.config.X日志记录器.X取文件路径() {
		if err := s.config.X日志记录器.X设置文件路径(s.config.X日志存储目录); err != nil {
			return err
		}
	}
	return nil
}

// SetLogger 设置日志记录器以承担日志记录职责。
// 注意，由于可能存在并发安全问题，因此不能在运行时设置。
func (s *X服务) X设置日志记录器(日志记录器 *日志类.Logger) {
	s.config.X日志记录器 = 日志记录器
}

// Logger 是 GetLogger 的别名。
func (s *X服务) Logger别名() *日志类.Logger {
	return s.config.X日志记录器
}

// SetLogLevel 通过级别字符串设置日志等级。
func (s *X服务) X设置日志开启记录等级(等级 string) {
	s.config.X日志记录等级 = 等级
}

// SetLogStdout 设置是否将日志内容输出到 stdout。
func (s *X服务) X设置日志开启输出到CMD(开启 bool) {
	s.config.X日志开启输出到CMD = 开启
}

// SetAccessLogEnabled 开启/关闭访问日志功能。
func (s *X服务) X设置日志开启访客记录(开启 bool) {
	s.config.X日志开启访客记录 = 开启
}

// SetErrorLogEnabled 用于启用/禁用错误日志。
func (s *X服务) X设置日志开启错误记录(开启 bool) {
	s.config.X日志开启错误记录 = 开启
}

// SetErrorStack 开启或关闭错误堆栈功能。
func (s *X服务) X设置日志开启错误堆栈记录(开启 bool) {
	s.config.X日志开启错误堆栈记录 = 开启
}

// GetLogPath 返回日志路径。
func (s *X服务) X取日志存储目录() string {
	return s.config.X日志存储目录
}

// IsAccessLogEnabled 检查访问日志是否已启用。
func (s *X服务) X日志访客记录是否已开启() bool {
	return s.config.X日志开启访客记录
}

// IsErrorLogEnabled 检查错误日志是否已启用。
func (s *X服务) X日志错误记录是否已开启() bool {
	return s.config.X日志开启错误记录
}
