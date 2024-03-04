// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gvar

import (
	"testing"
)

var varPtr = New(nil)

func Benchmark_Ptr_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Set(i)
	}
}

func Benchmark_Ptr_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Val()
	}
}

func Benchmark_Ptr_IsNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.IsNil()
	}
}

func Benchmark_Ptr_Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Bytes()
	}
}

func Benchmark_Ptr_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.String()
	}
}

func Benchmark_Ptr_Bool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Bool()
	}
}

func Benchmark_Ptr_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Int()
	}
}

func Benchmark_Ptr_Int8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Int8()
	}
}

func Benchmark_Ptr_Int16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Int16()
	}
}

func Benchmark_Ptr_Int32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Int32()
	}
}

func Benchmark_Ptr_Int64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Int64()
	}
}

func Benchmark_Ptr_Uint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Uint()
	}
}

func Benchmark_Ptr_Uint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Uint8()
	}
}

func Benchmark_Ptr_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Uint16()
	}
}

func Benchmark_Ptr_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Uint32()
	}
}

func Benchmark_Ptr_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Uint64()
	}
}

func Benchmark_Ptr_Float32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Float32()
	}
}

func Benchmark_Ptr_Float64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Float64()
	}
}

func Benchmark_Ptr_Ints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Ints()
	}
}

func Benchmark_Ptr_Strings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Strings()
	}
}

func Benchmark_Ptr_Floats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Floats()
	}
}

func Benchmark_Ptr_Interfaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.Interfaces()
	}
}
