// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类_test

import (
	"fmt"

	gstr "github.com/888go/goframe/text/gstr"
)

func ExampleCount() {
	var (
		str     = `goframe is very, very easy to use`
		substr1 = "goframe"
		substr2 = "very"
		result1 = gstr.X统计次数(str, substr1)
		result2 = gstr.X统计次数(str, substr2)
	)
	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 1
	// 2
}

func ExampleCountI() {
	var (
		str     = `goframe is very, very easy to use`
		substr1 = "GOFRAME"
		substr2 = "VERY"
		result1 = gstr.X统计次数并忽略大小写(str, substr1)
		result2 = gstr.X统计次数并忽略大小写(str, substr2)
	)
	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 1
	// 2
}

func ExampleToLower() {
	var (
		s      = `GOFRAME`
		result = gstr.X到小写(s)
	)
	fmt.Println(result)

	// Output:
	// goframe
}

func ExampleToUpper() {
	var (
		s      = `goframe`
		result = gstr.X到大写(s)
	)
	fmt.Println(result)

	// Output:
	// GOFRAME
}

func ExampleUcFirst() {
	var (
		s      = `hello`
		result = gstr.X到首字母大写(s)
	)
	fmt.Println(result)

	// Output:
	// Hello
}

func ExampleLcFirst() {
	var (
		str    = `Goframe`
		result = gstr.X到首字母小写(str)
	)
	fmt.Println(result)

	// Output:
	// goframe
}

func ExampleUcWords() {
	var (
		str    = `hello world`
		result = gstr.X到单词首字母大写(str)
	)
	fmt.Println(result)

	// Output:
	// Hello World
}

func ExampleIsLetterLower() {
	fmt.Println(gstr.X是否小写字符('a'))
	fmt.Println(gstr.X是否小写字符('A'))

	// Output:
	// true
	// false
}

func ExampleIsLetterUpper() {
	fmt.Println(gstr.X是否大写字符('A'))
	fmt.Println(gstr.X是否大写字符('a'))

	// Output:
	// true
	// false
}

func ExampleIsNumeric() {
	fmt.Println(gstr.X是否为数字("88"))
	fmt.Println(gstr.X是否为数字("3.1415926"))
	fmt.Println(gstr.X是否为数字("abc"))
	// Output:
	// true
	// true
	// false
}

func ExampleReverse() {
	var (
		str    = `123456`
		result = gstr.X反转字符(str)
	)
	fmt.Println(result)

	// Output:
	// 654321
}

func ExampleNumberFormat() {
	var (
		number       float64 = 123456
		decimals             = 2
		decPoint             = "."
		thousandsSep         = ","
		result               = gstr.X格式化数值(number, decimals, decPoint, thousandsSep)
	)
	fmt.Println(result)

	// Output:
	// 123,456.00
}

func ExampleChunkSplit() {
	var (
		body     = `1234567890`
		chunkLen = 2
		end      = "#"
		result   = gstr.X长度分割(body, chunkLen, end)
	)
	fmt.Println(result)

	// Output:
	// 12#34#56#78#90#
}

func ExampleCompare() {
	fmt.Println(gstr.X顺序比较("c", "c"))
	fmt.Println(gstr.X顺序比较("a", "b"))
	fmt.Println(gstr.X顺序比较("c", "b"))

	// Output:
	// 0
	// -1
	// 1
}

func ExampleEqual() {
	fmt.Println(gstr.X相等比较并忽略大小写(`A`, `a`))
	fmt.Println(gstr.X相等比较并忽略大小写(`A`, `A`))
	fmt.Println(gstr.X相等比较并忽略大小写(`A`, `B`))

	// Output:
	// true
	// true
	// false
}

func ExampleFields() {
	var (
		str    = `Hello World`
		result = gstr.X单词分割(str)
	)
	fmt.Printf(`%#v`, result)

	// Output:
	// []string{"Hello", "World"}
}

