// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 正则类_test

import (
	"strings"
	"testing"
	
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/888go/goframe/gregex"
)

var (
	PatternErr = `([\d+`
)

func Test_Quote(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := `[foo]` //`\[foo\]`
		t.Assert(正则类.X转义特殊符号(s1), `\[foo\]`)
	})
}

func Test_Validate(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var s1 = `(.+):(\d+)`
		t.Assert(正则类.X表达式验证(s1), nil)
		s1 = `((.+):(\d+)`
		t.Assert(正则类.X表达式验证(s1) == nil, false)
	})
}

func Test_IsMatch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var pattern = `(.+):(\d+)`
		s1 := []byte(`sfs:2323`)
		t.Assert(正则类.X是否匹配字节集(pattern, s1), true)
		s1 = []byte(`sfs2323`)
		t.Assert(正则类.X是否匹配字节集(pattern, s1), false)
		s1 = []byte(`sfs:`)
		t.Assert(正则类.X是否匹配字节集(pattern, s1), false)
		// error pattern
		t.Assert(正则类.X是否匹配字节集(PatternErr, s1), false)
	})
}

func Test_IsMatchString(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var pattern = `(.+):(\d+)`
		s1 := `sfs:2323`
		t.Assert(正则类.X是否匹配文本(pattern, s1), true)
		s1 = `sfs2323`
		t.Assert(正则类.X是否匹配文本(pattern, s1), false)
		s1 = `sfs:`
		t.Assert(正则类.X是否匹配文本(pattern, s1), false)
		// error pattern
		t.Assert(正则类.X是否匹配文本(PatternErr, s1), false)
	})
}

func Test_Match(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		s := "acbb" + wantSubs + "dd"
		subs, err := 正则类.X匹配字节集(re, []byte(s))
		t.AssertNil(err)
		if string(subs[0]) != wantSubs {
			t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[0], wantSubs)
		}
		if string(subs[1]) != "aab" {
			t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[1], "aab")
		}
		// error pattern
		_, err = 正则类.X匹配字节集(PatternErr, []byte(s))
		t.AssertNE(err, nil)
	})
}

func Test_MatchString(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		s := "acbb" + wantSubs + "dd"
		subs, err := 正则类.X匹配文本(re, s)
		t.AssertNil(err)
		if string(subs[0]) != wantSubs {
			t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[0], wantSubs)
		}
		if string(subs[1]) != "aab" {
			t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[1], "aab")
		}
		// error pattern
		_, err = 正则类.X匹配文本(PatternErr, s)
		t.AssertNE(err, nil)
	})
}

func Test_MatchAll(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		s := "acbb" + wantSubs + "dd"
		s = s + `其他的` + s
		subs, err := 正则类.X匹配全部字节集(re, []byte(s))
		t.AssertNil(err)
		if string(subs[0][0]) != wantSubs {
			t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[0][0], wantSubs)
		}
		if string(subs[0][1]) != "aab" {
			t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[0][1], "aab")
		}

		if string(subs[1][0]) != wantSubs {
			t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[1][0], wantSubs)
		}
		if string(subs[1][1]) != "aab" {
			t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[1][1], "aab")
		}
		// error pattern
		_, err = 正则类.X匹配全部字节集(PatternErr, []byte(s))
		t.AssertNE(err, nil)
	})
}

func Test_MatchAllString(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		s := "acbb" + wantSubs + "dd"
		subs, err := 正则类.X匹配全部文本(re, s+`其他的`+s)
		t.AssertNil(err)
		if string(subs[0][0]) != wantSubs {
			t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[0][0], wantSubs)
		}
		if string(subs[0][1]) != "aab" {
			t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[0][1], "aab")
		}

		if string(subs[1][0]) != wantSubs {
			t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[1][0], wantSubs)
		}
		if string(subs[1][1]) != "aab" {
			t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[1][1], "aab")
		}
		// error pattern
		_, err = 正则类.X匹配全部文本(PatternErr, s)
		t.AssertNE(err, nil)
	})
}

func Test_Replace(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		replace := "12345"
		s := "acbb" + wantSubs + "dd"
		wanted := "acbb" + replace + "dd"
		replacedStr, err := 正则类.X替换字节集(re, []byte(replace), []byte(s))
		t.AssertNil(err)
		if string(replacedStr) != wanted {
			t.Fatalf("regex:%s,old:%s; want %q", re, s, wanted)
		}
		// error pattern
		_, err = 正则类.X替换字节集(PatternErr, []byte(replace), []byte(s))
		t.AssertNE(err, nil)
	})
}

func Test_ReplaceString(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		replace := "12345"
		s := "acbb" + wantSubs + "dd"
		wanted := "acbb" + replace + "dd"
		replacedStr, err := 正则类.X替换文本(re, replace, s)
		t.AssertNil(err)
		if replacedStr != wanted {
			t.Fatalf("regex:%s,old:%s; want %q", re, s, wanted)
		}
		// error pattern
		_, err = 正则类.X替换文本(PatternErr, replace, s)
		t.AssertNE(err, nil)
	})
}

