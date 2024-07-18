// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package ghash_test//bm:哈希类_test

import (
	"testing"

	"github.com/gogf/gf/v2/encoding/ghash"
)

var (
	str = []byte("This is the test string for hash.")
)

func Benchmark_BKDR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.BKDR(str)
	}
}

func Benchmark_BKDR64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.BKDR64(str)
	}
}

func Benchmark_SDBM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.SDBM(str)
	}
}

func Benchmark_SDBM64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.SDBM64(str)
	}
}

func Benchmark_RS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.RS(str)
	}
}

func Benchmark_RS64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.RS64(str)
	}
}

func Benchmark_JS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.JS(str)
	}
}

func Benchmark_JS64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.JS64(str)
	}
}

func Benchmark_PJW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.PJW(str)
	}
}

func Benchmark_PJW64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.PJW64(str)
	}
}

func Benchmark_ELF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.ELF(str)
	}
}

func Benchmark_ELF64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.ELF64(str)
	}
}

func Benchmark_DJB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.DJB(str)
	}
}

func Benchmark_DJB64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.DJB64(str)
	}
}

func Benchmark_AP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.AP(str)
	}
}

func Benchmark_AP64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ghash.AP64(str)
	}
}