func ExampleHasPrefix() {
	var (
		s      = `Hello World`
		prefix = "Hello"
		result = gstr.X开头判断(s, prefix)
	)
	fmt.Println(result)

	// Output:
	// true
}

func ExampleHasSuffix() {
	var (
		s      = `my best love is goframe`
		prefix = "goframe"
		result = gstr.X末尾判断(s, prefix)
	)
	fmt.Println(result)

	// Output:
	// true
}

func ExampleCountWords() {
	var (
		str    = `goframe is very, very easy to use!`
		result = gstr.X统计单词数(str)
	)
	fmt.Printf(`%#v`, result)

	// Output:
	// map[string]int{"easy":1, "goframe":1, "is":1, "to":1, "use!":1, "very":1, "very,":1}
}

func ExampleCountChars() {
	var (
		str    = `goframe`
		result = gstr.X统计unicode字符数(str)
	)
	fmt.Println(result)

	// May Output:
	// map[a:1 e:1 f:1 g:1 m:1 o:1 r:1]
}

func ExampleWordWrap() {
	{
		var (
			str    = `A very long woooooooooooooooooord. and something`
			width  = 8
			br     = "\n"
			result = gstr.X按字符数量换行(str, width, br)
		)
		fmt.Println(result)
	}
	{
		var (
			str    = `The quick brown fox jumped over the lazy dog.`
			width  = 20
			br     = "<br />\n"
			result = gstr.X按字符数量换行(str, width, br)
		)
		fmt.Printf("%v", result)
	}

	// Output:
	// A very
	// long
	// woooooooooooooooooord.
	// and
	// something
	// The quick brown fox<br />
	// jumped over the lazy<br />
	// dog.
}

func ExampleLenRune() {
	var (
		str    = `GoFrame框架`
		result = gstr.X取字符长度(str)
	)
	fmt.Println(result)

	// Output:
	// 9
}

func ExampleRepeat() {
	var (
		input      = `goframe `
		multiplier = 3
		result     = gstr.X生成重复文本(input, multiplier)
	)
	fmt.Println(result)

	// Output:
	// goframe goframe goframe
}

func ExampleShuffle() {
	var (
		str    = `123456`
		result = gstr.X随机打散字符(str)
	)
	fmt.Println(result)

	// May Output:
	// 563214
}

func ExampleSplit() {
	var (
		str       = `a|b|c|d`
		delimiter = `|`
		result    = gstr.X分割(str, delimiter)
	)
	fmt.Printf(`%#v`, result)

	// Output:
	// []string{"a", "b", "c", "d"}
}

func ExampleSplitAndTrim() {
	var (
		str       = `a|b|||||c|d`
		delimiter = `|`
		result    = gstr.X分割并忽略空值(str, delimiter)
	)
	fmt.Printf(`%#v`, result)

	// Output:
	// []string{"a", "b", "c", "d"}
}

func ExampleJoin() {
	var (
		array  = []string{"goframe", "is", "very", "easy", "to", "use"}
		sep    = ` `
		result = gstr.X连接(array, sep)
	)
	fmt.Println(result)

	// Output:
	// goframe is very easy to use
}

func ExampleJoinAny() {
	var (
		sep    = `,`
		arr2   = []int{99, 73, 85, 66}
		result = gstr.X连接Any(arr2, sep)
	)
	fmt.Println(result)

	// Output:
	// 99,73,85,66
}

func ExampleExplode() {
	var (
		str       = `Hello World`
		delimiter = " "
		result    = gstr.Explode别名(delimiter, str)
	)
	fmt.Printf(`%#v`, result)

	// Output:
	// []string{"Hello", "World"}
}

func ExampleImplode() {
	var (
		pieces = []string{"goframe", "is", "very", "easy", "to", "use"}
		glue   = " "
		result = gstr.Implode别名(glue, pieces)
	)
	fmt.Println(result)

	// Output:
	// goframe is very easy to use
}

func ExampleChr() {
	var (
		ascii  = 65 // A
		result = gstr.X整数到ascii(ascii)
	)
	fmt.Println(result)

	// Output:
	// A
}

