
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
// Manager is a process manager maintaining multiple processes.
<原文结束>

# <翻译开始>
// Manager 是一个进程管理器，用于维护多个进程。
# <翻译结束>


<原文开始>
// Process id to Process object mapping.
<原文结束>

# <翻译开始>
// 进程id到进程对象的映射。
# <翻译结束>


<原文开始>
// NewManager creates and returns a new process manager.
<原文结束>

# <翻译开始>
// NewManager 创建并返回一个新的进程管理器。
# <翻译结束>


<原文开始>
// NewProcess creates and returns a Process object.
<原文结束>

# <翻译开始>
// NewProcess 创建并返回一个 Process 对象。
# <翻译结束>


<原文开始>
// GetProcess retrieves and returns a Process object.
// It returns nil if it does not find the process with given `pid`.
<原文结束>

# <翻译开始>
// GetProcess 获取并返回一个 Process 对象。
// 如果未找到给定 `pid` 的进程，则返回 nil。
# <翻译结束>


<原文开始>
// AddProcess adds a process to current manager.
// It does nothing if the process with given `pid` does not exist.
<原文结束>

# <翻译开始>
// AddProcess 将一个进程添加到当前管理器中。
// 如果给定 `pid` 的进程不存在，则不做任何操作。
# <翻译结束>


<原文开始>
// RemoveProcess removes a process from current manager.
<原文结束>

# <翻译开始>
// RemoveProcess 从当前管理器中移除一个进程。
# <翻译结束>


<原文开始>
// Processes retrieves and returns all processes in current manager.
<原文结束>

# <翻译开始>
// Processes 获取并返回当前管理器中的所有进程。
# <翻译结束>


<原文开始>
// Pids retrieves and returns all process id array in current manager.
<原文结束>

# <翻译开始>
// Pids 获取并返回当前管理器中的所有进程ID数组。
# <翻译结束>


<原文开始>
// WaitAll waits until all process exit.
<原文结束>

# <翻译开始>
// WaitAll等待直到所有进程退出。
# <翻译结束>


<原文开始>
// KillAll kills all processes in current manager.
<原文结束>

# <翻译开始>
// KillAll杀掉当前管理器中的所有进程。
# <翻译结束>


<原文开始>
// SignalAll sends a signal `sig` to all processes in current manager.
<原文结束>

# <翻译开始>
// SignalAll 向当前管理器中的所有进程发送信号 `sig`。
# <翻译结束>


<原文开始>
// Send sends data bytes to all processes in current manager.
<原文结束>

# <翻译开始>
// Send 向当前管理器中的所有进程发送 data 字节。
# <翻译结束>


<原文开始>
// SendTo sneds data bytes to specified processe in current manager.
<原文结束>

# <翻译开始>
// SendTo 向当前管理器中指定进程发送数据字节。
# <翻译结束>


<原文开始>
// Clear removes all processes in current manager.
<原文结束>

# <翻译开始>
// Clear 清除当前管理器中的所有进程。
# <翻译结束>


<原文开始>
// Size returns the size of processes in current manager.
<原文结束>

# <翻译开始>
// Size 返回当前管理器中进程的数量。
# <翻译结束>

