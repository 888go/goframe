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

func Test_ToLower(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		e1 := "abcdefg乱入的中文abcdefg"
		t.Assert(gstr.X到小写(s1), e1)
	})
}

func Test_ToUpper(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		e1 := "ABCDEFG乱入的中文ABCDEFG"
		t.Assert(gstr.X到大写(s1), e1)
	})
}

func Test_UcFirst(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		e1 := "AbcdEFG乱入的中文abcdefg"
		t.Assert(gstr.X到首字母大写(""), "")
		t.Assert(gstr.X到首字母大写(s1), e1)
		t.Assert(gstr.X到首字母大写(e1), e1)
	})
}

func Test_LcFirst(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "AbcdEFG乱入的中文abcdefg"
		e1 := "abcdEFG乱入的中文abcdefg"
		t.Assert(gstr.X到首字母小写(""), "")
		t.Assert(gstr.X到首字母小写(s1), e1)
		t.Assert(gstr.X到首字母小写(e1), e1)
	})
}

func Test_UcWords(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := "我爱GF: i love go frame"
		e1 := "我爱GF: I Love Go Frame"
		t.Assert(gstr.X到单词首字母大写(s1), e1)
	})
}

func Test_IsLetterLower(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X是否小写字符('a'), true)
		t.Assert(gstr.X是否小写字符('A'), false)
		t.Assert(gstr.X是否小写字符('1'), false)
	})
}

func Test_IsLetterUpper(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X是否大写字符('a'), false)
		t.Assert(gstr.X是否大写字符('A'), true)
		t.Assert(gstr.X是否大写字符('1'), false)
	})
}

func Test_IsNumeric(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X是否为数字("1a我"), false)
		t.Assert(gstr.X是否为数字("0123"), true)
		t.Assert(gstr.X是否为数字("我是中国人"), false)
		t.Assert(gstr.X是否为数字("1.2.3.4"), false)
	})
}

func Test_SubStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X按长度取文本("我爱GoFrame", 0), "我爱GoFrame")
		t.Assert(gstr.X按长度取文本("我爱GoFrame", 6), "GoFrame")
		t.Assert(gstr.X按长度取文本("我爱GoFrame", 6, 2), "Go")
		t.Assert(gstr.X按长度取文本("我爱GoFrame", -1, 30), "e")
		t.Assert(gstr.X按长度取文本("我爱GoFrame", 30, 30), "")
		t.Assert(gstr.X按长度取文本("abcdef", 0, -1), "abcde")
		t.Assert(gstr.X按长度取文本("abcdef", 2, -1), "cde")
		t.Assert(gstr.X按长度取文本("abcdef", 4, -4), "")
		t.Assert(gstr.X按长度取文本("abcdef", -3, -1), "de")
	})
}

func Test_SubStrRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X按长度取文本Unicode("我爱GoFrame", 0), "我爱GoFrame")
		t.Assert(gstr.X按长度取文本Unicode("我爱GoFrame", 2), "GoFrame")
		t.Assert(gstr.X按长度取文本Unicode("我爱GoFrame", 2, 2), "Go")
		t.Assert(gstr.X按长度取文本Unicode("我爱GoFrame", -1, 30), "e")
		t.Assert(gstr.X按长度取文本Unicode("我爱GoFrame", 30, 30), "")
		t.Assert(gstr.X按长度取文本Unicode("abcdef", 0, -1), "abcde")
		t.Assert(gstr.X按长度取文本Unicode("abcdef", 2, -1), "cde")
		t.Assert(gstr.X按长度取文本Unicode("abcdef", 4, -4), "")
		t.Assert(gstr.X按长度取文本Unicode("abcdef", -3, -1), "de")
		t.Assert(gstr.X按长度取文本Unicode("我爱GoFrame呵呵", -3, 100), "e呵呵")
		t.Assert(gstr.X按长度取文本Unicode("abcdef哈哈", -3, -1), "f哈")
		t.Assert(gstr.X按长度取文本Unicode("ab我爱GoFramecdef哈哈", -3, -1), "f哈")
		t.Assert(gstr.X按长度取文本Unicode("我爱GoFrame", 0, 3), "我爱G")
	})
}

