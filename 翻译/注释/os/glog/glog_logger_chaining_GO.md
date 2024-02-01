
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
// To is a chaining function,
// which redirects current logging content output to the specified `writer`.
<原文结束>

# <翻译开始>
// To 是一个链式函数，
// 它将当前日志内容输出重定向到指定的 `writer`。
# <翻译结束>


<原文开始>
// Path is a chaining function,
// which sets the directory path to `path` for current logging content output.
//
// Note that the parameter `path` is a directory path, not a file path.
<原文结束>

# <翻译开始>
// Path 是一个链式函数，
// 用于将当前日志内容输出的目标目录路径设置为 `path`。
//
// 注意，参数 `path` 应为一个目录路径，而非文件路径。
# <翻译结束>


<原文开始>
// Cat is a chaining function,
// which sets the category to `category` for current logging content output.
// Param `category` can be hierarchical, eg: module/user.
<原文结束>

# <翻译开始>
// Cat 是一个链式函数，
// 它用于为当前日志内容输出设置类别为 `category`。
// 参数 `category` 可以是层级式的，例如：module/user。
# <翻译结束>


<原文开始>
// File is a chaining function,
// which sets file name `pattern` for the current logging content output.
<原文结束>

# <翻译开始>
// File 是一个链式函数，
// 用于设置当前日志内容输出的文件名模式。
# <翻译结束>


<原文开始>
// Level is a chaining function,
// which sets logging level for the current logging content output.
<原文结束>

# <翻译开始>
// Level 是一个链式函数，
// 用于设置当前日志内容输出的记录级别。
# <翻译结束>


<原文开始>
// LevelStr is a chaining function,
// which sets logging level for the current logging content output using level string.
<原文结束>

# <翻译开始>
// LevelStr 是一个链式函数，
// 通过级别字符串设置当前日志内容输出的记录级别。
# <翻译结束>


<原文开始>
// Skip is a chaining function,
// which sets stack skip for the current logging content output.
// It also affects the caller file path checks when line number printing enabled.
<原文结束>

# <翻译开始>
// Skip 是一个链式函数，
// 用于设置当前日志内容输出时的堆栈跳过级别。
// 当启用行号打印时，它也会影响调用文件路径的检查。
# <翻译结束>


<原文开始>
// Stack is a chaining function,
// which sets stack options for the current logging content output .
<原文结束>

# <翻译开始>
// Stack 是一个链式函数，
// 用于为当前日志内容输出设置堆栈选项。
# <翻译结束>


<原文开始>
// StackWithFilter is a chaining function,
// which sets stack filter for the current logging content output .
<原文结束>

# <翻译开始>
// StackWithFilter 是一个链式函数，
// 用于为当前日志内容输出设置堆栈过滤器。
# <翻译结束>


<原文开始>
// Stdout is a chaining function,
// which enables/disables stdout for the current logging content output.
// It's enabled in default.
<原文结束>

# <翻译开始>
// Stdout 是一个链式函数，
// 用于启用/禁用当前日志内容输出到标准输出（stdout）。
// 默认情况下它是启用的。
# <翻译结束>


<原文开始>
// stdout printing is enabled if `enabled` is not passed.
<原文结束>

# <翻译开始>
// 如果未传递 `enabled`，则启用标准输出打印。
# <翻译结束>


<原文开始>
// Header is a chaining function,
// which enables/disables log header for the current logging content output.
// It's enabled in default.
<原文结束>

# <翻译开始>
// Header 是一个链式函数，
// 用于启用/禁用当前日志内容输出的头部信息。
// 默认情况下它是启用的。
# <翻译结束>


<原文开始>
// header is enabled if `enabled` is not passed.
<原文结束>

# <翻译开始>
// 如果未传递`enabled`参数，则启用header。
# <翻译结束>


<原文开始>
// Line is a chaining function,
// which enables/disables printing its caller file path along with its line number.
// The parameter `long` specified whether print the long absolute file path, eg: /a/b/c/d.go:23,
// or else short one: d.go:23.
<原文结束>

# <翻译开始>
// Line 是一个链式函数，
// 它用于启用/禁用在输出时附带调用者所在的文件路径及行号。
// 参数 `long` 指定了是否打印完整的绝对文件路径，例如：/a/b/c/d.go:23，
// 否则只打印简短形式：d.go:23。
# <翻译结束>


<原文开始>
// Async is a chaining function,
// which enables/disables async logging output feature.
<原文结束>

# <翻译开始>
// Async 是一个链式函数，
// 用于启用/禁用异步日志输出功能。
# <翻译结束>


<原文开始>
// async feature is enabled if `enabled` is not passed.
<原文结束>

# <翻译开始>
// 如果未传递`enabled`参数，则启用异步功能。
# <翻译结束>

