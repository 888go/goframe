// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类_test

import (
	"fmt"

	gfile "github.com/888go/goframe/os/gfile"
)

func ExampleGetContents() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

	// 它读取文件内容并返回字符串形式。
	// 如果读取失败，例如权限或IO错误，它将返回空字符串。
	// md5:47f28c87cbca1824
	fmt.Println(gfile.GetContents(tempFile))

	// Output:
	// goframe example content
}

func ExampleGetBytes() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

	// 它读取文件内容并返回一个字节切片。
	// 如果在读取时出现错误，如权限或IO错误，它将返回nil。
	// md5:5cdf3501c0f95f5e
	fmt.Println(gfile.GetBytes(tempFile))

	// Output:
	// [103 111 102 114 97 109 101 32 101 120 97 109 112 108 101 32 99 111 110 116 101 110 116]
}

func ExamplePutContents() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// 它创建并将内容字符串放入指定的文件路径中。
	// 如果目录不存在，它会自动递归创建。
	// md5:ed9205edb3fc637b
	gfile.PutContents(tempFile, "goframe example content")

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

	// Output:
	// goframe example content
}

func ExamplePutBytes() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutBytes(tempFile, []byte("goframe example content"))

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

	// Output:
	// goframe example content
}

func ExamplePutContentsAppend() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

	// 该函数创建并追加内容字符串到指定的文件路径。
	// 如果指定的目录不存在，它会自动递归创建。
	// md5:bc62171e0c6aaf77
	gfile.PutContentsAppend(tempFile, " append content")

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

	// Output:
	// goframe example content
	// goframe example content append content
}

func ExamplePutBytesAppend() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

	// write contents
	gfile.PutBytesAppend(tempFile, []byte(" append"))

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

	// Output:
	// goframe example content
	// goframe example content append
}

func ExampleGetNextCharOffsetByPath() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

	// read contents
	index := gfile.GetNextCharOffsetByPath(tempFile, 'f', 0)
	fmt.Println(index)

	// Output:
	// 2
}

func ExampleGetBytesTilCharByPath() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

	// read contents
	fmt.Println(gfile.GetBytesTilCharByPath(tempFile, 'f', 0))

	// Output:
	// [103 111 102] 2
}

func ExampleGetBytesByTwoOffsetsByPath() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

	// read contents
	fmt.Println(gfile.GetBytesByTwoOffsetsByPath(tempFile, 0, 7))

	// Output:
	// [103 111 102 114 97 109 101]
}

func ExampleReadLines() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "L1 goframe example content\nL2 goframe example content")

	// read contents
	gfile.ReadLines(tempFile, func(text string) error {
		// Process each line
		fmt.Println(text)
		return nil
	})

	// Output:
	// L1 goframe example content
	// L2 goframe example content
}

func ExampleReadLinesBytes() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "L1 goframe example content\nL2 goframe example content")

	// read contents
	gfile.ReadLinesBytes(tempFile, func(bytes []byte) error {
		// Process each line
		fmt.Println(bytes)
		return nil
	})

	// Output:
	// [76 49 32 103 111 102 114 97 109 101 32 101 120 97 109 112 108 101 32 99 111 110 116 101 110 116]
	// [76 50 32 103 111 102 114 97 109 101 32 101 120 97 109 112 108 101 32 99 111 110 116 101 110 116]
}
