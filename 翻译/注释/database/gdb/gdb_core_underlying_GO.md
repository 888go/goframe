
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//
# <翻译结束>


<原文开始>
// Query commits one query SQL to underlying driver and returns the execution result.
// It is most commonly used for data querying.
<原文结束>

# <翻译开始>
// Query 将一个查询SQL语句提交给底层驱动并返回执行结果。
// 这个方法最常用于数据查询。
# <翻译结束>


<原文开始>
// DoQuery commits the sql string and its arguments to underlying driver
// through given link object and returns the execution result.
<原文结束>

# <翻译开始>
// DoQuery 通过给定的link对象，将sql字符串及其参数提交到底层驱动，并返回执行结果。
# <翻译结束>







<原文开始>
// Firstly, check and retrieve transaction link from context.
<原文结束>

# <翻译开始>
// 首先，从上下文检查并检索交易链接。
# <翻译结束>


<原文开始>
// Or else it creates one from master node.
<原文结束>

# <翻译开始>
// 或者从主节点创建一个
# <翻译结束>


<原文开始>
// If current link is not transaction link, it checks and retrieves transaction from context.
<原文结束>

# <翻译开始>
// 如果当前链接不是事务链接，它会检查并从上下文中获取事务。
# <翻译结束>







<原文开始>
// Exec commits one query SQL to underlying driver and returns the execution result.
// It is most commonly used for data inserting and updating.
<原文结束>

# <翻译开始>
// Exec方法将一个SQL查询语句提交给底层驱动执行并返回执行结果。
// 该方法主要用于数据的插入和更新操作。
# <翻译结束>


<原文开始>
// DoExec commits the sql string and its arguments to underlying driver
// through given link object and returns the execution result.
<原文结束>

# <翻译开始>
// DoExec通过给定的link对象，将SQL字符串及其参数提交给底层驱动，并返回执行结果。
# <翻译结束>


<原文开始>
// DoFilter is a hook function, which filters the sql and its arguments before it's committed to underlying driver.
// The parameter `link` specifies the current database connection operation object. You can modify the sql
// string `sql` and its arguments `args` as you wish before they're committed to driver.
<原文结束>

# <翻译开始>
// DoFilter 是一个钩子函数，在 SQL 语句及其参数提交给底层驱动程序之前对其进行过滤。
// 参数 `link` 指定了当前数据库连接操作对象。您可以在 SQL 字符串 `sql` 和其参数 `args` 提交给驱动程序之前，根据需要自由修改它们。
# <翻译结束>


<原文开始>
// DoCommit commits current sql and arguments to underlying sql driver.
<原文结束>

# <翻译开始>
// DoCommit 将当前SQL语句及其参数提交给底层SQL驱动执行。
# <翻译结束>


<原文开始>
// Inject internal data into ctx, especially for transaction creating.
<原文结束>

# <翻译开始>
// 将内部数据注入到ctx中，特别是用于创建事务。
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
// Prepare 函数用于为后续查询或执行创建预编译语句。
// 从返回的语句可以并发地运行多个查询或执行操作。
// 当该语句不再需要时，调用者必须调用该语句的 Close 方法。
//
// 参数 `execOnMaster` 指定是否在主节点上执行 SQL，如果配置了主从模式，
// 则此参数为 false 时将在从节点上执行 SQL。
# <翻译结束>


<原文开始>
// DoPrepare calls prepare function on given link object and returns the statement object.
<原文结束>

# <翻译开始>
// DoPrepare在给定的link对象上调用prepare函数，并返回statement对象。
# <翻译结束>


<原文开始>
// DO NOT USE cancel function in prepare statement.
<原文结束>

# <翻译开始>
// **请勿在预处理语句中使用取消函数。**
# <翻译结束>


<原文开始>
// RowsToResult converts underlying data record type sql.Rows to Result type.
<原文结束>

# <翻译开始>
// RowsToResult 将底层数据记录类型 sql.Rows 转换为 Result 类型。
# <翻译结束>







<原文开始>
				// DO NOT use `gvar.New(nil)` here as it creates an initialized object
				// which will cause struct converting issue.
<原文结束>

# <翻译开始>
// **注意**：在此处不要使用 `gvar.New(nil)`，因为它会创建一个已初始化的对象，
// 这将会导致结构体转换问题。
# <翻译结束>


<原文开始>
// Common basic builtin types.
<原文结束>

# <翻译开始>
// 常见的基本内置类型。
# <翻译结束>


<原文开始>
// Other complex types, especially custom types.
<原文结束>

# <翻译开始>
// 其他复杂类型，特别是自定义类型。
# <翻译结束>












<原文开始>
// Transaction checks.
<原文结束>

# <翻译开始>
// 事务检查。
# <翻译结束>


<原文开始>
// SQL format and retrieve.
<原文结束>

# <翻译开始>
// SQL格式化并检索
# <翻译结束>


<原文开始>
// Execution cased by type.
<原文结束>

# <翻译开始>
// 根据类型执行。
# <翻译结束>


<原文开始>
// Column names and types.
<原文结束>

# <翻译开始>
// 列名和类型。
# <翻译结束>

