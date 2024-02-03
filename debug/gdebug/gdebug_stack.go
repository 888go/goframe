// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdebug

import (
	"bytes"
	"fmt"
	"runtime"
)

// PrintStack 将由 runtime.Stack 返回的堆栈跟踪信息打印到标准错误输出。
func PrintStack(skip ...int) {
	fmt.Print(Stack(skip...))
}

// Stack 函数返回调用该函数的 goroutine 格式化的堆栈跟踪信息。
// 它通过调用 runtime.Stack 并传入足够大的缓冲区来捕获整个跟踪信息。
func Stack(skip ...int) string {
	return StackWithFilter(nil, skip...)
}

// StackWithFilter 返回调用该函数的goroutine格式化后的堆栈跟踪。
// 它通过传入足够大的缓冲区来调用runtime.Stack，以便捕获整个跟踪信息。
//
// 参数`filter`用于过滤调用者路径。
func StackWithFilter(filters []string, skip ...int) string {
	return StackWithFilters(filters, skip...)
}

// StackWithFilters 返回调用该函数的goroutine格式化后的堆栈跟踪。
// 它通过传递足够大的缓冲区调用runtime.Stack，以捕获整个跟踪信息。
//
// 参数`filters`是一个字符串切片，用于过滤调用者路径。
//
// TODO 使用debug.Stack改进性能。
func StackWithFilters(filters []string, skip ...int) string {
	number := 0
	if len(skip) > 0 {
		number = skip[0]
	}
	var (
		name                  string
		space                 = "  "
		index                 = 1
		buffer                = bytes.NewBuffer(nil)
		ok                    = true
		pc, file, line, start = callerFromIndex(filters)
	)
	for i := start + number; i < maxCallerDepth; i++ {
		if i != start {
			pc, file, line, ok = runtime.Caller(i)
		}
		if ok {
			if filterFileByFilters(file, filters) {
				continue
			}
			if fn := runtime.FuncForPC(pc); fn == nil {
				name = "unknown"
			} else {
				name = fn.Name()
			}
			if index > 9 {
				space = " "
			}
			buffer.WriteString(fmt.Sprintf("%d.%s%s\n    %s:%d\n", index, space, name, file, line))
			index++
		} else {
			break
		}
	}
	return buffer.String()
}
