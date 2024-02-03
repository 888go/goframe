// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gerror

import (
	"fmt"
	"strings"
	
	"github.com/888go/goframe/errors/gcode"
)

// NewCode创建并返回一个错误，该错误包含错误代码和给定的文本。
func NewCode(code gcode.Code, text ...string) error {
	return &Error{
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// NewCodef返回一个错误，该错误包含错误代码，并使用给定的格式和参数进行格式化。
func NewCodef(code gcode.Code, format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// NewCodeSkip 创建并返回一个错误，该错误包含错误代码，并根据给定文本格式化而成。
// 参数`skip`用于指定跳过堆栈调用者的数量。
func NewCodeSkip(code gcode.Code, skip int, text ...string) error {
	return &Error{
		stack: callers(skip),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// NewCodeSkipf 返回一个具有错误代码且格式化为给定格式和参数的错误。
// 参数 `skip` 指定了要跳过的堆栈调用者数量。
func NewCodeSkipf(code gcode.Code, skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCode 将错误用代码和文本包装起来。
// 如果给出的 err 为 nil，则返回 nil。
func WrapCode(code gcode.Code, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// WrapCodef 将错误用代码和格式符包装起来。
// 如果给出的 `err` 为 nil，则返回 nil。
func WrapCodef(code gcode.Code, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCodeSkip 通过代码和文本包装错误。
// 如果给出的 err 为 nil，则返回 nil。
// 参数 `skip` 指定了需要跳过的堆栈调用者数量。
func WrapCodeSkip(code gcode.Code, skip int, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// WrapCodeSkipf 对给定的错误 err 进行包装，添加代码和格式化后的文本信息。如果给定的 err 为 nil，则返回 nil。
// 参数 `skip` 指定了需要跳过的堆栈调用者数量。
func WrapCodeSkipf(code gcode.Code, skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// Code 返回当前错误的错误代码。
// 如果该错误没有错误代码，或者未实现 Code 接口，它将返回 `CodeNil`。
func Code(err error) gcode.Code {
	if err == nil {
		return gcode.CodeNil
	}
	if e, ok := err.(ICode); ok {
		return e.Code()
	}
	if e, ok := err.(IUnwrap); ok {
		return Code(e.Unwrap())
	}
	return gcode.CodeNil
}

// HasCode 检查并报告 `err` 在其链式错误中是否包含 `code`。
func HasCode(err error, code gcode.Code) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(ICode); ok {
		return code == e.Code()
	}
	if e, ok := err.(IUnwrap); ok {
		return HasCode(e.Unwrap(), code)
	}
	return false
}
