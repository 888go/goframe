
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
// WhereOr adds "OR" condition to the where statement.
<原文结束>

# <翻译开始>
// WhereOr 向 WHERE 语句中添加“OR”条件。. md5:753c32f428b02541
# <翻译结束>


<原文开始>
// WhereOrf builds `OR` condition string using fmt.Sprintf and arguments.
<原文结束>

# <翻译开始>
// WhereOrf 使用fmt.Sprintf和参数构建`OR`条件字符串。. md5:aa04236f081a2885
# <翻译结束>


<原文开始>
// WhereOrf builds `OR` condition string using fmt.Sprintf and arguments.
// Eg:
// WhereOrf(`amount<? and status=%s`, "paid", 100)  => WHERE xxx OR `amount`<100 and status='paid'
// WhereOrf(`amount<%d and status=%s`, 100, "paid") => WHERE xxx OR `amount`<100 and status='paid'
<原文结束>

# <翻译开始>
// WhereOrf 使用 fmt.Sprintf 和参数构建 `OR` 条件字符串。
// 例如：
// WhereOrf(`amount<? and status=%s`, "paid", 100) => WHERE xxx OR `amount`<100 and status='paid'
// WhereOrf(`amount<%d and status=%s`, 100, "paid") => WHERE xxx OR `amount`<100 and status='paid' 
// 
// 这个函数用于构造一个 SQL 查询中的 `OR` 条件，其中 `%s` 和 `%d` 是占位符，会被传入的字符串和整数替换。它用于在 SQL 语句中添加额外的条件，当某个条件不满足时，返回的结果中包含 `OR` 连接的条件。
// md5:58173d98ee55b521
# <翻译结束>


<原文开始>
// WhereOrNot builds `column != value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrNot在`OR`条件下构建`column != value`语句。. md5:adc6d63e61bf279f
# <翻译结束>


<原文开始>
// WhereOrLT builds `column < value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrLT 在 `OR` 条件下构建 `column < value` 的语句。. md5:5517b3812e2c8e8b
# <翻译结束>


<原文开始>
// WhereOrLTE builds `column <= value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrLTE 在 OR 条件中构建 `column <= value` 语句。. md5:3b0287bd1f8030ce
# <翻译结束>


<原文开始>
// WhereOrGT builds `column > value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrGT在`OR`条件下构建`column > value`语句。. md5:2289d39bb82e521f
# <翻译结束>


<原文开始>
// WhereOrGTE builds `column >= value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrGTE在`OR`条件下构建`column >= value`语句。. md5:e178dd8cfc5661e5
# <翻译结束>


<原文开始>
// WhereOrBetween builds `column BETWEEN min AND max` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrBetween 用于构建 `column BETWEEN min AND max` 语句，并在 `OR` 条件下使用。. md5:90f98622a1fd5981
# <翻译结束>


<原文开始>
// WhereOrLike builds `column LIKE 'like'` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrLike 在 `OR` 条件中构建 `column LIKE 'like'` 语句。. md5:7a2d37411752fb51
# <翻译结束>


<原文开始>
// WhereOrIn builds `column IN (in)` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrIn 在`OR`条件下构建`column IN (in)`语句。. md5:4bb93b5ae9a5e887
# <翻译结束>


<原文开始>
// WhereOrNull builds `columns[0] IS NULL OR columns[1] IS NULL ...` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrNull 在 `OR` 条件下构建 `columns[0] IS NULL OR columns[1] IS NULL ...` 语句。. md5:08d38a60dc594441
# <翻译结束>


<原文开始>
// WhereOrNotBetween builds `column NOT BETWEEN min AND max` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrNotBetween 用于构建在 `OR` 条件下的 `column NOT BETWEEN min AND max` 语句。. md5:f20408e0126bbbab
# <翻译结束>


<原文开始>
// WhereOrNotLike builds `column NOT LIKE like` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrNotLike 在 OR 条件中构建 `column NOT LIKE like` 语句。. md5:751e840816119632
# <翻译结束>


<原文开始>
// WhereOrNotIn builds `column NOT IN (in)` statement.
<原文结束>

# <翻译开始>
// WhereOrNotIn构建`column NOT IN (in)`语句。. md5:433fd8a0f224fc24
# <翻译结束>


<原文开始>
// WhereOrNotNull builds `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrNotNull 构建 `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` 的 `OR` 条件语句。. md5:e122f662846a4ba4
# <翻译结束>

