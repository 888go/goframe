// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 错误类

import (
	"fmt"
	"strings"
	
	"github.com/888go/goframe/errors/gcode"
)

// NewCode创建并返回一个错误，该错误包含错误代码和给定的文本。
func X创建错误码(错误码 错误码类.Code, 错误文本 ...string) error {
	return &Error{
		stack: callers(),
		text:  strings.Join(错误文本, commaSeparatorSpace),
		code:  错误码,
	}
}

// NewCodef返回一个错误，该错误包含错误代码，并使用给定的格式和参数进行格式化。
func X创建错误码并格式化(错误码 错误码类.Code, 格式 string, 参数 ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(格式, 参数...),
		code:  错误码,
	}
}

// NewCodeSkip 创建并返回一个错误，该错误包含错误代码，并根据给定文本格式化而成。
// 参数`skip`用于指定跳过堆栈调用者的数量。
func X创建错误码并跳过堆栈(错误码 错误码类.Code, 跳过堆栈 int, 错误文本 ...string) error {
	return &Error{
		stack: callers(跳过堆栈),
		text:  strings.Join(错误文本, commaSeparatorSpace),
		code:  错误码,
	}
}

// NewCodeSkipf 返回一个具有错误代码且格式化为给定格式和参数的错误。
// 参数 `skip` 指定了要跳过的堆栈调用者数量。
func X创建错误码并跳过堆栈与格式化(错误码 错误码类.Code, 跳过堆栈 int, 格式 string, 参数 ...interface{}) error {
	return &Error{
		stack: callers(跳过堆栈),
		text:  fmt.Sprintf(格式, 参数...),
		code:  错误码,
	}
}

// WrapCode 将错误用代码和文本包装起来。
// 如果给出的 err 为 nil，则返回 nil。
func X多层错误码(错误码 错误码类.Code, 上一层错误 error, 错误文本 ...string) error {
	if 上一层错误 == nil {
		return nil
	}
	return &Error{
		error: 上一层错误,
		stack: callers(),
		text:  strings.Join(错误文本, commaSeparatorSpace),
		code:  错误码,
	}
}

// WrapCodef 将错误用代码和格式符包装起来。
// 如果给出的 `err` 为 nil，则返回 nil。
func X多层错误码并格式化(错误码 错误码类.Code, 上一层错误 error, 格式 string, 参数 ...interface{}) error {
	if 上一层错误 == nil {
		return nil
	}
	return &Error{
		error: 上一层错误,
		stack: callers(),
		text:  fmt.Sprintf(格式, 参数...),
		code:  错误码,
	}
}

// WrapCodeSkip 通过代码和文本包装错误。
// 如果给出的 err 为 nil，则返回 nil。
// 参数 `skip` 指定了需要跳过的堆栈调用者数量。
func X多层错误码并跳过堆栈(错误码 错误码类.Code, 跳过堆栈 int, 上一层错误 error, 错误文本 ...string) error {
	if 上一层错误 == nil {
		return nil
	}
	return &Error{
		error: 上一层错误,
		stack: callers(跳过堆栈),
		text:  strings.Join(错误文本, commaSeparatorSpace),
		code:  错误码,
	}
}

// WrapCodeSkipf 对给定的错误 err 进行包装，添加代码和格式化后的文本信息。如果给定的 err 为 nil，则返回 nil。
// 参数 `skip` 指定了需要跳过的堆栈调用者数量。
func X多层错误码并跳过堆栈与格式化(错误码 错误码类.Code, 跳过堆栈 int, 上一层错误 error, 格式 string, 参数 ...interface{}) error {
	if 上一层错误 == nil {
		return nil
	}
	return &Error{
		error: 上一层错误,
		stack: callers(跳过堆栈),
		text:  fmt.Sprintf(格式, 参数...),
		code:  错误码,
	}
}

// Code 返回当前错误的错误代码。
// 如果该错误没有错误代码，或者未实现 Code 接口，它将返回 `CodeNil`。
func X取错误码(错误 error) 错误码类.Code {
	if 错误 == nil {
		return 错误码类.CodeNil
	}
	if e, ok := 错误.(ICode); ok {
		return e.Code()
	}
	if e, ok := 错误.(IUnwrap); ok {
		return X取错误码(e.Unwrap())
	}
	return 错误码类.CodeNil
}

// HasCode 检查并报告 `err` 在其链式错误中是否包含 `code`。
func X是否包含错误码(错误 error, 错误码 错误码类.Code) bool {
	if 错误 == nil {
		return false
	}
	if e, ok := 错误.(ICode); ok {
		return 错误码 == e.Code()
	}
	if e, ok := 错误.(IUnwrap); ok {
		return X是否包含错误码(e.Unwrap(), 错误码)
	}
	return false
}
