// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"fmt"
	"regexp"
	
	"github.com/888go/goframe/os/gfile"
)

func ExampleReplaceFile() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = 文件类.X取临时目录("gfile_example_replace")
		tempFile = 文件类.X路径生成(tempDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "goframe example content")

	// read contents
	fmt.Println(文件类.X读文本(tempFile))

	// 它通过文件路径直接替换内容。
	文件类.X子文本替换("content", "replace word", tempFile)

	fmt.Println(文件类.X读文本(tempFile))

	// Output:
	// goframe example content
	// goframe example replace word
}

func ExampleReplaceFileFunc() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = 文件类.X取临时目录("gfile_example_replace")
		tempFile = 文件类.X路径生成(tempDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "goframe example 123")

	// read contents
	fmt.Println(文件类.X读文本(tempFile))

	// 它通过文件路径和回调函数直接替换内容。
	文件类.X子文本替换_函数(func(path, content string) string {
		// 用普通匹配替换
		reg, _ := regexp.Compile(`\d{3}`)
		return reg.ReplaceAllString(content, "[num]")
	}, tempFile)

	fmt.Println(文件类.X读文本(tempFile))

	// Output:
	// goframe example 123
	// goframe example [num]
}

func ExampleReplaceDir() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = 文件类.X取临时目录("gfile_example_replace")
		tempFile = 文件类.X路径生成(tempDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "goframe example content")

	// read contents
	fmt.Println(文件类.X读文本(tempFile))

	// 它递归地替换指定目录下所有文件的内容。
	文件类.X目录子文本替换("content", "replace word", tempDir, "gflie_example.txt", true)

	// read contents
	fmt.Println(文件类.X读文本(tempFile))

	// Output:
	// goframe example content
	// goframe example replace word
}

func ExampleReplaceDirFunc() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = 文件类.X取临时目录("gfile_example_replace")
		tempFile = 文件类.X路径生成(tempDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "goframe example 123")

	// read contents
	fmt.Println(文件类.X读文本(tempFile))

	// 它会递归地替换指定目录下所有文件的内容，使用自定义的回调函数进行替换。
	文件类.X目录子文本替换_函数(func(path, content string) string {
		// 用普通匹配替换
		reg, _ := regexp.Compile(`\d{3}`)
		return reg.ReplaceAllString(content, "[num]")
	}, tempDir, "gflie_example.txt", true)

	fmt.Println(文件类.X读文本(tempFile))

	// Output:
	// goframe example 123
	// goframe example [num]

}
