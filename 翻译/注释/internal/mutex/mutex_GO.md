
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
// Package mutex provides switch of concurrent safe feature for sync.Mutex.
<原文结束>

# <翻译开始>
// Package mutex 提供了对 sync.Mutex 的并发安全特性的开关控制。
# <翻译结束>


<原文开始>
// Mutex is a sync.Mutex with a switch for concurrent safe feature.
<原文结束>

# <翻译开始>
// Mutex 是一个带并发安全开关的 sync.Mutex，用于在并发场景中保证安全。
# <翻译结束>







<原文开始>
// New creates and returns a new *Mutex.
// The parameter `safe` is used to specify whether using this mutex in concurrent safety,
// which is false in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的 *Mutex。
// 参数 `safe` 用于指定是否在并发安全的情况下使用此互斥锁，默认为 false。
# <翻译结束>


<原文开始>
// Create creates and returns a new Mutex object.
// The parameter `safe` is used to specify whether using this mutex in concurrent safety,
// which is false in default.
<原文结束>

# <翻译开始>
// 创建并返回一个新的Mutex对象。
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

