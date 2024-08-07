// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类

import (
	"strings"

	garray "github.com/888go/goframe/container/garray"
)

// fileSortFunc 是用于文件的比较函数。
// 它按照顺序对数组进行排序：目录 -> 文件。
// 如果 `path1` 和 `path2` 是同一种类型，它会将它们作为字符串进行进一步排序。
// md5:f85e943b3d688062
func fileSortFunc(path1, path2 string) int {
	isDirPath1 := X是否存在目录(path1)
	isDirPath2 := X是否存在目录(path2)
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

// X排序 按照目录 -> 文件的顺序对 `files` 进行排序。
// 请注意，`files` 列表中的项应该是绝对路径。
// md5:78b3df91d9486a1b
func X排序(文件切片 []string) []string {
	array := garray.X创建文本排序并带排序函数(fileSortFunc)
	array.X入栈右(文件切片...)
	return array.X取切片()
}
