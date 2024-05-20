// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package pgsql

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gregex"
)

// DoFilter 在将 SQL 字符串提交给底层 SQL 驱动程序之前处理它。. md5:f9ff7431f1478cfb
func (d *Driver) DoFilter(
	ctx context.Context, link gdb.Link, sql string, args []interface{},
) (newSql string, newArgs []interface{}, err error) {
	var index int
	// 将占位符字符'?'转换为字符串 "$x"。. md5:a1e39f745b49128a
	newSql, err = gregex.ReplaceStringFunc(`\?`, sql, func(s string) string {
		index++
		return fmt.Sprintf(`$%d`, index)
	})
	if err != nil {
		return "", nil, err
	}
// 处理pgsql的jsonb功能支持，其中包含占位符字符'?'。
// 参考：
// https://github.com/gogf/gf/issues/1537
// https://www.postgresql.org/docs/12/functions-json.html
// md5:49874022abc6a281
	newSql, err = gregex.ReplaceStringFuncMatch(`(::jsonb([^\w\d]*)\$\d)`, newSql, func(match []string) string {
		return fmt.Sprintf(`::jsonb%s?`, match[2])
	})
	if err != nil {
		return "", nil, err
	}
	newSql, err = gregex.ReplaceString(` LIMIT (\d+),\s*(\d+)`, ` LIMIT $2 OFFSET $1`, newSql)
	if err != nil {
		return "", nil, err
	}
	newArgs = args
	return d.Core.DoFilter(ctx, link, newSql, newArgs)
}
