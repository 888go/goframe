// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"context"
	"database/sql"
)

// Stmt 是一个预处理语句。
// Stmt 安全地支持多个 goroutine 并发使用。
//
// 如果在一个 Tx 或 Conn 上准备了一个 Stmt，它将永久绑定到单一的底层连接上。如果该 Tx 或 Conn 关闭了，Stmt 将变得无法使用，并且所有操作都将返回错误。
// 如果在一个 DB 上准备了一个 Stmt，它将在 DB 的整个生命周期内保持可用。当 Stmt 需要在新的底层连接上执行时，它会自动在新连接上为自己做准备。
// md5:cf1394a31e88aa92
type Stmt struct {
	*sql.Stmt
	core *Core
	link Link
	sql  string
}

// ExecContext 使用给定的参数执行一个预编译的语句，并返回一个总结语句影响的结果。
// md5:aa3184ee6004a352
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

// QueryContext 使用给定的参数执行一个预处理查询语句，并将查询结果作为*Rows返回。
// md5:2dc9d8dfa34a3cb1
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

// QueryRowContext 使用给定的参数执行预处理查询语句。
// 如果在执行语句时发生错误，该错误将通过调用返回的*Row的Scan方法返回，该*Row始终非空。
// 如果查询没有选择任何行，*Row的Scan将返回ErrNoRows。
// 否则，*Row的Scan会扫描选中的第一行并丢弃其余部分。
// md5:a2aed2939b0e481c
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

// Exec 使用给定的参数执行预处理语句，并返回一个 Result 来概括该语句的影响。
// md5:7ff0a9731fb27dd1
func (s *Stmt) Exec(args ...interface{}) (sql.Result, error) {
	return s.ExecContext(context.Background(), args...)
}

// Query 使用给定的参数执行预编译的查询语句，并将查询结果返回为一个 *Rows 对象。
// md5:5028edd0d7332759
func (s *Stmt) Query(args ...interface{}) (*sql.Rows, error) {
	return s.QueryContext(context.Background(), args...)
}

// QueryRow 执行一个预处理的查询语句，使用给定的参数。如果在执行语句时发生错误，该错误将通过调用返回的 *Row 的 Scan 方法返回，*Row 总是非空的。如果查询没有选择任何行，*Row 的 Scan 将返回 ErrNoRows。否则，*Row 的 Scan 会扫描选中的第一行并忽略其余行。
// 
// 示例用法：
//
//	var name string
//	err := nameByUseridStmt.QueryRow(id).Scan(&name)
// md5:3b2851663bab6cd6
func (s *Stmt) QueryRow(args ...interface{}) *sql.Row {
	return s.QueryRowContext(context.Background(), args...)
}

// Close关闭语句。. md5:73ebf594ef1ccf7c
func (s *Stmt) Close() error {
	return s.Stmt.Close()
}
