
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
// PoolConn is a connection with pool feature for TCP.
// Note that it is NOT a pool or connection manager, it is just a TCP connection object.
<原文结束>

# <翻译开始>
// PoolConn 是具有连接池特性的 TCP 连接。
// 注意，它不是一个连接池或连接管理器，它仅仅是一个 TCP 连接对象。
# <翻译结束>


<原文开始>
// Underlying connection object.
<原文结束>

# <翻译开始>
// 基础连接对象。
# <翻译结束>


<原文开始>
// Connection pool, which is not a real connection pool, but a connection reusable pool.
<原文结束>

# <翻译开始>
// 连接池，虽然并非真正的连接池，但实现了连接复用的功能。
# <翻译结束>


<原文开始>
// Status of current connection, which is used to mark this connection usable or not.
<原文结束>

# <翻译开始>
// 当前连接的状态，用于标记此连接是否可用。
# <翻译结束>







<原文开始>
// Means it is now connective.
<原文结束>

# <翻译开始>
// 表示它现在已经连接。
# <翻译结束>


<原文开始>
// Means it should be closed and removed from pool.
<原文结束>

# <翻译开始>
// 这意味着它应该被关闭并从池中移除。
# <翻译结束>


<原文开始>
// addressPoolMap is a mapping for address to its pool object.
<原文结束>

# <翻译开始>
// addressPoolMap 是一个映射，用于将地址与其对应的池对象关联。
# <翻译结束>


<原文开始>
// NewPoolConn creates and returns a connection with pool feature.
<原文结束>

# <翻译开始>
// NewPoolConn 创建并返回一个具有连接池功能的连接。
# <翻译结束>


<原文开始>
// Close puts back the connection to the pool if it's active,
// or closes the connection if it's not active.
//
// Note that, if `c` calls Close function closing itself, `c` can not
// be used again.
<原文结束>

# <翻译开始>
// Close函数会将活跃的连接归还给连接池，如果该连接处于非活跃状态，则关闭该连接。
//
// 注意，如果通过`c`调用Close函数来关闭自身，则此后不能再使用`c`。
# <翻译结束>


<原文开始>
// Send writes data to the connection. It retrieves a new connection from its pool if it fails
// writing data.
<原文结束>

# <翻译开始>
// Send 将数据写入连接。如果写入数据失败，它将从其连接池中获取一个新的连接。
# <翻译结束>


<原文开始>
// Recv receives data from the connection.
<原文结束>

# <翻译开始>
// Recv 从连接中接收数据。
# <翻译结束>


<原文开始>
// RecvLine reads data from the connection until reads char '\n'.
// Note that the returned result does not contain the last char '\n'.
<原文结束>

# <翻译开始>
// RecvLine 从连接中读取数据，直到读取到字符 '\n'。
// 注意，返回的结果不包含最后的字符 '\n'。
# <翻译结束>


<原文开始>
// RecvTill reads data from the connection until reads bytes `til`.
// Note that the returned result contains the last bytes `til`.
<原文结束>

# <翻译开始>
// RecvTill从连接中读取数据，直到读取到指定字节序列`til`为止。
// 注意，返回的结果中包含最后的字节序列`til`。
# <翻译结束>


<原文开始>
// RecvWithTimeout reads data from the connection with timeout.
<原文结束>

# <翻译开始>
// RecvWithTimeout函数以超时方式从连接中读取数据。
# <翻译结束>


<原文开始>
// SendWithTimeout writes data to the connection with timeout.
<原文结束>

# <翻译开始>
// SendWithTimeout 函数在设定的超时时间内向连接写入数据。
# <翻译结束>


<原文开始>
// SendRecv writes data to the connection and blocks reading response.
<原文结束>

# <翻译开始>
// SendRecv 向连接写入数据，并阻塞等待读取响应。
# <翻译结束>


<原文开始>
// SendRecvWithTimeout writes data to the connection and reads response with timeout.
<原文结束>

# <翻译开始>
// SendRecvWithTimeout 函数向连接写入数据并设定超时读取响应。
# <翻译结束>


<原文开始>
// Default TTL for connection in the pool.
<原文结束>

# <翻译开始>
// Default TTL for connection in the pool.
// 连接池中连接的默认生存时间（TTL）。
# <翻译结束>

