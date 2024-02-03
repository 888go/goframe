// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gproc

import (
	"context"
	"io"
)

// MustShell 的行为与 Shell 相同，但如果出现任何错误，它会引发 panic。
func MustShell(ctx context.Context, cmd string, out io.Writer, in io.Reader) {
	if err := Shell(ctx, cmd, out, in); err != nil {
		panic(err)
	}
}

// MustShellRun 的行为与 ShellRun 相同，但当出现任何错误时，它会触发 panic。
func MustShellRun(ctx context.Context, cmd string) {
	if err := ShellRun(ctx, cmd); err != nil {
		panic(err)
	}
}

// MustShellExec 的执行方式与 ShellExec 相同，但当出现任何错误时，它会触发 panic。
func MustShellExec(ctx context.Context, cmd string, environment ...[]string) string {
	result, err := ShellExec(ctx, cmd, environment...)
	if err != nil {
		panic(err)
	}
	return result
}
