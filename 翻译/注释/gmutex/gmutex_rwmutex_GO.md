
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
// RWMutex is a high level RWMutex, which implements more rich features for mutex.
<原文结束>

# <翻译开始>
// RWMutex 是一种高级别的读写互斥锁，它为 mutex 实现了更多丰富的功能。
# <翻译结束>


<原文开始>
// LockFunc locks the mutex for writing with given callback function `f`.
// If there's a write/reading lock the mutex, it will block until the lock is released.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// LockFunc 使用给定的回调函数`f`对互斥锁进行写入锁定。
// 如果有其他goroutine正在对互斥锁进行写入或读取锁定，该函数将阻塞，直到锁被释放。
//
// 在`f`执行完毕后，它会自动释放锁。
# <翻译结束>


<原文开始>
// RLockFunc locks the mutex for reading with given callback function `f`.
// If there's a writing lock the mutex, it will block until the lock is released.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// RLockFunc 以给定的回调函数`f`对互斥锁进行读取锁定。
// 如果存在写入锁定的互斥锁，它将阻塞直到该锁被释放。
//
// 在`f`执行完毕后，它会自动释放锁。
# <翻译结束>


<原文开始>
// TryLockFunc tries locking the mutex for writing with given callback function `f`.
// it returns true immediately if success, or if there's a write/reading lock on the mutex,
// it returns false immediately.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// TryLockFunc 尝试以给定回调函数 `f` 对互斥锁进行写入锁定。
// 如果成功，则立即返回 true；如果互斥锁上存在写入/读取锁，它会立即返回 false。
//
// 在 `f` 执行完毕后释放锁。
# <翻译结束>


<原文开始>
// TryRLockFunc tries locking the mutex for reading with given callback function `f`.
// It returns true immediately if success, or if there's a writing lock on the mutex,
// it returns false immediately.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// TryRLockFunc 尝试以读模式锁定互斥锁并执行给定的回调函数 `f`。
// 若成功锁定，则立即返回 true；若互斥锁当前正被写模式锁定，则立即返回 false。
//
// 在 `f` 执行完毕后，它会自动释放该锁。
# <翻译结束>

