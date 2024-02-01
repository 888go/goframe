
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
// WhereOrPrefix performs as WhereOr, but it adds prefix to each field in where statement.
// Eg:
// WhereOrPrefix("order", "status", "paid")                        => WHERE xxx OR (`order`.`status`='paid')
// WhereOrPrefix("order", struct{Status:"paid", "channel":"bank"}) => WHERE xxx OR (`order`.`status`='paid' AND `order`.`channel`='bank')
<原文结束>

# <翻译开始>
// WhereOrPrefix 的行为类似于 WhereOr，但是它会在 where 语句中的每个字段前添加指定的前缀。
// 示例：
// WhereOrPrefix("order", "status", "paid")                        => WHERE xxx OR (`order`.`status`='已支付')
// WhereOrPrefix("order", struct{Status:"paid", Channel:"bank"}) => WHERE xxx OR (`order`.`status`='已支付' AND `order`.`channel`='银行')
# <翻译结束>


<原文开始>
// WhereOrPrefixNot builds `prefix.column != value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixNot 在“OR”条件下构建“prefix.column != value”语句。
# <翻译结束>


<原文开始>
// WhereOrPrefixLT builds `prefix.column < value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixLT 在`OR`条件下构建 `prefix.column < value` 语句。
# <翻译结束>


<原文开始>
// WhereOrPrefixLTE builds `prefix.column <= value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixLTE 在`OR`条件下构建 `prefix.column <= value` 语句。
# <翻译结束>


<原文开始>
// WhereOrPrefixGT builds `prefix.column > value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixGT 在“OR”条件下构建 `prefix.column > value` 语句。
# <翻译结束>


<原文开始>
// WhereOrPrefixGTE builds `prefix.column >= value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixGTE 在“OR”条件下构建“prefix.column >= value”语句。
# <翻译结束>


<原文开始>
// WhereOrPrefixBetween builds `prefix.column BETWEEN min AND max` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixBetween 在`OR`条件下构建 `prefix.column BETWEEN min AND max` 语句。
# <翻译结束>


<原文开始>
// WhereOrPrefixLike builds `prefix.column LIKE 'like'` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixLike 在`OR`条件下构建`prefix.column LIKE 'like'`语句。
# <翻译结束>


<原文开始>
// WhereOrPrefixIn builds `prefix.column IN (in)` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixIn 在“OR”条件下构建 `prefix.column IN (in)` 语句。
# <翻译结束>


<原文开始>
// WhereOrPrefixNull builds `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixNull 在`OR`条件下构建 `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` 语句。
# <翻译结束>


<原文开始>
// WhereOrPrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixNotBetween 在“OR”条件下构建 `prefix.column NOT BETWEEN min AND max` 语句。
# <翻译结束>


<原文开始>
// WhereOrPrefixNotLike builds `prefix.column NOT LIKE 'like'` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixNotLike 在`OR`条件下构建 `prefix.column NOT LIKE 'like'` 语句。
# <翻译结束>


<原文开始>
// WhereOrPrefixNotIn builds `prefix.column NOT IN (in)` statement.
<原文结束>

# <翻译开始>
// WhereOrPrefixNotIn 用于构建 `prefix.column NOT IN (in)` 语句。
# <翻译结束>


<原文开始>
// WhereOrPrefixNotNull builds `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixNotNull 用于构建 `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` 形式的 OR 条件语句。
# <翻译结束>

