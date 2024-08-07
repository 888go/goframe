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

var valueFloat = float64(1.23456789)

func Benchmark_Float_To_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(valueFloat)
	}
}

func Benchmark_Float_To_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数(valueFloat)
	}
}

func Benchmark_Float_To_Int8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数8位(valueFloat)
	}
}

func Benchmark_Float_To_Int16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数16位(valueFloat)
	}
}

func Benchmark_Float_To_Int32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数32位(valueFloat)
	}
}

func Benchmark_Float_To_Int64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数(valueFloat)
	}
}

func Benchmark_Float_To_Uint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数(valueFloat)
	}
}

func Benchmark_Float_To_Uint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数8位(valueFloat)
	}
}

func Benchmark_Float_To_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数16位(valueFloat)
	}
}

func Benchmark_Float_To_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数32位(valueFloat)
	}
}

func Benchmark_Float_To_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数64位(valueFloat)
	}
}

func Benchmark_Float_To_Float32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取小数32位(valueFloat)
	}
}

func Benchmark_Float_To_Float64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取小数64位(valueFloat)
	}
}

func Benchmark_Float_To_Time(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取时间(valueFloat)
	}
}

func Benchmark_Float_To_TimeDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取时长(valueFloat)
	}
}

func Benchmark_Float_To_Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取字节集(valueFloat)
	}
}

func Benchmark_Float_To_Strings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取文本切片(valueFloat)
	}
}

func Benchmark_Float_To_Ints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数切片(valueFloat)
	}
}

func Benchmark_Float_To_Floats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取小数切片(valueFloat)
	}
}

func Benchmark_Float_To_Interfaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取any切片(valueFloat)
	}
}
