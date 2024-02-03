// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package glist

import (
	"testing"
)

var (
	l = New(true)
)

func Benchmark_PushBack(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			l.PushBack(i)
			i++
		}
	})
}

func Benchmark_PushFront(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			l.PushFront(i)
			i++
		}
	})
}

func Benchmark_Len(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Len()
		}
	})
}

func Benchmark_PopFront(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.PopFront()
		}
	})
}

func Benchmark_PopBack(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.PopBack()
		}
	})
}
