// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gmutex

import "sync"

// RWMutex 是一个高级的 RWMutex，它为互斥锁提供了更多丰富的功能。. md5:25a009f424184cca
type RWMutex struct {
	sync.RWMutex
}

// LockFunc 使用给定的回调函数 `f` 对互斥锁进行写入锁定。
// 如果已经有写入或读取锁持有互斥锁，它将阻塞直到锁被释放。
//
// 在 `f` 执行完毕后，它会释放锁。
// md5:946a127ed090616d
func (m *RWMutex) LockFunc(f func()) {
	m.Lock()
	defer m.Unlock()
	f()
}

// RLockFunc 使用给定的回调函数`f`对互斥锁进行读取锁定。
// 如果有写入锁正在锁定互斥锁，它将阻塞直到锁被释放。
//
// 在执行完`f`后，它会释放锁。
// md5:e97751e4a44accd4
func (m *RWMutex) RLockFunc(f func()) {
	m.RLock()
	defer m.RUnlock()
	f()
}

// TryLockFunc尝试使用给定的回调函数`f`为写入锁定mutex。如果成功，它会立即返回true，或者如果mutex已经有写入或读取锁，它会立即返回false。
// 
// 在执行完`f`后，它会释放锁。
// md5:d12ccf3fb040146e
func (m *RWMutex) TryLockFunc(f func()) (result bool) {
	if m.TryLock() {
		result = true
		defer m.Unlock()
		f()
	}
	return
}

// TryRLockFunc尝试使用给定的回调函数`f`获取读取锁。
// 如果成功，它会立即返回true。如果互斥锁已被写入锁定，
// 它会立即返回false。
//
// 在`f`执行后释放锁。
// md5:138728820d3bbfa0
func (m *RWMutex) TryRLockFunc(f func()) (result bool) {
	if m.TryRLock() {
		result = true
		defer m.RUnlock()
		f()
	}
	return
}
