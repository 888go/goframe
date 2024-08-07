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

	gdb "github.com/888go/goframe/database/gdb"
)

// X底层删除 为表执行 "DELETE FROM ... " 语句。 md5:48e885baa804ab97
func (d *Driver) X底层删除(ctx context.Context, link gdb.Link, table string, condition string, args ...interface{}) (result sql.Result, err error) {
	ctx = d.injectNeedParsedSql(ctx)
	return d.Core.X底层删除(ctx, link, table, condition, args...)
}
