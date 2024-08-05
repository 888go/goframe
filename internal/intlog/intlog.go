// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包intlog为GoFrame开发内部日志提供支持，仅限于内部使用。 md5:f26290ffdfc859a6
package intlog

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/gogf/gf/v2/debug/gdebug"
	"github.com/gogf/gf/v2/internal/utils"
)

const (
	stackFilterKey = "/internal/intlog"
)

// Print 使用 fmt.Println 将 `v` 打印并附带换行。
// 参数 `v` 可以是多个变量。
// md5:38560020d990b134
func Print(ctx context.Context, v ...interface{}) {
	if !utils.IsDebugEnabled() {
		return
	}
	doPrint(ctx, fmt.Sprint(v...), false)
}

// Printf 使用 fmt.Printf 函数，按照格式 `format` 打印变量 `v`。
// 参数 `v` 可以是多个变量。
// md5:4791bf475aaad1f3
func Printf(ctx context.Context, format string, v ...interface{}) {
	if !utils.IsDebugEnabled() {
		return
	}
	doPrint(ctx, fmt.Sprintf(format, v...), false)
}

// Error 使用 fmt.Println 在新的一行打印 `v`。
// 参数 `v` 可以是多个变量。
// md5:25d1ed8d9d342a20
func Error(ctx context.Context, v ...interface{}) {
	if !utils.IsDebugEnabled() {
		return
	}
	doPrint(ctx, fmt.Sprint(v...), true)
}

// Errorf 使用 fmt.Printf 将格式化字符串 `format` 和变量 `v` 打印出来。 md5:d42411e861c4fdc7
func Errorf(ctx context.Context, format string, v ...interface{}) {
	if !utils.IsDebugEnabled() {
		return
	}
	doPrint(ctx, fmt.Sprintf(format, v...), true)
}

// PrintFunc 打印来自函数 `f` 的输出。
// 仅当调试模式启用时，它才会调用函数 `f`。
// md5:c3e57b3168c59983
func PrintFunc(ctx context.Context, f func() string) {
	if !utils.IsDebugEnabled() {
		return
	}
	s := f()
	if s == "" {
		return
	}
	doPrint(ctx, s, false)
}

// ErrorFunc 打印函数 `f` 的输出。
// 只有在调试模式启用时，它才会调用函数 `f`。
// md5:0fe2f7c87344bf65
func ErrorFunc(ctx context.Context, f func() string) {
	if !utils.IsDebugEnabled() {
		return
	}
	s := f()
	if s == "" {
		return
	}
	doPrint(ctx, s, true)
}

func doPrint(ctx context.Context, content string, stack bool) {
	if !utils.IsDebugEnabled() {
		return
	}
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString(time.Now().Format("2006-01-02 15:04:05.000"))
	buffer.WriteString(" [INTE] ")
	buffer.WriteString(file())
	buffer.WriteString(" ")
	if s := traceIdStr(ctx); s != "" {
		buffer.WriteString(s + " ")
	}
	buffer.WriteString(content)
	buffer.WriteString("\n")
	if stack {
		buffer.WriteString("Caller Stack:\n")
		buffer.WriteString(gdebug.StackWithFilter([]string{stackFilterKey}))
	}
	fmt.Print(buffer.String())
}

// traceIdStr 用于获取并返回日志输出的trace id字符串。 md5:de8f3fd19e246332
func traceIdStr(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	spanCtx := trace.SpanContextFromContext(ctx)
	if traceId := spanCtx.TraceID(); traceId.IsValid() {
		return "{" + traceId.String() + "}"
	}
	return ""
}

// file 函数返回调用者的文件名及其行号。 md5:28828a137896ad84
func file() string {
	_, p, l := gdebug.CallerWithFilter([]string{stackFilterKey})
	return fmt.Sprintf(`%s:%d`, filepath.Base(p), l)
}
