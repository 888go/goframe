
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
// Process is the struct for a single process.
<原文结束>

# <翻译开始>
// Process 是表示单个进程的结构体。 md5:f6524ce6eee4a18b
# <翻译结束>


<原文开始>
// NewProcess creates and returns a new Process.
<原文结束>

# <翻译开始>
// NewProcess 创建并返回一个新的 Process。 md5:dbd46312fa39f087
# <翻译结束>


<原文开始>
// Exclude of current binary path.
<原文结束>

# <翻译开始>
		// 排除当前二进制文件的路径。 md5:a174ba38ac49d432
# <翻译结束>


<原文开始>
// NewProcessCmd creates and returns a process with given command and optional environment variable array.
<原文结束>

# <翻译开始>
// NewProcessCmd 创建并返回一个具有给定命令和可选环境变量数组的进程。 md5:01376a1e29c9935e
# <翻译结束>


<原文开始>
// Start starts executing the process in non-blocking way.
// It returns the pid if success, or else it returns an error.
<原文结束>

# <翻译开始>
// Start 以非阻塞方式开始执行进程。
// 如果成功，它将返回进程ID（pid），否则返回一个错误。
// md5:4607fc00f35e6338
# <翻译结束>


<原文开始>
// OpenTelemetry for command.
<原文结束>

# <翻译开始>
	// 为命令提供OpenTelemetry。 md5:46407dd5b38f692f
# <翻译结束>


<原文开始>
// OpenTelemetry propagation.
<原文结束>

# <翻译开始>
	// OpenTelemetry 传播。 md5:aecf3a0cccd13f96
# <翻译结束>


<原文开始>
// On Windows, this works and doesn't work on other platforms
<原文结束>

# <翻译开始>
	// 在 Windows 系统中，这可以工作，但在其他平台则无法工作. md5:9aac240ca7d717fe
# <翻译结束>


<原文开始>
// Run executes the process in blocking way.
<原文结束>

# <翻译开始>
// Run以阻塞方式执行进程。 md5:aeab1ddf5fca3d31
# <翻译结束>


<原文开始>
// Pid retrieves and returns the PID for the process.
<原文结束>

# <翻译开始>
// Pid 获取并返回进程的PID。 md5:7f6e89391a9d1aac
# <翻译结束>


<原文开始>
// Send sends custom data to the process.
<原文结束>

# <翻译开始>
// Send 向进程发送自定义数据。 md5:cb2381344fb13fd4
# <翻译结束>


<原文开始>
// Release releases any resources associated with the Process p,
// rendering it unusable in the future.
// Release only needs to be called if Wait is not.
<原文结束>

# <翻译开始>
// Release 释放与进程 p 关联的任何资源，使其将来无法使用。
// 只有在不调用 Wait 的情况下才需要调用 Release。
// md5:f3540c25ba14f0ee
# <翻译结束>


<原文开始>
// Kill causes the Process to exit immediately.
<原文结束>

# <翻译开始>
// Kill 立即导致 Process 终止。 md5:4bacb16ab3b9aebe
# <翻译结束>


<原文开始>
// It ignores this error, just log it.
<原文结束>

# <翻译开始>
	// 它忽略这个错误，仅记录日志。 md5:578bff85a58d16e8
# <翻译结束>


<原文开始>
// Signal sends a signal to the Process.
// Sending Interrupt on Windows is not implemented.
<原文结束>

# <翻译开始>
// Signal 向进程发送一个信号。
// 在Windows上发送Interrupt信号未实现。
// md5:c1afe56a9d236095
# <翻译结束>

