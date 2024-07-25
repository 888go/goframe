// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gdb

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// SoftTimeType自定义定义软时间字段类型。 md5:dac7cb3a21ca2d1d
type SoftTimeType int

const (
	SoftTimeTypeAuto           SoftTimeType = 0 // (默认)根据表字段类型自动检测字段类型。 md5:92b9309fac6f4d09
	SoftTimeTypeTime           SoftTimeType = 1 // 使用日期时间作为字段值。 md5:fb9be1cf84e4192e
	SoftTimeTypeTimestamp      SoftTimeType = 2 // In unix seconds.
	SoftTimeTypeTimestampMilli SoftTimeType = 3 // In unix milliseconds.
	SoftTimeTypeTimestampMicro SoftTimeType = 4 // In unix microseconds.
	SoftTimeTypeTimestampNano  SoftTimeType = 5 // In unix nanoseconds.
)

// SoftTimeOption 是用于自定义 Model 的软时间功能的选项。 md5:fcc19f5ef8ad45e7
type SoftTimeOption struct {
	SoftTimeType SoftTimeType // 软时间字段的值类型。 md5:472088e64d8a928f
}

type softTimeMaintainer struct {
	*Model
}

type iSoftTimeMaintainer interface {
	GetFieldNameAndTypeForCreate(
		ctx context.Context, schema string, table string,
	) (fieldName string, fieldType LocalType)

	GetFieldNameAndTypeForUpdate(
		ctx context.Context, schema string, table string,
	) (fieldName string, fieldType LocalType)

	GetFieldNameAndTypeForDelete(
		ctx context.Context, schema string, table string,
	) (fieldName string, fieldType LocalType)

	GetValueByFieldTypeForCreateOrUpdate(
		ctx context.Context, fieldType LocalType, isDeletedField bool,
	) (dataValue any)

	GetDataByFieldNameAndTypeForDelete(
		ctx context.Context, fieldPrefix, fieldName string, fieldType LocalType,
	) (dataHolder string, dataValue any)

	GetWhereConditionForDelete(ctx context.Context) string
}

// getSoftFieldNameAndTypeCacheItem 是用于存储创建/更新/删除字段的内部结构体。 md5:df4233b79c5f6dad
type getSoftFieldNameAndTypeCacheItem struct {
	FieldName string
	FieldType LocalType
}

var (
	// 当创建记录时，用于自动填充的表的默认字段名称。 md5:58c0524feef22203
	createdFieldNames = []string{"created_at", "create_at"}
	// 用于记录更新时自动填充的表默认字段名称。 md5:dfaf612ced6164b4
	updatedFieldNames = []string{"updated_at", "update_at"}
	// 默认的表字段名，用于自动填充记录删除。 md5:82caa57d9d8aac21
	deletedFieldNames = []string{"deleted_at", "delete_at"}
)

// SoftTime 设置 SoftTimeOption 以自定义 Model 的软时间功能。 md5:6c4368abcd89e6b0
func (m *Model) SoftTime(option SoftTimeOption) *Model {
	model := m.getModel()
	model.softTimeOption = option
	return model
}

// Unscoped禁用插入、更新和删除操作的软时间特性。 md5:0fc4af29459bd61e
func (m *Model) Unscoped() *Model {
	model := m.getModel()
	model.unscoped = true
	return model
}

func (m *Model) softTimeMaintainer() iSoftTimeMaintainer {
	return &softTimeMaintainer{
		m,
	}
}

// GetFieldNameAndTypeForCreate 检查并返回用于记录创建时间的字段名。
// 如果没有用于存储创建时间的字段名，它将返回一个空字符串。
// 它会检查键名，无论大小写或包含字符 '-'、'_'、'.'、' '。 md5:c03150380846ea77
func (m *softTimeMaintainer) GetFieldNameAndTypeForCreate(
	ctx context.Context, schema string, table string,
) (fieldName string, fieldType LocalType) {
	// 检查是否禁用了此功能。 md5:413ae315bebe927f
	if m.db.GetConfig().TimeMaintainDisabled {
		return "", LocalTypeUndefined
	}
	tableName := ""
	if table != "" {
		tableName = table
	} else {
		tableName = m.tablesInit
	}
	config := m.db.GetConfig()
	if config.CreatedAt != "" {
		return m.getSoftFieldNameAndType(
			ctx, schema, tableName, []string{config.CreatedAt},
		)
	}
	return m.getSoftFieldNameAndType(
		ctx, schema, tableName, createdFieldNames,
	)
}

