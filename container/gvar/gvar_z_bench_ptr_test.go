// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package 泛型类

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
