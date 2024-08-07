// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package pgsql

import (
	"context"
	"database/sql"

	gdb "github.com/888go/goframe/database/gdb"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

// X底层插入 为给定的表插入或更新数据。 md5:2a62d01f344269b8
func (d *Driver) X底层插入(ctx context.Context, link gdb.Link, table string, list gdb.Map切片, option gdb.DoInsertOption) (result sql.Result, err error) {
	switch option.InsertOption {
	case gdb.InsertOptionReplace:
		return nil, gerror.X创建错误码(
			gcode.CodeNotSupported,
			`Replace operation is not supported by pgsql driver`,
		)

	case gdb.InsertOptionIgnore:
		return nil, gerror.X创建错误码(
			gcode.CodeNotSupported,
			`Insert ignore operation is not supported by pgsql driver`,
		)

	case gdb.InsertOptionDefault:
		tableFields, err := d.X取Core对象().X取DB对象().X取表字段信息Map(ctx, table)
		if err == nil {
			for _, field := range tableFields {
				if field.Key == "pri" {
					pkField := *field
					ctx = context.WithValue(ctx, internalPrimaryKeyInCtx, pkField)
					break
				}
			}
		}
	}
	return d.Core.X底层插入(ctx, link, table, list, option)
}
