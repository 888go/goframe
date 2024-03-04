// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gmutex_test

import (
	"sync"
	"testing"
	
	"github.com/888go/goframe/gmutex"
)

var (
	mu   = sync.Mutex{}
	rwmu = sync.RWMutex{}
	gmu  = gmutex.New()
)

func Benchmark_Mutex_LockUnlock(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			mu.Unlock()
		}
	})
}

func Benchmark_RWMutex_LockUnlock(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rwmu.Lock()
			rwmu.Unlock()
		}
	})
}

func Benchmark_RWMutex_RLockRUnlock(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rwmu.RLock()
			rwmu.RUnlock()
		}
	})
}

func Benchmark_GMutex_LockUnlock(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gmu.Lock()
			gmu.Unlock()
		}
	})
}

func Benchmark_GMutex_TryLock(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if gmu.TryLock() {
				gmu.Unlock()
			}
		}
	})
}

func Benchmark_GMutex_RLockRUnlock(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gmu.RLock()
			gmu.RUnlock()
		}
	})
}

func Benchmark_GMutex_TryRLock(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if gmu.TryRLock() {
				gmu.RUnlock()
			}
		}
	})
}
