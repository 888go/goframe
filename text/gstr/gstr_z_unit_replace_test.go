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

	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func Test_Replace(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		t.Assert(gstr.X替换(s1, "ab", "AB"), "ABcdEFG乱入的中文ABcdefg")
		t.Assert(gstr.X替换(s1, "EF", "ef"), "abcdefG乱入的中文abcdefg")
		t.Assert(gstr.X替换(s1, "MN", "mn"), s1)

		t.Assert(gstr.X切片替换(s1, g.X文本切片{
			"a", "A",
			"A", "-",
			"a",
		}), "-bcdEFG乱入的中文-bcdefg")

		t.Assert(gstr.Map替换(s1, g.MapStrStr{
			"a": "A",
			"G": "g",
		}), "AbcdEFg乱入的中文Abcdefg")
	})
}

func Test_ReplaceI_1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcd乱入的中文ABCD"
		s2 := "a"
		t.Assert(gstr.X替换并忽略大小写(s1, "ab", "aa"), "aacd乱入的中文aaCD")
		t.Assert(gstr.X替换并忽略大小写(s1, "ab", "aa", 0), "abcd乱入的中文ABCD")
		t.Assert(gstr.X替换并忽略大小写(s1, "ab", "aa", 1), "aacd乱入的中文ABCD")

		t.Assert(gstr.X替换并忽略大小写(s1, "abcd", "-"), "-乱入的中文-")
		t.Assert(gstr.X替换并忽略大小写(s1, "abcd", "-", 1), "-乱入的中文ABCD")

		t.Assert(gstr.X替换并忽略大小写(s1, "abcd乱入的", ""), "中文ABCD")
		t.Assert(gstr.X替换并忽略大小写(s1, "ABCD乱入的", ""), "中文ABCD")

		t.Assert(gstr.X替换并忽略大小写(s2, "A", "-"), "-")
		t.Assert(gstr.X替换并忽略大小写(s2, "a", "-"), "-")

		t.Assert(gstr.X切片替换并忽略大小写(s1, g.X文本切片{
			"abcd乱入的", "-",
			"-", "=",
			"a",
		}), "=中文ABCD")

		t.Assert(gstr.Map替换并忽略大小写(s1, g.MapStrStr{
			"ab": "-",
			"CD": "=",
		}), "-=乱入的中文-=")
	})
}

func Test_ReplaceI_2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X替换并忽略大小写("aaa", "A", "-a-"), `-a--a--a-`)
		t.Assert(gstr.X替换并忽略大小写("aaaa", "AA", "-"), `--`)
		t.Assert(gstr.X替换并忽略大小写("a a a", "A", "b"), `b b b`)
		t.Assert(gstr.X替换并忽略大小写("aaaaaa", "aa", "a"), `aaa`)
		t.Assert(gstr.X替换并忽略大小写("aaaaaa", "AA", "A"), `AAA`)
		t.Assert(gstr.X替换并忽略大小写("aaa", "A", "AA"), `AAAAAA`)
		t.Assert(gstr.X替换并忽略大小写("aaa", "A", "AA"), `AAAAAA`)
		t.Assert(gstr.X替换并忽略大小写("a duration", "duration", "recordduration"), `a recordduration`)
	})
	// With count parameter.
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X替换并忽略大小写("aaaaaa", "aa", "a", 2), `aaaa`)
		t.Assert(gstr.X替换并忽略大小写("aaaaaa", "AA", "A", 1), `Aaaaa`)
		t.Assert(gstr.X替换并忽略大小写("aaaaaa", "AA", "A", 3), `AAA`)
		t.Assert(gstr.X替换并忽略大小写("aaaaaa", "AA", "A", 4), `AAA`)
		t.Assert(gstr.X替换并忽略大小写("aaa", "A", "AA", 2), `AAAAa`)
		t.Assert(gstr.X替换并忽略大小写("aaa", "A", "AA", 3), `AAAAAA`)
		t.Assert(gstr.X替换并忽略大小写("aaa", "A", "AA", 4), `AAAAAA`)
	})
}
