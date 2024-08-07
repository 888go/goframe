// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
package 正则类_test

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/888go/goframe/frame/g"
	gregex "github.com/888go/goframe/text/gregex"
)

func ExampleIsMatch() {
	patternStr := `\d+`
	g.X调试输出(gregex.X是否匹配字节集(patternStr, []byte("hello 2022! hello gf!")))
	g.X调试输出(gregex.X是否匹配字节集(patternStr, nil))
	g.X调试输出(gregex.X是否匹配字节集(patternStr, []byte("hello gf!")))

	// Output:
	// true
	// false
	// false
}

func ExampleIsMatchString() {
	patternStr := `\d+`
	g.X调试输出(gregex.X是否匹配文本(patternStr, "hello 2022! hello gf!"))
	g.X调试输出(gregex.X是否匹配文本(patternStr, "hello gf!"))
	g.X调试输出(gregex.X是否匹配文本(patternStr, ""))

	// Output:
	// true
	// false
	// false
}

func ExampleMatch() {
	patternStr := `(\w+)=(\w+)`
	matchStr := "https://goframe.org/pages/viewpage.action?pageId=1114219&searchId=8QC5D1D2E!"
			// 这个方法寻找第一个匹配的索引. md5:d933ee1d74dd86c5
	result, err := gregex.X匹配字节集(patternStr, []byte(matchStr))
	g.X调试输出(result)
	g.X调试输出(err)

	// Output:
	// [
	//     "pageId=1114219",
	//     "pageId",
	//     "1114219",
	// ]
	// <nil>
}

func ExampleMatchString() {
	patternStr := `(\w+)=(\w+)`
	matchStr := "https://goframe.org/pages/viewpage.action?pageId=1114219&searchId=8QC5D1D2E!"
			// 这个方法寻找第一个匹配的索引. md5:d933ee1d74dd86c5
	result, err := gregex.X匹配文本(patternStr, matchStr)
	g.X调试输出(result)
	g.X调试输出(err)

	// Output:
	// [
	//     "pageId=1114219",
	//     "pageId",
	//     "1114219",
	// ]
	// <nil>
}

func ExampleMatchAll() {
	patternStr := `(\w+)=(\w+)`
	matchStr := "https://goframe.org/pages/viewpage.action?pageId=1114219&searchId=8QC5D1D2E!"
	result, err := gregex.X匹配全部字节集(patternStr, []byte(matchStr))
	g.X调试输出(result)
	g.X调试输出(err)

	// Output:
	//  [
	//     [
	//         "pageId=1114219",
	//         "pageId",
	//         "1114219",
	//     ],
	//     [
	//         "searchId=8QC5D1D2E",
	//         "searchId",
	//         "8QC5D1D2E",
	//     ],
	// ]
	// <nil>
}

func ExampleMatchAllString() {
	patternStr := `(\w+)=(\w+)`
	matchStr := "https://goframe.org/pages/viewpage.action?pageId=1114219&searchId=8QC5D1D2E!"
	result, err := gregex.X匹配全部文本(patternStr, matchStr)
	g.X调试输出(result)
	g.X调试输出(err)

	// Output:
	// [
	//     [
	//         "pageId=1114219",
	//         "pageId",
	//         "1114219",
	//     ],
	//     [
	//         "searchId=8QC5D1D2E",
	//         "searchId",
	//         "8QC5D1D2E",
	//     ],
	// ]
	// <nil>
}

func ExampleQuote() {
	result := gregex.X转义特殊符号(`[1-9]\d+`)
	fmt.Println(result)

	// Output:
	// \[1-9\]\\d\+
}

func ExampleReplace() {
	var (
		patternStr  = `\d+`
		str         = "hello gf 2020!"
		repStr      = "2021"
		result, err = gregex.X替换字节集(patternStr, []byte(repStr), []byte(str))
	)
	g.X调试输出(err)
	g.X调试输出(result)

	// Output:
	// <nil>
	// "hello gf 2021!"
}

