
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
// RWMutex is a high level RWMutex, which implements more rich features for mutex.
<原文结束>

# <翻译开始>
// RWMutex 是一个高级的 RWMutex，它为互斥锁提供了更多丰富的功能。. md5:25a009f424184cca
# <翻译结束>


<原文开始>
// LockFunc locks the mutex for writing with given callback function `f`.
// If there's a write/reading lock the mutex, it will block until the lock is released.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// LockFunc 使用给定的回调函数 `f` 对互斥锁进行写入锁定。
// 如果已经有写入或读取锁持有互斥锁，它将阻塞直到锁被释放。
//
// 在 `f` 执行完毕后，它会释放锁。
// md5:946a127ed090616d
# <翻译结束>


<原文开始>
// RLockFunc locks the mutex for reading with given callback function `f`.
// If there's a writing lock the mutex, it will block until the lock is released.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// RLockFunc 使用给定的回调函数`f`对互斥锁进行读取锁定。
// 如果有写入锁正在锁定互斥锁，它将阻塞直到锁被释放。
//
// 在执行完`f`后，它会释放锁。
// md5:e97751e4a44accd4
# <翻译结束>


<原文开始>
// TryLockFunc tries locking the mutex for writing with given callback function `f`.
// it returns true immediately if success, or if there's a write/reading lock on the mutex,
// it returns false immediately.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// TryLockFunc尝试使用给定的回调函数`f`为写入锁定mutex。如果成功，它会立即返回true，或者如果mutex已经有写入或读取锁，它会立即返回false。
// 
// 在执行完`f`后，它会释放锁。
// md5:d12ccf3fb040146e
# <翻译结束>


<原文开始>
// TryRLockFunc tries locking the mutex for reading with given callback function `f`.
// It returns true immediately if success, or if there's a writing lock on the mutex,
// it returns false immediately.
//
// It releases the lock after `f` is executed.
<原文结束>

# <翻译开始>
// TryRLockFunc尝试使用给定的回调函数`f`获取读取锁。
// 如果成功，它会立即返回true。如果互斥锁已被写入锁定，
// 它会立即返回false。
//
// 在`f`执行后释放锁。
// md5:138728820d3bbfa0
# <翻译结束>

