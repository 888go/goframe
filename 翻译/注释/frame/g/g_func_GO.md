
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
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Go creates a new asynchronous goroutine function with specified recover function.
//
// The parameter `recoverFunc` is called when any panic during executing of `goroutineFunc`.
// If `recoverFunc` is given nil, it ignores the panic from `goroutineFunc` and no panic will
// throw to parent goroutine.
//
// But, note that, if `recoverFunc` also throws panic, such panic will be thrown to parent goroutine.
<原文结束>

# <翻译开始>
// Go 创建了一个具有指定恢复函数的新的异步 goroutine 函数。
//
// 参数 `recoverFunc` 在执行 `goroutineFunc` 期间发生任何 panic 时被调用。
// 如果 `recoverFunc` 被赋予 nil，它将忽略来自 `goroutineFunc` 的 panic，并且不会向父级 goroutine 抛出 panic。
//
// 但是请注意，如果 `recoverFunc` 也抛出了 panic，这样的 panic 将会被抛给父级 goroutine。
# <翻译结束>







<原文开始>
// Wait is an alias of ghttp.Wait, which blocks until all the web servers shutdown.
// It's commonly used in multiple servers' situation.
<原文结束>

# <翻译开始>
// Wait 是 ghttp.Wait 的别名，它会阻塞直到所有 web 服务器关闭。
// 在多服务器场景中，它经常被使用。
# <翻译结束>


<原文开始>
// Listen is an alias of gproc.Listen, which handles the signals received and automatically
// calls registered signal handler functions.
// It blocks until shutdown signals received and all registered shutdown handlers done.
<原文结束>

# <翻译开始>
// Listen 是 gproc.Listen 的别名，用于处理接收到的信号并自动调用已注册的信号处理器函数。
// 它会阻塞直到接收到关闭信号且所有已注册的关闭处理器执行完毕。
# <翻译结束>


<原文开始>
// Dump dumps a variable to stdout with more manually readable.
<原文结束>

# <翻译开始>
// Dump 将一个变量以更易于人工阅读的方式输出到标准输出（stdout）
# <翻译结束>


<原文开始>
// DumpTo writes variables `values` as a string in to `writer` with more manually readable
<原文结束>

# <翻译开始>
// DumpTo 将变量 `values` 转换为字符串并写入到 `writer` 中，以更易于人工阅读的方式
# <翻译结束>


<原文开始>
// DumpWithType acts like Dump, but with type information.
// Also see Dump.
<原文结束>

# <翻译开始>
// DumpWithType 的行为类似于 Dump，但会包含类型信息。
// 也可参考 Dump。
# <翻译结束>


<原文开始>
// DumpWithOption returns variables `values` as a string with more manually readable.
<原文结束>

# <翻译开始>
// DumpWithOption 使用自定义选项返回变量 `values`，将其格式化为更易读的字符串形式。
# <翻译结束>


<原文开始>
// DumpJson pretty dumps json content to stdout.
<原文结束>

# <翻译开始>
// DumpJson 将 JSON 内容格式化输出到标准输出（stdout）。
# <翻译结束>


<原文开始>
// Throw throws an exception, which can be caught by TryCatch function.
<原文结束>

# <翻译开始>
// Throw 抛出一个异常，该异常可以被 TryCatch 函数捕获。
# <翻译结束>


<原文开始>
// Try implements try... logistics using internal panic...recover.
// It returns error if any exception occurs, or else it returns nil.
<原文结束>

# <翻译开始>
// Try 使用内部 panic...recover 实现 try...逻辑。
// 如果发生任何异常，它将返回错误，否则返回 nil。
# <翻译结束>


<原文开始>
// TryCatch implements try...catch... logistics using internal panic...recover.
// It automatically calls function `catch` if any exception occurs and passes the exception as an error.
//
// But, note that, if function `catch` also throws panic, the current goroutine will panic.
<原文结束>

# <翻译开始>
// TryCatch 通过内部 panic...recover 实现了类似 try...catch... 的错误处理逻辑。
// 当出现任何异常时，它会自动调用函数 `catch` 并将异常作为 error 参数传递。
//
// 但是请注意，如果函数 `catch` 本身也抛出了 panic，则当前的 goroutine 将同样触发 panic。
# <翻译结束>


<原文开始>
// IsNil checks whether given `value` is nil.
// Parameter `traceSource` is used for tracing to the source variable if given `value` is type
// of pointer that also points to a pointer. It returns nil if the source is nil when `traceSource`
// is true.
// Note that it might use reflect feature which affects performance a little.
<原文结束>

# <翻译开始>
// IsNil 检查给定的 `value` 是否为 nil。
// 参数 `traceSource` 用于在 `value` 类型为指向指针的指针时，追踪到源变量。如果当 `traceSource` 为真时源变量为 nil，则返回 nil。
// 注意，此函数可能使用 reflect 特性，这会对性能产生轻微影响。
# <翻译结束>


<原文开始>
// IsEmpty checks whether given `value` empty.
// It returns true if `value` is in: 0, nil, false, "", len(slice/map/chan) == 0.
// Or else it returns true.
//
// The parameter `traceSource` is used for tracing to the source variable if given `value` is type of pointer
// that also points to a pointer. It returns true if the source is empty when `traceSource` is true.
// Note that it might use reflect feature which affects performance a little.
<原文结束>

# <翻译开始>
// IsEmpty 检查给定的 `value` 是否为空。
// 当 `value` 为以下情形时返回 true：0、nil、false、""、slice/map/chan 的长度为 0。
// 否则返回 false。
// 参数 `traceSource` 用于在 `value` 类型为指针且指向另一个指针时，追踪到源变量。
// 当 `traceSource` 为 true 时，如果源变量为空，则返回 true。
// 注意，这可能使用 reflect 特性，会对性能造成轻微影响。
# <翻译结束>


<原文开始>
// RequestFromCtx retrieves and returns the Request object from context.
<原文结束>

# <翻译开始>
// RequestFromCtx 从 context 中检索并返回 Request 对象。
# <翻译结束>


<原文开始>
// NewVar returns a gvar.Var.
<原文结束>

# <翻译开始>
// NewVar 返回一个 gvar.Var。
# <翻译结束>

