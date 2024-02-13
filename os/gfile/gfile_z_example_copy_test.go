// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"fmt"
	
	"github.com/888go/goframe/os/gfile"
)

func ExampleCopy() {
	// init
	var (
		srcFileName = "gflie_example.txt"
		srcTempDir  = 文件类.X取临时目录("gfile_example_copy_src")
		srcTempFile = 文件类.X路径生成(srcTempDir, srcFileName)

		// copy file
		dstFileName = "gflie_example_copy.txt"
		dstTempFile = 文件类.X路径生成(srcTempDir, dstFileName)

		// copy dir
		dstTempDir = 文件类.X取临时目录("gfile_example_copy_dst")
	)

	// write contents
	文件类.X写入文本(srcTempFile, "goframe example copy")

	// copy file
	文件类.X复制(srcTempFile, dstTempFile)

	// 在复制文件后读取内容
	fmt.Println(文件类.X读文本(dstTempFile))

	// copy dir
	文件类.X复制(srcTempDir, dstTempDir)

	// 列表复制目录文件
	fList, _ := 文件类.X枚举并含子目录名(dstTempDir, "*", false)
	for _, v := range fList {
		fmt.Println(文件类.X路径取文件名(v))
	}

	// Output:
	// goframe example copy
	// gflie_example.txt
	// gflie_example_copy.txt
}