// GetFieldNameAndTypeForUpdate 检查并返回用于更新时间的字段名。如果没有用于存储更新时间的字段名，它将返回空字符串。它会检查带有或不带大小写、字符 '-'/'_'/'.'/' 的键。 md5:220eb56737359035
func (m *softTimeMaintainer) GetFieldNameAndTypeForUpdate(
	ctx context.Context, schema string, table string,
) (fieldName string, fieldType LocalType) {
	// 检查是否禁用了此功能。 md5:413ae315bebe927f
	if m.db.GetConfig().TimeMaintainDisabled {
		return "", LocalTypeUndefined
	}
	tableName := ""
	if table != "" {
		tableName = table
	} else {
		tableName = m.tablesInit
	}
	config := m.db.GetConfig()
	if config.UpdatedAt != "" {
		return m.getSoftFieldNameAndType(
			ctx, schema, tableName, []string{config.UpdatedAt},
		)
	}
	return m.getSoftFieldNameAndType(
		ctx, schema, tableName, updatedFieldNames,
	)
}

// GetFieldNameAndTypeForDelete 检查并返回记录删除时间的字段名。如果没有用于存储删除时间的字段名，它将返回空字符串。它会检查大小写敏感或不敏感，以及使用 '-'、'_'、'.' 或 ' ' 作为分隔符的键。 md5:f7c6b45838b970b0
func (m *softTimeMaintainer) GetFieldNameAndTypeForDelete(
	ctx context.Context, schema string, table string,
) (fieldName string, fieldType LocalType) {
	// 检查是否禁用了此功能。 md5:413ae315bebe927f
	if m.db.GetConfig().TimeMaintainDisabled {
		return "", LocalTypeUndefined
	}
	tableName := ""
	if table != "" {
		tableName = table
	} else {
		tableName = m.tablesInit
	}
	config := m.db.GetConfig()
	if config.DeletedAt != "" {
		return m.getSoftFieldNameAndType(
			ctx, schema, tableName, []string{config.DeletedAt},
		)
	}
	return m.getSoftFieldNameAndType(
		ctx, schema, tableName, deletedFieldNames,
	)
}

// getSoftFieldName 获取并返回表中可能键的字段名。 md5:e32e19240070c456
func (m *softTimeMaintainer) getSoftFieldNameAndType(
	ctx context.Context,
	schema string, table string, checkFiledNames []string,
) (fieldName string, fieldType LocalType) {
	var (
		cacheKey      = fmt.Sprintf(`getSoftFieldNameAndType:%s#%s#%s`, schema, table, strings.Join(checkFiledNames, "_"))
		cacheDuration = gcache.DurationNoExpire
		cacheFunc     = func(ctx context.Context) (value interface{}, err error) {
			// 忽略TableFields函数的错误。 md5:b488d48f86ec5aea
			fieldsMap, _ := m.TableFields(table, schema)
			if len(fieldsMap) > 0 {
				for _, checkFiledName := range checkFiledNames {
					fieldName, _ = gutil.MapPossibleItemByKey(
						gconv.Map(fieldsMap), checkFiledName,
					)
					if fieldName != "" {
						fieldType, _ = m.db.CheckLocalTypeForField(
							ctx, fieldsMap[fieldName].Type, nil,
						)
						var cacheItem = getSoftFieldNameAndTypeCacheItem{
							FieldName: fieldName,
							FieldType: fieldType,
						}
						return cacheItem, nil
					}
				}
			}
			return
		}
	)
	result, err := gcache.GetOrSetFunc(ctx, cacheKey, cacheFunc, cacheDuration)
	if err != nil {
		intlog.Error(ctx, err)
	}
	if result != nil {
		var cacheItem getSoftFieldNameAndTypeCacheItem
		if err = result.Scan(&cacheItem); err != nil {
			return "", ""
		}
		fieldName = cacheItem.FieldName
		fieldType = cacheItem.FieldType
	}
	return
}

