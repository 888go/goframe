
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
// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter `table` can be joined table and its joined condition,
// and also with its alias name.
//
// Eg:
// Model("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").LeftJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid").
<原文结束>

# <翻译开始>
// LeftJoin 在模型上执行 "LEFT JOIN ... ON ..." 语句。
// 参数 `table` 可以是连接的表及其连接条件，也可以包含其别名。
//
// 示例：
// Model("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").LeftJoin("SELECT xxx FROM xxx", "a", "a.uid=u.uid")
// md5:5f7464280da64004
# <翻译结束>


<原文开始>
// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter `table` can be joined table and its joined condition,
// and also with its alias name.
//
// Eg:
// Model("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").RightJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid").
<原文结束>

# <翻译开始>
// RightJoin 执行 "RIGHT JOIN ... ON ..." 语句在模型上。
// 参数 `table` 可以是待连接的表及其连接条件，
// 也可以包含表的别名。
//
// 例如：
// Model("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").RightJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid")
// md5:dbab2528fb37c84e
# <翻译结束>


<原文开始>
// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter `table` can be joined table and its joined condition,
// and also with its alias name。
//
// Eg:
// Model("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").InnerJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid").
<原文结束>

# <翻译开始>
// InnerJoin 在模型上执行 "INNER JOIN ... ON ..." 语句。
// 参数 `table` 可以是需要连接的表及其连接条件，同时也可包含别名名称。
//
// 例如：
// Model("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").InnerJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid")
// md5:eda419ad685c559d
# <翻译结束>


<原文开始>
// LeftJoinOnField performs as LeftJoin, but it joins both tables with the `same field name`.
//
// Eg:
// Model("order").LeftJoinOnField("user", "user_id")
// Model("order").LeftJoinOnField("product", "product_id").
<原文结束>

# <翻译开始>
// LeftJoinOnField 执行左连接，但使用两个表中的`相同字段名`进行连接。
//
// 例如：
// Model("order").LeftJoinOnField("user", "user_id")
// Model("order").LeftJoinOnField("product", "product_id")
// md5:5ace5bdef45c73d6
# <翻译结束>


<原文开始>
// RightJoinOnField performs as RightJoin, but it joins both tables with the `same field name`.
//
// Eg:
// Model("order").InnerJoinOnField("user", "user_id")
// Model("order").InnerJoinOnField("product", "product_id").
<原文结束>

# <翻译开始>
// RightJoinOnField 执行右连接，但使用相同的字段名称连接两个表。
//
// 例如：
// Model("order").InnerJoinOnField("user", "user_id")
// Model("order").InnerJoinOnField("product", "product_id")
// md5:ac8281e2d383e3d6
# <翻译结束>


<原文开始>
// InnerJoinOnField performs as InnerJoin, but it joins both tables with the `same field name`.
//
// Eg:
// Model("order").InnerJoinOnField("user", "user_id")
// Model("order").InnerJoinOnField("product", "product_id").
<原文结束>

# <翻译开始>
// InnerJoinOnField 的行为类似于 InnerJoin，但它使用的是具有`相同字段名`的两个表进行连接。
//
// 例如：
// Model("order").InnerJoinOnField("user", "user_id")
// Model("order").InnerJoinOnField("product", "product_id")
// md5:bdc954b5bcb8a9c5
# <翻译结束>


<原文开始>
// LeftJoinOnFields performs as LeftJoin. It specifies different fields and comparison operator.
//
// Eg:
// Model("user").LeftJoinOnFields("order", "id", "=", "user_id")
// Model("user").LeftJoinOnFields("order", "id", ">", "user_id")
// Model("user").LeftJoinOnFields("order", "id", "<", "user_id")
<原文结束>

# <翻译开始>
// LeftJoinOnFields 执行类似于 LeftJoin 的操作，但允许指定不同的字段和比较运算符。
//
// 例如：
// Model("user").LeftJoinOnFields("order", "id", "=", "user_id")
// Model("user").LeftJoinOnFields("order", "id", ">", "user_id")
// Model("user").LeftJoinOnFields("order", "id", "<", "user_id")
// md5:90ce0e2226eb4b30
# <翻译结束>


