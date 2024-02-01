
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
// Stmt 是一个预编译的SQL语句。
// Stmt 对象可以安全地被多个goroutine并发使用。
//
// 如果Stmt在一个Tx或Conn上预编译，那么它将永久绑定到单个底层连接。如果Tx或Conn关闭，
// Stmt将变得无法使用，所有操作都将返回错误。
// 如果Stmt在一个DB上预编译，那么只要DB存在，Stmt就会保持有效。当Stmt需要在新的底层连接上执行时，
// 它会自动在新连接上重新预编译自己。
# <翻译结束>


<原文开始>
// ExecContext executes a prepared statement with the given arguments and
// returns a Result summarizing the effect of the statement.
<原文结束>

# <翻译开始>
// ExecContext 在给定参数的情况下执行预编译的语句，并返回一个Result，该Result总结了该语句的效果。
# <翻译结束>


<原文开始>
// QueryContext executes a prepared query statement with the given arguments
// and returns the query results as a *Rows.
<原文结束>

# <翻译开始>
// QueryContext 执行带有给定参数的预编译查询语句，并将查询结果以 *Rows 类型返回。
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
// QueryRowContext 执行一个带有给定参数的预编译查询语句。
// 如果在执行该语句时发生错误，该错误将在对返回的 *Row 调用 Scan 时返回，且返回的 *Row 始终是非 nil 的。
// 若查询未选择任何行，则 *Row 的 Scan 将返回 ErrNoRows 错误。
// 否则，*Row 的 Scan 会扫描第一条被选择的行，并丢弃其余行。
# <翻译结束>


<原文开始>
// Exec executes a prepared statement with the given arguments and
// returns a Result summarizing the effect of the statement.
<原文结束>

# <翻译开始>
// Exec 执行已准备好的语句，使用给定的参数，并返回一个 Result 对象，该对象总结了该语句的效果。
# <翻译结束>


<原文开始>
// Query executes a prepared query statement with the given arguments
// and returns the query results as a *Rows.
<原文结束>

# <翻译开始>
// Query函数通过给定的参数执行预编译查询语句，并将查询结果以*Rows的形式返回。
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
// QueryRow 执行一个预编译的查询语句，同时传入给定的参数。
// 如果在执行语句过程中发生错误，该错误将在对返回的 *Row 调用 Scan 方法时返回（Scan 方法始终不为 nil）。
// 如果查询未选择任何行，则 *Row 的 Scan 方法将返回 ErrNoRows 错误。
// 否则，*Row 的 Scan 方法将扫描所选中的第一行数据，并丢弃其余数据。
//
// 示例用法：
//
//	var name string // 声明一个字符串变量 name
//	err := nameByUseridStmt.QueryRow(id).Scan(&name) // 使用预编译的查询语句并传入 id 参数，然后将查询结果扫描到变量 name 中。
# <翻译结束>


<原文开始>
// Close closes the statement.
<原文结束>

# <翻译开始>
// Close 关闭该语句。
# <翻译结束>

