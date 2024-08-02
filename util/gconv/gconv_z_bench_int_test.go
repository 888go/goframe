// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package 转换类

import (
	"testing"
)

var valueInt = 123456789

func Benchmark_Int_To_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(valueInt)
	}
}

func Benchmark_Int_To_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(valueInt)
	}
}

func Benchmark_Int_To_Int8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int8(valueInt)
	}
}

func Benchmark_Int_To_Int16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int16(valueInt)
	}
}

func Benchmark_Int_To_Int32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int32(valueInt)
	}
}

func Benchmark_Int_To_Int64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(valueInt)
	}
}

func Benchmark_Int_To_Uint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint(valueInt)
	}
}

func Benchmark_Int_To_Uint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint8(valueInt)
	}
}

func Benchmark_Int_To_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint16(valueInt)
	}
}

func Benchmark_Int_To_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32(valueInt)
	}
}

func Benchmark_Int_To_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint64(valueInt)
	}
}

func Benchmark_Int_To_Float32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float32(valueInt)
	}
}

func Benchmark_Int_To_Float64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64(valueInt)
	}
}

func Benchmark_Int_To_Time(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Time(valueInt)
	}
}

func Benchmark_Int_To_TimeDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Duration(valueInt)
	}
}

func Benchmark_Int_To_Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(valueInt)
	}
}

func Benchmark_Int_To_Strings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Strings(valueInt)
	}
}

func Benchmark_Int_To_Ints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ints(valueInt)
	}
}

func Benchmark_Int_To_Floats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Floats(valueInt)
	}
}

func Benchmark_Int_To_Interfaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Interfaces(valueInt)
	}
}
