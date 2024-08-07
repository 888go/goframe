// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"database/sql"

	gerror "github.com/888go/goframe/errors/gerror"
)

// SqlResult是SQL操作的执行结果。它还支持行影响的批量操作结果。
// md5:c89c5ab24627c936
type SqlResult struct {
	X原生sql行记录   sql.Result
	X影响行数 int64
}

// X取影响行数PANI 返回受影响的行数，如果发生任何错误，则会引发恐慌。 md5:be151685a0da2b44
func (r *SqlResult) X取影响行数PANI() int64 {
	rows, err := r.RowsAffected()
	if err != nil {
		err = gerror.X多层错误(err, `sql.Result.RowsAffected failed`)
		panic(err)
	}
	return rows
}

// X取最后插入IdPANI 返回最后的插入ID，如果发生任何错误，它将引发恐慌。 md5:bd23d169a4cb6738
func (r *SqlResult) X取最后插入IdPANI() int64 {
	id, err := r.LastInsertId()
	if err != nil {
		err = gerror.X多层错误(err, `sql.Result.LastInsertId failed`)
		panic(err)
	}
	return id
}

// RowsAffected 返回更新、插入或删除操作影响的行数。并非所有数据库或数据库驱动程序都支持此功能。
// 参见 sql.Result。
// md5:f41c8ccbf7344301
func (r *SqlResult) RowsAffected() (int64, error) {
	if r.X影响行数 > 0 {
		return r.X影响行数, nil
	}
	if r.X原生sql行记录 == nil {
		return 0, nil
	}
	return r.X原生sql行记录.RowsAffected()
}

// LastInsertId返回数据库对命令的响应生成的整数。通常，这将是在插入新行时来自“自动递增”列的。并非所有数据库都支持此功能，且此类语句的语法各不相同。
// 参见sql.Result。
// md5:7236c1ac3f4fc094
func (r *SqlResult) LastInsertId() (int64, error) {
	if r.X原生sql行记录 == nil {
		return 0, nil
	}
	return r.X原生sql行记录.LastInsertId()
}
