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

func ExampleScanDir() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_scan_dir")
		tempFile = gfile.X路径生成(tempDir, fileName)

		tempSubDir  = gfile.X路径生成(tempDir, "sub_dir")
		tempSubFile = gfile.X路径生成(tempSubDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "goframe example content")
	gfile.X写入文本(tempSubFile, "goframe example content")

		// 递归扫描目录. md5:6afcf6167831bfa3
	list, _ := gfile.X枚举并含子目录名(tempDir, "*", true)
	for _, v := range list {
		fmt.Println(gfile.X路径取文件名(v))
	}

	// Output:
	// gfile_example.txt
	// sub_dir
	// gfile_example.txt
}

func ExampleScanDirFile() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_scan_dir_file")
		tempFile = gfile.X路径生成(tempDir, fileName)

		tempSubDir  = gfile.X路径生成(tempDir, "sub_dir")
		tempSubFile = gfile.X路径生成(tempSubDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "goframe example content")
	gfile.X写入文本(tempSubFile, "goframe example content")

			// 递归扫描目录. md5:6afcf6167831bfa3
	list, _ := gfile.X枚举(tempDir, "*.txt", true)
	for _, v := range list {
		fmt.Println(gfile.X路径取文件名(v))
	}

	// Output:
	// gfile_example.txt
	// gfile_example.txt
}

func ExampleScanDirFunc() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_scan_dir_func")
		tempFile = gfile.X路径生成(tempDir, fileName)

		tempSubDir  = gfile.X路径生成(tempDir, "sub_dir")
		tempSubFile = gfile.X路径生成(tempSubDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "goframe example content")
	gfile.X写入文本(tempSubFile, "goframe example content")

		// 递归扫描目录. md5:6afcf6167831bfa3
	list, _ := gfile.X枚举并含子目录名_函数(tempDir, "*", true, func(path string) string {
		// ignores some files
		if gfile.X路径取文件名(path) == "gfile_example.txt" {
			return ""
		}
		return path
	})
	for _, v := range list {
		fmt.Println(gfile.X路径取文件名(v))
	}

	// Output:
	// sub_dir
}

func ExampleScanDirFileFunc() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_scan_dir_file_func")
		tempFile = gfile.X路径生成(tempDir, fileName)

		fileName1 = "gfile_example_ignores.txt"
		tempFile1 = gfile.X路径生成(tempDir, fileName1)

		tempSubDir  = gfile.X路径生成(tempDir, "sub_dir")
		tempSubFile = gfile.X路径生成(tempSubDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "goframe example content")
	gfile.X写入文本(tempFile1, "goframe example content")
	gfile.X写入文本(tempSubFile, "goframe example content")

			// 递归扫描目录. md5:6afcf6167831bfa3
	list, _ := gfile.X枚举_函数(tempDir, "*.txt", true, func(path string) string {
		// ignores some files
		if gfile.X路径取文件名(path) == "gfile_example_ignores.txt" {
			return ""
		}
		return path
	})
	for _, v := range list {
		fmt.Println(gfile.X路径取文件名(v))
	}

	// Output:
	// gfile_example.txt
	// gfile_example.txt
}
