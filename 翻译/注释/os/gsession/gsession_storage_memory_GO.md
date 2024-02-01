
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
// StorageMemory implements the Session Storage interface with memory.
<原文结束>

# <翻译开始>
// StorageMemory 实现了基于内存的 Session 存储接口。
# <翻译结束>


<原文开始>
	// cache is the memory data cache for session TTL,
	// which is available only if the Storage does not store any session data in synchronizing.
	// Please refer to the implements of StorageFile, StorageMemory and StorageRedis.
	//
	// Its value is type of `*gmap.StrAnyMap`.
<原文结束>

# <翻译开始>
// cache 是用于会话TTL（生存时间）的内存数据缓存，
// 只有在Storage在同步时不存储任何会话数据时才可用。
// 请参考StorageFile、StorageMemory和StorageRedis的具体实现。
//
// 其值的类型为`*gmap.StrAnyMap`。
# <翻译结束>


<原文开始>
// NewStorageMemory creates and returns a file storage object for session.
<原文结束>

# <翻译开始>
// NewStorageMemory 创建并返回一个用于存储session的内存文件存储对象。
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
// Retrieve memory session data from manager.
<原文结束>

# <翻译开始>
// 从管理器中检索内存会话数据。
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

