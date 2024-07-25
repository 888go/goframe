
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// WherePrefix performs as Where, but it adds prefix to each field in where statement.
// Eg:
// WherePrefix("order", "status", "paid")                        => WHERE `order`.`status`='paid'
// WherePrefix("order", struct{Status:"paid", "channel":"bank"}) => WHERE `order`.`status`='paid' AND `order`.`channel`='bank'
<原文结束>

# <翻译开始>
// WherePrefix 的功能类似于 Where，但它会在 where 语句中的每个字段前添加一个前缀。
// 例如：
// WherePrefix("order", "status", "paid")                        => WHERE `order`.`status`='paid'
// WherePrefix("order", struct{Status:"paid", "channel":"bank"}) => WHERE `order`.`status`='paid' AND `order`.`channel`='bank'
//
// 这段注释的中文翻译为：
//
// WherePrefix 的行为与 Where 相似，但它会在 where 子句里的每个字段前加上一个前缀。
// 例如：
// WherePrefix("order", "status", "paid")                        => 生成 WHERE `order`.`status`='paid'
// WherePrefix("order", struct{Status:"paid", "channel":"bank"}) => 生成 WHERE `order`.`status`='paid' AND `order`.`channel`='bank' md5:062302edb484784b
# <翻译结束>


<原文开始>
// WherePrefixLT builds `prefix.column < value` statement.
<原文结束>

# <翻译开始>
// WherePrefixLT 构建 `prefix.column < value` 语句。 md5:de5cb5259ef84499
# <翻译结束>


<原文开始>
// WherePrefixLTE builds `prefix.column <= value` statement.
<原文结束>

# <翻译开始>
// WherePrefixLTE构建`prefix.column <= value`语句。 md5:1c5d93f173a39b03
# <翻译结束>


<原文开始>
// WherePrefixGT builds `prefix.column > value` statement.
<原文结束>

# <翻译开始>
// WherePrefixGT构建`prefix.column > value`语句。 md5:61d5cbbb9f5422fd
# <翻译结束>


<原文开始>
// WherePrefixGTE builds `prefix.column >= value` statement.
<原文结束>

# <翻译开始>
// WherePrefixGTE 生成 `prefix.column >= value` 的语句。 md5:1b581ea600e215e7
# <翻译结束>


<原文开始>
// WherePrefixBetween builds `prefix.column BETWEEN min AND max` statement.
<原文结束>

# <翻译开始>
// WherePrefixBetween 构建 `prefix.column BETWEEN min AND max` 语句。 md5:e6176c55b8a31575
# <翻译结束>


<原文开始>
// WherePrefixLike builds `prefix.column LIKE like` statement.
<原文结束>

# <翻译开始>
// WherePrefixLike构建`prefix.column LIKE like`语句。 md5:baf08eac5c7dc2aa
# <翻译结束>


<原文开始>
// WherePrefixIn builds `prefix.column IN (in)` statement.
<原文结束>

# <翻译开始>
// WherePrefixIn 构建 `prefix.column IN (in)` 语句。 md5:fd691f634711ba7f
# <翻译结束>


<原文开始>
// WherePrefixNull builds `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` statement.
<原文结束>

# <翻译开始>
// WherePrefixNull 构建 `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` 语句。 md5:ac08bde96048fdce
# <翻译结束>


<原文开始>
// WherePrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement.
<原文结束>

# <翻译开始>
// WherePrefixNotBetween 构建 `prefix.column NOT BETWEEN min AND max` 语句。 md5:a16703b511af05c3
# <翻译结束>


<原文开始>
// WherePrefixNotLike builds `prefix.column NOT LIKE like` statement.
<原文结束>

# <翻译开始>
// WherePrefixNotLike构建`prefix.column NOT LIKE like`语句。 md5:083bd1d45c368a83
# <翻译结束>


<原文开始>
// WherePrefixNot builds `prefix.column != value` statement.
<原文结束>

# <翻译开始>
// WherePrefixNot构建`prefix.column != value`语句。 md5:c1366e00cd0da49e
# <翻译结束>


<原文开始>
// WherePrefixNotIn builds `prefix.column NOT IN (in)` statement.
<原文结束>

# <翻译开始>
// WherePrefixNotIn 用于构建 `prefix.column NOT IN (in)` 的SQL语句。 md5:3b790678c07a51fd
# <翻译结束>


<原文开始>
// WherePrefixNotNull builds `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` statement.
<原文结束>

# <翻译开始>
// WherePrefixNotNull 构建 `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` 语句。 md5:d5a307a7c3004dda
# <翻译结束>

