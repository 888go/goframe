// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package 文本类_test

import (
	"testing"

	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func Test_Pos(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		t.Assert(gstr.Pos(s1, "ab"), 0)
		t.Assert(gstr.Pos(s1, "ab", 2), 7)
		t.Assert(gstr.Pos(s1, "abd", 0), -1)
		t.Assert(gstr.Pos(s1, "e", -4), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(gstr.Pos(s1, "爱"), 3)
		t.Assert(gstr.Pos(s1, "C"), 6)
		t.Assert(gstr.Pos(s1, "China"), 6)
	})
}

func Test_PosRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		t.Assert(gstr.PosRune(s1, "ab"), 0)
		t.Assert(gstr.PosRune(s1, "ab", 2), 7)
		t.Assert(gstr.PosRune(s1, "abd", 0), -1)
		t.Assert(gstr.PosRune(s1, "e", -4), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(gstr.PosRune(s1, "爱"), 1)
		t.Assert(gstr.PosRune(s1, "C"), 2)
		t.Assert(gstr.PosRune(s1, "China"), 2)
	})
}

func Test_PosI(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		t.Assert(gstr.PosI(s1, "zz"), -1)
		t.Assert(gstr.PosI(s1, "ab"), 0)
		t.Assert(gstr.PosI(s1, "ef", 2), 4)
		t.Assert(gstr.PosI(s1, "abd", 0), -1)
		t.Assert(gstr.PosI(s1, "E", -4), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(gstr.PosI(s1, "爱"), 3)
		t.Assert(gstr.PosI(s1, "c"), 6)
		t.Assert(gstr.PosI(s1, "china"), 6)
	})
}

func Test_PosIRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		t.Assert(gstr.PosIRune(s1, "zz"), -1)
		t.Assert(gstr.PosIRune(s1, "ab"), 0)
		t.Assert(gstr.PosIRune(s1, "ef", 2), 4)
		t.Assert(gstr.PosIRune(s1, "abd", 0), -1)
		t.Assert(gstr.PosIRune(s1, "E", -4), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(gstr.PosIRune(s1, "爱"), 1)
		t.Assert(gstr.PosIRune(s1, "c"), 2)
		t.Assert(gstr.PosIRune(s1, "china"), 2)
	})
}

func Test_PosR(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		s2 := "abcdEFGz1cdeab"
		t.Assert(gstr.PosR(s1, "zz"), -1)
		t.Assert(gstr.PosR(s1, "ab"), 7)
		t.Assert(gstr.PosR(s2, "ab", -2), 0)
		t.Assert(gstr.PosR(s1, "ef"), 11)
		t.Assert(gstr.PosR(s1, "abd", 0), -1)
		t.Assert(gstr.PosR(s1, "e", -4), -1)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(gstr.PosR(s1, "爱"), 3)
		t.Assert(gstr.PosR(s1, "C"), 6)
		t.Assert(gstr.PosR(s1, "China"), 6)
	})
}

func Test_PosRRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		s2 := "abcdEFGz1cdeab"
		t.Assert(gstr.PosRRune(s1, "zz"), -1)
		t.Assert(gstr.PosRRune(s1, "ab"), 7)
		t.Assert(gstr.PosRRune(s2, "ab", -2), 0)
		t.Assert(gstr.PosRRune(s1, "ef"), 11)
		t.Assert(gstr.PosRRune(s1, "abd", 0), -1)
		t.Assert(gstr.PosRRune(s1, "e", -4), -1)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(gstr.PosRRune(s1, "爱"), 1)
		t.Assert(gstr.PosRRune(s1, "C"), 2)
		t.Assert(gstr.PosRRune(s1, "China"), 2)
	})
}

func Test_PosRI(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		s2 := "abcdEFGz1cdeab"
		t.Assert(gstr.PosRI(s1, "zz"), -1)
		t.Assert(gstr.PosRI(s1, "AB"), 7)
		t.Assert(gstr.PosRI(s2, "AB", -2), 0)
		t.Assert(gstr.PosRI(s1, "EF"), 11)
		t.Assert(gstr.PosRI(s1, "abd", 0), -1)
		t.Assert(gstr.PosRI(s1, "e", -5), 4)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(gstr.PosRI(s1, "爱"), 3)
		t.Assert(gstr.PosRI(s1, "C"), 19)
		t.Assert(gstr.PosRI(s1, "China"), 6)
	})
}

func Test_PosRIRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		s2 := "abcdEFGz1cdeab"
		t.Assert(gstr.PosRIRune(s1, "zz"), -1)
		t.Assert(gstr.PosRIRune(s1, "AB"), 7)
		t.Assert(gstr.PosRIRune(s2, "AB", -2), 0)
		t.Assert(gstr.PosRIRune(s1, "EF"), 11)
		t.Assert(gstr.PosRIRune(s1, "abd", 0), -1)
		t.Assert(gstr.PosRIRune(s1, "e", -5), 4)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(gstr.PosRIRune(s1, "爱"), 1)
		t.Assert(gstr.PosRIRune(s1, "C"), 15)
		t.Assert(gstr.PosRIRune(s1, "China"), 2)
	})
}
