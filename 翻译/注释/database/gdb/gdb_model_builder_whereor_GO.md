
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
// WhereOr adds "OR" condition to the where statement.
<原文结束>

# <翻译开始>
// WhereOr 向 where 语句添加“OR”条件。
# <翻译结束>


<原文开始>
// WhereOrf builds `OR` condition string using fmt.Sprintf and arguments.
<原文结束>

# <翻译开始>
// WhereOrf 使用 fmt.Sprintf 和参数构建 `OR` 条件字符串。
# <翻译结束>


<原文开始>
// WhereOrf builds `OR` condition string using fmt.Sprintf and arguments.
// Eg:
// WhereOrf(`amount<? and status=%s`, "paid", 100)  => WHERE xxx OR `amount`<100 and status='paid'
// WhereOrf(`amount<%d and status=%s`, 100, "paid") => WHERE xxx OR `amount`<100 and status='paid'
<原文结束>

# <翻译开始>
// WhereOrf 使用 fmt.Sprintf 和参数构建 `OR` 条件字符串。
// 示例：
// WhereOrf(`amount<? and status=%s`, "paid", 100)  => WHERE xxx OR `amount`<100 AND 状态='已支付'
// WhereOrf(`amount<%d and status=%s`, 100, "paid") => WHERE xxx OR `amount`<100 AND 状态='已支付'
// 注：这里的 "xxx" 是预设的其他条件部分，实际使用中会被相应的内容替换。
# <翻译结束>


<原文开始>
// WhereOrNot builds `column != value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrNot 用于构建在“OR”条件中的“column != value”语句。
# <翻译结束>


<原文开始>
// WhereOrLT builds `column < value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrLT 在“OR”条件中构建“column < value”语句。
# <翻译结束>


<原文开始>
// WhereOrLTE builds `column <= value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrLTE 用于构建在“OR”条件中的`column <= value`语句。
# <翻译结束>


<原文开始>
// WhereOrGT builds `column > value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrGT 在“OR”条件下构建“column > value”语句。
# <翻译结束>


<原文开始>
// WhereOrGTE builds `column >= value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrGTE 在“OR”条件下构建“column >= value”语句。
# <翻译结束>


<原文开始>
// WhereOrBetween builds `column BETWEEN min AND max` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrBetween 在`OR`条件下构建 `column BETWEEN min AND max` 语句。
# <翻译结束>


<原文开始>
// WhereOrLike builds `column LIKE 'like'` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrLike 在`OR`条件下构建`column LIKE 'like'`语句。
# <翻译结束>


<原文开始>
// WhereOrIn builds `column IN (in)` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrIn 在“OR”条件中构建“column IN (in)”语句。
# <翻译结束>


<原文开始>
// WhereOrNull builds `columns[0] IS NULL OR columns[1] IS NULL ...` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrNull 根据“或”条件构建 `columns[0] IS NULL OR columns[1] IS NULL ...` 语句。
# <翻译结束>


<原文开始>
// WhereOrNotBetween builds `column NOT BETWEEN min AND max` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrNotBetween 在“OR”条件中构建`column NOT BETWEEN min AND max`语句。
# <翻译结束>


<原文开始>
// WhereOrNotLike builds `column NOT LIKE like` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrNotLike 在“OR”条件下构建`column NOT LIKE like`语句。
# <翻译结束>


<原文开始>
// WhereOrNotIn builds `column NOT IN (in)` statement.
<原文结束>

# <翻译开始>
// WhereOrNotIn 用于构建 `column NOT IN (in)` 语句。
# <翻译结束>


<原文开始>
// WhereOrNotNull builds `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrNotNull 在`OR`条件下构建 `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` 语句。
# <翻译结束>

