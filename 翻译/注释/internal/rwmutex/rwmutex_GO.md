
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
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Package rwmutex provides switch of concurrent safety feature for sync.RWMutex.
<原文结束>

# <翻译开始>
// Package rwmutex 为 sync.RWMutex 提供了并发安全特性开关。
# <翻译结束>


<原文开始>
// RWMutex is a sync.RWMutex with a switch for concurrent safe feature.
// If its attribute *sync.RWMutex is not nil, it means it's in concurrent safety usage.
// Its attribute *sync.RWMutex is nil in default, which makes this struct mush lightweight.
<原文结束>

# <翻译开始>
// RWMutex 是一个带并发安全开关的 sync.RWMutex。
// 如果其属性 *sync.RWMutex 非空，表示它处于并发安全使用状态。
// 默认情况下，其属性 *sync.RWMutex 为空（nil），这使得该结构体极为轻量级。
// 这段代码的注释翻译成中文如下：
// ```go
// RWMutex 是对标准库sync.RWMutex的一种扩展，增加了并发安全特性切换功能。
// 若其成员变量 *sync.RWMutex 不为 nil，表明当前它正在以支持并发安全的方式使用。
// 默认情况下，其成员变量 *sync.RWMutex 初始化为 nil，这样的设计使得该结构体保持轻量。
# <翻译结束>







<原文开始>
// New creates and returns a new *RWMutex.
// The parameter `safe` is used to specify whether using this mutex in concurrent safety,
// which is false in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的 *RWMutex。
// 参数 `safe` 用于指定是否在并发安全的情况下使用此互斥锁，默认为 false。
# <翻译结束>


<原文开始>
// Create creates and returns a new RWMutex object.
// The parameter `safe` is used to specify whether using this mutex in concurrent safety,
// which is false in default.
<原文结束>

# <翻译开始>
// 创建并返回一个新的RWMutex对象。
// 参数`safe`用于指定是否在并发安全的情况下使用此互斥锁，默认为false。
# <翻译结束>


<原文开始>
// IsSafe checks and returns whether current mutex is in concurrent-safe usage.
<原文结束>

# <翻译开始>
// IsSafe 检查并返回当前互斥锁是否处于线程安全使用状态。
# <翻译结束>


<原文开始>
// Lock locks mutex for writing.
// It does nothing if it is not in concurrent-safe usage.
<原文结束>

# <翻译开始>
// Lock 用于对 mutex 进行写锁定。
// 如果不在并发安全使用场景下，该方法将不做任何操作。
# <翻译结束>


<原文开始>
// Unlock unlocks mutex for writing.
// It does nothing if it is not in concurrent-safe usage.
<原文结束>

# <翻译开始>
// Unlock用于解锁互斥锁以供写入。
// 如果未在并发安全的使用场景下，此操作将无任何效果。
# <翻译结束>


<原文开始>
// RLock locks mutex for reading.
// It does nothing if it is not in concurrent-safe usage.
<原文结束>

# <翻译开始>
// RLock 用于读取时锁定互斥锁。
// 如果未处于并发安全使用状态，则此操作不做任何事情。
# <翻译结束>


<原文开始>
// RUnlock unlocks mutex for reading.
// It does nothing if it is not in concurrent-safe usage.
<原文结束>

# <翻译开始>
// RUnlock 用于读取解锁。
// 如果未在并发安全的使用场景下，该方法将不做任何操作。
# <翻译结束>

