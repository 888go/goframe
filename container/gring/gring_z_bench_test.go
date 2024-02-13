// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 循环链表类_test

import (
	"testing"
	
	"github.com/888go/goframe/container/gring"
)

var length = 10000

var ringObject = 循环链表类.New(length, true)

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
