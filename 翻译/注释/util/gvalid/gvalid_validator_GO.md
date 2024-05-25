
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
// Validator is the validation manager for chaining operations.
<原文结束>

# <翻译开始>
// Validator是用于链式操作的验证管理器。 md5:4554cd1e10f5c88e
# <翻译结束>


<原文开始>
// I18n manager for error message translation.
<原文结束>

# <翻译开始>
// 用于错误消息翻译的国际化管理器。 md5:cc3a7d5d034e574f
# <翻译结束>


<原文开始>
// Validation data, which can be a map, struct or a certain value to be validated.
<原文结束>

# <翻译开始>
// 验证数据，可以是地图、结构体或待验证的某个值。 md5:e15200f8fa5aa3a2
# <翻译结束>


<原文开始>
// Associated data, which is usually a map, for union validation.
<原文结束>

# <翻译开始>
// 关联数据，通常是一个映射，用于联合验证。 md5:9888fdb467a95751
# <翻译结束>


<原文开始>
// Custom validation data.
<原文结束>

# <翻译开始>
// 自定义验证数据。 md5:35e94ac262edfe24
# <翻译结束>


<原文开始>
// Custom validation error messages, which can be string or type of CustomMsg.
<原文结束>

# <翻译开始>
// 自定义验证错误消息，可以是字符串或CustomMsg类型。 md5:c3507018b9e0da11
# <翻译结束>


<原文开始>
// ruleFuncMap stores custom rule functions for current Validator.
<原文结束>

# <翻译开始>
// ruleFuncMap 存储当前验证器的自定义规则函数。 md5:e2e248128b108117
# <翻译结束>


<原文开始>
// Using `assoc` as its validation source instead of attribute values from `Object`.
<原文结束>

# <翻译开始>
// 使用`assoc`作为验证源，而不是来自`Object`的属性值。 md5:2ecc0aebe1d9f9e0
# <翻译结束>


<原文开始>
// Stop validation after the first validation error.
<原文结束>

# <翻译开始>
// 在第一次验证错误后停止验证。 md5:78c177ee4a5553f1
# <翻译结束>


<原文开始>
// It tells the next validation using current value as an array and validates each of its element.
<原文结束>

# <翻译开始>
// 它使用当前值作为数组，并验证其每个元素，以便进行下一个验证。 md5:84d43a2805a14d90
# <翻译结束>


<原文开始>
// Case-Insensitive configuration for those rules that need value comparison.
<原文结束>

# <翻译开始>
// 大小写不敏感配置，适用于那些需要进行值比较的规则。 md5:12e4998422cd3091
# <翻译结束>


<原文开始>
// New creates and returns a new Validator.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的Validator.. md5:cca3c6d267bf0323
# <翻译结束>


<原文开始>
// Use default i18n manager.
<原文结束>

# <翻译开始>
// 使用默认的国际化管理器。 md5:89cb0f7e25a6ca81
# <翻译结束>


<原文开始>
// Custom rule function storing map.
<原文结束>

# <翻译开始>
// 自定义规则函数，用于存储映射。 md5:ac4fbe8b4302ecf3
# <翻译结束>


<原文开始>
// Run starts validating the given data with rules and messages.
<原文结束>

# <翻译开始>
// Run 开始根据规则和消息验证给定的数据。 md5:4345968979b93f1e
# <翻译结束>


<原文开始>
// Clone creates and returns a new Validator which is a shallow copy of current one.
<原文结束>

# <翻译开始>
// Clone 创建并返回一个新的Validator，它是当前对象的浅拷贝。 md5:3524ef480b75393c
# <翻译结束>


<原文开始>
// I18n sets the i18n manager for the validator.
<原文结束>

# <翻译开始>
// I18n 设置验证器的i18n管理器。 md5:aeb8eebb20995b34
# <翻译结束>


<原文开始>
// Bail sets the mark for stopping validation after the first validation error.
<原文结束>

