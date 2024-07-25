// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 包mutex为sync.Mutex提供了并发安全开关的功能。 md5:2d280f557de1d7a8
package mutex

import (
	"sync"
)

// Mutex是一个带有并发安全功能开关的sync.Mutex。 md5:8889db913bd4aa9f
type Mutex struct {
	// Underlying mutex.
	mutex *sync.Mutex
}

// New 创建并返回一个新的 *Mutex。
// 参数 `safe` 用于指定是否在并发安全情况下使用此互斥锁，其默认值为 false。 md5:4b9e38d55d8b7828
func New(safe ...bool) *Mutex {
	mu := Create(safe...)
	return &mu
}

// Create 创建并返回一个新的 Mutex 对象。
// 参数 `safe` 用于指定是否在并发安全的环境下使用此 mutex，默认为 false。 md5:a3db2fe6cfb0197f
func Create(safe ...bool) Mutex {
	if len(safe) > 0 && safe[0] {
		return Mutex{
			mutex: new(sync.Mutex),
		}
	}
	return Mutex{}
}

// IsSafe 检查并返回当前互斥锁是否在并发安全的使用中。 md5:1a2c4197eb3278b5
func (mu *Mutex) IsSafe() bool {
	return mu.mutex != nil
}

// Lock 为写入锁定互斥量。如果没有进行并发安全使用，它不会做任何事情。 md5:e7a0e420dc8d74c3
func (mu *Mutex) Lock() {
	if mu.mutex != nil {
		mu.mutex.Lock()
	}
}

// Unlock 为写操作解锁互斥锁。如果它不是在并发安全模式下使用，则不会做任何事情。 md5:ce0b3215f968f29c
func (mu *Mutex) Unlock() {
	if mu.mutex != nil {
		mu.mutex.Unlock()
	}
}
