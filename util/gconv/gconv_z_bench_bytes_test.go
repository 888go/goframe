// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 使用go test命令测试当前目录下所有.go文件，执行名称中包含"Benchmark_Bytes_To_"的基准测试，并开启内存使用统计

package 转换类

import (
	"testing"
	"unsafe"
	
	"github.com/888go/goframe/encoding/gbinary"
)

var valueBytes = 字节集类.Encode(123456789)

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
		X取整数(valueBytes)
	}
}

func Benchmark_Bytes_To_Int8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数8位(valueBytes)
	}
}

func Benchmark_Bytes_To_Int16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数16位(valueBytes)
	}
}

func Benchmark_Bytes_To_Int32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数32位(valueBytes)
	}
}

func Benchmark_Bytes_To_Int64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数(valueBytes)
	}
}

func Benchmark_Bytes_To_Uint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数(valueBytes)
	}
}

func Benchmark_Bytes_To_Uint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数8位(valueBytes)
	}
}

func Benchmark_Bytes_To_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数16位(valueBytes)
	}
}

func Benchmark_Bytes_To_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数32位(valueBytes)
	}
}

func Benchmark_Bytes_To_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取正整数64位(valueBytes)
	}
}

func Benchmark_Bytes_To_Float32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取小数32位(valueBytes)
	}
}

func Benchmark_Bytes_To_Float64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取小数64位(valueBytes)
	}
}

func Benchmark_Bytes_To_Time(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取时间(valueBytes)
	}
}

func Benchmark_Bytes_To_TimeDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取时长(valueBytes)
	}
}

func Benchmark_Bytes_To_Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取字节集(valueBytes)
	}
}

func Benchmark_Bytes_To_Strings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取文本数组(valueBytes)
	}
}

func Benchmark_Bytes_To_Ints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取整数数组(valueBytes)
	}
}

func Benchmark_Bytes_To_Floats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取小数数组(valueBytes)
	}
}

func Benchmark_Bytes_To_Interfaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X取any数组(valueBytes)
	}
}
