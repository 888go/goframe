
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
// WherePrefix performs as Where, but it adds prefix to each field in where statement.
// Eg:
// WherePrefix("order", "status", "paid")                        => WHERE `order`.`status`='paid'
// WherePrefix("order", struct{Status:"paid", "channel":"bank"}) => WHERE `order`.`status`='paid' AND `order`.`channel`='bank'
<原文结束>

# <翻译开始>
// WherePrefix 的行为类似于 Where，但它会在 where 语句中的每个字段前添加前缀。
// 示例：
// WherePrefix("order", "status", "paid")                        => WHERE `order`.`status`='paid'
// WherePrefix("order", struct{Status:"paid", Channel:"bank"}) => WHERE `order`.`status`='paid' AND `order`.`channel`='bank'
# <翻译结束>


<原文开始>
// WherePrefixLT builds `prefix.column < value` statement.
<原文结束>

# <翻译开始>
// WherePrefixLT 用于构建 `prefix.column < value` 的语句。
# <翻译结束>


<原文开始>
// WherePrefixLTE builds `prefix.column <= value` statement.
<原文结束>

# <翻译开始>
// WherePrefixLTE 用于构建 `prefix.column <= value` 的语句。
# <翻译结束>


<原文开始>
// WherePrefixGT builds `prefix.column > value` statement.
<原文结束>

# <翻译开始>
// WherePrefixGT 用于构建 `prefix.column > value` 的表达式语句。
# <翻译结束>


<原文开始>
// WherePrefixGTE builds `prefix.column >= value` statement.
<原文结束>

# <翻译开始>
// WherePrefixGTE 生成 `prefix.column >= value` 语句。
# <翻译结束>


<原文开始>
// WherePrefixBetween builds `prefix.column BETWEEN min AND max` statement.
<原文结束>

# <翻译开始>
// WherePrefixBetween 用于构建 `prefix.column BETWEEN min AND max` 的语句。
# <翻译结束>


<原文开始>
// WherePrefixLike builds `prefix.column LIKE like` statement.
<原文结束>

# <翻译开始>
// WherePrefixLike 用于构建 `prefix.column LIKE like` 语句。
# <翻译结束>


<原文开始>
// WherePrefixIn builds `prefix.column IN (in)` statement.
<原文结束>

# <翻译开始>
// WherePrefixIn 用于构建 `prefix.column IN (in)` 语句。
# <翻译结束>


<原文开始>
// WherePrefixNull builds `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` statement.
<原文结束>

# <翻译开始>
// WherePrefixNull 用于构建如 `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` 形式的语句。
# <翻译结束>


<原文开始>
// WherePrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement.
<原文结束>

# <翻译开始>
// WherePrefixNotBetween 用于构建 `prefix.column NOT BETWEEN min AND max` 的表达式语句。
# <翻译结束>


<原文开始>
// WherePrefixNotLike builds `prefix.column NOT LIKE like` statement.
<原文结束>

# <翻译开始>
// WherePrefixNotLike 用于构建 `prefix.column NOT LIKE like` 语句。
# <翻译结束>


<原文开始>
// WherePrefixNot builds `prefix.column != value` statement.
<原文结束>

# <翻译开始>
// WherePrefixNot 用于构建 `prefix.column != value` 的表达式语句。
# <翻译结束>


<原文开始>
// WherePrefixNotIn builds `prefix.column NOT IN (in)` statement.
<原文结束>

# <翻译开始>
// WherePrefixNotIn 用于构建 `prefix.column NOT IN (in)` 语句。
# <翻译结束>


<原文开始>
// WherePrefixNotNull builds `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` statement.
<原文结束>

# <翻译开始>
// WherePrefixNotNull 用于构建 `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` 语句。
# <翻译结束>