func Test_StrLimit(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X按长度取左边并带前缀("我爱GoFrame", 6), "我爱...")
		t.Assert(gstr.X按长度取左边并带前缀("我爱GoFrame", 6, ""), "我爱")
		t.Assert(gstr.X按长度取左边并带前缀("我爱GoFrame", 6, "**"), "我爱**")
		t.Assert(gstr.X按长度取左边并带前缀("我爱GoFrame", 8, ""), "我爱Go")
		t.Assert(gstr.X按长度取左边并带前缀("*", 4, ""), "*")
	})
}

func Test_StrLimitRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X按长度取左边并带前缀Unicode("我爱GoFrame", 2), "我爱...")
		t.Assert(gstr.X按长度取左边并带前缀Unicode("我爱GoFrame", 2, ""), "我爱")
		t.Assert(gstr.X按长度取左边并带前缀Unicode("我爱GoFrame", 2, "**"), "我爱**")
		t.Assert(gstr.X按长度取左边并带前缀Unicode("我爱GoFrame", 4, ""), "我爱Go")
		t.Assert(gstr.X按长度取左边并带前缀Unicode("*", 4, ""), "*")
	})
}

func Test_HasPrefix(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X开头判断("我爱GoFrame", "我爱"), true)
		t.Assert(gstr.X开头判断("en我爱GoFrame", "我爱"), false)
		t.Assert(gstr.X开头判断("en我爱GoFrame", "en"), true)
	})
}

func Test_HasSuffix(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X末尾判断("我爱GoFrame", "GoFrame"), true)
		t.Assert(gstr.X末尾判断("en我爱GoFrame", "a"), false)
		t.Assert(gstr.X末尾判断("GoFrame很棒", "棒"), true)
	})
}

func Test_Reverse(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X反转字符("我爱123"), "321爱我")
	})
}

func Test_NumberFormat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X格式化数值(1234567.8910, 2, ".", ","), "1,234,567.89")
		t.Assert(gstr.X格式化数值(1234567.8910, 2, "#", "/"), "1/234/567#89")
		t.Assert(gstr.X格式化数值(-1234567.8910, 2, "#", "/"), "-1/234/567#89")
	})
}

func Test_ChunkSplit(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X长度分割("1234", 1, "#"), "1#2#3#4#")
		t.Assert(gstr.X长度分割("我爱123", 1, "#"), "我#爱#1#2#3#")
		t.Assert(gstr.X长度分割("1234", 1, ""), "1\r\n2\r\n3\r\n4\r\n")
	})
}

func Test_SplitAndTrim(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := `

010    

020  

`
		a := gstr.X分割并忽略空值(s, "\n", "0")
		t.Assert(len(a), 2)
		t.Assert(a[0], "1")
		t.Assert(a[1], "2")
	})
}

func Test_Fields(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X单词分割("我爱 Go Frame"), []string{
			"我爱", "Go", "Frame",
		})
	})
}

func Test_CountWords(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X统计单词数("我爱 Go Go Go"), map[string]int{
			"Go": 3,
			"我爱": 1,
		})
	})
}

func Test_CountChars(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X统计unicode字符数("我爱 Go Go Go"), map[string]int{
			" ": 3,
			"G": 3,
			"o": 3,
			"我": 1,
			"爱": 1,
		})
		t.Assert(gstr.X统计unicode字符数("我爱 Go Go Go", true), map[string]int{
			"G": 3,
			"o": 3,
			"我": 1,
			"爱": 1,
		})
	})
}

func Test_LenRune(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X取字符长度("1234"), 4)
		t.Assert(gstr.X取字符长度("我爱GoFrame"), 9)
	})
}

