
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
// Session struct for storing single session data, which is bound to a single request.
// The Session struct is the interface with user, but the Storage is the underlying adapter designed interface
// for functionality implements.
<原文结束>

# <翻译开始>
// Session 结构体用于存储单个会话数据，它与单个请求绑定。
// Session 结构体是与用户交互的接口，而 Storage 是为功能实现设计的底层适配器接口。
# <翻译结束>


<原文开始>
// Session id. It retrieves the session if id is custom specified.
<原文结束>

# <翻译开始>
// 会话ID。如果自定义指定了id，它将获取该会话。
# <翻译结束>


<原文开始>
// Context for current session. Please note that, session lives along with context.
<原文结束>

# <翻译开始>
// 当前会话的上下文。请注意，会话与上下文共存。
# <翻译结束>


<原文开始>
// Current Session data, which is retrieved from Storage.
<原文结束>

# <翻译开始>
// 当前会话数据，从Storage中获取。
# <翻译结束>


<原文开始>
// Used to mark session is modified.
<原文结束>

# <翻译开始>
// 用于标记会话已被修改。
# <翻译结束>


<原文开始>
// Used to mark session is started.
<原文结束>

# <翻译开始>
// 用于标记会话已开始。
# <翻译结束>


<原文开始>
	// idFunc is a callback function used for creating custom session id.
	// This is called if session id is empty ever when session starts.
<原文结束>

# <翻译开始>
// idFunc 是一个用于创建自定义会话 ID 的回调函数。
// 当会话开始且会话 ID 为空时，将会调用这个函数。
# <翻译结束>


<原文开始>
// init does the lazy initialization for session, which retrieves the session if session id is specified,
// or else it creates a new empty session.
<原文结束>

# <翻译开始>
// init 执行会话的延迟初始化。如果指定了会话ID，则从存储中获取该会话；否则，创建一个新的空会话。
# <翻译结束>


<原文开始>
// Retrieve stored session data from storage.
<原文结束>

# <翻译开始>
// 从存储中检索已存储的会话数据。
# <翻译结束>


<原文开始>
// Use custom session id creating function.
<原文结束>

# <翻译开始>
// 使用自定义的会话ID创建函数。
# <翻译结束>


<原文开始>
// Use default session id creating function of storage.
<原文结束>

# <翻译开始>
// 使用存储的默认会话ID创建函数。
# <翻译结束>


<原文开始>
			// If session storage does not implements id generating functionality,
			// it then uses default session id creating function.
<原文结束>

# <翻译开始>
// 如果会话存储未实现ID生成功能，
// 则使用默认的会话ID创建函数。
# <翻译结束>


<原文开始>
// Close closes current session and updates its ttl in the session manager.
// If this session is dirty, it also exports it to storage.
//
// NOTE that this function must be called ever after a session request done.
<原文结束>

# <翻译开始>
// Close 关闭当前会话并在会话管理器中更新其TTL（Time To Live，生存时间）。
// 如果此会话是脏的（即已修改），它还会将其导出到存储中。
//
// 注意：此函数必须在每次会话请求完成后调用。
# <翻译结束>


<原文开始>
// Set sets key-value pair to this session.
<原文结束>

# <翻译开始>
// Set 将键值对设置到此会话中。
# <翻译结束>


<原文开始>
// SetMap batch sets the session using map.
<原文结束>

# <翻译开始>
// SetMap 批量使用 map 设置 session。
# <翻译结束>


<原文开始>
// Remove removes key along with its value from this session.
<原文结束>

# <翻译开始>
// Remove 从当前会话中移除指定键及其对应的值。
# <翻译结束>


<原文开始>
// RemoveAll deletes all key-value pairs from this session.
<原文结束>

# <翻译开始>
// RemoveAll 从该会话中删除所有键值对。
# <翻译结束>


<原文开始>
// Id returns the session id for this session.
// It creates and returns a new session id if the session id is not passed in initialization.
<原文结束>

# <翻译开始>
// Id 返回当前会话的 session id。
// 如果在初始化时没有传递 session id，它将创建并返回一个新的 session id。
# <翻译结束>


<原文开始>
// SetId sets custom session before session starts.
// It returns error if it is called after session starts.
<原文结束>

# <翻译开始>
// SetId 在会话开始前设置自定义会话ID。
// 如果在会话开始后调用，将返回错误。
# <翻译结束>


<原文开始>
// SetIdFunc sets custom session id creating function before session starts.
// It returns error if it is called after session starts.
<原文结束>

# <翻译开始>
// SetIdFunc 在会话开始前设置自定义的 session id 生成函数。 
// 如果在会话开始后调用，将返回错误。
# <翻译结束>


<原文开始>
// Data returns all data as map.
// Note that it's using value copy internally for concurrent-safe purpose.
<原文结束>

# <翻译开始>
// Data 返回所有数据作为映射。
// 注意，为了保证并发安全，内部使用了值复制。
# <翻译结束>


<原文开始>
// Size returns the size of the session.
<原文结束>

# <翻译开始>
// Size 返回会话的大小。
# <翻译结束>


<原文开始>
// Contains checks whether key exist in the session.
<原文结束>

# <翻译开始>
// Contains 检查键是否存在于会话中。
# <翻译结束>


<原文开始>
// IsDirty checks whether there's any data changes in the session.
<原文结束>

# <翻译开始>
// IsDirty 检查会话中是否存在任何数据更改。
# <翻译结束>


<原文开始>
// Get retrieves session value with given key.
// It returns `def` if the key does not exist in the session if `def` is given,
// or else it returns nil.
<原文结束>

# <翻译开始>
// Get 通过给定的键从会话中检索值。
// 如果提供了 `def`，当键在会话中不存在时，它将返回 `def`，
// 否则返回 nil。
# <翻译结束>


<原文开始>
// MustId performs as function Id, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustId 函数表现如同 Id 函数，但是当发生任何错误时，它会触发panic。
# <翻译结束>


<原文开始>
// MustGet performs as function Get, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustGet 函数表现如同 Get 函数，但当发生任何错误时，它会触发panic。
# <翻译结束>


<原文开始>
// MustSet performs as function Set, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustSet 函数表现如同 Set 函数，但是当发生任何错误时它会引发panic。
# <翻译结束>


<原文开始>
// MustSetMap performs as function SetMap, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustSetMap表现如同函数 SetMap，但是当发生任何错误时它会触发panic。
# <翻译结束>


<原文开始>
// MustContains performs as function Contains, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustContains 执行与 Contains 函数相同的功能，但是当发生任何错误时，它会触发panic。
# <翻译结束>


<原文开始>
// MustData performs as function Data, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustData 的行为类似于函数 Data，但是当发生任何错误时，它会触发 panic。
# <翻译结束>


<原文开始>
// MustSize performs as function Size, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustSize 执行与 Size 函数相同的功能，但是当发生任何错误时，它会触发 panic。
# <翻译结束>


<原文开始>
// MustRemove performs as function Remove, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustRemove 的行为与 Remove 函数相同，但是当发生任何错误时，它会触发 panic。
# <翻译结束>


<原文开始>
// Parent session Manager.
<原文结束>

# <翻译开始>
// 父级会话管理器。
# <翻译结束>


<原文开始>
// Session retrieving.
<原文结束>

# <翻译开始>
// 会话获取。
# <翻译结束>


<原文开始>
// Session id creation.
<原文结束>

# <翻译开始>
// 会话ID创建
# <翻译结束>


<原文开始>
// Remove data from memory.
<原文结束>

# <翻译开始>
// 从内存中移除数据。
# <翻译结束>

