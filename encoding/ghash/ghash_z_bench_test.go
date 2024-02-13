// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 哈希类_test

import (
	"testing"
	
	"github.com/888go/goframe/encoding/ghash"
)

var (
	str = []byte("This is the test string for hash.")
)

func Benchmark_BKDR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.BKDR(str)
	}
}

func Benchmark_BKDR64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.BKDR64(str)
	}
}

func Benchmark_SDBM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.SDBM(str)
	}
}

func Benchmark_SDBM64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.SDBM64(str)
	}
}

func Benchmark_RS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.RS(str)
	}
}

func Benchmark_RS64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.RS64(str)
	}
}

func Benchmark_JS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.JS(str)
	}
}

func Benchmark_JS64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.JS64(str)
	}
}

func Benchmark_PJW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.PJW(str)
	}
}

func Benchmark_PJW64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.PJW64(str)
	}
}

func Benchmark_ELF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.ELF(str)
	}
}

func Benchmark_ELF64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.ELF64(str)
	}
}

func Benchmark_DJB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.DJB(str)
	}
}

func Benchmark_DJB64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.DJB64(str)
	}
}

func Benchmark_AP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.AP(str)
	}
}

func Benchmark_AP64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		哈希类.AP64(str)
	}
}
