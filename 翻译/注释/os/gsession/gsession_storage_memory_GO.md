
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
// StorageMemory implements the Session Storage interface with memory.
<原文结束>

# <翻译开始>
// StorageMemory 使用内存实现了会话存储接口。 md5:1a9a78b3bd5a138b
# <翻译结束>


<原文开始>
	// cache is the memory data cache for session TTL,
	// which is available only if the Storage does not store any session data in synchronizing.
	// Please refer to the implements of StorageFile, StorageMemory and StorageRedis.
	//
	// Its value is type of `*gmap.StrAnyMap`.
<原文结束>

# <翻译开始>
// cache是用于session TTL的内存数据缓存， 
// 只有在Storage不存储任何会话数据时才可用（即不同步存储数据）。
// 请参考StorageFile, StorageMemory和StorageRedis的实现。
//
// 其值为`*gmap.StrAnyMap`类型。
// md5:c8273be50da58f8d
# <翻译结束>


<原文开始>
// NewStorageMemory creates and returns a file storage object for session.
<原文结束>

# <翻译开始>
// NewStorageMemory 创建并返回一个用于会话的内存存储对象。 md5:9b1b616d48dd808e
# <翻译结束>


<原文开始>
// RemoveAll deletes session from storage.
<原文结束>

# <翻译开始>
// RemoveAll 从存储中删除会话。 md5:488d9f9ca747e8e4
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
// Retrieve memory session data from manager.
<原文结束>

# <翻译开始>
// 从管理器中获取内存会话数据。 md5:9a3be5b3f3de62f6
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

