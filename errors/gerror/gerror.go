// 版权所有 GoFrame gf 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一个。

// Package gerror 提供了丰富的错误处理功能。
//
// 维护者请注意，
// 该包是一个基础包，除标准库和内部包外，**不应该**引入额外的包，以避免循环导入问题。
package 错误类

import (
	"github.com/888go/goframe/errors/gcode"
)

// IIs 是 Is 特性的接口。
type IIs interface {
	Error() string
	Is(target error) bool
}

// IEqual 是用于 Equal 功能的接口。
type IEqual interface {
	Error() string
	Equal(target error) bool
}

// ICode是Code功能的接口。
type ICode interface {
	Error() string
	Code() 错误码类.Code
}

// IStack 是 Stack 功能的接口。
type IStack interface {
	Error() string
	Stack() string
}

// ICause 是 Cause 功能的接口。
type ICause interface {
	Error() string
	Cause() error
}

// ICurrent 是用于“当前特性”的接口。
type ICurrent interface {
	Error() string
	Current() error
}

// IUnwrap 是用于 Unwrap 功能的接口。
type IUnwrap interface {
	Error() string
	Unwrap() error
}

const (
	// commaSeparatorSpace 是带空格的逗号分隔符。
	commaSeparatorSpace = ", "
)
