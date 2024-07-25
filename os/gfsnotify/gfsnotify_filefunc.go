// 版权声明：GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码受MIT许可证条款约束。如果此文件未附带MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。 md5:12b80d680e9de440

package gfsnotify

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
)

// fileDir 返回路径除最后一个元素之外的所有内容，通常为路径的目录。
// 在丢弃最后一个元素后，Dir 函数会对路径调用 Clean 方法，并移除尾随斜杠。
// 如果路径为空，Dir 返回"."。
// 如果路径仅由分隔符组成，Dir 返回单个分隔符。
// 返回的路径除非是根目录，否则不会以分隔符结尾。 md5:c4932e9f21542326
func fileDir(path string) string {
	return filepath.Dir(path)
}

// fileRealPath 将给定的 `path` 转换为其绝对路径，并检查文件路径是否存在。
// 如果文件不存在，返回空字符串。 md5:e30bc7542e1c332e
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

// fileExists 检查给定的 `path` 是否存在。 md5:88ff2b38709b04ab
func fileExists(path string) bool {
	if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// fileIsDir 检查给定的 `path` 是否为目录。 md5:e16eca80a5a42f13
func fileIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// fileAllDirs递归地返回给定`path`的所有子文件夹，包括自身。 md5:b6638d72b44fee31
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

// fileScanDir 返回给定 `path` 所有子文件的绝对路径，
// 如果参数 `recursive` 为 true，则递归扫描目录。 md5:871aa4d8d78c9ee4
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

// doFileScanDir 是一个内部方法，用于扫描目录
// 并返回未排序的文件的绝对路径列表。
//
// 模式参数 `pattern` 支持多个文件名模式，
// 使用 `,` 符号分隔多个模式。
//
// 如果给定的参数 `recursive` 为 true，它将递归地扫描目录。 md5:187616f6d86800cf
func doFileScanDir(path string, pattern string, recursive ...bool) ([]string, error) {
	var (
		list      []string
		file, err = os.Open(path)
	)
	if err != nil {
		err = gerror.Wrapf(err, `os.Open failed for path "%s"`, path)
		return nil, err
	}
	defer file.Close()
	names, err := file.Readdirnames(-1)
	if err != nil {
		err = gerror.Wrapf(err, `read directory files failed for path "%s"`, path)
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