func Test_ReplaceFun(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		//replace :="12345"
		s := "acbb" + wantSubs + "dd"
		wanted := "acbb[x" + wantSubs + "y]dd"
		wanted = "acbb" + "3个a" + "dd"
		replacedStr, err := 正则类.X替换字节集_函数(re, []byte(s), func(s []byte) []byte {
			if strings.Contains(string(s), "aaa") {
				return []byte("3个a")
			}
			return []byte("[x" + string(s) + "y]")
		})
		t.AssertNil(err)
		if string(replacedStr) != wanted {
			t.Fatalf("regex:%s,old:%s; want %q", re, s, wanted)
		}
		// error pattern
		_, err = 正则类.X替换字节集_函数(PatternErr, []byte(s), func(s []byte) []byte {
			return []byte("")
		})
		t.AssertNE(err, nil)
	})
}

func Test_ReplaceFuncMatch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := []byte("1234567890")
		p := `(\d{3})(\d{3})(.+)`
		s0, e0 := 正则类.ReplaceFuncMatch(p, s, func(match [][]byte) []byte {
			return match[0]
		})
		t.Assert(e0, nil)
		t.Assert(s0, s)
		s1, e1 := 正则类.ReplaceFuncMatch(p, s, func(match [][]byte) []byte {
			return match[1]
		})
		t.Assert(e1, nil)
		t.Assert(s1, []byte("123"))
		s2, e2 := 正则类.ReplaceFuncMatch(p, s, func(match [][]byte) []byte {
			return match[2]
		})
		t.Assert(e2, nil)
		t.Assert(s2, []byte("456"))
		s3, e3 := 正则类.ReplaceFuncMatch(p, s, func(match [][]byte) []byte {
			return match[3]
		})
		t.Assert(e3, nil)
		t.Assert(s3, []byte("7890"))
		// error pattern
		_, err := 正则类.ReplaceFuncMatch(PatternErr, s, func(match [][]byte) []byte {
			return match[3]
		})
		t.AssertNE(err, nil)
	})
}

func Test_ReplaceStringFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		//replace :="12345"
		s := "acbb" + wantSubs + "dd"
		wanted := "acbb[x" + wantSubs + "y]dd"
		wanted = "acbb" + "3个a" + "dd"
		replacedStr, err := 正则类.X替换文本_函数(re, s, func(s string) string {
			if strings.Contains(s, "aaa") {
				return "3个a"
			}
			return "[x" + s + "y]"
		})
		t.AssertNil(err)
		if replacedStr != wanted {
			t.Fatalf("regex:%s,old:%s; want %q", re, s, wanted)
		}
		// error pattern
		_, err = 正则类.X替换文本_函数(PatternErr, s, func(s string) string {
			return ""
		})
		t.AssertNE(err, nil)
	})
}

func Test_ReplaceStringFuncMatch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "1234567890"
		p := `(\d{3})(\d{3})(.+)`
		s0, e0 := 正则类.ReplaceStringFuncMatch(p, s, func(match []string) string {
			return match[0]
		})
		t.Assert(e0, nil)
		t.Assert(s0, s)
		s1, e1 := 正则类.ReplaceStringFuncMatch(p, s, func(match []string) string {
			return match[1]
		})
		t.Assert(e1, nil)
		t.Assert(s1, "123")
		s2, e2 := 正则类.ReplaceStringFuncMatch(p, s, func(match []string) string {
			return match[2]
		})
		t.Assert(e2, nil)
		t.Assert(s2, "456")
		s3, e3 := 正则类.ReplaceStringFuncMatch(p, s, func(match []string) string {
			return match[3]
		})
		t.Assert(e3, nil)
		t.Assert(s3, "7890")
		// error pattern
		_, err := 正则类.ReplaceStringFuncMatch(PatternErr, s, func(match []string) string {
			return ""
		})
		t.AssertNE(err, nil)
	})
}

func Test_Split(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		re := "a(a+b+)b"
		matched := "aaabb"
		item0 := "acbb"
		item1 := "dd"
		s := item0 + matched + item1
		t.Assert(正则类.X是否匹配文本(re, matched), true)
		items := 正则类.X分割(re, s) // 通过匹配的分隔符拆分字符串
		if items[0] != item0 {
			t.Fatalf("regex:%s,Split(%q) want %q", re, s, item0)
		}
		if items[1] != item1 {
			t.Fatalf("regex:%s,Split(%q) want %q", re, s, item0)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		re := "a(a+b+)b"
		notmatched := "aaxbb"
		item0 := "acbb"
		item1 := "dd"
		s := item0 + notmatched + item1
		t.Assert(正则类.X是否匹配文本(re, notmatched), false)
		items := 正则类.X分割(re, s) // 使用notmatched进行字符串分割，若未匹配则不进行分割
		if items[0] != s {
			t.Fatalf("regex:%s,Split(%q) want %q", re, s, item0)
		}
		// error pattern
		items = 正则类.X分割(PatternErr, s)
		t.AssertEQ(items, nil)

	})
}
