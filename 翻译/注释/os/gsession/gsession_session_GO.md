
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
// Session struct for storing single session data, which is bound to a single request.
// The Session struct is the interface with user, but the Storage is the underlying adapter designed interface
// for functionality implements.
<原文结束>

# <翻译开始>
// Session 结构体，用于存储单个会话数据，它与单个请求绑定。Session 结构体是与用户交互的接口，但 Storage 是底层适配器设计的接口，用于实现特定功能。
// md5:1d1b86dcb53a276e
# <翻译结束>


<原文开始>
// Session id. It retrieves the session if id is custom specified.
<原文结束>

# <翻译开始>
// 会话ID。如果指定了自定义的id，则用于获取会话。 md5:7fec8def5a7d635f
# <翻译结束>


<原文开始>
// Context for current session. Please note that, session lives along with context.
<原文结束>

# <翻译开始>
// 当前会话的上下文。请注意，会话与上下文一起存在。 md5:c61c2b0f3688fce4
# <翻译结束>


<原文开始>
// Current Session data, which is retrieved from Storage.
<原文结束>

# <翻译开始>
// 当前会话数据，从存储中获取。 md5:00861bf25945ffcd
# <翻译结束>


<原文开始>
// Used to mark session is modified.
<原文结束>

# <翻译开始>
// 用于标记会话已修改。 md5:6d1cc43a943b4c84
# <翻译结束>


<原文开始>
// Used to mark session is started.
<原文结束>

# <翻译开始>
// 用于标记会话已开始。 md5:213288ad3376abfc
# <翻译结束>


<原文开始>
// Parent session Manager.
<原文结束>

# <翻译开始>
// 父级会话管理器。 md5:6b920db9951d3f89
# <翻译结束>


<原文开始>
	// idFunc is a callback function used for creating custom session id.
	// This is called if session id is empty ever when session starts.
<原文结束>

# <翻译开始>
	// idFunc是一个用于创建自定义会话ID的回调函数。当会话开始时，如果会话ID为空，就会调用这个函数。
	// md5:83b853b0118605bd
# <翻译结束>


<原文开始>
// init does the lazy initialization for session, which retrieves the session if session id is specified,
// or else it creates a new empty session.
<原文结束>

# <翻译开始>
// init函数进行懒惰初始化session，如果指定了session ID，则获取session，否则创建一个新的空session。
// md5:60a85a2954d0a427
# <翻译结束>


<原文开始>
// Retrieve stored session data from storage.
<原文结束>

# <翻译开始>
		// 从存储中检索存储的会话数据。 md5:e6c3f2bdc143f93c
# <翻译结束>


<原文开始>
// Use custom session id creating function.
<原文结束>

# <翻译开始>
			// 使用自定义的会话ID创建函数。 md5:4c1e5f997d31f1b3
# <翻译结束>


<原文开始>
// Use default session id creating function of storage.
<原文结束>

# <翻译开始>
			// 使用存储的默认会话ID创建函数。 md5:11db3a1576d0231f
# <翻译结束>


<原文开始>
			// If session storage does not implements id generating functionality,
			// it then uses default session id creating function.
<原文结束>

# <翻译开始>
			// 如果会话存储不实现ID生成功能，那么它将使用默认的会话ID创建函数。
			// md5:4f2bd0ddc795fde4
# <翻译结束>


<原文开始>
// Close closes current session and updates its ttl in the session manager.
// If this session is dirty, it also exports it to storage.
//
// NOTE that this function must be called ever after a session request done.
<原文结束>

# <翻译开始>
// Close 方法关闭当前会话并在会话管理器中更新其TTL（生存时间）。
// 如果此会话已被修改（脏会话），它还会将该会话导出到存储中。
//
// 注意：此功能必须在每次会话请求完成后调用。
// md5:f68a83f493f4727a
# <翻译结束>


<原文开始>
// Set sets key-value pair to this session.
<原文结束>

# <翻译开始>
// Set 将键值对设置到这个会话中。 md5:09e1539c4a50fcfd
# <翻译结束>


<原文开始>
// SetMap batch sets the session using map.
<原文结束>

# <翻译开始>
// SetMap 批量使用映射设置会话。 md5:f55c78b98e85ba61
# <翻译结束>


<原文开始>
// Remove removes key along with its value from this session.
<原文结束>

# <翻译开始>
// Remove 从本次会话中移除指定的键及其对应的值。 md5:3dc440da200c0834
# <翻译结束>


