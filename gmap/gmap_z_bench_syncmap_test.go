// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package map类_test

import (
	"sync"
	"testing"
	
	"github.com/888go/goframe/gmap"
)

var gm = map类.X创建IntInt(true)

var sm = sync.Map{}

func Benchmark_GMapSet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			gm.X设置值(i, i)
			i++
		}
	})
}

func Benchmark_SyncMapSet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			sm.Store(i, i)
			i++
		}
	})
}

func Benchmark_GMapGet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			gm.X取值(i)
			i++
		}
	})
}

func Benchmark_SyncMapGet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			sm.Load(i)
			i++
		}
	})
}

func Benchmark_GMapRemove(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			gm.X删除(i)
			i++
		}
	})
}

func Benchmark_SyncMapRmove(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			sm.Delete(i)
			i++
		}
	})
}
