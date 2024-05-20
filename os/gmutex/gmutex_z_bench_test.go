// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gmutex_test

import (
	"sync"
	"testing"

	"github.com/gogf/gf/v2/os/gmutex"
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
