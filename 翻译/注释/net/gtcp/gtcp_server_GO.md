
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
// FreePortAddress marks the server listens using random free port.
<原文结束>

# <翻译开始>
// FreePortAddress 表示服务器使用随机空闲端口进行监听。
# <翻译结束>







<原文开始>
// Used for Server.listen concurrent safety. -- The golang test with data race checks this.
<原文结束>

# <翻译开始>
// 用于Server.listen方法的并发安全。-- 此段golang代码会通过数据竞争检测进行测试。
# <翻译结束>






















<原文开始>
// Map for name to server, for singleton purpose.
<原文结束>

# <翻译开始>
// 用于单例目的，映射名称到服务器的Map。
# <翻译结束>


<原文开始>
// GetServer returns the TCP server with specified `name`,
// or it returns a new normal TCP server named `name` if it does not exist.
// The parameter `name` is used to specify the TCP server
<原文结束>

# <翻译开始>
// GetServer 函数返回指定名称 `name` 的 TCP 服务器，
// 如果该服务器不存在，则返回一个新的普通 TCP 服务器并命名为 `name`。
// 参数 `name` 用于指定要获取的 TCP 服务器。
# <翻译结束>


<原文开始>
// NewServer creates and returns a new normal TCP server.
// The parameter `name` is optional, which is used to specify the instance name of the server.
<原文结束>

# <翻译开始>
// NewServer 创建并返回一个新的普通TCP服务器。
// 参数`name`是可选的，用于指定服务器实例的名称。
# <翻译结束>


<原文开始>
// NewServerTLS creates and returns a new TCP server with TLS support.
// The parameter `name` is optional, which is used to specify the instance name of the server.
<原文结束>

# <翻译开始>
// NewServerTLS 创建并返回一个带有 TLS 支持的新 TCP 服务器。
// 参数 `name` 是可选的，用于指定服务器实例的名称。
# <翻译结束>


<原文开始>
// NewServerKeyCrt creates and returns a new TCP server with TLS support.
// The parameter `name` is optional, which is used to specify the instance name of the server.
<原文结束>

# <翻译开始>
// NewServerKeyCrt 创建并返回一个带有 TLS 支持的新 TCP 服务器。
// 参数 `name` 是可选的，用于指定服务器实例名称。
# <翻译结束>


<原文开始>
// SetAddress sets the listening address for server.
<原文结束>

# <翻译开始>
// SetAddress 设置服务器的监听地址。
# <翻译结束>


<原文开始>
// GetAddress get the listening address for server.
<原文结束>

# <翻译开始>
// GetAddress 获取服务器的监听地址。
# <翻译结束>


<原文开始>
// SetHandler sets the connection handler for server.
<原文结束>

# <翻译开始>
// SetHandler 设置服务器的连接处理器。
# <翻译结束>


<原文开始>
// SetTLSKeyCrt sets the certificate and key file for TLS configuration of server.
<原文结束>

# <翻译开始>
// SetTLSKeyCrt 用于设置服务器TLS配置所需的证书和密钥文件。
# <翻译结束>


<原文开始>
// SetTLSConfig sets the TLS configuration of server.
<原文结束>

# <翻译开始>
// SetTLSConfig 设置服务器的 TLS 配置。
# <翻译结束>


<原文开始>
// Close closes the listener and shutdowns the server.
<原文结束>

# <翻译开始>
// Close 关闭监听器并关闭服务器。
# <翻译结束>


<原文开始>
// Run starts running the TCP Server.
<原文结束>

# <翻译开始>
// Run 开始运行 TCP 服务器。
# <翻译结束>


<原文开始>
// GetListenedAddress retrieves and returns the address string which are listened by current server.
<原文结束>

# <翻译开始>
// GetListenedAddress 获取并返回当前服务器监听的地址字符串。
# <翻译结束>


<原文开始>
// GetListenedPort retrieves and returns one port which is listened to by current server.
<原文结束>

# <翻译开始>
// GetListenedPort 获取并返回当前服务器正在监听的一个端口。
# <翻译结束>







<原文开始>
// Server is a TCP server.
<原文结束>

# <翻译开始>
// Server 是一个 TCP 服务器。
# <翻译结束>


<原文开始>
// TCP address listener.
<原文结束>

# <翻译开始>
// TCP地址监听器。
# <翻译结束>


<原文开始>
// Server listening address.
<原文结束>

# <翻译开始>
// 服务监听地址。
# <翻译结束>


<原文开始>
// Connection handler.
<原文结束>

# <翻译开始>
// 连接处理器
# <翻译结束>


<原文开始>
// TLS configuration.
<原文结束>

# <翻译开始>
// TLS 配置
# <翻译结束>

