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
	"regexp"

	gdb "github.com/888go/goframe/database/gdb"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gutil "github.com/888go/goframe/util/gutil"
)

var (
	tablesSqlTmp = `
SELECT
	c.relname
FROM
	pg_class c
INNER JOIN pg_namespace n ON
	c.relnamespace = n.oid
WHERE
	n.nspname = '%s'
	AND c.relkind IN ('r', 'p')
	%s
ORDER BY
	c.relname
`
)

func init() {
	var err error
	tablesSqlTmp, err = gdb.FormatMultiLineSqlToSingle(tablesSqlTmp)
	if err != nil {
		panic(err)
	}
}

// X取表名称切片 获取并返回当前模式下的表格列表。
//主要用于命令行工具链，用于自动生成模型。
// md5:bce161ba95454bf5
func (d *Driver) X取表名称切片(ctx context.Context, schema ...string) (tables []string, err error) {
	var (
		result     gdb.Result
		usedSchema = gutil.X取文本值或取默认值(d.X取当前节点配置().Namespace, schema...)
	)
	if usedSchema == "" {
		usedSchema = defaultSchema
	}
		// 不要将`usedSchema`作为`SlaveLink`函数的参数。 md5:283541defa4ac558
	link, err := d.X底层SlaveLink(schema...)
	if err != nil {
		return nil, err
	}

	useRelpartbound := ""
	if gstr.X版本号比较GNU格式(d.version(ctx, link), "10") >= 0 {
		useRelpartbound = "AND c.relpartbound IS NULL"
	}

	var query = fmt.Sprintf(
		tablesSqlTmp,
		usedSchema,
		useRelpartbound,
	)

	query, _ = gregex.X替换文本(`[\n\r\s]+`, " ", gstr.X过滤首尾符并含空白(query))
	result, err = d.X底层查询(ctx, link, query)
	if err != nil {
		return
	}
	for _, m := range result {
		for _, v := range m {
			tables = append(tables, v.String())
		}
	}
	return
}

// 检查并返回数据库版本。 md5:39cd1f37b14f728a
func (d *Driver) version(ctx context.Context, link gdb.Link) string {
	result, err := d.X底层查询(ctx, link, "SELECT version();")
	if err != nil {
		return ""
	}
	if len(result) > 0 {
		if v, ok := result[0]["version"]; ok {
			matches := regexp.MustCompile(`PostgreSQL (\d+\.\d+)`).FindStringSubmatch(v.String())
			if len(matches) >= 2 {
				return matches[1]
			}
		}
	}
	return ""
}
