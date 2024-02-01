
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
// initGroup initializes the group object of redis.
<原文结束>

# <翻译开始>
// initGroup 初始化redis的组对象
# <翻译结束>


<原文开始>
// SetAdapter changes the underlying adapter with custom adapter for current redis client.
<原文结束>

# <翻译开始>
// SetAdapter 更改当前Redis客户端的底层适配器，使用自定义适配器。
# <翻译结束>


<原文开始>
// GetAdapter returns the adapter that is set in current redis client.
<原文结束>

# <翻译开始>
// GetAdapter 返回当前Redis客户端中设置的适配器。
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
// Do send a command to the server and returns the received reply.
// It uses json.Marshal for struct/slice/map type values before committing them to redis.
<原文结束>

# <翻译开始>
// 向服务器发送命令并返回接收到的回复。
// 在将结构体/切片/映射类型值提交给redis之前，它使用json.Marshal进行序列化。
# <翻译结束>


<原文开始>
// MustConn performs as function Conn, but it panics if any error occurs internally.
<原文结束>

# <翻译开始>
// MustConn 的行为与函数 Conn 相同，但如果内部发生任何错误，它会触发 panic。
# <翻译结束>


<原文开始>
// MustDo performs as function Do, but it panics if any error occurs internally.
<原文结束>

# <翻译开始>
// MustDo 执行与函数 Do 相同的操作，但如果内部出现任何错误，它会触发 panic（异常）。
# <翻译结束>


<原文开始>
// Close closes current redis client, closes its connection pool and releases all its related resources.
<原文结束>

# <翻译开始>
// Close 关闭当前的 Redis 客户端，关闭其连接池并释放所有相关资源。
# <翻译结束>

