// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gerror

import (
	"runtime"
)

// stack 表示一个程序计数器栈。 md5:974ffc07f2ddbe11
type stack []uintptr

const (
	// maxStackDepth 标记错误回溯的最大堆栈深度。 md5:e4020e5e182a92f7
	maxStackDepth = 64
)

// Cause 返回 `err` 的根本原因错误。 md5:3719c97ae5cd6a94
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

// Stack 返回调用堆栈作为字符串。
// 如果 `err` 不支持堆栈信息，它将直接返回错误字符串。 md5:bcba4c4f440cb0a7
func Stack(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(IStack); ok {
		return e.Stack()
	}
	return err.Error()
}

// Current 创建并返回当前级别的错误。如果当前级别错误为 nil，则返回 nil。 md5:d8b26e22ec63a837
func Current(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(ICurrent); ok {
		return e.Current()
	}
	return err
}

// Unwrap返回下一级错误。
// 如果当前级别或下一级错误为nil，则返回nil。 md5:97894c3cda2b0c4e
func Unwrap(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(IUnwrap); ok {
		return e.Unwrap()
	}
	return nil
}

// HasStack 检查并报告 `err` 是否实现了接口 `gerror.IStack`。 md5:f8ab57042234eea7
func HasStack(err error) bool {
	_, ok := err.(IStack)
	return ok
}

// Equal 检查当前错误 `err` 是否等于错误 `target`。
// 请注意，在默认的 `Error` 比较逻辑中，如果两个错误的 `code` 和 `text` 都相同，则认为它们是相同的。 md5:adaa63023ba44f79
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

// Is 判断当前错误 `err` 是否在其嵌套错误中包含目标错误 `target`。这是为了实现从 Go 1.17 版本开始的标准库中的 errors.Is 接口。 md5:dfc92c8d3ba58133
func Is(err, target error) bool {
	if e, ok := err.(IIs); ok {
		return e.Is(target)
	}
	return false
}

// HasError是Is的别名，它更易于理解语义。 md5:d84dea05563aadb6
func HasError(err, target error) bool {
	return Is(err, target)
}

// callers 返回调用栈的调用者。
// 注意，这里只是获取调用者内存地址数组，并非调用者的具体信息。 md5:9c65fc07c2395a65
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
