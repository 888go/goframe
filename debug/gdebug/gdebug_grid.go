// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdebug

import (
	"regexp"
	"runtime"
	"strconv"
)

var (
	// gridRegex 是用于从堆栈信息中解析 goroutine ID 的正则表达式对象。
	gridRegex = regexp.MustCompile(`^\w+\s+(\d+)\s+`)
)

// GoroutineId 从堆栈信息中获取并返回当前 goroutine 的 ID。
// 需要注意的是，由于它使用了 runtime.Stack 函数，所以性能较低。
// 该函数通常用于调试目的。
func GoroutineId() int {
	buf := make([]byte, 26)
	runtime.Stack(buf, false)
	match := gridRegex.FindSubmatch(buf)
	id, _ := strconv.Atoi(string(match[1]))
	return id
}
