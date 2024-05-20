// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。. md5:81db3d7bd1ed4da8

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
