// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gerror
import (
	"fmt"
	
	"github.com/888go/goframe/errors/gcode"
	)
// New根据给定的文本创建并返回一个格式化后的错误。
func New(text string) error {
	return &Error{
		stack: callers(),
		text:  text,
		code:  gcode.CodeNil,
	}
}

// Newf 返回一个错误，其格式化输出为给定的 format 和 args。
func Newf(format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  gcode.CodeNil,
	}
}

// NewSkip 创建并返回一个根据给定文本格式化的错误。
// 参数 `skip` 指定了跳过堆栈调用者的数量。
func NewSkip(skip int, text string) error {
	return &Error{
		stack: callers(skip),
		text:  text,
		code:  gcode.CodeNil,
	}
}

// NewSkipf 返回一个格式化为给定格式和参数的错误。
// 参数 `skip` 指定了跳过的调用栈层数量。
func NewSkipf(skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  gcode.CodeNil,
	}
}

// Wrap 使用文本包装错误。如果给出的 err 为 nil，则返回 nil。
// 注意，它不会丢失被包装错误的错误码，因为它会继承该错误的错误码。
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

// Wrapf 返回一个错误，该错误在Wrapf调用的位置添加了堆栈跟踪信息以及格式化说明符。
// 如果给出的`err`为nil，则返回nil。
// 注意，它不会丢失被包装错误的错误码，因为它会继承该错误的错误码。
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

// WrapSkip 用于用文本包装错误。如果给定的 err 为 nil，则返回 nil。
// 参数 `skip` 指定了跳过的堆栈调用者数量。
// 注意，它不会丢失包装错误的错误码，因为它会继承该错误的错误码。
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

// WrapSkipf 通过给定的格式和参数对错误进行包装并添加文本描述。如果给定的err为nil，此函数将返回nil。
// 参数`skip`用于指定跳过堆栈调用者的数量。
// 注意，它不会丢失包装错误的错误码，因为它会继承该错误的错误码。
// 以下是翻译后更详尽的中文注释：
// ```go
// WrapSkipf 函数用于对传入的错误 err 进行包装，同时使用给定的 format 和 args 格式化输出附加的错误信息。
// 若传入的原始错误 err 为 nil，则 WrapSkipf 函数直接返回 nil。
// 参数 `skip` 指定了在获取堆栈信息时需要跳过的调用层级数量。
// 需要注意的是，WrapSkipf 在对错误进行包装的过程中，会保留原始错误的错误代码，因为它从原始错误中继承了该错误代码属性。
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
