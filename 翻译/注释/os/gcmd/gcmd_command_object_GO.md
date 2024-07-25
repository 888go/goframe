
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
// defaultValueTags is the struct tag names for default value storing.
<原文结束>

# <翻译开始>
	// defaultValueTags 是用于存储默认值的结构体标签名称。 md5:5ce6fb87a220db49
# <翻译结束>


<原文开始>
// NewFromObject creates and returns a root command object using given object.
<原文结束>

# <翻译开始>
// NewFromObject 使用给定的对象创建并返回一个根命令对象。 md5:3bdd362e3ec9f337
# <翻译结束>


<原文开始>
	// If given `object` is not pointer, it then creates a temporary one,
	// of which the value is `reflectValue`.
	// It then can retrieve all the methods both of struct/*struct.
<原文结束>

# <翻译开始>
	// 如果给定的`object`不是指针，那么它会创建一个临时的，其值为`reflectValue`。
	// 然后它可以获取结构体/`*struct`的所有方法。
	// md5:1e216cd9c7839ef2
# <翻译结束>


<原文开始>
// The `object` is the Meta attribute from business object, and the `name` is the command name,
// commonly from method name, which is used when no name tag is defined in Meta.
<原文结束>

# <翻译开始>
// `object`是业务对象的元属性，`name`是命令名称，通常来自方法名。当Meta中没有定义名称标签时使用这个。
// md5:044898119694fdf5
# <翻译结束>


<原文开始>
// Name field is necessary.
<原文结束>

# <翻译开始>
	// 名称字段是必需的。 md5:be70066859cce69e
# <翻译结束>


<原文开始>
// Necessary validation for input/output parameters and naming.
<原文结束>

# <翻译开始>
	// 对输入/输出参数及命名进行必要的验证。 md5:9e72ac9f4181fad5
# <翻译结束>


<原文开始>
// The input struct should be named as `xxxInput`.
<原文结束>

# <翻译开始>
	// 输入结构体应该命名为`xxxInput`。 md5:98fe3954e690f01f
# <翻译结束>


<原文开始>
// The output struct should be named as `xxxOutput`.
<原文结束>

# <翻译开始>
	// 输出结构体应该命名为`xxxOutput`。 md5:3c8e65a804dbe66c
# <翻译结束>


<原文开始>
// For input struct converting using priority tag.
<原文结束>

# <翻译开始>
	// 利用优先级标签进行输入结构体转换。 md5:501b3e1e29551f82
# <翻译结束>


<原文开始>
	// =============================================================================================
	// Create function that has value return.
	// =============================================================================================
<原文结束>

# <翻译开始>
	// =============================================================================================
	// 创建一个有返回值的函数。
	// =============================================================================================
	// md5:665eb5f5657321cc
# <翻译结束>


<原文开始>
// Use the left args to assign to input struct object.
<原文结束>

# <翻译开始>
			// 使用左面的参数来给输入结构体对象赋值。 md5:7a575ddae03dd0d6
# <翻译结束>


<原文开始>
// Read argument from command line index.
<原文结束>

# <翻译开始>
				// 从命令行索引读取参数。 md5:aa066778e00736af
# <翻译结束>


<原文开始>
// Read argument from command line option name.
<原文结束>

# <翻译开始>
				// 从命令行选项名称中读取参数。 md5:9fcf493c28108e47
# <翻译结束>


<原文开始>
							// Adapter with common user habits.
							// Eg:
							// `gf -f=0`: which parameter `f` is parsed as false
							// `gf -f=1`: which parameter `f` is parsed as true
<原文结束>

# <翻译开始>
							// 适配常见的用户习惯。
							// 例如：
							// `gf -f=0`：参数 `f` 被解析为 false
							// `gf -f=1`：参数 `f` 被解析为 true
							// md5:72432f87d0fc818b
# <翻译结束>


<原文开始>
// Default values from struct tag.
<原文结束>

# <翻译开始>
		// 来自结构体标签的默认值。 md5:13e4a73100683597
# <翻译结束>


<原文开始>
// Construct input parameters.
<原文结束>

# <翻译开始>
		// 构建输入参数。 md5:c74f2d54c503f98e
# <翻译结束>


<原文开始>
// Call handler with dynamic created parameter values.
<原文结束>

# <翻译开始>
		// 使用动态创建的参数值调用处理器。 md5:991efec71cdcc95a
# <翻译结束>


<原文开始>
// mergeDefaultStructValue merges the request parameters with default values from struct tag definition.
<原文结束>

# <翻译开始>
// mergeDefaultStructValue 将请求参数与结构体标签定义中的默认值合并。 md5:0a73ebb7f647201a
# <翻译结束>


<原文开始>
// If it already has value, it then ignores the default value.
<原文结束>

# <翻译开始>
			// 如果已经有值，那么它将忽略默认值。 md5:e95a88514a952418
# <翻译结束>

