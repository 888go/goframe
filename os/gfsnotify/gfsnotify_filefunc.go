// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT协议条款。如果随此文件未分发MIT协议副本，
// 您可以在https://github.com/gogf/gf获取一份。

package 文件监控类

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	
	"github.com/888go/goframe/errors/gerror"
)

// fileDir 返回路径中除最后一个元素之外的所有元素，通常是路径的目录部分。
// 在去掉最后一个元素后，Dir 会调用 Clean 来清理路径，并移除尾部的斜杠。
// 如果路径为空，Dir 返回 "."
// 如果路径完全由分隔符组成，Dir 返回单个分隔符。
// 返回的路径除非是根目录，否则不会以分隔符结尾。
// 此函数（fileDir）用于提取并返回给定路径的基本目录部分，在处理过程中会对路径进行规范化处理。具体规则如下：
// 1. 删除路径的最后一部分（通常为文件名或子目录名）。
// 2. 调用标准库中的 `Clean` 函数清理路径，去除末尾多余的斜杠。
// 3. 若路径为空，则返回当前目录（`.`）表示。
// 4. 若路径仅包含分隔符，则返回一个分隔符。
// 5. 返回的目录路径不以分隔符结尾，除非该路径指向的是根目录。
func fileDir(path string) string {
	return filepath.Dir(path)
}

// fileRealPath 将给定的 `path` 转换为绝对路径
// 并检查文件路径是否存在。
// 若文件不存在，则返回一个空字符串。
func fileRealPath(path string) string {
	p, err := filepath.Abs(path)
	if err != nil {
		return ""
	}
	if !fileExists(p) {
		return ""
	}
	return p
}

// fileExists 检查给定的 `path` 是否存在。
func fileExists(path string) bool {
	if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// fileIsDir 检查给定的 `path` 是否为一个目录。
func fileIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// fileAllDirs 返回给定 `path`（包括其自身）的所有子目录，递归遍历。
func fileAllDirs(path string) (list []string) {
	list = []string{path}
	file, err := os.Open(path)
	if err != nil {
		return list
	}
	defer file.Close()
	names, err := file.Readdirnames(-1)
	if err != nil {
		return list
	}
	for _, name := range names {
		tempPath := fmt.Sprintf("%s%s%s", path, string(filepath.Separator), name)
		if fileIsDir(tempPath) {
			if array := fileAllDirs(tempPath); len(array) > 0 {
				list = append(list, array...)
			}
		}
	}
	return
}

// fileScanDir 函数返回给定 `path` 下所有子文件的绝对路径，
// 如果给定参数 `recursive` 为 true，则会递归地扫描目录。
func fileScanDir(path string, pattern string, recursive ...bool) ([]string, error) {
	list, err := doFileScanDir(path, pattern, recursive...)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		sort.Strings(list)
	}
	return list, nil
}

// doFileScanDir 是一个内部方法，用于扫描目录并返回未排序的文件绝对路径列表。
//
// 参数`pattern`支持多个文件名模式，使用逗号 ',' 作为分隔符来指定多个模式。
//
// 如果给定的参数 `recursive` 为 true，则会递归地扫描目录。
func doFileScanDir(path string, pattern string, recursive ...bool) ([]string, error) {
	var (
		list      []string
		file, err = os.Open(path)
	)
	if err != nil {
		err = 错误类.X多层错误并格式化(err, `os.Open failed for path "%s"`, path)
		return nil, err
	}
	defer file.Close()
	names, err := file.Readdirnames(-1)
	if err != nil {
		err = 错误类.X多层错误并格式化(err, `read directory files failed for path "%s"`, path)
		return nil, err
	}
	filePath := ""
	for _, name := range names {
		filePath = fmt.Sprintf("%s%s%s", path, string(filepath.Separator), name)
		if fileIsDir(filePath) && len(recursive) > 0 && recursive[0] {
			array, _ := doFileScanDir(filePath, pattern, true)
			if len(array) > 0 {
				list = append(list, array...)
			}
		}
		for _, p := range strings.Split(pattern, ",") {
			if match, _ := filepath.Match(strings.TrimSpace(p), name); match {
				list = append(list, filePath)
			}
		}
	}
	return list, nil
}
