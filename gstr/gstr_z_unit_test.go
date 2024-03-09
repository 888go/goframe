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

func Test_ToLower(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		e1 := "abcdefg乱入的中文abcdefg"
		t.Assert(文本类.X到小写(s1), e1)
	})
}

func Test_ToUpper(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		e1 := "ABCDEFG乱入的中文ABCDEFG"
		t.Assert(文本类.X到大写(s1), e1)
	})
}

func Test_UcFirst(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		e1 := "AbcdEFG乱入的中文abcdefg"
		t.Assert(文本类.X到首字母大写(""), "")
		t.Assert(文本类.X到首字母大写(s1), e1)
		t.Assert(文本类.X到首字母大写(e1), e1)
	})
}

func Test_LcFirst(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "AbcdEFG乱入的中文abcdefg"
		e1 := "abcdEFG乱入的中文abcdefg"
		t.Assert(文本类.X到首字母小写(""), "")
		t.Assert(文本类.X到首字母小写(s1), e1)
		t.Assert(文本类.X到首字母小写(e1), e1)
	})
}

func Test_UcWords(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱GF: i love go frame"
		e1 := "我爱GF: I Love Go Frame"
		t.Assert(文本类.X到单词首字母大写(s1), e1)
	})
}

func Test_IsLetterLower(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X是否小写字符('a'), true)
		t.Assert(文本类.X是否小写字符('A'), false)
		t.Assert(文本类.X是否小写字符('1'), false)
	})
}

func Test_IsLetterUpper(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X是否大写字符('a'), false)
		t.Assert(文本类.X是否大写字符('A'), true)
		t.Assert(文本类.X是否大写字符('1'), false)
	})
}

func Test_IsNumeric(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X是否为数字("1a我"), false)
		t.Assert(文本类.X是否为数字("0123"), true)
		t.Assert(文本类.X是否为数字("我是中国人"), false)
		t.Assert(文本类.X是否为数字("1.2.3.4"), false)
	})
}

func Test_SubStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X按长度取文本("我爱GoFrame", 0), "我爱GoFrame")
		t.Assert(文本类.X按长度取文本("我爱GoFrame", 6), "GoFrame")
		t.Assert(文本类.X按长度取文本("我爱GoFrame", 6, 2), "Go")
		t.Assert(文本类.X按长度取文本("我爱GoFrame", -1, 30), "e")
		t.Assert(文本类.X按长度取文本("我爱GoFrame", 30, 30), "")
		t.Assert(文本类.X按长度取文本("abcdef", 0, -1), "abcde")
		t.Assert(文本类.X按长度取文本("abcdef", 2, -1), "cde")
		t.Assert(文本类.X按长度取文本("abcdef", 4, -4), "")
		t.Assert(文本类.X按长度取文本("abcdef", -3, -1), "de")
	})
}

func Test_SubStrRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X按长度取文本Unicode("我爱GoFrame", 0), "我爱GoFrame")
		t.Assert(文本类.X按长度取文本Unicode("我爱GoFrame", 2), "GoFrame")
		t.Assert(文本类.X按长度取文本Unicode("我爱GoFrame", 2, 2), "Go")
		t.Assert(文本类.X按长度取文本Unicode("我爱GoFrame", -1, 30), "e")
		t.Assert(文本类.X按长度取文本Unicode("我爱GoFrame", 30, 30), "")
		t.Assert(文本类.X按长度取文本Unicode("abcdef", 0, -1), "abcde")
		t.Assert(文本类.X按长度取文本Unicode("abcdef", 2, -1), "cde")
		t.Assert(文本类.X按长度取文本Unicode("abcdef", 4, -4), "")
		t.Assert(文本类.X按长度取文本Unicode("abcdef", -3, -1), "de")
		t.Assert(文本类.X按长度取文本Unicode("我爱GoFrame呵呵", -3, 100), "e呵呵")
		t.Assert(文本类.X按长度取文本Unicode("abcdef哈哈", -3, -1), "f哈")
		t.Assert(文本类.X按长度取文本Unicode("ab我爱GoFramecdef哈哈", -3, -1), "f哈")
		t.Assert(文本类.X按长度取文本Unicode("我爱GoFrame", 0, 3), "我爱G")
	})
}

