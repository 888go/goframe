// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile_test

import (
	"fmt"
	
	"github.com/888go/goframe/os/gfile"
)

func ExampleSize() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_size")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "0123456789")
	fmt.Println(gfile.Size(tempFile))

	// Output:
	// 10
}

func ExampleSizeFormat() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_size")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "0123456789")
	fmt.Println(gfile.SizeFormat(tempFile))

	// Output:
	// 10.00B
}

func ExampleReadableSize() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_size")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "01234567899876543210")
	fmt.Println(gfile.ReadableSize(tempFile))

	// Output:
	// 20.00B
}

func ExampleStrToSize() {
	size := gfile.StrToSize("100MB")
	fmt.Println(size)

	// Output:
	// 104857600
}

func ExampleFormatSize() {
	sizeStr := gfile.FormatSize(104857600)
	fmt.Println(sizeStr)
	sizeStr0 := gfile.FormatSize(1024)
	fmt.Println(sizeStr0)
	sizeStr1 := gfile.FormatSize(999999999999999999)
	fmt.Println(sizeStr1)

	// Output:
	// 100.00M
	// 1.00K
	// 888.18P
}
