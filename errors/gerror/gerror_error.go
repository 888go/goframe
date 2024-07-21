// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gerror

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
)

// Error 是自定义错误，用于附加功能。 md5:6111a19ebbc88a60
type Error struct {
	error error      // Wrapped error.
	stack stack      // Stack数组，用于记录在创建或包装此错误时的堆栈信息。 md5:a86aaae3cd5b8beb
	text  string     // 当创建错误时自定义错误文本，当错误代码不为nil时可能会为空。 md5:42ad3f2e459c445e
	code  gcode.Code // 如果需要，错误代码。 md5:92ec689bb1dbbb33
}

const (
	// 当前错误模块路径的过滤键。 md5:c4b4987b15b2caf0
	stackFilterKeyLocal = "/errors/gerror/gerror"
)

var (
	// goRootForFilter 用于开发环境中栈过滤的目的。 md5:df10489e92979e5e
	goRootForFilter = runtime.GOROOT()
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}
}

// Error 实现了 Error 接口，它返回所有的错误信息作为字符串。 md5:916d521fe191e82f
// ff:
// err:
func (err *Error) Error() string {
	if err == nil {
		return ""
	}
	errStr := err.text
	if errStr == "" && err.code != nil {
		errStr = err.code.Message()
	}
	if err.error != nil {
		if errStr != "" {
			errStr += ": "
		}
		errStr += err.error.Error()
	}
	return errStr
}

// Cause返回根本原因错误。 md5:c43631d8af1a0815
// ff:
// err:
func (err *Error) Cause() error {
	if err == nil {
		return nil
	}
	loop := err
	for loop != nil {
		if loop.error != nil {
			if e, ok := loop.error.(*Error); ok {
				// Internal Error struct.
				loop = e
			} else if e, ok := loop.error.(ICause); ok {
				// 实现ApiCause接口的其他错误。 md5:50f12dd0449fe932
				return e.Cause()
			} else {
				return loop.error
			}
		} else {
// 返回循环
//
// 以兼容 https://github.com/pkg/errors 中的 Case 情况。
// md5:a923900fc4a93e9d
			return errors.New(loop.text)
		}
	}
	return nil
}

// Current 创建并返回当前级别的错误。如果当前级别错误为 nil，则返回 nil。
// md5:d8b26e22ec63a837
// ff:
// err:
func (err *Error) Current() error {
	if err == nil {
		return nil
	}
	return &Error{
		error: nil,
		stack: err.stack,
		text:  err.text,
		code:  err.code,
	}
}

// Unwrap 是函数 `Next` 的别名。
// 它只是为了实现自 Go 1.17 版本的stdlib库中的 `errors.Unwrap`。
// md5:4ab7dcc4181801cd
// ff:
// err:
func (err *Error) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.error
}

// Equal 判断当前错误 `err` 是否等于目标错误 `target`。请注意，在默认的错误比较中，如果两个错误的 `code` 和 `text` 都相同，那么它们将被视为相等。
// md5:6256ec44e7b04b0e
// ff:
// err:
// target:
func (err *Error) Equal(target error) bool {
	if err == target {
		return true
	}
	// 代码应该保持不变。
	// 注意，如果两个错误的代码都为`nil`，那么它们也会被视为相等。
	// md5:9cd5037f48adc142
	if err.code != Code(target) {
		return false
	}
	// 文本内容应保持一致。 md5:950f9f350f074b9c
	if err.text != fmt.Sprintf(`%-s`, target) {
		return false
	}
	return true
}

	// Is 判断当前错误 `err` 是否在其嵌套错误中包含目标错误 `target`。这是为了实现从 Go 1.17 版本开始的标准库中的 errors.Is 接口。
	// md5:dfc92c8d3ba58133
// ff:
// err:
// target:
func (err *Error) Is(target error) bool {
	if Equal(err, target) {
		return true
	}
	nextErr := err.Unwrap()
	if nextErr == nil {
		return false
	}
	if Equal(nextErr, target) {
		return true
	}
	if e, ok := nextErr.(IIs); ok {
		return e.Is(target)
	}
	return false
}
