
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
// Package rwmutex provides switch of concurrent safety feature for sync.RWMutex.
<原文结束>

# <翻译开始>
// 包 rwmutex 为 sync.RWMutex 提供并发安全特性切换的功能。 md5:563f53220ab3eec8
# <翻译结束>


<原文开始>
// RWMutex is a sync.RWMutex with a switch for concurrent safe feature.
// If its attribute *sync.RWMutex is not nil, it means it's in concurrent safety usage.
// Its attribute *sync.RWMutex is nil in default, which makes this struct mush lightweight.
<原文结束>

# <翻译开始>
// RWMutex 是一个具有并发安全开关的 sync.RWMutex。
// 如果其 sync.RWMutex 类型的属性非空，表示它处于并发安全使用中。
// 默认情况下，它的 sync.RWMutex 属性为 nil，这使得该结构体更为轻量。
// md5:2d8d597983a75c36
# <翻译结束>


<原文开始>
// New creates and returns a new *RWMutex.
// The parameter `safe` is used to specify whether using this mutex in concurrent safety,
// which is false in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的 RWMutex 实例。
// 参数 `safe` 用于指定是否在并发环境中使用这个互斥锁，默认为 false，表示不安全。
// md5:e431e613f230b125
# <翻译结束>


<原文开始>
// Create creates and returns a new RWMutex object.
// The parameter `safe` is used to specify whether using this mutex in concurrent safety,
// which is false in default.
<原文结束>

# <翻译开始>
// Create 创建并返回一个新的 RWMutex 对象。
// 参数 `safe` 用于指定是否在并发安全模式下使用该互斥锁，其默认值为 false。
// md5:e40df278667779d2
# <翻译结束>


<原文开始>
// IsSafe checks and returns whether current mutex is in concurrent-safe usage.
<原文结束>

# <翻译开始>
// IsSafe 检查并返回当前互斥锁是否在并发安全的使用中。 md5:1a2c4197eb3278b5
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


<原文开始>
// RLock locks mutex for reading.
// It does nothing if it is not in concurrent-safe usage.
<原文结束>

# <翻译开始>
// RLock 用于对互斥锁进行读取锁定。
// 如果不是在并发安全的使用场景下，它不做任何操作。
// md5:61160c78e9bcccd5
# <翻译结束>


<原文开始>
// RUnlock unlocks mutex for reading.
// It does nothing if it is not in concurrent-safe usage.
<原文结束>

# <翻译开始>
// RUnlock 释放读取锁。
// 如果在非并发安全使用时，它将不执行任何操作。
// md5:834672a97d0bd47f
# <翻译结束>

