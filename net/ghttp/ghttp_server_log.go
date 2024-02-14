// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"fmt"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/instance"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/text/gstr"
)

// handleAccessLog 处理服务器的访问日志。
func (s *X服务) handleAccessLog(r *X请求) {
	if !s.X日志访客记录是否已开启() {
		return
	}
	var (
		scheme            = "http"
		proto             = r.Header.Get("X-Forwarded-Proto")
		loggerInstanceKey = fmt.Sprintf(`Acccess Logger Of Server:%s`, s.instance)
	)

	if r.TLS != nil || 文本类.X相等比较并忽略大小写(proto, "https") {
		scheme = "https"
	}
	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s"`,
		r.X响应.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.X结束时间-r.X开始时间)/1000,
		r.X取客户端IP地址(), r.Referer(), r.UserAgent(),
	)
	logger := instance.GetOrSetFuncLock(loggerInstanceKey, func() interface{} {
		l := s.Logger别名().X取副本()
		l.X设置文件名格式(s.config.X日志访客文件命名模式)
		l.X设置是否同时输出到终端(s.config.X日志开启输出到CMD)
		l.X设置是否输出级别(false)
		return l
	}).(*日志类.Logger)
	logger.X输出(r.Context别名(), content)
}

// handleErrorLog 处理服务器的错误日志。
func (s *X服务) handleErrorLog(err error, r *X请求) {
	// 如果错误日志自定义禁用，则此操作无任何效果。
	if !s.X日志错误记录是否已开启() {
		return
	}
	var (
		code              = 错误类.X取错误码(err)
		scheme            = "http"
		codeDetail        = code.Detail()
		proto             = r.Header.Get("X-Forwarded-Proto")
		loggerInstanceKey = fmt.Sprintf(`Error Logger Of Server:%s`, s.instance)
		codeDetailStr     string
	)
	if r.TLS != nil || 文本类.X相等比较并忽略大小写(proto, "https") {
		scheme = "https"
	}
	if codeDetail != nil {
		codeDetailStr = 文本类.X替换(fmt.Sprintf(`%+v`, codeDetail), "\n", " ")
	}
	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s", %d, "%s", "%+v"`,
		r.X响应.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.X结束时间-r.X开始时间)/1000,
		r.X取客户端IP地址(), r.Referer(), r.UserAgent(),
		code.Code(), code.Message(), codeDetailStr,
	)
	if s.config.X日志开启错误堆栈记录 {
		if stack := 错误类.X取文本(err); stack != "" {
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
		l.X设置文件名格式(s.config.X日志错误文件命名模式)
		l.X设置是否同时输出到终端(s.config.X日志开启输出到CMD)
		l.X设置是否输出级别(false)
		return l
	}).(*日志类.Logger)
	logger.Error(r.Context别名(), content)
}
