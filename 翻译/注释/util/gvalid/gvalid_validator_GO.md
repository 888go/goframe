
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
// Validator is the validation manager for chaining operations.
<原文结束>

# <翻译开始>
// Validator 是用于链接操作的验证管理器。
# <翻译结束>


<原文开始>
// I18n manager for error message translation.
<原文结束>

# <翻译开始>
// I18n 错误信息翻译管理器。
# <翻译结束>


<原文开始>
// Validation data, which can be a map, struct or a certain value to be validated.
<原文结束>

# <翻译开始>
// 验证数据，它可以是一个映射、结构体或需要进行验证的特定值。
# <翻译结束>


<原文开始>
// Associated data, which is usually a map, for union validation.
<原文结束>

# <翻译开始>
// 关联数据，通常是一个映射，用于联合验证。
# <翻译结束>







<原文开始>
// Custom validation error messages, which can be string or type of CustomMsg.
<原文结束>

# <翻译开始>
// 自定义验证错误消息，可以是字符串或CustomMsg类型的。
# <翻译结束>


<原文开始>
// ruleFuncMap stores custom rule functions for current Validator.
<原文结束>

# <翻译开始>
// ruleFuncMap 用于存储当前 Validator 的自定义规则函数。
# <翻译结束>


<原文开始>
// Stop validation after the first validation error.
<原文结束>

# <翻译开始>
// 在出现第一个验证错误后停止验证。
# <翻译结束>


<原文开始>
// It tells the next validation using current value as an array and validates each of its element.
<原文结束>

# <翻译开始>
// 它使用当前值作为数组并对其中的每个元素进行验证，用于告知接下来的校验。
# <翻译结束>


<原文开始>
// Case-Insensitive configuration for those rules that need value comparison.
<原文结束>

# <翻译开始>
// 对于那些需要值比较的规则，提供不区分大小写的配置。
# <翻译结束>


<原文开始>
// New creates and returns a new Validator.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的验证器。
# <翻译结束>







<原文开始>
// Custom rule function storing map.
<原文结束>

# <翻译开始>
// 自定义规则函数存储映射。
# <翻译结束>


<原文开始>
// Run starts validating the given data with rules and messages.
<原文结束>

# <翻译开始>
// Run 开始使用规则和消息验证给定的数据。
# <翻译结束>


<原文开始>
// Clone creates and returns a new Validator which is a shallow copy of current one.
<原文结束>

# <翻译开始>
// Clone 创建并返回一个新的 Validator，它是当前 Validator 的浅复制。
# <翻译结束>


<原文开始>
// I18n sets the i18n manager for the validator.
<原文结束>

# <翻译开始>
// I18n 为验证器设置国际化管理器。
# <翻译结束>


<原文开始>
// Bail sets the mark for stopping validation after the first validation error.
<原文结束>

# <翻译开始>
// Bail 设置标记，在出现第一个验证错误后停止验证。
# <翻译结束>


<原文开始>
// Foreach tells the next validation using current value as an array and validates each of its element.
// Note that this decorating rule takes effect just once for next validation rule, specially for single value validation.
<原文结束>

# <翻译开始>
// Foreach 注释：对于当前值作为数组进行处理，并对其每个元素进行验证。
// 注意，此装饰规则仅对下一次验证规则生效一次，特别是针对单个值的验证。
# <翻译结束>


<原文开始>
// Ci sets the mark for Case-Insensitive for those rules that need value comparison.
<原文结束>

# <翻译开始>
// Ci 设置标记，用于那些需要值比较的规则，实现大小写不敏感。
# <翻译结束>


<原文开始>
// Data is a chaining operation function, which sets validation data for current operation.
<原文结束>

# <翻译开始>
// Data 是一个链式操作函数，用于为当前操作设置验证数据。
# <翻译结束>


<原文开始>
// Assoc is a chaining operation function, which sets associated validation data for current operation.
// The optional parameter `assoc` is usually type of map, which specifies the parameter map used in union validation.
// Calling this function with `assoc` also sets `useAssocInsteadOfObjectAttributes` true
<原文结束>

# <翻译开始>
// Assoc 是一个链式操作函数，用于为当前操作设置关联验证数据。
// 可选参数 `assoc` 通常为 map 类型，用于指定联合验证中使用的参数映射。
// 当调用该函数并传入 `assoc` 参数时，同时会将 `useAssocInsteadOfObjectAttributes` 设置为 true。
# <翻译结束>


<原文开始>
// Rules is a chaining operation function, which sets custom validation rules for current operation.
<原文结束>

# <翻译开始>
// Rules 是一个链式操作函数，用于为当前操作设置自定义验证规则。
# <翻译结束>


<原文开始>
// Messages is a chaining operation function, which sets custom error messages for current operation.
// The parameter `messages` can be type of string/[]string/map[string]string. It supports sequence in error result
// if `rules` is type of []string.
<原文结束>

# <翻译开始>
// Messages 是一个链式操作函数，用于为当前操作设置自定义错误消息。
// 参数 `messages` 可以是 string/[]string/map[string]string 类型。如果 `rules` 类型为 []string，则支持在错误结果中按顺序展示消息。
// 更详细的翻译：
// ```go
// Messages 函数提供链式操作功能，允许为当前执行的操作设定个性化的错误信息。
// 其参数 `messages` 的类型可以是字符串(string)、字符串切片([]string)或字符串到字符串的映射(map[string]string)。
// 特别地，当 `rules` 参数为字符串切片([]string)类型时，Messages 函数能够支持按照规则数组中的顺序，在生成的错误结果中逐条展示相应的错误消息。
# <翻译结束>


<原文开始>
// RuleFunc registers one custom rule function to current Validator.
<原文结束>

# <翻译开始>
// RuleFunc 将一个自定义规则函数注册到当前的 Validator。
# <翻译结束>


<原文开始>
// RuleFuncMap registers multiple custom rule functions to current Validator.
<原文结束>

# <翻译开始>
// RuleFuncMap 将多个自定义规则函数注册到当前的 Validator。
# <翻译结束>


<原文开始>
// getCustomRuleFunc retrieves and returns the custom rule function for specified rule.
<原文结束>

# <翻译开始>
// getCustomRuleFunc 根据指定规则获取并返回自定义规则函数。
# <翻译结束>


<原文开始>
// checkRuleRequired checks and returns whether the given `rule` is required even it is nil or empty.
<原文结束>

# <翻译开始>
// checkRuleRequired 检查并返回给定的 `rule` 是否为必需，即使它是 nil 或为空。
# <翻译结束>







<原文开始>
// All custom validation rules are required rules.
<原文结束>

# <翻译开始>
// 所有自定义验证规则均为必填规则。
# <翻译结束>


<原文开始>
// Using `assoc` as its validation source instead of attribute values from `Object`.
<原文结束>

# <翻译开始>
// 使用 `assoc` 作为其验证源，而不是从 `Object` 中获取属性值。
# <翻译结束>


<原文开始>
// Custom validation data.
<原文结束>

# <翻译开始>
// 自定义验证数据。
# <翻译结束>


<原文开始>
// Use default i18n manager.
<原文结束>

# <翻译开始>
// 使用默认的国际化管理器。
# <翻译结束>


<原文开始>
// Default required rules.
<原文结束>

# <翻译开始>
// 默认必需的规则。
# <翻译结束>

