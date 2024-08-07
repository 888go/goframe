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

// Option 是创建错误时的选项。 md5:e4c01dc11cd0fde6
type Option struct {
	Error error      // Wrapped error if any.
	Stack bool       // 是否将堆栈信息记录到错误中。 md5:d22b9750d48ecbd8
	Text  string     // 由New*函数创建的错误文本。 md5:f5b9bbbd3516fe7c
	Code  gcode.Code // 如果需要，错误代码。 md5:92ec689bb1dbbb33
}

// NewWithOption 创建并返回一个带有选项的自定义错误。
// 这是创建错误的高级用法，通常在框架内部使用。
// md5:0727fc0de90f397f
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

// X弃用NewOption 创建并返回一个带有 Option 的自定义错误。
// 已弃用：请使用 NewWithOption 替代。
// md5:dba08610e527e0f9
func X弃用NewOption(选项 Option) error {
	return NewWithOption(选项)
}
