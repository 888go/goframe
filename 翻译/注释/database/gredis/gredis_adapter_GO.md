
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
// Adapter is an interface for universal redis operations.
<原文结束>

# <翻译开始>
// Adapter是通用Redis操作的接口。 md5:9c96b73f93ac5323
# <翻译结束>


<原文开始>
// AdapterGroup is an interface managing group operations for redis.
<原文结束>

# <翻译开始>
// AdapterGroup 是一个接口，用于管理 Redis 的组操作。 md5:f603a1b02c295995
# <翻译结束>


<原文开始>
// AdapterOperation is the core operation functions for redis.
// These functions can be easily overwritten by custom implements.
<原文结束>

# <翻译开始>
// AdapterOperation 是 Redis 的核心操作函数。
// 这些函数可以被自定义实现轻松覆盖。
// md5:6a3c39d3c764e39e
# <翻译结束>


<原文开始>
	// Do send a command to the server and returns the received reply.
	// It uses json.Marshal for struct/slice/map type values before committing them to redis.
<原文结束>

# <翻译开始>
	// 发送一个命令到服务器并返回接收到的回复。
	// 在将结构体/切片/映射类型的值提交到redis之前，它使用json.Marshal进行编码。
	// md5:5a464ca35e177113
# <翻译结束>


<原文开始>
	// Conn retrieves and returns a connection object for continuous operations.
	// Note that you should call Close function manually if you do not use this connection any further.
<原文结束>

# <翻译开始>
	// Conn 获取并返回一个用于连续操作的连接对象。
	// 请注意，如果您不再使用此连接，请手动调用 Close 函数。
	// md5:adf083088afcd372
# <翻译结束>


<原文开始>
// Close closes current redis client, closes its connection pool and releases all its related resources.
<原文结束>

# <翻译开始>
// Close 方法关闭当前Redis客户端，关闭其连接池并释放所有相关资源。 md5:bfd91d0269572038
# <翻译结束>


<原文开始>
// Conn is an interface of a connection from universal redis client.
<原文结束>

# <翻译开始>
// Conn 是一个通用 Redis 客户端连接的接口。 md5:75bf8588ab4ad4e1
# <翻译结束>


<原文开始>
// Close puts the connection back to connection pool.
<原文结束>

# <翻译开始>
// Close 将连接放回连接池。 md5:7cc2158c987fb9c1
# <翻译结束>


<原文开始>
// ConnCommand is an interface managing some operations bound to certain connection.
<原文结束>

# <翻译开始>
// ConnCommand是一个接口，用于管理与特定连接绑定的一些操作。 md5:25fa514417ce2230
# <翻译结束>


<原文开始>
	// Subscribe subscribes the client to the specified channels.
	// https://redis.io/commands/subscribe/
<原文结束>

# <翻译开始>
	// Subscribe 将客户端订阅到指定的频道。
	// 参考链接：https:	//redis.io/commands/subscribe/
	// md5:a7414ed1d330bfc7
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
	// PSubscribe 将客户端订阅给定的模式。
	//
	// 支持的glob风格模式：
	// - h?llo 订阅hello, hallo和hxllo
	// - h*llo 订阅hllo和heeeello
	// - h[ae]llo 订阅hello和hallo，但不订阅hillo
	//
	// 如果需要匹配特殊字符本身，请使用\进行转义。
	//
	// https:	//redis.io/commands/psubscribe/
	// md5:0bfeb7ebd0d003a7
# <翻译结束>


<原文开始>
// ReceiveMessage receives a single message of subscription from the Redis server.
<原文结束>

# <翻译开始>
// ReceiveMessage 从Redis服务器接收一个订阅的消息。 md5:dbf6509713a7b2b3
# <翻译结束>


<原文开始>
// Receive receives a single reply as gvar.Var from the Redis server.
<原文结束>

# <翻译开始>
// Receive 从Redis服务器接收一个作为gvar.Var的单个回复。 md5:c4dad7138865cef4
# <翻译结束>

