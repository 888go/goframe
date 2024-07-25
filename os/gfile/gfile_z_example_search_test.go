// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gfile_test

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gfile"
)

func ExampleSearch() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_search")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

	// search file
	realPath, _ := gfile.Search(fileName, tempDir)
	fmt.Println(gfile.Basename(realPath))

	// Output:
	// gfile_example.txt
}
