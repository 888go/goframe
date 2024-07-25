// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gerror

import (
	"fmt"

	"github.com/gogf/gf/v2/errors/gcode"
)

// New 创建并返回一个根据给定文本格式化的错误。 md5:de9ec7c958a945bb
func New(text string) error {
	return &Error{
		stack: callers(),
		text:  text,
		code:  gcode.CodeNil,
	}
}

// Newf 返回一个根据给定格式和参数格式化的错误。 md5:bd62f35687f8bc83
func Newf(format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  gcode.CodeNil,
	}
}

// NewSkip 创建并返回一个根据给定文本格式化的错误。参数 `skip` 指定了要跳过的调用者堆栈数量。 md5:22bec296ea4c17b0
func NewSkip(skip int, text string) error {
	return &Error{
		stack: callers(skip),
		text:  text,
		code:  gcode.CodeNil,
	}
}

// NewSkipf 返回一个按照给定格式和参数格式化的错误。
// 参数 `skip` 指定了跳过调用栈的层数。 md5:82d8fef84b9d2ba0
func NewSkipf(skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  gcode.CodeNil,
	}
}

// Wrap 使用文本包装错误。如果给定的 err 为 nil，则返回 nil。
// 注意，它不会丢失被包裹错误的错误码，因为它从被包裹的错误中继承了错误码。 md5:e04f9222b50c8938
func Wrap(err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  text,
		code:  Code(err),
	}
}

// Wrapf 会在调用 Wrapf 的位置为错误 err 添加一个堆栈跟踪信息，并使用格式化指定器。
// 如果给定的 `err` 为 nil，它将返回 nil。
// 请注意，它不会丢失被包装错误的错误代码，因为它从错误中继承了错误代码。 md5:cbfccfaa6fa0bee1
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  Code(err),
	}
}

// WrapSkip 使用文本包装错误。如果给定的 err 为 nil，它将返回 nil。
// 参数 `skip` 指定了跳过调用堆栈的层数。
// 注意，它不会丢失被包装错误的错误代码，因为它是从其继承错误代码的。 md5:5f87402ce06c586b
func WrapSkip(skip int, err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  text,
		code:  Code(err),
	}
}

// WrapSkipf 将错误用给定的格式和参数进行文本包装。如果给定的 err 为 nil，它将返回 nil。
// 参数 `skip` 指定了要跳过的调用栈层数。
// 注意，它不会丢失被包装错误的错误代码，因为它是从原始错误中继承错误代码的。 md5:82d4f5ae39c67b27
func WrapSkipf(skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  Code(err),
	}
}
