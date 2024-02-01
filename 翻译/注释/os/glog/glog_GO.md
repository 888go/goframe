
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
// Package glog implements powerful and easy-to-use leveled logging functionality.
<原文结束>

# <翻译开始>
// 包 glog 实现了强大且易于使用的分级日志功能。
# <翻译结束>


<原文开始>
// ILogger is the API interface for logger.
<原文结束>

# <翻译开始>
// ILogger 是 logger 的 API 接口。
# <翻译结束>


<原文开始>
// Ensure Logger implements ILogger.
<原文结束>

# <翻译开始>
// 确保 Logger 实现了 ILogger 接口。
# <翻译结束>


<原文开始>
// Default logger object, for package method usage.
<原文结束>

# <翻译开始>
// 默认日志器对象，用于包内方法的使用。
# <翻译结束>


<原文开始>
	// Goroutine pool for async logging output.
	// It uses only one asynchronous worker to ensure log sequence.
<原文结束>

# <翻译开始>
// Goroutine 池用于异步日志输出。
// 它仅使用一个异步工作者以确保日志按序输出。
# <翻译结束>


<原文开始>
	// defaultDebug enables debug level or not in default,
	// which can be configured using command option or system environment.
<原文结束>

# <翻译开始>
// defaultDebug 默认是否开启调试级别，可以通过命令行选项或系统环境进行配置。
# <翻译结束>


<原文开始>
// DefaultLogger returns the default logger.
<原文结束>

# <翻译开始>
// DefaultLogger 返回默认日志器。
# <翻译结束>


<原文开始>
// SetDefaultLogger sets the default logger for package glog.
// Note that there might be concurrent safety issue if calls this function
// in different goroutines.
<原文结束>

# <翻译开始>
// SetDefaultLogger 为 glog 包设置默认日志器。
// 注意，如果在不同 goroutine 中调用此函数可能存在并发安全问题。
# <翻译结束>

