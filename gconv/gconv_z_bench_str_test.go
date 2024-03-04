// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gconv

import (
	"testing"
)

var valueStr = "123456789"

func Benchmark_Str_To_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(valueStr)
	}
}

func Benchmark_Str_To_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(valueStr)
	}
}

func Benchmark_Str_To_Int8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int8(valueStr)
	}
}

func Benchmark_Str_To_Int16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int16(valueStr)
	}
}

func Benchmark_Str_To_Int32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int32(valueStr)
	}
}

func Benchmark_Str_To_Int64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(valueStr)
	}
}

func Benchmark_Str_To_Uint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint(valueStr)
	}
}

func Benchmark_Str_To_Uint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint8(valueStr)
	}
}

func Benchmark_Str_To_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint16(valueStr)
	}
}

func Benchmark_Str_To_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32(valueStr)
	}
}

func Benchmark_Str_To_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint64(valueStr)
	}
}

func Benchmark_Str_To_Float32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float32(valueStr)
	}
}

func Benchmark_Str_To_Float64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64(valueStr)
	}
}

func Benchmark_Str_To_Time(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Time(valueStr)
	}
}

func Benchmark_Str_To_TimeDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Duration(valueStr)
	}
}

func Benchmark_Str_To_Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(valueStr)
	}
}

func Benchmark_Str_To_Strings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Strings(valueStr)
	}
}

func Benchmark_Str_To_Ints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ints(valueStr)
	}
}

func Benchmark_Str_To_Floats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Floats(valueStr)
	}
}

func Benchmark_Str_To_Interfaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Interfaces(valueStr)
	}
}
