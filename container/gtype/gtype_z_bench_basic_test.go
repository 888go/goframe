// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gtype_test

import (
	"strconv"
	"sync/atomic"
	"testing"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/encoding/gbinary"
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
		it.Set(i)
	}
}

func BenchmarkInt_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it.Val()
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
		it32.Set(i)
	}
}

func BenchmarkInt32_Val(b *testing.B) {
	for i := int32(0); i < int32(b.N); i++ {
		it32.Val()
	}
}

func BenchmarkInt32_Add(b *testing.B) {
	for i := int32(0); i < int32(b.N); i++ {
		it32.Add(i)
	}
}

func BenchmarkInt64_Set(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		it64.Set(i)
	}
}

func BenchmarkInt64_Val(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		it64.Val()
	}
}

func BenchmarkInt64_Add(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		it64.Add(i)
	}
}

func BenchmarkUint_Set(b *testing.B) {
	for i := uint(0); i < uint(b.N); i++ {
		uit.Set(i)
	}
}

func BenchmarkUint_Val(b *testing.B) {
	for i := uint(0); i < uint(b.N); i++ {
		uit.Val()
	}
}

func BenchmarkUint_Add(b *testing.B) {
	for i := uint(0); i < uint(b.N); i++ {
		uit.Add(i)
	}
}

func BenchmarkUint32_Set(b *testing.B) {
	for i := uint32(0); i < uint32(b.N); i++ {
		uit32.Set(i)
	}
}

func BenchmarkUint32_Val(b *testing.B) {
	for i := uint32(0); i < uint32(b.N); i++ {
		uit32.Val()
	}
}

func BenchmarkUint32_Add(b *testing.B) {
	for i := uint32(0); i < uint32(b.N); i++ {
		uit32.Add(i)
	}
}

func BenchmarkUint64_Set(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		uit64.Set(i)
	}
}

func BenchmarkUint64_Val(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		uit64.Val()
	}
}

func BenchmarkUint64_Add(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		uit64.Add(i)
	}
}

func BenchmarkBool_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bl.Set(true)
	}
}

func BenchmarkBool_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bl.Val()
	}
}

func BenchmarkBool_Cas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bl.Cas(false, true)
	}
}

func BenchmarkString_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str.Set(strconv.Itoa(i))
	}
}

func BenchmarkString_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str.Val()
	}
}

func BenchmarkBytes_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vbytes.Set(gbinary.EncodeInt(i))
	}
}

func BenchmarkBytes_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vbytes.Val()
	}
}

func BenchmarkInterface_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inf.Set(i)
	}
}

func BenchmarkInterface_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inf.Val()
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
