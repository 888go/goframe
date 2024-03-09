// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 内存锁类_test

import (
	"testing"
	
	"github.com/888go/goframe/gmlock"
)

var (
	lockKey = "This is the lock key for gmlock."
)

func Benchmark_GMLock_Lock_Unlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		内存锁类.X写锁定(lockKey)
		内存锁类.X退出写锁定(lockKey)
	}
}

func Benchmark_GMLock_RLock_RUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		内存锁类.X读锁定(lockKey)
		内存锁类.X退出读锁定(lockKey)
	}
}

func Benchmark_GMLock_TryLock_Unlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if 内存锁类.X非阻塞写锁定(lockKey) {
			内存锁类.X退出写锁定(lockKey)
		}
	}
}

func Benchmark_GMLock_TryRLock_RUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if 内存锁类.X非阻塞读锁定(lockKey) {
			内存锁类.X退出读锁定(lockKey)
		}
	}
}
