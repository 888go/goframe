// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gmlock

import (
	"sync"
	
	"github.com/888go/goframe/container/gmap"
)

// Locker是一个基于内存的锁。
// 注意，locker中的互斥锁没有缓存过期机制。
// 当您不再需要使用某个互斥锁时，需要手动移除它。
type Locker struct {
	m *gmap.StrAnyMap
}

// New 创建并返回一个新的内存锁器。
// 内存锁器可以使用动态字符串键进行锁定和解锁。
func New() *Locker {
	return &Locker{
		m: gmap.NewStrAnyMap(true),
	}
}

// Lock 用于对`key`进行写锁锁定。
// 如果`key`当前存在写锁或读锁，
// 则该操作将会阻塞，直到锁被释放。
func (l *Locker) Lock(key string) {
	l.getOrNewMutex(key).Lock()
}

// TryLock 尝试以写锁方式锁定`key`，
// 如果成功则返回 true，如果`key`已有写锁或读锁，则返回 false。
func (l *Locker) TryLock(key string) bool {
	return l.getOrNewMutex(key).TryLock()
}

// Unlock 解锁指定`key`的写入锁。
func (l *Locker) Unlock(key string) {
	if v := l.m.Get(key); v != nil {
		v.(*sync.RWMutex).Unlock()
	}
}

// RLock 对“key”加读锁。
// 如果“key”上存在写锁，
// 它将阻塞直到写锁被释放。
func (l *Locker) RLock(key string) {
	l.getOrNewMutex(key).RLock()
}

// TryRLock 尝试对`key`使用读锁进行加锁。
// 如果加锁成功，返回true；如果`key`上存在写锁，则返回false。
func (l *Locker) TryRLock(key string) bool {
	return l.getOrNewMutex(key).TryRLock()
}

// RUnlock 解除对`key`的读取锁。
func (l *Locker) RUnlock(key string) {
	if v := l.m.Get(key); v != nil {
		v.(*sync.RWMutex).RUnlock()
	}
}

// LockFunc 对`key`使用写锁并调用回调函数`f`进行锁定。
// 如果`key`已经有写锁或读锁存在，将会阻塞直到锁被释放。
//
// 在`f`执行完毕后会自动释放锁。
func (l *Locker) LockFunc(key string, f func()) {
	l.Lock(key)
	defer l.Unlock(key)
	f()
}

// RLockFunc 对 `key` 加上读锁并执行回调函数 `f`。
// 如果 `key` 上存在写锁，则会阻塞直到该锁被释放。
//
// 在 `f` 执行完毕后，它将自动释放锁。
func (l *Locker) RLockFunc(key string, f func()) {
	l.RLock(key)
	defer l.RUnlock(key)
	f()
}

// TryLockFunc 尝试对`key`使用写锁并执行回调函数`f`。
// 如果成功获取锁，则返回 true，否则如果`key`已有写锁或读锁存在，则返回 false。
//
// 在`f`执行完毕后会自动释放锁。
func (l *Locker) TryLockFunc(key string, f func()) bool {
	if l.TryLock(key) {
		defer l.Unlock(key)
		f()
		return true
	}
	return false
}

// TryRLockFunc 尝试对`key`加读锁并执行回调函数`f`。
// 如果成功获取读锁则返回true，否则如果`key`已有写锁存在，则返回false。
//
// 在`f`执行完毕后会自动释放锁。
func (l *Locker) TryRLockFunc(key string, f func()) bool {
	if l.TryRLock(key) {
		defer l.RUnlock(key)
		f()
		return true
	}
	return false
}

// Remove 从 locker 中移除具有给定 `key` 的互斥锁。
func (l *Locker) Remove(key string) {
	l.m.Remove(key)
}

// Clear 从 locker 中移除所有互斥锁。
func (l *Locker) Clear() {
	l.m.Clear()
}

// getOrNewMutex根据给定的`key`返回其关联的互斥锁，
// 如果该互斥锁不存在，则创建并返回一个新的互斥锁。
func (l *Locker) getOrNewMutex(key string) *sync.RWMutex {
	return l.m.GetOrSetFuncLock(key, func() interface{} {
		return &sync.RWMutex{}
	}).(*sync.RWMutex)
}
