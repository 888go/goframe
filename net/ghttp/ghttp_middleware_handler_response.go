// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"net/http"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
)

// DefaultHandlerResponse 是 HandlerResponse 接口的默认实现。
type X默认处理器响应 struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

// MiddlewareHandlerResponse 是默认中间件处理处理器响应对象及其错误的接口。
func MiddlewareHandlerResponse(r *X请求) {
	r.X中间件管理器.Next()

	// 如果存在自定义缓冲区内容，则退出当前处理器。
	if r.X响应.X取缓冲区长度() > 0 {
		return
	}

	var (
		msg  string
		err  = r.X取错误信息()
		res  = r.X取响应对象及错误信息()
		code = 错误类.X取错误码(err)
	)
	if err != nil {
		if code == 错误码类.CodeNil {
			code = 错误码类.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.X响应.Status > 0 && r.X响应.Status != http.StatusOK {
			msg = http.StatusText(r.X响应.Status)
			switch r.X响应.Status {
			case http.StatusNotFound:
				code = 错误码类.CodeNotFound
			case http.StatusForbidden:
				code = 错误码类.CodeNotAuthorized
			default:
				code = 错误码类.CodeUnknown
			}
			// 它创建错误，以便其他中间件可以获取该错误。
			err = 错误类.X创建错误码(code, msg)
			r.X设置错误信息(err)
		} else {
			code = 错误码类.CodeOK
		}
	}

	r.X响应.X写响应缓冲区JSON(X默认处理器响应{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}
