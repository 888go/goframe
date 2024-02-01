// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr

// List2 使用`delimiter`分割`str`并将结果作为两个部分的字符串返回。
func List2(str, delimiter string) (part1, part2 string) {
	return doList2(delimiter, Split(str, delimiter))
}

// ListAndTrim2 使用`delimiter`分割并修剪`str`，然后将结果以两个部分的字符串形式返回。
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

// List3 使用`delimiter`分割`str`并将结果以三个部分的字符串形式返回。
func List3(str, delimiter string) (part1, part2, part3 string) {
	return doList3(delimiter, Split(str, delimiter))
}

// ListAndTrim3 以`delimiter`为分隔符对`str`进行分割并去除首尾空白字符，然后将结果返回为三个部分的字符串。
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
