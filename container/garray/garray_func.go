// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package garray

import "strings"

// defaultComparatorInt 用于整数比较。. md5:c6fa54bf179ad1f1
func defaultComparatorInt(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// defaultComparatorStr 用于字符串比较。. md5:031e6c286f826685
func defaultComparatorStr(a, b string) int {
	return strings.Compare(a, b)
}

// quickSortInt 是为整型数据实现的快速排序算法。. md5:f172b33dc88c6a22
func quickSortInt(values []int, comparator func(a, b int) int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		if comparator(values[i], mid) > 0 {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	quickSortInt(values[:head], comparator)
	quickSortInt(values[head+1:], comparator)
}

// quickSortStr 是用于字符串的快速排序算法实现。. md5:7df719bfdb420cc1
func quickSortStr(values []string, comparator func(a, b string) int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		if comparator(values[i], mid) > 0 {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	quickSortStr(values[:head], comparator)
	quickSortStr(values[head+1:], comparator)
}
