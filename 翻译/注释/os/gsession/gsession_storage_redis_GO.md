
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
// StorageRedis implements the Session Storage interface with redis.
<原文结束>

# <翻译开始>
// StorageRedis 使用Redis实现会话存储接口。 md5:df4e47711869aaf9
# <翻译结束>


<原文开始>
// Redis client for session storage.
<原文结束>

# <翻译开始>
// 用于session存储的Redis客户端。 md5:6ab8fcce48bcdda4
# <翻译结束>


<原文开始>
// Redis key prefix for session id.
<原文结束>

# <翻译开始>
// 会话ID的Redis键前缀。 md5:c0a31dd348ccaac3
# <翻译结束>


<原文开始>
// Updating TTL set for session id.
<原文结束>

# <翻译开始>
// 更新会话ID的TTL集合。 md5:5eacf4393666b98f
# <翻译结束>


<原文开始>
	// DefaultStorageRedisLoopInterval is the interval updating TTL for session ids
	// in last duration.
<原文结束>

# <翻译开始>
	// DefaultStorageRedisLoopInterval 是用于在最近一段时间内更新会话ID的TTL（生存时间）的间隔。
	// md5:5adbee0aa8ff1658
# <翻译结束>


<原文开始>
// NewStorageRedis creates and returns a redis storage object for session.
<原文结束>

# <翻译开始>
// NewStorageRedis 创建并返回一个用于session的redis存储对象。 md5:58528aab48b7daea
# <翻译结束>


<原文开始>
// Batch updates the TTL for session ids timely.
<原文结束>

# <翻译开始>
// 定期批量更新会话ID的TTL（时间到 live，生存时间）。 md5:81e845800fad5861
# <翻译结束>


<原文开始>
// RemoveAll deletes all key-value pairs from storage.
<原文结束>

# <翻译开始>
// RemoveAll 删除存储中的所有键值对。 md5:8b06607595d19a73
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
// GetSession 从存储中根据给定的会话ID获取会话数据，返回一个指向*gmap.StrAnyMap的指针。
//
// 参数`ttl`指定了此会话的有效期，如果超过有效期，则返回nil。参数`data`是当前存储在内存中的旧会话数据，对于某些存储方式，如果禁用了内存存储，它可能会为nil。
//
// 此函数在会话启动时会被调用。
// md5:01e56ce09d5fd934
# <翻译结束>


<原文开始>
// SetSession updates the data map for specified session id.
// This function is called ever after session, which is changed dirty, is closed.
// This copy all session data map from memory to storage.
<原文结束>

# <翻译开始>
// SetSession 根据指定的会话ID更新数据映射。
// 当某个被标记为脏（即发生过修改）的会话关闭后，将调用此函数。
// 该操作会将所有会话数据从内存复制到存储中。
// md5:1caa26989d884fa4
# <翻译结束>


<原文开始>
// UpdateTTL updates the TTL for specified session id.
// This function is called ever after session, which is not dirty, is closed.
// It just adds the session id to the async handling queue.
<原文结束>

# <翻译开始>
// UpdateTTL 更新指定会话ID的生存时间（TTL）。
// 当一个未被修改（非脏）的会话关闭后，此函数会被调用。
// 它只是将会话ID添加到异步处理队列中。
// md5:cc5ac287cbbc0eab
# <翻译结束>


<原文开始>
// doUpdateExpireForSession updates the TTL for session id.
<原文结束>

# <翻译开始>
// doUpdateExpireForSession 更新会话ID的过期时间。 md5:6a4d0561227fb192
# <翻译结束>


<原文开始>
// sessionIdToRedisKey converts and returns the redis key for given session id.
<原文结束>

# <翻译开始>
// sessionIdToRedisKey 将给定的会话ID转换并返回对应的Redis键。 md5:e18b9b593a10a025
# <翻译结束>

