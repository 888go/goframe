// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类

import (
	"path/filepath"
	"sort"

	gerror "github.com/888go/goframe/errors/gerror"
	gstr "github.com/888go/goframe/text/gstr"
)

const (
		// 扫描目录时的最大递归深度。 md5:6443b3b221d62366
	maxScanDepth = 100000
)

// X枚举并含子目录名 返回给定`path`下的所有子文件的绝对路径，
// 如果给定的参数`recursive`为true，则递归扫描目录。
//
// 模式参数`pattern`支持多个文件名模式，
// 使用`,`符号分隔多个模式。
// md5:1f662f1008f0113e
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

// X枚举并含子目录名_函数 返回给定`path`下的所有子文件的绝对路径。
// 如果参数`recursive`为真，它将递归扫描目录。
//
// 参数`pattern`支持多个文件名模式，使用逗号分隔多个模式。
//
// 参数`recursive`指定是否递归扫描`path`。默认情况下，它是false，表示不递归。
//
// 参数`handler`指定了处理`path`及其子目录下每个子文件路径的回调函数。如果`handler`返回空字符串，将忽略子文件路径，否则将子文件路径添加到结果切片中。
// md5:93774b4b752cee08
func X枚举并含子目录名_函数(目录 string, 匹配文件模式 string, 是否递归替换 bool, 回调函数 func(目录 string) string) ([]string, error) {
	list, err := doScanDir(0, 目录, 匹配文件模式, 是否递归替换, 回调函数)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		sort.Strings(list)
	}
	return list, nil
}

// X枚举 返回给定 `path` 所有子文件的绝对路径，
// 如果 `recursive` 参数为真，它会递归扫描目录。
//
// `pattern` 参数支持多个文件名模式，使用逗号 `,` 来分隔多个模式。
//
// 注意，它只返回文件，不包括目录。
// md5:1d9c6ada055eaa05
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

// X枚举_函数 返回给定 `path` 的所有子文件的绝对路径，
// 如果参数 `recursive` 为 true，则会递归扫描目录。
//
// `pattern` 参数支持多个文件名模式，使用逗号（',') 分隔多个模式。
//
// 参数 `recursive` 指定是否递归扫描 `path`，即如果子文件也是一个文件夹，它将把子文件路径添加到结果数组中。默认情况下为 false。
//
// 参数 `handler` 指定处理 `path` 和其子文件夹每个子文件路径的回调函数。如果 `handler` 返回空字符串，那么忽略该子文件路径；否则，将子文件路径添加到结果切片中。
//
// 注意，`handler` 中的参数 `path` 不是目录，而是文件。它只返回文件，不包括目录。
// md5:036965ff87c95b63
func X枚举_函数(目录 string, 匹配文件模式 string, 是否递归查找 bool, 匿名函数 func(目录 string) string) ([]string, error) {
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

// doScanDir 是一个内部方法，用于扫描目录并返回未排序的文件绝对路径列表。
//
// 模式参数 `pattern` 支持多个文件名模式，使用 ',' 符号分隔多个模式。
//
// 参数 `recursive` 指定是否递归扫描 `path`，即如果子文件也是一个文件夹，
// 则扫描其子文件并将文件路径追加到结果数组中。默认为 false。
//
// 参数 `handler` 指定一个回调函数，用于处理 `path` 及其子目录下的每个子文件路径。
// 如果 `handler` 返回空字符串，则忽略该子文件路径；否则，将子文件路径追加到结果切片中。
// md5:5f6bc88fb2ff75fe
func doScanDir(depth int, path string, pattern string, recursive bool, handler func(path string) string) ([]string, error) {
	if depth >= maxScanDepth {
		return nil, gerror.X创建并格式化("directory scanning exceeds max recursive depth: %d", maxScanDepth)
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
		err = gerror.X多层错误并格式化(err, `read directory files failed from path "%s"`, path)
		return nil, err
	}
	var (
		filePath string
		patterns = gstr.X分割并忽略空值(pattern, ",")
	)
	for _, name := range names {
		filePath = path + Separator + name
		if X是否存在目录(filePath) && recursive {
			array, _ := doScanDir(depth+1, filePath, pattern, true, handler)
			if len(array) > 0 {
				list = append(list, array...)
			}
		}
		// Handler filtering.
		if handler != nil {
			filePath = handler(filePath)
			if filePath == "" {
				continue
			}
		}
				// 如果满足模式，将其添加到结果列表中。 md5:11ed1569cf70af04
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
