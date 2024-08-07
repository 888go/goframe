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

func ExampleSearch() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_search")
		tempFile = gfile.X路径生成(tempDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "goframe example content")

	// search file
	realPath, _ := gfile.X查找(fileName, tempDir)
	fmt.Println(gfile.X路径取文件名(realPath))

	// Output:
	// gfile_example.txt
}
