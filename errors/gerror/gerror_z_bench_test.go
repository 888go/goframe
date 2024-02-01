// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gerror_test
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
		gerror.New("test")
	}
}

func Benchmark_Newf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gerror.Newf("%s", "test")
	}
}

func Benchmark_Wrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gerror.Wrap(baseError, "test")
	}
}

func Benchmark_Wrapf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gerror.Wrapf(baseError, "%s", "test")
	}
}

func Benchmark_NewSkip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gerror.NewSkip(1, "test")
	}
}

func Benchmark_NewSkipf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gerror.NewSkipf(1, "%s", "test")
	}
}

func Benchmark_NewCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gerror.NewCode(gcode.New(500, "", nil), "test")
	}
}

func Benchmark_NewCodef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gerror.NewCodef(gcode.New(500, "", nil), "%s", "test")
	}
}

func Benchmark_NewCodeSkip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gerror.NewCodeSkip(gcode.New(1, "", nil), 500, "test")
	}
}

func Benchmark_NewCodeSkipf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gerror.NewCodeSkipf(gcode.New(1, "", nil), 500, "%s", "test")
	}
}

func Benchmark_WrapCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gerror.WrapCode(gcode.New(500, "", nil), baseError, "test")
	}
}

func Benchmark_WrapCodef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gerror.WrapCodef(gcode.New(500, "", nil), baseError, "test")
	}
}
