
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
// RuleFunc is the custom function for data validation.
<原文结束>

# <翻译开始>
// RuleFunc 是用于数据验证的自定义函数。
# <翻译结束>


<原文开始>
// RuleFuncInput holds the input parameters that passed to custom rule function RuleFunc.
<原文结束>

# <翻译开始>
// RuleFuncInput 用于存储传递给自定义规则函数 RuleFunc 的输入参数。
# <翻译结束>


<原文开始>
// Rule specifies the validation rule string, like "required", "between:1,100", etc.
<原文结束>

# <翻译开始>
// Rule 指定验证规则字符串，如 "必填", "范围:1,100" 等等。
# <翻译结束>


<原文开始>
// Message specifies the custom error message or configured i18n message for this rule.
<原文结束>

# <翻译开始>
// Message 指定此规则的自定义错误消息或配置的 i18n 消息。
# <翻译结束>


<原文开始>
// Field specifies the field for this rule to validate.
<原文结束>

# <翻译开始>
// Field 指定该规则进行验证的字段。
# <翻译结束>


<原文开始>
// ValueType specifies the type of the value, which might be nil.
<原文结束>

# <翻译开始>
// ValueType 指定值的类型，该值可能为 nil。
# <翻译结束>


<原文开始>
// Value specifies the value for this rule to validate.
<原文结束>

# <翻译开始>
// Value 指定此规则验证的值。
# <翻译结束>


<原文开始>
	// Data specifies the `data` which is passed to the Validator. It might be a type of map/struct or a nil value.
	// You can ignore the parameter `Data` if you do not really need it in your custom validation rule.
<原文结束>

# <翻译开始>
// Data 指定传递给 Validator 的 `data`，它可以是 map 或 struct 类型，也可以是 nil 值。
// 如果在自定义验证规则中并不真正需要这个参数，你可以忽略 `Data`。
# <翻译结束>


<原文开始>
	// customRuleFuncMap stores the custom rule functions.
	// map[Rule]RuleFunc
<原文结束>

# <翻译开始>
// customRuleFuncMap 用于存储自定义规则函数。
// map[Rule]RuleFunc 表示键为 Rule 类型，值为 RuleFunc 类型的映射表。
# <翻译结束>


<原文开始>
// RegisterRule registers custom validation rule and function for package.
<原文结束>

# <翻译开始>
// RegisterRule 注册自定义验证规则及其相关函数，供包内使用。
# <翻译结束>


<原文开始>
// RegisterRuleByMap registers custom validation rules using map for package.
<原文结束>

# <翻译开始>
// RegisterRuleByMap 通过映射表为包注册自定义验证规则。
# <翻译结束>


<原文开始>
// GetRegisteredRuleMap returns all the custom registered rules and associated functions.
<原文结束>

# <翻译开始>
// GetRegisteredRuleMap 返回所有已注册的自定义规则及其关联的函数。
# <翻译结束>


<原文开始>
// DeleteRule deletes custom defined validation one or more rules and associated functions from global package.
<原文结束>

# <翻译开始>
// DeleteRule 从全局包中删除一个或多个自定义验证规则及其关联函数。
# <翻译结束>

