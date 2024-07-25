
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Conn is the TCP connection object.
<原文结束>

# <翻译开始>
// Conn是TCP连接对象。 md5:d539be0916b2b1b8
# <翻译结束>


<原文开始>
// Underlying TCP connection object.
<原文结束>

# <翻译开始>
// 基础TCP连接对象。 md5:b7b1c0f618f0d5c9
# <翻译结束>


<原文开始>
// Buffer reader for connection.
<原文结束>

# <翻译开始>
// 连接的缓冲读取器。 md5:490834adc66f0794
# <翻译结束>


<原文开始>
// Timeout point for reading.
<原文结束>

# <翻译开始>
// 读取的超时点。 md5:c3a775342f7a3b1e
# <翻译结束>


<原文开始>
// Timeout point for writing.
<原文结束>

# <翻译开始>
// 写入的超时点。 md5:27ca5c03a339a72e
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
// NewConn creates and returns a new connection with given address.
<原文结束>

# <翻译开始>
// NewConn 创建并返回一个与给定地址的新连接。 md5:b8d7c0b5ae5f53f0
# <翻译结束>


<原文开始>
// NewConnTLS creates and returns a new TLS connection
// with given address and TLS configuration.
<原文结束>

# <翻译开始>
// NewConnTLS 创建并返回一个新的TLS连接，使用给定的地址和TLS配置。 md5:a21dcb1cad67caa6
# <翻译结束>


<原文开始>
// NewConnKeyCrt creates and returns a new TLS connection
// with given address and TLS certificate and key files.
<原文结束>

# <翻译开始>
// NewConnKeyCrt 创建并返回一个新的带有给定地址和TLS证书及密钥文件的TLS连接。 md5:b79c43de9a5f13ce
# <翻译结束>


<原文开始>
// NewConnByNetConn creates and returns a TCP connection object with given net.Conn object.
<原文结束>

# <翻译开始>
// NewConnByNetConn 根据给定的net.Conn对象创建并返回一个TCP连接对象。 md5:88aad787fa32f138
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
// 注意，
// 1. 如果长度为0，表示它会从当前缓冲区接收数据并立即返回。
// 2. 如果长度小于0，表示它会接收连接中的所有数据，并返回直到没有更多的数据。如果决定接收缓冲区中的所有数据，开发者需要注意自行解析数据包。
// 3. 如果长度大于0，表示它会阻塞，直到接收到指定长度的数据。这是最常用的用于接收数据的长度值。 md5:75d42f229725a3f7
# <翻译结束>


<原文开始>
// Whether buffer reading timeout set.
<原文结束>

# <翻译开始>
// 是否设置了缓冲区读取超时。 md5:7a71357a4bf2a1e8
# <翻译结束>


<原文开始>
// It reads til `length` size if `length` is specified.
<原文结束>

# <翻译开始>
				// 如果指定了`length`，则读取`length`大小的内容。 md5:1216f8afe5006977
# <翻译结束>


<原文开始>
// If it exceeds the buffer size, it then automatically increases its buffer size.
<原文结束>

# <翻译开始>
					// 如果它超过了缓冲区的大小，那么它会自动增加其缓冲区的大小。 md5:2dfec7c7ce3557ab
# <翻译结束>


<原文开始>
// It returns immediately if received size is lesser than buffer size.
<原文结束>

# <翻译开始>
					// 如果接收到的大小小于缓冲区大小，它会立即返回。 md5:e98695cece9485a3
# <翻译结束>


<原文开始>
// Re-set the timeout when reading data.
<原文结束>

# <翻译开始>
			// 重新设置读取数据时的超时时间。 md5:a04ae0a806c5a3c6
# <翻译结束>


<原文开始>
// It fails even it retried.
<原文结束>

# <翻译开始>
				// 即使重试也会失败。 md5:7f32623c1f255555
# <翻译结束>


<原文开始>
// Just read once from buffer.
<原文结束>

# <翻译开始>
		// 从缓冲区只读一次。 md5:26f6ad162c136702
# <翻译结束>


<原文开始>
// RecvLine reads data from the connection until reads char '\n'.
// Note that the returned result does not contain the last char '\n'.
<原文结束>

# <翻译开始>
// RecvLine 从连接中读取数据，直到读取到字符 '\n'。注意，返回的结果不包含最后一个字符 '\n'。 md5:e8f4d38a9d0e03e2
# <翻译结束>


<原文开始>
// RecvTill reads data from the connection until reads bytes `til`.
// Note that the returned result contains the last bytes `til`.
<原文结束>

# <翻译开始>
// RecvTill 从连接中读取数据，直到读取到字节 `til` 为止。
// 注意，返回的结果中包含最后一个字节 `til`。 md5:3d5a6b2420bd7164
# <翻译结束>


<原文开始>
// RecvWithTimeout reads data from the connection with timeout.
<原文结束>

# <翻译开始>
// RecvWithTimeout 在连接上读取数据，带有超时设置。 md5:9c30fddfcddde9a2
# <翻译结束>


<原文开始>
// SendWithTimeout writes data to the connection with timeout.
<原文结束>

# <翻译开始>
// SendWithTimeout 在超时时间内向连接写入数据。 md5:776c26aa00723dd4
# <翻译结束>


<原文开始>
// SendRecv writes data to the connection and blocks reading response.
<原文结束>

# <翻译开始>
// SendRecv 向连接写入数据，并阻塞读取响应。 md5:a92dbf9e10bfe35b
# <翻译结束>


<原文开始>
// SendRecvWithTimeout writes data to the connection and reads response with timeout.
<原文结束>

# <翻译开始>
// SendRecvWithTimeout 向连接写入数据并带有超时时间地读取响应。 md5:154815490fa55262
# <翻译结束>


<原文开始>
// SetDeadline sets the deadline for current connection.
<原文结束>

# <翻译开始>
// SetDeadline 设置当前连接的截止时间。 md5:d3ab57d58a8a6b64
# <翻译结束>


<原文开始>
// SetDeadlineRecv sets the deadline of receiving for current connection.
<原文结束>

# <翻译开始>
// SetDeadlineRecv为当前连接设置接收截止期限。 md5:5cf236fbc63fc1f7
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
// 等待时间不能过长，否则可能会延迟从远程地址接收数据。 md5:54992dd21ce2360a
# <翻译结束>

