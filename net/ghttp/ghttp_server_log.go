// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"fmt"

	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/instance"
	glog "github.com/888go/goframe/os/glog"
	gstr "github.com/888go/goframe/text/gstr"
)

// handleAccessLog 处理服务器的访问日志。 md5:37f533b1207682a1
func (s *X服务) handleAccessLog(r *Request) {
	if !s.X日志访客记录是否已开启() {
		return
	}
	var (
		scheme            = r.GetSchema()
		loggerInstanceKey = fmt.Sprintf(`Acccess Logger Of Server:%s`, s.instance)
	)
	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s"`,
		r.X响应.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.LeaveTime.X取纳秒时长(r.EnterTime).Milliseconds())/1000,
		r.X取客户端IP地址(), r.Referer(), r.UserAgent(),
	)
	logger := instance.GetOrSetFuncLock(loggerInstanceKey, func() interface{} {
		l := s.Logger别名().X取副本()
		l.X设置文件名格式(s.config.AccessLogPattern)
		l.X设置是否同时输出到终端(s.config.LogStdout)
		l.X设置是否输出级别(false)
		return l
	}).(*glog.Logger)
	logger.X输出(r.Context别名(), content)
}

// handleErrorLog 处理服务器的错误日志。 md5:f4ba08c94e4f5e6f
func (s *X服务) handleErrorLog(err error, r *Request) {
		// 如果错误日志记录被特别禁用，则什么也不做。 md5:f40ab65302593bd7
	if !s.X日志错误记录是否已开启() {
		return
	}
	var (
		code              = gerror.X取错误码(err)
		scheme            = r.GetSchema()
		codeDetail        = code.Detail()
		loggerInstanceKey = fmt.Sprintf(`Error Logger Of Server:%s`, s.instance)
		codeDetailStr     string
	)
	if codeDetail != nil {
		codeDetailStr = gstr.X替换(fmt.Sprintf(`%+v`, codeDetail), "\n", " ")
	}
	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s", %d, "%s", "%+v"`,
		r.X响应.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.LeaveTime.X取纳秒时长(r.EnterTime))/1000,
		r.X取客户端IP地址(), r.Referer(), r.UserAgent(),
		code.Code(), code.Message(), codeDetailStr,
	)
	if s.config.ErrorStack {
		if stack := gerror.X取文本(err); stack != "" {
			content += "\nStack:\n" + stack
		} else {
			content += ", " + err.Error()
		}
	} else {
		content += ", " + err.Error()
	}
	logger := instance.GetOrSetFuncLock(loggerInstanceKey, func() interface{} {
		l := s.Logger别名().X取副本()
		l.X设置堆栈跟踪(false)
		l.X设置文件名格式(s.config.ErrorLogPattern)
		l.X设置是否同时输出到终端(s.config.LogStdout)
		l.X设置是否输出级别(false)
		return l
	}).(*glog.Logger)
	logger.Error(r.Context别名(), content)
}
