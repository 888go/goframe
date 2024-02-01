
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
// Field names to alias name map.
<原文结束>

# <翻译开始>
// 字段名称到别名名称的映射。
# <翻译结束>


<原文开始>
// It here must use gstructs.TagFields not gstructs.FieldMap to ensure error sequence.
<原文结束>

# <翻译开始>
// 在这里必须使用gstructs.TagFields而不是gstructs.FieldMap，以确保错误顺序的正确性。
# <翻译结束>


<原文开始>
// If there's no struct tag and validation rules, it does nothing and returns quickly.
<原文结束>

# <翻译开始>
// 如果没有结构体标签和验证规则，它将不做任何操作并快速返回。
# <翻译结束>


<原文开始>
// just for internally searching index purpose.
<原文结束>

# <翻译开始>
// 仅用于内部搜索索引目的。
# <翻译结束>


<原文开始>
// Custom rule error message map.
<原文结束>

# <翻译开始>
// 自定义规则错误消息映射。
# <翻译结束>


<原文开始>
// Ready to be validated data.
<原文结束>

# <翻译开始>
// 准备就绪，等待验证的数据。
# <翻译结束>


<原文开始>
	// Sequence tag: []sequence tag
	// Sequence has order for error results.
<原文结束>

# <翻译开始>
// 序列标签: []序列标签
// 序列对错误结果有顺序要求。
# <翻译结束>


<原文开始>
					// If length of custom messages is lesser than length of rules,
					// the rest rules use the default error messages.
<原文结束>

# <翻译开始>
// 如果自定义消息的长度小于规则的长度，
// 剩余的规则将使用默认错误消息。
# <翻译结束>


<原文开始>
	// Map type rules does not support sequence.
	// Format: map[key]rule
<原文结束>

# <翻译开始>
// map类型规则不支持序列。
// 格式：map[key]rule
# <翻译结束>


<原文开始>
// Input parameter map handling.
<原文结束>

# <翻译开始>
// 输入参数映射处理。
# <翻译结束>


<原文开始>
// Checks and extends the parameters map with struct alias tag.
<原文结束>

# <翻译开始>
// 检查并使用结构体别名标签扩展参数映射。
# <翻译结束>


<原文开始>
	// Merge the custom validation rules with rules in struct tag.
	// The custom rules has the most high priority that can overwrite the struct tag rules.
<原文结束>

# <翻译开始>
// 将自定义验证规则与结构体标签中的规则进行合并。
// 自定义规则具有最高优先级，可以覆盖结构体标签中的规则。
# <翻译结束>


<原文开始>
// It uses alias name of the attribute if its alias name tag exists.
<原文结束>

# <翻译开始>
// 如果属性存在别名标签，则它使用该属性的别名名称。
# <翻译结束>


<原文开始>
// It or else uses the attribute name directly.
<原文结束>

# <翻译开始>
// 如果不使用属性名称直接作为键，则使用它
# <翻译结束>


<原文开始>
// It uses the alias name from validation rule.
<原文结束>

# <翻译开始>
// 它使用了验证规则中的别名名称。
# <翻译结束>


<原文开始>
		// It here extends the params map using alias names.
		// Note that the variable `name` might be alias name or attribute name.
<原文结束>

# <翻译开始>
// 这里通过别名扩展params映射。
// 注意，变量`name`可能是别名或属性名。
# <翻译结束>


<原文开始>
				// If there's alias name,
				// use alias name as its key and remove the field name key.
<原文结束>

# <翻译开始>
// 如果存在别名名称，
// 则使用别名名称作为键，并移除字段名称键。
# <翻译结束>


<原文开始>
// The input rules can overwrite the rules in struct tag.
<原文结束>

# <翻译开始>
// 输入的规则可以覆盖结构体标签中的规则。
# <翻译结束>


<原文开始>
				// If length of custom messages is lesser than length of rules,
				// the rest rules use the default error messages.
<原文结束>

# <翻译开始>
// 如果自定义消息的长度小于规则的长度，
// 剩余的规则将使用默认错误消息。
# <翻译结束>


<原文开始>
	// Custom error messages,
	// which have the most priority than `rules` and struct tag.
<原文结束>

# <翻译开始>
// 自定义错误消息，
// 这些错误消息具有比`rules`和结构体标签更高的优先级。
# <翻译结束>


<原文开始>
// Overwrite the key of field name.
<原文结束>

# <翻译开始>
// 覆盖字段名称的键。
# <翻译结束>


<原文开始>
// Temporary variable for value.
<原文结束>

# <翻译开始>
// 用于临时存储值的变量。
# <翻译结束>


<原文开始>
// It checks the struct recursively if its attribute is a struct/struct slice.
<原文结束>

# <翻译开始>
// 它递归检查结构体，如果其属性是结构体/结构体切片。
# <翻译结束>


<原文开始>
// No validation interface implements check.
<原文结束>

# <翻译开始>
// 没有验证接口实现了check。
# <翻译结束>


<原文开始>
// No validation field tag check.
<原文结束>

# <翻译开始>
// 不进行字段标签验证检查。
# <翻译结束>


<原文开始>
// It merges the errors into single error map.
<原文结束>

# <翻译开始>
// 它将错误合并到单个错误映射中。
# <翻译结束>


<原文开始>
			// The `field.TagValue` is the alias name of field.Name().
			// Eg, value from struct tag `p`.
<原文结束>

# <翻译开始>
// `field.TagValue` 是 `field.Name()` 的别名。
// 例如，来自结构体标签 `p` 中的值。
# <翻译结束>


<原文开始>
// Recursively check attribute slice/map.
<原文结束>

# <翻译开始>
// 递归检查属性切片/映射。
# <翻译结束>







<原文开始>
// The following logic is the same as some of CheckMap but with sequence support.
<原文结束>

# <翻译开始>
// 下面的逻辑与 CheckMap 的部分功能相同，但增加了对序列的支持。
# <翻译结束>


<原文开始>
// Empty json string checks according to mapping field kind.
<原文结束>

# <翻译开始>
// 根据映射字段类型检查空的json字符串。
# <翻译结束>


<原文开始>
// It checks each rule and its value in loop.
<原文结束>

# <翻译开始>
// 它在循环中检查每一条规则及其对应的值。
# <翻译结束>


<原文开始>
			// ============================================================
			// Only in map and struct validations:
			// If value is nil or empty string and has no required* rules,
			// it clears the error message.
			// ============================================================
<原文结束>

# <翻译开始>
// ============================================================
// 仅在map和struct验证中：
// 如果值为nil或空字符串且没有required*规则，
// 它将清除错误消息。
// ============================================================
# <翻译结束>












<原文开始>
// The `name` is different from `attribute alias`, which is used for validation only.
<原文结束>

# <翻译开始>
// `name`与用于验证的`attribute alias`不同。
# <翻译结束>

