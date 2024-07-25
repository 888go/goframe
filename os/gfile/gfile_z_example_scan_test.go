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

func ExampleScanDir() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_scan_dir")
		tempFile = gfile.Join(tempDir, fileName)

		tempSubDir  = gfile.Join(tempDir, "sub_dir")
		tempSubFile = gfile.Join(tempSubDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")
	gfile.PutContents(tempSubFile, "goframe example content")

		// 递归扫描目录. md5:6afcf6167831bfa3
	list, _ := gfile.ScanDir(tempDir, "*", true)
	for _, v := range list {
		fmt.Println(gfile.Basename(v))
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
		tempDir  = gfile.Temp("gfile_example_scan_dir_file")
		tempFile = gfile.Join(tempDir, fileName)

		tempSubDir  = gfile.Join(tempDir, "sub_dir")
		tempSubFile = gfile.Join(tempSubDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")
	gfile.PutContents(tempSubFile, "goframe example content")

			// 递归扫描目录. md5:6afcf6167831bfa3
	list, _ := gfile.ScanDirFile(tempDir, "*.txt", true)
	for _, v := range list {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// gfile_example.txt
	// gfile_example.txt
}

func ExampleScanDirFunc() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_scan_dir_func")
		tempFile = gfile.Join(tempDir, fileName)

		tempSubDir  = gfile.Join(tempDir, "sub_dir")
		tempSubFile = gfile.Join(tempSubDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")
	gfile.PutContents(tempSubFile, "goframe example content")

		// 递归扫描目录. md5:6afcf6167831bfa3
	list, _ := gfile.ScanDirFunc(tempDir, "*", true, func(path string) string {
		// ignores some files
		if gfile.Basename(path) == "gfile_example.txt" {
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
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_scan_dir_file_func")
		tempFile = gfile.Join(tempDir, fileName)

		fileName1 = "gfile_example_ignores.txt"
		tempFile1 = gfile.Join(tempDir, fileName1)

		tempSubDir  = gfile.Join(tempDir, "sub_dir")
		tempSubFile = gfile.Join(tempSubDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")
	gfile.PutContents(tempFile1, "goframe example content")
	gfile.PutContents(tempSubFile, "goframe example content")

			// 递归扫描目录. md5:6afcf6167831bfa3
	list, _ := gfile.ScanDirFileFunc(tempDir, "*.txt", true, func(path string) string {
		// ignores some files
		if gfile.Basename(path) == "gfile_example_ignores.txt" {
			return ""
		}
		return path
	})
	for _, v := range list {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// gfile_example.txt
	// gfile_example.txt
}
