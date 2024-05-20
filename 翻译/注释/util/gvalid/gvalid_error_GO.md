
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
// Error is the validation error for validation result.
<原文结束>

# <翻译开始>
// Error是验证结果的错误。. md5:333865ca9d205dfa
# <翻译结束>


<原文开始>
// validationError is the validation error for validation result.
<原文结束>

# <翻译开始>
// validationError 是验证结果的验证错误。. md5:b67f2d45170f86ce
# <翻译结束>


<原文开始>
// Rules by sequence, which is used for keeping error sequence only.
<原文结束>

# <翻译开始>
// 按顺序的规则，仅用于保持错误顺序。. md5:865d75142a03d16d
# <翻译结束>


<原文开始>
// Error map:map[field]map[rule]message
<原文结束>

# <翻译开始>
// 错误信息映射：map字段到map规则到消息. md5:57934a019c99d928
# <翻译结束>


<原文开始>
// The first error rule key(empty in default).
<原文结束>

# <翻译开始>
// 第一个错误规则键（默认为空）。. md5:19b132d9be7a2e96
# <翻译结束>


<原文开始>
// The first error rule value(nil in default).
<原文结束>

# <翻译开始>
// 第一个错误规则的值（默认为nil）。. md5:282d9086842ac373
# <翻译结束>


<原文开始>
// newValidationError creates and returns a validation error.
<原文结束>

# <翻译开始>
// newValidationError 创建并返回一个验证错误。. md5:60829ca804e6f83e
# <翻译结束>


<原文开始>
// Filter repeated sequence rules.
<原文结束>

# <翻译开始>
// 过滤重复序列规则。. md5:7a7958b11e315baa
# <翻译结束>


<原文开始>
// newValidationErrorByStr creates and returns a validation error by string.
<原文结束>

# <翻译开始>
// newValidationErrorByStr 通过字符串创建并返回一个验证错误。. md5:f8649a2f7b8f4b7c
# <翻译结束>


<原文开始>
// Code returns the error code of current validation error.
<原文结束>

# <翻译开始>
// Code 返回当前验证错误的错误代码。. md5:e3c1f143cc6ab020
# <翻译结束>


<原文开始>
// Map returns the first error message as map.
<原文结束>

# <翻译开始>
// Map 返回第一个错误消息作为映射。. md5:a50660d08282062c
# <翻译结束>


<原文开始>
// Maps returns all error messages as map.
<原文结束>

# <翻译开始>
// Maps返回所有的错误消息作为映射。. md5:3018cad54a77010b
# <翻译结束>


<原文开始>
// Items retrieves and returns error items array in sequence if possible,
// or else it returns error items with no sequence .
<原文结束>

# <翻译开始>
// Items 如果可能，按顺序检索并返回错误项数组，否则返回无序的错误项。
// md5:cb51d4d0fa07a635
# <翻译结束>


<原文开始>
// FirstItem returns the field name and error messages for the first validation rule error.
<原文结束>

# <翻译开始>
// FirstItem 返回第一个验证规则错误的字段名称和错误消息。. md5:f1a0ce09f39c751b
# <翻译结束>


<原文开始>
// FirstRule returns the first error rule and message string.
<原文结束>

# <翻译开始>
// FirstRule 返回第一个错误规则及其消息字符串。. md5:ba540411a8e82a5d
# <翻译结束>


<原文开始>
// FirstError returns the first error message as string.
// Note that the returned message might be different if it has no sequence.
<原文结束>

# <翻译开始>
// FirstError 返回第一个错误消息作为字符串。
// 注意，如果没有错误序列，返回的消息可能会有所不同。
// md5:194a5e5551fbb1e3
# <翻译结束>


<原文开始>
// Current is alis of FirstError, which implements interface gerror.iCurrent.
<原文结束>

# <翻译开始>
// Current是FirstError的别名，实现了gerror.iCurrent接口。. md5:0a09fda4e8417f2c
# <翻译结束>


<原文开始>
// String returns all error messages as string, multiple error messages joined using char ';'.
<原文结束>

# <翻译开始>
// String 返回所有错误信息作为一个字符串，多个错误消息使用分号 ';' 连接。. md5:d6ac7d8c7c8a6a03
# <翻译结束>


<原文开始>
// Error implements interface of error.Error.
<原文结束>

# <翻译开始>
// Error 实现了 error 接口的 Error 方法。. md5:6b9d58fee5a72399
# <翻译结束>


<原文开始>
// Strings returns all error messages as string array.
<原文结束>

# <翻译开始>
// Strings 将所有的错误消息返回为字符串数组。. md5:63f084a27bc91b14
# <翻译结束>


<原文开始>
// validation error checks.
<原文结束>

# <翻译开始>
// 验证错误检查。. md5:f68965da177b50ef
# <翻译结束>

