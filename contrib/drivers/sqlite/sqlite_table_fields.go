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

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gutil"
)

	// TableFields 获取并返回当前模式下指定表的字段信息。
	//
	// 参见 DriverMysql.TableFields。
	// md5:7f7a75c67e38ad22
// ff:
// d:
// ctx:
// table:
// schema:
// fields:
// err:
func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result     gdb.Result
		link       gdb.Link
		usedSchema = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	result, err = d.DoSelect(ctx, link, fmt.Sprintf(`PRAGMA TABLE_INFO(%s)`, d.QuoteWord(table)))
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	for i, m := range result {
		mKey := ""
		if m["pk"].Bool() {
			mKey = "pri"
		}
		fields[m["name"].String()] = &gdb.TableField{
			Index:   i,
			Name:    m["name"].String(),
			Type:    m["type"].String(),
			Key:     mKey,
			Default: m["dflt_value"].Val(),
			Null:    !m["notnull"].Bool(),
		}
	}
	return fields, nil
}
