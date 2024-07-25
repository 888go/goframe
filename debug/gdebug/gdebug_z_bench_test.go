// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package gdebug

import (
	"runtime"
	"runtime/debug"
	"testing"
)

func Benchmark_BinVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinVersion()
	}
}

func Benchmark_BinVersionMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinVersionMd5()
	}
}

func Benchmark_RuntimeCaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.Caller(0)
	}
}

func Benchmark_RuntimeFuncForPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.FuncForPC(11010101)
	}
}

func Benchmark_callerFromIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callerFromIndex(nil)
	}
}

func Benchmark_Stack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Stack()
	}
}

func Benchmark_StackOfStdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		debug.Stack()
	}
}

func Benchmark_StackWithFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StackWithFilter([]string{"test"})
	}
}

func Benchmark_Caller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Caller()
	}
}

func Benchmark_CallerWithFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerWithFilter([]string{"test"})
	}
}

func Benchmark_CallerFilePath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerFilePath()
	}
}

func Benchmark_CallerDirectory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerDirectory()
	}
}

func Benchmark_CallerFileLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerFileLine()
	}
}

func Benchmark_CallerFileLineShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerFileLineShort()
	}
}

func Benchmark_CallerFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerFunction()
	}
}

func Benchmark_CallerPackage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerPackage()
	}
}
