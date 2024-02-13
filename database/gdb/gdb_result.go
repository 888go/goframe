// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"database/sql"
	
	"github.com/888go/goframe/errors/gerror"
)

// SqlResult 是 SQL 操作的执行结果。
// 同时，它还支持批量操作的结果，其中包含 rowsAffected（受影响的行数）。
type Sql执行结果 struct {
	X原生sql行记录   sql.Result
	X影响行数 int64
}

// MustGetAffected 返回受影响的行数，如果发生任何错误，它会触发panic（即程序崩溃）。
func (r *Sql执行结果) X取影响行数PANI() int64 {
	rows, err := r.RowsAffected()
	if err != nil {
		err = 错误类.X多层错误(err, `sql.Result.RowsAffected failed`)
		panic(err)
	}
	return rows
}

// MustGetInsertId 返回最后一次插入的ID，如果发生任何错误，它会引发panic。
func (r *Sql执行结果) X取最后插入IdPANI() int64 {
	id, err := r.LastInsertId()
	if err != nil {
		err = 错误类.X多层错误(err, `sql.Result.LastInsertId failed`)
		panic(err)
	}
	return id
}

// 2024-01-21 X取影响行数(),这个方法不能翻译, 翻译后返回的结构体会导致未实现RowsAffected方法.
// RowsAffected 返回由更新、插入或删除操作影响的行数。
// 但请注意，并非所有数据库或数据库驱动程序都支持此功能。
// 另外，请参阅 sql.Result。
func (r *Sql执行结果) RowsAffected() (int64, error) {
	if r.X影响行数 > 0 {
		return r.X影响行数, nil
	}
	if r.X原生sql行记录 == nil {
		return 0, nil
	}
	return r.X原生sql行记录.RowsAffected()
}

// 2024-01-21 X取最后插入id() 这个方法不能翻译, 翻译后返回的结构体会导致未实现LastInsertId方法.
// LastInsertId 返回数据库对命令执行后生成的整数。通常，这个值来自于插入新行时“自动递增”列。并非所有数据库都支持此特性，并且此类语句的语法也各不相同。
// 另外，请参阅 sql.Result。
func (r *Sql执行结果) LastInsertId() (int64, error) {
	if r.X原生sql行记录 == nil {
		return 0, nil
	}
	return r.X原生sql行记录.LastInsertId()
}
