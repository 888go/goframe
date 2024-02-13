// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 哈希类_test

import (
	"testing"
	
	"github.com/888go/goframe/encoding/ghash"
	"github.com/888go/goframe/test/gtest"
)

var (
	strBasic = []byte("This is the test string for hash.")
)

func Test_BKDR(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint32(200645773)
		j := 哈希类.BKDR(strBasic)
		t.Assert(j, x)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint64(4214762819217104013)
		j := 哈希类.BKDR64(strBasic)
		t.Assert(j, x)
	})
}

func Test_SDBM(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint32(1069170245)
		j := 哈希类.SDBM(strBasic)
		t.Assert(j, x)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint64(9881052176572890693)
		j := 哈希类.SDBM64(strBasic)
		t.Assert(j, x)
	})
}

func Test_RS(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint32(1944033799)
		j := 哈希类.RS(strBasic)
		t.Assert(j, x)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint64(13439708950444349959)
		j := 哈希类.RS64(strBasic)
		t.Assert(j, x)
	})
}

func Test_JS(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint32(498688898)
		j := 哈希类.JS(strBasic)
		t.Assert(j, x)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint64(13410163655098759877)
		j := 哈希类.JS64(strBasic)
		t.Assert(j, x)
	})
}

func Test_PJW(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint32(7244206)
		j := 哈希类.PJW(strBasic)
		t.Assert(j, x)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint64(31150)
		j := 哈希类.PJW64(strBasic)
		t.Assert(j, x)
	})
}

func Test_ELF(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint32(7244206)
		j := 哈希类.ELF(strBasic)
		t.Assert(j, x)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint64(31150)
		j := 哈希类.ELF64(strBasic)
		t.Assert(j, x)
	})
}

func Test_DJB(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint32(959862602)
		j := 哈希类.DJB(strBasic)
		t.Assert(j, x)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint64(2519720351310960458)
		j := 哈希类.DJB64(strBasic)
		t.Assert(j, x)
	})
}

func Test_AP(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint32(3998202516)
		j := 哈希类.AP(strBasic)
		t.Assert(j, x)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		x := uint64(2531023058543352243)
		j := 哈希类.AP64(strBasic)
		t.Assert(j, x)
	})
}
