// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"fmt"
	
	"github.com/888go/goframe/os/gfile"
)

func ExampleSize() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = 文件类.X取临时目录("gfile_example_size")
		tempFile = 文件类.X路径生成(tempDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "0123456789")
	fmt.Println(文件类.X取大小(tempFile))

	// Output:
	// 10
}

func ExampleSizeFormat() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = 文件类.X取临时目录("gfile_example_size")
		tempFile = 文件类.X路径生成(tempDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "0123456789")
	fmt.Println(文件类.X取大小并易读格式(tempFile))

	// Output:
	// 10.00B
}

func ExampleReadableSize() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = 文件类.X取临时目录("gfile_example_size")
		tempFile = 文件类.X路径生成(tempDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "01234567899876543210")
	fmt.Println(文件类.ReadableSize别名(tempFile))

	// Output:
	// 20.00B
}

func ExampleStrToSize() {
	size := 文件类.X易读格式转字节长度("100MB")
	fmt.Println(size)

	// Output:
	// 104857600
}

func ExampleFormatSize() {
	sizeStr := 文件类.X字节长度转易读格式(104857600)
	fmt.Println(sizeStr)
	sizeStr0 := 文件类.X字节长度转易读格式(1024)
	fmt.Println(sizeStr0)
	sizeStr1 := 文件类.X字节长度转易读格式(999999999999999999)
	fmt.Println(sizeStr1)

	// Output:
	// 100.00M
	// 1.00K
	// 888.18P
}
