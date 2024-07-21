
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
// Field names to alias name map.
<原文结束>

# <翻译开始>
// 字段名称到别名映射的地图。 md5:cd4fee2326581f83
# <翻译结束>


<原文开始>
// It here must use gstructs.TagFields not gstructs.FieldMap to ensure error sequence.
<原文结束>

# <翻译开始>
// 此处必须使用gstructs.TagFields而不是gstructs.FieldMap，以确保错误顺序。 md5:7f18271929bfd060
# <翻译结束>


<原文开始>
// If there's no struct tag and validation rules, it does nothing and returns quickly.
<原文结束>

# <翻译开始>
// 如果没有结构体标签和验证规则，它什么也不做，快速返回。 md5:62043578db8966a6
# <翻译结束>


<原文开始>
// just for internally searching index purpose.
<原文结束>

# <翻译开始>
// 只是为了内部搜索索引目的。 md5:d89a7eb74470680a
# <翻译结束>


<原文开始>
// Custom rule error message map.
<原文结束>

# <翻译开始>
// 自定义规则错误消息映射。 md5:1489786362b1ed0d
# <翻译结束>


<原文开始>
// Ready to be validated data.
<原文结束>

# <翻译开始>
// 准备进行验证的数据。 md5:8be527aae62e9c7b
# <翻译结束>


<原文开始>
	// Sequence tag: []sequence tag
	// Sequence has order for error results.
<原文结束>

# <翻译开始>
	// 序列标签：[]序列标签
	// 序列中错误结果的顺序是有意义的。
	// md5:3ffc642de1ce88d6
# <翻译结束>


<原文开始>
					// If length of custom messages is lesser than length of rules,
					// the rest rules use the default error messages.
<原文结束>

# <翻译开始>
					// 如果自定义消息的长度小于规则的长度，那么剩余的规则将使用默认的错误消息。
					// md5:ada20f4d064fc46a
# <翻译结束>


<原文开始>
	// Map type rules does not support sequence.
	// Format: map[key]rule
<原文结束>

# <翻译开始>
	// 地图类型规则不支持序列。
	// 格式：map[key]rule
	// md5:e0dce9966e6f4666
# <翻译结束>


<原文开始>
// Input parameter map handling.
<原文结束>

# <翻译开始>
// 处理输入参数映射。 md5:4e321b50a9e44d75
# <翻译结束>


<原文开始>
// Checks and extends the parameters map with struct alias tag.
<原文结束>

# <翻译开始>
// 检查并使用结构体别名标签扩展参数映射。 md5:d6fb47b89d8c4795
# <翻译结束>


<原文开始>
	// Merge the custom validation rules with rules in struct tag.
	// The custom rules has the most high priority that can overwrite the struct tag rules.
<原文结束>

# <翻译开始>
	// 将自定义验证规则与结构体标签中的规则合并。
	// 自定义规则具有最高优先级，可以覆盖结构体标签中的规则。
	// md5:327ef56dd9382e55
# <翻译结束>


<原文开始>
// The `name` is different from `attribute alias`, which is used for validation only.
<原文结束>

# <翻译开始>
// `name`与`attribute alias`不同，后者仅用于验证。 md5:01baece1f454b49d
# <翻译结束>


<原文开始>
// It uses alias name of the attribute if its alias name tag exists.
<原文结束>

# <翻译开始>
// 如果属性存在别名标签，它将使用属性的别名名称。 md5:6b80790d910fc981
# <翻译结束>


<原文开始>
// It or else uses the attribute name directly.
<原文结束>

# <翻译开始>
// 否则，它直接使用属性名称。 md5:4b45cf8fb32210b5
# <翻译结束>


<原文开始>
// It uses the alias name from validation rule.
<原文结束>

# <翻译开始>
// 它使用了验证规则中的别名名称。 md5:e12f1e225f531883
# <翻译结束>


