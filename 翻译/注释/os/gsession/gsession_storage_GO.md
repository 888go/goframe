
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
// Storage is the interface definition for session storage.
<原文结束>

# <翻译开始>
// Storage 是会话存储的接口定义。
# <翻译结束>


<原文开始>
	// New creates a custom session id.
	// This function can be used for custom session creation.
<原文结束>

# <翻译开始>
// New 创建一个自定义会话ID。
// 该函数可用于创建自定义会话。
# <翻译结束>


<原文开始>
	// Get retrieves and returns certain session value with given key.
	// It returns nil if the key does not exist in the session.
<原文结束>

# <翻译开始>
// Get 函数通过给定的键获取并返回特定的 session 值。
// 如果键在 session 中不存在，则返回 nil。
# <翻译结束>


<原文开始>
// GetSize retrieves and returns the size of key-value pairs from storage.
<原文结束>

# <翻译开始>
// GetSize 从存储中获取并返回键值对的大小。
# <翻译结束>


<原文开始>
// Data retrieves all key-value pairs as map from storage.
<原文结束>

# <翻译开始>
// Data 从存储中检索所有的键值对并以map形式返回。
# <翻译结束>


<原文开始>
	// Set sets one key-value session pair to the storage.
	// The parameter `ttl` specifies the TTL for the session id.
<原文结束>

# <翻译开始>
// Set 将一个键值对会话数据设置到存储中。
// 参数 `ttl` 指定该会话ID的生存时间（TTL，Time To Live）。
# <翻译结束>


<原文开始>
	// SetMap batch sets key-value session pairs as map to the storage.
	// The parameter `ttl` specifies the TTL for the session id.
<原文结束>

# <翻译开始>
// SetMap 批量将键值对形式的session设置到存储中。
// 参数 `ttl` 指定该session id的有效期（TTL，Time To Live）。
# <翻译结束>


<原文开始>
// Remove deletes key-value pair from specified session from storage.
<原文结束>

# <翻译开始>
// Remove 从存储中删除指定会话的键值对。
# <翻译结束>


<原文开始>
// RemoveAll deletes session from storage.
<原文结束>

# <翻译开始>
// RemoveAll 从存储中删除session。
# <翻译结束>


<原文开始>
	// GetSession returns the session data as `*gmap.StrAnyMap` for given session from storage.
	//
	// The parameter `ttl` specifies the TTL for this session.
	// The parameter `data` is the current old session data stored in memory,
	// and for some storage it might be nil if memory storage is disabled.
	//
	// This function is called ever when session starts.
	// It returns nil if the session does not exist or its TTL is expired.
<原文结束>

# <翻译开始>
// GetSession 从存储中获取给定会话的 session 数据，并以 `*gmap.StrAnyMap` 类型返回。
//
// 参数 `ttl` 指定了本次会话的生存时间（TTL）。
// 参数 `data` 是当前存储在内存中的旧会话数据，
// 如果禁用了内存存储，对于某些存储方式而言，这个参数可能会是 nil。
//
// 在每次会话启动时都会调用此函数。
// 如果会话不存在或者其 TTL 已经过期，则返回 nil。
# <翻译结束>


<原文开始>
	// SetSession updates the data for specified session id.
	// This function is called ever after session, which is changed dirty, is closed.
	// This copy all session data map from memory to storage.
<原文结束>

# <翻译开始>
// SetSession 更新指定会话 ID 的数据。
// 当发生更改且变为脏状态的会话关闭后，将调用此函数。
// 此函数将内存中所有会话数据映射复制到存储中。
# <翻译结束>


<原文开始>
	// UpdateTTL updates the TTL for specified session id.
	// This function is called ever after session, which is not dirty, is closed.
<原文结束>

# <翻译开始>
// UpdateTTL 更新指定会话ID的TTL（生存时间）。
// 此函数在非脏数据会话关闭后调用。
# <翻译结束>

