// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package mutex 提供了对 sync.Mutex 的并发安全特性的开关控制。
package mutex

import (
	"sync"
)

// Mutex 是一个带并发安全开关的 sync.Mutex，用于在并发场景中保证安全。
type Mutex struct {
	// Underlying mutex.
	mutex *sync.Mutex
}

// New 创建并返回一个新的 *Mutex。
// 参数 `safe` 用于指定是否在并发安全的情况下使用此互斥锁，默认为 false。
func New(safe ...bool) *Mutex {
	mu := Create(safe...)
	return &mu
}

// 创建并返回一个新的Mutex对象。
// 参数`safe`用于指定是否在并发安全的情况下使用此互斥锁，默认为false。
func Create(safe ...bool) Mutex {
	if len(safe) > 0 && safe[0] {
		return Mutex{
			mutex: new(sync.Mutex),
		}
	}
	return Mutex{}
}

// IsSafe 检查并返回当前互斥锁是否处于线程安全使用状态。
func (mu *Mutex) IsSafe() bool {
	return mu.mutex != nil
}

// Lock 用于对 mutex 进行写锁定。
// 如果不在并发安全使用场景下，该方法将不做任何操作。
func (mu *Mutex) Lock() {
	if mu.mutex != nil {
		mu.mutex.Lock()
	}
}

// Unlock用于解锁互斥锁以供写入。
// 如果未在并发安全的使用场景下，此操作将无任何效果。
func (mu *Mutex) Unlock() {
	if mu.mutex != nil {
		mu.mutex.Unlock()
	}
}
