// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

import (
	"fmt"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/instance"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/text/gstr"
)

// handleAccessLog 处理服务器的访问日志。
func (s *Server) handleAccessLog(r *Request) {
	if !s.IsAccessLogEnabled() {
		return
	}
	var (
		scheme            = "http"
		proto             = r.Header.Get("X-Forwarded-Proto")
		loggerInstanceKey = fmt.Sprintf(`Acccess Logger Of Server:%s`, s.instance)
	)

	if r.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}
	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s"`,
		r.Response.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.LeaveTime-r.EnterTime)/1000,
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

// handleErrorLog 处理服务器的错误日志。
func (s *Server) handleErrorLog(err error, r *Request) {
	// 如果错误日志自定义禁用，则此操作无任何效果。
	if !s.IsErrorLogEnabled() {
		return
	}
	var (
		code              = gerror.Code(err)
		scheme            = "http"
		codeDetail        = code.Detail()
		proto             = r.Header.Get("X-Forwarded-Proto")
		loggerInstanceKey = fmt.Sprintf(`Error Logger Of Server:%s`, s.instance)
		codeDetailStr     string
	)
	if r.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}
	if codeDetail != nil {
		codeDetailStr = gstr.Replace(fmt.Sprintf(`%+v`, codeDetail), "\n", " ")
	}
	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s", %d, "%s", "%+v"`,
		r.Response.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.LeaveTime-r.EnterTime)/1000,
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
