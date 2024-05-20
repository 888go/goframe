
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31
# <翻译结束>


<原文开始>
// Query commits one query SQL to underlying driver and returns the execution result.
// It is most commonly used for data querying.
<原文结束>

# <翻译开始>
// Query 向底层驱动提交一个查询SQL并返回执行结果。
// 它最常用于数据查询。
// md5:06bbbfc29aa3894b
# <翻译结束>


<原文开始>
// DoQuery commits the sql string and its arguments to underlying driver
// through given link object and returns the execution result.
<原文结束>

# <翻译开始>
// DoQuery 通过给定的链接对象将SQL字符串及其参数提交给底层驱动，并返回执行结果。
// md5:af7bdcd1a2074bc0
# <翻译结束>


<原文开始>
// Firstly, check and retrieve transaction link from context.
<原文结束>

# <翻译开始>
// 首先，从上下文中检查并获取交易链接。. md5:9ac4c60388fa960d
# <翻译结束>


<原文开始>
// Or else it creates one from master node.
<原文结束>

# <翻译开始>
// 否则，它将从主节点创建一个。. md5:4bd14606783b43fc
# <翻译结束>


<原文开始>
// If current link is not transaction link, it checks and retrieves transaction from context.
<原文结束>

# <翻译开始>
// 如果当前链接不是事务链接，它会检查并从上下文中检索事务。. md5:e3c484ab061699a1
# <翻译结束>


<原文开始>
// SQL format and retrieve.
<原文结束>

# <翻译开始>
// SQL格式化和获取。. md5:815f530302ec8a7e
# <翻译结束>


<原文开始>
// Exec commits one query SQL to underlying driver and returns the execution result.
// It is most commonly used for data inserting and updating.
<原文结束>

# <翻译开始>
// Exec 将一个查询 SQL 执行到底层驱动并返回执行结果。它最常用于数据插入和更新。
// md5:6f9ddc85964b9797
# <翻译结束>


<原文开始>
// DoExec commits the sql string and its arguments to underlying driver
// through given link object and returns the execution result.
<原文结束>

# <翻译开始>
// DoExec 通过给定的链接对象将 sql 字符串及其参数提交到底层驱动，并返回执行结果。
// md5:947bd2b83e751e10
# <翻译结束>


<原文开始>
// DoFilter is a hook function, which filters the sql and its arguments before it's committed to underlying driver.
// The parameter `link` specifies the current database connection operation object. You can modify the sql
// string `sql` and its arguments `args` as you wish before they're committed to driver.
<原文结束>

# <翻译开始>
// DoFilter 是一个钩子函数，它在 SQL 语句及其参数提交给底层驱动之前进行过滤。
// 参数 `link` 指定当前数据库连接的操作对象。在 SQL 语句 `sql` 及其参数 `args` 被提交给驱动之前，您可以根据需要随意修改它们。
// md5:41118fbc4e6c5562
# <翻译结束>


<原文开始>
// DoCommit commits current sql and arguments to underlying sql driver.
<原文结束>

# <翻译开始>
// DoCommit 将当前SQL和参数提交给底层SQL驱动程序。. md5:7cf9b1f6f4d9d2cb
# <翻译结束>


<原文开始>
// Execution cased by type.
<原文结束>

# <翻译开始>
// 根据类型执行。. md5:4f8bf756ef2da0c5
# <翻译结束>


<原文开始>
// Prepare creates a prepared statement for later queries or executions.
// Multiple queries or executions may be run concurrently from the
// returned statement.
// The caller must call the statement's Close method
// when the statement is no longer needed.
//
// The parameter `execOnMaster` specifies whether executing the sql on master node,
// or else it executes the sql on slave node if master-slave configured.
<原文结束>

# <翻译开始>
// Prepare 准备一个预编译语句，供后续查询或执行使用。
// 可以从返回的语句对象并发运行多个查询或执行。
// 当不再需要语句时，调用者必须调用 statement 的 Close 方法。
// 
// 参数 `execOnMaster` 指定是否在主节点上执行 SQL，如果配置了主从复制，则在从节点上执行。
// md5:639eebcae369b0a2
# <翻译结束>


<原文开始>
// DoPrepare calls prepare function on given link object and returns the statement object.
<原文结束>

# <翻译开始>
// DoPrepare 会调用给定链接对象上的prepare函数，并返回语句对象。. md5:bae03ede256987bd
# <翻译结束>


<原文开始>
// DO NOT USE cancel function in prepare statement.
<原文结束>

# <翻译开始>
// 不要在预处理语句中使用取消函数。. md5:5e529fe5094c7942
# <翻译结束>


<原文开始>
// FormatUpsert formats and returns SQL clause part for upsert statement.
// In default implements, this function performs upsert statement for MySQL like:
// `INSERT INTO ... ON DUPLICATE KEY UPDATE x=VALUES(z),m=VALUES(y)...`
<原文结束>

# <翻译开始>
// FormatUpsert 格式化并返回用于 UPSERT 语句的 SQL 子句部分。
// 在默认实现中，此函数执行类似 MySQL 的 UPSERT 语句：
// `INSERT INTO ... ON DUPLICATE KEY UPDATE x=VALUES(z),m=VALUES(y)...`
// md5:c1c6d7b14661682b
# <翻译结束>


<原文开始>
// If it's SAVE operation, do not automatically update the creating time.
<原文结束>

# <翻译开始>
// 如果是SAVE操作，不要自动更新创建时间。. md5:409c9c162d30afae
# <翻译结束>


<原文开始>
// RowsToResult converts underlying data record type sql.Rows to Result type.
<原文结束>

# <翻译开始>
// RowsToResult 将底层的 sql.Rows 数据记录类型转换为 Result 类型。. md5:ae9065176ef07b2e
# <翻译结束>


<原文开始>
// Column names and types.
<原文结束>

# <翻译开始>
// 列名和类型。. md5:51cafb00c4482aba
# <翻译结束>


<原文开始>
				// DO NOT use `gvar.New(nil)` here as it creates an initialized object
				// which will cause struct converting issue.
<原文结束>

# <翻译开始>
// 不要在这里使用 `gvar.New(nil)`，因为它会创建一个已初始化的对象，
// 这将导致结构体转换问题。
// md5:4df778f025fefd53
# <翻译结束>


<原文开始>
// Common basic builtin types.
<原文结束>

# <翻译开始>
// 常见的基本内置类型。. md5:4c57bcc430188806
# <翻译结束>


<原文开始>
// Other complex types, especially custom types.
<原文结束>

# <翻译开始>
// 其他复杂类型，特别是自定义类型。. md5:5d9bae215068a0c1
# <翻译结束>

