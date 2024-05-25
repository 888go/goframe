
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
// Name specifies the name of parameter `value`, which might be the custom tag name of the parameter.
<原文结束>

# <翻译开始>
// Name 指定了参数 `value` 的名称，它可能是参数的自定义标签名。 md5:31bb3221a3724b81
# <翻译结束>


<原文开始>
// Value specifies the value for the rules to be validated.
<原文结束>

# <翻译开始>
// Value 指定要验证的规则的值。 md5:7ee2f1438ebf6afb
# <翻译结束>


<原文开始>
// ValueType specifies the type of the value, mainly used for value type id retrieving.
<原文结束>

# <翻译开始>
// ValueType 指定了值的类型，主要用于获取值类型的ID。 md5:299dde27f54dddba
# <翻译结束>


<原文开始>
// Rule specifies the validation rules string, like "required", "required|between:1,100", etc.
<原文结束>

# <翻译开始>
// Rule 指定验证规则字符串，如 "required", "required|between:1,100" 等。 md5:2ae63fb98a26734b
# <翻译结束>


<原文开始>
// Messages specifies the custom error messages for this rule from parameters input, which is usually type of map/slice.
<原文结束>

# <翻译开始>
// Messages从参数输入（通常为map或slice类型）中指定该规则的自定义错误消息。 md5:6aa72de28be7f730
# <翻译结束>


<原文开始>
// DataRaw specifies the `raw data` which is passed to the Validator. It might be type of map/struct or a nil value.
<原文结束>

# <翻译开始>
// DataRaw 指定传递给Validator的`原始数据`。它可能是map/结构体类型或nil值。 md5:ea0ae3bdb176793b
# <翻译结束>


<原文开始>
// DataMap specifies the map that is converted from `dataRaw`. It is usually used internally
<原文结束>

# <翻译开始>
// DataMap 指定了从 `dataRaw` 转换而来的映射。它通常在内部使用. md5:50285f5b09df4771
# <翻译结束>


<原文开始>
// doCheckValue does the really rules validation for single key-value.
<原文结束>

# <翻译开始>
// doCheckValue 对单个键值对执行实际的规则验证。 md5:9032f66341668b1c
# <翻译结束>


<原文开始>
// If there's no validation rules, it does nothing and returns quickly.
<原文结束>

# <翻译开始>
// 如果没有验证规则，它什么也不做并迅速返回。 md5:bc52d29571b990f7
# <翻译结束>


<原文开始>
// It converts value to string and then does the validation.
<原文结束>

# <翻译开始>
// 它将值转换为字符串，然后进行验证。 md5:2687e35bf141700c
# <翻译结束>


<原文开始>
// Do not trim it as the space is also part of the value.
<原文结束>

# <翻译开始>
// 不要删除空白，因为空白也是值的一部分。 md5:149754fbc3e60837
# <翻译结束>


<原文开始>
// Custom error messages handling.
<原文结束>

# <翻译开始>
// 自定义错误消息处理。 md5:034ef969034ce61c
# <翻译结束>


<原文开始>
	// Handle the char '|' in the rule,
	// which makes this rule separated into multiple rules.
<原文结束>

# <翻译开始>
// 处理规则中的字符'|'，这使得该规则被分成多个子规则。
// md5:11aa0a7f39f13bef
# <翻译结束>


<原文开始>
			// ============================ SPECIAL ============================
			// Special `regex` and `not-regex` rules.
			// Merge the regex pattern if there are special chars, like ':', '|', in pattern.
			// ============================ SPECIAL ============================
<原文结束>

# <翻译开始>
// =========================== 特殊 ===========================
// 对于特殊的正则表达式 (`regex`) 和非正则表达式 (`not-regex`) 规则。
// 如果模式中包含特殊字符，如 ':' 或 '|'，则合并正则表达式模式。
// =========================== 特殊 ===========================
// md5:8f3bcac9a314de33
# <翻译结束>


<原文开始>
// rule key like "max" in rule "max: 6"
<原文结束>

# <翻译开始>
// 规则键，如规则 "max: 6" 中的 "max". md5:b9eff8d7691a084c
# <翻译结束>


<原文开始>
// rule pattern is like "6" in rule:"max:6"
<原文结束>

# <翻译开始>
// 规则模式类似于 "6" 在规则 "max:6" 中. md5:7766c1e829f5f940
# <翻译结束>


