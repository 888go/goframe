
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
// RuleFunc is the custom function for data validation.
<原文结束>

# <翻译开始>
// RuleFunc 是用于数据验证的自定义函数。 md5:7988c41777832ac1
# <翻译结束>


<原文开始>
// RuleFuncInput holds the input parameters that passed to custom rule function RuleFunc.
<原文结束>

# <翻译开始>
// RuleFuncInput 是传递给自定义规则函数 RuleFunc 的输入参数。 md5:071da67c908f30a9
# <翻译结束>


<原文开始>
// Rule specifies the validation rule string, like "required", "between:1,100", etc.
<原文结束>

# <翻译开始>
// Rule 定义了验证规则字符串，例如 "required"、"between:1,100" 等等。 md5:0903f4201c9e300d
# <翻译结束>


<原文开始>
// Message specifies the custom error message or configured i18n message for this rule.
<原文结束>

# <翻译开始>
// Message 指定此规则的自定义错误消息或配置的 i18n 消息。 md5:407649d2c7943432
# <翻译结束>


<原文开始>
// Field specifies the field for this rule to validate.
<原文结束>

# <翻译开始>
// Field 指定此规则要验证的字段。 md5:b21049696367d3c3
# <翻译结束>


<原文开始>
// ValueType specifies the type of the value, which might be nil.
<原文结束>

# <翻译开始>
// ValueType 指定了值的类型，可能为 nil。 md5:b1ad5cfd9a152a1d
# <翻译结束>


<原文开始>
// Value specifies the value for this rule to validate.
<原文结束>

# <翻译开始>
// Value 指定此规则用于验证的值。 md5:29bdb57107181fe6
# <翻译结束>


<原文开始>
	// Data specifies the `data` which is passed to the Validator. It might be a type of map/struct or a nil value.
	// You can ignore the parameter `Data` if you do not really need it in your custom validation rule.
<原文结束>

# <翻译开始>
	// Data 指定了传递给Validator的数据，它可以是map/结构体类型或nil值。如果你的自定义验证规则不需要这个参数，可以忽略它。
	// md5:fd9ebb5b1bdabe03
# <翻译结束>


<原文开始>
	// customRuleFuncMap stores the custom rule functions.
	// map[Rule]RuleFunc
<原文结束>

# <翻译开始>
	// customRuleFuncMap 存储自定义规则函数。
	// map[Rule]RuleFunc
	// md5:ddde03f9fa92aae7
# <翻译结束>


<原文开始>
// RegisterRule registers custom validation rule and function for package.
<原文结束>

# <翻译开始>
// RegisterRule 为包注册自定义验证规则和函数。 md5:bb0c3971adfb8935
# <翻译结束>


<原文开始>
// RegisterRuleByMap registers custom validation rules using map for package.
<原文结束>

# <翻译开始>
// RegisterRuleByMap 通过映射为包注册自定义验证规则。 md5:6f3ae52bddfd4a24
# <翻译结束>


<原文开始>
// GetRegisteredRuleMap returns all the custom registered rules and associated functions.
<原文结束>

# <翻译开始>
// GetRegisteredRuleMap 返回所有自定义注册的规则及其关联的函数。 md5:3abbd0fbfe9f3c51
# <翻译结束>


<原文开始>
// DeleteRule deletes custom defined validation one or more rules and associated functions from global package.
<原文结束>

# <翻译开始>
// DeleteRule 从全局包中删除一个或多个自定义定义的验证规则及其关联函数。 md5:474d821f8f0b7fdc
# <翻译结束>

