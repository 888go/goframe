	// 版权归GoFrame作者(https:	//goframe.org)所有。保留所有权利。
	//
	// 本源代码形式受MIT许可证条款约束。
	// 如果未随本文件一同分发MIT许可证副本，
	// 您可以在https:	//github.com/gogf/gf处获取。
	// md5:a9832f33b234e3f3

package rwmutex_test

import (
	"testing"

	"github.com/gogf/gf/v2/internal/rwmutex"
)

var (
	safeLock   = rwmutex.New(true)
	unsafeLock = rwmutex.New(false)
)

func Benchmark_Safe_LockUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		safeLock.Lock()
		safeLock.Unlock()
	}
}

func Benchmark_Safe_RLockRUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		safeLock.RLock()
		safeLock.RUnlock()
	}
}

func Benchmark_UnSafe_LockUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unsafeLock.Lock()
		unsafeLock.Unlock()
	}
}

func Benchmark_UnSafe_RLockRUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unsafeLock.RLock()
		unsafeLock.RUnlock()
	}
}
