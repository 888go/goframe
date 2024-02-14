// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"context"
	"database/sql/driver"
	"reflect"
	"strings"
	"time"
	
	"github.com/888go/goframe/encoding/gbinary"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

// GetFieldTypeStr 通过名称检索并返回特定字段的字段类型字符串。
func (c *Core) X取字段类型(上下文 context.Context, 字段名称, 表名称, 数据库名称 string) string {
	field := c.X取字段信息对象(上下文, 字段名称, 表名称, 数据库名称)
	if field != nil {
		return field.Type
	}
	return ""
}

// GetFieldType通过名称获取并返回特定字段的字段类型对象。
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

// ConvertDataForRecord 是一个非常重要的函数，用于将任何要作为记录插入到表/集合中的数据进行转换。
//
// 参数 `value` 应为 *map、map、*struct 或 struct 类型。对于 struct，它支持嵌套的 struct 定义。
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
			return nil, 错误类.X多层错误并格式化(err, `ConvertDataForRecord failed for value: %#v`, fieldValue)
		}
	}
	return data, nil
}

// ConvertValueForField 将值转换为目标记录字段的类型。
// 参数`fieldType`是目标记录字段。
// 参数`fieldValue`是要提交到记录字段的值。
func (c *Core) X底层ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) {
	var (
		err            error
		convertedValue = fieldValue
	)
	// 如果`value`实现了接口`driver.Valuer`，那么它将使用该接口进行值转换。
	if valuer, ok := fieldValue.(driver.Valuer); ok {
		if convertedValue, err = valuer.Value(); err != nil {
			if err != nil {
				return nil, err
			}
		}
		return convertedValue, nil
	}
	// 默认值转换
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
		// 它应当忽略 bytes 类型。
		if _, ok := fieldValue.([]byte); !ok {
			// 将值转换为JSON。
			convertedValue, err = json.Marshal(fieldValue)
			if err != nil {
				return nil, err
			}
		}

	case reflect.Struct:
		switch r := fieldValue.(type) {
// 如果时间是零值，那么将其更新为nil，
// 这将会把该值以"null"的形式插入/更新到数据库中。
		case time.Time:
			if r.IsZero() {
				convertedValue = nil
			}

		case 时间类.Time:
			if r.IsZero() {
				convertedValue = nil
			} else {
				convertedValue = r.Time
			}

		case *时间类.Time:
			if r.IsZero() {
				convertedValue = nil
			} else {
				convertedValue = r.Time
			}

		case *time.Time:
			// Nothing to do.

		case X增减, *X增减:
			// Nothing to do.

		default:
// 如果`value`实现了接口iNil，
// 检查其IsNil()函数，如果得到true，
// 则将该值作为"null"插入/更新到数据库中。
// 在Go语言中，这段代码的注释描述了当变量value实现了名为iNil的接口时，会进一步调用其IsNil()方法进行判断。若该方法返回true，则会在对数据库进行操作时，将这个变量值视为"null"进行插入或更新操作。
			if v, ok := fieldValue.(iNil); ok && v.X是否为Nil() {
				convertedValue = nil
			} else if s, ok := fieldValue.(iString); ok {
				// 默认情况下使用字符串转换
				convertedValue = s.String()
			} else {
				// 将值转换为JSON。
				convertedValue, err = json.Marshal(fieldValue)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return convertedValue, nil
}

// CheckLocalTypeForField 检查并返回给定数据库类型对应的本地类型。
func (c *Core) X底层CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (LocalType, error) {
	var (
		typeName    string
		typePattern string
	)
	match, _ := 正则类.X匹配文本(`(.+?)\((.+)\)`, fieldType)
	if len(match) == 3 {
		typeName = 文本类.X过滤首尾符并含空白(match[1])
		typePattern = 文本类.X过滤首尾符并含空白(match[2])
	} else {
		typeName = 文本类.X分割(fieldType, " ")[0]
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
		if 文本类.X是否包含并忽略大小写(fieldType, "unsigned") {
			return LocalTypeUint, nil
		}
		return LocalTypeInt, nil

	case
		fieldTypeBigInt,
		fieldTypeBigint,
		fieldTypeBigserial:
		if 文本类.X是否包含并忽略大小写(fieldType, "unsigned") {
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
		// 建议使用bit(1)作为布尔值。
		if typePattern == "1" {
			return LocalTypeBool, nil
		}
		s := 转换类.String(fieldValue)
		// mssql 是一个表示真或假的字符串。
		if strings.EqualFold(s, "true") || strings.EqualFold(s, "false") {
			return LocalTypeBool, nil
		}
		if 文本类.X是否包含并忽略大小写(fieldType, "unsigned") {
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
		// 自动检测字段类型，通过键匹配实现。
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
			if 文本类.X是否包含并忽略大小写(fieldType, "unsigned") {
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

// ConvertValueForLocal 将值根据数据库字段类型名称转换为本地 Golang 类型的值。
// 参数 `fieldType` 为小写形式，例如：
// `float(5,2)`、`unsigned double(5,2)`、`decimal(10,2)`、`char(45)`、`varchar(100)` 等。
func (c *Core) X底层ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) {
// 如果没有检索到类型，它会直接返回 `fieldValue`，使用其原始数据类型，
// 因为 `fieldValue` 是 interface{} 类型。
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
		return 转换类.X取字节集(fieldValue), nil

	case LocalTypeInt:
		return 转换类.X取整数(转换类.String(fieldValue)), nil

	case LocalTypeUint:
		return 转换类.X取正整数(转换类.String(fieldValue)), nil

	case LocalTypeInt64:
		return 转换类.X取整数64位(转换类.String(fieldValue)), nil

	case LocalTypeUint64:
		return 转换类.X取正整数64位(转换类.String(fieldValue)), nil

	case LocalTypeInt64Bytes:
		return 字节集类.BeDecodeToInt64(转换类.X取字节集(fieldValue)), nil

	case LocalTypeUint64Bytes:
		return 字节集类.BeDecodeToUint64(转换类.X取字节集(fieldValue)), nil

	case LocalTypeFloat32:
		return 转换类.X取小数32位(转换类.String(fieldValue)), nil

	case LocalTypeFloat64:
		return 转换类.X取小数64位(转换类.String(fieldValue)), nil

	case LocalTypeBool:
		s := 转换类.String(fieldValue)
		// mssql 是一个表示真或假的字符串。
		if strings.EqualFold(s, "true") {
			return 1, nil
		}
		if strings.EqualFold(s, "false") {
			return 0, nil
		}
		return 转换类.X取布尔(fieldValue), nil

	case LocalTypeDate:
		// 仅日期，不含时间。
		if t, ok := fieldValue.(time.Time); ok {
			return 时间类.X创建并按Time(t).X取格式文本("Y-m-d"), nil
		}
		t, _ := 时间类.X转换文本(转换类.String(fieldValue))
		return t.X取格式文本("Y-m-d"), nil

	case LocalTypeDatetime:
		if t, ok := fieldValue.(time.Time); ok {
			return 时间类.X创建并按Time(t), nil
		}
		t, _ := 时间类.X转换文本(转换类.String(fieldValue))
		return t, nil

	default:
		return 转换类.String(fieldValue), nil
	}
}

// mappingAndFilterData 自动将映射键映射到表字段，并移除所有非给定表字段的键值对。
func (c *Core) mappingAndFilterData(ctx context.Context, schema, table string, data map[string]interface{}, filter bool) (map[string]interface{}, error) {
	fieldsMap, err := c.db.X取表字段信息Map(ctx, c.guessPrimaryTableName(table), schema)
	if err != nil {
		return nil, err
	}
	fieldsKeyMap := make(map[string]interface{}, len(fieldsMap))
	for k := range fieldsMap {
		fieldsKeyMap[k] = nil
	}
	// 自动将数据键映射到表字段名称。
	var foundKey string
	for dataKey, dataValue := range data {
		if _, ok := fieldsKeyMap[dataKey]; !ok {
			foundKey, _ = 工具类.MapPossibleItemByKey(fieldsKeyMap, dataKey)
			if foundKey != "" {
				if _, ok = data[foundKey]; !ok {
					data[foundKey] = dataValue
				}
				delete(data, dataKey)
			}
		}
	}
// 数据过滤。
// 删除所有具有错误字段名的键值对。
	if filter {
		for dataKey := range data {
			if _, ok := fieldsMap[dataKey]; !ok {
				delete(data, dataKey)
			}
		}
	}
	return data, nil
}
