// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gerror

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	
	"github.com/888go/goframe/errors/gcode"
)

// Error 是自定义错误类型，用于提供额外功能。
type Error struct {
	error error      // Wrapped error.
	stack stack      // Stack 数组，用于记录当此错误创建或被包裹时的堆栈信息。
	text  string     // 当 Error 被创建时的自定义错误文本，如果其代码不为空，则可能为空。
	code  gcode.Code // 如果必要，此处为错误代码。
}

const (
	// 用于当前错误模块路径的过滤键。
	stackFilterKeyLocal = "/errors/gerror/gerror"
)

var (
	// goRootForFilter 用于在开发环境目的中进行堆栈过滤。
	goRootForFilter = runtime.GOROOT()
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}
}

// Error 实现了 Error 接口，它将所有错误以字符串形式返回。
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

// Cause 返回根本原因错误。
func (err *Error) Cause() error {
	if err == nil {
		return nil
	}
	loop := err
	for loop != nil {
		if loop.error != nil {
			if e, ok := loop.error.(*Error); ok {
				// 内部错误结构体。
				loop = e
			} else if e, ok := loop.error.(ICause); ok {
				// 其他实现了ApiCause接口的错误类型。
				return e.Cause()
			} else {
				return loop.error
			}
		} else {
// 返回循环
//
// 为了与 https://github.com/pkg/errors 包中的 Case 兼容。
			return errors.New(loop.text)
		}
	}
	return nil
}

// Current函数创建并返回当前层级的错误信息。
// 如果当前层级的错误信息为nil，则该函数返回nil。
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
// 它仅为实现从 Go 1.17 版本开始的stdlib errors.Unwrap 接口而存在。
func (err *Error) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.error
}

// Equal 判断当前错误 `err` 是否与目标错误 `target` 相等。
// 请注意，在默认的 `Error` 比较方式下，
// 如果两个错误的 `code` 和 `text` 都相同，则认为它们是相同的错误。
func (err *Error) Equal(target error) bool {
	if err == target {
		return true
	}
// 代码应当保持一致。
// 注意，如果两个错误的code都是`nil`，则认为它们也是相等的。
	if err.code != Code(target) {
		return false
	}
	// 文本内容应该保持一致。
	if err.text != fmt.Sprintf(`%-s`, target) {
		return false
	}
	return true
}

// Is 报告当前错误 `err` 在其链式错误中是否包含错误 `target`。
// 这只是为了实现从 Go 1.17 版本开始的stdlib errors.Is功能。
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
