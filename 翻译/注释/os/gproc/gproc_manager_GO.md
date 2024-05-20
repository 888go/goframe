
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
// Manager is a process manager maintaining multiple processes.
<原文结束>

# <翻译开始>
// Manager 是一个管理多个进程的进程管理器。. md5:608ec304d3cca78a
# <翻译结束>


<原文开始>
// Process id to Process object mapping.
<原文结束>

# <翻译开始>
// 进程ID到进程对象的映射。. md5:e1aabd18695c16fa
# <翻译结束>


<原文开始>
// NewManager creates and returns a new process manager.
<原文结束>

# <翻译开始>
// NewManager 创建并返回一个新的进程管理器。. md5:bfef06576c70f94f
# <翻译结束>


<原文开始>
// NewProcess creates and returns a Process object.
<原文结束>

# <翻译开始>
// NewProcess 创建并返回一个进程对象。. md5:41e1fd6b109e05e7
# <翻译结束>


<原文开始>
// GetProcess retrieves and returns a Process object.
// It returns nil if it does not find the process with given `pid`.
<原文结束>

# <翻译开始>
// GetProcess 获取并返回一个Process对象。
// 如果找不到具有给定`pid`的进程，它将返回nil。
// md5:d5b11d4d0e9fa1a3
# <翻译结束>


<原文开始>
// AddProcess adds a process to current manager.
// It does nothing if the process with given `pid` does not exist.
<原文结束>

# <翻译开始>
// AddProcess 向当前管理器添加一个进程。
// 如果给定的 `pid` 对应的进程不存在，它不会做任何操作。
// md5:c51d5832fb1ce691
# <翻译结束>


<原文开始>
// RemoveProcess removes a process from current manager.
<原文结束>

# <翻译开始>
// RemoveProcess 从当前管理器中移除一个进程。. md5:0076407de3a7d26a
# <翻译结束>


<原文开始>
// Processes retrieves and returns all processes in current manager.
<原文结束>

# <翻译开始>
// Processes 获取并返回当前管理器中的所有进程。. md5:30ac76e5c68d45de
# <翻译结束>


<原文开始>
// Pids retrieves and returns all process id array in current manager.
<原文结束>

# <翻译开始>
// Pids 获取并返回当前管理器中的所有进程ID数组。. md5:a5ef21ec52c87400
# <翻译结束>


<原文开始>
// WaitAll waits until all process exit.
<原文结束>

# <翻译开始>
// WaitAll等待直到所有进程退出。. md5:1d27f65463fe8c00
# <翻译结束>


<原文开始>
// KillAll kills all processes in current manager.
<原文结束>

# <翻译开始>
// KillAll 在当前管理器中杀死所有进程。. md5:337f683854b75187
# <翻译结束>


<原文开始>
// SignalAll sends a signal `sig` to all processes in current manager.
<原文结束>

# <翻译开始>
// SignalAll 向当前管理器中的所有进程发送信号 `sig`。. md5:64ce0027dcad8808
# <翻译结束>


<原文开始>
// Send sends data bytes to all processes in current manager.
<原文结束>

# <翻译开始>
// Send 将数据字节发送到当前管理器中的所有进程。. md5:05d5ed3b0a5c7e3e
# <翻译结束>


<原文开始>
// SendTo sneds data bytes to specified processe in current manager.
<原文结束>

# <翻译开始>
// SendTo 向当前管理器中的指定进程发送数据字节。. md5:b477f09d2f5cca5f
# <翻译结束>


<原文开始>
// Clear removes all processes in current manager.
<原文结束>

# <翻译开始>
// Clear 会清除当前管理器中的所有进程。. md5:26053a86c2f65b33
# <翻译结束>


<原文开始>
// Size returns the size of processes in current manager.
<原文结束>

# <翻译开始>
// Size 返回当前管理器中进程的数量。. md5:ffaeaa3ed9b66ed1
# <翻译结束>

