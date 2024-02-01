
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
// SigHandler defines a function type for signal handling.
<原文结束>

# <翻译开始>
// SigHandler 定义了一个用于信号处理的函数类型。
# <翻译结束>


<原文开始>
	// Use internal variable to guarantee concurrent safety
	// when multiple Listen happen.
<原文结束>

# <翻译开始>
// 使用内部变量来保证并发安全性
// 当发生多个Listen操作时。
# <翻译结束>


<原文开始>
// AddSigHandler adds custom signal handler for custom one or more signals.
<原文结束>

# <翻译开始>
// AddSigHandler 为一个或多个自定义信号添加自定义处理函数。
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
// AddSigHandlerShutdown 添加自定义信号处理器以处理关闭信号：
// syscall.SIGINT（中断信号）
// syscall.SIGQUIT（退出并生成 core 文件信号）
// syscall.SIGKILL（强制终止信号，无法被捕获或忽略）
// syscall.SIGTERM（软件终止信号）
// syscall.SIGABRT（异常终止信号，如调用 abort 函数时触发）
# <翻译结束>


<原文开始>
// Listen blocks and does signal listening and handling.
<原文结束>

# <翻译开始>
// Listen阻塞并执行信号监听和处理。
# <翻译结束>


<原文开始>
// If it is shutdown signal, it exits this signal listening.
<原文结束>

# <翻译开始>
// 如果接收到的是关闭信号，则退出该信号监听。
# <翻译结束>


<原文开始>
// Wait until signal handlers done.
<原文结束>

# <翻译开始>
// 等待，直到信号处理器完成。
# <翻译结束>

