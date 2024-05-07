// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类

import (
	"path/filepath"
	"sort"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
)

const (
	// 目录扫描的最大递归深度。
	maxScanDepth = 100000
)

// ScanDir 返回给定路径`path`下所有子文件的绝对路径，
// 若给定参数`recursive`为真，则会递归扫描目录。
//
// 参数`pattern`支持多种文件名模式，
// 可以使用逗号 ',' 作为分隔符来指定多个模式。
func X枚举并含子目录名(目录 string, 匹配文件模式 string, 是否递归替换 ...bool) ([]string, error) {
	isRecursive := false
	if len(是否递归替换) > 0 {
		isRecursive = 是否递归替换[0]
	}
	list, err := doScanDir(0, 目录, 匹配文件模式, isRecursive, nil)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		sort.Strings(list)
	}
	return list, nil
}

// ScanDirFunc 函数返回给定 `path` 下所有子文件的绝对路径，
// 如果给定参数 `recursive` 为 true，则会递归扫描目录。
//
// 参数 `pattern` 支持多个文件名模式，使用 ',' 符号分隔多个模式。
//
// 参数 `recursive` 指定是否递归扫描 `path`，这意味着如果子文件也是一个文件夹，它会扫描其下的子文件并将文件路径追加到结果切片中，默认情况下为 false。
//
// 参数 `handler` 指定了处理 `path` 及其子文件夹下每个子文件路径的回调函数。如果 `handler` 返回空字符串，则忽略该子文件路径，否则将其子文件路径追加到结果切片中。
func X枚举并含子目录名_函数(目录 string, 匹配文件模式 string, 是否递归替换 bool, 回调函数 func(路径 string) string) ([]string, error) {
	list, err := doScanDir(0, 目录, 匹配文件模式, 是否递归替换, 回调函数)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		sort.Strings(list)
	}
	return list, nil
}

// ScanDirFile 返回给定 `path` 下所有子文件的绝对路径，
// 如果给定参数 `recursive` 为 true，则会递归扫描目录。
//
// 参数 `pattern` 支持多个文件名模式，
// 使用 `,` 符号来分隔多个模式。
//
// 注意，它只返回文件，不包括目录。
func X枚举(目录 string, 匹配文件模式 string, 是否递归查找 ...bool) ([]string, error) {
	isRecursive := false
	if len(是否递归查找) > 0 {
		isRecursive = 是否递归查找[0]
	}
	list, err := doScanDir(0, 目录, 匹配文件模式, isRecursive, func(path string) string {
		if X是否存在目录(path) {
			return ""
		}
		return path
	})
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		sort.Strings(list)
	}
	return list, nil
}

// ScanDirFileFunc 函数返回给定路径 `path` 下所有子文件的绝对路径。
// 如果给定参数 `recursive` 为 true，则会递归扫描目录。
//
// 参数 `pattern` 支持多个文件名模式，使用 ',' 符号分隔多个模式。
//
// 参数 `recursive` 指定是否递归扫描 `path`，这意味着如果子文件也是一个文件夹，则会扫描其下级文件并将文件路径添加到结果切片中，默认情况下为 false。
//
// 参数 `handler` 指定了处理 `path` 及其子文件夹下每个子文件路径的回调函数。如果 `handler` 返回空字符串，则忽略该子文件路径，否则将子文件路径追加到结果切片中。
//
// 注意，传给 `handler` 的参数 `path` 不是一个目录而是一个文件。
// 此函数仅返回文件（不包括目录）。
func X枚举_函数(目录 string, 匹配文件模式 string, 是否递归查找 bool, 匿名函数 func(路径 string) string) ([]string, error) {
	list, err := doScanDir(0, 目录, 匹配文件模式, 是否递归查找, func(path string) string {
		if X是否存在目录(path) {
			return ""
		}
		return 匿名函数(path)
	})
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		sort.Strings(list)
	}
	return list, nil
}

// doScanDir 是一个内部方法，用于扫描目录并返回未排序的绝对路径文件列表。
// 参数 `pattern` 支持多个文件名模式，使用 ',' 符号分隔多个模式。
// 参数 `recursive` 指定是否递归扫描 `path`，这意味着它会扫描其子文件，如果子文件也是一个文件夹，则将子文件路径追加到结果切片中。默认情况下，recursive 为 false。
// 参数 `handler` 指定了处理 `path` 及其子文件夹下每个子文件路径的回调函数。如果 `handler` 返回空字符串，则忽略该子文件路径，否则将其追加到结果切片中。
func doScanDir(depth int, path string, pattern string, recursive bool, handler func(path string) string) ([]string, error) {
	if depth >= maxScanDepth {
		return nil, gerror.Newf("directory scanning exceeds max recursive depth: %d", maxScanDepth)
	}
	var (
		list      []string
		file, err = X打开并按只读模式(path)
	)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	names, err := file.Readdirnames(-1)
	if err != nil {
		err = gerror.Wrapf(err, `read directory files failed from path "%s"`, path)
		return nil, err
	}
	var (
		filePath string
		patterns = gstr.SplitAndTrim(pattern, ",")
	)
	for _, name := range names {
		filePath = path + Separator + name
		if X是否存在目录(filePath) && recursive {
			array, _ := doScanDir(depth+1, filePath, pattern, true, handler)
			if len(array) > 0 {
				list = append(list, array...)
			}
		}
		// 处理器筛选功能
		if handler != nil {
			filePath = handler(filePath)
			if filePath == "" {
				continue
			}
		}
		// 如果满足模式，则将其添加到结果列表中。
		for _, p := range patterns {
			if match, _ := filepath.Match(p, name); match {
				if filePath = X取绝对路径(filePath); filePath != "" {
					list = append(list, filePath)
				}
			}
		}
	}
	return list, nil
}
