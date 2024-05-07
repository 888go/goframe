// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package 泛型类

import (
	"testing"
)

var varPtr = X创建(nil)

func Benchmark_Ptr_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X设置值(i)
	}
}

func Benchmark_Ptr_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取值()
	}
}

func Benchmark_Ptr_IsNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X是否为Nil()
	}
}

func Benchmark_Ptr_Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取字节集()
	}
}

func Benchmark_Ptr_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.String()
	}
}

func Benchmark_Ptr_Bool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取布尔()
	}
}

func Benchmark_Ptr_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取整数()
	}
}

func Benchmark_Ptr_Int8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取整数8位()
	}
}

func Benchmark_Ptr_Int16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取整数16位()
	}
}

func Benchmark_Ptr_Int32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取整数32位()
	}
}

func Benchmark_Ptr_Int64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取整数64位()
	}
}

func Benchmark_Ptr_Uint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取正整数()
	}
}

func Benchmark_Ptr_Uint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取正整数8位()
	}
}

func Benchmark_Ptr_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取正整数16位()
	}
}

func Benchmark_Ptr_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取正整数32位()
	}
}

func Benchmark_Ptr_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取正整数64位()
	}
}

func Benchmark_Ptr_Float32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取小数32位()
	}
}

func Benchmark_Ptr_Float64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取小数64位()
	}
}

func Benchmark_Ptr_Ints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取整数切片()
	}
}

func Benchmark_Ptr_Strings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取文本切片()
	}
}

func Benchmark_Ptr_Floats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取小数切片()
	}
}

func Benchmark_Ptr_Interfaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		varPtr.X取any切片()
	}
}
