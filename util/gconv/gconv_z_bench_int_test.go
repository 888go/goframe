// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gconv
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
