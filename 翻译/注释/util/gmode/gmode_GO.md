
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
// Package gmode provides release mode management for project.
//
// It uses string to mark the mode instead of integer, which is convenient for configuration.
<原文结束>

# <翻译开始>
// Package gmode 为项目提供发布模式管理功能。
//
// 它使用字符串而非整数来标记模式，这种方式对于配置更为方便。
# <翻译结束>


<原文开始>
// Note that `currentMode` is not concurrent safe.
<原文结束>

# <翻译开始>
// 注意，`currentMode` 不是线程安全的。
# <翻译结束>


<原文开始>
// Set sets the mode for current application.
<原文结束>

# <翻译开始>
// Set 设置当前应用程序的模式。
# <翻译结束>


<原文开始>
// SetDevelop sets current mode DEVELOP for current application.
<原文结束>

# <翻译开始>
// SetDevelop 将当前应用程序的当前模式设置为 DEVELOP 模式。
# <翻译结束>


<原文开始>
// SetTesting sets current mode TESTING for current application.
<原文结束>

# <翻译开始>
// SetTesting 设置当前应用程序的当前模式为 TESTING。
# <翻译结束>


<原文开始>
// SetStaging sets current mode STAGING for current application.
<原文结束>

# <翻译开始>
// SetStaging 将当前应用的当前模式设置为 STAGING。
# <翻译结束>


<原文开始>
// SetProduct sets current mode PRODUCT for current application.
<原文结束>

# <翻译开始>
// SetProduct 将当前应用的模式设置为PRODUCT模式。
# <翻译结束>


<原文开始>
// Mode returns current application mode set.
<原文结束>

# <翻译开始>
// Mode 返回当前设置的应用程序模式。
# <翻译结束>


<原文开始>
// If current mode is not set, do this auto check.
<原文结束>

# <翻译开始>
// 如果当前模式未设置，则执行此自动检查。
# <翻译结束>


<原文开始>
// Mode configured from command argument of environment.
<原文结束>

# <翻译开始>
// Mode 由命令行参数或环境变量配置。
# <翻译结束>


<原文开始>
// If there are source codes found, it's in develop mode, or else in product mode.
<原文结束>

# <翻译开始>
// 如果找到了源代码，那么处于开发模式；否则，处于产品模式。
# <翻译结束>


<原文开始>
// IsDevelop checks and returns whether current application is running in DEVELOP mode.
<原文结束>

# <翻译开始>
// IsDevelop 检查并返回当前应用程序是否在开发模式下运行。
# <翻译结束>


<原文开始>
// IsTesting checks and returns whether current application is running in TESTING mode.
<原文结束>

# <翻译开始>
// IsTesting 检查并返回当前应用程序是否处于测试模式运行。
# <翻译结束>


<原文开始>
// IsStaging checks and returns whether current application is running in STAGING mode.
<原文结束>

# <翻译开始>
// IsStaging 检查并返回当前应用程序是否在 STAGING（暂存）模式下运行。
# <翻译结束>


<原文开始>
// IsProduct checks and returns whether current application is running in PRODUCT mode.
<原文结束>

# <翻译开始>
// IsProduct 检查并返回当前应用程序是否在PRODUCT模式下运行。
# <翻译结束>

