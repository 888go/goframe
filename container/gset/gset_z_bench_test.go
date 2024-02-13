// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package 集合类_test

import (
	"strconv"
	"testing"
	
	"github.com/888go/goframe/container/gset"
)

var intSet = 集合类.X创建整数(true)

var anySet = 集合类.NewSet别名(true)

var strSet = 集合类.X创建文本(true)

var intSetUnsafe = 集合类.X创建整数()

var anySetUnsafe = 集合类.NewSet别名()

var strSetUnsafe = 集合类.X创建文本()

func Benchmark_IntSet_Add(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intSet.X加入(i)
			i++
		}
	})
}

func Benchmark_IntSet_Contains(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intSet.X是否存在(i)
			i++
		}
	})
}

func Benchmark_IntSet_Remove(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intSet.X删除(i)
			i++
		}
	})
}

func Benchmark_AnySet_Add(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			anySet.X加入(i)
			i++
		}
	})
}

func Benchmark_AnySet_Contains(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			anySet.X是否存在(i)
			i++
		}
	})
}

func Benchmark_AnySet_Remove(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			anySet.X删除(i)
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。
func Benchmark_StrSet_Add(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strSet.X加入(strconv.Itoa(i))
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。
func Benchmark_StrSet_Contains(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strSet.X是否存在(strconv.Itoa(i))
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。
func Benchmark_StrSet_Remove(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strSet.X删除(strconv.Itoa(i))
			i++
		}
	})
}

func Benchmark_Unsafe_IntSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSetUnsafe.X加入(i)
	}
}

func Benchmark_Unsafe_IntSet_Contains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSetUnsafe.X是否存在(i)
	}
}

func Benchmark_Unsafe_IntSet_Remove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSetUnsafe.X删除(i)
	}
}

func Benchmark_Unsafe_AnySet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anySetUnsafe.X加入(i)
	}
}

func Benchmark_Unsafe_AnySet_Contains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anySetUnsafe.X是否存在(i)
	}
}

func Benchmark_Unsafe_AnySet_Remove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anySetUnsafe.X删除(i)
	}
}

// 注意，字符串转换会带来额外的性能开销。
func Benchmark_Unsafe_StrSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strSetUnsafe.X加入(strconv.Itoa(i))
	}
}

// 注意，字符串转换会带来额外的性能开销。
func Benchmark_Unsafe_StrSet_Contains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strSetUnsafe.X是否存在(strconv.Itoa(i))
	}
}

// 注意，字符串转换会带来额外的性能开销。
func Benchmark_Unsafe_StrSet_Remove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strSetUnsafe.X删除(strconv.Itoa(i))
	}
}