func ExampleReplaceFunc() {
	// 与 [ExampleReplaceFunc] 相比，
	// 结果包含了所有使用匹配函数的子模式的`pattern`
	// md5:3cc683990c37065c
	result, err := gregex.ReplaceFuncMatch(`(\d+)~(\d+)`, []byte("hello gf 2018~2020!"), func(match [][]byte) []byte {
		g.X调试输出(match)
		match[2] = []byte("2021")
		return bytes.Join(match[1:], []byte("~"))
	})
	g.X调试输出(result)
	g.X调试输出(err)

	// Output:
	// [
	//     "2018~2020",
	//     "2018",
	//     "2020",
	// ]
	// "hello gf 2018~2021!"
	// <nil>
}

func ExampleReplaceFuncMatch() {
	var (
		patternStr = `(\d+)~(\d+)`
		str        = "hello gf 2018~2020!"
	)
	// 与[ExampleReplaceFunc]不同的是，
	// 结果包含了使用匹配函数的所有子模式的`pattern`。
	// md5:1b711898b19df13d
	result, err := gregex.ReplaceFuncMatch(patternStr, []byte(str), func(match [][]byte) []byte {
		g.X调试输出(match)
		match[2] = []byte("2021")
		return bytes.Join(match[1:], []byte("-"))
	})
	g.X调试输出(result)
	g.X调试输出(err)

	// Output:
	// [
	//     "2018~2020",
	//     "2018",
	//     "2020",
	// ]
	// "hello gf 2018-2021!"
	// <nil>
}

func ExampleReplaceString() {
	patternStr := `\d+`
	str := "hello gf 2020!"
	replaceStr := "2021"
	result, err := gregex.X替换文本(patternStr, replaceStr, str)

	g.X调试输出(result)
	g.X调试输出(err)

	// Output:
	// "hello gf 2021!"
	// <nil>
}

func ExampleReplaceStringFunc() {
	replaceStrMap := map[string]string{
		"2020": "2021",
	}
	// 当常规语句可以匹配多个结果时
	// 函数可以进一步控制需要修改的值
	// md5:453f5d05c5806c71
	result, err := gregex.X替换文本_函数(`\d+`, `hello gf 2018~2020!`, func(b string) string {
		g.X调试输出(b)
		if replaceStr, ok := replaceStrMap[b]; ok {
			return replaceStr
		}
		return b
	})
	g.X调试输出(result)
	g.X调试输出(err)

	result, err = gregex.X替换文本_函数(`[a-z]*`, "gf@goframe.org", strings.ToUpper)
	g.X调试输出(result)
	g.X调试输出(err)

	// Output:
	// "2018"
	// "2020"
	// "hello gf 2018~2021!"
	// <nil>
	// "GF@GOFRAME.ORG"
	// <nil>
}

func ExampleReplaceStringFuncMatch() {
	var (
		patternStr = `([A-Z])\w+`
		str        = "hello Golang 2018~2021!"
	)
	// 与[ExampleReplaceFunc]不同的是，
	// 结果包含了使用匹配函数的所有子模式的`pattern`。
	// md5:1b711898b19df13d
	result, err := gregex.ReplaceStringFuncMatch(patternStr, str, func(match []string) string {
		g.X调试输出(match)
		match[0] = "Gf"
		return match[0]
	})
	g.X调试输出(result)
	g.X调试输出(err)

	// Output:
	// [
	//     "Golang",
	//     "G",
	// ]
	// "hello Gf 2018~2021!"
	// <nil>
}

func ExampleSplit() {
	patternStr := `\d+`
	str := "hello2020gf"
	result := gregex.X分割(patternStr, str)
	g.X调试输出(result)

	// Output:
	// [
	//     "hello",
	//     "gf",
	// ]
}

func ExampleValidate() {
	// Valid match statement
	fmt.Println(gregex.X表达式验证(`\d+`))
	// Mismatched statement
	fmt.Println(gregex.X表达式验证(`[a-9]\d+`))

	// Output:
	// <nil>
	// regexp.Compile failed for pattern "[a-9]\d+": error parsing regexp: invalid character class range: `a-9`
}
