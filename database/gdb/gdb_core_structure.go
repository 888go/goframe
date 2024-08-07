// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"context"
	"database/sql/driver"
	"reflect"
	"strings"
	"time"

	gbinary "github.com/888go/goframe/encoding/gbinary"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	gtime "github.com/888go/goframe/os/gtime"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

// X取字段类型 通过名称检索并返回指定字段的字段类型字符串。 md5:aeb8d310c854c45a
func (c *Core) X取字段类型(上下文 context.Context, 字段名称, 表名称, 数据库名称 string) string {
	field := c.X取字段信息对象(上下文, 字段名称, 表名称, 数据库名称)
	if field != nil {
		return field.X类型
	}
	return ""
}

// X取字段信息对象 通过字段名获取并返回该字段的类型对象。 md5:eeebff59dbaf1064
func (c *Core) X取字段信息对象(上下文 context.Context, 字段名称, 表名称, 数据库名称 string) *TableField {
	fieldsMap, err := c.db.X取表字段信息Map(上下文, 表名称, 数据库名称)
	if err != nil {
		intlog.Errorf(
			上下文,
			`TableFields failed for table "%s", schema "%s": %+v`,
			表名称, 数据库名称, err,
		)
		return nil
	}
	for tableFieldName, tableField := range fieldsMap {
		if tableFieldName == 字段名称 {
			return tableField
		}
	}
	return nil
}

// X底层ConvertDataForRecord 是一个非常重要的函数，用于将任何数据转换为
// 以便将其作为记录插入到表或集合中。
//
// 参数 `value` 应为 *map/map/*struct/struct 类型。
// 对于结构体，它支持嵌入式结构体定义。
// md5:27b867ec3a1c3c1d
func (c *Core) X底层ConvertDataForRecord(上下文 context.Context, 值 interface{}, 表名称 string) (map[string]interface{}, error) {
	var (
		err  error
		data = X转换到Map(值, true)
	)
	for fieldName, fieldValue := range data {
		data[fieldName], err = c.db.X底层ConvertValueForField(
			上下文,
			c.X取字段类型(上下文, fieldName, 表名称, c.X取默认数据库名称()),
			fieldValue,
		)
		if err != nil {
			return nil, gerror.X多层错误并格式化(err, `ConvertDataForRecord failed for value: %#v`, fieldValue)
		}
	}
	return data, nil
}

