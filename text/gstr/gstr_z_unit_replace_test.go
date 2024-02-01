// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package gstr_test
import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	)

func Test_Replace(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		t.Assert(gstr.Replace(s1, "ab", "AB"), "ABcdEFG乱入的中文ABcdefg")
		t.Assert(gstr.Replace(s1, "EF", "ef"), "abcdefG乱入的中文abcdefg")
		t.Assert(gstr.Replace(s1, "MN", "mn"), s1)

		t.Assert(gstr.ReplaceByArray(s1, g.ArrayStr{
			"a", "A",
			"A", "-",
			"a",
		}), "-bcdEFG乱入的中文-bcdefg")

		t.Assert(gstr.ReplaceByMap(s1, g.MapStrStr{
			"a": "A",
			"G": "g",
		}), "AbcdEFg乱入的中文Abcdefg")
	})
}

func Test_ReplaceI_1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcd乱入的中文ABCD"
		s2 := "a"
		t.Assert(gstr.ReplaceI(s1, "ab", "aa"), "aacd乱入的中文aaCD")
		t.Assert(gstr.ReplaceI(s1, "ab", "aa", 0), "abcd乱入的中文ABCD")
		t.Assert(gstr.ReplaceI(s1, "ab", "aa", 1), "aacd乱入的中文ABCD")

		t.Assert(gstr.ReplaceI(s1, "abcd", "-"), "-乱入的中文-")
		t.Assert(gstr.ReplaceI(s1, "abcd", "-", 1), "-乱入的中文ABCD")

		t.Assert(gstr.ReplaceI(s1, "abcd乱入的", ""), "中文ABCD")
		t.Assert(gstr.ReplaceI(s1, "ABCD乱入的", ""), "中文ABCD")

		t.Assert(gstr.ReplaceI(s2, "A", "-"), "-")
		t.Assert(gstr.ReplaceI(s2, "a", "-"), "-")

		t.Assert(gstr.ReplaceIByArray(s1, g.ArrayStr{
			"abcd乱入的", "-",
			"-", "=",
			"a",
		}), "=中文ABCD")

		t.Assert(gstr.ReplaceIByMap(s1, g.MapStrStr{
			"ab": "-",
			"CD": "=",
		}), "-=乱入的中文-=")
	})
}

func Test_ReplaceI_2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.ReplaceI("aaa", "A", "-a-"), `-a--a--a-`)
		t.Assert(gstr.ReplaceI("aaaa", "AA", "-"), `--`)
		t.Assert(gstr.ReplaceI("a a a", "A", "b"), `b b b`)
		t.Assert(gstr.ReplaceI("aaaaaa", "aa", "a"), `aaa`)
		t.Assert(gstr.ReplaceI("aaaaaa", "AA", "A"), `AAA`)
		t.Assert(gstr.ReplaceI("aaa", "A", "AA"), `AAAAAA`)
		t.Assert(gstr.ReplaceI("aaa", "A", "AA"), `AAAAAA`)
		t.Assert(gstr.ReplaceI("a duration", "duration", "recordduration"), `a recordduration`)
	})
	// 带有 count 参数。
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.ReplaceI("aaaaaa", "aa", "a", 2), `aaaa`)
		t.Assert(gstr.ReplaceI("aaaaaa", "AA", "A", 1), `Aaaaa`)
		t.Assert(gstr.ReplaceI("aaaaaa", "AA", "A", 3), `AAA`)
		t.Assert(gstr.ReplaceI("aaaaaa", "AA", "A", 4), `AAA`)
		t.Assert(gstr.ReplaceI("aaa", "A", "AA", 2), `AAAAa`)
		t.Assert(gstr.ReplaceI("aaa", "A", "AA", 3), `AAAAAA`)
		t.Assert(gstr.ReplaceI("aaa", "A", "AA", 4), `AAAAAA`)
	})
}
