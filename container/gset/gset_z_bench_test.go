// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package 集合类_test

import (
	"strconv"
	"testing"

	gset "github.com/888go/goframe/container/gset"
)

var intSet = gset.X创建整数(true)

var anySet = gset.NewSet别名(true)

var strSet = gset.X创建文本(true)

var intSetUnsafe = gset.X创建整数()

var anySetUnsafe = gset.NewSet别名()

var strSetUnsafe = gset.X创建文本()

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

// 注意，字符串转换会带来额外的性能开销。 md5:1a28b26d80944830
func Benchmark_StrSet_Add(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strSet.X加入(strconv.Itoa(i))
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。 md5:1a28b26d80944830
func Benchmark_StrSet_Contains(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strSet.X是否存在(strconv.Itoa(i))
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。 md5:1a28b26d80944830
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

// 注意，字符串转换会带来额外的性能开销。 md5:1a28b26d80944830
func Benchmark_Unsafe_StrSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strSetUnsafe.X加入(strconv.Itoa(i))
	}
}

// 注意，字符串转换会带来额外的性能开销。 md5:1a28b26d80944830
func Benchmark_Unsafe_StrSet_Contains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strSetUnsafe.X是否存在(strconv.Itoa(i))
	}
}

// 注意，字符串转换会带来额外的性能开销。 md5:1a28b26d80944830
func Benchmark_Unsafe_StrSet_Remove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strSetUnsafe.X删除(strconv.Itoa(i))
	}
}
