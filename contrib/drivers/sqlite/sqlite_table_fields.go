// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package sqlite

import (
	"context"
	"fmt"

	gdb "github.com/888go/goframe/database/gdb"
	gutil "github.com/888go/goframe/util/gutil"
)

// X取表字段信息Map 获取并返回当前模式下指定表的字段信息。
//
// 参见 DriverMysql.X取表字段信息Map。
// md5:7f7a75c67e38ad22
func (d *Driver) X取表字段信息Map(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result     gdb.Result
		link       gdb.Link
		usedSchema = gutil.X取文本值或取默认值(d.X取默认数据库名称(), schema...)
	)
	if link, err = d.X底层SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	result, err = d.X底层查询(ctx, link, fmt.Sprintf(`PRAGMA TABLE_INFO(%s)`, d.X底层QuoteWord(table)))
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	for i, m := range result {
		mKey := ""
		if m["pk"].X取布尔() {
			mKey = "pri"
		}
		fields[m["name"].String()] = &gdb.TableField{
			Index:   i,
			X名称:    m["name"].String(),
			X类型:    m["type"].String(),
			Key:     mKey,
			Default: m["dflt_value"].X取值(),
			Null:    !m["notnull"].X取布尔(),
		}
	}
	return fields, nil
}
