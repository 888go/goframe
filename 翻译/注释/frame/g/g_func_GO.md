
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
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
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
// Go 创建一个新的异步 goroutine 函数，并指定了恢复函数。
//
// 参数 `recoverFunc` 在 `goroutineFunc` 执行过程中发生任何 panic 时被调用。
// 如果 `recoverFunc` 为 nil，它将忽略 `goroutineFunc` 中的 panic，且不会向父 goroutine 抛出 panic。
//
// 但是要注意，如果 `recoverFunc` 自身也抛出 panic，这个 panic 将会被抛给父 goroutine。
// md5:3820395064a9e843
# <翻译结束>


<原文开始>
// NewVar returns a gvar.Var.
<原文结束>

# <翻译开始>
// NewVar返回一个gvar.Var。. md5:384cbf5dc86048e6
# <翻译结束>


<原文开始>
// Wait is an alias of ghttp.Wait, which blocks until all the web servers shutdown.
// It's commonly used in multiple servers' situation.
<原文结束>

# <翻译开始>
// Wait 是对 ghttp.Wait 的别名，它会阻塞直到所有的网络服务器停止运行。在多服务器情况下，这通常会被使用。
// md5:efbd28e068404766
# <翻译结束>


<原文开始>
// Listen is an alias of gproc.Listen, which handles the signals received and automatically
// calls registered signal handler functions.
// It blocks until shutdown signals received and all registered shutdown handlers done.
<原文结束>

# <翻译开始>
// Listen 是 gproc.Listen 的别名，用于处理接收到的信号，并自动调用已注册的信号处理函数。
// 它会阻塞直到接收到关闭信号，并且所有已注册的关闭处理函数执行完毕。
// md5:566b27b7da3e69b8
# <翻译结束>


<原文开始>
// Dump dumps a variable to stdout with more manually readable.
<原文结束>

# <翻译开始>
// Dump 将变量以更易读的格式输出到标准输出。. md5:e19cf0b857ffd1c6
# <翻译结束>


<原文开始>
// DumpTo writes variables `values` as a string in to `writer` with more manually readable
<原文结束>

# <翻译开始>
// DumpTo 将变量 `values` 作为字符串写入到 `writer` 中，提供更易人工阅读的格式. md5:68fd8fc9ea0dfc4b
# <翻译结束>


<原文开始>
// DumpWithType acts like Dump, but with type information.
// Also see Dump.
<原文结束>

# <翻译开始>
// DumpWithType 类似于 Dump，但带有类型信息。同时参阅 Dump。
// md5:faabab79589d38a3
# <翻译结束>


<原文开始>
// DumpWithOption returns variables `values` as a string with more manually readable.
<原文结束>

# <翻译开始>
// DumpWithOption 函数将变量 `values` 以更易于人工阅读的字符串形式返回。. md5:99fec3f0f209dcf7
# <翻译结束>


<原文开始>
// DumpJson pretty dumps json content to stdout.
<原文结束>

# <翻译开始>
// DumpJson 将 JSON 内容以美化的方式输出到标准输出。. md5:9f4c95e099395360
# <翻译结束>


<原文开始>
// Throw throws an exception, which can be caught by TryCatch function.
<原文结束>

# <翻译开始>
// Throw抛出一个异常，可以被TryCatch函数捕获。. md5:ade0a15f9238635b
# <翻译结束>


<原文开始>
// Try implements try... logistics using internal panic...recover.
// It returns error if any exception occurs, or else it returns nil.
<原文结束>

# <翻译开始>
// Try 使用内部的 panic...recover 实现 try...catch 逻辑。如果发生任何异常，它会返回错误；否则返回 nil。
// md5:7c85aa857bb16fca
# <翻译结束>


<原文开始>
// TryCatch implements try...catch... logistics using internal panic...recover.
// It automatically calls function `catch` if any exception occurs and passes the exception as an error.
//
// But, note that, if function `catch` also throws panic, the current goroutine will panic.
<原文结束>

# <翻译开始>
// TryCatch 使用内部的 panic...recover 机制来实现 try...catch... 的逻辑。
// 如果发生任何异常，它会自动调用函数 `catch` 并将异常作为错误传递。
//
// 但请注意，如果函数 `catch` 本身也引发了 panic，当前的 goroutine 会触发 panic。
// md5:c9fae3297a82421f
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
// 参数 `traceSource` 用于在 `value` 为指向指针的指针类型时，追踪到源变量。当 `traceSource` 为真且源为 nil 时，它会返回 nil。
// 请注意，该功能可能使用反射，这可能稍微影响性能。
// md5:1a86ab3bedd8914d
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
// 如果 `value` 等于：0，nil，false，""，切片、映射或通道的长度为0，则返回 true。
// 否则，返回 true。
// 
// 参数 `traceSource` 用于追踪，如果给定的 `value` 是指向指针的指针类型，它会追踪到源变量。当 `traceSource` 为 true 时，如果源为空，将返回 true。
// 注意，它可能使用反射特性，这可能会影响性能。
// md5:7262defa92ce37cb
# <翻译结束>


<原文开始>
// RequestFromCtx retrieves and returns the Request object from context.
<原文结束>

# <翻译开始>
// RequestFromCtx 从上下文中检索并返回Request对象。. md5:c247eac3d031fb2b
# <翻译结束>

