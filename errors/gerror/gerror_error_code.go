// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 错误类

import (
	gcode "github.com/888go/goframe/errors/gcode"
)

// Code 返回错误代码。
// 如果没有错误代码，它将返回 CodeNil。
// md5:ff28e1e7c152b8de
func (err *Error) Code() gcode.Code {
	if err == nil {
		return gcode.CodeNil
	}
	if err.code == gcode.CodeNil {
		return X取错误码(err.Unwrap())
	}
	return err.code
}

// SetCode 使用给定的代码更新内部代码。 md5:e53ddd795359c2bb
func (err *Error) SetCode(code gcode.Code) {
	if err == nil {
		return
	}
	err.code = code
}