func Test_StrLimit(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X按长度取左边并带前缀("我爱GoFrame", 6), "我爱...")
		t.Assert(文本类.X按长度取左边并带前缀("我爱GoFrame", 6, ""), "我爱")
		t.Assert(文本类.X按长度取左边并带前缀("我爱GoFrame", 6, "**"), "我爱**")
		t.Assert(文本类.X按长度取左边并带前缀("我爱GoFrame", 8, ""), "我爱Go")
		t.Assert(文本类.X按长度取左边并带前缀("*", 4, ""), "*")
	})
}

func Test_StrLimitRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X按长度取左边并带前缀Unicode("我爱GoFrame", 2), "我爱...")
		t.Assert(文本类.X按长度取左边并带前缀Unicode("我爱GoFrame", 2, ""), "我爱")
		t.Assert(文本类.X按长度取左边并带前缀Unicode("我爱GoFrame", 2, "**"), "我爱**")
		t.Assert(文本类.X按长度取左边并带前缀Unicode("我爱GoFrame", 4, ""), "我爱Go")
		t.Assert(文本类.X按长度取左边并带前缀Unicode("*", 4, ""), "*")
	})
}

func Test_HasPrefix(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X开头判断("我爱GoFrame", "我爱"), true)
		t.Assert(文本类.X开头判断("en我爱GoFrame", "我爱"), false)
		t.Assert(文本类.X开头判断("en我爱GoFrame", "en"), true)
	})
}

func Test_HasSuffix(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X末尾判断("我爱GoFrame", "GoFrame"), true)
		t.Assert(文本类.X末尾判断("en我爱GoFrame", "a"), false)
		t.Assert(文本类.X末尾判断("GoFrame很棒", "棒"), true)
	})
}

func Test_Reverse(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X反转字符("我爱123"), "321爱我")
	})
}

func Test_NumberFormat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X格式化数值(1234567.8910, 2, ".", ","), "1,234,567.89")
		t.Assert(文本类.X格式化数值(1234567.8910, 2, "#", "/"), "1/234/567#89")
		t.Assert(文本类.X格式化数值(-1234567.8910, 2, "#", "/"), "-1/234/567#89")
	})
}

func Test_ChunkSplit(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X长度分割("1234", 1, "#"), "1#2#3#4#")
		t.Assert(文本类.X长度分割("我爱123", 1, "#"), "我#爱#1#2#3#")
		t.Assert(文本类.X长度分割("1234", 1, ""), "1\r\n2\r\n3\r\n4\r\n")
	})
}

func Test_SplitAndTrim(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := `

010    

020  

`
		a := 文本类.X分割并忽略空值(s, "\n", "0")
		t.Assert(len(a), 2)
		t.Assert(a[0], "1")
		t.Assert(a[1], "2")
	})
}

func Test_Fields(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X单词分割("我爱 Go Frame"), []string{
			"我爱", "Go", "Frame",
		})
	})
}

func Test_CountWords(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X统计单词数("我爱 Go Go Go"), map[string]int{
			"Go": 3,
			"我爱": 1,
		})
	})
}

func Test_CountChars(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X统计unicode字符数("我爱 Go Go Go"), map[string]int{
			" ": 3,
			"G": 3,
			"o": 3,
			"我": 1,
			"爱": 1,
		})
		t.Assert(文本类.X统计unicode字符数("我爱 Go Go Go", true), map[string]int{
			"G": 3,
			"o": 3,
			"我": 1,
			"爱": 1,
		})
	})
}

func Test_LenRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X取字符长度("1234"), 4)
		t.Assert(文本类.X取字符长度("我爱GoFrame"), 9)
	})
}

