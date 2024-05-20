
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
// Storage is the interface definition for session storage.
<原文结束>

# <翻译开始>
// Storage是会话存储的接口定义。. md5:3c03cfdd3299edcc
# <翻译结束>


<原文开始>
	// New creates a custom session id.
	// This function can be used for custom session creation.
<原文结束>

# <翻译开始>
// New 创建一个自定义会话ID。
// 此函数可用于自定义会话创建。
// md5:bf8b403018c5c6df
# <翻译结束>


<原文开始>
	// Get retrieves and returns certain session value with given key.
	// It returns nil if the key does not exist in the session.
<原文结束>

# <翻译开始>
// Get 通过给定的键获取并返回会话中的特定值。
// 如果键在会话中不存在，则返回nil。
// md5:2584a452a5632118
# <翻译结束>


<原文开始>
// GetSize retrieves and returns the size of key-value pairs from storage.
<原文结束>

# <翻译开始>
// GetSize 从存储中获取并返回键值对的大小。. md5:2c41726f18e2cd04
# <翻译结束>


<原文开始>
// Data retrieves all key-value pairs as map from storage.
<原文结束>

# <翻译开始>
// Data 从存储中获取所有的键值对并将其作为映射返回。. md5:7160c6695dcc211b
# <翻译结束>


<原文开始>
	// Set sets one key-value session pair to the storage.
	// The parameter `ttl` specifies the TTL for the session id.
<原文结束>

# <翻译开始>
// Set 将一个键值对设置到存储中。
// 参数 `ttl` 指定了会话 ID 的过期时间。
// md5:f141e9b5de211364
# <翻译结束>


<原文开始>
	// SetMap batch sets key-value session pairs as map to the storage.
	// The parameter `ttl` specifies the TTL for the session id.
<原文结束>

# <翻译开始>
// SetMap 批量将键值对设置为存储中的会话映射。参数 `ttl` 指定会话 ID 的过期时间。
// md5:be3d6b9412b66e49
# <翻译结束>


<原文开始>
// Remove deletes key-value pair from specified session from storage.
<原文结束>

# <翻译开始>
// Remove 从存储中删除指定会话中的键值对。. md5:3887f6d1acd56ad6
# <翻译结束>


<原文开始>
// RemoveAll deletes session from storage.
<原文结束>

# <翻译开始>
// RemoveAll 从存储中删除会话。. md5:488d9f9ca747e8e4
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
// GetSession 从存储中返回给定会话的数据，数据类型为 `*gmap.StrAnyMap`。
//
// 参数 `ttl` 指定了该会话的生存时间（TTL）。
// 参数 `data` 是当前存储在内存中的旧会话数据，对于某些存储，如果禁用了内存存储，此参数可能为 `nil`。
//
// 这个函数会在每次会话开始时被调用。
// 如果会话不存在或者其 TTL 已过期，它将返回 `nil`。
// md5:a495b20f42259c94
# <翻译结束>


<原文开始>
	// SetSession updates the data for specified session id.
	// This function is called ever after session, which is changed dirty, is closed.
	// This copy all session data map from memory to storage.
<原文结束>

# <翻译开始>
// SetSession 更新指定会话ID的数据。
// 在关闭已更改的会话后，都会调用此函数。这个函数将内存中的所有会话数据映射复制到存储中。
// md5:16766d7e58c61924
# <翻译结束>


<原文开始>
	// UpdateTTL updates the TTL for specified session id.
	// This function is called ever after session, which is not dirty, is closed.
<原文结束>

# <翻译开始>
// UpdateTTL 更新指定会话ID的TTL（时间到 live）。
// 在非脏会话关闭后，将调用此函数。
// md5:29eae01946af2846
# <翻译结束>

