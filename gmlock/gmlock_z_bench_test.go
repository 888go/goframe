// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gmlock_test

import (
	"testing"
	
	"github.com/888go/goframe/gmlock"
)

var (
	lockKey = "This is the lock key for gmlock."
)

func Benchmark_GMLock_Lock_Unlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gmlock.Lock(lockKey)
		gmlock.Unlock(lockKey)
	}
}

func Benchmark_GMLock_RLock_RUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gmlock.RLock(lockKey)
		gmlock.RUnlock(lockKey)
	}
}

func Benchmark_GMLock_TryLock_Unlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if gmlock.TryLock(lockKey) {
			gmlock.Unlock(lockKey)
		}
	}
}

func Benchmark_GMLock_TryRLock_RUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if gmlock.TryRLock(lockKey) {
			gmlock.RUnlock(lockKey)
		}
	}
}
