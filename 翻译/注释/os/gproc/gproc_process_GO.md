
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
// Process is the struct for a single process.
<原文结束>

# <翻译开始>
// Process是用于单个进程的结构体。
# <翻译结束>


<原文开始>
// NewProcess creates and returns a new Process.
<原文结束>

# <翻译开始>
// NewProcess 创建并返回一个新的 Process。
# <翻译结束>


<原文开始>
// Exclude of current binary path.
<原文结束>

# <翻译开始>
// 排除当前二进制文件路径。
# <翻译结束>


<原文开始>
// NewProcessCmd creates and returns a process with given command and optional environment variable array.
<原文结束>

# <翻译开始>
// NewProcessCmd根据给定的命令和可选的环境变量数组创建并返回一个进程。
# <翻译结束>


<原文开始>
// Start starts executing the process in non-blocking way.
// It returns the pid if success, or else it returns an error.
<原文结束>

# <翻译开始>
// Start以非阻塞方式启动进程执行。
// 如果成功，返回pid；否则返回错误。
# <翻译结束>


<原文开始>
// Run executes the process in blocking way.
<原文结束>

# <翻译开始>
// Run以阻塞方式执行进程。
# <翻译结束>


<原文开始>
// Pid retrieves and returns the PID for the process.
<原文结束>

# <翻译开始>
// Pid 获取并返回当前进程的PID（进程标识符）
# <翻译结束>


<原文开始>
// Send sends custom data to the process.
<原文结束>

# <翻译开始>
// Send 向进程发送自定义数据。
# <翻译结束>


<原文开始>
// Release releases any resources associated with the Process p,
// rendering it unusable in the future.
// Release only needs to be called if Wait is not.
<原文结束>

# <翻译开始>
// Release 会释放与进程p关联的任何资源，
// 使其在未来无法使用。
// 只有在不调用Wait的情况下，才需要调用Release。
# <翻译结束>


<原文开始>
// Kill causes the Process to exit immediately.
<原文结束>

# <翻译开始>
// Kill 导致 Process 立即退出。
# <翻译结束>


<原文开始>
// It ignores this error, just log it.
<原文结束>

# <翻译开始>
// 它忽略这个错误，仅将其记录到日志中。
# <翻译结束>


<原文开始>
// Signal sends a signal to the Process.
// Sending Interrupt on Windows is not implemented.
<原文结束>

# <翻译开始>
// Signal 向 Process 发送一个信号。
// 在 Windows 系统上发送 Interrupt 信号尚未实现。
# <翻译结束>


<原文开始>
// OpenTelemetry for command.
<原文结束>

# <翻译开始>
// OpenTelemetry 用于命令。
# <翻译结束>


<原文开始>
// OpenTelemetry propagation.
<原文结束>

# <翻译开始>
// OpenTelemetry 传播
# <翻译结束>

