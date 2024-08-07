// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 内存锁类_test

import (
	"testing"

	gmlock "github.com/888go/goframe/os/gmlock"
)

var (
	lockKey = "This is the lock key for gmlock."
)

func Benchmark_GMLock_Lock_Unlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gmlock.X写锁定(lockKey)
		gmlock.X退出写锁定(lockKey)
	}
}

func Benchmark_GMLock_RLock_RUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gmlock.X读锁定(lockKey)
		gmlock.X退出读锁定(lockKey)
	}
}

func Benchmark_GMLock_TryLock_Unlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if gmlock.X非阻塞写锁定(lockKey) {
			gmlock.X退出写锁定(lockKey)
		}
	}
}

func Benchmark_GMLock_TryRLock_RUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if gmlock.X非阻塞读锁定(lockKey) {
			gmlock.X退出读锁定(lockKey)
		}
	}
}
