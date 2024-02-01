
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Shell executes command `cmd` synchronously with given input pipe `in` and output pipe `out`.
// The command `cmd` reads the input parameters from input pipe `in`, and writes its output automatically
// to output pipe `out`.
<原文结束>

# <翻译开始>
// Shell 同步执行命令 `cmd`，并使用给定的输入管道 `in` 和输出管道 `out`。
// 命令 `cmd` 从输入管道 `in` 读取输入参数，并自动将其输出写入输出管道 `out`。
# <翻译结束>


<原文开始>
// ShellRun executes given command `cmd` synchronously and outputs the command result to the stdout.
<原文结束>

# <翻译开始>
// ShellRun 同步执行给定的命令 `cmd`，并将命令结果输出到标准输出（stdout）。
# <翻译结束>


<原文开始>
// ShellExec executes given command `cmd` synchronously and returns the command result.
<原文结束>

# <翻译开始>
// ShellExec 同步执行给定的命令 `cmd`，并返回命令执行结果。
# <翻译结束>


<原文开始>
// parseCommand parses command `cmd` into slice arguments.
//
// Note that it just parses the `cmd` for "cmd.exe" binary in windows, but it is not necessary
// parsing the `cmd` for other systems using "bash"/"sh" binary.
<原文结束>

# <翻译开始>
// parseCommand 将命令 `cmd` 解析为切片参数。
//
// 注意，它仅针对 Windows 中的 "cmd.exe" 二进制文件解析 `cmd`，但对于使用 "bash"/"sh" 二进制文件的其他系统，没有必要对 `cmd` 进行解析。
# <翻译结束>


<原文开始>
// Just for "cmd.exe" in windows.
<原文结束>

# <翻译开始>
// 仅为Windows中的"cmd.exe"设计。
# <翻译结束>


<原文开始>
// It should remove the first quote char.
<原文结束>

# <翻译开始>
// 它应该移除第一个引号字符。
# <翻译结束>


<原文开始>
// It should remove the last quote char.
<原文结束>

# <翻译开始>
// 它应该移除最后一个引号字符。
# <翻译结束>


<原文开始>
// getShell returns the shell command depending on current working operating system.
// It returns "cmd.exe" for windows, and "bash" or "sh" for others.
<原文结束>

# <翻译开始>
// getShell 函数根据当前操作系统返回相应的 shell 命令。
// 对于 Windows 系统，它返回 "cmd.exe"；对于其他系统，返回 "bash" 或 "sh"。
# <翻译结束>


<原文开始>
// Check the default binary storage path.
<原文结束>

# <翻译开始>
// 检查默认的二进制存储路径。
# <翻译结束>


<原文开始>
// getShellOption returns the shell option depending on current working operating system.
// It returns "/c" for windows, and "-c" for others.
<原文结束>

# <翻译开始>
// getShellOption 根据当前操作系统返回相应的 shell 选项。
// 对于 Windows 系统，返回 "/c"；对于其他系统，返回 "-c"。
# <翻译结束>


<原文开始>
// tracingEnvFromCtx converts OpenTelemetry propagation data as environment variables.
<原文结束>

# <翻译开始>
// tracingEnvFromCtx 将 OpenTelemetry 传播数据转换为环境变量。
# <翻译结束>


<原文开始>
// Else search the env PATH.
<原文结束>

# <翻译开始>
// 否则在环境变量PATH中搜索。
# <翻译结束>

