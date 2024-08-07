// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类_test

import (
	"fmt"
	"regexp"

	gfile "github.com/888go/goframe/os/gfile"
)

func ExampleReplaceFile() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_replace")
		tempFile = gfile.X路径生成(tempDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "goframe example content")

	// read contents
	fmt.Println(gfile.X读文本(tempFile))

		// 它通过文件路径直接替换内容。 md5:2a3205bdc3de2657
	gfile.X子文本替换("content", "replace word", tempFile)

	fmt.Println(gfile.X读文本(tempFile))

	// Output:
	// goframe example content
	// goframe example replace word
}

func ExampleReplaceFileFunc() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_replace")
		tempFile = gfile.X路径生成(tempDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "goframe example 123")

	// read contents
	fmt.Println(gfile.X读文本(tempFile))

		// 它通过文件路径和回调函数直接替换内容。 md5:7962223be9a9a643
	gfile.X子文本替换_函数(func(path, content string) string {
						// 使用常规匹配替换. md5:30a3741f5800de5e
		reg, _ := regexp.Compile(`\d{3}`)
		return reg.ReplaceAllString(content, "[num]")
	}, tempFile)

	fmt.Println(gfile.X读文本(tempFile))

	// Output:
	// goframe example 123
	// goframe example [num]
}

func ExampleReplaceDir() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_replace")
		tempFile = gfile.X路径生成(tempDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "goframe example content")

	// read contents
	fmt.Println(gfile.X读文本(tempFile))

		// 它会递归地替换指定目录下所有文件的内容。 md5:20439a8528d54108
	gfile.X目录子文本替换("content", "replace word", tempDir, "gfile_example.txt", true)

	// read contents
	fmt.Println(gfile.X读文本(tempFile))

	// Output:
	// goframe example content
	// goframe example replace word
}

func ExampleReplaceDirFunc() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_replace")
		tempFile = gfile.X路径生成(tempDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "goframe example 123")

	// read contents
	fmt.Println(gfile.X读文本(tempFile))

		// 它会递归地用自定义回调函数替换指定目录下所有文件的内容。 md5:9186cb76d9407085
	gfile.X目录子文本替换_函数(func(path, content string) string {
						// 使用常规匹配替换. md5:30a3741f5800de5e
		reg, _ := regexp.Compile(`\d{3}`)
		return reg.ReplaceAllString(content, "[num]")
	}, tempDir, "gfile_example.txt", true)

	fmt.Println(gfile.X读文本(tempFile))

	// Output:
	// goframe example 123
	// goframe example [num]

}
