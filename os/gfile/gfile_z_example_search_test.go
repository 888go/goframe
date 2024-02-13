// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"fmt"
	
	"github.com/888go/goframe/os/gfile"
)

func ExampleSearch() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = 文件类.X取临时目录("gfile_example_search")
		tempFile = 文件类.X路径生成(tempDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "goframe example content")

	// search file
	realPath, _ := 文件类.X查找(fileName, tempDir)
	fmt.Println(文件类.X路径取文件名(realPath))

	// Output:
	// gflie_example.txt
}
