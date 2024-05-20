
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
// Package gmlock implements a concurrent-safe memory-based locker.
<原文结束>

# <翻译开始>
// 包gmlock实现了一个基于内存的并发安全锁。. md5:46d1d349c220f670
# <翻译结束>


<原文开始>
// Lock locks the `key` with writing lock.
// If there's a write/reading lock the `key`,
// it will blocks until the lock is released.
<原文结束>

# <翻译开始>
// Lock 以写锁方式锁定 `key`。
// 如果有写锁或读锁正在锁定 `key`，
// 它将阻塞直到锁被释放。
// md5:8dcc0b1e059e3831
# <翻译结束>


<原文开始>
// TryLock tries locking the `key` with writing lock,
// it returns true if success, or if there's a write/reading lock the `key`,
// it returns false.
<原文结束>

# <翻译开始>
// TryLock 尝试对`key`加写锁，如果成功则返回true。如果`key`已有写锁或读锁，则返回false。
// md5:e4f172d7beca094d
# <翻译结束>


<原文开始>
// Unlock unlocks the writing lock of the `key`.
<原文结束>

# <翻译开始>
// Unlock 解锁对`key`的写入锁。. md5:b54a3ae386cfa500
# <翻译结束>


<原文开始>
// RLock locks the `key` with reading lock.
// If there's a writing lock on `key`,
// it will blocks until the writing lock is released.
<原文结束>

# <翻译开始>
// RLock 使用读锁锁定`key`。如果`key`上有一个写锁，它将阻塞直到写锁被释放。
// md5:f45f660f368bbb78
# <翻译结束>


<原文开始>
// TryRLock tries locking the `key` with reading lock.
// It returns true if success, or if there's a writing lock on `key`, it returns false.
<原文结束>

# <翻译开始>
// TryRLock 尝试使用读取锁对 `key` 进行加锁。
// 如果成功，则返回true；如果 `key` 上存在写入锁，则返回false。
// md5:8733aa161c104b87
# <翻译结束>


<原文开始>
// RUnlock unlocks the reading lock of the `key`.
<原文结束>

# <翻译开始>
// RUnlock 释放对 `key` 的读取锁。. md5:d4f823abaa858783
# <翻译结束>


<原文开始>
// LockFunc locks the `key` with writing lock and callback function `f`.
// If there's a write/reading lock the `key`,
// it will blocks until the lock is released.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// LockFunc 使用写入锁锁定`key`，并使用回调函数`f`。如果`key`已有写入或读取锁，它将阻塞直到锁被释放。
// 
// 在`f`执行后，它会释放锁。
// md5:3e35c1977b58dac3
# <翻译结束>


<原文开始>
// RLockFunc locks the `key` with reading lock and callback function `f`.
// If there's a writing lock the `key`,
// it will blocks until the lock is released.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// RLockFunc 使用读取锁对 `key` 进行锁定，并调用回调函数 `f`。
// 如果 `key` 上有写入锁，它将阻塞直到锁被释放。
//
// 在 `f` 执行完毕后，它会释放锁。
// md5:e4a03ce9029d1911
# <翻译结束>


<原文开始>
// TryLockFunc locks the `key` with writing lock and callback function `f`.
// It returns true if success, or else if there's a write/reading lock the `key`, it return false.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// TryLockFunc 使用写锁锁定`key`并执行回调函数`f`。
// 如果操作成功，返回true；如果`key`已被写锁或读锁占用，则返回false。
//
// 在回调函数`f`执行完毕后，它会释放锁。
// md5:a016db0c6b2bc67e
# <翻译结束>


<原文开始>
// TryRLockFunc locks the `key` with reading lock and callback function `f`.
// It returns true if success, or else if there's a writing lock the `key`, it returns false.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// TryRLockFunc 使用读取锁尝试锁定`key`，并调用回调函数`f`。
// 如果成功，它返回true；如果`key`已有写入锁，则返回false。
//
// 在`f`执行完毕后释放锁。
// md5:527ef8bb470bd8fd
# <翻译结束>


<原文开始>
// Remove removes mutex with given `key`.
<原文结束>

# <翻译开始>
// Remove 删除具有给定`key`的互斥锁。. md5:a8a4db10705ecf7a
# <翻译结束>

