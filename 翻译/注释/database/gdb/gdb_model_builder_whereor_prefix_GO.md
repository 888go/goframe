
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
// WhereOrPrefix performs as WhereOr, but it adds prefix to each field in where statement.
// Eg:
// WhereOrPrefix("order", "status", "paid")                        => WHERE xxx OR (`order`.`status`='paid')
// WhereOrPrefix("order", struct{Status:"paid", "channel":"bank"}) => WHERE xxx OR (`order`.`status`='paid' AND `order`.`channel`='bank')
<原文结束>

# <翻译开始>
// WhereOrPrefix 的功能类似于 WhereOr，但它会在 WHERE 子句中的每个字段前添加一个前缀。
// 例如：
// WhereOrPrefix("order", "status", "paid")                        => WHERE xxx OR (`order`.`status`='paid')
// WhereOrPrefix("order", struct{Status:"paid", "channel":"bank"}) => WHERE xxx OR (`order`.`status`='paid' AND `order`.`channel`='bank') 
// 
// 这意味着 WhereOrPrefix 函数允许你在一个 WHERE 子句中指定多个条件，并且自动为这些条件的字段名加上一个指定的前缀，以便清晰地指向某个表或结构。它可以处理单个字段值的情况，也可以处理包含多个键值对的结构体，以生成更复杂的逻辑组合。
// md5:2358d43541f472f5
# <翻译结束>


<原文开始>
// WhereOrPrefixNot builds `prefix.column != value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixNot 在 OR 条件中构建 `prefix.column != value` 语句。 md5:385a9f9fb58b8fc3
# <翻译结束>


<原文开始>
// WhereOrPrefixLT builds `prefix.column < value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixLT在"OR"条件下构建`prefix.column < value`语句。 md5:c1a6baf94f553043
# <翻译结束>


<原文开始>
// WhereOrPrefixLTE builds `prefix.column <= value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixLTE 在“OR”条件下构建 `prefix.column <= value` 语句。 md5:77072877b38f04a8
# <翻译结束>


<原文开始>
// WhereOrPrefixGT builds `prefix.column > value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixGT 在 `OR` 条件下构建 `prefix.column > value` 的语句。 md5:d34b5bdc0e6b2fa8
# <翻译结束>


<原文开始>
// WhereOrPrefixGTE builds `prefix.column >= value` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixGTE 在 OR 条件中构建 `prefix.column >= value` 语句。 md5:d652ca0304ac835e
# <翻译结束>


<原文开始>
// WhereOrPrefixBetween builds `prefix.column BETWEEN min AND max` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixBetween在`OR`条件下构建`prefix.column BETWEEN min AND max`语句。 md5:d7adaf273fa5681b
# <翻译结束>


<原文开始>
// WhereOrPrefixLike builds `prefix.column LIKE 'like'` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixLike在`OR`条件下构建`prefix.column LIKE 'like'`语句。 md5:c975b47e3a5cc2c1
# <翻译结束>


<原文开始>
// WhereOrPrefixIn builds `prefix.column IN (in)` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixIn 用于构建 `prefix.column IN (in)` 形式的 SQL 语句，在 `OR` 条件下。 md5:18e0cf5cc971267d
# <翻译结束>


<原文开始>
// WhereOrPrefixNull builds `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixNull 在"OR"条件中构建`prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...`语句。 md5:facf88eb72f3d299
# <翻译结束>


<原文开始>
// WhereOrPrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixNotBetween在`OR`条件下构建`prefix.column NOT BETWEEN min AND max`语句。 md5:15259f135308893b
# <翻译结束>


<原文开始>
// WhereOrPrefixNotLike builds `prefix.column NOT LIKE 'like'` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixNotLike 在 `OR` 条件下构建 `prefix.column NOT LIKE 'like'` 语句。 md5:2785cbc79e811104
# <翻译结束>


<原文开始>
// WhereOrPrefixNotIn builds `prefix.column NOT IN (in)` statement.
<原文结束>

# <翻译开始>
// WhereOrPrefixNotIn 用于构建 `prefix.column NOT IN (in)` 的SQL语句。 md5:bd296110bb5635a1
# <翻译结束>


<原文开始>
// WhereOrPrefixNotNull builds `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` statement in `OR` conditions.
<原文结束>

# <翻译开始>
// WhereOrPrefixNotNull 在`OR`条件中构建`prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...`语句。 md5:9ecd3bffabf47cb7
# <翻译结束>

