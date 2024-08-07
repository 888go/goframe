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

func ExampleCopy() {
	// init
	var (
		srcFileName = "gfile_example.txt"
		srcTempDir  = gfile.X取临时目录("gfile_example_copy_src")
		srcTempFile = gfile.X路径生成(srcTempDir, srcFileName)

		// copy file
		dstFileName = "gfile_example_copy.txt"
		dstTempFile = gfile.X路径生成(srcTempDir, dstFileName)

		// copy dir
		dstTempDir = gfile.X取临时目录("gfile_example_copy_dst")
	)

	// write contents
	gfile.X写入文本(srcTempFile, "goframe example copy")

	// copy file
	gfile.X复制(srcTempFile, dstTempFile)

			// 读取复制文件后的内容. md5:1645f6025a796e52
	fmt.Println(gfile.X读文本(dstTempFile))

	// copy dir
	gfile.X复制(srcTempDir, dstTempDir)

	// list copy dir file
	fList, _ := gfile.X枚举并含子目录名(dstTempDir, "*", false)
	for _, v := range fList {
		fmt.Println(gfile.X路径取文件名(v))
	}

	// Output:
	// goframe example copy
	// gfile_example.txt
	// gfile_example_copy.txt
}
