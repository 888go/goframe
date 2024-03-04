// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
package gregex_test

import (
	"bytes"
	"fmt"
	"strings"
	
	"github.com/gogf/gf/v2/frame/g"
	"github.com/888go/goframe/gregex"
)

func ExampleIsMatch() {
	patternStr := `\d+`
	g.Dump(gregex.IsMatch(patternStr, []byte("hello 2022! hello gf!")))
	g.Dump(gregex.IsMatch(patternStr, nil))
	g.Dump(gregex.IsMatch(patternStr, []byte("hello gf!")))

	// Output:
	// true
	// false
	// false
}

func ExampleIsMatchString() {
	patternStr := `\d+`
	g.Dump(gregex.IsMatchString(patternStr, "hello 2022! hello gf!"))
	g.Dump(gregex.IsMatchString(patternStr, "hello gf!"))
	g.Dump(gregex.IsMatchString(patternStr, ""))

	// Output:
	// true
	// false
	// false
}

func ExampleMatch() {
	patternStr := `(\w+)=(\w+)`
	matchStr := "https://goframe.org/pages/viewpage.action?pageId=1114219&searchId=8QC5D1D2E!"
	// 这个方法查找第一个匹配的索引
	result, err := gregex.Match(patternStr, []byte(matchStr))
	g.Dump(result)
	g.Dump(err)

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
	// 这个方法查找第一个匹配的索引
	result, err := gregex.MatchString(patternStr, matchStr)
	g.Dump(result)
	g.Dump(err)

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
	result, err := gregex.MatchAll(patternStr, []byte(matchStr))
	g.Dump(result)
	g.Dump(err)

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
	result, err := gregex.MatchAllString(patternStr, matchStr)
	g.Dump(result)
	g.Dump(err)

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
	result := gregex.Quote(`[1-9]\d+`)
	fmt.Println(result)

	// Output:
	// \[1-9\]\\d\+
}

func ExampleReplace() {
	var (
		patternStr  = `\d+`
		str         = "hello gf 2020!"
		repStr      = "2021"
		result, err = gregex.Replace(patternStr, []byte(repStr), []byte(str))
	)
	g.Dump(err)
	g.Dump(result)

	// Output:
	// <nil>
	// "hello gf 2021!"
}

func ExampleReplaceFunc() {
// 与[ExampleReplaceFunc]相反
// 结果包含所有使用匹配函数的子模式的`pattern`
	result, err := gregex.ReplaceFuncMatch(`(\d+)~(\d+)`, []byte("hello gf 2018~2020!"), func(match [][]byte) []byte {
		g.Dump(match)
		match[2] = []byte("2021")
		return bytes.Join(match[1:], []byte("~"))
	})
	g.Dump(result)
	g.Dump(err)

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
// 与 [ExampleReplaceFunc] 相反
// 结果包含所有使用匹配函数的子模式的 `pattern`
	result, err := gregex.ReplaceFuncMatch(patternStr, []byte(str), func(match [][]byte) []byte {
		g.Dump(match)
		match[2] = []byte("2021")
		return bytes.Join(match[1:], []byte("-"))
	})
	g.Dump(result)
	g.Dump(err)

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
	result, err := gregex.ReplaceString(patternStr, replaceStr, str)

	g.Dump(result)
	g.Dump(err)

	// Output:
	// "hello gf 2021!"
	// <nil>
}

func ExampleReplaceStringFunc() {
	replaceStrMap := map[string]string{
		"2020": "2021",
	}
// 当常规语句可以匹配多个结果时
// 可以使用func进行进一步控制，以便确定需要修改的值
	result, err := gregex.ReplaceStringFunc(`\d+`, `hello gf 2018~2020!`, func(b string) string {
		g.Dump(b)
		if replaceStr, ok := replaceStrMap[b]; ok {
			return replaceStr
		}
		return b
	})
	g.Dump(result)
	g.Dump(err)

	result, err = gregex.ReplaceStringFunc(`[a-z]*`, "gf@goframe.org", strings.ToUpper)
	g.Dump(result)
	g.Dump(err)

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
// 与 [ExampleReplaceFunc] 相反
// 结果包含所有使用匹配函数的子模式的 `pattern`
	result, err := gregex.ReplaceStringFuncMatch(patternStr, str, func(match []string) string {
		g.Dump(match)
		match[0] = "Gf"
		return match[0]
	})
	g.Dump(result)
	g.Dump(err)

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
	result := gregex.Split(patternStr, str)
	g.Dump(result)

	// Output:
	// [
	//     "hello",
	//     "gf",
	// ]
}

func ExampleValidate() {
	// 有效的匹配语句
	fmt.Println(gregex.Validate(`\d+`))
	// 不匹配的语句
	fmt.Println(gregex.Validate(`[a-9]\d+`))

	// Output:
	// <nil>
	// regexp.Compile failed for pattern "[a-9]\d+": error parsing regexp: invalid character class range: `a-9`
}
