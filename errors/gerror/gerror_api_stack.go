// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gerror
import (
	"runtime"
	)
// stack代表一个程序计数器的堆栈。
type stack []uintptr

const (
	// maxStackDepth 标记了错误回溯时的最大堆栈深度。
	maxStackDepth = 64
)

// Cause 返回 `err` 的根本原因错误。
func Cause(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(ICause); ok {
		return e.Cause()
	}
	if e, ok := err.(IUnwrap); ok {
		return Cause(e.Unwrap())
	}
	return err
}

// Stack 返回调用栈信息作为字符串。
// 如果`err`不支持堆栈信息，则直接返回错误字符串。
func Stack(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(IStack); ok {
		return e.Stack()
	}
	return err.Error()
}

// Current函数创建并返回当前层级的错误信息。
// 如果当前层级的错误信息为nil，则该函数返回nil。
func Current(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(ICurrent); ok {
		return e.Current()
	}
	return err
}

// Unwrap 返回下一层级的错误。
// 如果当前层级的错误或下一层级的错误为 nil，它将返回 nil。
func Unwrap(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(IUnwrap); ok {
		return e.Unwrap()
	}
	return nil
}

// HasStack 检查并报告 `err` 是否实现了接口 `gerror.IStack`。
func HasStack(err error) bool {
	_, ok := err.(IStack)
	return ok
}

// Equal 判断当前错误 `err` 是否等于目标错误 `target`。
// 请注意，在默认的 `Error` 比较逻辑中，
// 如果两个错误的 `code` 和 `text` 都相同，则认为它们是相同的错误。
func Equal(err, target error) bool {
	if err == target {
		return true
	}
	if e, ok := err.(IEqual); ok {
		return e.Equal(target)
	}
	if e, ok := target.(IEqual); ok {
		return e.Equal(err)
	}
	return false
}

// Is 报告当前错误 `err` 在其链式错误中是否包含错误 `target`。
// 这只是为了实现从 Go 1.17 版本开始的stdlib errors.Is功能。
func Is(err, target error) bool {
	if e, ok := err.(IIs); ok {
		return e.Is(target)
	}
	return false
}

// HasError 是 Is 的别名，它具有更易于理解的语义。
func HasError(err, target error) bool {
	return Is(err, target)
}

// callers 返回调用栈的调用者信息。
// 注意，此处它仅获取调用者内存地址数组，并非调用者详细信息。
func callers(skip ...int) stack {
	var (
		pcs [maxStackDepth]uintptr
		n   = 3
	)
	if len(skip) > 0 {
		n += skip[0]
	}
	return pcs[:runtime.Callers(n, pcs[:])]
}
