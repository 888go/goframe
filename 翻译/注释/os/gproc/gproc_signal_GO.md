
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
// SigHandler defines a function type for signal handling.
<原文结束>

# <翻译开始>
// SigHandler 定义了一个用于处理信号的函数类型。 md5:d7319108f37510cd
# <翻译结束>


<原文开始>
	// Use internal variable to guarantee concurrent safety
	// when multiple Listen happen.
<原文结束>

# <翻译开始>
	// 使用内部变量来保证当发生多个监听（Listen）时的并发安全。 md5:7c2a9e90bd5be8da
# <翻译结束>


<原文开始>
// AddSigHandler adds custom signal handler for custom one or more signals.
<原文结束>

# <翻译开始>
// AddSigHandler 为自定义的一个或多个信号添加自定义的信号处理器。 md5:996226c8d75ebdf5
# <翻译结束>


<原文开始>
// AddSigHandlerShutdown adds custom signal handler for shutdown signals:
// syscall.SIGINT,
// syscall.SIGQUIT,
// syscall.SIGKILL,
// syscall.SIGTERM,
// syscall.SIGABRT.
<原文结束>

# <翻译开始>
// AddSigHandlerShutdown 为关闭信号添加自定义信号处理器：
// - syscall.SIGINT（中断信号，通常由Ctrl+C触发）
// - syscall.SIGQUIT（退出信号，通常通过Ctrl+\触发）
// - syscall.SIGKILL（杀死信号，不可被捕获或忽略，用于强制终止进程）
// - syscall.SIGTERM（终止信号，用来请求程序正常退出）
// - syscall.SIGABRT（异常终止信号，通常由调用abort函数产生，用于指示严重错误） md5:6fd417c58f499e80
# <翻译结束>


<原文开始>
// Listen blocks and does signal listening and handling.
<原文结束>

# <翻译开始>
// Listen 会阻塞并进行信号监听和处理。 md5:2425bc5d9026c36f
# <翻译结束>


<原文开始>
// If it is shutdown signal, it exits this signal listening.
<原文结束>

# <翻译开始>
		// 如果它是退出信号，它将退出此信号监听。 md5:9b1cb86f40c5e361
# <翻译结束>


<原文开始>
// Wait until signal handlers done.
<原文结束>

# <翻译开始>
			// 等待信号处理器完成。 md5:4d1ee13d17a0a193
# <翻译结束>