func Test_Repeat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X生成重复文本("go", 3), "gogogo")
		t.Assert(gstr.X生成重复文本("好的", 3), "好的好的好的")
	})
}

func Test_Str(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X取右边并含关键字("name@example.com", "@"), "@example.com")
		t.Assert(gstr.X取右边并含关键字("name@example.com", ""), "")
		t.Assert(gstr.X取右边并含关键字("name@example.com", "z"), "")
	})
}

func Test_StrEx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X取右边("name@example.com", "@"), "example.com")
		t.Assert(gstr.X取右边("name@example.com", ""), "")
		t.Assert(gstr.X取右边("name@example.com", "z"), "")
	})
}

func Test_StrTill(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X取左边并含关键字("name@example.com", "@"), "name@")
		t.Assert(gstr.X取左边并含关键字("name@example.com", ""), "")
		t.Assert(gstr.X取左边并含关键字("name@example.com", "z"), "")
	})
}

func Test_StrTillEx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X取左边("name@example.com", "@"), "name")
		t.Assert(gstr.X取左边("name@example.com", ""), "")
		t.Assert(gstr.X取左边("name@example.com", "z"), "")
	})
}

func Test_Shuffle(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(gstr.X随机打散字符("123456")), 6)
	})
}

func Test_Split(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X分割("1.2", "."), []string{"1", "2"})
		t.Assert(gstr.X分割("我爱 - GoFrame", " - "), []string{"我爱", "GoFrame"})
	})
}

func Test_Join(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X连接([]string{"我爱", "GoFrame"}, " - "), "我爱 - GoFrame")
	})
}

func Test_Explode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.Explode别名(" - ", "我爱 - GoFrame"), []string{"我爱", "GoFrame"})
	})
}

func Test_Implode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.Implode别名(" - ", []string{"我爱", "GoFrame"}), "我爱 - GoFrame")
	})
}

func Test_Chr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X整数到ascii(65), "A")
	})
}

func Test_Ord(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.Ord("A"), 65)
	})
}

func Test_HideStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X替换中间字符("15928008611", 40, "*"), "159****8611")
		t.Assert(gstr.X替换中间字符("john@kohg.cn", 40, "*"), "jo*n@kohg.cn")
		t.Assert(gstr.X替换中间字符("张三", 50, "*"), "张*")
		t.Assert(gstr.X替换中间字符("张小三", 50, "*"), "张*三")
		t.Assert(gstr.X替换中间字符("欧阳小三", 50, "*"), "欧**三")
	})
}

func Test_Nl2Br(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X替换换行符("1\n2"), "1<br>2")
		t.Assert(gstr.X替换换行符("1\r\n2"), "1<br>2")
		t.Assert(gstr.X替换换行符("1\r\n2", true), "1<br />2")
	})
}

func Test_AddSlashes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X转义(`1'2"3\`), `1\'2\"3\\`)
	})
}

func Test_StripSlashes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X转义还原(`1\'2\"3\\`), `1'2"3\`)
	})
}

func Test_QuoteMeta(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X转义并按字符(`.\+*?[^]($)`), `\.\\\+\*\?\[\^\]\(\$\)`)
		t.Assert(gstr.X转义并按字符(`.\+*中国?[^]($)`), `\.\\\+\*中国\?\[\^\]\(\$\)`)
		t.Assert(gstr.X转义并按字符(`.''`, `'`), `.\'\'`)
		t.Assert(gstr.X转义并按字符(`中国.''`, `'`), `中国.\'\'`)
	})
}

func Test_Count(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "abcdaAD"
		t.Assert(gstr.X统计次数(s, "0"), 0)
		t.Assert(gstr.X统计次数(s, "a"), 2)
		t.Assert(gstr.X统计次数(s, "b"), 1)
		t.Assert(gstr.X统计次数(s, "d"), 1)
	})
}

