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

var anyAnyMap = gmap.X创建AnyAny(true)

var intIntMap = gmap.X创建IntInt(true)

var intAnyMap = gmap.X创建IntAny(true)

var intStrMap = gmap.X创建IntStr(true)

var strIntMap = gmap.X创建StrInt(true)

var strAnyMap = gmap.X创建StrAny(true)

var strStrMap = gmap.X创建StrStr(true)

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

// 注意，字符串转换会带来额外的性能开销。 md5:1a28b26d80944830
func Benchmark_StrIntMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strIntMap.X设置值(strconv.Itoa(i), i)
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。 md5:1a28b26d80944830
func Benchmark_StrAnyMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strAnyMap.X设置值(strconv.Itoa(i), i)
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。 md5:1a28b26d80944830
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

// 注意，字符串转换会带来额外的性能开销。 md5:1a28b26d80944830
func Benchmark_StrIntMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strIntMap.X取值(strconv.Itoa(i))
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。 md5:1a28b26d80944830
func Benchmark_StrAnyMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strAnyMap.X取值(strconv.Itoa(i))
			i++
		}
	})
}

// 注意，字符串转换会带来额外的性能开销。 md5:1a28b26d80944830
func Benchmark_StrStrMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			strStrMap.X取值(strconv.Itoa(i))
			i++
		}
	})
}