# <翻译开始>
// Bail设置在遇到第一个验证错误后停止验证的标记。 md5:219188161ae03b77
# <翻译结束>


<原文开始>
// Foreach tells the next validation using current value as an array and validates each of its element.
// Note that this decorating rule takes effect just once for next validation rule, specially for single value validation.
<原文结束>

# <翻译开始>
// Foreach 通知下一个验证器将当前值作为数组对待，并验证它的每个元素。
// 注意，此装饰规则仅对下一个验证规则生效一次，特别适用于单值验证。
// md5:59e49ab195827b14
# <翻译结束>


<原文开始>
// Ci sets the mark for Case-Insensitive for those rules that need value comparison.
<原文结束>

# <翻译开始>
// Ci 设置标记，表示对于需要值比较的规则进行不区分大小写的处理。 md5:a248130276497a1f
# <翻译结束>


<原文开始>
// Data is a chaining operation function, which sets validation data for current operation.
<原文结束>

# <翻译开始>
// Data是一个链式操作函数，为当前操作设置验证数据。 md5:4bbfa1bb8271d34e
# <翻译结束>


<原文开始>
// Assoc is a chaining operation function, which sets associated validation data for current operation.
// The optional parameter `assoc` is usually type of map, which specifies the parameter map used in union validation.
// Calling this function with `assoc` also sets `useAssocInsteadOfObjectAttributes` true
<原文结束>

# <翻译开始>
// Assoc是一个链式操作函数，为当前操作设置关联验证数据。
// 可选参数`assoc`通常类型为map，用于指定并联合验证时使用的参数映射。
// 使用带有`assoc`调用此函数也会将`useAssocInsteadOfObjectAttributes`设置为true。
// md5:45823829185f6ad6
# <翻译结束>


<原文开始>
// Rules is a chaining operation function, which sets custom validation rules for current operation.
<原文结束>

# <翻译开始>
// Rules 是一个链接操作函数，用于为当前操作设置自定义验证规则。 md5:20d3aa2d271b3575
# <翻译结束>


<原文开始>
// Messages is a chaining operation function, which sets custom error messages for current operation.
// The parameter `messages` can be type of string/[]string/map[string]string. It supports sequence in error result
// if `rules` is type of []string.
<原文结束>

# <翻译开始>
// Messages 是一个链式操作函数，用于为当前操作设置自定义错误消息。
// 参数 `messages` 可以为 string/[]string/map[string]string 类型。如果 `rules` 类型为 []string，它支持在错误结果中按顺序显示消息。
// md5:442bfbf7d1878c37
# <翻译结束>


<原文开始>
// RuleFunc registers one custom rule function to current Validator.
<原文结束>

# <翻译开始>
// RuleFunc将一个自定义规则函数注册到当前Validator。 md5:3733cab7b3035ce3
# <翻译结束>


<原文开始>
// RuleFuncMap registers multiple custom rule functions to current Validator.
<原文结束>

# <翻译开始>
// RuleFuncMap 将多个自定义规则函数注册到当前Validator。 md5:38d8a4ac760a431a
# <翻译结束>


<原文开始>
// getCustomRuleFunc retrieves and returns the custom rule function for specified rule.
<原文结束>

# <翻译开始>
// getCustomRuleFunc 获取并返回指定规则的自定义规则函数。 md5:8a82be67553011ed
# <翻译结束>


<原文开始>
// checkRuleRequired checks and returns whether the given `rule` is required even it is nil or empty.
<原文结束>

# <翻译开始>
// checkRuleRequired 检查并返回给定的 `rule` 是否即使为 nil 或空，也是必需的。 md5:8dd4a95af0752f7f
# <翻译结束>


<原文开始>
// Default required rules.
<原文结束>

# <翻译开始>
// 默认所需的规则。 md5:7047f401aaa9d537
# <翻译结束>


<原文开始>
// All custom validation rules are required rules.
<原文结束>

# <翻译开始>
// 所有自定义验证规则都是必填规则。 md5:58545e43bcc00d45
# <翻译结束>

