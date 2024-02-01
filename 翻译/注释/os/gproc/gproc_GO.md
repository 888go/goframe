
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
// Package gproc implements management and communication for processes.
<原文结束>

# <翻译开始>
// 包gproc实现了对进程的管理和通信功能。
# <翻译结束>


<原文开始>
// processPid is the pid of current process.
<原文结束>

# <翻译开始>
// processPid 是当前进程的进程ID。
# <翻译结束>


<原文开始>
// processStartTime is the start time of current process.
<原文结束>

# <翻译开始>
// processStartTime 是当前进程的启动时间。
# <翻译结束>


<原文开始>
// Pid returns the pid of current process.
<原文结束>

# <翻译开始>
// Pid 返回当前进程的进程ID。
# <翻译结束>


<原文开始>
// PPid returns the custom parent pid if exists, or else it returns the system parent pid.
<原文结束>

# <翻译开始>
// PPid 返回自定义父进程ID（如果存在），否则返回系统父进程ID。
# <翻译结束>


<原文开始>
// PPidOS returns the system parent pid of current process.
// Note that the difference between PPidOS and PPid function is that the PPidOS returns
// the system ppid, but the PPid functions may return the custom pid by gproc if the custom
// ppid exists.
<原文结束>

# <翻译开始>
// PPidOS 返回当前进程的系统父进程ID。
// 注意，PPidOS 和 PPid 函数之间的区别在于：PPidOS 返回的是系统的父进程ID，
// 但若存在自定义父进程ID，PPid 函数可能会返回由 gproc 提供的自定义进程ID。
# <翻译结束>


<原文开始>
// IsChild checks and returns whether current process is a child process.
// A child process is forked by another gproc process.
<原文结束>

# <翻译开始>
// IsChild 检查并返回当前进程是否为子进程。
// 子进程是由另一个 gproc 进程 fork（派生）出来的。
# <翻译结束>


<原文开始>
// SetPPid sets custom parent pid for current process.
<原文结束>

# <翻译开始>
// SetPPid 设置当前进程的自定义父进程ID。
# <翻译结束>


<原文开始>
// StartTime returns the start time of current process.
<原文结束>

# <翻译开始>
// StartTime 返回当前进程的启动时间。
# <翻译结束>


<原文开始>
// Uptime returns the duration which current process has been running
<原文结束>

# <翻译开始>
// Uptime 返回当前进程已经运行的时间间隔
# <翻译结束>


<原文开始>
// SearchBinary searches the binary `file` in current working folder and PATH environment.
<原文结束>

# <翻译开始>
// SearchBinary 在当前工作目录和PATH环境变量中搜索二进制文件 `file`。
# <翻译结束>


<原文开始>
// Check if it is absolute path of exists at current working directory.
<原文结束>

# <翻译开始>
// 检查给定路径是否为绝对路径，或者在当前工作目录中是否存在。
# <翻译结束>


<原文开始>
// SearchBinaryPath searches the binary `file` in PATH environment.
<原文结束>

# <翻译开始>
// SearchBinaryPath在PATH环境变量中搜索二进制文件`file`。
# <翻译结束>