// '103' 在ASCII码中代表字母 'g'. md5:e664c0cbb63cb67e
func ExampleOrd() {
	var (
		str    = `goframe`
		result = gstr.Ord(str)
	)

	fmt.Println(result)

	// Output:
	// 103
}

func ExampleHideStr() {
	var (
		str     = `13800138000`
		percent = 40
		hide    = `*`
		result  = gstr.X替换中间字符(str, percent, hide)
	)
	fmt.Println(result)

	// Output:
	// 138****8000
}

func ExampleNl2Br() {
	var (
		str = `goframe
is
very
easy
to
use`
		result = gstr.X替换换行符(str)
	)

	fmt.Println(result)

	// Output:
	// goframe<br>is<br>very<br>easy<br>to<br>use
}

func ExampleAddSlashes() {
	var (
		str    = `'aa'"bb"cc\r\n\d\t`
		result = gstr.X转义(str)
	)

	fmt.Println(result)

	// Output:
	// \'aa\'\"bb\"cc\\r\\n\\d\\t
}

func ExampleStripSlashes() {
	var (
		str    = `C:\\windows\\GoFrame\\test`
		result = gstr.X转义还原(str)
	)
	fmt.Println(result)

	// Output:
	// C:\windows\GoFrame\test
}

func ExampleQuoteMeta() {
	{
		var (
			str    = `.\+?[^]()`
			result = gstr.X转义并按字符(str)
		)
		fmt.Println(result)
	}
	{
		var (
			str    = `https://goframe.org/pages/viewpage.action?pageId=1114327`
			result = gstr.X转义并按字符(str)
		)
		fmt.Println(result)
	}

	// Output:
	// \.\\\+\?\[\^\]\(\)
	// https://goframe\.org/pages/viewpage\.action\?pageId=1114327

}

// array
func ExampleSearchArray() {
	var (
		array  = []string{"goframe", "is", "very", "nice"}
		str    = `goframe`
		result = gstr.X切片查找(array, str)
	)
	fmt.Println(result)

	// Output:
	// 0
}

func ExampleInArray() {
	var (
		a      = []string{"goframe", "is", "very", "easy", "to", "use"}
		s      = "goframe"
		result = gstr.X切片是否存在(a, s)
	)
	fmt.Println(result)

	// Output:
	// true
}

func ExamplePrefixArray() {
	var (
		strArray = []string{"tom", "lily", "john"}
	)

	gstr.X切片加前缀(strArray, "classA_")

	fmt.Println(strArray)

	// Output:
	// [classA_tom classA_lily classA_john]
}

// case
func ExampleCaseCamel() {
	var (
		str    = `hello world`
		result = gstr.X命名转换到首字母大写驼峰(str)
	)
	fmt.Println(result)

	// Output:
	// HelloWorld
}

func ExampleCaseCamelLower() {
	var (
		str    = `hello world`
		result = gstr.X命名转换到首字母小写驼峰(str)
	)
	fmt.Println(result)

	// Output:
	// helloWorld
}

func ExampleCaseSnake() {
	var (
		str    = `hello world`
		result = gstr.X命名转换到全小写蛇形(str)
	)
	fmt.Println(result)

	// Output:
	// hello_world
}

func ExampleCaseSnakeScreaming() {
	var (
		str    = `hello world`
		result = gstr.X命名转换到大写蛇形(str)
	)
	fmt.Println(result)

	// Output:
	// HELLO_WORLD
}

func ExampleCaseSnakeFirstUpper() {
	var (
		str    = `RGBCodeMd5`
		result = gstr.X命名转换到全小写蛇形2(str)
	)
	fmt.Println(result)

	// Output:
	// rgb_code_md5
}

func ExampleCaseKebab() {
	var (
		str    = `hello world`
		result = gstr.X命名转换到小写短横线(str)
	)
	fmt.Println(result)

	// Output:
	// hello-world
}

