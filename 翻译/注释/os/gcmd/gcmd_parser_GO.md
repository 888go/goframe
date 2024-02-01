
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//
# <翻译结束>


<原文开始>
// ParserOption manages the parsing options.
<原文结束>

# <翻译开始>
// ParserOption 管理解析选项。
# <翻译结束>


<原文开始>
// Marks options parsing in case-sensitive way.
<原文结束>

# <翻译开始>
// 以区分大小写的方式标记选项解析
# <翻译结束>


<原文开始>
// Whether stops parsing and returns error if invalid option passed.
<原文结束>

# <翻译开始>
// 如果传递了无效的选项，则停止解析并返回错误。
# <翻译结束>












<原文开始>
// Command function map for function handler.
<原文结束>

# <翻译开始>
// Command 函数映射，用于函数处理器。
# <翻译结束>


<原文开始>
// ParserFromCtx retrieves and returns Parser from context.
<原文结束>

# <翻译开始>
// ParserFromCtx 从上下文中检索并返回 Parser。
# <翻译结束>


<原文开始>
// Parse creates and returns a new Parser with os.Args and supported options.
//
// Note that the parameter `supportedOptions` is as [option name: need argument], which means
// the value item of `supportedOptions` indicates whether corresponding option name needs argument or not.
//
// The optional parameter `strict` specifies whether stops parsing and returns error if invalid option passed.
<原文结束>

# <翻译开始>
// Parse 创建并返回一个新的 Parser，其参数为 os.Args 以及支持的选项。
//
// 注意，参数 `supportedOptions` 形式为 [选项名: 是否需要参数]，这意味着
// `supportedOptions` 的值项表示对应的选项名是否需要参数。
//
// 可选参数 `strict` 指定在遇到无效选项时是否停止解析并返回错误。
# <翻译结束>


<原文开始>
// ParseArgs creates and returns a new Parser with given arguments and supported options.
//
// Note that the parameter `supportedOptions` is as [option name: need argument], which means
// the value item of `supportedOptions` indicates whether corresponding option name needs argument or not.
//
// The optional parameter `strict` specifies whether stops parsing and returns error if invalid option passed.
<原文结束>

# <翻译开始>
// ParseArgs 创建并返回一个新的解析器，该解析器包含给定的参数及支持的选项。
//
// 注意，参数 `supportedOptions` 形式为 [选项名称: 是否需要参数]，这意味着
// `supportedOptions` 中的值项表示对应的选项名称是否需要参数。
//
// 可选参数 `strict` 指定了当遇到无效选项时，是否停止解析并返回错误。
# <翻译结束>







<原文开始>
// parseMultiOption parses option to multiple valid options like: --dav.
// It returns nil if given option is not multi-option.
<原文结束>

# <翻译开始>
// parseMultiOption解析选项为多个有效选项，如：--dav。
// 如果给定的选项不是多选项，则返回nil。
# <翻译结束>







<原文开始>
// setOptionValue sets the option value for name and according alias.
<原文结束>

# <翻译开始>
// setOptionValue 为名称name及其别名设置选项值。
# <翻译结束>


<原文开始>
// GetOpt returns the option value named `name` as gvar.Var.
<原文结束>

# <翻译开始>
// GetOpt 函数返回名为 `name` 的选项值，类型为 gvar.Var。
# <翻译结束>


<原文开始>
// GetOptAll returns all parsed options.
<原文结束>

# <翻译开始>
// GetOptAll 返回所有已解析的选项。
# <翻译结束>


<原文开始>
// GetArg returns the argument at `index` as gvar.Var.
<原文结束>

# <翻译开始>
// GetArg 返回位于`index`处的参数作为gvar.Var类型。
# <翻译结束>


<原文开始>
// GetArgAll returns all parsed arguments.
<原文结束>

# <翻译开始>
// GetArgAll 返回所有已解析的参数。
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
# <翻译结束>


<原文开始>
// User passed supported options, like: map[string]bool{"name,n":true}
<原文结束>

# <翻译开始>
// 用户传入支持的选项，格式如：map[string]bool{"name,n":true}
// （该行代码注释表明，用户可以通过一个字符串到布尔值的映射表来传递支持的选项，其中键（key）为选项名，可能包含逗号分隔的别名，值（value）为true表示该选项被启用或选中。例如，在这个示例中，“name,n”表示选项名为"name"且其别名为 "n"，并且这个选项是被支持并启用的。）
# <翻译结束>


<原文开始>
// Option [OptionName:WhetherNeedArgument], like: map[string]bool{"name":true, "n":true}
<原文结束>

# <翻译开始>
// Option [选项名称:是否需要参数], 格式如：map[string]bool{"name":true, "n":true}
// 其中，键(如"name"和"n")代表选项名称，值(true或false)表示该选项在使用时是否需要携带参数。
# <翻译结束>







<原文开始>
// Parser for arguments.
<原文结束>

# <翻译开始>
// 参数解析器
# <翻译结束>


<原文开始>
// As name described.
<原文结束>

# <翻译开始>
// 如名称所述。
# <翻译结束>

