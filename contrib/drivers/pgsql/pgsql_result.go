// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package pgsql


import (
	"database/sql"
	)
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
