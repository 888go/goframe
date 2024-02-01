
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
// Default connection timeout.
<原文结束>

# <翻译开始>
// 默认连接超时时间
# <翻译结束>







<原文开始>
// (Byte) Buffer size for reading.
<原文结束>

# <翻译开始>
// (字节) 读取时的缓冲区大小。
# <翻译结束>


<原文开始>
// NewNetConn creates and returns a net.Conn with given address like "127.0.0.1:80".
// The optional parameter `timeout` specifies the timeout for dialing connection.
<原文结束>

# <翻译开始>
// NewNetConn 创建并返回一个 net.Conn，其地址格式如 "127.0.0.1:80"。
// 可选参数 `timeout` 指定了拨号连接的超时时间。
# <翻译结束>


<原文开始>
// NewNetConnTLS creates and returns a TLS net.Conn with given address like "127.0.0.1:80".
// The optional parameter `timeout` specifies the timeout for dialing connection.
<原文结束>

# <翻译开始>
// NewNetConnTLS 根据给定的地址（如 "127.0.0.1:80"）创建并返回一个 TLS 安全连接 net.Conn。
// 可选参数 `timeout` 指定了建立连接时的超时时间。
# <翻译结束>


<原文开始>
// NewNetConnKeyCrt creates and returns a TLS net.Conn with given TLS certificate and key files
// and address like "127.0.0.1:80". The optional parameter `timeout` specifies the timeout for
// dialing connection.
<原文结束>

# <翻译开始>
// NewNetConnKeyCrt 根据给定的 TLS 证书和密钥文件以及类似 "127.0.0.1:80" 的地址创建并返回一个 TLS 网络连接（net.Conn）。可选参数 `timeout` 指定了建立连接时的超时时间。
# <翻译结束>


<原文开始>
// Send creates connection to `address`, writes `data` to the connection and then closes the connection.
// The optional parameter `retry` specifies the retry policy when fails in writing data.
<原文结束>

# <翻译开始>
// Send 函数创建到 `address` 的连接，将 `data` 数据写入该连接，然后关闭连接。
// 可选参数 `retry` 指定了在写入数据失败时的重试策略。
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
// SendRecv 创建到 `address` 的连接，将 `data` 写入连接，接收响应，然后关闭连接。
//
// 参数 `length` 指定等待接收的字节数量。如果 `length` 为 -1，则接收所有缓冲区内容并返回。
//
// 可选参数 `retry` 指定了在写入数据失败时重试策略。
# <翻译结束>


<原文开始>
// SendWithTimeout does Send logic with writing timeout limitation.
<原文结束>

# <翻译开始>
// SendWithTimeout 在写入超时限制下执行Send逻辑。
# <翻译结束>


<原文开始>
// SendRecvWithTimeout does SendRecv logic with reading timeout limitation.
<原文结束>

# <翻译开始>
// SendRecvWithTimeout 在读取超时限制下执行SendRecv逻辑。
# <翻译结束>


<原文开始>
// isTimeout checks whether given `err` is a timeout error.
<原文结束>

# <翻译开始>
// isTimeout 检查给定的 `err` 是否为超时错误。
# <翻译结束>


<原文开始>
// LoadKeyCrt creates and returns a TLS configuration object with given certificate and key files.
<原文结束>

# <翻译开始>
// LoadKeyCrt 通过给定的证书和密钥文件创建并返回一个 TLS 配置对象。
# <翻译结束>


<原文开始>
// MustGetFreePort performs as GetFreePort, but it panics is any error occurs.
<原文结束>

# <翻译开始>
// MustGetFreePort 的行为与 GetFreePort 相同，但是，如果发生任何错误，它会触发 panic。
# <翻译结束>


<原文开始>
// GetFreePort retrieves and returns a port that is free.
<原文结束>

# <翻译开始>
// GetFreePort 获取并返回一个可用的端口。
# <翻译结束>


<原文开始>
// GetFreePorts retrieves and returns specified number of ports that are free.
<原文结束>

# <翻译开始>
// GetFreePorts 获取并返回指定数量的空闲端口。
# <翻译结束>







<原文开始>
// Default retry interval.
<原文结束>

# <翻译开始>
// 默认重试间隔。
# <翻译结束>

