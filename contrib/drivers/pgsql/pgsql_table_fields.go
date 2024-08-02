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

	gdb "github.com/888go/goframe/database/gdb"
	gutil "github.com/888go/goframe/util/gutil"
)

var (
	tableFieldsSqlTmp = `
SELECT a.attname AS field, t.typname AS type,a.attnotnull as null,
    (case when d.contype = 'p' then 'pri' when d.contype = 'u' then 'uni' else '' end)  as key
      ,ic.column_default as default_value,b.description as comment
      ,coalesce(character_maximum_length, numeric_precision, -1) as length
      ,numeric_scale as scale
FROM pg_attribute a
         left join pg_class c on a.attrelid = c.oid
         left join pg_constraint d on d.conrelid = c.oid and a.attnum = d.conkey[1]
         left join pg_description b ON a.attrelid=b.objoid AND a.attnum = b.objsubid
         left join pg_type t ON a.atttypid = t.oid
         left join information_schema.columns ic on ic.column_name = a.attname and ic.table_name = c.relname
WHERE c.relname = '%s' and a.attisdropped is false and a.attnum > 0
ORDER BY a.attnum`
)

func init() {
	var err error
	tableFieldsSqlTmp, err = gdb.FormatMultiLineSqlToSingle(tableFieldsSqlTmp)
	if err != nil {
		panic(err)
	}
}

// TableFields 获取并返回当前模式中指定表的字段信息。 md5:920febaff284f5e7
func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result     gdb.Result
		link       gdb.Link
		usedSchema = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
				// TODO：id结果重复了吗？. md5:1a107e32ccbc7928
		structureSql = fmt.Sprintf(tableFieldsSqlTmp, table)
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	result, err = d.DoSelect(ctx, link, structureSql)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	var (
		index = 0
		name  string
		ok    bool
	)
	for _, m := range result {
		name = m["field"].String()
				// 过滤重复的字段。 md5:e2695e711875512a
		if _, ok = fields[name]; ok {
			continue
		}
		fields[name] = &gdb.TableField{
			Index:   index,
			Name:    name,
			Type:    m["type"].String(),
			Null:    !m["null"].Bool(),
			Key:     m["key"].String(),
			Default: m["default_value"].Val(),
			Comment: m["comment"].String(),
		}
		index++
	}
	return fields, nil
}