// X底层ConvertValueForField 将值转换为记录字段的类型。
// 参数 `fieldType` 是目标记录字段。
// 参数 `fieldValue` 是要写入记录字段的值。
// md5:196c02c9f6cf3380
func (c *Core) X底层ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) {
	var (
		err            error
		convertedValue = fieldValue
	)
		// 如果`value`实现了`driver.Valuer`接口，那么它会使用该接口进行值的转换。 md5:ba72a317b79988a2
	if valuer, ok := fieldValue.(driver.Valuer); ok {
		if convertedValue, err = valuer.Value(); err != nil {
			if err != nil {
				return nil, err
			}
		}
		return convertedValue, nil
	}
		// 默认值转换。 md5:99c30401ccc62832
	var (
		rvValue = reflect.ValueOf(fieldValue)
		rvKind  = rvValue.Kind()
	)
	for rvKind == reflect.Ptr {
		rvValue = rvValue.Elem()
		rvKind = rvValue.Kind()
	}
	switch rvKind {
	case reflect.Slice, reflect.Array, reflect.Map:
				// 它应该忽略字节类型。 md5:48e040f1decb1a5e
		if _, ok := fieldValue.([]byte); !ok {
							// 将值转换为JSON。 md5:f4977a0ae972c910
			convertedValue, err = json.Marshal(fieldValue)
			if err != nil {
				return nil, err
			}
		}

	case reflect.Struct:
		switch r := fieldValue.(type) {
		// 如果时间是零值，它将更新为nil，
		// 这样在数据库中插入或更新的值将会是"null"。
		// md5:058aebae61025f37
		case time.Time:
			if r.IsZero() {
				convertedValue = nil
			}

		case gtime.Time:
			if r.IsZero() {
				convertedValue = nil
			} else {
				convertedValue = r.Time
			}

		case *gtime.Time:
			if r.IsZero() {
				convertedValue = nil
			} else {
				convertedValue = r.Time
			}

		case *time.Time:
			// Nothing to do.

		case Counter, *Counter:
			// Nothing to do.

		default:
			// 如果`value`实现了iNil接口，
			// 检查其IsNil()函数，如果返回true，
			// 将把该值插入/更新到数据库中作为"null"。
			// md5:b2415061d93829e6
			if v, ok := fieldValue.(iNil); ok && v.X是否为Nil() {
				convertedValue = nil
			} else if s, ok := fieldValue.(iString); ok {
								// 默认使用字符串转换。 md5:36cba4c54f848f87
				convertedValue = s.String()
			} else {
								// 将值转换为JSON。 md5:f4977a0ae972c910
				convertedValue, err = json.Marshal(fieldValue)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return convertedValue, nil
}

// X底层CheckLocalTypeForField 检查并返回与给定数据库类型相对应的本地类型。 md5:d3191e6393b7e531
func (c *Core) X底层CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (LocalType, error) {
	var (
		typeName    string
		typePattern string
	)
	match, _ := gregex.X匹配文本(`(.+?)\((.+)\)`, fieldType)
	if len(match) == 3 {
		typeName = gstr.X过滤首尾符并含空白(match[1])
		typePattern = gstr.X过滤首尾符并含空白(match[2])
	} else {
		typeName = gstr.X分割(fieldType, " ")[0]
	}

	typeName = strings.ToLower(typeName)

	switch typeName {
	case
		fieldTypeBinary,
		fieldTypeVarbinary,
		fieldTypeBlob,
		fieldTypeTinyblob,
		fieldTypeMediumblob,
		fieldTypeLongblob:
		return LocalTypeBytes, nil

	case
		fieldTypeInt,
		fieldTypeTinyint,
		fieldTypeSmallInt,
		fieldTypeSmallint,
		fieldTypeMediumInt,
		fieldTypeMediumint,
		fieldTypeSerial:
		if gstr.X是否包含并忽略大小写(fieldType, "unsigned") {
			return LocalTypeUint, nil
		}
		return LocalTypeInt, nil

	case
		fieldTypeBigInt,
		fieldTypeBigint,
		fieldTypeBigserial:
		if gstr.X是否包含并忽略大小写(fieldType, "unsigned") {
			return LocalTypeUint64, nil
		}
		return LocalTypeInt64, nil

	case
		fieldTypeReal:
		return LocalTypeFloat32, nil

	case
		fieldTypeDecimal,
		fieldTypeMoney,
		fieldTypeNumeric,
		fieldTypeSmallmoney:
		return LocalTypeString, nil
	case
		fieldTypeFloat,
		fieldTypeDouble:
		return LocalTypeFloat64, nil

	case
		fieldTypeBit:
				// 建议使用 bit(1) 作为布尔值。 md5:5be00c9e8395ea93
		if typePattern == "1" {
			return LocalTypeBool, nil
		}
		s := gconv.String(fieldValue)
				// mssql 是一个true|false类型的字符串。 md5:6d4dbdb95d9adfa1
		if strings.EqualFold(s, "true") || strings.EqualFold(s, "false") {
			return LocalTypeBool, nil
		}
		if gstr.X是否包含并忽略大小写(fieldType, "unsigned") {
			return LocalTypeUint64Bytes, nil
		}
		return LocalTypeInt64Bytes, nil

	case
		fieldTypeBool:
		return LocalTypeBool, nil

	case
		fieldTypeDate:
		return LocalTypeDate, nil

	case
		fieldTypeDatetime,
		fieldTypeTimestamp,
		fieldTypeTimestampz:
		return LocalTypeDatetime, nil

	case
		fieldTypeJson:
		return LocalTypeJson, nil

	case
		fieldTypeJsonb:
		return LocalTypeJsonb, nil

	default:
				// 自动检测字段类型，通过键匹配。 md5:138e4aeac8d26d8a
		switch {
		case strings.Contains(typeName, "text") || strings.Contains(typeName, "char") || strings.Contains(typeName, "character"):
			return LocalTypeString, nil

		case strings.Contains(typeName, "float") || strings.Contains(typeName, "double") || strings.Contains(typeName, "numeric"):
			return LocalTypeFloat64, nil

		case strings.Contains(typeName, "bool"):
			return LocalTypeBool, nil

		case strings.Contains(typeName, "binary") || strings.Contains(typeName, "blob"):
			return LocalTypeBytes, nil

		case strings.Contains(typeName, "int"):
			if gstr.X是否包含并忽略大小写(fieldType, "unsigned") {
				return LocalTypeUint, nil
			}
			return LocalTypeInt, nil

		case strings.Contains(typeName, "time"):
			return LocalTypeDatetime, nil

		case strings.Contains(typeName, "date"):
			return LocalTypeDatetime, nil

		default:
			return LocalTypeString, nil
		}
	}
}

// X底层ConvertValueForLocal 根据从数据库中获取的字段类型名称，将值转换为Go语言中的本地类型。
// 参数 `fieldType` 为小写格式，例如：
// `float(5,2)`，`unsigned double(5,2)`，`decimal(10,2)`，`char(45)`，`varchar(100)` 等。
// md5:7e1ede2b68158e31
func (c *Core) X底层ConvertValueForLocal(
	ctx context.Context, fieldType string, fieldValue interface{},
) (interface{}, error) {
	// 如果没有获取到类型，则直接返回`fieldValue`，
	// 利用其原始数据类型，因为`fieldValue`是`interface{}`类型的。
	// md5:62cf4d391c9da4f2
	if fieldType == "" {
		return fieldValue, nil
	}
	typeName, err := c.db.X底层CheckLocalTypeForField(ctx, fieldType, fieldValue)
	if err != nil {
		return nil, err
	}
	switch typeName {
	case LocalTypeBytes:
		var typeNameStr = string(typeName)
		if strings.Contains(typeNameStr, "binary") || strings.Contains(typeNameStr, "blob") {
			return fieldValue, nil
		}
		return gconv.X取字节集(fieldValue), nil

	case LocalTypeInt:
		return gconv.X取整数(gconv.String(fieldValue)), nil

	case LocalTypeUint:
		return gconv.X取正整数(gconv.String(fieldValue)), nil

	case LocalTypeInt64:
		return gconv.X取整数64位(gconv.String(fieldValue)), nil

	case LocalTypeUint64:
		return gconv.X取正整数64位(gconv.String(fieldValue)), nil

	case LocalTypeInt64Bytes:
		return gbinary.BeDecodeToInt64(gconv.X取字节集(fieldValue)), nil

	case LocalTypeUint64Bytes:
		return gbinary.BeDecodeToUint64(gconv.X取字节集(fieldValue)), nil

	case LocalTypeFloat32:
		return gconv.X取小数32位(gconv.String(fieldValue)), nil

	case LocalTypeFloat64:
		return gconv.X取小数64位(gconv.String(fieldValue)), nil

	case LocalTypeBool:
		s := gconv.String(fieldValue)
				// mssql 是一个true|false类型的字符串。 md5:6d4dbdb95d9adfa1
		if strings.EqualFold(s, "true") {
			return 1, nil
		}
		if strings.EqualFold(s, "false") {
			return 0, nil
		}
		return gconv.X取布尔(fieldValue), nil

	case LocalTypeDate:
		// Date without time.
		if t, ok := fieldValue.(time.Time); ok {
			return gtime.X创建并按Time(t).X取格式文本("Y-m-d"), nil
		}
		t, _ := gtime.X转换文本(gconv.String(fieldValue))
		return t.X取格式文本("Y-m-d"), nil

	case LocalTypeDatetime:
		if t, ok := fieldValue.(time.Time); ok {
			return gtime.X创建并按Time(t), nil
		}
		t, _ := gtime.X转换文本(gconv.String(fieldValue))
		return t, nil

	default:
		return gconv.String(fieldValue), nil
	}
}

// mappingAndFilterData 自动将映射键映射到表格字段，并删除所有不是给定表格字段的键值对。
// md5:27fc8e27d4d4a389
func (c *Core) mappingAndFilterData(ctx context.Context, schema, table string, data map[string]interface{}, filter bool) (map[string]interface{}, error) {
	fieldsMap, err := c.db.X取表字段信息Map(ctx, c.guessPrimaryTableName(table), schema)
	if err != nil {
		return nil, err
	}
	fieldsKeyMap := make(map[string]interface{}, len(fieldsMap))
	for k := range fieldsMap {
		fieldsKeyMap[k] = nil
	}
		// 自动将数据键映射到表格字段名。 md5:bdc9aa8a688bb975
	var foundKey string
	for dataKey, dataValue := range data {
		if _, ok := fieldsKeyMap[dataKey]; !ok {
			foundKey, _ = gutil.MapPossibleItemByKey(fieldsKeyMap, dataKey)
			if foundKey != "" {
				if _, ok = data[foundKey]; !ok {
					data[foundKey] = dataValue
				}
				delete(data, dataKey)
			}
		}
	}
	// 数据过滤。
	// 它会删除所有具有错误字段名的键值对。
	// md5:24aafcb1699db80c
	if filter {
		for dataKey := range data {
			if _, ok := fieldsMap[dataKey]; !ok {
				delete(data, dataKey)
			}
		}
	}
	return data, nil
}
