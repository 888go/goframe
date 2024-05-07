// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 文本类_test

import (
	"testing"

	"github.com/888go/goframe/gstr"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Replace(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		t.Assert(文本类.X替换(s1, "ab", "AB"), "ABcdEFG乱入的中文ABcdefg")
		t.Assert(文本类.X替换(s1, "EF", "ef"), "abcdefG乱入的中文abcdefg")
		t.Assert(文本类.X替换(s1, "MN", "mn"), s1)

		t.Assert(文本类.X切片替换(s1, g.ArrayStr{
			"a", "A",
			"A", "-",
			"a",
		}), "-bcdEFG乱入的中文-bcdefg")

		t.Assert(文本类.Map替换(s1, g.MapStrStr{
			"a": "A",
			"G": "g",
		}), "AbcdEFg乱入的中文Abcdefg")
	})
}

func Test_ReplaceI_1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcd乱入的中文ABCD"
		s2 := "a"
		t.Assert(文本类.X替换并忽略大小写(s1, "ab", "aa"), "aacd乱入的中文aaCD")
		t.Assert(文本类.X替换并忽略大小写(s1, "ab", "aa", 0), "abcd乱入的中文ABCD")
		t.Assert(文本类.X替换并忽略大小写(s1, "ab", "aa", 1), "aacd乱入的中文ABCD")

		t.Assert(文本类.X替换并忽略大小写(s1, "abcd", "-"), "-乱入的中文-")
		t.Assert(文本类.X替换并忽略大小写(s1, "abcd", "-", 1), "-乱入的中文ABCD")

		t.Assert(文本类.X替换并忽略大小写(s1, "abcd乱入的", ""), "中文ABCD")
		t.Assert(文本类.X替换并忽略大小写(s1, "ABCD乱入的", ""), "中文ABCD")

		t.Assert(文本类.X替换并忽略大小写(s2, "A", "-"), "-")
		t.Assert(文本类.X替换并忽略大小写(s2, "a", "-"), "-")

		t.Assert(文本类.X切片替换并忽略大小写(s1, g.ArrayStr{
			"abcd乱入的", "-",
			"-", "=",
			"a",
		}), "=中文ABCD")

		t.Assert(文本类.Map替换并忽略大小写(s1, g.MapStrStr{
			"ab": "-",
			"CD": "=",
		}), "-=乱入的中文-=")
	})
}

func Test_ReplaceI_2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X替换并忽略大小写("aaa", "A", "-a-"), `-a--a--a-`)
		t.Assert(文本类.X替换并忽略大小写("aaaa", "AA", "-"), `--`)
		t.Assert(文本类.X替换并忽略大小写("a a a", "A", "b"), `b b b`)
		t.Assert(文本类.X替换并忽略大小写("aaaaaa", "aa", "a"), `aaa`)
		t.Assert(文本类.X替换并忽略大小写("aaaaaa", "AA", "A"), `AAA`)
		t.Assert(文本类.X替换并忽略大小写("aaa", "A", "AA"), `AAAAAA`)
		t.Assert(文本类.X替换并忽略大小写("aaa", "A", "AA"), `AAAAAA`)
		t.Assert(文本类.X替换并忽略大小写("a duration", "duration", "recordduration"), `a recordduration`)
	})
	// 带有 count 参数。
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X替换并忽略大小写("aaaaaa", "aa", "a", 2), `aaaa`)
		t.Assert(文本类.X替换并忽略大小写("aaaaaa", "AA", "A", 1), `Aaaaa`)
		t.Assert(文本类.X替换并忽略大小写("aaaaaa", "AA", "A", 3), `AAA`)
		t.Assert(文本类.X替换并忽略大小写("aaaaaa", "AA", "A", 4), `AAA`)
		t.Assert(文本类.X替换并忽略大小写("aaa", "A", "AA", 2), `AAAAa`)
		t.Assert(文本类.X替换并忽略大小写("aaa", "A", "AA", 3), `AAAAAA`)
		t.Assert(文本类.X替换并忽略大小写("aaa", "A", "AA", 4), `AAAAAA`)
	})
}