<原文开始>
// RemoveAll deletes all key-value pairs from this session.
<原文结束>

# <翻译开始>
// RemoveAll 从该会话中删除所有键值对。 md5:6ca756339a9f18b5
# <翻译结束>


<原文开始>
// Remove data from memory.
<原文结束>

# <翻译开始>
	// 从内存中移除数据。 md5:47322b1cdcaf7596
# <翻译结束>


<原文开始>
// Id returns the session id for this session.
// It creates and returns a new session id if the session id is not passed in initialization.
<原文结束>

# <翻译开始>
// Id 返回此会话的会话标识符。
// 如果在初始化时未传递会话标识符，则创建并返回新的会话标识符。
// md5:c1a4c6b98633e656
# <翻译结束>


<原文开始>
// SetId sets custom session before session starts.
// It returns error if it is called after session starts.
<原文结束>

# <翻译开始>
// SetId 在会话开始前设置自定义会话。如果在会话已经开始后调用，将返回错误。
// md5:cf8fd98a6cd07079
# <翻译结束>


<原文开始>
// SetIdFunc sets custom session id creating function before session starts.
// It returns error if it is called after session starts.
<原文结束>

# <翻译开始>
// SetIdFunc 在会话开始前设置自定义会话ID生成函数。
// 如果在会话已经开始后调用它，将返回错误。
// md5:07c5962c3c68bf37
# <翻译结束>


<原文开始>
// Data returns all data as map.
// Note that it's using value copy internally for concurrent-safe purpose.
<原文结束>

# <翻译开始>
// Data 将所有数据作为映射返回。
// 请注意，为了并发安全，它内部使用了值拷贝。
// md5:a37827aba4dd5df4
# <翻译结束>


<原文开始>
// Size returns the size of the session.
<原文结束>

# <翻译开始>
// Size返回会话的大小。 md5:072795e87a3938d1
# <翻译结束>


<原文开始>
// Contains checks whether key exist in the session.
<原文结束>

# <翻译开始>
// Contains 检查键是否存在于会话中。 md5:7a03d1ea75cda393
# <翻译结束>


<原文开始>
// IsDirty checks whether there's any data changes in the session.
<原文结束>

# <翻译开始>
// IsDirty 检查会话中是否有数据变更。 md5:2a726ce013b067fe
# <翻译结束>


<原文开始>
// Get retrieves session value with given key.
// It returns `def` if the key does not exist in the session if `def` is given,
// or else it returns nil.
<原文结束>

# <翻译开始>
// Get 通过给定的键获取 session 值。
// 如果键在 session 中不存在且提供了 `def`，则返回 `def`，
// 否则返回 nil。
// md5:893a612d87b25ee2
# <翻译结束>


<原文开始>
// MustId performs as function Id, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustId 行为就像Id函数一样，但如果发生任何错误，它会引发恐慌。 md5:a51e8673adaf6727
# <翻译结束>


<原文开始>
// MustGet performs as function Get, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGet执行与Get相同的功能，但如果发生任何错误，它将引发恐慌。 md5:bdc72a85510733d5
# <翻译结束>


<原文开始>
// MustSet performs as function Set, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustSet 的功能与 Set 函数相同，但如果发生任何错误，它会直接 panic。 md5:06fa308e1636bcfa
# <翻译结束>


<原文开始>
// MustSetMap performs as function SetMap, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustSetMap 行为类似于函数 SetMap，但如果发生任何错误则会引发 panic。 md5:3d54948e22292bcf
# <翻译结束>


<原文开始>
// MustContains performs as function Contains, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustContains执行Contains函数的功能，但如果发生任何错误，它将引发恐慌。 md5:b9f29f0374157bc5
# <翻译结束>


<原文开始>
// MustData performs as function Data, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustData 执行与函数 Data 相同的操作，但如果发生任何错误，它将引发恐慌。 md5:ae01e79f6f27c9fe
# <翻译结束>


<原文开始>
// MustSize performs as function Size, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustSize 的行为与 Size 函数相同，但如果发生任何错误，它会直接 panic。 md5:d9d8c4724cdd0db4
# <翻译结束>


<原文开始>
// MustRemove performs as function Remove, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustRemove 行为与函数 Remove 相同，但如果发生任何错误则会引发恐慌。 md5:76bd8c9cb1e6223b
# <翻译结束>

