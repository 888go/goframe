// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

// Package gerror 提供了丰富的错误处理功能。
// 
// 对于维护者，请注意，这个包是一个基础包，不应该导入除标准包和内部包以外的额外包，以避免循环导入。
// md5:ee2cf0d4a8603263
package gerror

import (
	"github.com/gogf/gf/v2/errors/gcode"
)

// IIs 是 Is 特性接口。. md5:c4d4591e57d60306
type IIs interface {
	Error() string
	Is(target error) bool
}

// IEqual 是实现相等特性（Equal feature）的接口。. md5:734ec9a0f4399a8f
type IEqual interface {
	Error() string
	Equal(target error) bool
}

// ICode是代码功能的接口。. md5:97b69f390ed9b25c
type ICode interface {
	Error() string
	Code() gcode.Code
}

// IStack是栈功能的接口。. md5:bdce0c99c1f4dfb1
type IStack interface {
	Error() string
	Stack() string
}

// ICause 是Cause特性的接口。. md5:f98814a45dc166a2
type ICause interface {
	Error() string
	Cause() error
}

// ICurrent 是表示 Current 功能的接口。. md5:df6566b1d7d711b5
type ICurrent interface {
	Error() string
	Current() error
}

// IUnwrap是Unwrap功能的接口。. md5:71a8ef0c90103c1e
type IUnwrap interface {
	Error() string
	Unwrap() error
}

const (
	// commaSeparatorSpace 是带有空格的逗号分隔符。. md5:ef2597553e27a102
	commaSeparatorSpace = ", "
)
