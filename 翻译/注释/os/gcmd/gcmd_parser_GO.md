
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31
# <翻译结束>


<原文开始>
// ParserOption manages the parsing options.
<原文结束>

# <翻译开始>
// ParserOption负责管理解析选项。. md5:6294496b49d5c3bb
# <翻译结束>


<原文开始>
// Marks options parsing in case-sensitive way.
<原文结束>

# <翻译开始>
// 以区分大小写的方式标记选项解析。. md5:8d9524f23421dc60
# <翻译结束>


<原文开始>
// Whether stops parsing and returns error if invalid option passed.
<原文结束>

# <翻译开始>
// 如果传入了无效的选项，是否停止解析并返回错误。. md5:2564adc332d2fd51
# <翻译结束>


<原文开始>
// User passed supported options, like: map[string]bool{"name,n":true}
<原文结束>

# <翻译开始>
// 用户传递的受支持选项，如：map[string]bool{"name,n":true}. md5:ae5a3d920682c314
# <翻译结束>


<原文开始>
// Option [OptionName:WhetherNeedArgument], like: map[string]bool{"name":true, "n":true}
<原文结束>

# <翻译开始>
// 选项 [OptionName:是否需要参数]，例如：map[string]bool{"name":true, "n":true}. md5:d57dd0851c5ab783
# <翻译结束>


<原文开始>
// Command function map for function handler.
<原文结束>

# <翻译开始>
// 函数处理程序的命令函数映射。. md5:0061f09955d9b987
# <翻译结束>


<原文开始>
// ParserFromCtx retrieves and returns Parser from context.
<原文结束>

# <翻译开始>
// ParserFromCtx 从上下文中检索并返回解析器。. md5:260bf6b7d06ebc7c
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
// Parse 创建并返回一个新的Parser，使用os.Args和受支持的选项。
//
// 请注意，参数`supportedOptions`是[key: need argument]形式，其中
// `supportedOptions`的值项表示相应的选项名是否需要参数。
//
// 可选参数`strict`指定如果遇到无效选项时，是否停止解析并返回错误。
// md5:136e728aecd2a3b5
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
// ParseArgs 创建并返回一个新的Parser，具有给定的参数和支持的选项。
// 
// 注意，参数`supportedOptions`是一个[选项名称: 需要参数]的映射，这意味着`supportedOptions`的值项表示对应选项名称是否需要参数。
// 
// 可选参数`strict`指定是否在遇到无效选项时停止解析并返回错误。
// md5:5c367c6c4d6d78be
# <翻译结束>


<原文开始>
// parseMultiOption parses option to multiple valid options like: --dav.
// It returns nil if given option is not multi-option.
<原文结束>

# <翻译开始>
// parseMultiOption 解析多个有效选项，如：--dav。如果给定的选项不是多选项，它将返回nil。
// md5:d70d0f096bf48cc4
# <翻译结束>


<原文开始>
// setOptionValue sets the option value for name and according alias.
<原文结束>

# <翻译开始>
// setOptionValue 为名称和相应的别名设置选项值。. md5:9b55fb71d527f4c6
# <翻译结束>


<原文开始>
// Accurate option name match.
<原文结束>

# <翻译开始>
// 准确的选项名称匹配。. md5:92eb07ef58b2270c
# <翻译结束>


<原文开始>
// Fuzzy option name match.
<原文结束>

# <翻译开始>
// 模糊选项名称匹配。. md5:84dde7ce64941c27
# <翻译结束>


<原文开始>
// GetOpt returns the option value named `name` as gvar.Var.
<原文结束>

# <翻译开始>
// GetOpt 作为gvar.Var返回名为`name`的选项值。. md5:1859b868ee779be0
# <翻译结束>


<原文开始>
// GetOptAll returns all parsed options.
<原文结束>

# <翻译开始>
// GetOptAll 返回所有已解析的选项。. md5:6de4d266d8991786
# <翻译结束>


<原文开始>
// GetArg returns the argument at `index` as gvar.Var.
<原文结束>

# <翻译开始>
// GetArg 作为gvar.Var返回索引为`index`的参数。. md5:12ea2f8d74c6370d
# <翻译结束>


<原文开始>
// GetArgAll returns all parsed arguments.
<原文结束>

# <翻译开始>
// GetArgAll 返回所有解析的参数。. md5:85cc0fd5995d4878
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。. md5:43c3b36e60a18f9a
# <翻译结束>

