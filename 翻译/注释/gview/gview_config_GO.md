
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
// Config is the configuration object for template engine.
<原文结束>

# <翻译开始>
// Config 是用于模板引擎的配置对象。
# <翻译结束>


<原文开始>
// Searching array for path, NOT concurrent-safe for performance purpose.
<原文结束>

# <翻译开始>
// 为了性能考虑，以下代码在数组中搜索路径，但并不保证并发安全。
# <翻译结束>


<原文开始>
// Global template variables including configuration.
<原文结束>

# <翻译开始>
// 全局模板变量，包括配置信息。
# <翻译结束>


<原文开始>
// Default template file for parsing.
<原文结束>

# <翻译开始>
// 默认用于解析的模板文件。
# <翻译结束>


<原文开始>
// Automatically encodes and provides safe html output, which is good for avoiding XSS.
<原文结束>

# <翻译开始>
// 自动进行编码并提供安全的HTML输出，有助于避免XSS攻击。
# <翻译结束>


<原文开始>
// DefaultConfig creates and returns a configuration object with default configurations.
<原文结束>

# <翻译开始>
// DefaultConfig 创建并返回一个包含默认配置的配置对象。
# <翻译结束>


<原文开始>
// SetConfig sets the configuration for view.
<原文结束>

# <翻译开始>
// SetConfig 设置视图的配置。
# <翻译结束>


<原文开始>
	// Clear global template object cache.
	// It's just cache, do not hesitate clearing it.
<原文结束>

# <翻译开始>
// 清除全局模板对象缓存。
// 这只是缓存，不必犹豫去清除它。
# <翻译结束>


<原文开始>
// SetConfigWithMap set configurations with map for the view.
<原文结束>

# <翻译开始>
// SetConfigWithMap 使用map设置视图的相关配置。
# <翻译结束>


<原文开始>
	// The m now is a shallow copy of m.
	// Any changes to m does not affect the original one.
	// A little tricky, isn't it?
<原文结束>

# <翻译开始>
// 现在的m是m的一个浅拷贝。
// 对m的任何改动都不会影响原始的那个m。
// 有点小巧妙，不是吗？
# <翻译结束>


<原文开始>
// Most common used configuration support for single view path.
<原文结束>

# <翻译开始>
// 最常用的单视图路径配置支持。
# <翻译结束>


<原文开始>
// SetPath sets the template directory path for template file search.
// The parameter `path` can be absolute or relative path, but absolute path is suggested.
<原文结束>

# <翻译开始>
// SetPath 设置模板文件搜索的目录路径。
// 参数 `path` 可以是绝对路径或相对路径，但建议使用绝对路径。
# <翻译结束>


<原文开始>
// AddPath adds an absolute or relative path to the search paths.
<原文结束>

# <翻译开始>
// AddPath 将一个绝对路径或相对路径添加到搜索路径中。
# <翻译结束>


<原文开始>
// realPath should be type of folder.
<原文结束>

# <翻译开始>
// realPath 应为文件夹类型。
# <翻译结束>


<原文开始>
// Assigns binds multiple global template variables to current view object.
// Note that it's not concurrent-safe, which means it would panic
// if it's called in multiple goroutines in runtime.
<原文结束>

# <翻译开始>
// Assigns 将多个全局模板变量绑定到当前视图对象。
// 注意，它不是并发安全的，这意味着如果在运行时多个goroutine中调用它，将会引发panic。
# <翻译结束>


<原文开始>
// Assign binds a global template variable to current view object.
// Note that it's not concurrent-safe, which means it would panic
// if it's called in multiple goroutines in runtime.
<原文结束>

# <翻译开始>
// Assign 将全局模板变量绑定到当前视图对象。
// 注意，它不是并发安全的，这意味着如果在运行时多个goroutine中调用它，将会导致panic。
# <翻译结束>


<原文开始>
// SetDefaultFile sets default template file for parsing.
<原文结束>

# <翻译开始>
// SetDefaultFile 设置用于解析的默认模板文件。
# <翻译结束>


<原文开始>
// GetDefaultFile returns default template file for parsing.
<原文结束>

# <翻译开始>
// GetDefaultFile 返回用于解析的默认模板文件。
# <翻译结束>


<原文开始>
// SetDelimiters sets customized delimiters for template parsing.
<原文结束>

# <翻译开始>
// SetDelimiters 设置用于模板解析的自定义分隔符。
# <翻译结束>


<原文开始>
// SetAutoEncode enables/disables automatically html encoding feature.
// When AutoEncode feature is enables, view engine automatically encodes and provides safe html output,
// which is good for avoid XSS.
<原文结束>

# <翻译开始>
// SetAutoEncode 用于开启或关闭自动HTML编码功能。
// 当自动编码功能开启时，视图引擎会自动进行编码并提供安全的HTML输出，
// 这有助于避免XSS（跨站脚本攻击）漏洞。
# <翻译结束>


<原文开始>
// BindFunc registers customized global template function named `name`
// with given function `function` to current view object.
// The `name` is the function name which can be called in template content.
<原文结束>

# <翻译开始>
// BindFunc 注册一个名为 `name` 的自定义全局模板函数到当前视图对象，
// 使用给定的 `function` 函数。在模板内容中，`name` 是可以被调用的函数名。
# <翻译结束>


<原文开始>
// BindFuncMap registers customized global template functions by map to current view object.
// The key of map is the template function name
// and the value of map is the address of customized function.
<原文结束>

# <翻译开始>
// BindFuncMap 通过映射注册自定义的全局模板函数到当前视图对象。
// 映射的键是模板函数名称
// 映射的值是自定义函数的地址。
# <翻译结束>


<原文开始>
// SetI18n binds i18n manager to current view engine.
<原文结束>

# <翻译开始>
// SetI18n 将 i18n 管理器绑定到当前视图引擎。
# <翻译结束>


<原文开始>
// Custom template delimiters.
<原文结束>

# <翻译开始>
// 自定义模板分隔符。
# <翻译结束>


<原文开始>
// I18n manager for the view.
<原文结束>

# <翻译开始>
// 视图的国际化管理器。
# <翻译结束>


<原文开始>
// Should be a directory.
<原文结束>

# <翻译开始>
// 应该是一个目录。
# <翻译结束>


<原文开始>
// Repeated path adding check.
<原文结束>

# <翻译开始>
// 重复路径添加检查。
# <翻译结束>

