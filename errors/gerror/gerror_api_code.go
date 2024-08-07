// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 错误类

import (
	"fmt"
	"strings"

	gcode "github.com/888go/goframe/errors/gcode"
)

// X创建错误码 创建并返回一个具有错误代码和给定文本的错误。 md5:5f88f8ae1151acac
func X创建错误码(错误码 gcode.Code, text ...string) error {
	return &Error{
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  错误码,
	}
}

// X创建错误码并格式化 返回一个具有错误代码，并按照给定格式和参数格式化的错误。 md5:bb6b90ee5a4ce175
func X创建错误码并格式化(错误码 gcode.Code, format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  错误码,
	}
}

// X创建错误码并跳过堆栈 创建并返回一个带有错误码的错误，该错误根据给定的文本格式化。
// 参数 `skip` 指定了跳过的堆栈调用者数量。
// md5:5c3aabed2ce89e0c
func X创建错误码并跳过堆栈(错误码 gcode.Code, skip int, text ...string) error {
	return &Error{
		stack: callers(skip),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  错误码,
	}
}

// X创建错误码并跳过堆栈与格式化 返回一个具有指定错误代码和格式化参数的错误。
// 参数 `skip` 指定了要跳过的调用堆栈数量。
// md5:ccd3b74e8b4f8acc
func X创建错误码并跳过堆栈与格式化(错误码 gcode.Code, skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  错误码,
	}
}

// X多层错误码 将错误与代码和文本一起包装。
// 如果给定的 err 为 nil，它将返回 nil。
// md5:5e09a5ffb6fa4e21
func X多层错误码(错误码 gcode.Code, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  错误码,
	}
}

// X多层错误码并格式化 将错误与代码和格式化占位符一起包装。
// 如果给定的 `err` 为 nil，它将返回 nil。
// md5:ef3a7436eb342ff6
func X多层错误码并格式化(错误码 gcode.Code, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  错误码,
	}
}

// X多层错误码并跳过堆栈 用于给错误附加代码和文本信息。
// 如果给定的err为nil，该函数将返回nil。
// 参数 `skip` 指定了要跳过的堆栈调用者数量。
// md5:5ee348edd866b587
func X多层错误码并跳过堆栈(错误码 gcode.Code, skip int, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  错误码,
	}
}

// X多层错误码并跳过堆栈与格式化 使用给定的格式和参数将错误包装成带有代码和文本的错误。
// 如果给定的err为nil，它将返回nil。
// 参数`skip`指定了要跳过的调用者堆栈的数量。
// md5:00fbaefc556da645
func X多层错误码并跳过堆栈与格式化(错误码 gcode.Code, skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  错误码,
	}
}

// X取错误码 函数返回当前错误的错误代码。
// 如果它没有错误代码并且也没有实现 X取错误码 接口，它将返回 CodeNil。
// md5:33b7429f6f7b3dbe
func X取错误码(错误 error) gcode.Code {
	if 错误 == nil {
		return gcode.CodeNil
	}
	if e, ok := 错误.(ICode); ok {
		return e.Code()
	}
	if e, ok := 错误.(IUnwrap); ok {
		return X取错误码(e.Unwrap())
	}
	return gcode.CodeNil
}

// X是否包含错误码 检查并报告 `err` 的链式错误中是否包含 `code`。 md5:5d1b8286d1872717
func X是否包含错误码(错误 error, 错误码 gcode.Code) bool {
	if 错误 == nil {
		return false
	}
	if e, ok := 错误.(ICode); ok && 错误码 == e.Code() {
		return true
	}
	if e, ok := 错误.(IUnwrap); ok {
		return X是否包含错误码(e.Unwrap(), 错误码)
	}
	return false
}
