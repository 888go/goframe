// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package intlog 提供内部日志功能，仅用于 GoFrame 开发使用。
package intlog

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"time"
	
	"go.opentelemetry.io/otel/trace"
	
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/internal/utils"
)

const (
	stackFilterKey = "/internal/intlog"
)

// Print 通过 fmt.Println 打印 `v`（并附带换行）。参数 `v` 可以是多个变量。
func Print(ctx context.Context, v ...interface{}) {
	if !utils.IsDebugEnabled() {
		return
	}
	doPrint(ctx, fmt.Sprint(v...), false)
}

// Printf 使用 fmt.Printf 格式化并打印 `v`。其中参数 `v` 可以是多个变量。
// ```go
// Printf 函数通过 fmt.Printf 格式化输出 `v`。
// 注意，这里的 `v` 参数可以接受多个变量。
func Printf(ctx context.Context, format string, v ...interface{}) {
	if !utils.IsDebugEnabled() {
		return
	}
	doPrint(ctx, fmt.Sprintf(format, v...), false)
}

// Error 使用 fmt.Println 打印 `v`（并附带换行）。参数 `v` 可以是多个变量。
func Error(ctx context.Context, v ...interface{}) {
	if !utils.IsDebugEnabled() {
		return
	}
	doPrint(ctx, fmt.Sprint(v...), true)
}

// Errorf使用fmt.Printf格式化方式打印变量v，格式字符串为format。
func Errorf(ctx context.Context, format string, v ...interface{}) {
	if !utils.IsDebugEnabled() {
		return
	}
	doPrint(ctx, fmt.Sprintf(format, v...), true)
}

// PrintFunc 用于打印函数 `f` 的输出结果。
// 只有在调试模式开启时，才会调用函数 `f`。
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

// ErrorFunc 用于打印函数 `f` 的输出结果。
// 只有在调试模式开启的情况下，才会调用函数 `f`。
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

// traceIdStr 用于获取并返回日志输出的追踪ID字符串。
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

// file 返回调用文件名及其行号。
func file() string {
	_, p, l := gdebug.CallerWithFilter([]string{stackFilterKey})
	return fmt.Sprintf(`%s:%d`, filepath.Base(p), l)
}
