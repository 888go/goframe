// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package rwmutex_test

import (
	"testing"
	
	"github.com/888go/goframe/internal/rwmutex"
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
