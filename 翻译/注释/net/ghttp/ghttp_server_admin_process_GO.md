
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
// Allow executing management command after server starts after this interval in milliseconds.
<原文结束>

# <翻译开始>
	// 允许在服务器启动后经过此毫秒间隔执行管理命令。 md5:0a7e1d2b4fe2af39
# <翻译结束>


<原文开始>
// serverActionLocker is the locker for server administration operations.
<原文结束>

# <翻译开始>
	// serverActionLocker 是用于服务器管理操作的锁。 md5:3de592f90d7f4ae4
# <翻译结束>


<原文开始>
// serverActionLastTime is timestamp in milliseconds of last administration operation.
<原文结束>

# <翻译开始>
	// serverActionLastTime 是上一次管理操作的时间戳（以毫秒为单位）。 md5:df445bcc172577e2
# <翻译结束>


<原文开始>
// serverProcessStatus is the server status for operation of current process.
<原文结束>

# <翻译开始>
	// serverProcessStatus是当前进程运行时的服务器状态。 md5:3d55829242522190
# <翻译结束>


<原文开始>
// RestartAllServer restarts all the servers of the process gracefully.
// The optional parameter `newExeFilePath` specifies the new binary file for creating process.
<原文结束>

# <翻译开始>
// RestartAllServer 优雅地重启进程中的所有服务器。
// 可选参数 `newExeFilePath` 指定了用于创建进程的新二进制文件路径。
// md5:cd148e150eddefe2
# <翻译结束>


<原文开始>
// ShutdownAllServer shuts down all servers of current process gracefully.
<原文结束>

# <翻译开始>
// ShutdownAllServer 优雅地关闭当前进程中的所有服务器。 md5:1eb1bf001c79c66c
# <翻译结束>


<原文开始>
// checkProcessStatus checks the server status of current process.
<原文结束>

# <翻译开始>
// checkProcessStatus 检查当前进程的服务器状态。 md5:f49e9c4fdac4de86
# <翻译结束>


<原文开始>
// checkActionFrequency checks the operation frequency.
// It returns error if it is too frequency.
<原文结束>

# <翻译开始>
// checkActionFrequency 检查操作频率。
// 如果频率过高，它会返回错误。
// md5:b5db2b4c0ba2cdf7
# <翻译结束>


<原文开始>
// forkReloadProcess creates a new child process and copies the fd to child process.
<原文结束>

# <翻译开始>
// forkReloadProcess 创建一个新的子进程，并将文件描述符复制到子进程中。 md5:5de49cf62f76603e
# <翻译结束>


<原文开始>
// forkRestartProcess creates a new server process.
<原文结束>

# <翻译开始>
// forkRestartProcess 创建一个新的服务器进程。 md5:f786ce6758d0d9ed
# <翻译结束>


<原文开始>
// getServerFdMap returns all the servers name to file descriptor mapping as map.
<原文结束>

# <翻译开始>
// getServerFdMap 返回所有服务器名称到文件描述符映射的map。 md5:dd5b6c5b0372c1b6
# <翻译结束>


<原文开始>
// bufferToServerFdMap converts binary content to fd map.
<原文结束>

# <翻译开始>
// bufferToServerFdMap 将二进制内容转换为fd映射。 md5:f02ae7f98f43f216
# <翻译结束>


<原文开始>
// restartWebServers restarts all servers.
<原文结束>

# <翻译开始>
// restartWebServers 重启所有服务器。 md5:cad06ab5309d1725
# <翻译结束>


<原文开始>
		// Controlled by web page.
		// It should ensure the response wrote to client and then close all servers gracefully.
<原文结束>

# <翻译开始>
		// 由网页控制。
		// 它应该确保响应已写入客户端，然后优雅地关闭所有服务器。
		// md5:a5b2bfe1eb0f3681
# <翻译结束>


<原文开始>
// shutdownWebServersGracefully gracefully shuts down all servers.
<原文结束>

# <翻译开始>
// shutdownWebServersGracefully 延长关闭所有服务器。 md5:032a0ea9c1919f82
# <翻译结束>


<原文开始>
// forceCloseWebServers forced shuts down all servers.
<原文结束>

# <翻译开始>
// forceCloseWebServers 强制关闭所有服务器。 md5:e7c5bd88a9acbd9e
# <翻译结束>


<原文开始>
// handleProcessMessage receives and handles the message from processes,
// which are commonly used for graceful reloading feature.
<原文结束>

# <翻译开始>
// handleProcessMessage 接收并处理来自进程的消息，
// 这通常用于实现优雅重启功能。
// md5:80ebd3c82cd48199
# <翻译结束>

