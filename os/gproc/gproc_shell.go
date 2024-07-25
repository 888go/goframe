// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gproc

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/gogf/gf/v2/os/gfile"
)

// Shell 函数同步地执行命令 `cmd`，并使用给定的输入管道 `in` 和输出管道 `out`。
// 命令 `cmd` 从输入管道 `in` 读取输入参数，并自动将其输出写入到输出管道 `out`。 md5:6690cb7819cb1af8
func Shell(ctx context.Context, cmd string, out io.Writer, in io.Reader) error {
	p := NewProcess(
		getShell(),
		append([]string{getShellOption()}, parseCommand(cmd)...),
	)
	p.Stdin = in
	p.Stdout = out
	return p.Run(ctx)
}

// ShellRun 同步执行给定的命令 `cmd`，并将命令结果输出到stdout。 md5:b97833e7f1598d90
func ShellRun(ctx context.Context, cmd string) error {
	p := NewProcess(
		getShell(),
		append([]string{getShellOption()}, parseCommand(cmd)...),
	)
	return p.Run(ctx)
}

// ShellExec 同步执行给定命令 `cmd` 并返回命令结果。 md5:218406708403afde
func ShellExec(ctx context.Context, cmd string, environment ...[]string) (result string, err error) {
	var (
		buf = bytes.NewBuffer(nil)
		p   = NewProcess(
			getShell(),
			append([]string{getShellOption()}, parseCommand(cmd)...),
			environment...,
		)
	)
	p.Stdout = buf
	p.Stderr = buf
	err = p.Run(ctx)
	result = buf.String()
	return
}

// parseCommand 将命令 `cmd` 解析为参数切片。
//
// 注意，它只为 Windows 系统中的 "cmd.exe" 命令解析 `cmd`，而对于使用 "bash" 或 "sh" 命令的其他系统，这并不是必需的。 md5:22f6d16c6637aeee
func parseCommand(cmd string) (args []string) {
	return []string{cmd}
}

// getShell 根据当前操作系统返回相应的shell命令。
// 对于Windows系统，它返回 "cmd.exe"；对于其他系统，则返回 "bash" 或 "sh"。 md5:9b8e621dfd22db86
func getShell() string {
	switch runtime.GOOS {
	case "windows":
		return SearchBinary("cmd.exe")

	default:
		// 检查默认的二进制存储路径。 md5:11d55faa0b1f45a3
		if gfile.Exists("/bin/bash") {
			return "/bin/bash"
		}
		if gfile.Exists("/bin/sh") {
			return "/bin/sh"
		}
		// 否则，在环境PATH中搜索。 md5:73695e9885dbcbe8
		path := SearchBinary("bash")
		if path == "" {
			path = SearchBinary("sh")
		}
		return path
	}
}

// getShellOption 根据当前工作操作系统返回shell选项。
// 对于Windows，返回"/c"，对于其他系统，返回"-c"。 md5:e3515e6516946346
func getShellOption() string {
	switch runtime.GOOS {
	case "windows":
		return "/c"

	default:
		return "-c"
	}
}

// tracingEnvFromCtx 将 OpenTelemetry 传播数据转换为环境变量。 md5:ca513c78879da082
func tracingEnvFromCtx(ctx context.Context) []string {
	var (
		a = make([]string, 0)
		m = make(map[string]string)
	)
	otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(m))
	for k, v := range m {
		a = append(a, fmt.Sprintf(`%s=%s`, k, v))
	}
	return a
}
