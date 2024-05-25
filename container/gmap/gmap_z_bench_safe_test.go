// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。. md5:81db3d7bd1ed4da8

package gmap_test

import (
	"strconv"
	"testing"

	"github.com/gogf/gf/v2/container/gmap"
)

var anyAnyMap = gmap.NewAnyAnyMap(true)

var intIntMap = gmap.NewIntIntMap(true)

var intAnyMap = gmap.NewIntAnyMap(true)

var intStrMap = gmap.NewIntStrMap(true)

var strIntMap = gmap.NewStrIntMap(true)

var strAnyMap = gmap.NewStrAnyMap(true)

var strStrMap = gmap.NewStrStrMap(true)

func Benchmark_IntIntMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intIntMap.Set(i, i)
			i++
		}
	})
}

func Benchmark_IntAnyMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intAnyMap.Set(i, i)
			i++
		}
	})
}

func Benchmark_IntStrMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intStrMap.Set(i, "123456789")
			i++
		}
	})
}

func Benchmark_AnyAnyMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			anyAnyMap.Set(i, i)
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。. md5:1a28b26d80944830
func Benchmark_StrIntMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strIntMap.Set(strconv.Itoa(i), i)
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。. md5:1a28b26d80944830
func Benchmark_StrAnyMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strAnyMap.Set(strconv.Itoa(i), i)
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。. md5:1a28b26d80944830
func Benchmark_StrStrMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strStrMap.Set(strconv.Itoa(i), "123456789")
			i++
		}
	})
}

func Benchmark_IntIntMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intIntMap.Get(i)
			i++
		}
	})
}

func Benchmark_IntAnyMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intAnyMap.Get(i)
			i++
		}
	})
}

func Benchmark_IntStrMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intStrMap.Get(i)
			i++
		}
	})
}

func Benchmark_AnyAnyMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			anyAnyMap.Get(i)
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。. md5:1a28b26d80944830
func Benchmark_StrIntMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strIntMap.Get(strconv.Itoa(i))
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。. md5:1a28b26d80944830
func Benchmark_StrAnyMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strAnyMap.Get(strconv.Itoa(i))
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。. md5:1a28b26d80944830
func Benchmark_StrStrMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strStrMap.Get(strconv.Itoa(i))
			i++
		}
	})
}
