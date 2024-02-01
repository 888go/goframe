
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
// Adapter is an interface for universal redis operations.
<原文结束>

# <翻译开始>
// Adapter 是一个用于通用 Redis 操作的接口。
# <翻译结束>


<原文开始>
	// Do send a command to the server and returns the received reply.
	// It uses json.Marshal for struct/slice/map type values before committing them to redis.
<原文结束>

# <翻译开始>
// 向服务器发送命令并返回接收到的回复。
// 在将结构体、切片或映射类型值提交到redis前，它使用json.Marshal进行序列化。
# <翻译结束>


<原文开始>
	// Conn retrieves and returns a connection object for continuous operations.
	// Note that you should call Close function manually if you do not use this connection any further.
<原文结束>

# <翻译开始>
// Conn 获取并返回一个用于连续操作的连接对象。
// 注意，如果你不再使用此连接，应手动调用 Close 函数。
# <翻译结束>


<原文开始>
// Close closes current redis client, closes its connection pool and releases all its related resources.
<原文结束>

# <翻译开始>
// Close 关闭当前的 Redis 客户端，关闭其连接池并释放所有相关的资源。
# <翻译结束>


<原文开始>
// Conn is an interface of a connection from universal redis client.
<原文结束>

# <翻译开始>
// Conn 是一个通用 Redis 客户端连接的接口。
# <翻译结束>


<原文开始>
// Close puts the connection back to connection pool.
<原文结束>

# <翻译开始>
// Close将连接放回连接池。
# <翻译结束>


<原文开始>
// AdapterGroup is an interface managing group operations for redis.
<原文结束>

# <翻译开始>
// AdapterGroup 是一个接口，用于管理针对 Redis 的组操作。
# <翻译结束>


<原文开始>
// ConnCommand is an interface managing some operations bound to certain connection.
<原文结束>

# <翻译开始>
// ConnCommand 是一个接口，用于管理与特定连接相关的一些操作。
# <翻译结束>


<原文开始>
	// Subscribe subscribes the client to the specified channels.
	// https://redis.io/commands/subscribe/
<原文结束>

# <翻译开始>
// Subscribe 订阅函数，使客户端订阅指定的频道。
// 参考文档：https://redis.io/commands/subscribe/
# <翻译结束>


<原文开始>
	// PSubscribe subscribes the client to the given patterns.
	//
	// Supported glob-style patterns:
	// - h?llo subscribes to hello, hallo and hxllo
	// - h*llo subscribes to hllo and heeeello
	// - h[ae]llo subscribes to hello and hallo, but not hillo
	//
	// Use \ to escape special characters if you want to match them verbatim.
	//
	// https://redis.io/commands/psubscribe/
<原文结束>

# <翻译开始>
// PSubscribe 订阅客户端到给定的模式。
//
// 支持的glob风格模式：
// - h?llo 订阅hello, hallo和hxllo
// - h*llo 订阅hllo和heeeello
// - h[ae]llo 订阅hello和hallo，但不订阅hillo
//
// 如果你想精确匹配特殊字符，请使用\进行转义。
//
// 参考文档：https://redis.io/commands/psubscribe/
# <翻译结束>


<原文开始>
// ReceiveMessage receives a single message of subscription from the Redis server.
<原文结束>

# <翻译开始>
// ReceiveMessage 从 Redis 服务器接收订阅的单条消息。
# <翻译结束>


<原文开始>
// Receive receives a single reply as gvar.Var from the Redis server.
<原文结束>

# <翻译开始>
// Receive 从Redis服务器接收单个回复作为gvar.Var。
# <翻译结束>

