// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdebug

import (
	"bytes"
	"fmt"
	"runtime"
)

// PrintStack 将运行时堆栈跟踪信息输出到标准错误。 md5:91aa0db11ba5ad2a
// ff:
// skip:
func PrintStack(skip ...int) {
	fmt.Print(Stack(skip...))
}

// Stack 返回调用它的goroutine的格式化堆栈跟踪。它调用runtime.Stack，使用足够大的缓冲区来捕获整个跟踪。
// md5:1115a268fe979d5c
// ff:
// skip:
func Stack(skip ...int) string {
	return StackWithFilter(nil, skip...)
}

// StackWithFilter 返回调用它的goroutine的格式化堆栈跟踪。
// 它使用足够大的缓冲区调用runtime.Stack以捕获整个跟踪。
// 
// 参数`filter`用于过滤调用者路径。
// md5:5342cfca7c1801ab
// ff:
// filters:
// skip:
func StackWithFilter(filters []string, skip ...int) string {
	return StackWithFilters(filters, skip...)
}

// StackWithFilters 返回调用它的goroutine的格式化堆栈跟踪。
// 它使用足够大的缓冲区调用runtime.Stack以捕获整个跟踪。
//
// 参数 `filters` 是一个字符串切片，用于过滤调用者的路径。
//
// 待办：使用debug.Stack来提高性能。
// md5:febf8524b3fa5e97
// ff:
// filters:
// skip:
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
