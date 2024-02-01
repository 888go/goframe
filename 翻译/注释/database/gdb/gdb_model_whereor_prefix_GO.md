
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
// See WhereBuilder.WhereOrPrefix.
<原文结束>

# <翻译开始>
// WhereOrPrefix 执行的功能与 WhereOr 相同，但会在 where 语句中的每个字段前添加指定的前缀。
// 请参阅 WhereBuilder.WhereOrPrefix。
# <翻译结束>


<原文开始>
// WhereOrPrefixLT builds `prefix.column < value` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixLT.
<原文结束>

# <翻译开始>
// WhereOrPrefixLT 用于构建 `prefix.column < value` 形式的表达式，并将其以 `OR` 条件方式组合。 
// 详情请参阅 WhereBuilder.WhereOrPrefixLT 方法。
# <翻译结束>


<原文开始>
// WhereOrPrefixLTE builds `prefix.column <= value` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixLTE.
<原文结束>

# <翻译开始>
// WhereOrPrefixLTE 用于构建 `prefix.column <= value` 形式的 OR 条件语句。
// 请参阅 WhereBuilder.WhereOrPrefixLTE 方法。
# <翻译结束>


<原文开始>
// WhereOrPrefixGT builds `prefix.column > value` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixGT.
<原文结束>

# <翻译开始>
// WhereOrPrefixGT 用于构建 `prefix.column > value` 形式的表达式，并将其以 `OR` 条件的方式加入到语句中。
// 详情请参阅 WhereBuilder.WhereOrPrefixGT。
# <翻译结束>


<原文开始>
// WhereOrPrefixGTE builds `prefix.column >= value` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixGTE.
<原文结束>

# <翻译开始>
// WhereOrPrefixGTE 用于构建 `prefix.column >= value` 形式的 OR 条件语句。
// 请参考 WhereBuilder.WhereOrPrefixGTE。
# <翻译结束>


<原文开始>
// WhereOrPrefixBetween builds `prefix.column BETWEEN min AND max` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixBetween.
<原文结束>

# <翻译开始>
// WhereOrPrefixBetween 用于构建在“OR”条件中的 `prefix.column BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereOrPrefixBetween。
# <翻译结束>


<原文开始>
// WhereOrPrefixLike builds `prefix.column LIKE like` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixLike.
<原文结束>

# <翻译开始>
// WhereOrPrefixLike 在“OR”条件下构建 `prefix.column LIKE like` 语句。
// 参见 WhereBuilder.WhereOrPrefixLike。
# <翻译结束>


<原文开始>
// WhereOrPrefixIn builds `prefix.column IN (in)` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixIn.
<原文结束>

# <翻译开始>
// WhereOrPrefixIn 在`OR`条件中构建 `prefix.column IN (in)` 语句。
// 参见 WhereBuilder.WhereOrPrefixIn。
# <翻译结束>


<原文开始>
// WhereOrPrefixNull builds `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixNull.
<原文结束>

# <翻译开始>
// WhereOrPrefixNull 用于构建在`OR`条件中的 `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` 语句。
// 参见 WhereBuilder.WhereOrPrefixNull。
# <翻译结束>


<原文开始>
// WhereOrPrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixNotBetween.
<原文结束>

# <翻译开始>
// WhereOrPrefixNotBetween 在“OR”条件下构建 `prefix.column NOT BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereOrPrefixNotBetween。
# <翻译结束>


<原文开始>
// WhereOrPrefixNotLike builds `prefix.column NOT LIKE like` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixNotLike.
<原文结束>

# <翻译开始>
// WhereOrPrefixNotLike 用于构建 `prefix.column NOT LIKE like` 语句并在 `OR` 条件中使用。
// 参见 WhereBuilder.WhereOrPrefixNotLike。
# <翻译结束>


<原文开始>
// WhereOrPrefixNotIn builds `prefix.column NOT IN (in)` statement.
// See WhereBuilder.WhereOrPrefixNotIn.
<原文结束>

# <翻译开始>
// WhereOrPrefixNotIn 用于构建 `prefix.column NOT IN (in)` 语句。
// 请参阅 WhereBuilder.WhereOrPrefixNotIn。
# <翻译结束>


<原文开始>
// WhereOrPrefixNotNull builds `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixNotNull.
<原文结束>

# <翻译开始>
// WhereOrPrefixNotNull 用于构建在`OR`条件中的 `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` 语句。
// 参见 WhereBuilder.WhereOrPrefixNotNull。
# <翻译结束>


<原文开始>
// WhereOrPrefixNot builds `prefix.column != value` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixNot.
<原文结束>

# <翻译开始>
// WhereOrPrefixNot 在`OR`条件中构建 `prefix.column != value` 语句。
// 参见 WhereBuilder.WhereOrPrefixNot。
# <翻译结束>

