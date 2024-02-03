// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile_test

import (
	"fmt"
	
	"github.com/888go/goframe/os/gfile"
)

func ExampleScanDir() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_scan_dir")
		tempFile = gfile.Join(tempDir, fileName)

		tempSubDir  = gfile.Join(tempDir, "sub_dir")
		tempSubFile = gfile.Join(tempSubDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")
	gfile.PutContents(tempSubFile, "goframe example content")

	// 递归扫描目录
	list, _ := gfile.ScanDir(tempDir, "*", true)
	for _, v := range list {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// gflie_example.txt
	// sub_dir
	// gflie_example.txt
}

func ExampleScanDirFile() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_scan_dir_file")
		tempFile = gfile.Join(tempDir, fileName)

		tempSubDir  = gfile.Join(tempDir, "sub_dir")
		tempSubFile = gfile.Join(tempSubDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")
	gfile.PutContents(tempSubFile, "goframe example content")

	// 递归扫描目录 exclusive of directories
	list, _ := gfile.ScanDirFile(tempDir, "*.txt", true)
	for _, v := range list {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// gflie_example.txt
	// gflie_example.txt
}

func ExampleScanDirFunc() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_scan_dir_func")
		tempFile = gfile.Join(tempDir, fileName)

		tempSubDir  = gfile.Join(tempDir, "sub_dir")
		tempSubFile = gfile.Join(tempSubDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")
	gfile.PutContents(tempSubFile, "goframe example content")

	// 递归扫描目录
	list, _ := gfile.ScanDirFunc(tempDir, "*", true, func(path string) string {
		// 忽略某些文件
		if gfile.Basename(path) == "gflie_example.txt" {
			return ""
		}
		return path
	})
	for _, v := range list {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// sub_dir
}

func ExampleScanDirFileFunc() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_scan_dir_file_func")
		tempFile = gfile.Join(tempDir, fileName)

		fileName1 = "gflie_example_ignores.txt"
		tempFile1 = gfile.Join(tempDir, fileName1)

		tempSubDir  = gfile.Join(tempDir, "sub_dir")
		tempSubFile = gfile.Join(tempSubDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")
	gfile.PutContents(tempFile1, "goframe example content")
	gfile.PutContents(tempSubFile, "goframe example content")

	// 递归扫描目录 exclusive of directories
	list, _ := gfile.ScanDirFileFunc(tempDir, "*.txt", true, func(path string) string {
		// 忽略某些文件
		if gfile.Basename(path) == "gflie_example_ignores.txt" {
			return ""
		}
		return path
	})
	for _, v := range list {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// gflie_example.txt
	// gflie_example.txt
}
