
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
// StorageBase is a base implement for Session Storage.
<原文结束>

# <翻译开始>
// StorageBase 是 Session 存储的一个基础实现。
# <翻译结束>


<原文开始>
// New creates a session id.
// This function can be used for custom session creation.
<原文结束>

# <翻译开始>
// New 创建一个会话ID。
// 该函数可用于自定义会话创建。
# <翻译结束>


<原文开始>
// Get retrieves certain session value with given key.
// It returns nil if the key does not exist in the session.
<原文结束>

# <翻译开始>
// Get 通过给定的键获取特定会话值。
// 如果键在会话中不存在，则返回空值（nil）。
# <翻译结束>


<原文开始>
// Data retrieves all key-value pairs as map from storage.
<原文结束>

# <翻译开始>
// Data 从存储中检索所有的键值对并以map形式返回。
# <翻译结束>


<原文开始>
// GetSize retrieves the size of key-value pairs from storage.
<原文结束>

# <翻译开始>
// GetSize 从存储中检索键值对的大小。
# <翻译结束>


<原文开始>
// Set sets key-value session pair to the storage.
// The parameter `ttl` specifies the TTL for the session id (not for the key-value pair).
<原文结束>

# <翻译开始>
// Set 将键值对会话设置到存储中。
// 参数 `ttl` 指定了会话ID的生存时间（并非针对键值对）。
# <翻译结束>


<原文开始>
// SetMap batch sets key-value session pairs with map to the storage.
// The parameter `ttl` specifies the TTL for the session id(not for the key-value pair).
<原文结束>

# <翻译开始>
// SetMap 批量将键值对集合设置到存储中。
// 参数 `ttl` 指定的是会话ID的TTL（生存时间），而不是键值对的生存时间。
# <翻译结束>


<原文开始>
// Remove deletes key with its value from storage.
<原文结束>

# <翻译开始>
// Remove 从存储中删除指定键及其对应的值。
# <翻译结束>


<原文开始>
// RemoveAll deletes session from storage.
<原文结束>

# <翻译开始>
// RemoveAll 从存储中删除session。
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

