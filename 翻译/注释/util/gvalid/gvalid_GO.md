
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Package gvalid implements powerful and useful data/form validation functionality.
<原文结束>

# <翻译开始>
// 包gvalid实现了强大且实用的数据/表单验证功能。 md5:e037cf7a2dd78c4c
# <翻译结束>


<原文开始>
// CustomMsg is the custom error message type,
// like: map[field] => string|map[rule]string
<原文结束>

# <翻译开始>
// CustomMsg 是自定义错误消息类型，例如：map[field] => string|map[rule]string md5:7623c4a13054d811
# <翻译结束>


<原文开始>
// fieldRule defined the alias name and rule string for specified field.
<原文结束>

# <翻译开始>
// fieldRule 定义了指定字段的别名名称和规则字符串。 md5:54ed236de61abbbc
# <翻译结束>


<原文开始>
// Alias name for the field.
<原文结束>

# <翻译开始>
// 字段的别名名称。 md5:c25526bbdb1ad925
# <翻译结束>


<原文开始>
// Rule string like: "max:6"
<原文结束>

# <翻译开始>
// 规则字符串格式，例如："max:6". md5:347084141516f2db
# <翻译结束>


<原文开始>
// Is this rule is from gmeta.Meta, which marks it as whole struct rule.
<原文结束>

# <翻译开始>
// 是否此规则来自gmeta.Meta，这表明它是整个结构体规则。 md5:32e5c4d3da4488ad
# <翻译结束>


<原文开始>
// Original kind of struct field, which is used for parameter type checks.
<原文结束>

# <翻译开始>
// 原始的结构体字段类型，用于参数类型检查。 md5:425bcc151df78ab4
# <翻译结束>


<原文开始>
// Type of struct field, which is used for parameter type checks.
<原文结束>

# <翻译开始>
// 结构体字段的类型，用于参数类型检查。 md5:04ad279a1f0f2f8d
# <翻译结束>


<原文开始>
// iNoValidation is an interface that marks current struct not validated by package `gvalid`.
<原文结束>

# <翻译开始>
// iNoValidation 是一个接口，用于标记当前结构体不被 `gvalid` 包验证。 md5:b80ca2c2ec6a9a3e
# <翻译结束>


<原文开始>
// regular expression pattern for single validation rule.
<原文结束>

# <翻译开始>
// 单个验证规则的正则表达式模式。 md5:38a42987a367551c
# <翻译结束>


<原文开始>
// rule name for internal invalid rules validation error.
<原文结束>

# <翻译开始>
// 用于内部无效规则验证错误的规则名称。 md5:4f1137b2b0bcbf87
# <翻译结束>


<原文开始>
// rule name for internal invalid params validation error.
<原文结束>

# <翻译开始>
// 用于内部无效参数验证错误的规则名称。 md5:891768620779d9d5
# <翻译结束>


<原文开始>
// rule name for internal invalid object validation error.
<原文结束>

# <翻译开始>
// 内部无效对象验证错误的规则名称。 md5:d86405319c6f33d2
# <翻译结束>


<原文开始>
// error map key for internal errors.
<原文结束>

# <翻译开始>
// error map 中用于内部错误的键。 md5:e69cd3c42d326301
# <翻译结束>


<原文开始>
// default rule name for i18n error message format if no i18n message found for specified error rule.
<原文结束>

# <翻译开始>
// 如果为指定的错误规则找不到i18n消息，则为i18n错误消息格式设置的默认规则名称。 md5:c390e0867660c4ca
# <翻译结束>


<原文开始>
// prefix string for each rule configuration in i18n content.
<原文结束>

# <翻译开始>
// i18n内容中每个规则配置的前缀字符串。 md5:0f1f87e48f3229f2
# <翻译结束>


<原文开始>
// no validation tag name for struct attribute.
<原文结束>

# <翻译开始>
// 结构体属性缺少验证标签名称。 md5:eb3058b8dac711c4
# <翻译结束>


<原文开始>
// the name for rule "regex"
<原文结束>

# <翻译开始>
// "regex"规则的名称. md5:4a58a8d172eb6158
# <翻译结束>


<原文开始>
// the name for rule "not-regex"
<原文结束>

# <翻译开始>
// "not-regex" 规则的名称. md5:f7723458e5697b5d
# <翻译结束>


<原文开始>
// the name for rule "foreach"
<原文结束>

# <翻译开始>
// "foreach"规则的名称. md5:3d97e3f1ec27986c
# <翻译结束>


<原文开始>
// the name for rule "bail"
<原文结束>

# <翻译开始>
// "bail"规则的名称. md5:e9d4d005416cc4b3
# <翻译结束>


<原文开始>
// Empty json string for array type.
<原文结束>

# <翻译开始>
// 空的json字符串，用于数组类型。 md5:977af1a23874089e
# <翻译结束>


<原文开始>
// Empty json string for object type.
<原文结束>

# <翻译开始>
// 对象类型的空json字符串。 md5:5c45918837cd2fe1
# <翻译结束>


<原文开始>
// requiredRulesPrefix specifies the rule prefix that must be validated even the value is empty (nil or empty).
<原文结束>

# <翻译开始>
// requiredRulesPrefix 指定必须验证的规则前缀，即使值为空（nil或空字符串）。 md5:be7bfaed0613daec
# <翻译结束>


<原文开始>
	// defaultErrorMessages is the default error messages.
	// Note that these messages are synchronized from ./i18n/en/validation.toml .
<原文结束>

# <翻译开始>
	// defaultErrorMessages 是默认的错误信息。
	// 注意，这些信息是从 ./i18n/en/validation.toml 文件同步而来的。 md5:373f31d6c37a48f9
# <翻译结束>


<原文开始>
// structTagPriority specifies the validation tag priority array.
<原文结束>

# <翻译开始>
	// structTagPriority 指定结构体标签的验证优先级数组。 md5:f41cb86d701dc7f1
# <翻译结束>


<原文开始>
// aliasNameTagPriority specifies the alias tag priority array.
<原文结束>

# <翻译开始>
	// aliasNameTagPriority 指定别名标签优先级数组。 md5:8c51a2951426a6c4
# <翻译结束>


<原文开始>
// all internal error keys.
<原文结束>

# <翻译开始>
	// 所有内部错误键。 md5:d6981aa171db620c
# <翻译结束>


<原文开始>
	// regular expression object for single rule
	// which is compiled just once and of repeatable usage.
<原文结束>

# <翻译开始>
	// 单个规则的正则表达式对象
	// 它仅编译一次，可重复使用。 md5:5d3b8b54080f71ba
# <翻译结束>


<原文开始>
	// decorativeRuleMap defines all rules that are just marked rules which have neither functional meaning
	// nor error messages.
<原文结束>

# <翻译开始>
	// decorativeRuleMap 定义了所有仅具有标记规则，既没有功能意义也没有错误信息的规则。 md5:d98db5ea3aaff41f
# <翻译结束>


<原文开始>
// ParseTagValue parses one sequence tag to field, rule and error message.
// The sequence tag is like: [alias@]rule[...#msg...]
<原文结束>

# <翻译开始>
// ParseTagValue 解析一个序列标签到字段、规则和错误消息。
// 序列标签的格式为：[别名@]规则[...#消息...] md5:c1a14088e6940223
# <翻译结束>


<原文开始>
// GetTags returns the validation tags.
<原文结束>

# <翻译开始>
// GetTags 返回验证标签。 md5:58fb30086314fe05
# <翻译结束>