<原文开始>
		// It here extends the params map using alias names.
		// Note that the variable `name` might be alias name or attribute name.
<原文结束>

# <翻译开始>
		// 这里使用别名名称扩展params映射。
		// 注意变量`name`可能是别名名称或属性名称。
		// md5:67115358a00d1d8c
# <翻译结束>


<原文开始>
				// If there's alias name,
				// use alias name as its key and remove the field name key.
<原文结束>

# <翻译开始>
				// 如果有别名名称，
				// 使用别名名称作为其键，删除字段名称键。
				// md5:4454688e693edd27
# <翻译结束>


<原文开始>
// The input rules can overwrite the rules in struct tag.
<原文结束>

# <翻译开始>
// 输入的规则可以覆盖结构标签中的规则。 md5:2f6a125f9ce31c45
# <翻译结束>


<原文开始>
	// Custom error messages,
	// which have the most priority than `rules` and struct tag.
<原文结束>

# <翻译开始>
	// 自定义错误消息，
	// 其优先级高于 `rules` 和结构体标签。
	// md5:9ed0fde7c514e9ef
# <翻译结束>


<原文开始>
// Overwrite the key of field name.
<原文结束>

# <翻译开始>
// 替换字段名称的键。 md5:b77535a09a21fa98
# <翻译结束>


<原文开始>
// Temporary variable for value.
<原文结束>

# <翻译开始>
// 临时变量，用于存储值。 md5:5c2a7c202b7bf486
# <翻译结束>


<原文开始>
// It checks the struct recursively if its attribute is a struct/struct slice.
<原文结束>

# <翻译开始>
// 它会递归地检查结构体，以确定其属性是否为结构体或结构体切片。 md5:5461f7e14c7ea93f
# <翻译结束>


<原文开始>
// No validation interface implements check.
<原文结束>

# <翻译开始>
// 没有验证接口实现检查。 md5:ac37246e66f75369
# <翻译结束>


<原文开始>
// No validation field tag check.
<原文结束>

# <翻译开始>
// 不进行验证字段标签检查。 md5:70499912c8f45a5d
# <翻译结束>


<原文开始>
// The attributes of embedded struct are considered as direct attributes of its parent struct.
<原文结束>

# <翻译开始>
//嵌入的结构体的属性被视为其父结构体的直接属性。 md5:157af1f27cab37d6
# <翻译结束>


<原文开始>
// It merges the errors into single error map.
<原文结束>

# <翻译开始>
// 它将错误合并为单个错误映射。 md5:56fe32c627a507ee
# <翻译结束>


<原文开始>
			// The `field.TagValue` is the alias name of field.Name().
			// Eg, value from struct tag `p`.
<原文结束>

# <翻译开始>
			// `field.TagValue`是field.Name()的别名。
			// 例如，从结构体标签`p`获取的值。
			// md5:0b34b40285c5eb31
# <翻译结束>


<原文开始>
// Recursively check attribute slice/map.
<原文结束>

# <翻译开始>
// 递归检查属性切片/映射。 md5:da454f5b75f8a7ba
# <翻译结束>


<原文开始>
// The following logic is the same as some of CheckMap but with sequence support.
<原文结束>

# <翻译开始>
// 下面的逻辑与CheckMap中的一些逻辑相同，但增加了序列支持。 md5:8a807868be68e3c3
# <翻译结束>


<原文开始>
// Empty json string checks according to mapping field kind.
<原文结束>

# <翻译开始>
// 根据映射字段类型检查空的json字符串。 md5:e4223594884df6e0
# <翻译结束>


<原文开始>
// It checks each rule and its value in loop.
<原文结束>

# <翻译开始>
// 它在循环中检查每个规则及其值。 md5:5ab8f96747fbcec4
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
			// 仅在映射和结构体验证中：
			// 如果值为nil或空字符串，并且没有required*规则，
			// 则清除错误信息。
			// ============================================================
			// md5:4632db837be49942
# <翻译结束>

