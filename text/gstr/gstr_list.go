// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstr

// List2：使用`delimiter`将`str`分割，并将结果作为两个部分的字符串返回。 md5:7d263c861d943343
func List2(str, delimiter string) (part1, part2 string) {
	return doList2(delimiter, Split(str, delimiter))
}

// ListAndTrim2 使用 `delimiter` 分割并修剪 `str`，然后将结果作为两个部分的字符串返回。 md5:8cd76102c10490b7
func ListAndTrim2(str, delimiter string) (part1, part2 string) {
	return doList2(delimiter, SplitAndTrim(str, delimiter))
}

func doList2(delimiter string, array []string) (part1, part2 string) {
	switch len(array) {
	case 0:
		return "", ""
	case 1:
		return array[0], ""
	case 2:
		return array[0], array[1]
	default:
		return array[0], Join(array[1:], delimiter)
	}
}

// List3 使用 `delimiter` 将 `str` 分割，并将结果作为三个部分的字符串返回。 md5:dc7213002b271e82
func List3(str, delimiter string) (part1, part2, part3 string) {
	return doList3(delimiter, Split(str, delimiter))
}

// ListAndTrim3：使用`delimiter`分割`str`，并将结果作为三个部分的字符串返回。 md5:f669baf0550fad04
func ListAndTrim3(str, delimiter string) (part1, part2, part3 string) {
	return doList3(delimiter, SplitAndTrim(str, delimiter))
}

func doList3(delimiter string, array []string) (part1, part2, part3 string) {
	switch len(array) {
	case 0:
		return "", "", ""
	case 1:
		return array[0], "", ""
	case 2:
		return array[0], array[1], ""
	case 3:
		return array[0], array[1], array[2]
	default:
		return array[0], array[1], Join(array[2:], delimiter)
	}
}
