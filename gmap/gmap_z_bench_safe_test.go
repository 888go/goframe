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

var anyAnyMap = map类.X创建AnyAny(true)

var intIntMap = map类.X创建IntInt(true)

var intAnyMap = map类.X创建IntAny(true)

var intStrMap = map类.X创建IntStr(true)

var strIntMap = map类.X创建StrInt(true)

var strAnyMap = map类.X创建StrAny(true)

var strStrMap = map类.X创建StrStr(true)

func Benchmark_IntIntMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intIntMap.X设置值(i, i)
			i++
		}
	})
}

func Benchmark_IntAnyMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intAnyMap.X设置值(i, i)
			i++
		}
	})
}

func Benchmark_IntStrMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intStrMap.X设置值(i, "123456789")
			i++
		}
	})
}

func Benchmark_AnyAnyMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			anyAnyMap.X设置值(i, i)
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。
func Benchmark_StrIntMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strIntMap.X设置值(strconv.Itoa(i), i)
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。
func Benchmark_StrAnyMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strAnyMap.X设置值(strconv.Itoa(i), i)
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。
func Benchmark_StrStrMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strStrMap.X设置值(strconv.Itoa(i), "123456789")
			i++
		}
	})
}

func Benchmark_IntIntMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intIntMap.X取值(i)
			i++
		}
	})
}

func Benchmark_IntAnyMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intAnyMap.X取值(i)
			i++
		}
	})
}

func Benchmark_IntStrMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			intStrMap.X取值(i)
			i++
		}
	})
}

func Benchmark_AnyAnyMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			anyAnyMap.X取值(i)
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。
func Benchmark_StrIntMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strIntMap.X取值(strconv.Itoa(i))
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。
func Benchmark_StrAnyMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strAnyMap.X取值(strconv.Itoa(i))
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。
func Benchmark_StrStrMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strStrMap.X取值(strconv.Itoa(i))
			i++
		}
	})
}
