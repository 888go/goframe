
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
// Package builtin implements built-in validation rules.
//
// Referred to Laravel validation:
// https://laravel.com/docs/master/validation#available-validation-rules
<原文结束>

# <翻译开始>
// Package builtin 实现了内置的验证规则。
//
// 参考了 Laravel 验证：
// https://laravel.com/docs/master/validation#available-validation-rules
// （该链接为 Laravel 框架关于可用验证规则的文档）
# <翻译结束>


<原文开始>
// Name returns the builtin name of the rule.
<原文结束>

# <翻译开始>
// Name 返回规则的内置名称。
# <翻译结束>


<原文开始>
// Message returns the default error message of the rule.
<原文结束>

# <翻译开始>
// Message 返回该规则的默认错误消息。
# <翻译结束>


<原文开始>
// Run starts running the rule, it returns nil if successful, or else an error.
<原文结束>

# <翻译开始>
// Run 开始运行规则，如果运行成功则返回 nil，否则返回错误。
# <翻译结束>







<原文开始>
// ValueType specifies the type of the value, which might be nil.
<原文结束>

# <翻译开始>
// ValueType 指定了值的类型，该值可能为 nil。
# <翻译结束>


<原文开始>
// Value specifies the value for this rule to validate.
<原文结束>

# <翻译开始>
// Value 指定此规则验证的值。
# <翻译结束>


<原文开始>
// Message specifies the custom error message or configured i18n message for this rule.
<原文结束>

# <翻译开始>
// Message 指定了该规则的自定义错误消息或配置好的国际化（i18n）消息。
# <翻译结束>


<原文开始>
// Option provides extra configuration for validation rule.
<原文结束>

# <翻译开始>
// Option 提供了验证规则的额外配置选项。
# <翻译结束>


<原文开始>
// CaseInsensitive indicates that it does Case-Insensitive comparison in string.
<原文结束>

# <翻译开始>
// CaseInsensitive 表示在进行字符串比较时采用不区分大小写的方式。
# <翻译结束>


<原文开始>
// ruleMap stores all builtin validation rules.
<原文结束>

# <翻译开始>
// ruleMap 存储所有内置验证规则。
# <翻译结束>


<原文开始>
// Register registers builtin rule into manager.
<原文结束>

# <翻译开始>
// Register 将内建规则注册到管理器中。
# <翻译结束>


<原文开始>
// GetRule retrieves and returns rule by `name`.
<原文结束>

# <翻译开始>
// GetRule 通过`name`检索并返回规则。
# <翻译结束>


<原文开始>
// RuleKey is like the "max" in rule "max: 6"
<原文结束>

# <翻译开始>
// RuleKey 类似于规则中的 "max"，如 "max: 6" 中的 "max"
# <翻译结束>


<原文开始>
// RulePattern is like "6" in rule:"max:6"
<原文结束>

# <翻译开始>
// RulePattern 类似于规则 "max:6" 中的 "6"
# <翻译结束>


<原文开始>
// Data specifies the `data` which is passed to the Validator.
<原文结束>

# <翻译开始>
// Data 指定传递给 Validator 的 `data`。
# <翻译结束>


<原文开始>
// The field name of Value.
<原文结束>

# <翻译开始>
// Value的字段名称。
# <翻译结束>