func Test_Repeat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X生成重复文本("go", 3), "gogogo")
		t.Assert(文本类.X生成重复文本("好的", 3), "好的好的好的")
	})
}

func Test_Str(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X取右边并含关键字("name@example.com", "@"), "@example.com")
		t.Assert(文本类.X取右边并含关键字("name@example.com", ""), "")
		t.Assert(文本类.X取右边并含关键字("name@example.com", "z"), "")
	})
}

func Test_StrEx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X取右边("name@example.com", "@"), "example.com")
		t.Assert(文本类.X取右边("name@example.com", ""), "")
		t.Assert(文本类.X取右边("name@example.com", "z"), "")
	})
}

func Test_StrTill(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X取左边并含关键字("name@example.com", "@"), "name@")
		t.Assert(文本类.X取左边并含关键字("name@example.com", ""), "")
		t.Assert(文本类.X取左边并含关键字("name@example.com", "z"), "")
	})
}

func Test_StrTillEx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X取左边("name@example.com", "@"), "name")
		t.Assert(文本类.X取左边("name@example.com", ""), "")
		t.Assert(文本类.X取左边("name@example.com", "z"), "")
	})
}

func Test_Shuffle(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(文本类.X随机打散字符("123456")), 6)
	})
}

func Test_Split(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X分割("1.2", "."), []string{"1", "2"})
		t.Assert(文本类.X分割("我爱 - GoFrame", " - "), []string{"我爱", "GoFrame"})
	})
}

func Test_Join(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X连接([]string{"我爱", "GoFrame"}, " - "), "我爱 - GoFrame")
	})
}

func Test_Explode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.Explode别名(" - ", "我爱 - GoFrame"), []string{"我爱", "GoFrame"})
	})
}

func Test_Implode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.Implode别名(" - ", []string{"我爱", "GoFrame"}), "我爱 - GoFrame")
	})
}

func Test_Chr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X整数到ascii(65), "A")
	})
}

func Test_Ord(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.Ord("A"), 65)
	})
}

func Test_HideStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X替换中间字符("15928008611", 40, "*"), "159****8611")
		t.Assert(文本类.X替换中间字符("john@kohg.cn", 40, "*"), "jo*n@kohg.cn")
		t.Assert(文本类.X替换中间字符("张三", 50, "*"), "张*")
		t.Assert(文本类.X替换中间字符("张小三", 50, "*"), "张*三")
		t.Assert(文本类.X替换中间字符("欧阳小三", 50, "*"), "欧**三")
	})
}

func Test_Nl2Br(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X替换换行符("1\n2"), "1<br>2")
		t.Assert(文本类.X替换换行符("1\r\n2"), "1<br>2")
		t.Assert(文本类.X替换换行符("1\r\n2", true), "1<br />2")
	})
}

func Test_AddSlashes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X转义(`1'2"3\`), `1\'2\"3\\`)
	})
}

func Test_StripSlashes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X转义还原(`1\'2\"3\\`), `1'2"3\`)
	})
}

func Test_QuoteMeta(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X转义并按字符(`.\+*?[^]($)`), `\.\\\+\*\?\[\^\]\(\$\)`)
		t.Assert(文本类.X转义并按字符(`.\+*中国?[^]($)`), `\.\\\+\*中国\?\[\^\]\(\$\)`)
		t.Assert(文本类.X转义并按字符(`.''`, `'`), `.\'\'`)
		t.Assert(文本类.X转义并按字符(`中国.''`, `'`), `中国.\'\'`)
	})
}

func Test_Count(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "abcdaAD"
		t.Assert(文本类.X统计次数(s, "0"), 0)
		t.Assert(文本类.X统计次数(s, "a"), 2)
		t.Assert(文本类.X统计次数(s, "b"), 1)
		t.Assert(文本类.X统计次数(s, "d"), 1)
	})
}

