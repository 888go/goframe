// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。. md5:81db3d7bd1ed4da8

package gset_test

import (
	"strconv"
	"testing"

	"github.com/gogf/gf/v2/container/gset"
)

var intSet = gset.NewIntSet(true)

var anySet = gset.NewSet(true)

var strSet = gset.NewStrSet(true)

var intSetUnsafe = gset.NewIntSet()

var anySetUnsafe = gset.NewSet()

var strSetUnsafe = gset.NewStrSet()

func Benchmark_IntSet_Add(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intSet.Add(i)
			i++
		}
	})
}

func Benchmark_IntSet_Contains(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intSet.Contains(i)
			i++
		}
	})
}

func Benchmark_IntSet_Remove(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intSet.Remove(i)
			i++
		}
	})
}

func Benchmark_AnySet_Add(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			anySet.Add(i)
			i++
		}
	})
}

func Benchmark_AnySet_Contains(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			anySet.Contains(i)
			i++
		}
	})
}

func Benchmark_AnySet_Remove(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			anySet.Remove(i)
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。. md5:1a28b26d80944830
func Benchmark_StrSet_Add(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strSet.Add(strconv.Itoa(i))
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。. md5:1a28b26d80944830
func Benchmark_StrSet_Contains(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strSet.Contains(strconv.Itoa(i))
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。. md5:1a28b26d80944830
func Benchmark_StrSet_Remove(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strSet.Remove(strconv.Itoa(i))
			i++
		}
	})
}

func Benchmark_Unsafe_IntSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSetUnsafe.Add(i)
	}
}

func Benchmark_Unsafe_IntSet_Contains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSetUnsafe.Contains(i)
	}
}

func Benchmark_Unsafe_IntSet_Remove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSetUnsafe.Remove(i)
	}
}

func Benchmark_Unsafe_AnySet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anySetUnsafe.Add(i)
	}
}

func Benchmark_Unsafe_AnySet_Contains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anySetUnsafe.Contains(i)
	}
}

func Benchmark_Unsafe_AnySet_Remove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anySetUnsafe.Remove(i)
	}
}

// 注意，字符串转换会带来额外的性能开销。. md5:1a28b26d80944830
func Benchmark_Unsafe_StrSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strSetUnsafe.Add(strconv.Itoa(i))
	}
}

// 注意，字符串转换会带来额外的性能开销。. md5:1a28b26d80944830
func Benchmark_Unsafe_StrSet_Contains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strSetUnsafe.Contains(strconv.Itoa(i))
	}
}

// 注意，字符串转换会带来额外的性能开销。. md5:1a28b26d80944830
func Benchmark_Unsafe_StrSet_Remove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strSetUnsafe.Remove(strconv.Itoa(i))
	}
}
