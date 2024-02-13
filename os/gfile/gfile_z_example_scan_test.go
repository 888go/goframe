// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"fmt"
	
	"github.com/888go/goframe/os/gfile"
)

func ExampleScanDir() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = 文件类.X取临时目录("gfile_example_scan_dir")
		tempFile = 文件类.X路径生成(tempDir, fileName)

		tempSubDir  = 文件类.X路径生成(tempDir, "sub_dir")
		tempSubFile = 文件类.X路径生成(tempSubDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "goframe example content")
	文件类.X写入文本(tempSubFile, "goframe example content")

	// 递归扫描目录
	list, _ := 文件类.X枚举并含子目录名(tempDir, "*", true)
	for _, v := range list {
		fmt.Println(文件类.X路径取文件名(v))
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
		tempDir  = 文件类.X取临时目录("gfile_example_scan_dir_file")
		tempFile = 文件类.X路径生成(tempDir, fileName)

		tempSubDir  = 文件类.X路径生成(tempDir, "sub_dir")
		tempSubFile = 文件类.X路径生成(tempSubDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "goframe example content")
	文件类.X写入文本(tempSubFile, "goframe example content")

	// 递归扫描目录 exclusive of directories
	list, _ := 文件类.X枚举(tempDir, "*.txt", true)
	for _, v := range list {
		fmt.Println(文件类.X路径取文件名(v))
	}

	// Output:
	// gflie_example.txt
	// gflie_example.txt
}

func ExampleScanDirFunc() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = 文件类.X取临时目录("gfile_example_scan_dir_func")
		tempFile = 文件类.X路径生成(tempDir, fileName)

		tempSubDir  = 文件类.X路径生成(tempDir, "sub_dir")
		tempSubFile = 文件类.X路径生成(tempSubDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "goframe example content")
	文件类.X写入文本(tempSubFile, "goframe example content")

	// 递归扫描目录
	list, _ := 文件类.X枚举并含子目录名_函数(tempDir, "*", true, func(path string) string {
		// 忽略某些文件
		if 文件类.X路径取文件名(path) == "gflie_example.txt" {
			return ""
		}
		return path
	})
	for _, v := range list {
		fmt.Println(文件类.X路径取文件名(v))
	}

	// Output:
	// sub_dir
}

func ExampleScanDirFileFunc() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = 文件类.X取临时目录("gfile_example_scan_dir_file_func")
		tempFile = 文件类.X路径生成(tempDir, fileName)

		fileName1 = "gflie_example_ignores.txt"
		tempFile1 = 文件类.X路径生成(tempDir, fileName1)

		tempSubDir  = 文件类.X路径生成(tempDir, "sub_dir")
		tempSubFile = 文件类.X路径生成(tempSubDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "goframe example content")
	文件类.X写入文本(tempFile1, "goframe example content")
	文件类.X写入文本(tempSubFile, "goframe example content")

	// 递归扫描目录 exclusive of directories
	list, _ := 文件类.X枚举_函数(tempDir, "*.txt", true, func(path string) string {
		// 忽略某些文件
		if 文件类.X路径取文件名(path) == "gflie_example_ignores.txt" {
			return ""
		}
		return path
	})
	for _, v := range list {
		fmt.Println(文件类.X路径取文件名(v))
	}

	// Output:
	// gflie_example.txt
	// gflie_example.txt
}
