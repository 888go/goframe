
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
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Allow executing management command after server starts after this interval in milliseconds.
<原文结束>

# <翻译开始>
// 在服务器启动后，允许在此间隔（以毫秒为单位）后执行管理命令。
# <翻译结束>


<原文开始>
// serverActionLocker is the locker for server administration operations.
<原文结束>

# <翻译开始>
// serverActionLocker 是针对服务器管理操作的锁。
# <翻译结束>


<原文开始>
// serverActionLastTime is timestamp in milliseconds of last administration operation.
<原文结束>

# <翻译开始>
// serverActionLastTime 是上一次管理操作发生时的时间戳，单位为毫秒。
# <翻译结束>


<原文开始>
// serverProcessStatus is the server status for operation of current process.
<原文结束>

# <翻译开始>
// serverProcessStatus 是当前进程运行操作的服务器状态。
# <翻译结束>


<原文开始>
// RestartAllServer restarts all the servers of the process gracefully.
// The optional parameter `newExeFilePath` specifies the new binary file for creating process.
<原文结束>

# <翻译开始>
// RestartAllServer 将优雅地重启进程中的所有服务器。
// 可选参数 `newExeFilePath` 指定了用于创建新进程的二进制文件。
# <翻译结束>


<原文开始>
// ShutdownAllServer shuts down all servers of current process gracefully.
<原文结束>

# <翻译开始>
// ShutdownAllServer 将当前进程中的所有服务器优雅地关闭。
# <翻译结束>


<原文开始>
// checkProcessStatus checks the server status of current process.
<原文结束>

# <翻译开始>
// checkProcessStatus 检查当前进程的服务器状态。
# <翻译结束>


<原文开始>
// checkActionFrequency checks the operation frequency.
// It returns error if it is too frequency.
<原文结束>

# <翻译开始>
// checkActionFrequency 检查操作频率。
// 如果操作过于频繁，则返回错误。
# <翻译结束>


<原文开始>
// forkReloadProcess creates a new child process and copies the fd to child process.
<原文结束>

# <翻译开始>
// forkReloadProcess 创建一个新的子进程，并将文件描述符(fd)复制到子进程中。
# <翻译结束>


<原文开始>
// forkRestartProcess creates a new server process.
<原文结束>

# <翻译开始>
// forkRestartProcess 创建一个新的服务进程。
# <翻译结束>


<原文开始>
// getServerFdMap returns all the servers name to file descriptor mapping as map.
<原文结束>

# <翻译开始>
// getServerFdMap 返回一个映射，其中包含了所有服务器名称到文件描述符的映射关系，以map形式返回。
# <翻译结束>


<原文开始>
// bufferToServerFdMap converts binary content to fd map.
<原文结束>

# <翻译开始>
// bufferToServerFdMap 将二进制内容转换为文件描述符映射。
# <翻译结束>


<原文开始>
// restartWebServers restarts all servers.
<原文结束>

# <翻译开始>
// 重启Web服务器 restartWebServers 函数会重启所有服务器。
# <翻译结束>







<原文开始>
		// Controlled by web page.
		// It should ensure the response wrote to client and then close all servers gracefully.
<原文结束>

# <翻译开始>
// 由网页控制。
// 应确保响应已写入客户端，然后优雅地关闭所有服务器。
# <翻译结束>


<原文开始>
// shutdownWebServersGracefully gracefully shuts down all servers.
<原文结束>

# <翻译开始>
// shutdownWebServersGracefully 优雅地关闭所有服务器。
# <翻译结束>


<原文开始>
// forceCloseWebServers forced shuts down all servers.
<原文结束>

# <翻译开始>
// forceCloseWebServers 强制关闭所有服务器。
# <翻译结束>


<原文开始>
// handleProcessMessage receives and handles the message from processes,
// which are commonly used for graceful reloading feature.
<原文结束>

# <翻译开始>
// handleProcessMessage 接收并处理来自进程的消息，
// 这通常用于优雅重载功能。
# <翻译结束>


<原文开始>
// Controlled by signal.
<原文结束>

# <翻译开始>
// 由信号控制。
# <翻译结束>

