
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
// PoolConn is a connection with pool feature for TCP.
// Note that it is NOT a pool or connection manager, it is just a TCP connection object.
<原文结束>

# <翻译开始>
// PoolConn 是一个具有连接池特性的TCP连接。
// 注意，它本身并不是一个连接池或连接管理器，而只是一个TCP连接对象。
// md5:ba1f65d3a4240a38
# <翻译结束>


<原文开始>
// Underlying connection object.
<原文结束>

# <翻译开始>
// 底层连接对象。. md5:b967bfe4f6e1fc27
# <翻译结束>


<原文开始>
// Connection pool, which is not a real connection pool, but a connection reusable pool.
<原文结束>

# <翻译开始>
// 连接池，其实不是一个真正的连接池，而是一个连接的重用池。. md5:a26d386f822f05df
# <翻译结束>


<原文开始>
// Status of current connection, which is used to mark this connection usable or not.
<原文结束>

# <翻译开始>
// 当前连接的状态，用于标记此连接是否可用。. md5:12cd6637c98f2ef8
# <翻译结束>


<原文开始>
// Default TTL for connection in the pool.
<原文结束>

# <翻译开始>
// 连接池中连接的默认TTL（生存时间）。. md5:0bf58836ef7b5d32
# <翻译结束>


<原文开始>
// Means it is now connective.
<原文结束>

# <翻译开始>
// 意味着它现在是连接状态。. md5:fc77bf21979d8581
# <翻译结束>


<原文开始>
// Means it should be closed and removed from pool.
<原文结束>

# <翻译开始>
// 表示应该关闭并从池中移除。. md5:0a8b237e0bb2ab9a
# <翻译结束>


<原文开始>
// addressPoolMap is a mapping for address to its pool object.
<原文结束>

# <翻译开始>
// addressPoolMap 是一个将地址映射到其池对象的映射。. md5:8e4ae3e1f1fdc0a6
# <翻译结束>


<原文开始>
// NewPoolConn creates and returns a connection with pool feature.
<原文结束>

# <翻译开始>
// NewPoolConn 创建并返回一个具有连接池特性的连接。. md5:ee2281aa2be8c181
# <翻译结束>


<原文开始>
// Close puts back the connection to the pool if it's active,
// or closes the connection if it's not active.
//
// Note that, if `c` calls Close function closing itself, `c` can not
// be used again.
<原文结束>

# <翻译开始>
// Close将活动的连接放回池中，如果连接未活动，则将其关闭。
//
// 请注意，如果`c`调用Close函数关闭自身，那么`c`将无法再次使用。
// md5:8596872730e65b10
# <翻译结束>


<原文开始>
// Send writes data to the connection. It retrieves a new connection from its pool if it fails
// writing data.
<原文结束>

# <翻译开始>
// Send 将数据写入连接。如果写数据失败，它会从其池中获取一个新的连接。
// md5:a5cfc10ec76d87e1
# <翻译结束>


<原文开始>
// Recv receives data from the connection.
<原文结束>

# <翻译开始>
// Recv 从连接中接收数据。. md5:d32a0574f5be517a
# <翻译结束>


<原文开始>
// RecvLine reads data from the connection until reads char '\n'.
// Note that the returned result does not contain the last char '\n'.
<原文结束>

# <翻译开始>
// RecvLine 从连接中读取数据，直到读取到字符 '\n'。注意，返回的结果不包含最后一个字符 '\n'。
// md5:e8f4d38a9d0e03e2
# <翻译结束>


<原文开始>
// RecvTill reads data from the connection until reads bytes `til`.
// Note that the returned result contains the last bytes `til`.
<原文结束>

# <翻译开始>
// RecvTill 从连接中读取数据，直到读取到字节 `til` 为止。
// 注意，返回的结果中包含最后一个字节 `til`。
// md5:3d5a6b2420bd7164
# <翻译结束>


<原文开始>
// RecvWithTimeout reads data from the connection with timeout.
<原文结束>

# <翻译开始>
// RecvWithTimeout 在连接上读取数据，带有超时设置。. md5:9c30fddfcddde9a2
# <翻译结束>


<原文开始>
// SendWithTimeout writes data to the connection with timeout.
<原文结束>

# <翻译开始>
// SendWithTimeout 在超时时间内向连接写入数据。. md5:776c26aa00723dd4
# <翻译结束>


<原文开始>
// SendRecv writes data to the connection and blocks reading response.
<原文结束>

# <翻译开始>
// SendRecv 向连接写入数据，并阻塞读取响应。. md5:a92dbf9e10bfe35b
# <翻译结束>


<原文开始>
// SendRecvWithTimeout writes data to the connection and reads response with timeout.
<原文结束>

# <翻译开始>
// SendRecvWithTimeout 向连接写入数据并带有超时时间地读取响应。. md5:154815490fa55262
# <翻译结束>

