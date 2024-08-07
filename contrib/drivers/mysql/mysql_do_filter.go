// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql

import (
	"context"

	gdb "github.com/888go/goframe/database/gdb"
)

// X底层DoFilter 在将SQL提交到数据库之前进行处理。 md5:e56455a7432db765
func (d *Driver) X底层DoFilter(ctx context.Context, link gdb.Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	return d.Core.X底层DoFilter(ctx, link, sql, args)
}
