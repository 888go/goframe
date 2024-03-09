// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 互斥锁类

import (
	"sync"
)

// Mutex 是一个高级别的互斥锁，它为互斥锁实现了更多丰富的功能。
type X互斥锁 struct {
	sync.Mutex
}

// LockFunc 使用给定的回调函数`f`对互斥锁进行写入锁定。
// 如果有其他goroutine正在对互斥锁进行写入或读取锁定，该函数将阻塞，直到锁被释放。
//
// 在`f`执行完毕后，它会自动释放锁。
func (m *X互斥锁) X锁定_函数(回调函数 func()) {
	m.Lock()
	defer m.Unlock()
	回调函数()
}

// TryLockFunc 尝试以给定回调函数 `f` 对互斥锁进行写入锁定。
// 如果成功，则立即返回 true；如果互斥锁上存在写入/读取锁，它会立即返回 false。
//
// 在 `f` 执行完毕后释放锁。
func (m *X互斥锁) X非阻塞锁定_函数(回调函数 func()) (结果 bool) {
	if m.TryLock() {
		结果 = true
		defer m.Unlock()
		回调函数()
	}
	return
}
