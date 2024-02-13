// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 错误类_test

import (
	"errors"
	"testing"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
)

var (
	// 用于Wrap*函数基准测试的基础错误。
	baseError = errors.New("test")
)

func Benchmark_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		错误类.X创建("test")
	}
}

func Benchmark_Newf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		错误类.X创建并格式化("%s", "test")
	}
}

func Benchmark_Wrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		错误类.X多层错误(baseError, "test")
	}
}

func Benchmark_Wrapf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		错误类.X多层错误并格式化(baseError, "%s", "test")
	}
}

func Benchmark_NewSkip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		错误类.X创建并跳过堆栈(1, "test")
	}
}

func Benchmark_NewSkipf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		错误类.X创建并跳过堆栈与格式化(1, "%s", "test")
	}
}

func Benchmark_NewCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		错误类.X创建错误码(错误码类.New(500, "", nil), "test")
	}
}

func Benchmark_NewCodef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		错误类.X创建错误码并格式化(错误码类.New(500, "", nil), "%s", "test")
	}
}

func Benchmark_NewCodeSkip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		错误类.X创建错误码并跳过堆栈(错误码类.New(1, "", nil), 500, "test")
	}
}

func Benchmark_NewCodeSkipf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		错误类.X创建错误码并跳过堆栈与格式化(错误码类.New(1, "", nil), 500, "%s", "test")
	}
}

func Benchmark_WrapCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		错误类.X多层错误码(错误码类.New(500, "", nil), baseError, "test")
	}
}

func Benchmark_WrapCodef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		错误类.X多层错误码并格式化(错误码类.New(500, "", nil), baseError, "test")
	}
}
