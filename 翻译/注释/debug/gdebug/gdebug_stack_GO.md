
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// PrintStack prints to standard error the stack trace returned by runtime.Stack.
<原文结束>

# <翻译开始>
// PrintStack 将运行时堆栈跟踪信息输出到标准错误。 md5:91aa0db11ba5ad2a
# <翻译结束>


<原文开始>
// Stack returns a formatted stack trace of the goroutine that calls it.
// It calls runtime.Stack with a large enough buffer to capture the entire trace.
<原文结束>

# <翻译开始>
// Stack 返回调用它的goroutine的格式化堆栈跟踪。它调用runtime.Stack，使用足够大的缓冲区来捕获整个跟踪。 md5:1115a268fe979d5c
# <翻译结束>


<原文开始>
// StackWithFilter returns a formatted stack trace of the goroutine that calls it.
// It calls runtime.Stack with a large enough buffer to capture the entire trace.
//
// The parameter `filter` is used to filter the path of the caller.
<原文结束>

# <翻译开始>
// StackWithFilter 返回调用它的goroutine的格式化堆栈跟踪。
// 它使用足够大的缓冲区调用runtime.Stack以捕获整个跟踪。
//
// 参数`filter`用于过滤调用者路径。 md5:5342cfca7c1801ab
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
// StackWithFilters 返回调用它的goroutine的格式化堆栈跟踪。
// 它使用足够大的缓冲区调用runtime.Stack以捕获整个跟踪。
//
// 参数 `filters` 是一个字符串切片，用于过滤调用者的路径。
//
// 待办：使用debug.Stack来提高性能。 md5:febf8524b3fa5e97
# <翻译结束>

