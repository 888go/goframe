// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 进程类

import (
	"context"
	"io"
)

// MustShell 行为与 Shell 相同，但如果发生任何错误则会引发 panic。 md5:8ffd357cf1ea4dbc
func MustShell(ctx context.Context, cmd string, out io.Writer, in io.Reader) {
	if err := Shell(ctx, cmd, out, in); err != nil {
		panic(err)
	}
}

// MustShellRun 执行与 ShellRun 相同的操作，但如果发生任何错误，则会引发恐慌。 md5:b0e6d628208193e7
func MustShellRun(ctx context.Context, cmd string) {
	if err := ShellRun(ctx, cmd); err != nil {
		panic(err)
	}
}

// MustShellExec 执行类似于 ShellExec，但如果发生任何错误，它将引发恐慌。 md5:9754cecde7636273
func MustShellExec(ctx context.Context, cmd string, environment ...[]string) string {
	result, err := ShellExec(ctx, cmd, environment...)
	if err != nil {
		panic(err)
	}
	return result
}
