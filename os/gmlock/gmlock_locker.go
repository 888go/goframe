// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 内存锁类

import (
	"sync"

	gmap "github.com/888go/goframe/container/gmap"
)

// Locker 是一个基于内存的锁。
// 注意，Lockers 中的 mutex 没有缓存过期机制。
// 当你不再需要使用某个锁时，需要手动删除它。
// md5:330f85347bba3cc8
type Locker struct {
	m *gmap.StrAnyMap
}

// New 创建并返回一个新的内存锁。
// 这个内存锁能够使用动态字符串键进行锁定和解锁。
// md5:37aaba9921b3e711
func New() *Locker {
	return &Locker{
		m: gmap.NewStrAnyMap(true),
	}
}

// Lock 以写锁方式锁定`key`。
// 如果`key`已有写锁或读锁，它会阻塞直到锁被释放。
// md5:7b2d56ac41ec0a40
func (l *Locker) Lock(key string) {
	l.getOrNewMutex(key).Lock()
}

// TryLock尝试使用写入锁锁定`key`，如果成功返回true，如果`key`已经有写入或读取锁则返回false。
// md5:1e86a7888ed1621a
func (l *Locker) TryLock(key string) bool {
	return l.getOrNewMutex(key).TryLock()
}

// Unlock 解锁对`key`的写入锁。 md5:b54a3ae386cfa500
func (l *Locker) Unlock(key string) {
	if v := l.m.Get(key); v != nil {
		v.(*sync.RWMutex).Unlock()
	}
}

// RLock 使用读锁锁定`key`。如果`key`上有一个写锁，它将阻塞直到写锁被释放。
// md5:f45f660f368bbb78
func (l *Locker) RLock(key string) {
	l.getOrNewMutex(key).RLock()
}

// TryRLock 尝试使用读取锁对 `key` 进行加锁。
// 如果成功，则返回true；如果 `key` 上存在写入锁，则返回false。
// md5:8733aa161c104b87
func (l *Locker) TryRLock(key string) bool {
	return l.getOrNewMutex(key).TryRLock()
}

// RUnlock 释放对 `key` 的读取锁。 md5:d4f823abaa858783
func (l *Locker) RUnlock(key string) {
	if v := l.m.Get(key); v != nil {
		v.(*sync.RWMutex).RUnlock()
	}
}

// LockFunc 使用写入锁锁定`key`，并使用回调函数`f`。
// 如果`key`已有写入或读取锁，它将阻塞直到锁被释放。
//
// 在执行完`f`后，它会释放锁。
// md5:fc66c542fa813208
func (l *Locker) LockFunc(key string, f func()) {
	l.Lock(key)
	defer l.Unlock(key)
	f()
}

// RLockFunc 使用读取锁对`key`进行加锁，并执行回调函数`f`。
// 如果`key`已被写入锁锁定，
// 则会阻塞直到该锁被释放。
//
// 在`f`执行完毕后，它将释放锁。
// md5:3f30fb5d911cd5e7
func (l *Locker) RLockFunc(key string, f func()) {
	l.RLock(key)
	defer l.RUnlock(key)
	f()
}

// TryLockFunc 使用写锁锁定`key`并执行回调函数`f`。
// 如果操作成功，返回true；如果`key`已被写锁或读锁占用，则返回false。
//
// 在回调函数`f`执行完毕后，它会释放锁。
// md5:a016db0c6b2bc67e
func (l *Locker) TryLockFunc(key string, f func()) bool {
	if l.TryLock(key) {
		defer l.Unlock(key)
		f()
		return true
	}
	return false
}

// TryRLockFunc 使用读取锁尝试锁定`key`，并调用回调函数`f`。
// 如果成功，它返回true；如果`key`已有写入锁，则返回false。
//
// 在`f`执行完毕后释放锁。
// md5:527ef8bb470bd8fd
func (l *Locker) TryRLockFunc(key string, f func()) bool {
	if l.TryRLock(key) {
		defer l.RUnlock(key)
		f()
		return true
	}
	return false
}

// Remove 根据给定的 `key` 从 locker 中移除互斥锁。 md5:e557087320b6a672
func (l *Locker) Remove(key string) {
	l.m.Remove(key)
}

// Clear 从 locker 中移除所有互斥锁。 md5:6e7b8ead4ad69f9d
func (l *Locker) Clear() {
	l.m.Clear()
}

// getOrNewMutex 如果给定的`key`存在，则返回对应的互斥锁，否则创建一个新的并返回。
// md5:08c35eff58386554
func (l *Locker) getOrNewMutex(key string) *sync.RWMutex {
	return l.m.GetOrSetFuncLock(key, func() interface{} {
		return &sync.RWMutex{}
	}).(*sync.RWMutex)
}
