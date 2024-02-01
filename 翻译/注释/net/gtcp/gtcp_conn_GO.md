
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
// Conn is the TCP connection object.
<原文结束>

# <翻译开始>
// Conn 是 TCP 连接对象。
# <翻译结束>


<原文开始>
// Underlying TCP connection object.
<原文结束>

# <翻译开始>
// 底层TCP连接对象。
# <翻译结束>


<原文开始>
// Buffer reader for connection.
<原文结束>

# <翻译开始>
// 连接的缓冲读取器
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
// 默认读取缓冲区的间隔时间
# <翻译结束>


<原文开始>
// NewConn creates and returns a new connection with given address.
<原文结束>

# <翻译开始>
// NewConn 根据给定的地址创建并返回一个新的连接。
# <翻译结束>


<原文开始>
// NewConnTLS creates and returns a new TLS connection
// with given address and TLS configuration.
<原文结束>

# <翻译开始>
// NewConnTLS 根据给定的地址和 TLS 配置创建并返回一个新的 TLS 连接。
# <翻译结束>


<原文开始>
// NewConnKeyCrt creates and returns a new TLS connection
// with given address and TLS certificate and key files.
<原文结束>

# <翻译开始>
// NewConnKeyCrt 根据给定的地址和 TLS 证书及密钥文件创建并返回一个新的 TLS 连接。
# <翻译结束>


<原文开始>
// NewConnByNetConn creates and returns a TCP connection object with given net.Conn object.
<原文结束>

# <翻译开始>
// NewConnByNetConn 根据给定的 net.Conn 对象创建并返回一个 TCP 连接对象。
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
// 即使重试后仍然失败。
# <翻译结束>


<原文开始>
// Recv receives and returns data from the connection.
//
// Note that,
//  1. If length = 0, which means it receives the data from current buffer and returns immediately.
//  2. If length < 0, which means it receives all data from connection and returns it until no data
//     from connection. Developers should notice the package parsing yourself if you decide receiving
//     all data from buffer.
//  3. If length > 0, which means it blocks reading data from connection until length size was received.
//     It is the most commonly used length value for data receiving.
<原文结束>

# <翻译开始>
// Recv 从连接中接收并返回数据。
//
// 注意：
//  1. 如果 length 等于 0，表示它会从当前缓冲区接收数据并立即返回。
//  2. 如果 length 小于 0，表示它会从连接中接收所有数据并返回，直到没有更多的数据从连接中获取为止。开发者需要注意，
//     如果决定从缓冲区接收所有数据，则需要自行处理包解析问题。
//  3. 如果 length 大于 0，表示它会阻塞读取连接中的数据，直到接收到长度为 length 的数据为止。这是最常见的用于接收数据的长度值。
# <翻译结束>


<原文开始>
// Whether buffer reading timeout set.
<原文结束>

# <翻译开始>
// 是否设置了缓冲区读取超时。
# <翻译结束>


<原文开始>
// It reads til `length` size if `length` is specified.
<原文结束>

# <翻译开始>
// 如果指定了`length`，则读取到指定长度为止。
# <翻译结束>


<原文开始>
// If it exceeds the buffer size, it then automatically increases its buffer size.
<原文结束>

# <翻译开始>
// 如果超过缓冲区大小，它会自动增加其缓冲区大小。
# <翻译结束>


<原文开始>
// It returns immediately if received size is lesser than buffer size.
<原文结束>

# <翻译开始>
// 如果接收到的大小小于缓冲区大小，则它会立即返回。
# <翻译结束>


<原文开始>
// Re-set the timeout when reading data.
<原文结束>

# <翻译开始>
// 在读取数据时重新设置超时时间。
# <翻译结束>







<原文开始>
// Just read once from buffer.
<原文结束>

# <翻译开始>
// 仅从缓冲区读取一次。
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
// SetDeadline sets the deadline for current connection.
<原文结束>

# <翻译开始>
// SetDeadline 设置当前连接的截止时间。
# <翻译结束>


<原文开始>
// SetDeadlineRecv sets the deadline of receiving for current connection.
<原文结束>

# <翻译开始>
// SetDeadlineRecv为当前连接设置接收的截止时间。
# <翻译结束>


<原文开始>
// SetDeadlineSend sets the deadline of sending for current connection.
<原文结束>

# <翻译开始>
// SetDeadlineSend 为当前连接设置发送的截止时间。
# <翻译结束>


<原文开始>
// SetBufferWaitRecv sets the buffer waiting timeout when reading all data from connection.
// The waiting duration cannot be too long which might delay receiving data from remote address.
<原文结束>

# <翻译开始>
// SetBufferWaitRecv 设置从连接中读取所有数据时的缓冲等待超时时间。
// 等待时长不能过长，否则可能导致接收远程地址数据延迟。
# <翻译结束>


<原文开始>
// Timeout point for reading.
<原文结束>

# <翻译开始>
// 读取操作的超时点。
# <翻译结束>


<原文开始>
// Timeout point for writing.
<原文结束>

# <翻译开始>
// 写入操作的超时点。
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

