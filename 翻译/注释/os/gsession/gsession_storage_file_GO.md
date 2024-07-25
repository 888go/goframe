
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
// StorageFile implements the Session Storage interface with file system.
<原文结束>

# <翻译开始>
// StorageFile实现了使用文件系统作为会话存储的接口。 md5:bae13bc406aa3178
# <翻译结束>


<原文开始>
// Session file storage folder path.
<原文结束>

# <翻译开始>
// 会话文件存储文件夹路径。 md5:a07352e7e4b2ee5d
# <翻译结束>


<原文开始>
// Used when enable crypto feature.
<原文结束>

# <翻译开始>
// 在启用加密功能时使用。 md5:e2b00bed77b9f059
# <翻译结束>


<原文开始>
// To be batched updated session id set.
<原文结束>

# <翻译开始>
// 用于批量更新的会话ID集合。 md5:704d1ffa9a08b42d
# <翻译结束>


<原文开始>
// NewStorageFile creates and returns a file storage object for session.
<原文结束>

# <翻译开始>
// NewStorageFile 创建并返回一个用于会话的文件存储对象。 md5:047619bd552117d1
# <翻译结束>


<原文开始>
// timelyUpdateSessionTTL batch updates the TTL for sessions timely.
<原文结束>

# <翻译开始>
// timelyUpdateSessionTTL 批量及时更新会话的超时时间。 md5:8d440f6681b47013
# <翻译结束>


<原文开始>
// Batch updating sessions.
<原文结束>

# <翻译开始>
	// 批量更新会话。 md5:db1f90067d27cc66
# <翻译结束>


<原文开始>
// timelyClearExpiredSessionFile deletes all expired files timely.
<原文结束>

# <翻译开始>
// timelyClearExpiredSessionFile 及时删除所有过期的文件。 md5:5f02dbf03c17d4ca
# <翻译结束>


<原文开始>
// SetCryptoKey sets the crypto key for session storage.
// The crypto key is used when crypto feature is enabled.
<原文结束>

# <翻译开始>
// SetCryptoKey 设置会话存储的加密密钥。
// 当启用加密功能时，将使用此加密密钥。
// md5:dbc53d710307bd28
# <翻译结束>


<原文开始>
// SetCryptoEnabled enables/disables the crypto feature for session storage.
<原文结束>

# <翻译开始>
// SetCryptoEnabled 启用/禁用会话存储的加密功能。 md5:14228b4577da32ec
# <翻译结束>


<原文开始>
// sessionFilePath returns the storage file path for given session id.
<原文结束>

# <翻译开始>
// sessionFilePath 根据给定的会话ID返回存储文件的路径。 md5:9cec805dff8d12a7
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
// It updates the TTL only if the session file already exists.
<原文结束>

# <翻译开始>
	// 只有当会话文件已经存在时，它才会更新TTL。 md5:a9223056bbc67ae2
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
// updateSessionTTL updates the TTL for specified session id.
<原文结束>

# <翻译开始>
// updateSessionTTL 更新指定会话ID的超时时间。 md5:1cff3164e4ca8226
# <翻译结束>


<原文开始>
// Read the session file updated timestamp in milliseconds.
<原文结束>

# <翻译开始>
	// 读取会话文件更新的时间戳（以毫秒为单位）。 md5:e3b93f5cb9dd863b
# <翻译结束>


<原文开始>
// Remove expired session file.
<原文结束>

# <翻译开始>
	// 删除过期的会话文件。 md5:f3e7a080ff4d0135
# <翻译结束>

