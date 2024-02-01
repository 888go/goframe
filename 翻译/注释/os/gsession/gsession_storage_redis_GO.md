
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
// StorageRedis implements the Session Storage interface with redis.
<原文结束>

# <翻译开始>
// StorageRedis 使用 Redis 实现了 Session 存储接口。
# <翻译结束>


<原文开始>
// Redis client for session storage.
<原文结束>

# <翻译开始>
// Redis客户端用于会话存储。
# <翻译结束>


<原文开始>
// Redis key prefix for session id.
<原文结束>

# <翻译开始>
// Redis中用于session id的键前缀。
# <翻译结束>


<原文开始>
// Updating TTL set for session id.
<原文结束>

# <翻译开始>
// 更新给定会话ID的TTL集合。
# <翻译结束>


<原文开始>
	// DefaultStorageRedisLoopInterval is the interval updating TTL for session ids
	// in last duration.
<原文结束>

# <翻译开始>
// DefaultStorageRedisLoopInterval 是更新在最后持续时间段内会话ID的TTL（生存时间）的间隔。
# <翻译结束>


<原文开始>
// NewStorageRedis creates and returns a redis storage object for session.
<原文结束>

# <翻译开始>
// NewStorageRedis 创建并返回一个用于session的redis存储对象。
# <翻译结束>


<原文开始>
// Batch updates the TTL for session ids timely.
<原文结束>

# <翻译开始>
// 批量及时更新会话ID的TTL（生存时间）
# <翻译结束>


<原文开始>
// RemoveAll deletes all key-value pairs from storage.
<原文结束>

# <翻译开始>
// RemoveAll 从存储中删除所有键值对。
# <翻译结束>


<原文开始>
// GetSession returns the session data as *gmap.StrAnyMap for given session id from storage.
//
// The parameter `ttl` specifies the TTL for this session, and it returns nil if the TTL is exceeded.
// The parameter `data` is the current old session data stored in memory,
// and for some storage it might be nil if memory storage is disabled.
//
// This function is called ever when session starts.
<原文结束>

# <翻译开始>
// GetSession 通过给定的 session id 从存储中获取 session 数据，并以 *gmap.StrAnyMap 类型返回。
//
// 参数 `ttl` 指定了该 session 的生存时间（TTL），若生存时间已过，则返回 nil。
// 参数 `data` 是当前存储在内存中的旧 session 数据，如果禁用了内存存储，对于某些存储方式，此参数可能为 nil。
//
// 当每次 session 开始时，都会调用这个函数。
# <翻译结束>


<原文开始>
// SetSession updates the data map for specified session id.
// This function is called ever after session, which is changed dirty, is closed.
// This copy all session data map from memory to storage.
<原文结束>

# <翻译开始>
// SetSession 更新指定会话 ID 的数据映射。
// 在每次已标记为脏的、发生改变的会话关闭后，都会调用此函数。
// 此函数将内存中的所有会话数据映射复制到存储中。
# <翻译结束>


<原文开始>
// UpdateTTL updates the TTL for specified session id.
// This function is called ever after session, which is not dirty, is closed.
// It just adds the session id to the async handling queue.
<原文结束>

# <翻译开始>
// UpdateTTL 更新指定会话ID的TTL（生存时间）。
// 此函数在非脏数据会话关闭后调用。
// 它只是将该会话ID添加到异步处理队列中。
# <翻译结束>


<原文开始>
// doUpdateExpireForSession updates the TTL for session id.
<原文结束>

# <翻译开始>
// doUpdateExpireForSession 更新会话ID的TTL（生存时间）
# <翻译结束>


<原文开始>
// sessionIdToRedisKey converts and returns the redis key for given session id.
<原文结束>

# <翻译开始>
// sessionIdToRedisKey 将给定的session id转换并返回其对应的redis键。
# <翻译结束>