<原文开始>
// Ignore logic executing for marked rules.
<原文结束>

# <翻译开始>
// 忽略标记规则的执行逻辑。 md5:34f3e7a7cffba70b
# <翻译结束>


<原文开始>
// As it marks `foreach`, so it converts the value to slice.
<原文结束>

# <翻译开始>
// 由于它标记了 `foreach`，所以它会将值转换为切片。 md5:9f599bb9b2fe0bba
# <翻译结束>


<原文开始>
// Reset `foreach` rule as it only takes effect just once for next rule.
<原文结束>

# <翻译开始>
// 重置 `foreach` 规则，因为它只对下一条规则生效一次。 md5:8c7dd94030559037
# <翻译结束>


<原文开始>
// Custom validation rules.
<原文结束>

# <翻译开始>
// 自定义验证规则。 md5:fbd7800af1a73578
# <翻译结束>


<原文开始>
// Builtin validation rules.
<原文结束>

# <翻译开始>
// 内置验证规则。 md5:4f4f87cac993a840
# <翻译结束>


<原文开始>
// It never comes across here.
<原文结束>

# <翻译开始>
// 它永远不会出现在这里。 md5:1b17e9ac7d650245
# <翻译结束>


<原文开始>
// Error variable replacement for error message.
<原文结束>

# <翻译开始>
// 用于错误信息的错误变量替换。 md5:c424d98305e44662
# <翻译结束>


<原文开始>
// Field name of the `value`.
<原文结束>

# <翻译开始>
// `value` 的字段名称。 md5:c75900d2041a10e5
# <翻译结束>


<原文开始>
// Current validating value.
<原文结束>

# <翻译开始>
// 当前验证的值。 md5:17abd56cedea072f
# <翻译结束>


<原文开始>
// The variable part of the rule.
<原文结束>

# <翻译开始>
// 规则的变量部分。 md5:1463434d04a94902
# <翻译结束>


<原文开始>
// The same as `{field}`. It is deprecated.
<原文结束>

# <翻译开始>
// 与 `{field}` 相同。此用法已废弃。 md5:0ceaca304a2589af
# <翻译结束>


<原文开始>
// The error should have stack info to indicate the error position.
<原文结束>

# <翻译开始>
// 错误应该包含堆栈信息，以指示错误的位置。 md5:bef4a94931ed384c
# <翻译结束>


<原文开始>
// The error should have error code that is `gcode.CodeValidationFailed`.
<原文结束>

# <翻译开始>
// 错误应该有错误代码，该代码为 `gcode.CodeValidationFailed`。 md5:b54af62f83c4db11
# <翻译结束>


<原文开始>
// TODO it's better using interface?
<原文结束>

# <翻译开始>
// TODO 使用接口可能更好？. md5:04cb382580755c3a
# <翻译结束>


<原文开始>
				// If it is with error and there's bail rule,
				// it then does not continue validating for left rules.
<原文结束>

# <翻译开始>
// 如果存在错误并且有放弃规则，
// 则不再继续验证剩余的规则。
// md5:746db6c03bb62206
# <翻译结束>


<原文开始>
// Struct/map/slice type which to be recursively validated.
<原文结束>

# <翻译开始>
// 将要递归验证的结构体/映射/切片类型。 md5:ae6984d7ba567001
# <翻译结束>


<原文开始>
// Struct/map/slice kind to be asserted in following switch case.
<原文结束>

# <翻译开始>
// 要在接下来的开关语句中进行断言的结构体/映射/切片类型。 md5:b683235f95d7aae1
# <翻译结束>


<原文开始>
// The validated failed error map.
<原文结束>

# <翻译开始>
// 验证失败的错误映射。 md5:e0888bbfb505d641
# <翻译结束>


<原文开始>
// The validated failed rule in sequence.
<原文结束>

# <翻译开始>
// 依次验证失败的规则。 md5:5e8c03560ecc4a22
# <翻译结束>


<原文开始>
// Ignore data, assoc, rules and messages from parent.
<原文结束>

# <翻译开始>
// 忽略父级的数据、关联、规则和消息。 md5:27ad0097eee0432e
# <翻译结束>


<原文开始>
// It merges the errors into single error map.
<原文结束>

# <翻译开始>
// 它将错误合并为单个错误映射。 md5:56fe32c627a507ee
# <翻译结束>

