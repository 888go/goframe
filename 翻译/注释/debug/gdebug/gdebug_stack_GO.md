
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// PrintStack prints to standard error the stack trace returned by runtime.Stack.
<原文结束>

# <翻译开始>
// PrintStack 将由 runtime.Stack 返回的堆栈跟踪信息打印到标准错误输出。
# <翻译结束>


<原文开始>
// Stack returns a formatted stack trace of the goroutine that calls it.
// It calls runtime.Stack with a large enough buffer to capture the entire trace.
<原文结束>

# <翻译开始>
// Stack 函数返回调用该函数的 goroutine 格式化的堆栈跟踪信息。
// 它通过调用 runtime.Stack 并传入足够大的缓冲区来捕获整个跟踪信息。
# <翻译结束>


<原文开始>
// StackWithFilter returns a formatted stack trace of the goroutine that calls it.
// It calls runtime.Stack with a large enough buffer to capture the entire trace.
//
// The parameter `filter` is used to filter the path of the caller.
<原文结束>

# <翻译开始>
// StackWithFilter 返回调用该函数的goroutine格式化后的堆栈跟踪。
// 它通过传入足够大的缓冲区来调用runtime.Stack，以便捕获整个跟踪信息。
//
// 参数`filter`用于过滤调用者路径。
# <翻译结束>


<原文开始>
// StackWithFilters returns a formatted stack trace of the goroutine that calls it.
// It calls runtime.Stack with a large enough buffer to capture the entire trace.
//
// The parameter `filters` is a slice of strings, which are used to filter the path of the
// caller.
//
// TODO Improve the performance using debug.Stack.
<原文结束>

# <翻译开始>
// StackWithFilters 返回调用该函数的goroutine格式化后的堆栈跟踪。
// 它通过传递足够大的缓冲区调用runtime.Stack，以捕获整个跟踪信息。
//
// 参数`filters`是一个字符串切片，用于过滤调用者路径。
//
// TODO 使用debug.Stack改进性能。
# <翻译结束>