func Test_CountI(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "abcdaAD"
		t.Assert(gstr.X统计次数并忽略大小写(s, "0"), 0)
		t.Assert(gstr.X统计次数并忽略大小写(s, "a"), 3)
		t.Assert(gstr.X统计次数并忽略大小写(s, "b"), 1)
		t.Assert(gstr.X统计次数并忽略大小写(s, "d"), 2)
	})
}

func Test_Compare(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X顺序比较("a", "b"), -1)
		t.Assert(gstr.X顺序比较("a", "a"), 0)
		t.Assert(gstr.X顺序比较("b", "a"), 1)
	})
}

func Test_Equal(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X相等比较并忽略大小写("a", "A"), true)
		t.Assert(gstr.X相等比较并忽略大小写("a", "a"), true)
		t.Assert(gstr.X相等比较并忽略大小写("b", "a"), false)
	})
}

func Test_Contains(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X是否包含("abc", "a"), true)
		t.Assert(gstr.X是否包含("abc", "A"), false)
		t.Assert(gstr.X是否包含("abc", "ab"), true)
		t.Assert(gstr.X是否包含("abc", "abc"), true)
	})
}

func Test_ContainsI(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X是否包含并忽略大小写("abc", "a"), true)
		t.Assert(gstr.X是否包含并忽略大小写("abc", "A"), true)
		t.Assert(gstr.X是否包含并忽略大小写("abc", "Ab"), true)
		t.Assert(gstr.X是否包含并忽略大小写("abc", "ABC"), true)
		t.Assert(gstr.X是否包含并忽略大小写("abc", "ABCD"), false)
		t.Assert(gstr.X是否包含并忽略大小写("abc", "D"), false)
	})
}

func Test_ContainsAny(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X是否包含Any("abc", "a"), true)
		t.Assert(gstr.X是否包含Any("abc", "cd"), true)
		t.Assert(gstr.X是否包含Any("abc", "de"), false)
		t.Assert(gstr.X是否包含Any("abc", "A"), false)
	})
}

func Test_SubStrFrom(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.SubStrFrom别名("我爱GoFrameGood", `G`), "GoFrameGood")
		t.Assert(gstr.SubStrFrom别名("我爱GoFrameGood", `GG`), "")
		t.Assert(gstr.SubStrFrom别名("我爱GoFrameGood", `我`), "我爱GoFrameGood")
		t.Assert(gstr.SubStrFrom别名("我爱GoFrameGood", `Frame`), "FrameGood")
	})
}

func Test_SubStrFromEx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.SubStrFromEx别名("我爱GoFrameGood", `Go`), "FrameGood")
		t.Assert(gstr.SubStrFromEx别名("我爱GoFrameGood", `GG`), "")
		t.Assert(gstr.SubStrFromEx别名("我爱GoFrameGood", `我`), "爱GoFrameGood")
		t.Assert(gstr.SubStrFromEx别名("我爱GoFrameGood", `Frame`), `Good`)
	})
}

func Test_SubStrFromR(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X取右边并倒找与含关键字("我爱GoFrameGood", `G`), "Good")
		t.Assert(gstr.X取右边并倒找与含关键字("我爱GoFrameGood", `GG`), "")
		t.Assert(gstr.X取右边并倒找与含关键字("我爱GoFrameGood", `我`), "我爱GoFrameGood")
		t.Assert(gstr.X取右边并倒找与含关键字("我爱GoFrameGood", `Frame`), "FrameGood")
	})
}

func Test_SubStrFromREx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X取右边并倒找("我爱GoFrameGood", `G`), "ood")
		t.Assert(gstr.X取右边并倒找("我爱GoFrameGood", `GG`), "")
		t.Assert(gstr.X取右边并倒找("我爱GoFrameGood", `我`), "爱GoFrameGood")
		t.Assert(gstr.X取右边并倒找("我爱GoFrameGood", `Frame`), `Good`)
	})
}
