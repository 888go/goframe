// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package clickhouse

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	gdb "github.com/888go/goframe/database/gdb"
)

// DoInsert 为给定的表插入或更新数据。 md5:2a62d01f344269b8
func (d *Driver) DoInsert(
	ctx context.Context, link gdb.Link, table string, list gdb.List, option gdb.DoInsertOption,
) (result sql.Result, err error) {
	var (
		keys        []string // Field names.
		valueHolder = make([]string, 0)
	)
		// 处理字段名和占位符。 md5:a4c2e01bfbec2f37
	for k := range list[0] {
		keys = append(keys, k)
		valueHolder = append(valueHolder, "?")
	}
		// 准备批量结果指针。 md5:dfc8aa8bb292f9d5
	var (
		charL, charR = d.Core.GetChars()
		keysStr      = charL + strings.Join(keys, charR+","+charL) + charR
		holderStr    = strings.Join(valueHolder, ",")
		tx           gdb.TX
		stmt         *gdb.Stmt
	)
	tx, err = d.Core.Begin(ctx)
	if err != nil {
		return
	}
		// 使用`defer`确保事务将被提交或回滚。 md5:f7e6a525062b3162
	defer func() {
		if err == nil {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}()
	stmt, err = tx.Prepare(fmt.Sprintf(
		"INSERT INTO %s(%s) VALUES (%s)",
		d.QuotePrefixTableName(table), keysStr,
		holderStr,
	))
	if err != nil {
		return
	}
	for i := 0; i < len(list); i++ {
				// 将被提交给底层数据库驱动程序的值。 md5:d30c8d96f40663c3
		params := make([]interface{}, 0)
		for _, k := range keys {
			params = append(params, list[i][k])
		}
				// Prepare 只允许在由 clickhouse 打开的事务中执行一次. md5:b763067296709df3
		result, err = stmt.ExecContext(ctx, params...)
		if err != nil {
			return
		}
	}
	return
}
