// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql

import (
	"context"
	"fmt"

	gdb "github.com/888go/goframe/database/gdb"
	gutil "github.com/888go/goframe/util/gutil"
)

var (
	tableFieldsSqlByMariadb = `
SELECT
	c.COLUMN_NAME AS 'Field',
	( CASE WHEN ch.CHECK_CLAUSE LIKE 'json_valid%%' THEN 'json' ELSE c.COLUMN_TYPE END ) AS 'Type',
	c.COLLATION_NAME AS 'Collation',
	c.IS_NULLABLE AS 'Null',
	c.COLUMN_KEY AS 'Key',
	( CASE WHEN c.COLUMN_DEFAULT = 'NULL' OR c.COLUMN_DEFAULT IS NULL THEN NULL ELSE c.COLUMN_DEFAULT END) AS 'Default',
	c.EXTRA AS 'Extra',
	c.PRIVILEGES AS 'Privileges',
	c.COLUMN_COMMENT AS 'Comment' 
FROM
	information_schema.COLUMNS AS c
	LEFT JOIN information_schema.CHECK_CONSTRAINTS AS ch ON c.TABLE_NAME = ch.TABLE_NAME 
	AND c.COLUMN_NAME = ch.CONSTRAINT_NAME 
WHERE
	c.TABLE_SCHEMA = '%s' 
	AND c.TABLE_NAME = '%s'
	ORDER BY c.ORDINAL_POSITION`
)

func init() {
	var err error
	tableFieldsSqlByMariadb, err = gdb.FormatMultiLineSqlToSingle(tableFieldsSqlByMariadb)
	if err != nil {
		panic(err)
	}
}

// X取表字段信息Map 获取并返回当前模式指定表的字段信息。
// 
// 参数 `link` 是可选的，如果为 nil，则自动获取一个原始 SQL 连接，用于执行必要的 SQL 查询。
// 
// 它返回一个包含字段名及其对应字段的映射。由于映射是无序的，TableField 结构体有一个 "Index" 字段，标记其在字段中的顺序。
// 
// 该方法使用缓存功能来提高性能，直到进程重启，缓存永不过期。
// md5:38bed6cd2a065572
func (d *Driver) X取表字段信息Map(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result         gdb.Result
		link           gdb.Link
		usedSchema     = gutil.X取文本值或取默认值(d.X取默认数据库名称(), schema...)
		tableFieldsSql string
	)
	if link, err = d.X底层SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	dbType := d.X取当前节点配置().Type
	switch dbType {
	case "mariadb":
		tableFieldsSql = fmt.Sprintf(tableFieldsSqlByMariadb, usedSchema, table)
	default:
		tableFieldsSql = fmt.Sprintf(`SHOW FULL COLUMNS FROM %s`, d.X底层QuoteWord(table))
	}

	result, err = d.X底层查询(
		ctx, link,
		tableFieldsSql,
	)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	for i, m := range result {
		fields[m["Field"].String()] = &gdb.TableField{
			Index:   i,
			X名称:    m["Field"].String(),
			X类型:    m["Type"].String(),
			Null:    m["Null"].X取布尔(),
			Key:     m["Key"].String(),
			Default: m["Default"].X取值(),
			X额外:   m["Extra"].String(),
			Comment: m["Comment"].String(),
		}
	}
	return fields, nil
}
