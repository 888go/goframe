// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package 安全变量类_test

import (
	"strconv"
	"sync/atomic"
	"testing"

	gtype "github.com/888go/goframe/container/gtype"
	gbinary "github.com/888go/goframe/encoding/gbinary"
)

var (
	it     = gtype.NewInt()
	it32   = gtype.NewInt32()
	it64   = gtype.NewInt64()
	uit    = gtype.NewUint()
	uit32  = gtype.NewUint32()
	uit64  = gtype.NewUint64()
	bl     = gtype.NewBool()
	vbytes = gtype.NewBytes()
	str    = gtype.NewString()
	inf    = gtype.NewInterface()
	at     = atomic.Value{}
)

func BenchmarkInt_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it.X设置值(i)
	}
}

func BenchmarkInt_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it.X取值()
	}
}

func BenchmarkInt_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it.Add(i)
	}
}

func BenchmarkInt_Cas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it.Cas(i, i)
	}
}

func BenchmarkInt32_Set(b *testing.B) {
	for i := int32(0); i < int32(b.N); i++ {
		it32.X设置值(i)
	}
}

func BenchmarkInt32_Val(b *testing.B) {
	for i := int32(0); i < int32(b.N); i++ {
		it32.X取值()
	}
}

func BenchmarkInt32_Add(b *testing.B) {
	for i := int32(0); i < int32(b.N); i++ {
		it32.Add(i)
	}
}

func BenchmarkInt64_Set(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		it64.X设置值(i)
	}
}

func BenchmarkInt64_Val(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		it64.X取值()
	}
}

func BenchmarkInt64_Add(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		it64.Add(i)
	}
}

func BenchmarkUint_Set(b *testing.B) {
	for i := uint(0); i < uint(b.N); i++ {
		uit.X设置值(i)
	}
}

func BenchmarkUint_Val(b *testing.B) {
	for i := uint(0); i < uint(b.N); i++ {
		uit.X取值()
	}
}

func BenchmarkUint_Add(b *testing.B) {
	for i := uint(0); i < uint(b.N); i++ {
		uit.Add(i)
	}
}

func BenchmarkUint32_Set(b *testing.B) {
	for i := uint32(0); i < uint32(b.N); i++ {
		uit32.X设置值(i)
	}
}

func BenchmarkUint32_Val(b *testing.B) {
	for i := uint32(0); i < uint32(b.N); i++ {
		uit32.X取值()
	}
}

func BenchmarkUint32_Add(b *testing.B) {
	for i := uint32(0); i < uint32(b.N); i++ {
		uit32.Add(i)
	}
}

func BenchmarkUint64_Set(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		uit64.X设置值(i)
	}
}

func BenchmarkUint64_Val(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		uit64.X取值()
	}
}

func BenchmarkUint64_Add(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		uit64.Add(i)
	}
}

func BenchmarkBool_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bl.X设置值(true)
	}
}

func BenchmarkBool_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bl.X取值()
	}
}

func BenchmarkBool_Cas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bl.Cas(false, true)
	}
}

func BenchmarkString_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str.X设置值(strconv.Itoa(i))
	}
}

func BenchmarkString_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str.X取值()
	}
}

func BenchmarkBytes_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vbytes.X设置值(gbinary.EncodeInt(i))
	}
}

func BenchmarkBytes_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vbytes.X取值()
	}
}

func BenchmarkInterface_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inf.X设置值(i)
	}
}

func BenchmarkInterface_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inf.X取值()
	}
}

func BenchmarkAtomicValue_Store(b *testing.B) {
	for i := 0; i < b.N; i++ {
		at.Store(i)
	}
}

func BenchmarkAtomicValue_Load(b *testing.B) {
	for i := 0; i < b.N; i++ {
		at.Load()
	}
}
