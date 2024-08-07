// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 错误类

import (
	"fmt"

	gcode "github.com/888go/goframe/errors/gcode"
)

// X创建 创建并返回一个根据给定文本格式化的错误。 md5:de9ec7c958a945bb
func X创建(错误文本 string) error {
	return &Error{
		stack: callers(),
		text:  错误文本,
		code:  gcode.CodeNil,
	}
}

// X创建并格式化 返回一个根据给定格式和参数格式化的错误。 md5:bd62f35687f8bc83
func X创建并格式化(格式 string, 参数 ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(格式, 参数...),
		code:  gcode.CodeNil,
	}
}

// X创建并跳过堆栈 创建并返回一个根据给定文本格式化的错误。参数 `skip` 指定了要跳过的调用者堆栈数量。
// md5:22bec296ea4c17b0
func X创建并跳过堆栈(跳过堆栈 int, 错误文本 string) error {
	return &Error{
		stack: callers(跳过堆栈),
		text:  错误文本,
		code:  gcode.CodeNil,
	}
}

// X创建并跳过堆栈与格式化 返回一个按照给定格式和参数格式化的错误。
// 参数 `skip` 指定了跳过调用栈的层数。
// md5:82d8fef84b9d2ba0
func X创建并跳过堆栈与格式化(跳过堆栈 int, 格式 string, 参数 ...interface{}) error {
	return &Error{
		stack: callers(跳过堆栈),
		text:  fmt.Sprintf(格式, 参数...),
		code:  gcode.CodeNil,
	}
}

// X多层错误 使用文本包装错误。如果给定的 err 为 nil，则返回 nil。
// 注意，它不会丢失被包裹错误的错误码，因为它从被包裹的错误中继承了错误码。
// md5:e04f9222b50c8938
func X多层错误(上一层错误 error, 错误文本 string) error {
	if 上一层错误 == nil {
		return nil
	}
	return &Error{
		error: 上一层错误,
		stack: callers(),
		text:  错误文本,
		code:  X取错误码(上一层错误),
	}
}

// X多层错误并格式化 会在调用 X多层错误并格式化 的位置为错误 err 添加一个堆栈跟踪信息，并使用格式化指定器。
// 如果给定的 `err` 为 nil，它将返回 nil。
// 请注意，它不会丢失被包装错误的错误代码，因为它从错误中继承了错误代码。
// md5:cbfccfaa6fa0bee1
func X多层错误并格式化(上一层错误 error, 格式 string, 参数 ...interface{}) error {
	if 上一层错误 == nil {
		return nil
	}
	return &Error{
		error: 上一层错误,
		stack: callers(),
		text:  fmt.Sprintf(格式, 参数...),
		code:  X取错误码(上一层错误),
	}
}

// X多层错误并跳过堆栈 使用文本包装错误。如果给定的 err 为 nil，它将返回 nil。
// 参数 `skip` 指定了跳过调用堆栈的层数。
// 注意，它不会丢失被包装错误的错误代码，因为它是从其继承错误代码的。
// md5:5f87402ce06c586b
func X多层错误并跳过堆栈(跳过堆栈 int, 上一层错误 error, 错误文本 string) error {
	if 上一层错误 == nil {
		return nil
	}
	return &Error{
		error: 上一层错误,
		stack: callers(跳过堆栈),
		text:  错误文本,
		code:  X取错误码(上一层错误),
	}
}

// X多层错误并跳过堆栈与格式化 将错误用给定的格式和参数进行文本包装。如果给定的 err 为 nil，它将返回 nil。
// 参数 `skip` 指定了要跳过的调用栈层数。
// 注意，它不会丢失被包装错误的错误代码，因为它是从原始错误中继承错误代码的。
// md5:82d4f5ae39c67b27
func X多层错误并跳过堆栈与格式化(跳过堆栈 int, 上一层错误 error, 格式 string, 参数 ...interface{}) error {
	if 上一层错误 == nil {
		return nil
	}
	return &Error{
		error: 上一层错误,
		stack: callers(跳过堆栈),
		text:  fmt.Sprintf(格式, 参数...),
		code:  X取错误码(上一层错误),
	}
}
