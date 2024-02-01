
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
// Package gvalid implements powerful and useful data/form validation functionality.
<原文结束>

# <翻译开始>
// Package gvalid 提供了强大且实用的数据/表单验证功能。
# <翻译结束>


<原文开始>
// CustomMsg is the custom error message type,
// like: map[field] => string|map[rule]string
<原文结束>

# <翻译开始>
// CustomMsg 是自定义错误消息类型，
// 形如：map[field] => string|map[rule]string
// （注释解释：此代码定义了一个名为 CustomMsg 的自定义错误消息类型，它是一个映射结构，键为字段名，值可以是字符串或另一个映射结构，其中内部映射的键为规则名，值为字符串。这个类型用于根据特定字段和规则生成或存储自定义的错误消息。）
# <翻译结束>


<原文开始>
// fieldRule defined the alias name and rule string for specified field.
<原文结束>

# <翻译开始>
// fieldRule 定义了特定字段的别名名称和规则字符串。
# <翻译结束>







<原文开始>
// Is this rule is from gmeta.Meta, which marks it as whole struct rule.
<原文结束>

# <翻译开始>
// 这个规则是否来自gmeta.Meta，将其标记为整个结构体规则。
# <翻译结束>


<原文开始>
// Original kind of struct field, which is used for parameter type checks.
<原文结束>

# <翻译开始>
// 原始的结构体字段类型，用于参数类型检查。
# <翻译结束>


<原文开始>
// Type of struct field, which is used for parameter type checks.
<原文结束>

# <翻译开始>
// 结构体字段的类型，用于参数类型的检查。
# <翻译结束>


<原文开始>
// iNoValidation is an interface that marks current struct not validated by package `gvalid`.
<原文结束>

# <翻译开始>
// iNoValidation 是一个接口，用于标记当前结构体未通过 `gvalid` 包进行验证。
# <翻译结束>


<原文开始>
// regular expression pattern for single validation rule.
<原文结束>

# <翻译开始>
// 正则表达式模式，用于单个验证规则。
# <翻译结束>


<原文开始>
// rule name for internal invalid rules validation error.
<原文结束>

# <翻译开始>
// 规则名称，用于内部无效规则验证错误。
# <翻译结束>


<原文开始>
// rule name for internal invalid params validation error.
<原文结束>

# <翻译开始>
// 内部无效参数验证错误的规则名称。
# <翻译结束>


<原文开始>
// rule name for internal invalid object validation error.
<原文结束>

# <翻译开始>
// 内部无效对象验证错误的规则名称。
# <翻译结束>


<原文开始>
// error map key for internal errors.
<原文结束>

# <翻译开始>
// 错误映射键，用于内部错误。
# <翻译结束>


<原文开始>
// default rule name for i18n error message format if no i18n message found for specified error rule.
<原文结束>

# <翻译开始>
// 如果未找到指定错误规则的国际化消息，则为此提供默认规则名称，用于国际化错误消息格式。
# <翻译结束>


<原文开始>
// prefix string for each rule configuration in i18n content.
<原文结束>

# <翻译开始>
// 在i18n内容中，每一项规则配置前缀字符串。
# <翻译结束>


<原文开始>
// no validation tag name for struct attribute.
<原文结束>

# <翻译开始>
// 对结构体属性没有验证标签名称。
# <翻译结束>


<原文开始>
// Empty json string for array type.
<原文结束>

# <翻译开始>
// 空的json字符串，用于数组类型。
# <翻译结束>


<原文开始>
// Empty json string for object type.
<原文结束>

# <翻译开始>
// 空的JSON字符串，用于对象类型。
# <翻译结束>


<原文开始>
// requiredRulesPrefix specifies the rule prefix that must be validated even the value is empty (nil or empty).
<原文结束>

# <翻译开始>
// requiredRulesPrefix 指定即使值为空（nil 或空）也必须进行验证的规则前缀。
# <翻译结束>


<原文开始>
	// defaultErrorMessages is the default error messages.
	// Note that these messages are synchronized from ./i18n/en/validation.toml .
<原文结束>

# <翻译开始>
// defaultErrorMessages 是默认的错误消息集合。
// 注意，这些消息是从 ./i18n/en/validation.toml 文件同步过来的。
# <翻译结束>


<原文开始>
// structTagPriority specifies the validation tag priority array.
<原文结束>

# <翻译开始>
// structTagPriority 指定了验证标签优先级数组。
# <翻译结束>


<原文开始>
// aliasNameTagPriority specifies the alias tag priority array.
<原文结束>

# <翻译开始>
// aliasNameTagPriority 指定了别名标签优先级数组。
# <翻译结束>







<原文开始>
	// regular expression object for single rule
	// which is compiled just once and of repeatable usage.
<原文结束>

# <翻译开始>
// 正则表达式对象，用于单个规则
// 该对象仅编译一次，且可用于重复使用。
# <翻译结束>


<原文开始>
	// decorativeRuleMap defines all rules that are just marked rules which have neither functional meaning
	// nor error messages.
<原文结束>

# <翻译开始>
// decorativeRuleMap 定义了所有仅作为标记规则的规则，这些规则既没有功能含义，也没有错误消息。
# <翻译结束>


<原文开始>
// ParseTagValue parses one sequence tag to field, rule and error message.
// The sequence tag is like: [alias@]rule[...#msg...]
<原文结束>

# <翻译开始>
// ParseTagValue解析一个序列标签到字段、规则和错误消息。
// 序列标签格式类似：[别名@]规则[...#消息...]
# <翻译结束>


<原文开始>
// GetTags returns the validation tags.
<原文结束>

# <翻译开始>
// GetTags 返回验证标签。
# <翻译结束>












<原文开始>
// the name for rule "not-regex"
<原文结束>

# <翻译开始>
// "not-regex"规则的名称
# <翻译结束>


<原文开始>
// the name for rule "foreach"
<原文结束>

# <翻译开始>
// “foreach”规则的名称
# <翻译结束>












<原文开始>
// Alias name for the field.
<原文结束>

# <翻译开始>
// 字段的别名名称。
# <翻译结束>


<原文开始>
// Rule string like: "max:6"
<原文结束>

# <翻译开始>
// Rule 字符串格式如："max:6"
# <翻译结束>


<原文开始>
// the name for rule "regex"
<原文结束>

# <翻译开始>
// 正则规则的名称为 "regex"
# <翻译结束>


<原文开始>
// the name for rule "bail"
<原文结束>

# <翻译开始>
// "bail"规则的名称
# <翻译结束>


<原文开始>
// the name for rule "ci"
<原文结束>

# <翻译开始>
// "ci"规则的名称
# <翻译结束>


<原文开始>
// all internal error keys.
<原文结束>

# <翻译开始>
// 所有内部错误键。
# <翻译结束>

