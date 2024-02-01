
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
// Error is the validation error for validation result.
<原文结束>

# <翻译开始>
// Error 是验证结果的错误信息。
# <翻译结束>


<原文开始>
// validationError is the validation error for validation result.
<原文结束>

# <翻译开始>
// validationError 是验证结果的验证错误。
# <翻译结束>


<原文开始>
// Rules by sequence, which is used for keeping error sequence only.
<原文结束>

# <翻译开始>
// 按照顺序排列的规则，仅用于保留错误序列。
# <翻译结束>


<原文开始>
// Error map:map[field]map[rule]message
<原文结束>

# <翻译开始>
// 错误映射：map[字段]map[规则]消息
# <翻译结束>


<原文开始>
// The first error rule key(empty in default).
<原文结束>

# <翻译开始>
// 第一条错误规则键（默认为空）
# <翻译结束>


<原文开始>
// The first error rule value(nil in default).
<原文结束>

# <翻译开始>
// 第一条错误规则值（默认为nil）
# <翻译结束>


<原文开始>
// newValidationError creates and returns a validation error.
<原文结束>

# <翻译开始>
// newValidationError 创建并返回一个验证错误。
# <翻译结束>


<原文开始>
// Filter repeated sequence rules.
<原文结束>

# <翻译开始>
// 过滤重复的序列规则。
# <翻译结束>







<原文开始>
// newValidationErrorByStr creates and returns a validation error by string.
<原文结束>

# <翻译开始>
// newValidationErrorByStr 通过字符串创建并返回一个验证错误。
# <翻译结束>


<原文开始>
// Code returns the error code of current validation error.
<原文结束>

# <翻译开始>
// Code 返回当前验证错误的错误代码。
# <翻译结束>


<原文开始>
// Map returns the first error message as map.
<原文结束>

# <翻译开始>
// Map 返回第一个错误消息作为映射（map）。
# <翻译结束>


<原文开始>
// Maps returns all error messages as map.
<原文结束>

# <翻译开始>
// Maps 将所有错误消息以映射形式返回。
# <翻译结束>


<原文开始>
// Items retrieves and returns error items array in sequence if possible,
// or else it returns error items with no sequence .
<原文结束>

# <翻译开始>
// Items 函数尝试按顺序检索并返回错误项数组，如果无法按顺序获取，
// 则返回无特定顺序的错误项数组。
# <翻译结束>


<原文开始>
// FirstItem returns the field name and error messages for the first validation rule error.
<原文结束>

# <翻译开始>
// FirstItem 返回第一个验证规则错误的字段名称和错误消息。
# <翻译结束>


<原文开始>
// FirstRule returns the first error rule and message string.
<原文结束>

# <翻译开始>
// FirstRule 返回第一个错误规则及其消息字符串。
# <翻译结束>


<原文开始>
// FirstError returns the first error message as string.
// Note that the returned message might be different if it has no sequence.
<原文结束>

# <翻译开始>
// FirstError 返回第一个错误信息作为字符串。
// 注意，如果没有顺序，返回的消息可能会不同。
# <翻译结束>


<原文开始>
// Current is alis of FirstError, which implements interface gerror.iCurrent.
<原文结束>

# <翻译开始>
// Current 是 FirstError 的别名，实现了 gerror.iCurrent 接口。
# <翻译结束>


<原文开始>
// String returns all error messages as string, multiple error messages joined using char ';'.
<原文结束>

# <翻译开始>
// String 将所有错误消息作为字符串返回，多个错误消息之间使用字符 ';' 连接。
# <翻译结束>


<原文开始>
// Error implements interface of error.Error.
<原文结束>

# <翻译开始>
// Error 实现了 error 接口的 Error 方法。
# <翻译结束>


<原文开始>
// Strings returns all error messages as string array.
<原文结束>

# <翻译开始>
// Strings 将所有错误消息作为字符串数组返回。
# <翻译结束>












<原文开始>
// Delete repeated rule.
<原文结束>

# <翻译开始>
// 删除重复的规则。
# <翻译结束>


<原文开始>
// validation error checks.
<原文结束>

# <翻译开始>
// 验证错误检查。
# <翻译结束>


<原文开始>
// internal error checks.
<原文结束>

# <翻译开始>
// 内部错误检查
# <翻译结束>

