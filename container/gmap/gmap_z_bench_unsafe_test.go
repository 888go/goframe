// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gmap_test

import (
	"strconv"
	"testing"
	
	"github.com/888go/goframe/container/gmap"
)

var anyAnyMapUnsafe = gmap.New()

var intIntMapUnsafe = gmap.NewIntIntMap()

var intAnyMapUnsafe = gmap.NewIntAnyMap()

var intStrMapUnsafe = gmap.NewIntStrMap()

var strIntMapUnsafe = gmap.NewStrIntMap()

var strAnyMapUnsafe = gmap.NewStrAnyMap()

var strStrMapUnsafe = gmap.NewStrStrMap()

// 写入基准测试

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

// 阅读基准测试。

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
