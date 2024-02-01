
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
// defaultValueTags is the struct tag names for default value storing.
<原文结束>

# <翻译开始>
// defaultValueTags 是用于存储默认值的结构体标签名称。
# <翻译结束>


<原文开始>
// NewFromObject creates and returns a root command object using given object.
<原文结束>

# <翻译开始>
// NewFromObject 通过给定的对象创建并返回一个根命令对象。
# <翻译结束>


<原文开始>
	// If given `object` is not pointer, it then creates a temporary one,
	// of which the value is `reflectValue`.
	// It then can retrieve all the methods both of struct/*struct.
<原文结束>

# <翻译开始>
// 如果给定的`object`不是指针，它会创建一个临时指针，
// 其指向值为`reflectValue`。
// 然后可以获取结构体（包括结构体指针）的所有方法。
// 这段代码注释的翻译如下：
// ```go
// 如果传入的`object`不是一个指针类型，
// 则会创建一个临时指针变量，该指针指向`reflectValue`。
// 这样就可以获取到结构体及其指针类型的全部方法。
# <翻译结束>












<原文开始>
// The `object` is the Meta attribute from business object, and the `name` is the command name,
// commonly from method name, which is used when no name tag is defined in Meta.
<原文结束>

# <翻译开始>
// `object` 是业务对象中的 Meta 属性，而 `name` 是命令名称，
// 通常来源于方法名，在 Meta 中未定义 name 标签时使用。
# <翻译结束>







<原文开始>
// Necessary validation for input/output parameters and naming.
<原文结束>

# <翻译开始>
// 对输入/输出参数和命名进行必要的验证。
# <翻译结束>


<原文开始>
// The input struct should be named as `xxxInput`.
<原文结束>

# <翻译开始>
// 输入结构体应命名为 `xxxInput`。
# <翻译结束>


<原文开始>
// The output struct should be named as `xxxOutput`.
<原文结束>

# <翻译开始>
// 输出结构体的名称应命名为`xxxOutput`。
# <翻译结束>












<原文开始>
	// =============================================================================================
	// Create function that has value return.
	// =============================================================================================
<原文结束>

# <翻译开始>
// =============================================================================================
// 创建一个具有返回值的函数。
// =============================================================================================
# <翻译结束>







<原文开始>
// Read argument from command line index.
<原文结束>

# <翻译开始>
// 从命令行参数索引读取参数。
# <翻译结束>


<原文开始>
// Read argument from command line option name.
<原文结束>

# <翻译开始>
// 从命令行选项名称读取参数。
# <翻译结束>


<原文开始>
							// Adapter with common user habits.
							// Eg:
							// `gf -f=0`: which parameter `f` is parsed as false
							// `gf -f=1`: which parameter `f` is parsed as true
<原文结束>

# <翻译开始>
// 用户习惯适配器，包含通用用户习惯设置。
// 例如：
// `gf -f=0`：参数`f`将被解析为false
// `gf -f=1`：参数`f`将被解析为true
# <翻译结束>


<原文开始>
// Default values from struct tag.
<原文结束>

# <翻译开始>
// 结构体标签中的默认值。
# <翻译结束>


<原文开始>
// Construct input parameters.
<原文结束>

# <翻译开始>
// 构造输入参数。
# <翻译结束>







<原文开始>
// Call handler with dynamic created parameter values.
<原文结束>

# <翻译开始>
// 使用动态创建的参数值调用处理器。
# <翻译结束>


<原文开始>
// mergeDefaultStructValue merges the request parameters with default values from struct tag definition.
<原文结束>

# <翻译开始>
// mergeDefaultStructValue 将请求参数与来自结构体标签定义的默认值进行合并。
# <翻译结束>


<原文开始>
// If it already has value, it then ignores the default value.
<原文结束>

# <翻译开始>
// 如果它已经有了值，则会忽略默认值。
# <翻译结束>


<原文开始>
// Root command creating.
<原文结束>

# <翻译开始>
// 创建根命令
# <翻译结束>


<原文开始>
// Sub command creating.
<原文结束>

# <翻译开始>
// 子命令创建
# <翻译结束>


<原文开始>
// Name field is necessary.
<原文结束>

# <翻译开始>
// Name字段是必需的。
# <翻译结束>


<原文开始>
// Handle orphan options.
<原文结束>

# <翻译开始>
// 处理孤立选项。
# <翻译结束>


<原文开始>
// Parameters validation.
<原文结束>

# <翻译开始>
// 参数验证。
# <翻译结束>

