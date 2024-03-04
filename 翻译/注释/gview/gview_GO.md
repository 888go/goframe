
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
// Package gview implements a template engine based on text/template.
//
// Reserved template variable names:
// I18nLanguage: Assign this variable to define i18n language for each page.
<原文结束>

# <翻译开始>
// Package gview 实现了一个基于 text/template 的模板引擎。
//
// 预留的模板变量名称：
// I18nLanguage: 将此变量赋值以在每一页上定义 i18n 语言。
// 这段 Go 代码注释翻译成中文后为：
// ```go
// 包 gview 实现了一个基于 text/template 标准库的模板引擎功能。
//
// 已保留的模板变量名称：
// I18nLanguage：将该变量进行赋值，以便在每个页面上定义国际化的（i18n）语言设置。
# <翻译结束>


<原文开始>
// View object for template engine.
<原文结束>

# <翻译开始>
// 模板引擎的视图对象。
# <翻译结束>


<原文开始>
// Searching array for path, NOT concurrent-safe for performance purpose.
<原文结束>

# <翻译开始>
// 为了性能考虑，以下代码在数组中搜索路径，但并不保证并发安全。
# <翻译结束>


<原文开始>
// Global template function map.
<原文结束>

# <翻译开始>
// 全局模板函数映射。
# <翻译结束>


<原文开始>
// Extra configuration for the view.
<原文结束>

# <翻译开始>
// 额外的视图配置
# <翻译结束>


<原文开始>
// Params is type for template params.
<原文结束>

# <翻译开始>
// Params 是模板参数的类型。
# <翻译结束>


<原文开始>
// FuncMap is type for custom template functions.
<原文结束>

# <翻译开始>
// FuncMap 是用于自定义模板函数的类型。
# <翻译结束>


<原文开始>
// checkAndInitDefaultView checks and initializes the default view object.
// The default view object will be initialized just once.
<原文结束>

# <翻译开始>
// checkAndInitDefaultView 检查并初始化默认视图对象。
// 默认视图对象仅会被初始化一次。
# <翻译结束>


<原文开始>
// ParseContent parses the template content directly using the default view object
// and returns the parsed content.
<原文结束>

# <翻译开始>
// ParseContent 使用默认视图对象直接解析模板内容，
// 并返回已解析的内容。
# <翻译结束>


<原文开始>
// New returns a new view object.
// The parameter `path` specifies the template directory path to load template files.
<原文结束>

# <翻译开始>
// New返回一个新的视图对象。
// 参数`path`用于指定加载模板文件的模板目录路径。
# <翻译结束>


<原文开始>
// Customized dir path from env/cmd.
<原文结束>

# <翻译开始>
// 从环境变量/命令行自定义目录路径。
# <翻译结束>


<原文开始>
// Global template variables.
<原文结束>

# <翻译开始>
// 全局模板变量。
# <翻译结束>


<原文开始>
// Default view object.
<原文结束>

# <翻译开始>
// 默认视图对象。
# <翻译结束>


<原文开始>
// Dir path of working dir.
<原文结束>

# <翻译开始>
// Dir：工作目录的路径。
# <翻译结束>


<原文开始>
// Dir path of binary.
<原文结束>

# <翻译开始>
// Dir 二进制文件的路径。
# <翻译结束>


<原文开始>
// Dir path of main package.
<原文结束>

# <翻译开始>
// Dir：主包的路径。
# <翻译结束>


<原文开始>
// default build-in variables.
<原文结束>

# <翻译开始>
// 默认内置变量
# <翻译结束>


<原文开始>
// default build-in functions.
<原文结束>

# <翻译开始>
// 默认内置函数
# <翻译结束>