// GetWhereConditionForDelete 用于检索并返回软删除的条件字符串。它支持多表字符串，例如：
// "user u, user_detail ud" - "用户 u 和 user_detail ud"
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid)" - "用户 u 左连接 user_detail ud，连接条件为 ud.uid 等于 u.uid"
// "user LEFT JOIN user_detail ON(user_detail.uid=user.uid)" - "用户左连接 user_detail，连接条件为 user_detail.uid 等于 user.uid"
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid) LEFT JOIN user_stats us ON(us.uid=u.uid)" - "用户 u 先左连接 user_detail ud，再连接 user_stats us，连接条件为 us.uid 等于 u.uid" md5:f2c849c59f2ab188
func (m *softTimeMaintainer) GetWhereConditionForDelete(ctx context.Context) string {
	if m.unscoped {
		return ""
	}
	conditionArray := garray.NewStrArray()
	if gstr.Contains(m.tables, " JOIN ") {
		// Base table.
		tableMatch, _ := gregex.MatchString(`(.+?) [A-Z]+ JOIN`, m.tables)
		conditionArray.Append(m.getConditionOfTableStringForSoftDeleting(ctx, tableMatch[1]))
		// 多个连接的表，排除包含字符'('和')'的子查询SQL。 md5:a9edf50410c73b2c
		tableMatches, _ := gregex.MatchAllString(`JOIN ([^()]+?) ON`, m.tables)
		for _, match := range tableMatches {
			conditionArray.Append(m.getConditionOfTableStringForSoftDeleting(ctx, match[1]))
		}
	}
	if conditionArray.Len() == 0 && gstr.Contains(m.tables, ",") {
		// Multiple base tables.
		for _, s := range gstr.SplitAndTrim(m.tables, ",") {
			conditionArray.Append(m.getConditionOfTableStringForSoftDeleting(ctx, s))
		}
	}
	conditionArray.FilterEmpty()
	if conditionArray.Len() > 0 {
		return conditionArray.Join(" AND ")
	}
	// Only one table.
	fieldName, fieldType := m.GetFieldNameAndTypeForDelete(ctx, "", m.tablesInit)
	if fieldName != "" {
		return m.getConditionByFieldNameAndTypeForSoftDeleting(ctx, "", fieldName, fieldType)
	}
	return ""
}

// getConditionOfTableStringForSoftDeleting 的功能如其名称所述。
// `s` 的示例包括：
// - `test`.`demo` as b
// - `test`.`demo` b
// - `demo`
// - demo md5:ffb3e23129e1b6db
func (m *softTimeMaintainer) getConditionOfTableStringForSoftDeleting(ctx context.Context, s string) string {
	var (
		table  string
		schema string
		array1 = gstr.SplitAndTrim(s, " ")
		array2 = gstr.SplitAndTrim(array1[0], ".")
	)
	if len(array2) >= 2 {
		table = array2[1]
		schema = array2[0]
	} else {
		table = array2[0]
	}
	fieldName, fieldType := m.GetFieldNameAndTypeForDelete(ctx, schema, table)
	if fieldName == "" {
		return ""
	}
	if len(array1) >= 3 {
		return m.getConditionByFieldNameAndTypeForSoftDeleting(ctx, array1[2], fieldName, fieldType)
	}
	if len(array1) >= 2 {
		return m.getConditionByFieldNameAndTypeForSoftDeleting(ctx, array1[1], fieldName, fieldType)
	}
	return m.getConditionByFieldNameAndTypeForSoftDeleting(ctx, table, fieldName, fieldType)
}

