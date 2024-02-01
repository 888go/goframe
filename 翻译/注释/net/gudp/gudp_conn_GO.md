
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
// Conn handles the UDP connection.
<原文结束>

# <翻译开始>
// Conn 处理UDP连接。
# <翻译结束>


<原文开始>
// Timeout point for reading data.
<原文结束>

# <翻译开始>
// 读取数据的超时点。
# <翻译结束>


<原文开始>
// Timeout point for writing data.
<原文结束>

# <翻译开始>
// 数据写入的超时点。
# <翻译结束>


<原文开始>
// Interval duration for reading buffer.
<原文结束>

# <翻译开始>
// 读取缓冲区的间隔时长。
# <翻译结束>


<原文开始>
// Default interval for reading buffer.
<原文结束>

# <翻译开始>
// 默认读取缓冲区的间隔时间。
# <翻译结束>


<原文开始>
// NewConn creates UDP connection to `remoteAddress`.
// The optional parameter `localAddress` specifies the local address for connection.
<原文结束>

# <翻译开始>
// NewConn创建一个到`remoteAddress`的UDP连接。
// 可选参数`localAddress`用于指定连接的本地地址。
# <翻译结束>


<原文开始>
// NewConnByNetConn creates an UDP connection object with given *net.UDPConn object.
<原文结束>

# <翻译开始>
// NewConnByNetConn 通过给定的 *net.UDPConn 对象创建一个 UDP 连接对象。
# <翻译结束>


<原文开始>
// Send writes data to remote address.
<原文结束>

# <翻译开始>
// Send 将数据写入远程地址。
# <翻译结束>


<原文开始>
// Still failed even after retrying.
<原文结束>

# <翻译开始>
// 即便重试后仍然失败。
# <翻译结束>


<原文开始>
// Recv receives and returns data from remote address.
// The parameter `buffer` is used for customizing the receiving buffer size. If `buffer` <= 0,
// it uses the default buffer size, which is 1024 byte.
//
// There's package border in UDP protocol, we can receive a complete package if specified
// buffer size is big enough. VERY NOTE that we should receive the complete package in once
// or else the leftover package data would be dropped.
<原文结束>

# <翻译开始>
// Recv 从远程地址接收并返回数据。
// 参数`buffer`用于自定义接收缓冲区大小。如果`buffer` <= 0，
// 则使用默认的缓冲区大小，即1024字节。
//
// 在UDP协议中有包边界的限制，如果指定的缓冲区大小足够大，我们可以一次性接收一个完整的包。
// 非常注意的是，我们应该一次性接收完整个包，否则剩余的包数据会被丢弃。
# <翻译结束>


<原文开始>
// Current remote address for reading
<原文结束>

# <翻译开始>
// 当前用于读取的远程地址
# <翻译结束>


<原文开始>
// SendRecv writes data to connection and blocks reading response.
<原文结束>

# <翻译开始>
// SendRecv 向连接写入数据，并阻塞等待读取响应。
# <翻译结束>


<原文开始>
// RecvWithTimeout reads data from remote address with timeout.
<原文结束>

# <翻译开始>
// RecvWithTimeout在指定超时时间内从远程地址读取数据。
# <翻译结束>


<原文开始>
// SendWithTimeout writes data to connection with timeout.
<原文结束>

# <翻译开始>
// SendWithTimeout 在设定的超时时间内向连接写入数据。
# <翻译结束>


<原文开始>
// SendRecvWithTimeout writes data to connection and reads response with timeout.
<原文结束>

# <翻译开始>
// SendRecvWithTimeout 将数据写入连接并在超时时间内读取响应。
# <翻译结束>


<原文开始>
// SetDeadline sets the read and write deadlines associated with the connection.
<原文结束>

# <翻译开始>
// SetDeadline为连接设置读取和写入的截止时间。
# <翻译结束>


<原文开始>
// SetDeadlineRecv sets the read deadline associated with the connection.
<原文结束>

# <翻译开始>
// SetDeadlineRecv 设置与连接关联的读取截止时间。
# <翻译结束>


<原文开始>
// SetDeadlineSend sets the deadline of sending for current connection.
<原文结束>

# <翻译开始>
// SetDeadlineSend 设置当前连接的发送截止时间。
# <翻译结束>


<原文开始>
// SetBufferWaitRecv sets the buffer waiting timeout when reading all data from connection.
// The waiting duration cannot be too long which might delay receiving data from remote address.
<原文结束>

# <翻译开始>
// SetBufferWaitRecv 设置从连接读取所有数据时的缓冲等待超时时间。
// 等待时长不能过长，否则可能延迟从远程地址接收数据。
# <翻译结束>


<原文开始>
// RemoteAddr returns the remote address of current UDP connection.
// Note that it cannot use c.conn.RemoteAddr() as it is nil.
<原文结束>

# <翻译开始>
// RemoteAddr 返回当前 UDP 连接的远程地址。
// 注意，由于 c.conn.RemoteAddr() 为 nil，所以不能使用它。
# <翻译结束>


<原文开始>
// Underlying UDP connection.
<原文结束>

# <翻译开始>
// 基础的UDP连接。
# <翻译结束>


<原文开始>
// (Byte)Buffer size.
<原文结束>

# <翻译开始>
// (字节)缓冲区大小。
# <翻译结束>


<原文开始>
// Connection closed.
<原文结束>

# <翻译开始>
// 连接已关闭。
# <翻译结束>


<原文开始>
// It fails even it retried.
<原文结束>

# <翻译开始>
// 即使重试了也会失败。
# <翻译结束>

