
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
// Locker is a memory based locker.
// Note that there's no cache expire mechanism for mutex in locker.
// You need remove certain mutex manually when you do not want use it anymore.
<原文结束>

# <翻译开始>
// Locker 是一个基于内存的锁。
// 注意，Lockers 中的 mutex 没有缓存过期机制。
// 当你不再需要使用某个锁时，需要手动删除它。
// md5:330f85347bba3cc8
# <翻译结束>


<原文开始>
// New creates and returns a new memory locker.
// A memory locker can lock/unlock with dynamic string key.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的内存锁。
// 这个内存锁能够使用动态字符串键进行锁定和解锁。
// md5:37aaba9921b3e711
# <翻译结束>


<原文开始>
// Lock locks the `key` with writing lock.
// If there's a write/reading lock the `key`,
// it will block until the lock is released.
<原文结束>

# <翻译开始>
// Lock 以写锁方式锁定`key`。
// 如果`key`已有写锁或读锁，它会阻塞直到锁被释放。
// md5:7b2d56ac41ec0a40
# <翻译结束>


<原文开始>
// TryLock tries locking the `key` with writing lock,
// it returns true if success, or it returns false if there's a writing/reading lock the `key`.
<原文结束>

# <翻译开始>
// TryLock尝试使用写入锁锁定`key`，如果成功返回true，如果`key`已经有写入或读取锁则返回false。
// md5:1e86a7888ed1621a
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
// it will block until the lock is released.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// LockFunc 使用写入锁锁定`key`，并使用回调函数`f`。
// 如果`key`已有写入或读取锁，它将阻塞直到锁被释放。
//
// 在执行完`f`后，它会释放锁。
// md5:fc66c542fa813208
# <翻译结束>


<原文开始>
// RLockFunc locks the `key` with reading lock and callback function `f`.
// If there's a writing lock the `key`,
// it will block until the lock is released.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// RLockFunc 使用读取锁对`key`进行加锁，并执行回调函数`f`。
// 如果`key`已被写入锁锁定，
// 则会阻塞直到该锁被释放。
//
// 在`f`执行完毕后，它将释放锁。
// md5:3f30fb5d911cd5e7
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
// Remove removes mutex with given `key` from locker.
<原文结束>

# <翻译开始>
// Remove 根据给定的 `key` 从 locker 中移除互斥锁。. md5:e557087320b6a672
# <翻译结束>


<原文开始>
// Clear removes all mutexes from locker.
<原文结束>

# <翻译开始>
// Clear 从 locker 中移除所有互斥锁。. md5:6e7b8ead4ad69f9d
# <翻译结束>


<原文开始>
// getOrNewMutex returns the mutex of given `key` if it exists,
// or else creates and returns a new one.
<原文结束>

# <翻译开始>
// getOrNewMutex 如果给定的`key`存在，则返回对应的互斥锁，否则创建一个新的并返回。
// md5:08c35eff58386554
# <翻译结束>

