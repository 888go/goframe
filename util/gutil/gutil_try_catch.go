// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gutil

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// Throw throws out an exception, which can be caught be TryCatch or recover.

// ff:异常输出
// exception:消息
func Throw(exception interface{}) {
	panic(exception)
}

// Try implements try... logistics using internal panic...recover.
// It returns error if any exception occurs, or else it returns nil.

// ff:异常捕捉
// err:错误
// try:处理函数
// ctx:上下文
// ctx:上下文
func Try(ctx context.Context, try func(ctx context.Context)) (err error) {
	if try == nil {
		return
	}
	defer func() {
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok && gerror.HasStack(v) {
				err = v
			} else {
				err = gerror.NewCodef(gcode.CodeInternalPanic, "%+v", exception)
			}
		}
	}()
	try(ctx)
	return
}

// TryCatch implements `try...catch..`. logistics using internal `panic...recover`.
// It automatically calls function `catch` if any exception occurs and passes the exception as an error.
// If `catch` is given nil, it ignores the panic from `try` and no panic will throw to parent goroutine.
//
// But, note that, if function `catch` also throws panic, the current goroutine will panic.

// ff:异常捕捉并带异常处理
// catch:异常处理函数
// exception:错误
// ctx:上下文
// try:处理函数
// ctx:上下文
// ctx:上下文
func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error)) {
	if try == nil {
		return
	}
	if exception := Try(ctx, try); exception != nil && catch != nil {
		catch(ctx, exception)
	}
}
