// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package 转换类

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
		X取整数(valueStr)
	}
}

func Benchmark_Str_To_Int8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数8位(valueStr)
	}
}

func Benchmark_Str_To_Int16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数16位(valueStr)
	}
}

func Benchmark_Str_To_Int32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数32位(valueStr)
	}
}

func Benchmark_Str_To_Int64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数(valueStr)
	}
}

func Benchmark_Str_To_Uint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数(valueStr)
	}
}

func Benchmark_Str_To_Uint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数8位(valueStr)
	}
}

func Benchmark_Str_To_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数16位(valueStr)
	}
}

func Benchmark_Str_To_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数32位(valueStr)
	}
}

func Benchmark_Str_To_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数64位(valueStr)
	}
}

func Benchmark_Str_To_Float32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取小数32位(valueStr)
	}
}

func Benchmark_Str_To_Float64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取小数64位(valueStr)
	}
}

func Benchmark_Str_To_Time(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取时间(valueStr)
	}
}

func Benchmark_Str_To_TimeDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取时长(valueStr)
	}
}

func Benchmark_Str_To_Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取字节集(valueStr)
	}
}

func Benchmark_Str_To_Strings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取文本数组(valueStr)
	}
}

func Benchmark_Str_To_Ints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数数组(valueStr)
	}
}

func Benchmark_Str_To_Floats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取小数数组(valueStr)
	}
}

func Benchmark_Str_To_Interfaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取any数组(valueStr)
	}
}
