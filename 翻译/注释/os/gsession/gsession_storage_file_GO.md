
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
// StorageFile implements the Session Storage interface with file system.
<原文结束>

# <翻译开始>
// StorageFile实现了使用文件系统作为Session存储接口。
# <翻译结束>


<原文开始>
// Session file storage folder path.
<原文结束>

# <翻译开始>
// 会话文件存储文件夹路径。
# <翻译结束>


<原文开始>
// Used when enable crypto feature.
<原文结束>

# <翻译开始>
// 当启用加密功能时使用。
# <翻译结束>


<原文开始>
// To be batched updated session id set.
<原文结束>

# <翻译开始>
// 待批量更新的会话ID集合。
# <翻译结束>


<原文开始>
// NewStorageFile creates and returns a file storage object for session.
<原文结束>

# <翻译开始>
// NewStorageFile 创建并返回一个用于存储session的文件存储对象。
# <翻译结束>


<原文开始>
// timelyUpdateSessionTTL batch updates the TTL for sessions timely.
<原文结束>

# <翻译开始>
// timelyUpdateSessionTTL 批量及时更新会话的TTL（生存时间）
# <翻译结束>


<原文开始>
// timelyClearExpiredSessionFile deletes all expired files timely.
<原文结束>

# <翻译开始>
// 定时清理过期会话文件，及时删除所有已过期的文件。
# <翻译结束>


<原文开始>
// SetCryptoKey sets the crypto key for session storage.
// The crypto key is used when crypto feature is enabled.
<原文结束>

# <翻译开始>
// SetCryptoKey 设置会话存储的加密密钥。
// 当启用加密功能时，会使用此加密密钥。
# <翻译结束>


<原文开始>
// SetCryptoEnabled enables/disables the crypto feature for session storage.
<原文结束>

# <翻译开始>
// SetCryptoEnabled 用于启用/禁用会话存储的加密功能。
# <翻译结束>


<原文开始>
// sessionFilePath returns the storage file path for given session id.
<原文结束>

# <翻译开始>
// sessionFilePath根据给定的session id返回存储文件路径。
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
// It updates the TTL only if the session file already exists.
<原文结束>

# <翻译开始>
// 如果会话文件已经存在，则仅更新TTL（生存时间）
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
// updateSessionTTL updates the TTL for specified session id.
<原文结束>

# <翻译开始>
// updateSessionTTL 更新指定会话ID的TTL（生存时间）
# <翻译结束>


<原文开始>
// Read the session file updated timestamp in milliseconds.
<原文结束>

# <翻译开始>
// 读取会话文件更新的毫秒级时间戳。
# <翻译结束>


<原文开始>
// Remove expired session file.
<原文结束>

# <翻译开始>
// 移除过期的会话文件。
# <翻译结束>


<原文开始>
// Batch updating sessions.
<原文结束>

# <翻译开始>
// 批量更新会话。
# <翻译结束>

