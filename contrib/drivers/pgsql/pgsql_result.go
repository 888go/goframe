// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package pgsql

import "database/sql"

type Result struct {
	sql.Result
	affected          int64
	lastInsertId      int64
	lastInsertIdError error
}

func (pgr Result) RowsAffected() (int64, error) {
	return pgr.affected, nil
}

func (pgr Result) LastInsertId() (int64, error) {
	return pgr.lastInsertId, pgr.lastInsertIdError
}
