
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
// Config is the configuration object for template engine.
<原文结束>

# <翻译开始>
// Config是模板引擎的配置对象。. md5:0c7a20a5c1f534d4
# <翻译结束>


<原文开始>
// Searching array for path, NOT concurrent-safe for performance purpose.
<原文结束>

# <翻译开始>
// 在数组中搜索路径，为了性能原因，非并发安全。. md5:536357ec68a07213
# <翻译结束>


<原文开始>
// Global template variables including configuration.
<原文结束>

# <翻译开始>
// 全局模板变量，包括配置信息。. md5:5f96c7a35c11b4b2
# <翻译结束>


<原文开始>
// Default template file for parsing.
<原文结束>

# <翻译开始>
// 默认的模板文件用于解析。. md5:41607c84f42fcf9d
# <翻译结束>


<原文开始>
// Custom template delimiters.
<原文结束>

# <翻译开始>
// 自定义模板分隔符。. md5:0a97ca0eda8842d4
# <翻译结束>


<原文开始>
// Automatically encodes and provides safe html output, which is good for avoiding XSS.
<原文结束>

# <翻译开始>
// 自动进行编码并提供安全的HTML输出，这对于防止XSS攻击很有帮助。. md5:ec33e2ef01aaf3d3
# <翻译结束>


<原文开始>
// I18n manager for the view.
<原文结束>

# <翻译开始>
// 视图的国际化管理器。. md5:7c90b657f5c4c28b
# <翻译结束>


<原文开始>
// DefaultConfig creates and returns a configuration object with default configurations.
<原文结束>

# <翻译开始>
// DefaultConfig 创建并返回一个使用默认配置的配置对象。. md5:27f0cf63ebd5dd9e
# <翻译结束>


<原文开始>
// SetConfig sets the configuration for view.
<原文结束>

# <翻译开始>
// SetConfig 设置视图的配置。. md5:44d304b99e74e865
# <翻译结束>


<原文开始>
	// Clear global template object cache.
	// It's just cache, do not hesitate clearing it.
<原文结束>

# <翻译开始>
// 清除全局模板对象缓存。
// 这只是一个缓存，不要犹豫清空它。
// md5:51c51fe68d143dd8
# <翻译结束>


<原文开始>
// SetConfigWithMap set configurations with map for the view.
<原文结束>

# <翻译开始>
// SetConfigWithMap 使用映射为视图设置配置。. md5:1e1d667c3b2ace2b
# <翻译结束>


<原文开始>
	// The m now is a shallow copy of m.
	// Any changes to m does not affect the original one.
	// A little tricky, isn't it?
<原文结束>

# <翻译开始>
// m 现在是 m 的浅拷贝。
// 对 m 的任何修改都不会影响原始对象。
// 这有点巧妙，不是吗？
// md5:4d1dd38c4db57a79
# <翻译结束>


<原文开始>
// Most common used configuration support for single view path.
<原文结束>

# <翻译开始>
// 最常用的单视图路径配置支持。. md5:4ebc24cd15a30d35
# <翻译结束>


<原文开始>
// SetPath sets the template directory path for template file search.
// The parameter `path` can be absolute or relative path, but absolute path is suggested.
<原文结束>

# <翻译开始>
// SetPath 设置模板文件搜索的目录路径。参数 `path` 可以是绝对路径或相对路径，但建议使用绝对路径。
// md5:abd751ab819d28b6
# <翻译结束>


<原文开始>
// Repeated path adding check.
<原文结束>

# <翻译开始>
// 重复路径添加检查。. md5:e210e91d65ec4857
# <翻译结束>


<原文开始>
// AddPath adds an absolute or relative path to the search paths.
<原文结束>

# <翻译开始>
// AddPath 向搜索路径中添加一个绝对或相对路径。. md5:d279479528c86f4e
# <翻译结束>


<原文开始>
// realPath should be type of folder.
<原文结束>

# <翻译开始>
// realPath 应该是文件夹类型的路径。. md5:8b57fae1c1158ae9
# <翻译结束>


<原文开始>
// Assigns binds multiple global template variables to current view object.
// Note that it's not concurrent-safe, which means it would panic
// if it's called in multiple goroutines in runtime.
<原文结束>

# <翻译开始>
// 将多个全局模板变量绑定到当前视图对象。需要注意的是，它不是并发安全的，这意味着如果在运行时从多个goroutine中调用它，会导致panic。
// md5:b31929b349e74390
# <翻译结束>


<原文开始>
// Assign binds a global template variable to current view object.
// Note that it's not concurrent-safe, which means it would panic
// if it's called in multiple goroutines in runtime.
<原文结束>

# <翻译开始>
// Assign 将全局模板变量绑定到当前视图对象。需要注意的是，它不是线程安全的，这意味着如果在运行时从多个goroutine中调用它，会导致panic。
// md5:7043c41fc2b3a0c3
# <翻译结束>


<原文开始>
// SetDefaultFile sets default template file for parsing.
<原文结束>

# <翻译开始>
// SetDefaultFile 为解析设置默认的模板文件。. md5:17f210ece0d189f6
# <翻译结束>


<原文开始>
// GetDefaultFile returns default template file for parsing.
<原文结束>

# <翻译开始>
// GetDefaultFile 返回默认的模板文件，用于解析。. md5:f72bb2dc04f3d4a4
# <翻译结束>


<原文开始>
// SetDelimiters sets customized delimiters for template parsing.
<原文结束>

# <翻译开始>
// SetDelimiters 设置模板解析的自定义分隔符。. md5:a09465c3518f1023
# <翻译结束>


<原文开始>
// SetAutoEncode enables/disables automatically html encoding feature.
// When AutoEncode feature is enables, view engine automatically encodes and provides safe html output,
// which is good for avoid XSS.
<原文结束>

# <翻译开始>
// SetAutoEncode 启用/禁用自动 HTML 编码功能。
// 当 AutoEncode 功能启用时，视图引擎会自动编码并提供安全的 HTML 输出，这对于防止 XSS 攻击很有好处。
// md5:cd0107f5d2170f4f
# <翻译结束>


<原文开始>
// BindFunc registers customized global template function named `name`
// with given function `function` to current view object.
// The `name` is the function name which can be called in template content.
<原文结束>

# <翻译开始>
// BindFunc 向当前视图对象注册一个名为 `name` 的自定义全局模板函数，
// 使用提供的 `function` 函数。其中，`name` 是在模板内容中可被调用的函数名。
// md5:20f79a4c8d0ba97a
# <翻译结束>


<原文开始>
// BindFuncMap registers customized global template functions by map to current view object.
// The key of map is the template function name
// and the value of map is the address of customized function.
<原文结束>

# <翻译开始>
// BindFuncMap 将自定义的全局模板函数通过映射注册到当前视图对象中。
// 映射的键是模板函数名称，
// 映射的值是自定义函数的地址。
// md5:2fe9bab0463cef27
# <翻译结束>


<原文开始>
// SetI18n binds i18n manager to current view engine.
<原文结束>

# <翻译开始>
// SetI18n 将i18n管理器绑定到当前视图引擎。. md5:8d1b88bd87c041ba
# <翻译结束>

