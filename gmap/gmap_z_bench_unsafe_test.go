// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package map类_test

import (
	"strconv"
	"testing"
	
	"github.com/888go/goframe/gmap"
)

var anyAnyMapUnsafe = map类.X创建()

var intIntMapUnsafe = map类.X创建IntInt()

var intAnyMapUnsafe = map类.X创建IntAny()

var intStrMapUnsafe = map类.X创建IntStr()

var strIntMapUnsafe = map类.X创建StrInt()

var strAnyMapUnsafe = map类.X创建StrAny()

var strStrMapUnsafe = map类.X创建StrStr()

// 写入基准测试

func Benchmark_Unsafe_IntIntMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intIntMapUnsafe.X设置值(i, i)
	}
}

func Benchmark_Unsafe_IntAnyMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intAnyMapUnsafe.X设置值(i, i)
	}
}

func Benchmark_Unsafe_IntStrMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intStrMapUnsafe.X设置值(i, strconv.Itoa(i))
	}
}

func Benchmark_Unsafe_AnyAnyMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anyAnyMapUnsafe.X设置值(i, i)
	}
}

func Benchmark_Unsafe_StrIntMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strIntMapUnsafe.X设置值(strconv.Itoa(i), i)
	}
}

func Benchmark_Unsafe_StrAnyMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strAnyMapUnsafe.X设置值(strconv.Itoa(i), i)
	}
}

func Benchmark_Unsafe_StrStrMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStrMapUnsafe.X设置值(strconv.Itoa(i), strconv.Itoa(i))
	}
}

// 阅读基准测试。

func Benchmark_Unsafe_IntIntMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intIntMapUnsafe.X取值(i)
	}
}

func Benchmark_Unsafe_IntAnyMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intAnyMapUnsafe.X取值(i)
	}
}

func Benchmark_Unsafe_IntStrMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intStrMapUnsafe.X取值(i)
	}
}

func Benchmark_Unsafe_AnyAnyMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anyAnyMapUnsafe.X取值(i)
	}
}

func Benchmark_Unsafe_StrIntMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strIntMapUnsafe.X取值(strconv.Itoa(i))
	}
}

func Benchmark_Unsafe_StrAnyMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strAnyMapUnsafe.X取值(strconv.Itoa(i))
	}
}

func Benchmark_Unsafe_StrStrMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStrMapUnsafe.X取值(strconv.Itoa(i))
	}
}
