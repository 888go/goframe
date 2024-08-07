// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 哈希类_test

import (
	"testing"

	ghash "github.com/888go/goframe/encoding/ghash"
	gtest "github.com/888go/goframe/test/gtest"
)

var (
	strBasic = []byte("This is the test string for hash.")
)

func Test_BKDR(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		x := uint32(200645773)
		j := ghash.BKDR(strBasic)
		t.Assert(j, x)
	})
	gtest.C(t, func(t *gtest.T) {
		x := uint64(4214762819217104013)
		j := ghash.BKDR64(strBasic)
		t.Assert(j, x)
	})
}

func Test_SDBM(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		x := uint32(1069170245)
		j := ghash.SDBM(strBasic)
		t.Assert(j, x)
	})
	gtest.C(t, func(t *gtest.T) {
		x := uint64(9881052176572890693)
		j := ghash.SDBM64(strBasic)
		t.Assert(j, x)
	})
}

func Test_RS(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		x := uint32(1944033799)
		j := ghash.RS(strBasic)
		t.Assert(j, x)
	})
	gtest.C(t, func(t *gtest.T) {
		x := uint64(13439708950444349959)
		j := ghash.RS64(strBasic)
		t.Assert(j, x)
	})
}

func Test_JS(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		x := uint32(498688898)
		j := ghash.JS(strBasic)
		t.Assert(j, x)
	})
	gtest.C(t, func(t *gtest.T) {
		x := uint64(13410163655098759877)
		j := ghash.JS64(strBasic)
		t.Assert(j, x)
	})
}

func Test_PJW(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		x := uint32(7244206)
		j := ghash.PJW(strBasic)
		t.Assert(j, x)
	})
	gtest.C(t, func(t *gtest.T) {
		x := uint64(31150)
		j := ghash.PJW64(strBasic)
		t.Assert(j, x)
	})
}

func Test_ELF(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		x := uint32(7244206)
		j := ghash.ELF(strBasic)
		t.Assert(j, x)
	})
	gtest.C(t, func(t *gtest.T) {
		x := uint64(31150)
		j := ghash.ELF64(strBasic)
		t.Assert(j, x)
	})
}

func Test_DJB(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		x := uint32(959862602)
		j := ghash.DJB(strBasic)
		t.Assert(j, x)
	})
	gtest.C(t, func(t *gtest.T) {
		x := uint64(2519720351310960458)
		j := ghash.DJB64(strBasic)
		t.Assert(j, x)
	})
}

func Test_AP(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		x := uint32(3998202516)
		j := ghash.AP(strBasic)
		t.Assert(j, x)
	})
	gtest.C(t, func(t *gtest.T) {
		x := uint64(2531023058543352243)
		j := ghash.AP64(strBasic)
		t.Assert(j, x)
	})
}
