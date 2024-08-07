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

	gdb "github.com/888go/goframe/database/gdb"
	gregex "github.com/888go/goframe/text/gregex"
	gutil "github.com/888go/goframe/util/gutil"
)

const (
	tableFieldsColumns = `name,position,default_expression,comment,type,is_in_partition_key,is_in_sorting_key,is_in_primary_key,is_in_sampling_key`
)

// X取表字段信息Map 获取并返回当前模式指定表的字段信息。也可以参考 DriverMysql.X取表字段信息Map。
// md5:2ca710808274dcba
func (d *Driver) X取表字段信息Map(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result     gdb.Result
		link       gdb.Link
		usedSchema = gutil.X取文本值或取默认值(d.X取默认数据库名称(), schema...)
	)
	if link, err = d.X底层SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	var (
		getColumnsSql = fmt.Sprintf(
			"select %s from `system`.columns c where `table` = '%s'",
			tableFieldsColumns, table,
		)
	)
	result, err = d.X底层查询(ctx, link, getColumnsSql)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	for _, m := range result {
		var (
			isNull    = false
			fieldType = m["type"].String()
		)
				// 在ClickHouse中，字段类型like是可空的整数（Nullable(int)）. md5:42a10ecf6628471b
		fieldsResult, _ := gregex.X匹配文本(`^Nullable\((.*?)\)`, fieldType)
		if len(fieldsResult) == 2 {
			isNull = true
			fieldType = fieldsResult[1]
		}
		position := m["position"].X取整数()
		if result[0]["position"].X取整数() != 0 {
			position -= 1
		}
		fields[m["name"].String()] = &gdb.TableField{
			Index:   position,
			X名称:    m["name"].String(),
			Default: m["default_expression"].X取值(),
			Comment: m["comment"].String(),
						// 键:     m["Key"] 的字符串表示,. md5:e3714fd2a741c0a1
			X类型: fieldType,
			Null: isNull,
		}
	}
	return fields, nil
}
