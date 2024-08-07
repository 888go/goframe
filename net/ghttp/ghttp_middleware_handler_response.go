// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"net/http"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

// DefaultHandlerResponse是HandlerResponse的默认实现。 md5:9340fa71a0d9e8f7
type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

// MiddlewareHandlerResponse 是默认的处理handler响应对象及其错误的中间件。 md5:d59676d7f703b4b1
func MiddlewareHandlerResponse(r *Request) {
	r.X中间件管理器.Next()

		// 存在自定义缓冲区内容，然后退出当前处理器。 md5:fd21f1b41f115a81
	if r.X响应.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.X取错误信息()
		res  = r.X取响应对象及错误信息()
		code = gerror.X取错误码(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.X响应.Status > 0 && r.X响应.Status != http.StatusOK {
			msg = http.StatusText(r.X响应.Status)
			switch r.X响应.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
						// 由于其他中间件可以获取到这个错误，所以它会产生错误。 md5:36a5d15e82de8d66
			err = gerror.X创建错误码(code, msg)
			r.X设置错误信息(err)
		} else {
			code = gcode.CodeOK
		}
	}

	r.X响应.X写响应缓冲区JSON(DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}
