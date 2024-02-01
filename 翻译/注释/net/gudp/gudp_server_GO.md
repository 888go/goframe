
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
// FreePortAddress marks the server listens using random free port.
<原文结束>

# <翻译开始>
// FreePortAddress 表示服务器使用随机空闲端口进行监听。
# <翻译结束>


<原文开始>
// Used for Server.listen concurrent safety. -- The golang test with data race checks this.
<原文结束>

# <翻译开始>
// 用于保证Server.listen方法的并发安全性。-- 该Go语言测试会通过数据竞争检查进行验证。
# <翻译结束>


<原文开始>
// UDP server connection object.
<原文结束>

# <翻译开始>
// UDP服务器连接对象。
# <翻译结束>


<原文开始>
// UDP server listening address.
<原文结束>

# <翻译开始>
// UDP服务器监听地址。
# <翻译结束>


<原文开始>
// serverMapping is used for instance name to its UDP server mappings.
<原文结束>

# <翻译开始>
// serverMapping 用于实例名与其UDP服务器之间的映射。
# <翻译结束>


<原文开始>
// GetServer creates and returns an UDP server instance with given name.
<原文结束>

# <翻译开始>
// GetServer 根据给定名称创建并返回一个UDP服务器实例。
# <翻译结束>


<原文开始>
// NewServer creates and returns an UDP server.
// The optional parameter `name` is used to specify its name, which can be used for
// GetServer function to retrieve its instance.
<原文结束>

# <翻译开始>
// NewServer 创建并返回一个 UDP 服务器。
// 可选参数 `name` 用于指定其名称，可用于 GetServer 函数来获取其实例。
# <翻译结束>


<原文开始>
// SetAddress sets the server address for UDP server.
<原文结束>

# <翻译开始>
// SetAddress 设置 UDP 服务器的服务器地址。
# <翻译结束>


<原文开始>
// SetHandler sets the connection handler for UDP server.
<原文结束>

# <翻译开始>
// SetHandler 设置 UDP 服务器的连接处理器。
# <翻译结束>


<原文开始>
// Close closes the connection.
// It will make server shutdowns immediately.
<原文结束>

# <翻译开始>
// Close 关闭连接。
// 它将使服务器立即关闭。
# <翻译结束>


<原文开始>
// Run starts listening UDP connection.
<原文结束>

# <翻译开始>
// Run 开始监听 UDP 连接。
# <翻译结束>


<原文开始>
// GetListenedAddress retrieves and returns the address string which are listened by current server.
<原文结束>

# <翻译开始>
// GetListenedAddress 获取并返回当前服务器正在监听的地址字符串。
# <翻译结束>


<原文开始>
// GetListenedPort retrieves and returns one port which is listened to by current server.
<原文结束>

# <翻译开始>
// GetListenedPort 获取并返回当前服务器正在监听的一个端口。
# <翻译结束>


<原文开始>
// Server is the UDP server.
<原文结束>

# <翻译开始>
// Server 是 UDP 服务器。
# <翻译结束>


<原文开始>
// Handler for UDP connection.
<原文结束>

# <翻译开始>
// 处理UDP连接的处理器。
# <翻译结束>

