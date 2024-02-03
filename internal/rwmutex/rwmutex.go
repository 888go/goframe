// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package rwmutex 为 sync.RWMutex 提供了并发安全特性开关。
package rwmutex

import (
	"sync"
)

// RWMutex 是一个带并发安全开关的 sync.RWMutex。
// 如果其属性 *sync.RWMutex 非空，表示它处于并发安全使用状态。
// 默认情况下，其属性 *sync.RWMutex 为空（nil），这使得该结构体极为轻量级。
// 这段代码的注释翻译成中文如下：
// ```go
// RWMutex 是对标准库sync.RWMutex的一种扩展，增加了并发安全特性切换功能。
// 若其成员变量 *sync.RWMutex 不为 nil，表明当前它正在以支持并发安全的方式使用。
// 默认情况下，其成员变量 *sync.RWMutex 初始化为 nil，这样的设计使得该结构体保持轻量。
type RWMutex struct {
	// Underlying mutex.
	mutex *sync.RWMutex
}

// New 创建并返回一个新的 *RWMutex。
// 参数 `safe` 用于指定是否在并发安全的情况下使用此互斥锁，默认为 false。
func New(safe ...bool) *RWMutex {
	mu := Create(safe...)
	return &mu
}

// 创建并返回一个新的RWMutex对象。
// 参数`safe`用于指定是否在并发安全的情况下使用此互斥锁，默认为false。
func Create(safe ...bool) RWMutex {
	if len(safe) > 0 && safe[0] {
		return RWMutex{
			mutex: new(sync.RWMutex),
		}
	}
	return RWMutex{}
}

// IsSafe 检查并返回当前互斥锁是否处于线程安全使用状态。
func (mu *RWMutex) IsSafe() bool {
	return mu.mutex != nil
}

// Lock 用于对 mutex 进行写锁定。
// 如果不在并发安全使用场景下，该方法将不做任何操作。
func (mu *RWMutex) Lock() {
	if mu.mutex != nil {
		mu.mutex.Lock()
	}
}

// Unlock用于解锁互斥锁以供写入。
// 如果未在并发安全的使用场景下，此操作将无任何效果。
func (mu *RWMutex) Unlock() {
	if mu.mutex != nil {
		mu.mutex.Unlock()
	}
}

// RLock 用于读取时锁定互斥锁。
// 如果未处于并发安全使用状态，则此操作不做任何事情。
func (mu *RWMutex) RLock() {
	if mu.mutex != nil {
		mu.mutex.RLock()
	}
}

// RUnlock 用于读取解锁。
// 如果未在并发安全的使用场景下，该方法将不做任何操作。
func (mu *RWMutex) RUnlock() {
	if mu.mutex != nil {
		mu.mutex.RUnlock()
	}
}
