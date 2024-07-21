// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package sqlite

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
)

// DoFilter 在将 SQL 字符串提交给底层 SQL 驱动程序之前处理它。 md5:f9ff7431f1478cfb
// ff:
// d:
// ctx:
// link:
// sql:
// args:
// newSql:
// newArgs:
// err:
func (d *Driver) DoFilter(ctx context.Context, link gdb.Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	// 针对sqlite的特殊插入/忽略操作。 md5:7cfea509103b3cc2
	switch {
	case gstr.HasPrefix(sql, gdb.InsertOperationIgnore):
		sql = "INSERT OR IGNORE" + sql[len(gdb.InsertOperationIgnore):]

	case gstr.HasPrefix(sql, gdb.InsertOperationReplace):
		sql = "INSERT OR REPLACE" + sql[len(gdb.InsertOperationReplace):]
	}
	return d.Core.DoFilter(ctx, link, sql, args)
}
