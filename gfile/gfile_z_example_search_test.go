// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile_test

import (
	"fmt"
	
	"github.com/888go/goframe/gfile"
)

func ExampleSearch() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_search")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

	// search file
	realPath, _ := gfile.Search(fileName, tempDir)
	fmt.Println(gfile.Basename(realPath))

	// Output:
	// gflie_example.txt
}
