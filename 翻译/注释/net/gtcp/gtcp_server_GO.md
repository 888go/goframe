
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
// FreePortAddress marks the server listens using random free port.
<原文结束>

# <翻译开始>
	// FreePortAddress 标记服务器使用随机的空闲端口进行监听。 md5:16e8ca0633c4a135
# <翻译结束>


<原文开始>
// Server is a TCP server.
<原文结束>

# <翻译开始>
// Server 是一个 TCP 服务器。 md5:15abc757287261ea
# <翻译结束>


<原文开始>
// Used for Server.listen concurrent safety. -- The golang test with data race checks this.
<原文结束>

# <翻译开始>
// 用于 Server.listen 的并发安全。-- Go 语言的竞态条件检测会检查这个。 md5:d330fd21b35ec6b2
# <翻译结束>


<原文开始>
// Server listening address.
<原文结束>

# <翻译开始>
// 服务器监听地址。 md5:c8adda00f51a60d8
# <翻译结束>


<原文开始>
// Map for name to server, for singleton purpose.
<原文结束>

# <翻译开始>
// 用于单例目的，存储名称到服务器的映射。 md5:8e877c386766a97c
# <翻译结束>


<原文开始>
// GetServer returns the TCP server with specified `name`,
// or it returns a new normal TCP server named `name` if it does not exist.
// The parameter `name` is used to specify the TCP server
<原文结束>

# <翻译开始>
// GetServer 返回指定名称的 TCP 服务器，如果不存在，则返回一个新创建的默认名为 `name` 的 TCP 服务器。参数 `name` 用于指定 TCP 服务器的名称。
// md5:f6bb57410cf2ca98
# <翻译结束>


<原文开始>
// NewServer creates and returns a new normal TCP server.
// The parameter `name` is optional, which is used to specify the instance name of the server.
<原文结束>

# <翻译开始>
// NewServer 创建并返回一个新的普通TCP服务器。
// 参数 `name` 是可选的，用于指定服务器的实例名称。
// md5:ce4abdc7a25f75da
# <翻译结束>


<原文开始>
// NewServerTLS creates and returns a new TCP server with TLS support.
// The parameter `name` is optional, which is used to specify the instance name of the server.
<原文结束>

# <翻译开始>
// NewServerTLS 创建并返回一个支持TLS的TCP服务器。
// 参数`name`是可选的，用于指定服务器的实例名称。
// md5:102d98ca307029b3
# <翻译结束>


<原文开始>
// NewServerKeyCrt creates and returns a new TCP server with TLS support.
// The parameter `name` is optional, which is used to specify the instance name of the server.
<原文结束>

# <翻译开始>
// NewServerKeyCrt 创建并返回一个支持TLS的TCP服务器。
// 参数 `name` 是可选的，用于指定服务器的实例名称。
// md5:65a6856829628fe8
# <翻译结束>


<原文开始>
// SetAddress sets the listening address for server.
<原文结束>

# <翻译开始>
// SetAddress 设置服务器的监听地址。 md5:35306d25b7cbc244
# <翻译结束>


<原文开始>
// GetAddress get the listening address for server.
<原文结束>

# <翻译开始>
// GetAddress 获取服务器的监听地址。 md5:6085c2f0086d87f9
# <翻译结束>


<原文开始>
// SetHandler sets the connection handler for server.
<原文结束>

# <翻译开始>
// SetHandler 设置服务器的连接处理器。 md5:10bacdc88ff59cee
# <翻译结束>


<原文开始>
// SetTLSKeyCrt sets the certificate and key file for TLS configuration of server.
<原文结束>

# <翻译开始>
// SetTLSKeyCrt 设置服务器TLS配置的证书和密钥文件。 md5:dd19415f9056b27d
# <翻译结束>


<原文开始>
// SetTLSConfig sets the TLS configuration of server.
<原文结束>

# <翻译开始>
// SetTLSConfig 设置服务器的TLS配置。 md5:02f67dcfad23906c
# <翻译结束>


<原文开始>
// Close closes the listener and shutdowns the server.
<原文结束>

# <翻译开始>
// Close 方法关闭监听器并停止服务器。 md5:494fcac465675910
# <翻译结束>


<原文开始>
// Run starts running the TCP Server.
<原文结束>

# <翻译开始>
// Run 开始运行TCP服务器。 md5:b107bdcd45f1ccdc
# <翻译结束>


<原文开始>
// GetListenedAddress retrieves and returns the address string which are listened by current server.
<原文结束>

# <翻译开始>
// GetListenedAddress 获取并返回当前服务器所监听的地址字符串。 md5:51d352ffec9dc329
# <翻译结束>


<原文开始>
// GetListenedPort retrieves and returns one port which is listened to by current server.
<原文结束>

# <翻译开始>
// GetListenedPort 获取并返回当前服务器监听的其中一个端口。 md5:98e33a51d8d8309c
# <翻译结束>

