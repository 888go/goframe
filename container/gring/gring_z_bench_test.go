// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package 循环链表类_test

import (
	"testing"

	gring "github.com/888go/goframe/container/gring"
)

var length = 10000

var ringObject = gring.New(length, true)

func BenchmarkRing_Put(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			ringObject.Put(i)
			i++
		}
	})
}

func BenchmarkRing_Next(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			ringObject.Next()
			i++
		}
	})
}

func BenchmarkRing_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			ringObject.X设置值(i)
			i++
		}
	})
}

func BenchmarkRing_Len(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			ringObject.Len()
			i++
		}
	})
}

func BenchmarkRing_Cap(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			ringObject.Cap()
			i++
		}
	})
}
