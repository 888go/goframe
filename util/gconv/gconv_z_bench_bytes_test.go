// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 使用`go test`命令测试所有`.go`文件，专注于"Benchmark_Bytes_To_*"命名的基准测试，并在运行时显示内存使用情况。 md5:de663064a4460648

package gconv

import (
	"testing"
	"unsafe"

	"github.com/gogf/gf/v2/encoding/gbinary"
)

var valueBytes = gbinary.Encode(123456789)

func Benchmark_Bytes_To_String_Normal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(valueBytes)
	}
}

func Benchmark_Bytes_To_String_Unsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = *(*string)(unsafe.Pointer(&valueBytes))
	}
}

func Benchmark_Bytes_To_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(valueBytes)
	}
}

func Benchmark_Bytes_To_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(valueBytes)
	}
}

func Benchmark_Bytes_To_Int8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int8(valueBytes)
	}
}

func Benchmark_Bytes_To_Int16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int16(valueBytes)
	}
}

func Benchmark_Bytes_To_Int32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int32(valueBytes)
	}
}

func Benchmark_Bytes_To_Int64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(valueBytes)
	}
}

func Benchmark_Bytes_To_Uint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint(valueBytes)
	}
}

func Benchmark_Bytes_To_Uint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint8(valueBytes)
	}
}

func Benchmark_Bytes_To_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint16(valueBytes)
	}
}

func Benchmark_Bytes_To_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32(valueBytes)
	}
}

func Benchmark_Bytes_To_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint64(valueBytes)
	}
}

func Benchmark_Bytes_To_Float32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float32(valueBytes)
	}
}

func Benchmark_Bytes_To_Float64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64(valueBytes)
	}
}

func Benchmark_Bytes_To_Time(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Time(valueBytes)
	}
}

func Benchmark_Bytes_To_TimeDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Duration(valueBytes)
	}
}

func Benchmark_Bytes_To_Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(valueBytes)
	}
}

func Benchmark_Bytes_To_Strings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Strings(valueBytes)
	}
}

func Benchmark_Bytes_To_Ints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ints(valueBytes)
	}
}

func Benchmark_Bytes_To_Floats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Floats(valueBytes)
	}
}

func Benchmark_Bytes_To_Interfaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Interfaces(valueBytes)
	}
}
