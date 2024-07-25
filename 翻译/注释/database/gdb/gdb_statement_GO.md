
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
// Stmt is a prepared statement.
// A Stmt is safe for concurrent use by multiple goroutines.
//
// If a Stmt is prepared on a Tx or Conn, it will be bound to a single
// underlying connection forever. If the Tx or Conn closes, the Stmt will
// become unusable and all operations will return an error.
// If a Stmt is prepared on a DB, it will remain usable for the lifetime of the
// DB. When the Stmt needs to execute on a new underlying connection, it will
// prepare itself on the new connection automatically.
<原文结束>

# <翻译开始>
// Stmt 是一个预处理语句。
// Stmt 安全地支持多个 goroutine 并发使用。
//
// 如果在一个 Tx 或 Conn 上准备了一个 Stmt，它将永久绑定到单一的底层连接上。如果该 Tx 或 Conn 关闭了，Stmt 将变得无法使用，并且所有操作都将返回错误。
// 如果在一个 DB 上准备了一个 Stmt，它将在 DB 的整个生命周期内保持可用。当 Stmt 需要在新的底层连接上执行时，它会自动在新连接上为自己做准备。
// md5:cf1394a31e88aa92
# <翻译结束>


<原文开始>
// ExecContext executes a prepared statement with the given arguments and
// returns a Result summarizing the effect of the statement.
<原文结束>

# <翻译开始>
// ExecContext 使用给定的参数执行一个预编译的语句，并返回一个总结语句影响的结果。
// md5:aa3184ee6004a352
# <翻译结束>


<原文开始>
// QueryContext executes a prepared query statement with the given arguments
// and returns the query results as a *Rows.
<原文结束>

# <翻译开始>
// QueryContext 使用给定的参数执行一个预处理查询语句，并将查询结果作为*Rows返回。
// md5:2dc9d8dfa34a3cb1
# <翻译结束>


<原文开始>
// QueryRowContext executes a prepared query statement with the given arguments.
// If an error occurs during the execution of the statement, that error will
// be returned by a call to Scan on the returned *Row, which is always non-nil.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
<原文结束>

# <翻译开始>
// QueryRowContext 使用给定的参数执行预处理查询语句。
// 如果在执行语句时发生错误，该错误将通过调用返回的*Row的Scan方法返回，该*Row始终非空。
// 如果查询没有选择任何行，*Row的Scan将返回ErrNoRows。
// 否则，*Row的Scan会扫描选中的第一行并丢弃其余部分。
// md5:a2aed2939b0e481c
# <翻译结束>


<原文开始>
// Exec executes a prepared statement with the given arguments and
// returns a Result summarizing the effect of the statement.
<原文结束>

# <翻译开始>
// Exec 使用给定的参数执行预处理语句，并返回一个 Result 来概括该语句的影响。
// md5:7ff0a9731fb27dd1
# <翻译结束>


<原文开始>
// Query executes a prepared query statement with the given arguments
// and returns the query results as a *Rows.
<原文结束>

# <翻译开始>
// Query 使用给定的参数执行预编译的查询语句，并将查询结果返回为一个 *Rows 对象。
// md5:5028edd0d7332759
# <翻译结束>


<原文开始>
// QueryRow executes a prepared query statement with the given arguments.
// If an error occurs during the execution of the statement, that error will
// be returned by a call to Scan on the returned *Row, which is always non-nil.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
//
// Example usage:
//
//	var name string
//	err := nameByUseridStmt.QueryRow(id).Scan(&name)
<原文结束>

# <翻译开始>
// QueryRow 执行一个预处理的查询语句，使用给定的参数。如果在执行语句时发生错误，该错误将通过调用返回的 *Row 的 Scan 方法返回，*Row 总是非空的。如果查询没有选择任何行，*Row 的 Scan 将返回 ErrNoRows。否则，*Row 的 Scan 会扫描选中的第一行并忽略其余行。
// 
// 示例用法：
//
//	var name string
//	err := nameByUseridStmt.QueryRow(id).Scan(&name)
// md5:3b2851663bab6cd6
# <翻译结束>


<原文开始>
// Close closes the statement.
<原文结束>

# <翻译开始>
// Close关闭语句。 md5:73ebf594ef1ccf7c
# <翻译结束>

