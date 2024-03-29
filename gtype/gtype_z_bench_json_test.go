// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行Go语言测试，对所有.go文件进行测试，并且仅针对名称中包含"_Json"的基准测试（benchmark）进行执行，同时显示内存分配统计信息。
// 以下是逐行详细解释：
// ```go
// 使用go test命令来运行测试
// 测试的文件为当前目录下所有的.go文件（即`*.go`）
// `-bench=".+\_Json"` 参数表示仅执行那些基准测试函数名称中包含"_Json"的基准测试
// `-benchmem` 参数表示在输出的基准测试结果中，包含内存分配统计信息

package 安全变量类_test

import (
	"testing"
	
	"github.com/888go/goframe/gtype"
	"github.com/888go/goframe/gtype/internal/json"
)

var (
	vBool      = 安全变量类.NewBool()
	vByte      = 安全变量类.NewByte()
	vBytes     = 安全变量类.NewBytes()
	vFloat32   = 安全变量类.NewFloat32()
	vFloat64   = 安全变量类.NewFloat64()
	vInt       = 安全变量类.NewInt()
	vInt32     = 安全变量类.NewInt32()
	vInt64     = 安全变量类.NewInt64()
	vInterface = 安全变量类.NewInterface()
	vString    = 安全变量类.NewString()
	vUint      = 安全变量类.NewUint()
	vUint32    = 安全变量类.NewUint32()
	vUint64    = 安全变量类.NewUint64()
)

func Benchmark_Bool_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vBool)
	}
}

func Benchmark_Byte_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vByte)
	}
}

func Benchmark_Bytes_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vBytes)
	}
}

func Benchmark_Float32_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vFloat32)
	}
}

func Benchmark_Float64_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vFloat64)
	}
}

func Benchmark_Int_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vInt)
	}
}

func Benchmark_Int32_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vInt32)
	}
}

func Benchmark_Int64_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vInt64)
	}
}

func Benchmark_Interface_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vInterface)
	}
}

func Benchmark_String_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vString)
	}
}

func Benchmark_Uint_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vUint)
	}
}

func Benchmark_Uint32_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vUint64)
	}
}