<原文开始>
// RightJoinOnFields performs as RightJoin. It specifies different fields and comparison operator.
//
// Eg:
// Model("user").RightJoinOnFields("order", "id", "=", "user_id")
// Model("user").RightJoinOnFields("order", "id", ">", "user_id")
// Model("user").RightJoinOnFields("order", "id", "<", "user_id")
<原文结束>

# <翻译开始>
// RightJoinOnFields 执行右连接操作。它指定了不同的字段和比较运算符。
//
// 例如：
// User("user").RightJoinOnFields("order", "id", "=", "user_id")
// User("user").RightJoinOnFields("order", "id", ">", "user_id")
// User("user").RightJoinOnFields("order", "id", "<", "user_id") 
// 
// 这里，`RightJoinOnFields` 是一个 Go 代码中的函数，用于在查询数据库时执行右连接操作，并且允许用户自定义连接的字段和比较操作。第一个参数是模型名（如 "user"），接下来的参数包括要连接的表名、字段名以及连接条件（等于、大于或小于）。
// md5:563f2b0f155fc829
# <翻译结束>


<原文开始>
// InnerJoinOnFields performs as InnerJoin. It specifies different fields and comparison operator.
//
// Eg:
// Model("user").InnerJoinOnFields("order", "id", "=", "user_id")
// Model("user").InnerJoinOnFields("order", "id", ">", "user_id")
// Model("user").InnerJoinOnFields("order", "id", "<", "user_id")
<原文结束>

# <翻译开始>
// InnerJoinOnFields 执行 InnerJoin 操作。它指定了不同的字段和比较运算符。
// 
// 例如：
// Model("user").InnerJoinOnFields("order", "id", "=", "user_id")
// Model("user").InnerJoinOnFields("order", "id", ">", "user_id")
// Model("user").InnerJoinOnFields("order", "id", "<", "user_id") 
// 
// 这段代码是在 Go 语言中定义了一个方法，用于在两个数据表之间执行内连接（InnerJoin），并允许用户指定连接的字段以及比较运算符。例如，`"user".InnerJoinOnFields("order", "id", "=", "user_id")` 表示连接 "user" 表和 "order" 表，通过 "id" 字段进行等号（=）匹配。其他示例展示了使用大于（>）和小于（<）运算符的情况。
// md5:0499f4b5bbbc2016
# <翻译结束>


<原文开始>
// doJoin does "LEFT/RIGHT/INNER JOIN ... ON ..." statement on the model.
// The parameter `tableOrSubQueryAndJoinConditions` can be joined table and its joined condition,
// and also with its alias name.
//
// Eg:
// Model("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid>u.uid")
// Model("user", "u").InnerJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid")
// Related issues:
// https://github.com/gogf/gf/issues/1024
<原文结束>

# <翻译开始>
// doJoin 在模型上执行 "LEFT/RIGHT/INNER JOIN ... ON ..." 语句。
// 参数 `tableOrSubQueryAndJoinConditions` 可以是待连接的表及其连接条件，
// 同时也可以包含表的别名。
//
// 例如：
// Model("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid>u.uid")
// Model("user", "u").InnerJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid")
// 相关问题讨论：
// https://github.com/gogf/gf/issues/1024
// md5:7b792ce187933a04
# <翻译结束>


<原文开始>
// Check the first parameter table or sub-query.
<原文结束>

# <翻译开始>
	// 检查第一个参数，是否为表格或子查询。 md5:0493998c4b03304e
# <翻译结束>


<原文开始>
// Generate join condition statement string.
<原文结束>

# <翻译开始>
	// 生成连接条件的字符串表达式。 md5:54f67a1d882ecd10
# <翻译结束>


<原文开始>
// getTableNameByPrefixOrAlias checks and returns the table name if `prefixOrAlias` is an alias of a table,
// it or else returns the `prefixOrAlias` directly.
<原文结束>

# <翻译开始>
// getTableNameByPrefixOrAlias 检查`prefixOrAlias`是否是某个表的别名，如果是，则返回该表的实际名称，否则直接返回`prefixOrAlias`。
// md5:ab423b9e1e0ad0ca
# <翻译结束>


<原文开始>
// isSubQuery checks and returns whether given string a sub-query sql string.
<原文结束>

# <翻译开始>
// isSubQuery 检查并返回给定的字符串是否为子查询SQL语句。 md5:0921761c51f20650
# <翻译结束>

