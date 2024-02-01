// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb
import (
	"context"
	"database/sql"
	)
// Stmt 是一个预编译的SQL语句。
// Stmt 对象可以安全地被多个goroutine并发使用。
//
// 如果Stmt在一个Tx或Conn上预编译，那么它将永久绑定到单个底层连接。如果Tx或Conn关闭，
// Stmt将变得无法使用，所有操作都将返回错误。
// 如果Stmt在一个DB上预编译，那么只要DB存在，Stmt就会保持有效。当Stmt需要在新的底层连接上执行时，
// 它会自动在新连接上重新预编译自己。
type Stmt struct {
	*sql.Stmt
	core *Core
	link Link
	sql  string
}

// ExecContext 在给定参数的情况下执行预编译的语句，并返回一个Result，该Result总结了该语句的效果。
func (s *Stmt) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	out, err := s.core.db.DoCommit(ctx, DoCommitInput{
		Stmt:          s.Stmt,
		Link:          s.link,
		Sql:           s.sql,
		Args:          args,
		Type:          SqlTypeStmtExecContext,
		IsTransaction: s.link.IsTransaction(),
	})
	return out.Result, err
}

// QueryContext 执行带有给定参数的预编译查询语句，并将查询结果以 *Rows 类型返回。
func (s *Stmt) QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error) {
	out, err := s.core.db.DoCommit(ctx, DoCommitInput{
		Stmt:          s.Stmt,
		Link:          s.link,
		Sql:           s.sql,
		Args:          args,
		Type:          SqlTypeStmtQueryContext,
		IsTransaction: s.link.IsTransaction(),
	})
	if err != nil {
		return nil, err
	}
	if out.RawResult != nil {
		return out.RawResult.(*sql.Rows), err
	}
	return nil, nil
}

// QueryRowContext 执行一个带有给定参数的预编译查询语句。
// 如果在执行该语句时发生错误，该错误将在对返回的 *Row 调用 Scan 时返回，且返回的 *Row 始终是非 nil 的。
// 若查询未选择任何行，则 *Row 的 Scan 将返回 ErrNoRows 错误。
// 否则，*Row 的 Scan 会扫描第一条被选择的行，并丢弃其余行。
func (s *Stmt) QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row {
	out, err := s.core.db.DoCommit(ctx, DoCommitInput{
		Stmt:          s.Stmt,
		Link:          s.link,
		Sql:           s.sql,
		Args:          args,
		Type:          SqlTypeStmtQueryContext,
		IsTransaction: s.link.IsTransaction(),
	})
	if err != nil {
		panic(err)
	}
	if out.RawResult != nil {
		return out.RawResult.(*sql.Row)
	}
	return nil
}

// Exec 执行已准备好的语句，使用给定的参数，并返回一个 Result 对象，该对象总结了该语句的效果。
func (s *Stmt) Exec(args ...interface{}) (sql.Result, error) {
	return s.ExecContext(context.Background(), args...)
}

// Query函数通过给定的参数执行预编译查询语句，并将查询结果以*Rows的形式返回。
func (s *Stmt) Query(args ...interface{}) (*sql.Rows, error) {
	return s.QueryContext(context.Background(), args...)
}

// QueryRow 执行一个预编译的查询语句，同时传入给定的参数。
// 如果在执行语句过程中发生错误，该错误将在对返回的 *Row 调用 Scan 方法时返回（Scan 方法始终不为 nil）。
// 如果查询未选择任何行，则 *Row 的 Scan 方法将返回 ErrNoRows 错误。
// 否则，*Row 的 Scan 方法将扫描所选中的第一行数据，并丢弃其余数据。
//
// 示例用法：
//
//	var name string // 声明一个字符串变量 name
//	err := nameByUseridStmt.QueryRow(id).Scan(&name) // 使用预编译的查询语句并传入 id 参数，然后将查询结果扫描到变量 name 中。
func (s *Stmt) QueryRow(args ...interface{}) *sql.Row {
	return s.QueryRowContext(context.Background(), args...)
}

// Close 关闭该语句。
func (s *Stmt) Close() error {
	return s.Stmt.Close()
}