func ExampleCaseKebabScreaming() {
	var (
		str    = `hello world`
		result = gstr.X命名转换到大写驼峰短横线(str)
	)
	fmt.Println(result)

	// Output:
	// HELLO-WORLD
}

func ExampleCaseDelimited() {
	var (
		str    = `hello world`
		del    = byte('-')
		result = gstr.X命名转换按符号(str, del)
	)
	fmt.Println(result)

	// Output:
	// hello-world
}

func ExampleCaseDelimitedScreaming() {
	{
		var (
			str    = `hello world`
			del    = byte('-')
			result = gstr.X命名转换按符号与大小写(str, del, true)
		)
		fmt.Println(result)
	}
	{
		var (
			str    = `hello world`
			del    = byte('-')
			result = gstr.X命名转换按符号与大小写(str, del, false)
		)
		fmt.Println(result)
	}

	// Output:
	// HELLO-WORLD
	// hello-world
}

// contain
func ExampleContains() {
	{
		var (
			str    = `Hello World`
			substr = `Hello`
			result = gstr.X是否包含(str, substr)
		)
		fmt.Println(result)
	}
	{
		var (
			str    = `Hello World`
			substr = `hello`
			result = gstr.X是否包含(str, substr)
		)
		fmt.Println(result)
	}

	// Output:
	// true
	// false
}

func ExampleContainsI() {
	var (
		str     = `Hello World`
		substr  = "hello"
		result1 = gstr.X是否包含(str, substr)
		result2 = gstr.X是否包含并忽略大小写(str, substr)
	)
	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// false
	// true
}

func ExampleContainsAny() {
	{
		var (
			s      = `goframe`
			chars  = "g"
			result = gstr.X是否包含Any(s, chars)
		)
		fmt.Println(result)
	}
	{
		var (
			s      = `goframe`
			chars  = "G"
			result = gstr.X是否包含Any(s, chars)
		)
		fmt.Println(result)
	}

	// Output:
	// true
	// false
}

// convert
func ExampleOctStr() {
	var (
		str    = `\346\200\241`
		result = gstr.X八进制到文本(str)
	)
	fmt.Println(result)

	// Output:
	// 怡
}

// domain
func ExampleIsSubDomain() {
	var (
		subDomain  = `s.goframe.org`
		mainDomain = `goframe.org`
		result     = gstr.X是否为子域名(subDomain, mainDomain)
	)
	fmt.Println(result)

	// Output:
	// true
}

// levenshtein
func ExampleLevenshtein() {
	var (
		str1    = "Hello World"
		str2    = "hallo World"
		costIns = 1
		costRep = 1
		costDel = 1
		result  = gstr.Levenshtein(str1, str2, costIns, costRep, costDel)
	)
	fmt.Println(result)

	// Output:
	// 2
}

