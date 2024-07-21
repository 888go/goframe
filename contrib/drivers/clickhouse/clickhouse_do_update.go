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

	"github.com/gogf/gf/v2/database/gdb"
)

// DoUpdate 对该表执行 "UPDATE ... " 语句。 md5:d99c5c0beb2de28e
// ff:
// d:
// ctx:
// link:
// table:
// data:
// condition:
// args:
// result:
// err:
func (d *Driver) DoUpdate(ctx context.Context, link gdb.Link, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error) {
	ctx = d.injectNeedParsedSql(ctx)
	return d.Core.DoUpdate(ctx, link, table, data, condition, args...)
}
