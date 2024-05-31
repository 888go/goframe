// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gstr

// List2 Split the `str` with `delimiter` and returns the result as two parts string.

// ff:分割2份
// part2:返回值2
// part1:返回值1
// delimiter:分隔符
// str:待分割文本
func List2(str, delimiter string) (part1, part2 string) {
	return doList2(delimiter, Split(str, delimiter))
}

// ListAndTrim2 SplitAndTrim the `str` with `delimiter` and returns the result as two parts string.

// ff:分割2份并忽略空值
// part2:返回值2
// part1:返回值1
// delimiter:分隔符
// str:待分割文本
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

// List3 Split the `str` with `delimiter` and returns the result as three parts string.

// ff:分割3份
// part3:返回值3
// part2:返回值2
// part1:返回值1
// delimiter:分隔符
// str:待分割文本
func List3(str, delimiter string) (part1, part2, part3 string) {
	return doList3(delimiter, Split(str, delimiter))
}

// ListAndTrim3 SplitAndTrim the `str` with `delimiter` and returns the result as three parts string.

// ff:分割3份并忽略空值
// part3:返回值3
// part2:返回值2
// part1:返回值1
// delimiter:分隔符
// str:待分割文本
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
