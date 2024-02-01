
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
// pathType is the type for i18n file path.
<原文结束>

# <翻译开始>
// pathType 是用于国际化文件路径的类型。
# <翻译结束>


<原文开始>
// Manager for i18n contents, it is concurrent safe, supporting hot reload.
<原文结束>

# <翻译开始>
// i18n内容的管理器，它是并发安全的，并支持热重载。
# <翻译结束>

















<原文开始>
// Options is used for i18n object configuration.
<原文结束>

# <翻译开始>
// Options 用于i18n对象配置。
# <翻译结束>












<原文开始>
// Delimiters for variable parsing.
<原文结束>

# <翻译开始>
// 变量解析的分隔符。
# <翻译结束>







<原文开始>
// defaultDelimiters defines the default language if user does not specify in options.
<原文结束>

# <翻译开始>
// defaultDelimiters 定义了如果用户在选项中未指定时的默认分隔符语言。
# <翻译结束>


<原文开始>
// defaultDelimiters defines the default key variable delimiters.
<原文结束>

# <翻译开始>
// defaultDelimiters 定义了默认的关键字变量分隔符。
# <翻译结束>


<原文开始>
// i18n files searching folders.
<原文结束>

# <翻译开始>
// i18n文件搜索目录
# <翻译结束>


<原文开始>
// New creates and returns a new i18n manager.
// The optional parameter `option` specifies the custom options for i18n manager.
// It uses a default one if it's not passed.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的 i18n 管理器。
// 可选参数 `option` 指定 i18n 管理器的自定义选项。
// 若未传递该参数，则使用默认选项。
# <翻译结束>


<原文开始>
// To avoid of the source path of GoFrame: github.com/gogf/i18n/gi18n
<原文结束>

# <翻译开始>
// 为避免引用GoFrame的源路径：github.com/gogf/i18n/gi18n
# <翻译结束>


<原文开始>
// checkPathType checks and returns the path type for given directory path.
<原文结束>

# <翻译开始>
// checkPathType 对给定的目录路径进行检查并返回其路径类型。
# <翻译结束>


<原文开始>
// SetPath sets the directory path storing i18n files.
<原文结束>

# <翻译开始>
// SetPath 设置存储 i18n 文件的目录路径。
# <翻译结束>


<原文开始>
// Reset the manager after path changed.
<原文结束>

# <翻译开始>
// 在路径改变后重置管理器。
# <翻译结束>


<原文开始>
// SetLanguage sets the language for translator.
<原文结束>

# <翻译开始>
// SetLanguage 设置翻译器的语言。
# <翻译结束>


<原文开始>
// SetDelimiters sets the delimiters for translator.
<原文结束>

# <翻译开始>
// SetDelimiters 设置翻译器的分隔符。
# <翻译结束>


<原文开始>
// T is alias of Translate for convenience.
<原文结束>

# <翻译开始>
// T 是 Translate 的别名，用于提供便利。
# <翻译结束>


<原文开始>
// Tf is alias of TranslateFormat for convenience.
<原文结束>

# <翻译开始>
// Tf 是 TranslateFormat 的别名，用于提供便利。
# <翻译结束>


<原文开始>
// TranslateFormat translates, formats and returns the `format` with configured language
// and given `values`.
<原文结束>

# <翻译开始>
// TranslateFormat 将根据配置的语言和给定的 `values` 对 `format` 进行翻译、格式化并返回结果。
# <翻译结束>


<原文开始>
// Translate translates `content` with configured language.
<原文结束>

# <翻译开始>
// Translate 使用配置的语言对`content`进行翻译。
# <翻译结束>







<原文开始>
// Parse content as variables container.
<原文结束>

# <翻译开始>
// 将内容解析为变量容器。
# <翻译结束>


<原文开始>
			// return match[1] will return the content between delimiters
			// return match[0] will return the original content
<原文结束>

# <翻译开始>
// 返回match[1]将返回分隔符之间的内容
// 返回match[0]将返回原始内容
# <翻译结束>


<原文开始>
// GetContent retrieves and returns the configured content for given key and specified language.
// It returns an empty string if not found.
<原文结束>

# <翻译开始>
// GetContent 函数根据给定的键和指定的语言获取并返回配置的内容。
// 如果未找到，则返回一个空字符串。
# <翻译结束>


<原文开始>
// reset reset data of the manager.
<原文结束>

# <翻译开始>
// reset 重置管理器的数据。
# <翻译结束>


<原文开始>
// init initializes the manager for lazy initialization design.
// The i18n manager is only initialized once.
<原文结束>

# <翻译开始>
// init 用于实现延迟初始化设计，初始化i18n管理器。
// i18n管理器只初始化一次。
# <翻译结束>


<原文开始>
// If the data is not nil, means it's already initialized.
<原文结束>

# <翻译开始>
// 如果数据不为nil，表示它已经被初始化过。
# <翻译结束>


<原文开始>
// Monitor changes of i18n files for hot reload feature.
<原文结束>

# <翻译开始>
// 监控i18n文件的更改以实现热重载功能。
# <翻译结束>


<原文开始>
// Any changes of i18n files, clear the data.
<原文结束>

# <翻译开始>
// 如果i18n文件有任何更改，清空数据。
# <翻译结束>







<原文开始>
// Pattern for regex parsing.
<原文结束>

# <翻译开始>
// 正则表达式解析的模式。
# <翻译结束>


<原文开始>
// Path type for i18n files.
<原文结束>

# <翻译开始>
// Path 类型用于 i18n 文件。
# <翻译结束>


<原文开始>
// configuration options.
<原文结束>

# <翻译开始>
// 配置选项
# <翻译结束>


<原文开始>
// I18n files storage path.
<原文结束>

# <翻译开始>
// 国际化文件存储路径。
# <翻译结束>


<原文开始>
// Default local language.
<原文结束>

# <翻译开始>
// 默认本地语言。
# <翻译结束>


<原文开始>
// Resource for i18n files.
<原文结束>

# <翻译开始>
// i18n文件资源。
# <翻译结束>


<原文开始>
// Parse content as name.
<原文结束>

# <翻译开始>
// 将内容解析为名称。
# <翻译结束>

