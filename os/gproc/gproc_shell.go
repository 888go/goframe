// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 进程类

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"runtime"
	
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/text/gstr"
)

// Shell 同步执行命令 `cmd`，并使用给定的输入管道 `in` 和输出管道 `out`。
// 命令 `cmd` 从输入管道 `in` 读取输入参数，并自动将其输出写入输出管道 `out`。
func Shell(ctx context.Context, cmd string, out io.Writer, in io.Reader) error {
	p := NewProcess(
		getShell(),
		append([]string{getShellOption()}, parseCommand(cmd)...),
	)
	p.Stdin = in
	p.Stdout = out
	return p.Run(ctx)
}

// ShellRun 同步执行给定的命令 `cmd`，并将命令结果输出到标准输出（stdout）。
func ShellRun(ctx context.Context, cmd string) error {
	p := NewProcess(
		getShell(),
		append([]string{getShellOption()}, parseCommand(cmd)...),
	)
	return p.Run(ctx)
}

// ShellExec 同步执行给定的命令 `cmd`，并返回命令执行结果。
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

// parseCommand 将命令 `cmd` 解析为切片参数。
//
// 注意，它仅针对 Windows 中的 "cmd.exe" 二进制文件解析 `cmd`，但对于使用 "bash"/"sh" 二进制文件的其他系统，没有必要对 `cmd` 进行解析。
func parseCommand(cmd string) (args []string) {
	if runtime.GOOS != "windows" {
		return []string{cmd}
	}
	// 仅为Windows中的"cmd.exe"设计。
	var argStr string
	var firstChar, prevChar, lastChar1, lastChar2 byte
	array := 文本类.X分割并忽略空值(cmd, " ")
	for _, v := range array {
		if len(argStr) > 0 {
			argStr += " "
		}
		firstChar = v[0]
		lastChar1 = v[len(v)-1]
		lastChar2 = 0
		if len(v) > 1 {
			lastChar2 = v[len(v)-2]
		}
		if prevChar == 0 && (firstChar == '"' || firstChar == '\'') {
			// 它应该移除第一个引号字符。
			argStr += v[1:]
			prevChar = firstChar
		} else if prevChar != 0 && lastChar2 != '\\' && lastChar1 == prevChar {
			// 它应该移除最后一个引号字符。
			argStr += v[:len(v)-1]
			args = append(args, argStr)
			argStr = ""
			prevChar = 0
		} else if len(argStr) > 0 {
			argStr += v
		} else {
			args = append(args, v)
		}
	}
	return
}

// getShell 函数根据当前操作系统返回相应的 shell 命令。
// 对于 Windows 系统，它返回 "cmd.exe"；对于其他系统，返回 "bash" 或 "sh"。
func getShell() string {
	switch runtime.GOOS {
	case "windows":
		return SearchBinary("cmd.exe")

	default:
		// 检查默认的二进制存储路径。
		if 文件类.X是否存在("/bin/bash") {
			return "/bin/bash"
		}
		if 文件类.X是否存在("/bin/sh") {
			return "/bin/sh"
		}
		// 否则在环境变量PATH中搜索。
		path := SearchBinary("bash")
		if path == "" {
			path = SearchBinary("sh")
		}
		return path
	}
}

// getShellOption 根据当前操作系统返回相应的 shell 选项。
// 对于 Windows 系统，返回 "/c"；对于其他系统，返回 "-c"。
func getShellOption() string {
	switch runtime.GOOS {
	case "windows":
		return "/c"

	default:
		return "-c"
	}
}

// tracingEnvFromCtx 将 OpenTelemetry 传播数据转换为环境变量。
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
