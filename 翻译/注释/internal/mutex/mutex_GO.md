
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
// Package mutex provides switch of concurrent safe feature for sync.Mutex.
<原文结束>

# <翻译开始>
// 包mutex为sync.Mutex提供了并发安全开关的功能。. md5:2d280f557de1d7a8
# <翻译结束>


<原文开始>
// Mutex is a sync.Mutex with a switch for concurrent safe feature.
<原文结束>

# <翻译开始>
// Mutex是一个带有并发安全功能开关的sync.Mutex。. md5:8889db913bd4aa9f
# <翻译结束>


<原文开始>
// New creates and returns a new *Mutex.
// The parameter `safe` is used to specify whether using this mutex in concurrent safety,
// which is false in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的 *Mutex。
// 参数 `safe` 用于指定是否在并发安全情况下使用此互斥锁，其默认值为 false。
// md5:4b9e38d55d8b7828
# <翻译结束>


<原文开始>
// Create creates and returns a new Mutex object.
// The parameter `safe` is used to specify whether using this mutex in concurrent safety,
// which is false in default.
<原文结束>

# <翻译开始>
// Create 创建并返回一个新的 Mutex 对象。
// 参数 `safe` 用于指定是否在并发安全的环境下使用此 mutex，默认为 false。
// md5:a3db2fe6cfb0197f
# <翻译结束>


<原文开始>
// IsSafe checks and returns whether current mutex is in concurrent-safe usage.
<原文结束>

# <翻译开始>
// IsSafe 检查并返回当前互斥锁是否在并发安全的使用中。. md5:1a2c4197eb3278b5
# <翻译结束>


<原文开始>
// Lock locks mutex for writing.
// It does nothing if it is not in concurrent-safe usage.
<原文结束>

# <翻译开始>
// Lock 为写入锁定互斥量。如果没有进行并发安全使用，它不会做任何事情。
// md5:e7a0e420dc8d74c3
# <翻译结束>


<原文开始>
// Unlock unlocks mutex for writing.
// It does nothing if it is not in concurrent-safe usage.
<原文结束>

# <翻译开始>
// Unlock 为写操作解锁互斥锁。如果它不是在并发安全模式下使用，则不会做任何事情。
// md5:ce0b3215f968f29c
# <翻译结束>

