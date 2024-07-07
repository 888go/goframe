// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gerror

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
)

// NewCode creates and returns an error that has error code and given text.
// ff:创建错误码
// code:错误码
// text:
func NewCode(code gcode.Code, text ...string) error {
	return &Error{
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// NewCodef returns an error that has error code and formats as the given format and args.
// ff:创建错误码并格式化
// code:错误码
// format:
// args:
func NewCodef(code gcode.Code, format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// NewCodeSkip creates and returns an error which has error code and is formatted from given text.
// The parameter `skip` specifies the stack callers skipped amount.
// ff:创建错误码并跳过堆栈
// code:错误码
// skip:
// text:
func NewCodeSkip(code gcode.Code, skip int, text ...string) error {
	return &Error{
		stack: callers(skip),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// NewCodeSkipf returns an error that has error code and formats as the given format and args.
// The parameter `skip` specifies the stack callers skipped amount.
// ff:创建错误码并跳过堆栈与格式化
// code:错误码
// skip:
// format:
// args:
func NewCodeSkipf(code gcode.Code, skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCode wraps error with code and text.
// It returns nil if given err is nil.
// ff:多层错误码
// code:错误码
// err:
// text:
func WrapCode(code gcode.Code, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// WrapCodef wraps error with code and format specifier.
// It returns nil if given `err` is nil.
// ff:多层错误码并格式化
// code:错误码
// err:
// format:
// args:
func WrapCodef(code gcode.Code, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCodeSkip wraps error with code and text.
// It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
// ff:多层错误码并跳过堆栈
// code:错误码
// skip:
// err:
// text:
func WrapCodeSkip(code gcode.Code, skip int, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// WrapCodeSkipf wraps error with code and text that is formatted with given format and args.
// It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
// ff:多层错误码并跳过堆栈与格式化
// code:错误码
// skip:
// err:
// format:
// args:
func WrapCodeSkipf(code gcode.Code, skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// Code returns the error code of `current error`.
// It returns `CodeNil` if it has no error code neither it does not implement interface Code.
// ff:取错误码
// err:错误
func Code(err error) gcode.Code {
	if err == nil {
		return gcode.CodeNil
	}
	if e, ok := err.(ICode); ok {
		return e.Code()
	}
	if e, ok := err.(IUnwrap); ok {
		return Code(e.Unwrap())
	}
	return gcode.CodeNil
}

// HasCode checks and reports whether `err` has `code` in its chaining errors.
// ff:是否包含错误码
// err:错误
// code:错误码
func HasCode(err error, code gcode.Code) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(ICode); ok && code == e.Code() {
		return true
	}
	if e, ok := err.(IUnwrap); ok {
		return HasCode(e.Unwrap(), code)
	}
	return false
}
