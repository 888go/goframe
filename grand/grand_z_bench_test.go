// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 随机类_test

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	mathRand "math/rand"
	"testing"

	"github.com/888go/goframe/grand"
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
		随机类.X区间整数(0, 99)
	}
}

func Benchmark_Perm10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		随机类.X整数切片(10)
	}
}

func Benchmark_Perm100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		随机类.X整数切片(100)
	}
}

func Benchmark_Rand_N1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		随机类.X区间整数(0, 99)
	}
}

func Benchmark_Rand_N2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		随机类.X区间整数(0, 999999999)
	}
}

func Benchmark_B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		随机类.X字节集(16)
	}
}

func Benchmark_S(b *testing.B) {
	for i := 0; i < b.N; i++ {
		随机类.X文本(16)
	}
}

func Benchmark_S_Symbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		随机类.X文本(16, true)
	}
}

func Benchmark_Str(b *testing.B) {
	for i := 0; i < b.N; i++ {
		随机类.X从文本生成文本(strForStr, 16)
	}
}

func Benchmark_Symbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		随机类.X特殊字符文本(16)
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
