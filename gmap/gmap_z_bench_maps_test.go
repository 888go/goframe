// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package map类_test

import (
	"testing"
	
	"github.com/888go/goframe/gmap"
	"github.com/gogf/gf/v2/util/gutil"
)

var hashMap = map类.X创建(true)

var listMap = map类.X创建链表mp(true)

var treeMap = map类.X创建红黑树Map(gutil.ComparatorInt, true)

func Benchmark_HashMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			hashMap.X设置值(i, i)
			i++
		}
	})
}

func Benchmark_ListMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			listMap.X设置值(i, i)
			i++
		}
	})
}

func Benchmark_TreeMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			treeMap.Set(i, i)
			i++
		}
	})
}

func Benchmark_HashMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			hashMap.X取值(i)
			i++
		}
	})
}

func Benchmark_ListMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			listMap.X取值(i)
			i++
		}
	})
}

func Benchmark_TreeMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			treeMap.Get(i)
			i++
		}
	})
}
