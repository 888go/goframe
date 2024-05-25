
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
// Conn handles the UDP connection.
<原文结束>

# <翻译开始>
// Conn 处理 UDP 连接。 md5:3d72ff914b3663e1
# <翻译结束>


<原文开始>
// Underlying UDP connection.
<原文结束>

# <翻译开始>
// 底层UDP连接。 md5:a4de01bc082c3b97
# <翻译结束>


<原文开始>
// Timeout point for reading data.
<原文结束>

# <翻译开始>
// 读取数据的超时点。 md5:115720ebb2ad57ae
# <翻译结束>


<原文开始>
// Timeout point for writing data.
<原文结束>

# <翻译开始>
// 写入数据的超时点。 md5:950ee5a16a0f64fc
# <翻译结束>


<原文开始>
// Interval duration for reading buffer.
<原文结束>

# <翻译开始>
// 读取缓冲区的间隔持续时间。 md5:373de61f6fb41d4d
# <翻译结束>


<原文开始>
// Default interval for reading buffer.
<原文结束>

# <翻译开始>
// 读取缓冲区的默认间隔时间。 md5:f0502056d8bc6f66
# <翻译结束>


<原文开始>
// NewConn creates UDP connection to `remoteAddress`.
// The optional parameter `localAddress` specifies the local address for connection.
<原文结束>

# <翻译开始>
// NewConn 创建到 `remoteAddress` 的 UDP 连接。
// 可选参数 `localAddress` 指定连接的本地地址。
// md5:d5e06df7ea2ee28d
# <翻译结束>


<原文开始>
// NewConnByNetConn creates an UDP connection object with given *net.UDPConn object.
<原文结束>

# <翻译开始>
// NewConnByNetConn 使用给定的 *net.UDPConn 对象创建一个UDP连接对象。 md5:8cbe128848b49656
# <翻译结束>


<原文开始>
// Send writes data to remote address.
<原文结束>

# <翻译开始>
// Send 将数据写入远程地址。 md5:445009019bd4a1a8
# <翻译结束>


<原文开始>
// Still failed even after retrying.
<原文结束>

# <翻译开始>
// 重试后仍然失败。 md5:b819d69935ab7496
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
// 参数 `buffer` 用于自定义接收缓冲区大小。如果 `buffer` <= 0，将使用默认的缓冲区大小，即1024字节。
//
// UDP协议存在分包边界，如果指定的缓冲区大小足够大，我们可以接收到一个完整的数据包。非常重要的是，必须一次性接收完整个包，否则剩下的包数据将会丢失。
// md5:190b81cc02f9d449
# <翻译结束>


<原文开始>
// Current remote address for reading
<原文结束>

# <翻译开始>
// 当前用于读取的远程地址. md5:7bff72f1a6b8b788
# <翻译结束>


<原文开始>
// It fails even it retried.
<原文结束>

# <翻译开始>
// 即使重试也会失败。 md5:7f32623c1f255555
# <翻译结束>


<原文开始>
// SendRecv writes data to connection and blocks reading response.
<原文结束>

# <翻译开始>
// SendRecv 向连接写入数据并阻塞读取响应。 md5:8416afe592163603
# <翻译结束>


<原文开始>
// RecvWithTimeout reads data from remote address with timeout.
<原文结束>

# <翻译开始>
// RecvWithTimeout 带超时限制地从远程地址读取数据。 md5:9e229854a65f6de2
# <翻译结束>


<原文开始>
// SendWithTimeout writes data to connection with timeout.
<原文结束>

# <翻译开始>
// SendWithTimeout 在连接中写入数据，并设置超时时间。 md5:d15d51d6004b2a6a
# <翻译结束>


<原文开始>
// SendRecvWithTimeout writes data to connection and reads response with timeout.
<原文结束>

# <翻译开始>
// SendRecvWithTimeout 向连接写入数据，并在超时时间内读取响应。 md5:6aa7751868598fb2
# <翻译结束>


<原文开始>
// SetDeadline sets the read and write deadlines associated with the connection.
<原文结束>

# <翻译开始>
// SetDeadline 设置与连接相关的读写超时截止时间。 md5:e0438bc956760556
# <翻译结束>


<原文开始>
// SetDeadlineRecv sets the read deadline associated with the connection.
<原文结束>

# <翻译开始>
// SetDeadlineRecv 设置与连接关联的读取截止时间。 md5:763094b16fe580fe
# <翻译结束>


<原文开始>
// SetDeadlineSend sets the deadline of sending for current connection.
<原文结束>

# <翻译开始>
// SetDeadlineSend 设置当前连接的发送截止期限。 md5:9f0d98d0e6beda95
# <翻译结束>


<原文开始>
// SetBufferWaitRecv sets the buffer waiting timeout when reading all data from connection.
// The waiting duration cannot be too long which might delay receiving data from remote address.
<原文结束>

# <翻译开始>
// SetBufferWaitRecv 设置从连接读取所有数据时的缓冲等待超时时间。
// 等待时间不能过长，否则可能会延迟从远程地址接收数据。
// md5:54992dd21ce2360a
# <翻译结束>


<原文开始>
// RemoteAddr returns the remote address of current UDP connection.
// Note that it cannot use c.conn.RemoteAddr() as it is nil.
<原文结束>

# <翻译开始>
// RemoteAddr 返回当前UDP连接的远程地址。
// 请注意，它不能使用c.conn.RemoteAddr()，因为该值为nil。
// md5:0a785ae4cb967a81
# <翻译结束>

