// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/instance"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

// handleAccessLog 处理服务器的访问日志。 md5:37f533b1207682a1
func (s *Server) handleAccessLog(r *Request) {
	if !s.IsAccessLogEnabled() {
		return
	}
	var (
		scheme            = r.GetSchema()
		loggerInstanceKey = fmt.Sprintf(`Acccess Logger Of Server:%s`, s.instance)
	)
	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s"`,
		r.Response.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.LeaveTime.Sub(r.EnterTime).Milliseconds())/1000,
		r.GetClientIp(), r.Referer(), r.UserAgent(),
	)
	logger := instance.GetOrSetFuncLock(loggerInstanceKey, func() interface{} {
		l := s.Logger().Clone()
		l.SetFile(s.config.AccessLogPattern)
		l.SetStdoutPrint(s.config.LogStdout)
		l.SetLevelPrint(false)
		return l
	}).(*glog.Logger)
	logger.Print(r.Context(), content)
}

// handleErrorLog 处理服务器的错误日志。 md5:f4ba08c94e4f5e6f
func (s *Server) handleErrorLog(err error, r *Request) {
	// 如果错误日志记录被特别禁用，则什么也不做。 md5:f40ab65302593bd7
	if !s.IsErrorLogEnabled() {
		return
	}
	var (
		code              = gerror.Code(err)
		scheme            = r.GetSchema()
		codeDetail        = code.Detail()
		loggerInstanceKey = fmt.Sprintf(`Error Logger Of Server:%s`, s.instance)
		codeDetailStr     string
	)
	if codeDetail != nil {
		codeDetailStr = gstr.Replace(fmt.Sprintf(`%+v`, codeDetail), "\n", " ")
	}
	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s", %d, "%s", "%+v"`,
		r.Response.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.LeaveTime.Sub(r.EnterTime))/1000,
		r.GetClientIp(), r.Referer(), r.UserAgent(),
		code.Code(), code.Message(), codeDetailStr,
	)
	if s.config.ErrorStack {
		if stack := gerror.Stack(err); stack != "" {
			content += "\nStack:\n" + stack
		} else {
			content += ", " + err.Error()
		}
	} else {
		content += ", " + err.Error()
	}
	logger := instance.GetOrSetFuncLock(loggerInstanceKey, func() interface{} {
		l := s.Logger().Clone()
		l.SetStack(false)
		l.SetFile(s.config.ErrorLogPattern)
		l.SetStdoutPrint(s.config.LogStdout)
		l.SetLevelPrint(false)
		return l
	}).(*glog.Logger)
	logger.Error(r.Context(), content)
}
