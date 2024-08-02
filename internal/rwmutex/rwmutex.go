// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包 rwmutex 为 sync.RWMutex 提供并发安全特性切换的功能。 md5:563f53220ab3eec8
package rwmutex

import (
	"sync"
)

// RWMutex 是一个具有并发安全开关的 sync.RWMutex。
// 如果其 sync.RWMutex 类型的属性非空，表示它处于并发安全使用中。
// 默认情况下，它的 sync.RWMutex 属性为 nil，这使得该结构体更为轻量。
// md5:2d8d597983a75c36
type RWMutex struct {
	// Underlying mutex.
	mutex *sync.RWMutex
}

// New 创建并返回一个新的 RWMutex 实例。
// 参数 `safe` 用于指定是否在并发环境中使用这个互斥锁，默认为 false，表示不安全。
// md5:e431e613f230b125
func New(safe ...bool) *RWMutex {
	mu := Create(safe...)
	return &mu
}

// Create 创建并返回一个新的 RWMutex 对象。
// 参数 `safe` 用于指定是否在并发安全模式下使用该互斥锁，其默认值为 false。
// md5:e40df278667779d2
func Create(safe ...bool) RWMutex {
	if len(safe) > 0 && safe[0] {
		return RWMutex{
			mutex: new(sync.RWMutex),
		}
	}
	return RWMutex{}
}

// IsSafe 检查并返回当前互斥锁是否在并发安全的使用中。 md5:1a2c4197eb3278b5
func (mu *RWMutex) IsSafe() bool {
	return mu.mutex != nil
}

// Lock 为写入锁定互斥量。如果没有进行并发安全使用，它不会做任何事情。
// md5:e7a0e420dc8d74c3
func (mu *RWMutex) Lock() {
	if mu.mutex != nil {
		mu.mutex.Lock()
	}
}

// Unlock 为写操作解锁互斥锁。如果它不是在并发安全模式下使用，则不会做任何事情。
// md5:ce0b3215f968f29c
func (mu *RWMutex) Unlock() {
	if mu.mutex != nil {
		mu.mutex.Unlock()
	}
}

// RLock 用于对互斥锁进行读取锁定。
// 如果不是在并发安全的使用场景下，它不做任何操作。
// md5:61160c78e9bcccd5
func (mu *RWMutex) RLock() {
	if mu.mutex != nil {
		mu.mutex.RLock()
	}
}

// RUnlock 释放读取锁。
// 如果在非并发安全使用时，它将不执行任何操作。
// md5:834672a97d0bd47f
func (mu *RWMutex) RUnlock() {
	if mu.mutex != nil {
		mu.mutex.RUnlock()
	}
}
