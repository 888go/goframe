// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gmutex


import (
	"sync"
	)
// RWMutex 是一种高级别的读写互斥锁，它为 mutex 实现了更多丰富的功能。
type RWMutex struct {
	sync.RWMutex
}

// LockFunc 使用给定的回调函数`f`对互斥锁进行写入锁定。
// 如果有其他goroutine正在对互斥锁进行写入或读取锁定，该函数将阻塞，直到锁被释放。
//
// 在`f`执行完毕后，它会自动释放锁。
func (m *RWMutex) LockFunc(f func()) {
	m.Lock()
	defer m.Unlock()
	f()
}

// RLockFunc 以给定的回调函数`f`对互斥锁进行读取锁定。
// 如果存在写入锁定的互斥锁，它将阻塞直到该锁被释放。
//
// 在`f`执行完毕后，它会自动释放锁。
func (m *RWMutex) RLockFunc(f func()) {
	m.RLock()
	defer m.RUnlock()
	f()
}

// TryLockFunc 尝试以给定回调函数 `f` 对互斥锁进行写入锁定。
// 如果成功，则立即返回 true；如果互斥锁上存在写入/读取锁，它会立即返回 false。
//
// 在 `f` 执行完毕后释放锁。
func (m *RWMutex) TryLockFunc(f func()) (result bool) {
	if m.TryLock() {
		result = true
		defer m.Unlock()
		f()
	}
	return
}

// TryRLockFunc 尝试以读模式锁定互斥锁并执行给定的回调函数 `f`。
// 若成功锁定，则立即返回 true；若互斥锁当前正被写模式锁定，则立即返回 false。
//
// 在 `f` 执行完毕后，它会自动释放该锁。
func (m *RWMutex) TryRLockFunc(f func()) (result bool) {
	if m.TryRLock() {
		result = true
		defer m.RUnlock()
		f()
	}
	return
}
