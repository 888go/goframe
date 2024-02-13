// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 错误类

import (
	"github.com/888go/goframe/errors/gcode"
)

// Code 返回错误代码。
// 如果没有错误代码，它将返回 CodeNil。
func (err *Error) Code() 错误码类.Code {
	if err == nil {
		return 错误码类.CodeNil
	}
	if err.code == 错误码类.CodeNil {
		return X取错误码(err.Unwrap())
	}
	return err.code
}

// SetCode 将给定的代码用于更新内部代码。
func (err *Error) SetCode(code 错误码类.Code) {
	if err == nil {
		return
	}
	err.code = code
}
