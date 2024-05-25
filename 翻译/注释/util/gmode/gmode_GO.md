
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
// Package gmode provides release mode management for project.
//
// It uses string to mark the mode instead of integer, which is convenient for configuration.
<原文结束>

# <翻译开始>
// gmode 包为项目提供发布模式管理功能。
//
// 它使用字符串而非整数来标记模式，便于配置。
// md5:716bfe2e364994bd
# <翻译结束>


<原文开始>
// Note that `currentMode` is not concurrent safe.
<原文结束>

# <翻译开始>
// 请注意，`currentMode` 不是并发安全的。 md5:71c58ea0b37583d8
# <翻译结束>


<原文开始>
// Set sets the mode for current application.
<原文结束>

# <翻译开始>
// Set 设置当前应用程序的模式。 md5:523a8be23f0521ca
# <翻译结束>


<原文开始>
// SetDevelop sets current mode DEVELOP for current application.
<原文结束>

# <翻译开始>
// SetDevelop 将当前应用的模式设置为 DEVELOP。 md5:7d10eaf834b69114
# <翻译结束>


<原文开始>
// SetTesting sets current mode TESTING for current application.
<原文结束>

# <翻译开始>
// SetTesting 将当前应用程序的运行模式设置为测试模式。 md5:3e3ee802bec9b04e
# <翻译结束>


<原文开始>
// SetStaging sets current mode STAGING for current application.
<原文结束>

# <翻译开始>
// SetStaging 将当前应用的模式设置为 STAGING。 md5:c8e3fac819c1d0b9
# <翻译结束>


<原文开始>
// SetProduct sets current mode PRODUCT for current application.
<原文结束>

# <翻译开始>
// SetProduct 将当前应用设置为PRODUCT模式。 md5:e681c0b16f9b2bf0
# <翻译结束>


<原文开始>
// Mode returns current application mode set.
<原文结束>

# <翻译开始>
// Mode 返回当前设置的应用模式。 md5:76410fdca2d2e6a9
# <翻译结束>


<原文开始>
// If current mode is not set, do this auto check.
<原文结束>

# <翻译开始>
// 如果当前模式未设置，则执行此自动检查。 md5:a6a2104d461130ba
# <翻译结束>


<原文开始>
// Mode configured from command argument of environment.
<原文结束>

# <翻译开始>
// 从命令行参数或环境设置的模式。 md5:9b483ed8828c68a6
# <翻译结束>


<原文开始>
// If there are source codes found, it's in develop mode, or else in product mode.
<原文结束>

# <翻译开始>
// 如果找到源代码，则为开发模式，否则为产品模式。 md5:4feb614f3843b3df
# <翻译结束>


<原文开始>
// IsDevelop checks and returns whether current application is running in DEVELOP mode.
<原文结束>

# <翻译开始>
// IsDevelop 检查并返回当前应用是否正在运行在DEVELOP模式下。 md5:577f79eacdd2c47e
# <翻译结束>


<原文开始>
// IsTesting checks and returns whether current application is running in TESTING mode.
<原文结束>

# <翻译开始>
// IsTesting 检查并返回当前应用是否正在以测试模式运行。 md5:3411dababba12269
# <翻译结束>


<原文开始>
// IsStaging checks and returns whether current application is running in STAGING mode.
<原文结束>

# <翻译开始>
// IsStaging 检查并返回当前应用程序是否在 STAGING 模式下运行。 md5:99c92c19f9a8925f
# <翻译结束>


<原文开始>
// IsProduct checks and returns whether current application is running in PRODUCT mode.
<原文结束>

# <翻译开始>
// IsProduct 检查并返回当前应用程序是否正在PRODUCT模式下运行。 md5:cf7849a0659f26bb
# <翻译结束>

