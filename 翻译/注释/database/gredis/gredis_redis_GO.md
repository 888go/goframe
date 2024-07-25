
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
// initGroup initializes the group object of redis.
<原文结束>

# <翻译开始>
// initGroup 初始化Redis的组对象。 md5:f3c3ba5dbd6196a2
# <翻译结束>


<原文开始>
// SetAdapter changes the underlying adapter with custom adapter for current redis client.
<原文结束>

# <翻译开始>
// SetAdapter 将当前 Redis 客户端的底层适配器替换为自定义适配器。 md5:f503f97750dcd95a
# <翻译结束>


<原文开始>
// GetAdapter returns the adapter that is set in current redis client.
<原文结束>

# <翻译开始>
// GetAdapter 返回当前 Redis 客户端设置的适配器。 md5:c46228b935b43204
# <翻译结束>


<原文开始>
// Conn retrieves and returns a connection object for continuous operations.
// Note that you should call Close function manually if you do not use this connection any further.
<原文结束>

# <翻译开始>
// Conn 获取并返回一个用于连续操作的连接对象。
// 请注意，如果您不再使用此连接，请手动调用 Close 函数。
// md5:b0379f4ab8131447
# <翻译结束>


<原文开始>
// Do send a command to the server and returns the received reply.
// It uses json.Marshal for struct/slice/map type values before committing them to redis.
<原文结束>

# <翻译开始>
// Do 向服务器发送命令并返回接收到的回复。
// 它在将结构体、切片或映射类型值提交到Redis之前，使用json.Marshal进行序列化。
// md5:bbe59d4e1ff07fa3
# <翻译结束>


<原文开始>
// MustConn performs as function Conn, but it panics if any error occurs internally.
<原文结束>

# <翻译开始>
// MustConn 表现如同 Conn 函数，但是如果内部发生任何错误，它将引发 panic。 md5:555eb0f8f348b94c
# <翻译结束>


<原文开始>
// MustDo performs as function Do, but it panics if any error occurs internally.
<原文结束>

# <翻译开始>
// MustDo 执行与 Do 相同的操作，但如果内部出现任何错误，它将引发 panic。 md5:0d30101f0e9e6a4e
# <翻译结束>


<原文开始>
// Close closes current redis client, closes its connection pool and releases all its related resources.
<原文结束>

# <翻译开始>
// Close 方法关闭当前Redis客户端，关闭其连接池并释放所有相关资源。 md5:bfd91d0269572038
# <翻译结束>

