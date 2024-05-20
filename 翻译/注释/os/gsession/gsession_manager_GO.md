
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
// Storage interface for session storage.
<原文结束>

# <翻译开始>
// 会话存储的接口。. md5:75545bf4dbeae018
# <翻译结束>


<原文开始>
// New creates and returns a new session manager.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的会话管理器。. md5:f41a96ed9e4273e4
# <翻译结束>


<原文开始>
// It uses StorageFile in default.
<原文结束>

# <翻译开始>
// 默认使用 StorageFile。. md5:a8eedc355767eccd
# <翻译结束>


<原文开始>
// New creates or fetches the session for given session id.
// The parameter `sessionId` is optional, it creates a new one if not it's passed
// depending on Storage.New.
<原文结束>

# <翻译开始>
// New 为给定的 session ID 创建或获取会话。
// 参数 `sessionId` 是可选的，如果未提供，则根据 Storage.New 的行为创建新的会话。
// md5:4d84930c3cbf9027
# <翻译结束>


<原文开始>
// SetStorage sets the session storage for manager.
<原文结束>

# <翻译开始>
// SetStorage 设置管理器的会话存储。. md5:9fe6b2a9a6ae9990
# <翻译结束>


<原文开始>
// GetStorage returns the session storage of current manager.
<原文结束>

# <翻译开始>
// GetStorage 返回当前会话管理器的存储对象。. md5:43cdd2b5155f8389
# <翻译结束>


<原文开始>
// SetTTL the TTL for the session manager.
<原文结束>

# <翻译开始>
// SetTTL 设置会话管理器的生存时间（TTL）。. md5:bba913d23693cf2a
# <翻译结束>


<原文开始>
// GetTTL returns the TTL of the session manager.
<原文结束>

# <翻译开始>
// GetTTL 返回会话管理器的TTL（时间到 live，生存时间）。. md5:d0733ac8b424fbe1
# <翻译结束>

