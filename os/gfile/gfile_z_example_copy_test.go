// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile_test
import (
	"fmt"
	
	"github.com/888go/goframe/os/gfile"
	)

func ExampleCopy() {
	// init
	var (
		srcFileName = "gflie_example.txt"
		srcTempDir  = gfile.Temp("gfile_example_copy_src")
		srcTempFile = gfile.Join(srcTempDir, srcFileName)

		// copy file
		dstFileName = "gflie_example_copy.txt"
		dstTempFile = gfile.Join(srcTempDir, dstFileName)

		// copy dir
		dstTempDir = gfile.Temp("gfile_example_copy_dst")
	)

	// write contents
	gfile.PutContents(srcTempFile, "goframe example copy")

	// copy file
	gfile.Copy(srcTempFile, dstTempFile)

	// 在复制文件后读取内容
	fmt.Println(gfile.GetContents(dstTempFile))

	// copy dir
	gfile.Copy(srcTempDir, dstTempDir)

	// 列表复制目录文件
	fList, _ := gfile.ScanDir(dstTempDir, "*", false)
	for _, v := range fList {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// goframe example copy
	// gflie_example.txt
	// gflie_example_copy.txt
}
