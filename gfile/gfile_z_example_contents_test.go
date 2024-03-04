// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile_test

import (
	"fmt"
	
	"github.com/888go/goframe/gfile"
)

func ExampleGetContents() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

// 它读取并以字符串形式返回文件内容。
// 如果读取失败（例如，由于权限或IO错误），则返回空字符串。
	fmt.Println(gfile.GetContents(tempFile))

	// Output:
	// goframe example content
}

func ExampleGetBytes() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

// 它读取并以 []byte 类型返回文件内容。
// 如果在读取过程中出现错误（如权限错误或IO错误），则返回 nil。
	fmt.Println(gfile.GetBytes(tempFile))

	// Output:
	// [103 111 102 114 97 109 101 32 101 120 97 109 112 108 101 32 99 111 110 116 101 110 116]
}

func ExamplePutContents() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

// 它创建并将内容字符串写入指定的文件路径。
// 如果目录不存在，它会自动递归地创建目录。
	gfile.PutContents(tempFile, "goframe example content")

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

	// Output:
	// goframe example content
}

func ExamplePutBytes() {
	// init
	var (
		fileName = "gflie_example.txt"
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
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_content")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

// 它创建并把内容字符串追加到指定的文件路径中。
// 如果目录不存在，它会自动递归地创建目录。
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
		fileName = "gflie_example.txt"
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
		fileName = "gflie_example.txt"
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
		fileName = "gflie_example.txt"
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
		fileName = "gflie_example.txt"
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
		fileName = "gflie_example.txt"
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
		fileName = "gflie_example.txt"
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