// GetDataByFieldNameAndTypeForDelete 用于在软删除场景下，根据指定的字段名和类型创建并返回占位符和值。 md5:276be24343264681
func (m *softTimeMaintainer) GetDataByFieldNameAndTypeForDelete(
	ctx context.Context, fieldPrefix, fieldName string, fieldType LocalType,
) (dataHolder string, dataValue any) {
	var (
		quotedFieldPrefix = m.db.GetCore().QuoteWord(fieldPrefix)
		quotedFieldName   = m.db.GetCore().QuoteWord(fieldName)
	)
	if quotedFieldPrefix != "" {
		quotedFieldName = fmt.Sprintf(`%s.%s`, quotedFieldPrefix, quotedFieldName)
	}
	dataHolder = fmt.Sprintf(`%s=?`, quotedFieldName)
	dataValue = m.GetValueByFieldTypeForCreateOrUpdate(ctx, fieldType, false)
	return
}

func (m *softTimeMaintainer) getConditionByFieldNameAndTypeForSoftDeleting(
	ctx context.Context, fieldPrefix, fieldName string, fieldType LocalType,
) string {
	var (
		quotedFieldPrefix = m.db.GetCore().QuoteWord(fieldPrefix)
		quotedFieldName   = m.db.GetCore().QuoteWord(fieldName)
	)
	if quotedFieldPrefix != "" {
		quotedFieldName = fmt.Sprintf(`%s.%s`, quotedFieldPrefix, quotedFieldName)
	}
	switch m.softTimeOption.SoftTimeType {
	case SoftTimeTypeAuto:
		switch fieldType {
		case LocalTypeDate, LocalTypeDatetime:
			return fmt.Sprintf(`%s IS NULL`, quotedFieldName)
		case LocalTypeInt, LocalTypeUint, LocalTypeInt64, LocalTypeUint64, LocalTypeBool:
			return fmt.Sprintf(`%s=0`, quotedFieldName)
		default:
			intlog.Errorf(
				ctx,
				`invalid field type "%s" of field name "%s" with prefix "%s" for soft deleting condition`,
				fieldType, fieldName, fieldPrefix,
			)
		}

	case SoftTimeTypeTime:
		return fmt.Sprintf(`%s IS NULL`, quotedFieldName)

	default:
		return fmt.Sprintf(`%s=0`, quotedFieldName)
	}
	return ""
}

// GetValueByFieldTypeForCreateOrUpdate 为创建或更新操作创建并返回指定字段类型的值。 md5:263c89f2a7abf2da
func (m *softTimeMaintainer) GetValueByFieldTypeForCreateOrUpdate(
	ctx context.Context, fieldType LocalType, isDeletedField bool,
) any {
	var value any
	if isDeletedField {
		switch fieldType {
		case LocalTypeDate, LocalTypeDatetime:
			value = nil
		default:
			value = 0
		}
		return value
	}
	switch m.softTimeOption.SoftTimeType {
	case SoftTimeTypeAuto:
		switch fieldType {
		case LocalTypeDate, LocalTypeDatetime:
			value = gtime.Now()
		case LocalTypeInt, LocalTypeUint, LocalTypeInt64, LocalTypeUint64:
			value = gtime.Timestamp()
		case LocalTypeBool:
			value = 1
		default:
			intlog.Errorf(
				ctx,
				`invalid field type "%s" for soft deleting data`,
				fieldType,
			)
		}

	default:
		switch fieldType {
		case LocalTypeBool:
			value = 1
		default:
			value = m.createValueBySoftTimeOption(isDeletedField)
		}
	}
	return value
}

func (m *softTimeMaintainer) createValueBySoftTimeOption(isDeletedField bool) any {
	var value any
	if isDeletedField {
		switch m.softTimeOption.SoftTimeType {
		case SoftTimeTypeTime:
			value = nil
		default:
			value = 0
		}
		return value
	}
	switch m.softTimeOption.SoftTimeType {
	case SoftTimeTypeTime:
		value = gtime.Now()
	case SoftTimeTypeTimestamp:
		value = gtime.Timestamp()
	case SoftTimeTypeTimestampMilli:
		value = gtime.TimestampMilli()
	case SoftTimeTypeTimestampMicro:
		value = gtime.TimestampMicro()
	case SoftTimeTypeTimestampNano:
		value = gtime.TimestampNano()
	default:
		panic(gerror.NewCodef(
			gcode.CodeInternalPanic,
			`unrecognized SoftTimeType "%d"`, m.softTimeOption.SoftTimeType,
		))
	}
	return value
}
