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

func ExampleSize() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_size")
		tempFile = gfile.X路径生成(tempDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "0123456789")
	fmt.Println(gfile.X取大小(tempFile))

	// Output:
	// 10
}

func ExampleSizeFormat() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_size")
		tempFile = gfile.X路径生成(tempDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "0123456789")
	fmt.Println(gfile.X取大小并易读格式(tempFile))

	// Output:
	// 10.00B
}

func ExampleReadableSize() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_size")
		tempFile = gfile.X路径生成(tempDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "01234567899876543210")
	fmt.Println(gfile.ReadableSize别名(tempFile))

	// Output:
	// 20.00B
}

func ExampleStrToSize() {
	size := gfile.X易读格式转字节长度("100MB")
	fmt.Println(size)

	// Output:
	// 104857600
}

func ExampleFormatSize() {
	sizeStr := gfile.X字节长度转易读格式(104857600)
	fmt.Println(sizeStr)
	sizeStr0 := gfile.X字节长度转易读格式(1024)
	fmt.Println(sizeStr0)
	sizeStr1 := gfile.X字节长度转易读格式(999999999999999999)
	fmt.Println(sizeStr1)

	// Output:
	// 100.00M
	// 1.00K
	// 888.18P
}
