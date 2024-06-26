// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gerror

import (
	"runtime"
)

// stack represents a stack of program counters.
type stack []uintptr

const (
	// maxStackDepth marks the max stack depth for error back traces.
	maxStackDepth = 64
)

// Cause returns the root cause error of `err`.

// ff:取根错误
// err:错误
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

// Stack returns the stack callers as string.
// It returns the error string directly if the `err` does not support stacks.

// ff:取文本
// err:错误
func Stack(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(IStack); ok {
		return e.Stack()
	}
	return err.Error()
}

// Current creates and returns the current level error.
// It returns nil if current level error is nil.

// ff:取当前错误
// err:错误
func Current(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(ICurrent); ok {
		return e.Current()
	}
	return err
}

// Unwrap returns the next level error.
// It returns nil if current level error or the next level error is nil.

// ff:取下一层错误
// err:错误
func Unwrap(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(IUnwrap); ok {
		return e.Unwrap()
	}
	return nil
}

// HasStack checks and reports whether `err` implemented interface `gerror.IStack`.

// ff:判断是否带堆栈
// err:错误
func HasStack(err error) bool {
	_, ok := err.(IStack)
	return ok
}

// Equal reports whether current error `err` equals to error `target`.
// Please note that, in default comparison logic for `Error`,
// the errors are considered the same if both the `code` and `text` of them are the same.

// ff:是否相等
// target:待比较
// err:
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

// Is reports whether current error `err` has error `target` in its chaining errors.
// It is just for implements for stdlib errors.Is from Go version 1.17.

// ff:是否包含
// target:待比较
// err:
func Is(err, target error) bool {
	if e, ok := err.(IIs); ok {
		return e.Is(target)
	}
	return false
}

// HasError is alias of Is, which more easily understanding semantics.

// ff:HasError别名
// target:
// err:
func HasError(err, target error) bool {
	return Is(err, target)
}

// callers returns the stack callers.
// Note that it here just retrieves the caller memory address array not the caller information.
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
