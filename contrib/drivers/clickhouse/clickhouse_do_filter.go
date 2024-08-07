// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package clickhouse

import (
	"context"
	"fmt"
	"strings"

	gdb "github.com/888go/goframe/database/gdb"
	gregex "github.com/888go/goframe/text/gregex"
)

// X底层DoFilter 在将SQL提交到数据库之前进行处理。 md5:e56455a7432db765
func (d *Driver) X底层DoFilter(
	ctx context.Context, link gdb.Link, originSql string, args []interface{},
) (newSql string, newArgs []interface{}, err error) {
	if len(args) == 0 {
		return originSql, args, nil
	}
		// 将占位符字符'?'转换为字符串 "$x"。 md5:a1e39f745b49128a
	var index int
	originSql, _ = gregex.X替换文本_函数(`\?`, originSql, func(s string) string {
		index++
		return fmt.Sprintf(`$%d`, index)
	})

		// 只有通过框架生成的SQL才会被处理。 md5:fe793f1543ad8481
	if !d.getNeedParsedSqlFromCtx(ctx) {
		return originSql, args, nil
	}

			// 将标准SQL替换为ClickHouse SQL语法. md5:eb710cc2dce6880f
	modeRes, err := gregex.X匹配文本(filterTypePattern, strings.TrimSpace(originSql))
	if err != nil {
		return "", nil, err
	}
	if len(modeRes) == 0 {
		return originSql, args, nil
	}

		// 只有删除/更新语句需要过滤条件. md5:a43c6e48c79fc525
	switch strings.ToUpper(modeRes[0]) {
	case "UPDATE":
		// MySQL 示例：UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE 条件]
		// Clickhouse 示例：ALTER TABLE [db.]table UPDATE column1 = expr1[, ...] WHERE filter_expr
		// 
		// 这段代码是针对两种数据库系统的更新语句的注释。在MySQL中，`UPDATE` 用于更新表中的数据，设置指定字段的新值，并可选地使用 `WHERE` 子句来限制更新的行。在Clickhouse中，`ALTER TABLE` 用于更新表中的列，将列的值设置为表达式（expr1），并且需要一个过滤表达式（filter_expr）来确定哪些行会被更新。其中 `[db.]table` 表示可以包含数据库名的表名。
		// md5:d201a8d0c4df9319
		newSql, err = gregex.ReplaceStringFuncMatch(
			updateFilterPattern, originSql,
			func(s []string) string {
				return fmt.Sprintf("ALTER TABLE %s UPDATE", s[1])
			},
		)
		if err != nil {
			return "", nil, err
		}
		return newSql, args, nil

	case "DELETE":
		// MySQL 示例：DELETE FROM 表名 [WHERE 子句]
		// Clickhouse 示例：ALTER TABLE [db.]表名 [ON CLUSTER 集群名称] DELETE WHERE 过滤表达式
		// md5:b83ff5334ed70fb6
		newSql, err = gregex.ReplaceStringFuncMatch(
			deleteFilterPattern, originSql,
			func(s []string) string {
				return fmt.Sprintf("ALTER TABLE %s DELETE", s[1])
			},
		)
		if err != nil {
			return "", nil, err
		}
		return newSql, args, nil

	}
	return originSql, args, nil
}

func (d *Driver) getNeedParsedSqlFromCtx(ctx context.Context) bool {
	if ctx.Value(needParsedSqlInCtx) != nil {
		return true
	}
	return false
}
