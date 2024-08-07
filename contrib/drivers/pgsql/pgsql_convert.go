// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package pgsql

import (
	"context"
	"strings"

	gdb "github.com/888go/goframe/database/gdb"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// X底层CheckLocalTypeForField 检查并返回给定数据库类型对应的本地Go语言类型。 md5:f8aef7c5b09aa9c8
func (d *Driver) X底层CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (gdb.LocalType, error) {
	var typeName string
	match, _ := gregex.X匹配文本(`(.+?)\((.+)\)`, fieldType)
	if len(match) == 3 {
		typeName = gstr.X过滤首尾符并含空白(match[1])
	} else {
		typeName = fieldType
	}
	typeName = strings.ToLower(typeName)
	switch typeName {
	case
				// 对于pgsql，int2等于smallint。 md5:5d75ff3e1cf74f36
		"int2",
						// 对于 PostgreSQL，int4 表示整数. md5:0e9fb5268eeec552
		"int4":
		return gdb.LocalTypeInt, nil

	case
			// 对于 PostgreSQL，int8 等同于 bigint. md5:4717ef91027dfe75.
		"int8":
		return gdb.LocalTypeInt64, nil

	case
		"_int2",
		"_int4":
		return gdb.LocalTypeIntSlice, nil

	case
		"_int8":
		return gdb.LocalTypeInt64Slice, nil

	default:
		return d.Core.X底层CheckLocalTypeForField(ctx, fieldType, fieldValue)
	}
}

// X底层ConvertValueForLocal 根据从数据库中获取的字段类型名称，将值转换为Go语言中的本地类型。
// 参数 `fieldType` 为小写格式，例如：
// `float(5,2)`，`unsigned double(5,2)`，`decimal(10,2)`，`char(45)`，`varchar(100)` 等。
// md5:7e1ede2b68158e31
func (d *Driver) X底层ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) {
	typeName, _ := gregex.X替换文本(`\(.+\)`, "", fieldType)
	typeName = strings.ToLower(typeName)
	switch typeName {
		// 对于pgsql，int2等于smallint，int4等于integer。 md5:9a03a0c9b626da62
	case "int2", "int4":
		return gconv.X取整数(gconv.String(fieldValue)), nil

		// 对于 PostgreSQL，int8 等同于 bigint. md5:4717ef91027dfe75..
	case "int8":
		return gconv.X取整数64位(gconv.String(fieldValue)), nil

	// Int32 slice.
	case
		"_int2", "_int4":
		return gconv.X取整数切片(
			gstr.Map替换(gconv.String(fieldValue),
				map[string]string{
					"{": "[",
					"}": "]",
				},
			),
		), nil

	// Int64 slice.
	case
		"_int8":
		return gconv.X取整数64位切片(
			gstr.Map替换(gconv.String(fieldValue),
				map[string]string{
					"{": "[",
					"}": "]",
				},
			),
		), nil

	default:
		return d.Core.X底层ConvertValueForLocal(ctx, fieldType, fieldValue)
	}
}
