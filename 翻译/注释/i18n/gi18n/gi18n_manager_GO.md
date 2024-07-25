
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
// pathType is the type for i18n file path.
<原文结束>

# <翻译开始>
// pathType 是用于i18n文件路径的类型。 md5:1aa056f2406cd3a6
# <翻译结束>


<原文开始>
// Manager for i18n contents, it is concurrent safe, supporting hot reload.
<原文结束>

# <翻译开始>
// i18n内容的管理器，它是并发安全的，支持热重载。 md5:9c519435bec8f5ad
# <翻译结束>


<原文开始>
// Pattern for regex parsing.
<原文结束>

# <翻译开始>
// 正则表达式解析的模式。 md5:1d0109d4850fd141
# <翻译结束>


<原文开始>
// Path type for i18n files.
<原文结束>

# <翻译开始>
// i18n 文件的路径类型。 md5:5b086fc46e36729e
# <翻译结束>


<原文开始>
// Options is used for i18n object configuration.
<原文结束>

# <翻译开始>
// Options 用于国际化对象的配置。 md5:029f81136c5c3e6a
# <翻译结束>


<原文开始>
// I18n files storage path.
<原文结束>

# <翻译开始>
// 国际化文件的存储路径。 md5:67cec25950dc6464
# <翻译结束>


<原文开始>
// Default local language.
<原文结束>

# <翻译开始>
// 默认本地语言。 md5:41ccf6a6028cf49d
# <翻译结束>


<原文开始>
// Delimiters for variable parsing.
<原文结束>

# <翻译开始>
// 变量解析的定界符。 md5:355db5afb17acaf5
# <翻译结束>


<原文开始>
// Resource for i18n files.
<原文结束>

# <翻译开始>
// i18n文件的资源。 md5:611cd5c408223400
# <翻译结束>


<原文开始>
// defaultLanguage defines the default language if user does not specify in options.
<原文结束>

# <翻译开始>
	// defaultLanguage 定义了如果用户在选项中未指定，默认的语言。 md5:37b426a695c48d49
# <翻译结束>


<原文开始>
// defaultDelimiters defines the default key variable delimiters.
<原文结束>

# <翻译开始>
	// defaultDelimiters 定义了默认的键变量分隔符。 md5:98706258206bfd9a
# <翻译结束>


<原文开始>
// i18n files searching folders.
<原文结束>

# <翻译开始>
	// 国际化文件搜索目录。 md5:cf8914abf6ec0557
# <翻译结束>


<原文开始>
// New creates and returns a new i18n manager.
// The optional parameter `option` specifies the custom options for i18n manager.
// It uses a default one if it's not passed.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的国际化管理器。
// 可选参数 `option` 用于指定国际化管理器的自定义选项。
// 如果未传递该参数，它将使用默认选项。 md5:79f31dcd2ff8cf56
# <翻译结束>


<原文开始>
// To avoid of the source path of GoFrame: github.com/gogf/i18n/gi18n
<原文结束>

# <翻译开始>
			// 为了避免GoFrame的源路径：github.com/gogf/i18n/gi18n. md5:2eecc4478ca65bd7
# <翻译结束>


<原文开始>
// checkPathType checks and returns the path type for given directory path.
<原文结束>

# <翻译开始>
// checkPathType 检查并返回给定目录路径的路径类型。 md5:101af7b8de6f50f8
# <翻译结束>


<原文开始>
// SetPath sets the directory path storing i18n files.
<原文结束>

# <翻译开始>
// SetPath 设置存储i18n文件的目录路径。 md5:b39e1d244949dcf8
# <翻译结束>


<原文开始>
// Reset the manager after path changed.
<原文结束>

# <翻译开始>
	// 路径改变后重置管理器。 md5:1f0260d8d112184d
# <翻译结束>


<原文开始>
// SetLanguage sets the language for translator.
<原文结束>

# <翻译开始>
// SetLanguage 设置翻译器的语言。 md5:50b09b0bb0944dc1
# <翻译结束>


<原文开始>
// SetDelimiters sets the delimiters for translator.
<原文结束>

# <翻译开始>
// SetDelimiters 为翻译器设置分隔符。 md5:f84b046b11204dc7
# <翻译结束>


<原文开始>
// T is alias of Translate for convenience.
<原文结束>

# <翻译开始>
// T 是为了方便而对 Translate 的别名。 md5:c07a6fa99a429eb3
# <翻译结束>


<原文开始>
// Tf is alias of TranslateFormat for convenience.
<原文结束>

# <翻译开始>
// Tf是TranslateFormat的别名，为了方便起见。 md5:bdb209b24c669f5a
# <翻译结束>


<原文开始>
// TranslateFormat translates, formats and returns the `format` with configured language
// and given `values`.
<原文结束>

# <翻译开始>
// TranslateFormat 使用配置的语言和给定的 `values` 对 `format` 进行翻译、格式化并返回结果。 md5:2806a81d6db86c7f
# <翻译结束>


<原文开始>
// Translate translates `content` with configured language.
<原文结束>

# <翻译开始>
// Translate 使用配置的语言翻译`content`。 md5:8f8b7d32e0b26a99
# <翻译结束>


<原文开始>
// Parse content as variables container.
<原文结束>

# <翻译开始>
	// 解析内容作为变量容器。 md5:6fab6ca886fe327a
# <翻译结束>


<原文开始>
			// return match[1] will return the content between delimiters
			// return match[0] will return the original content
<原文结束>

# <翻译开始>
			// 返回match[1] 将返回分隔符之间的内容
			// 返回match[0] 将返回原始内容 md5:3dd48230b02f1348
# <翻译结束>


<原文开始>
// GetContent retrieves and returns the configured content for given key and specified language.
// It returns an empty string if not found.
<原文结束>

# <翻译开始>
// GetContent 获取并返回给定键和指定语言的配置内容。
// 如果未找到，将返回一个空字符串。 md5:c64a3a803ac07e38
# <翻译结束>


<原文开始>
// reset reset data of the manager.
<原文结束>

# <翻译开始>
// reset 重置管理器的数据。 md5:582ac65a0b066583
# <翻译结束>


<原文开始>
// init initializes the manager for lazy initialization design.
// The i18n manager is only initialized once.
<原文结束>

# <翻译开始>
// init 初始化管理器，用于延迟初始化设计。
// 国际化(i18n)管理器仅初始化一次。 md5:b3e5cf7f018d1485
# <翻译结束>


<原文开始>
// If the data is not nil, means it's already initialized.
<原文结束>

# <翻译开始>
	// 如果数据不为nil，表示它已经初始化。 md5:8d4d8b324fc9951a
# <翻译结束>


<原文开始>
// Monitor changes of i18n files for hot reload feature.
<原文结束>

# <翻译开始>
		// 监控i18n文件的变化，以实现热重载功能。 md5:feeb4e0abb048a7b
# <翻译结束>


<原文开始>
// Any changes of i18n files, clear the data.
<原文结束>

# <翻译开始>
			// 对i18n文件的任何更改，都会清空数据。 md5:fbcc6de55a881c92
# <翻译结束>