// parse
func ExampleParse() {
	{
		var (
			str       = `v1=m&v2=n`
			result, _ = gstr.X参数解析(str)
		)
		fmt.Println(result)
	}
	{
		var (
			str       = `v[a][a]=m&v[a][b]=n`
			result, _ = gstr.X参数解析(str)
		)
		fmt.Println(result)
	}
	{
				// 遍历嵌套切片的格式尚未得到支持。 md5:4ba6b8e024a1bae0
		var str = `v[][]=m&v[][]=n`
		result, err := gstr.X参数解析(str)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
	{
				// 这将产生一个错误。 md5:a04ca00c0d3445fa
		var str = `v=m&v[a]=n`
		result, err := gstr.X参数解析(str)
		if err != nil {
			println(err)
		}
		fmt.Println(result)
	}
	{
		var (
			str       = `a .[[b=c`
			result, _ = gstr.X参数解析(str)
		)
		fmt.Println(result)
	}

	// May Output:
	// map[v1:m v2:n]
	// map[v:map[a:map[a:m b:n]]]
	// map[v:map[]]
	// Error: expected type 'map[string]interface{}' for key 'v', but got 'string'
	// map[]
	// map[a___[b:c]
}

// pos
func ExamplePos() {
	var (
		haystack = `Hello World`
		needle   = `World`
		result   = gstr.X查找(haystack, needle)
	)
	fmt.Println(result)

	// Output:
	// 6
}

func ExamplePosRune() {
	var (
		haystack = `GoFrame是一款模块化、高性能、企业级的Go基础开发框架`
		needle   = `Go`
		posI     = gstr.X查找Unicode(haystack, needle)
		posR     = gstr.X倒找Unicode(haystack, needle)
	)
	fmt.Println(posI)
	fmt.Println(posR)

	// Output:
	// 0
	// 22
}

func ExamplePosI() {
	var (
		haystack = `goframe is very, very easy to use`
		needle   = `very`
		posI     = gstr.X查找并忽略大小写(haystack, needle)
		posR     = gstr.X倒找(haystack, needle)
	)
	fmt.Println(posI)
	fmt.Println(posR)

	// Output:
	// 11
	// 17
}

func ExamplePosIRune() {
	{
		var (
			haystack    = `GoFrame是一款模块化、高性能、企业级的Go基础开发框架`
			needle      = `高性能`
			startOffset = 10
			result      = gstr.X查找并忽略大小写Unicode(haystack, needle, startOffset)
		)
		fmt.Println(result)
	}
	{
		var (
			haystack    = `GoFrame是一款模块化、高性能、企业级的Go基础开发框架`
			needle      = `高性能`
			startOffset = 30
			result      = gstr.X查找并忽略大小写Unicode(haystack, needle, startOffset)
		)
		fmt.Println(result)
	}

	// Output:
	// 14
	// -1
}

func ExamplePosR() {
	var (
		haystack = `goframe is very, very easy to use`
		needle   = `very`
		posI     = gstr.X查找并忽略大小写(haystack, needle)
		posR     = gstr.X倒找(haystack, needle)
	)
	fmt.Println(posI)
	fmt.Println(posR)

	// Output:
	// 11
	// 17
}

func ExamplePosRRune() {
	var (
		haystack = `GoFrame是一款模块化、高性能、企业级的Go基础开发框架`
		needle   = `Go`
		posI     = gstr.X查找并忽略大小写Unicode(haystack, needle)
		posR     = gstr.X倒找Unicode(haystack, needle)
	)
	fmt.Println(posI)
	fmt.Println(posR)

	// Output:
	// 0
	// 22
}

func ExamplePosRI() {
	var (
		haystack = `goframe is very, very easy to use`
		needle   = `VERY`
		posI     = gstr.X查找并忽略大小写(haystack, needle)
		posR     = gstr.X倒找并忽略大小写(haystack, needle)
	)
	fmt.Println(posI)
	fmt.Println(posR)

	// Output:
	// 11
	// 17
}

func ExamplePosRIRune() {
	var (
		haystack = `GoFrame是一款模块化、高性能、企业级的Go基础开发框架`
		needle   = `GO`
		posI     = gstr.X查找并忽略大小写Unicode(haystack, needle)
		posR     = gstr.X倒找并忽略大小写Unicode(haystack, needle)
	)
	fmt.Println(posI)
	fmt.Println(posR)

	// Output:
	// 0
	// 22
}

// replace
func ExampleReplace() {
	var (
		origin  = `golang is very nice!`
		search  = `golang`
		replace = `goframe`
		result  = gstr.X替换(origin, search, replace)
	)
	fmt.Println(result)

	// Output:
	// goframe is very nice!
}

func ExampleReplaceI() {
	var (
		origin  = `golang is very nice!`
		search  = `GOLANG`
		replace = `goframe`
		result  = gstr.X替换并忽略大小写(origin, search, replace)
	)
	fmt.Println(result)

	// Output:
	// goframe is very nice!
}

func ExampleReplaceByArray() {
	{
		var (
			origin = `golang is very nice`
			array  = []string{"lang", "frame"}
			result = gstr.X切片替换(origin, array)
		)
		fmt.Println(result)
	}
	{
		var (
			origin = `golang is very good`
			array  = []string{"golang", "goframe", "good", "nice"}
			result = gstr.X切片替换(origin, array)
		)
		fmt.Println(result)
	}

	// Output:
	// goframe is very nice
	// goframe is very nice
}

func ExampleReplaceIByArray() {
	var (
		origin = `golang is very Good`
		array  = []string{"Golang", "goframe", "GOOD", "nice"}
		result = gstr.X切片替换并忽略大小写(origin, array)
	)

	fmt.Println(result)

	// Output:
	// goframe is very nice
}

func ExampleReplaceByMap() {
	{
		var (
			origin   = `golang is very nice`
			replaces = map[string]string{
				"lang": "frame",
			}
			result = gstr.Map替换(origin, replaces)
		)
		fmt.Println(result)
	}
	{
		var (
			origin   = `golang is very good`
			replaces = map[string]string{
				"golang": "goframe",
				"good":   "nice",
			}
			result = gstr.Map替换(origin, replaces)
		)
		fmt.Println(result)
	}

	// Output:
	// goframe is very nice
	// goframe is very nice
}

func ExampleReplaceIByMap() {
	var (
		origin   = `golang is very nice`
		replaces = map[string]string{
			"Lang": "frame",
		}
		result = gstr.Map替换并忽略大小写(origin, replaces)
	)
	fmt.Println(result)

	// Output:
	// goframe is very nice
}

// similartext
func ExampleSimilarText() {
	var (
		first   = `AaBbCcDd`
		second  = `ad`
		percent = 0.80
		result  = gstr.X取相似度(first, second, &percent)
	)
	fmt.Println(result)

	// Output:
	// 2
}

// soundex
func ExampleSoundex() {
	var (
		str1    = `Hello`
		str2    = `Hallo`
		result1 = gstr.X取soundex码(str1)
		result2 = gstr.X取soundex码(str2)
	)
	fmt.Println(result1, result2)

	// Output:
	// H400 H400
}

// str
func ExampleStr() {
	var (
		haystack = `xxx.jpg`
		needle   = `.`
		result   = gstr.X取右边并含关键字(haystack, needle)
	)
	fmt.Println(result)

	// Output:
	// .jpg
}

func ExampleStrEx() {
	var (
		haystack = `https://goframe.org/index.html?a=1&b=2`
		needle   = `?`
		result   = gstr.X取右边(haystack, needle)
	)
	fmt.Println(result)

	// Output:
	// a=1&b=2
}

func ExampleStrTill() {
	var (
		haystack = `https://goframe.org/index.html?test=123456`
		needle   = `?`
		result   = gstr.X取左边并含关键字(haystack, needle)
	)
	fmt.Println(result)

	// Output:
	// https://goframe.org/index.html?
}

func ExampleStrTillEx() {
	var (
		haystack = `https://goframe.org/index.html?test=123456`
		needle   = `?`
		result   = gstr.X取左边(haystack, needle)
	)
	fmt.Println(result)

	// Output:
	// https://goframe.org/index.html
}

// substr
func ExampleSubStr() {
	var (
		str    = `1234567890`
		start  = 0
		length = 4
		subStr = gstr.X按长度取文本(str, start, length)
	)
	fmt.Println(subStr)

	// Output:
	// 1234
}

func ExampleSubStrRune() {
	var (
		str    = `GoFrame是一款模块化、高性能、企业级的Go基础开发框架。`
		start  = 14
		length = 3
		subStr = gstr.X按长度取文本Unicode(str, start, length)
	)
	fmt.Println(subStr)

	// Output:
	// 高性能
}

func ExampleStrLimit() {
	var (
		str    = `123456789`
		length = 3
		suffix = `...`
		result = gstr.X按长度取左边并带前缀(str, length, suffix)
	)
	fmt.Println(result)

	// Output:
	// 123...
}

func ExampleStrLimitRune() {
	var (
		str    = `GoFrame是一款模块化、高性能、企业级的Go基础开发框架。`
		length = 17
		suffix = "..."
		result = gstr.X按长度取左边并带前缀Unicode(str, length, suffix)
	)
	fmt.Println(result)

	// Output:
	// GoFrame是一款模块化、高性能...
}

func ExampleSubStrFrom() {
	var (
		str  = "我爱GoFrameGood"
		need = `爱`
	)

	fmt.Println(gstr.SubStrFrom别名(str, need))

	// Output:
	// 爱GoFrameGood
}

func ExampleSubStrFromEx() {
	var (
		str  = "我爱GoFrameGood"
		need = `爱`
	)

	fmt.Println(gstr.SubStrFromEx别名(str, need))

	// Output:
	// GoFrameGood
}

func ExampleSubStrFromR() {
	var (
		str  = "我爱GoFrameGood"
		need = `Go`
	)

	fmt.Println(gstr.X取右边并倒找与含关键字(str, need))

	// Output:
	// Good
}

func ExampleSubStrFromREx() {
	var (
		str  = "我爱GoFrameGood"
		need = `Go`
	)

	fmt.Println(gstr.X取右边并倒找(str, need))

	// Output:
	// od
}

// trim
func ExampleTrim() {
	var (
		str           = `*Hello World*`
		characterMask = "*"
		result        = gstr.X过滤首尾符并含空白(str, characterMask)
	)
	fmt.Println(result)

	// Output:
	// Hello World
}

func ExampleTrimStr() {
	var (
		str    = `Hello World`
		cut    = "World"
		count  = -1
		result = gstr.X过滤首尾(str, cut, count)
	)
	fmt.Println(result)

	// Output:
	// Hello
}

func ExampleTrimLeft() {
	var (
		str           = `*Hello World*`
		characterMask = "*"
		result        = gstr.X过滤首字符并含空白(str, characterMask)
	)
	fmt.Println(result)

	// Output:
	// Hello World*
}

func ExampleTrimLeftStr() {
	var (
		str    = `**Hello World**`
		cut    = "*"
		count  = 1
		result = gstr.X过滤首字符(str, cut, count)
	)
	fmt.Println(result)

	// Output:
	// *Hello World**
}

func ExampleTrimRight() {
	var (
		str           = `**Hello World**`
		characterMask = "*def" // 一个字节切片，包含以下元素：星号（`*`），字母 d，字母 e，和字母 f。 md5:c33233e03897a326
		result        = gstr.X过滤尾字符并含空白(str, characterMask)
	)
	fmt.Println(result)

	// Output:
	// **Hello Worl
}

func ExampleTrimRightStr() {
	var (
		str    = `Hello World!`
		cut    = "!"
		count  = -1
		result = gstr.X过滤尾字符(str, cut, count)
	)
	fmt.Println(result)

	// Output:
	// Hello World
}

func ExampleTrimAll() {
	var (
		str           = `*Hello World*`
		characterMask = "*"
		result        = gstr.X过滤所有字符并含空白(str, characterMask)
	)
	fmt.Println(result)

	// Output:
	// HelloWorld
}

// version
func ExampleCompareVersion() {
	fmt.Println(gstr.X版本号比较GNU格式("v2.11.9", "v2.10.8"))
	fmt.Println(gstr.X版本号比较GNU格式("1.10.8", "1.19.7"))
	fmt.Println(gstr.X版本号比较GNU格式("2.8.beta", "2.8"))

	// Output:
	// 1
	// -1
	// 0
}

func ExampleCompareVersionGo() {
	fmt.Println(gstr.X版本号比较GO格式("v2.11.9", "v2.10.8"))
	fmt.Println(gstr.X版本号比较GO格式("v4.20.1", "v4.20.1+incompatible"))
	fmt.Println(gstr.X版本号比较GO格式(
		"v0.0.2-20180626092158-b2ccc119800e",
		"v1.0.1-20190626092158-b2ccc519800e",
	))

	// Output:
	// 1
	// 1
	// -1
}
