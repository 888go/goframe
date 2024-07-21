// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gfile_test

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gfile"
)

func ExampleCopy() {
	// init
	var (
		srcFileName = "gfile_example.txt"
		srcTempDir  = gfile.Temp("gfile_example_copy_src")
		srcTempFile = gfile.Join(srcTempDir, srcFileName)

		// copy file
		dstFileName = "gfile_example_copy.txt"
		dstTempFile = gfile.Join(srcTempDir, dstFileName)

		// copy dir
		dstTempDir = gfile.Temp("gfile_example_copy_dst")
	)

	// write contents
	gfile.PutContents(srcTempFile, "goframe example copy")

	// copy file
	gfile.Copy(srcTempFile, dstTempFile)

	// 读取复制文件后的内容. md5:1645f6025a796e52
	fmt.Println(gfile.GetContents(dstTempFile))

	// copy dir
	gfile.Copy(srcTempDir, dstTempDir)

	// list copy dir file
	fList, _ := gfile.ScanDir(dstTempDir, "*", false)
	for _, v := range fList {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// goframe example copy
	// gfile_example.txt
	// gfile_example_copy.txt
}