func Test_CountI(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "abcdaAD"
		t.Assert(文本类.X统计次数并忽略大小写(s, "0"), 0)
		t.Assert(文本类.X统计次数并忽略大小写(s, "a"), 3)
		t.Assert(文本类.X统计次数并忽略大小写(s, "b"), 1)
		t.Assert(文本类.X统计次数并忽略大小写(s, "d"), 2)
	})
}

func Test_Compare(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X顺序比较("a", "b"), -1)
		t.Assert(文本类.X顺序比较("a", "a"), 0)
		t.Assert(文本类.X顺序比较("b", "a"), 1)
	})
}

func Test_Equal(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X相等比较并忽略大小写("a", "A"), true)
		t.Assert(文本类.X相等比较并忽略大小写("a", "a"), true)
		t.Assert(文本类.X相等比较并忽略大小写("b", "a"), false)
	})
}

func Test_Contains(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X是否包含("abc", "a"), true)
		t.Assert(文本类.X是否包含("abc", "A"), false)
		t.Assert(文本类.X是否包含("abc", "ab"), true)
		t.Assert(文本类.X是否包含("abc", "abc"), true)
	})
}

func Test_ContainsI(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X是否包含并忽略大小写("abc", "a"), true)
		t.Assert(文本类.X是否包含并忽略大小写("abc", "A"), true)
		t.Assert(文本类.X是否包含并忽略大小写("abc", "Ab"), true)
		t.Assert(文本类.X是否包含并忽略大小写("abc", "ABC"), true)
		t.Assert(文本类.X是否包含并忽略大小写("abc", "ABCD"), false)
		t.Assert(文本类.X是否包含并忽略大小写("abc", "D"), false)
	})
}

func Test_ContainsAny(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X是否包含Any("abc", "a"), true)
		t.Assert(文本类.X是否包含Any("abc", "cd"), true)
		t.Assert(文本类.X是否包含Any("abc", "de"), false)
		t.Assert(文本类.X是否包含Any("abc", "A"), false)
	})
}

func Test_SubStrFrom(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.SubStrFrom别名("我爱GoFrameGood", `G`), "GoFrameGood")
		t.Assert(文本类.SubStrFrom别名("我爱GoFrameGood", `GG`), "")
		t.Assert(文本类.SubStrFrom别名("我爱GoFrameGood", `我`), "我爱GoFrameGood")
		t.Assert(文本类.SubStrFrom别名("我爱GoFrameGood", `Frame`), "FrameGood")
	})
}

func Test_SubStrFromEx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.SubStrFromEx别名("我爱GoFrameGood", `Go`), "FrameGood")
		t.Assert(文本类.SubStrFromEx别名("我爱GoFrameGood", `GG`), "")
		t.Assert(文本类.SubStrFromEx别名("我爱GoFrameGood", `我`), "爱GoFrameGood")
		t.Assert(文本类.SubStrFromEx别名("我爱GoFrameGood", `Frame`), `Good`)
	})
}

func Test_SubStrFromR(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X取右边并倒找与含关键字("我爱GoFrameGood", `G`), "Good")
		t.Assert(文本类.X取右边并倒找与含关键字("我爱GoFrameGood", `GG`), "")
		t.Assert(文本类.X取右边并倒找与含关键字("我爱GoFrameGood", `我`), "我爱GoFrameGood")
		t.Assert(文本类.X取右边并倒找与含关键字("我爱GoFrameGood", `Frame`), "FrameGood")
	})
}

func Test_SubStrFromREx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(文本类.X取右边并倒找("我爱GoFrameGood", `G`), "ood")
		t.Assert(文本类.X取右边并倒找("我爱GoFrameGood", `GG`), "")
		t.Assert(文本类.X取右边并倒找("我爱GoFrameGood", `我`), "爱GoFrameGood")
		t.Assert(文本类.X取右边并倒找("我爱GoFrameGood", `Frame`), `Good`)
	})
}
