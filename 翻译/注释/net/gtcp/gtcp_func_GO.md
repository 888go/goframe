
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
// Default connection timeout.
<原文结束>

# <翻译开始>
// 默认的连接超时时间。. md5:f32319a8522e8f90
# <翻译结束>


<原文开始>
// Default retry interval.
<原文结束>

# <翻译开始>
// 默认重试间隔。. md5:d53e6b260a9e594d
# <翻译结束>


<原文开始>
// (Byte) Buffer size for reading.
<原文结束>

# <翻译开始>
// 读取的字节缓冲区大小。. md5:3bb21d80469c9916
# <翻译结束>


<原文开始>
// NewNetConn creates and returns a net.Conn with given address like "127.0.0.1:80".
// The optional parameter `timeout` specifies the timeout for dialing connection.
<原文结束>

# <翻译开始>
// NewNetConn 创建并返回一个具有给定地址（如 "127.0.0.1:80"）的net.Conn。
// 可选参数 `timeout` 指定了建立连接的超时时间。
// md5:2e0124ac2d5ba04b
# <翻译结束>


<原文开始>
// NewNetConnTLS creates and returns a TLS net.Conn with given address like "127.0.0.1:80".
// The optional parameter `timeout` specifies the timeout for dialing connection.
<原文结束>

# <翻译开始>
// NewNetConnTLS 创建并返回一个具有给定地址（如 "127.0.0.1:80"）的 TLS net.Conn。
// 可选参数 `timeout` 指定了建立连接时的超时时间。
// md5:5eb25eb4d9f5078a
# <翻译结束>


<原文开始>
// NewNetConnKeyCrt creates and returns a TLS net.Conn with given TLS certificate and key files
// and address like "127.0.0.1:80". The optional parameter `timeout` specifies the timeout for
// dialing connection.
<原文结束>

# <翻译开始>
// NewNetConnKeyCrt 创建并返回一个带有给定TLS证书和密钥文件的TLS net.Conn，地址类似于"127.0.0.1:80"。可选参数`timeout`指定了连接超时时间。
// md5:232eecc025740731
# <翻译结束>


<原文开始>
// Send creates connection to `address`, writes `data` to the connection and then closes the connection.
// The optional parameter `retry` specifies the retry policy when fails in writing data.
<原文结束>

# <翻译开始>
// Send 建立连接到 `address`，向连接写入 `data`，然后关闭连接。
// 可选参数 `retry` 指定在写入数据失败时的重试策略。
// md5:657cbdf2b2958d6f
# <翻译结束>


<原文开始>
// SendRecv creates connection to `address`, writes `data` to the connection, receives response
// and then closes the connection.
//
// The parameter `length` specifies the bytes count waiting to receive. It receives all buffer content
// and returns if `length` is -1.
//
// The optional parameter `retry` specifies the retry policy when fails in writing data.
<原文结束>

# <翻译开始>
// SendRecv 会创建到 `address` 的连接，向该连接写入 `data`，接收响应，
// 然后关闭连接。
//
// 参数 `length` 指定等待接收的字节数量。如果 `length` 为 -1，则接收缓冲区的所有内容并返回。
//
// 可选参数 `retry` 指定了在写入数据失败时的重试策略。
// md5:2f0794c80f81d806
# <翻译结束>


<原文开始>
// SendWithTimeout does Send logic with writing timeout limitation.
<原文结束>

# <翻译开始>
// SendWithTimeout 在发送逻辑中添加了写入超时的限制。. md5:3ede704cb632bc5e
# <翻译结束>


<原文开始>
// SendRecvWithTimeout does SendRecv logic with reading timeout limitation.
<原文结束>

# <翻译开始>
// SendRecvWithTimeout 在限制读取超时的情况下执行SendRecv逻辑。. md5:a0b595ec27ab2abf
# <翻译结束>


<原文开始>
// isTimeout checks whether given `err` is a timeout error.
<原文结束>

# <翻译开始>
// isTimeout 检查给定的 `err` 是否是超时错误。. md5:c277cc8323b1a413
# <翻译结束>


<原文开始>
// LoadKeyCrt creates and returns a TLS configuration object with given certificate and key files.
<原文结束>

# <翻译开始>
// LoadKeyCrt 根据给定的证书和密钥文件创建并返回一个 TLS 配置对象。. md5:e31385756c06b0a4
# <翻译结束>


<原文开始>
// MustGetFreePort performs as GetFreePort, but it panics is any error occurs.
<原文结束>

# <翻译开始>
// MustGetFreePort 的行为与 GetFreePort 相同，但如果发生任何错误，它会直接 panic。. md5:a1ae43bc1faffe59
# <翻译结束>


<原文开始>
// GetFreePort retrieves and returns a port that is free.
<原文结束>

# <翻译开始>
// GetFreePort 获取并返回一个空闲的端口号。. md5:52dbf7a2d6e71da6
# <翻译结束>


<原文开始>
// GetFreePorts retrieves and returns specified number of ports that are free.
<原文结束>

# <翻译开始>
// GetFreePorts 获取并返回指定数量的空闲端口。. md5:ea99fb15b5bbc0fb
# <翻译结束>

