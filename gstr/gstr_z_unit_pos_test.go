// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 文本类_test

import (
	"testing"
	
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/888go/goframe/gstr"
)

func Test_Pos(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		t.Assert(文本类.X查找(s1, "ab"), 0)
		t.Assert(文本类.X查找(s1, "ab", 2), 7)
		t.Assert(文本类.X查找(s1, "abd", 0), -1)
		t.Assert(文本类.X查找(s1, "e", -4), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(文本类.X查找(s1, "爱"), 3)
		t.Assert(文本类.X查找(s1, "C"), 6)
		t.Assert(文本类.X查找(s1, "China"), 6)
	})
}

func Test_PosRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		t.Assert(文本类.X查找Unicode(s1, "ab"), 0)
		t.Assert(文本类.X查找Unicode(s1, "ab", 2), 7)
		t.Assert(文本类.X查找Unicode(s1, "abd", 0), -1)
		t.Assert(文本类.X查找Unicode(s1, "e", -4), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(文本类.X查找Unicode(s1, "爱"), 1)
		t.Assert(文本类.X查找Unicode(s1, "C"), 2)
		t.Assert(文本类.X查找Unicode(s1, "China"), 2)
	})
}

func Test_PosI(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		t.Assert(文本类.X查找并忽略大小写(s1, "zz"), -1)
		t.Assert(文本类.X查找并忽略大小写(s1, "ab"), 0)
		t.Assert(文本类.X查找并忽略大小写(s1, "ef", 2), 4)
		t.Assert(文本类.X查找并忽略大小写(s1, "abd", 0), -1)
		t.Assert(文本类.X查找并忽略大小写(s1, "E", -4), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(文本类.X查找并忽略大小写(s1, "爱"), 3)
		t.Assert(文本类.X查找并忽略大小写(s1, "c"), 6)
		t.Assert(文本类.X查找并忽略大小写(s1, "china"), 6)
	})
}

func Test_PosIRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		t.Assert(文本类.X查找并忽略大小写Unicode(s1, "zz"), -1)
		t.Assert(文本类.X查找并忽略大小写Unicode(s1, "ab"), 0)
		t.Assert(文本类.X查找并忽略大小写Unicode(s1, "ef", 2), 4)
		t.Assert(文本类.X查找并忽略大小写Unicode(s1, "abd", 0), -1)
		t.Assert(文本类.X查找并忽略大小写Unicode(s1, "E", -4), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(文本类.X查找并忽略大小写Unicode(s1, "爱"), 1)
		t.Assert(文本类.X查找并忽略大小写Unicode(s1, "c"), 2)
		t.Assert(文本类.X查找并忽略大小写Unicode(s1, "china"), 2)
	})
}

func Test_PosR(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		s2 := "abcdEFGz1cdeab"
		t.Assert(文本类.X倒找(s1, "zz"), -1)
		t.Assert(文本类.X倒找(s1, "ab"), 7)
		t.Assert(文本类.X倒找(s2, "ab", -2), 0)
		t.Assert(文本类.X倒找(s1, "ef"), 11)
		t.Assert(文本类.X倒找(s1, "abd", 0), -1)
		t.Assert(文本类.X倒找(s1, "e", -4), -1)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(文本类.X倒找(s1, "爱"), 3)
		t.Assert(文本类.X倒找(s1, "C"), 6)
		t.Assert(文本类.X倒找(s1, "China"), 6)
	})
}

func Test_PosRRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		s2 := "abcdEFGz1cdeab"
		t.Assert(文本类.X倒找Unicode(s1, "zz"), -1)
		t.Assert(文本类.X倒找Unicode(s1, "ab"), 7)
		t.Assert(文本类.X倒找Unicode(s2, "ab", -2), 0)
		t.Assert(文本类.X倒找Unicode(s1, "ef"), 11)
		t.Assert(文本类.X倒找Unicode(s1, "abd", 0), -1)
		t.Assert(文本类.X倒找Unicode(s1, "e", -4), -1)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(文本类.X倒找Unicode(s1, "爱"), 1)
		t.Assert(文本类.X倒找Unicode(s1, "C"), 2)
		t.Assert(文本类.X倒找Unicode(s1, "China"), 2)
	})
}

func Test_PosRI(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		s2 := "abcdEFGz1cdeab"
		t.Assert(文本类.X倒找并忽略大小写(s1, "zz"), -1)
		t.Assert(文本类.X倒找并忽略大小写(s1, "AB"), 7)
		t.Assert(文本类.X倒找并忽略大小写(s2, "AB", -2), 0)
		t.Assert(文本类.X倒找并忽略大小写(s1, "EF"), 11)
		t.Assert(文本类.X倒找并忽略大小写(s1, "abd", 0), -1)
		t.Assert(文本类.X倒找并忽略大小写(s1, "e", -5), 4)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(文本类.X倒找并忽略大小写(s1, "爱"), 3)
		t.Assert(文本类.X倒找并忽略大小写(s1, "C"), 19)
		t.Assert(文本类.X倒找并忽略大小写(s1, "China"), 6)
	})
}

func Test_PosRIRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFGabcdefg"
		s2 := "abcdEFGz1cdeab"
		t.Assert(文本类.X倒找并忽略大小写Unicode(s1, "zz"), -1)
		t.Assert(文本类.X倒找并忽略大小写Unicode(s1, "AB"), 7)
		t.Assert(文本类.X倒找并忽略大小写Unicode(s2, "AB", -2), 0)
		t.Assert(文本类.X倒找并忽略大小写Unicode(s1, "EF"), 11)
		t.Assert(文本类.X倒找并忽略大小写Unicode(s1, "abd", 0), -1)
		t.Assert(文本类.X倒找并忽略大小写Unicode(s1, "e", -5), 4)
	})
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱China very much"
		t.Assert(文本类.X倒找并忽略大小写Unicode(s1, "爱"), 1)
		t.Assert(文本类.X倒找并忽略大小写Unicode(s1, "C"), 15)
		t.Assert(文本类.X倒找并忽略大小写Unicode(s1, "China"), 2)
	})
}
