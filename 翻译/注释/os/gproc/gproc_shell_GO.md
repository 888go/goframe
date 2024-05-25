
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Shell executes command `cmd` synchronously with given input pipe `in` and output pipe `out`.
// The command `cmd` reads the input parameters from input pipe `in`, and writes its output automatically
// to output pipe `out`.
<原文结束>

# <翻译开始>
// Shell 函数同步地执行命令 `cmd`，并使用给定的输入管道 `in` 和输出管道 `out`。
// 命令 `cmd` 从输入管道 `in` 读取输入参数，并自动将其输出写入到输出管道 `out`。
// md5:6690cb7819cb1af8
# <翻译结束>


<原文开始>
// ShellRun executes given command `cmd` synchronously and outputs the command result to the stdout.
<原文结束>

# <翻译开始>
// ShellRun 同步执行给定的命令 `cmd`，并将命令结果输出到stdout。 md5:b97833e7f1598d90
# <翻译结束>


<原文开始>
// ShellExec executes given command `cmd` synchronously and returns the command result.
<原文结束>

# <翻译开始>
// ShellExec 同步执行给定命令 `cmd` 并返回命令结果。 md5:218406708403afde
# <翻译结束>


<原文开始>
// parseCommand parses command `cmd` into slice arguments.
//
// Note that it just parses the `cmd` for "cmd.exe" binary in windows, but it is not necessary
// parsing the `cmd` for other systems using "bash"/"sh" binary.
<原文结束>

# <翻译开始>
// parseCommand 将命令 `cmd` 解析为参数切片。
//
// 注意，它只为 Windows 系统中的 "cmd.exe" 命令解析 `cmd`，而对于使用 "bash" 或 "sh" 命令的其他系统，这并不是必需的。
// md5:22f6d16c6637aeee
# <翻译结束>


<原文开始>
// getShell returns the shell command depending on current working operating system.
// It returns "cmd.exe" for windows, and "bash" or "sh" for others.
<原文结束>

# <翻译开始>
// getShell 根据当前操作系统返回相应的shell命令。
// 对于Windows系统，它返回 "cmd.exe"；对于其他系统，则返回 "bash" 或 "sh"。
// md5:9b8e621dfd22db86
# <翻译结束>


<原文开始>
// Check the default binary storage path.
<原文结束>

# <翻译开始>
// 检查默认的二进制存储路径。 md5:11d55faa0b1f45a3
# <翻译结束>


<原文开始>
// Else search the env PATH.
<原文结束>

# <翻译开始>
// 否则，在环境PATH中搜索。 md5:73695e9885dbcbe8
# <翻译结束>


<原文开始>
// getShellOption returns the shell option depending on current working operating system.
// It returns "/c" for windows, and "-c" for others.
<原文结束>

# <翻译开始>
// getShellOption 根据当前工作操作系统返回shell选项。
// 对于Windows，返回"/c"，对于其他系统，返回"-c"。
// md5:e3515e6516946346
# <翻译结束>


<原文开始>
// tracingEnvFromCtx converts OpenTelemetry propagation data as environment variables.
<原文结束>

# <翻译开始>
// tracingEnvFromCtx 将 OpenTelemetry 传播数据转换为环境变量。 md5:ca513c78879da082
# <翻译结束>

