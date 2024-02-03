// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gerror

import (
	"github.com/888go/goframe/errors/gcode"
)

// Option 是用于创建错误的选项。
type Option struct {
	Error error      // 如果存在，则为包装后的错误。
	Stack bool       // 是否将堆栈信息记录到错误中。
	Text  string     // 错误文本，由 New* 函数创建。
	Code  gcode.Code // 如果必要，此处为错误代码。
}

// NewWithOption 根据选项创建并返回一个自定义错误。
// 这是用于创建错误的高级用法，常在框架内部使用。
func NewWithOption(option Option) error {
	err := &Error{
		error: option.Error,
		text:  option.Text,
		code:  option.Code,
	}
	if option.Stack {
		err.stack = callers()
	}
	return err
}

// NewOption 创建并返回一个带有 Option 的自定义错误。
// 已弃用：请改用 NewWithOption。
func NewOption(option Option) error {
	return NewWithOption(option)
}
