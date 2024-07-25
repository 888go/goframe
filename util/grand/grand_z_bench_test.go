// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package grand_test

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	mathRand "math/rand"
	"testing"

	"github.com/gogf/gf/v2/util/grand"
)

var (
	buffer         = make([]byte, 8)
	randBuffer4    = make([]byte, 4)
	randBuffer1024 = make([]byte, 1024)
	strForStr      = "我爱GoFrame"
)

func Benchmark_Math_Rand_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathRand.Int()
	}
}

func Benchmark_CryptoRand_Buffer4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cryptoRand.Read(randBuffer4)
	}
}

func Benchmark_CryptoRand_Buffer1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cryptoRand.Read(randBuffer1024)
	}
}

func Benchmark_GRand_Intn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grand.N(0, 99)
	}
}

func Benchmark_Perm10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grand.Perm(10)
	}
}

func Benchmark_Perm100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grand.Perm(100)
	}
}

func Benchmark_Rand_N1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grand.N(0, 99)
	}
}

func Benchmark_Rand_N2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grand.N(0, 999999999)
	}
}

func Benchmark_B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grand.B(16)
	}
}

func Benchmark_S(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grand.S(16)
	}
}

func Benchmark_S_Symbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grand.S(16, true)
	}
}

func Benchmark_Str(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grand.Str(strForStr, 16)
	}
}

func Benchmark_Symbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grand.Symbols(16)
	}
}

func Benchmark_Uint32Converting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary.LittleEndian.Uint32([]byte{1, 1, 1, 1})
	}
}

func Benchmark_CryptoRand_Buffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := cryptoRand.Read(buffer); err == nil {
			binary.LittleEndian.Uint64(buffer)
		}
	}
}
