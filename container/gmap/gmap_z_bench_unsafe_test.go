// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package map类_test

import (
	"strconv"
	"testing"

	gmap "github.com/888go/goframe/container/gmap"
)

var anyAnyMapUnsafe = gmap.New()

var intIntMapUnsafe = gmap.NewIntIntMap()

var intAnyMapUnsafe = gmap.NewIntAnyMap()

var intStrMapUnsafe = gmap.NewIntStrMap()

var strIntMapUnsafe = gmap.NewStrIntMap()

var strAnyMapUnsafe = gmap.NewStrAnyMap()

var strStrMapUnsafe = gmap.NewStrStrMap()

// Writing benchmarks.

func Benchmark_Unsafe_IntIntMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intIntMapUnsafe.Set(i, i)
	}
}

func Benchmark_Unsafe_IntAnyMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intAnyMapUnsafe.Set(i, i)
	}
}

func Benchmark_Unsafe_IntStrMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intStrMapUnsafe.Set(i, strconv.Itoa(i))
	}
}

func Benchmark_Unsafe_AnyAnyMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anyAnyMapUnsafe.Set(i, i)
	}
}

func Benchmark_Unsafe_StrIntMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strIntMapUnsafe.Set(strconv.Itoa(i), i)
	}
}

func Benchmark_Unsafe_StrAnyMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strAnyMapUnsafe.Set(strconv.Itoa(i), i)
	}
}

func Benchmark_Unsafe_StrStrMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStrMapUnsafe.Set(strconv.Itoa(i), strconv.Itoa(i))
	}
}

// Reading benchmarks.

func Benchmark_Unsafe_IntIntMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intIntMapUnsafe.Get(i)
	}
}

func Benchmark_Unsafe_IntAnyMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intAnyMapUnsafe.Get(i)
	}
}

func Benchmark_Unsafe_IntStrMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intStrMapUnsafe.Get(i)
	}
}

func Benchmark_Unsafe_AnyAnyMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anyAnyMapUnsafe.Get(i)
	}
}

func Benchmark_Unsafe_StrIntMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strIntMapUnsafe.Get(strconv.Itoa(i))
	}
}

func Benchmark_Unsafe_StrAnyMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strAnyMapUnsafe.Get(strconv.Itoa(i))
	}
}

func Benchmark_Unsafe_StrStrMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStrMapUnsafe.Get(strconv.Itoa(i))
	}
}
