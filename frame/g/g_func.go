// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package g

import (
	"context"
	"io"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/util/gutil"
)

// Go creates a new asynchronous goroutine function with specified recover function.
//
// The parameter `recoverFunc` is called when any panic during executing of `goroutineFunc`.
// If `recoverFunc` is given nil, it ignores the panic from `goroutineFunc` and no panic will
// throw to parent goroutine.
//
// But, note that, if `recoverFunc` also throws panic, such panic will be thrown to parent goroutine.

// ff:
func Go(
	ctx context.Context,
	goroutineFunc func(ctx context.Context),
	recoverFunc func(ctx context.Context, exception error),
) {
	gutil.Go(ctx, goroutineFunc, recoverFunc)
}

// NewVar returns a gvar.Var.

// ff:X泛型类
// safe:并发安全
// i:值
func NewVar(i interface{}, safe ...bool) *Var {
	return gvar.New(i, safe...)
}

// Wait is an alias of ghttp.Wait, which blocks until all the web servers shutdown.
// It's commonly used in multiple servers' situation.

// ff:Http类等待所有服务完成
func Wait() {
	ghttp.Wait()
}

// Listen is an alias of gproc.Listen, which handles the signals received and automatically
// calls registered signal handler functions.
// It blocks until shutdown signals received and all registered shutdown handlers done.

// ff:
func Listen() {
	gproc.Listen()
}

// Dump dumps a variable to stdout with more manually readable.

// ff:调试输出
// values:值s
func Dump(values ...interface{}) {
	gutil.Dump(values...)
}

// DumpTo writes variables `values` as a string in to `writer` with more manually readable

// ff:调试输出到Writer
// option:选项
// value:值
// writer:
func DumpTo(writer io.Writer, value interface{}, option gutil.DumpOption) {
	gutil.DumpTo(writer, value, option)
}

// DumpWithType acts like Dump, but with type information.
// Also see Dump.

// ff:调试输出并带类型
// values:值s
func DumpWithType(values ...interface{}) {
	gutil.DumpWithType(values...)
}

// DumpWithOption returns variables `values` as a string with more manually readable.

// ff:调试输出并带选项
// option:选项
// value:值s
func DumpWithOption(value interface{}, option gutil.DumpOption) {
	gutil.DumpWithOption(value, option)
}

// DumpJson pretty dumps json content to stdout.

// ff:调试输出json
// value:
func DumpJson(value any) {
	gutil.DumpJson(value)
}

// Throw throws an exception, which can be caught by TryCatch function.

// ff:异常输出
// exception:消息
func Throw(exception interface{}) {
	gutil.Throw(exception)
}

// Try implements try... logistics using internal panic...recover.
// It returns error if any exception occurs, or else it returns nil.

// ff:异常捕捉
// err:错误
// try:处理函数
// ctx:上下文
// ctx:上下文
func Try(ctx context.Context, try func(ctx context.Context)) (err error) {
	return gutil.Try(ctx, try)
}

// TryCatch implements try...catch... logistics using internal panic...recover.
// It automatically calls function `catch` if any exception occurs and passes the exception as an error.
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
	gutil.TryCatch(ctx, try, catch)
}

// IsNil checks whether given `value` is nil.
// Parameter `traceSource` is used for tracing to the source variable if given `value` is type
// of pointer that also points to a pointer. It returns nil if the source is nil when `traceSource`
// is true.
// Note that it might use reflect feature which affects performance a little.

// ff:是否为Nil
// yx:true
// traceSource:
// value:
func IsNil(value interface{}, traceSource ...bool) bool {
	return empty.IsNil(value, traceSource...)
}

// IsEmpty checks whether given `value` empty.
// It returns true if `value` is in: 0, nil, false, "", len(slice/map/chan) == 0.
// Or else it returns true.
//
// The parameter `traceSource` is used for tracing to the source variable if given `value` is type of pointer
// that also points to a pointer. It returns true if the source is empty when `traceSource` is true.
// Note that it might use reflect feature which affects performance a little.

// ff:是否为空
// traceSource:追踪到源变量
// value:值
func IsEmpty(value interface{}, traceSource ...bool) bool {
	return empty.IsEmpty(value, traceSource...)
}

// RequestFromCtx retrieves and returns the Request object from context.

// ff:Http类上下文取请求对象
// ctx:上下文
func RequestFromCtx(ctx context.Context) *ghttp.Request {
	return ghttp.RequestFromCtx(ctx)
}
