
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
// Value specifies the value for the rules to be validated.
<原文结束>

# <翻译开始>
// Value 指定待验证规则的值。
# <翻译结束>


<原文开始>
// ValueType specifies the type of the value, mainly used for value type id retrieving.
<原文结束>

# <翻译开始>
// ValueType 指定值的类型，主要用于获取值的类型标识。
# <翻译结束>


<原文开始>
// Messages specifies the custom error messages for this rule from parameters input, which is usually type of map/slice.
<原文结束>

# <翻译开始>
// Messages 指定了该规则从输入参数（通常为 map 或 slice 类型）获取的自定义错误消息。
# <翻译结束>


<原文开始>
// doCheckValue does the really rules validation for single key-value.
<原文结束>

# <翻译开始>
// doCheckValue 对单个键值执行真正的规则验证。
# <翻译结束>


<原文开始>
// If there's no validation rules, it does nothing and returns quickly.
<原文结束>

# <翻译开始>
// 如果没有验证规则，它将不做任何操作并迅速返回。
# <翻译结束>


<原文开始>
// It converts value to string and then does the validation.
<原文结束>

# <翻译开始>
// 它将值转换为字符串，然后进行验证。
# <翻译结束>


<原文开始>
// Do not trim it as the space is also part of the value.
<原文结束>

# <翻译开始>
// 不要进行修剪操作，因为空格也是值的一部分。
# <翻译结束>


<原文开始>
// Custom error messages handling.
<原文结束>

# <翻译开始>
// 自定义错误消息处理。
# <翻译结束>


<原文开始>
	// Handle the char '|' in the rule,
	// which makes this rule separated into multiple rules.
<原文结束>

# <翻译开始>
// 处理规则中的字符' | '，
// 这使得该规则被分割为多个规则。
# <翻译结束>


<原文开始>
			// ============================ SPECIAL ============================
			// Special `regex` and `not-regex` rules.
			// Merge the regex pattern if there are special chars, like ':', '|', in pattern.
			// ============================ SPECIAL ============================
<原文结束>

# <翻译开始>
// ================================== 特殊规则 ==================================
// 特殊的 `regex` 和 `not-regex` 规则。
// 如果模式中包含特殊字符（如 ':'、'|' 等），则合并正则表达式模式。
// ================================== 特殊规则 ==================================
# <翻译结束>







<原文开始>
// Ignore logic executing for marked rules.
<原文结束>

# <翻译开始>
// 忽略已标记规则的执行逻辑。
# <翻译结束>


<原文开始>
// As it marks `foreach`, so it converts the value to slice.
<原文结束>

# <翻译开始>
// 因为此处标记了 `foreach`，所以它会将值转换为切片。
# <翻译结束>


<原文开始>
// Reset `foreach` rule as it only takes effect just once for next rule.
<原文结束>

# <翻译开始>
// 重置`foreach`规则，因为它只为下一条规则生效一次。
# <翻译结束>












<原文开始>
// It never comes across here.
<原文结束>

# <翻译开始>
// 这里永远不会执行到。
# <翻译结束>







<原文开始>
// Error variable replacement for error message.
<原文结束>

# <翻译开始>
// 错误变量替换用于错误消息。
# <翻译结束>







<原文开始>
// The variable part of the rule.
<原文结束>

# <翻译开始>
// 规则的可变部分。
# <翻译结束>


<原文开始>
// The error should have stack info to indicate the error position.
<原文结束>

# <翻译开始>
// 该错误应包含堆栈信息以指示错误位置。
# <翻译结束>


<原文开始>
// The error should have error code that is `gcode.CodeValidationFailed`.
<原文结束>

# <翻译开始>
// 错误应具有错误代码 `gcode.CodeValidationFailed`。
# <翻译结束>


<原文开始>
				// If it is with error and there's bail rule,
				// it then does not continue validating for left rules.
<原文结束>

# <翻译开始>
// 如果遇到错误且存在中断规则，
// 则不再继续验证剩余规则。
# <翻译结束>







<原文开始>
// Struct/map/slice type which to be recursively validated.
<原文结束>

# <翻译开始>
// 需要递归验证的结构体/映射/切片类型。
# <翻译结束>


<原文开始>
// Struct/map/slice kind to be asserted in following switch case.
<原文结束>

# <翻译开始>
// 在接下来的switch case中，需要断言的结构体/映射/切片类型。
# <翻译结束>


<原文开始>
// The validated failed error map.
<原文结束>

# <翻译开始>
// 验证失败的错误映射。
# <翻译结束>


<原文开始>
// The validated failed rule in sequence.
<原文结束>

# <翻译开始>
// 验证失败的规则按顺序排列。
# <翻译结束>


<原文开始>
// Ignore data, assoc, rules and messages from parent.
<原文结束>

# <翻译开始>
// 忽略来自父级的数据、关联、规则和消息。
# <翻译结束>


<原文开始>
// It merges the errors into single error map.
<原文结束>

# <翻译开始>
// 它将错误合并成单个错误映射。
# <翻译结束>


<原文开始>
// Name specifies the name of parameter `value`.
<原文结束>

# <翻译开始>
// Name 指定参数 `value` 的名称。
# <翻译结束>


<原文开始>
// Rule specifies the validation rules string, like "required", "required|between:1,100", etc.
<原文结束>

# <翻译开始>
// Rule 指定验证规则字符串，如 "required", "required|between:1,100" 等。
# <翻译结束>


<原文开始>
// DataRaw specifies the `raw data` which is passed to the Validator. It might be type of map/struct or a nil value.
<原文结束>

# <翻译开始>
// DataRaw 指定传递给验证器的 `原始数据`，其类型可以是 map 或 struct，也可以是 nil 值。
# <翻译结束>


<原文开始>
// DataMap specifies the map that is converted from `dataRaw`. It is usually used internally
<原文结束>

# <翻译开始>
// DataMap 指定了从 `dataRaw` 转换而来的映射（map）。它通常用于内部实现
// ```go
// DataMap 代表由 `dataRaw` 转化而来的数据映射，主要用于内部使用
# <翻译结束>


<原文开始>
// rule key like "max" in rule "max: 6"
<原文结束>

# <翻译开始>
// rule key 类似于规则 "max: 6" 中的 "max"
# <翻译结束>


<原文开始>
// rule pattern is like "6" in rule:"max:6"
<原文结束>

# <翻译开始>
// rule pattern 是规则中的模式部分，例如在规则 "max:6" 中的 "6"
# <翻译结束>







<原文开始>
// The same as `{field}`. It is deprecated.
<原文结束>

# <翻译开始>
// 与 `{field}` 相同。已被弃用。
# <翻译结束>


<原文开始>
// split single rule.
<原文结束>

# <翻译开始>
// 分割单个规则。
# <翻译结束>


<原文开始>
// Custom validation rules.
<原文结束>

# <翻译开始>
// 自定义验证规则。
# <翻译结束>


<原文开始>
// Builtin validation rules.
<原文结束>

# <翻译开始>
// 内置验证规则。
# <翻译结束>


<原文开始>
// Field name of the `value`.
<原文结束>

# <翻译开始>
// `value`的字段名称。
# <翻译结束>


<原文开始>
// Current validating value.
<原文结束>

# <翻译开始>
// 当前验证中的值。
# <翻译结束>


<原文开始>
// Value to be validated.
<原文结束>

# <翻译开始>
// 需要验证的值。
# <翻译结束>

