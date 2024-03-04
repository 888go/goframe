
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
// Package gmlock implements a concurrent-safe memory-based locker.
<原文结束>

# <翻译开始>
// Package gmlock 实现了一个基于内存的并发安全锁。
# <翻译结束>


<原文开始>
// Lock locks the `key` with writing lock.
// If there's a write/reading lock the `key`,
// it will blocks until the lock is released.
<原文结束>

# <翻译开始>
// Lock 对`key`使用写入锁进行加锁。
// 如果有对`key`的写入/读取锁存在，
// 它将阻塞直到锁被释放。
# <翻译结束>


<原文开始>
// TryLock tries locking the `key` with writing lock,
// it returns true if success, or if there's a write/reading lock the `key`,
// it returns false.
<原文结束>

# <翻译开始>
// TryLock 尝试对`key`进行写锁锁定，
// 如果成功则返回 true，如果`key`存在写锁或读锁，则返回 false。
# <翻译结束>


<原文开始>
// Unlock unlocks the writing lock of the `key`.
<原文结束>

# <翻译开始>
// Unlock 解锁指定`key`的写入锁。
# <翻译结束>


<原文开始>
// RLock locks the `key` with reading lock.
// If there's a writing lock on `key`,
// it will blocks until the writing lock is released.
<原文结束>

# <翻译开始>
// RLock 对“key”加读锁。
// 如果“key”上存在写锁，
// 它将阻塞直到写锁被释放。
# <翻译结束>


<原文开始>
// TryRLock tries locking the `key` with reading lock.
// It returns true if success, or if there's a writing lock on `key`, it returns false.
<原文结束>

# <翻译开始>
// TryRLock 尝试对`key`使用读锁进行加锁。
// 如果加锁成功，返回true；如果`key`上存在写锁，则返回false。
# <翻译结束>


<原文开始>
// RUnlock unlocks the reading lock of the `key`.
<原文结束>

# <翻译开始>
// RUnlock 解除对`key`的读取锁。
# <翻译结束>


<原文开始>
// LockFunc locks the `key` with writing lock and callback function `f`.
// If there's a write/reading lock the `key`,
// it will blocks until the lock is released.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// LockFunc 对`key`使用写锁并调用回调函数`f`进行锁定。
// 如果有对`key`的写锁/读锁存在，
// 它将阻塞直到锁被释放。
//
// 在`f`执行完毕后，它会自动释放锁。
# <翻译结束>


<原文开始>
// RLockFunc locks the `key` with reading lock and callback function `f`.
// If there's a writing lock the `key`,
// it will blocks until the lock is released.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// RLockFunc 对 `key` 加上读锁并执行回调函数 `f`。
// 如果存在对 `key` 的写锁，则会阻塞直到该锁被释放。
//
// 在 `f` 执行完毕后，它将自动释放锁。
# <翻译结束>


<原文开始>
// TryLockFunc locks the `key` with writing lock and callback function `f`.
// It returns true if success, or else if there's a write/reading lock the `key`, it return false.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// TryLockFunc 尝试对`key`使用写锁并执行回调函数`f`。
// 如果成功获取锁，则返回 true，否则如果`key`已有写锁或读锁存在，则返回 false。
//
// 在`f`执行完毕后会自动释放锁。
# <翻译结束>


<原文开始>
// TryRLockFunc locks the `key` with reading lock and callback function `f`.
// It returns true if success, or else if there's a writing lock the `key`, it returns false.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// TryRLockFunc 尝试对`key`加读锁并执行回调函数`f`。
// 如果成功获取读锁则返回true，否则如果`key`已有写锁存在，则返回false。
//
// 在`f`执行完毕后会自动释放锁。
# <翻译结束>


<原文开始>
// Remove removes mutex with given `key`.
<原文结束>

# <翻译开始>
// Remove 通过给定的`key`移除互斥锁。
# <翻译结束>






