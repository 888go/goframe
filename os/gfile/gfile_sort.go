// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile
import (
	"strings"
	
	"github.com/888go/goframe/container/garray"
	)
// fileSortFunc 是用于文件的比较函数。
// 它按以下顺序对数组进行排序：目录 -> 文件。
// 如果 `path1` 和 `path2` 是相同类型，则按照字符串顺序对它们进行排序。
func fileSortFunc(path1, path2 string) int {
	isDirPath1 := IsDir(path1)
	isDirPath2 := IsDir(path2)
	if isDirPath1 && !isDirPath2 {
		return -1
	}
	if !isDirPath1 && isDirPath2 {
		return 1
	}
	if n := strings.Compare(path1, path2); n != 0 {
		return n
	} else {
		return -1
	}
}

// SortFiles 将 `files` 按照以下顺序进行排序：目录 -> 文件。
// 注意，`files` 中的项应当是绝对路径。
func SortFiles(files []string) []string {
	array := garray.NewSortedStrArrayComparator(fileSortFunc)
	array.Add(files...)
	return array.Slice()
}
