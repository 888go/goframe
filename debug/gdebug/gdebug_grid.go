// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdebug

import (
	"regexp"
	"runtime"
	"strconv"
)

var (
	// gridRegex 是用于从堆栈信息中解析goroutine ID的正则表达式对象。. md5:682403f05aacc855
	gridRegex = regexp.MustCompile(`^\w+\s+(\d+)\s+`)
)

// GoroutineId 从堆栈信息中获取并返回当前goroutine的ID。
// 需要特别注意的是，由于它使用了runtime.Stack函数，因此性能较低。
// 通常用于调试目的。
// md5:c6453659dbcae88d
func GoroutineId() int {
	buf := make([]byte, 26)
	runtime.Stack(buf, false)
	match := gridRegex.FindSubmatch(buf)
	id, _ := strconv.Atoi(string(match[1]))
	return id
}
