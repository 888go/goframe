// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

import (
	"fmt"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

var (
	createdFieldNames = []string{"created_at", "create_at"} // 默认表字段名称，用于自动填充创建日期和时间。
	updatedFieldNames = []string{"updated_at", "update_at"} // 默认表字段名称，用于自动填充更新的日期和时间。
	deletedFieldNames = []string{"deleted_at", "delete_at"} // 默认表字段名称，用于自动填充删除时间戳。
)

// Unscoped 禁用在插入、更新和删除选项时自动更新时间的特性。
func (m *Model) Unscoped() *Model {
	model := m.getModel()
	model.unscoped = true
	return model
}

// getSoftFieldNameCreate 检查并返回记录创建时间的字段名称。
// 如果没有存储创建时间的字段名称，则返回一个空字符串。
// 它会检查包含或不包含大小写、字符 '-'/'_'/'.'/' ' 的键。
func (m *Model) getSoftFieldNameCreated(schema string, table string) string {
	// 它用于检查该特性是否已禁用。
	if m.db.GetConfig().TimeMaintainDisabled {
		return ""
	}
	tableName := ""
	if table != "" {
		tableName = table
	} else {
		tableName = m.tablesInit
	}
	config := m.db.GetConfig()
	if config.CreatedAt != "" {
		return m.getSoftFieldName(schema, tableName, []string{config.CreatedAt})
	}
	return m.getSoftFieldName(schema, tableName, createdFieldNames)
}

// getSoftFieldNameUpdate 检查并返回记录更新时间所对应的字段名称。
// 如果没有存储更新时间的字段名称，则返回一个空字符串。
// 它会检查包含或不包含大小写字符、'-'、'_'、'.'/' '等字符的关键字。
func (m *Model) getSoftFieldNameUpdated(schema string, table string) (field string) {
	// 它用于检查该特性是否已禁用。
	if m.db.GetConfig().TimeMaintainDisabled {
		return ""
	}
	tableName := ""
	if table != "" {
		tableName = table
	} else {
		tableName = m.tablesInit
	}
	config := m.db.GetConfig()
	if config.UpdatedAt != "" {
		return m.getSoftFieldName(schema, tableName, []string{config.UpdatedAt})
	}
	return m.getSoftFieldName(schema, tableName, updatedFieldNames)
}

// getSoftFieldNameDelete 检查并返回记录删除时间所使用的字段名。
// 如果没有存储删除时间的字段名，则返回一个空字符串。
// 它会检查包含或不包含大小写、字符 '-'/'_'/'.'/' ' 的键。
func (m *Model) getSoftFieldNameDeleted(schema string, table string) (field string) {
	// 它用于检查该特性是否已禁用。
	if m.db.GetConfig().TimeMaintainDisabled {
		return ""
	}
	tableName := ""
	if table != "" {
		tableName = table
	} else {
		tableName = m.tablesInit
	}
	config := m.db.GetConfig()
	if config.DeletedAt != "" {
		return m.getSoftFieldName(schema, tableName, []string{config.DeletedAt})
	}
	return m.getSoftFieldName(schema, tableName, deletedFieldNames)
}

// getSoftFieldName 获取并返回表中可能键的字段名称。
func (m *Model) getSoftFieldName(schema string, table string, keys []string) (field string) {
	// 忽略 TableFields 函数返回的错误。
	fieldsMap, _ := m.TableFields(table, schema)
	if len(fieldsMap) > 0 {
		for _, key := range keys {
			field, _ = gutil.MapPossibleItemByKey(
				gconv.Map(fieldsMap), key,
			)
			if field != "" {
				return
			}
		}
	}
	return
}

// getConditionForSoftDeleting 获取并返回用于软删除的条件字符串。
// 它支持多种表字符串，例如：
// "user u, user_detail ud" // 多个表别名定义
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid)" // 左连接查询语句
// "user LEFT JOIN user_detail ON(user_detail.uid=user.uid)" // 简化的左连接查询语句
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid) LEFT JOIN user_stats us ON(us.uid=u.uid)" // 多表左连接查询语句
// 该函数用于根据给定的多表查询条件，生成适用于软删除操作的SQL条件子句。
func (m *Model) getConditionForSoftDeleting() string {
	if m.unscoped {
		return ""
	}
	conditionArray := garray.NewStrArray()
	if gstr.Contains(m.tables, " JOIN ") {
		// Base table.
		match, _ := gregex.MatchString(`(.+?) [A-Z]+ JOIN`, m.tables)
		conditionArray.Append(m.getConditionOfTableStringForSoftDeleting(match[1]))
		// 多表连接，排除包含 '(' 和 ')' 字符的子查询SQL语句。
		matches, _ := gregex.MatchAllString(`JOIN ([^()]+?) ON`, m.tables)
		for _, match := range matches {
			conditionArray.Append(m.getConditionOfTableStringForSoftDeleting(match[1]))
		}
	}
	if conditionArray.Len() == 0 && gstr.Contains(m.tables, ",") {
		// 多个基础表。
		for _, s := range gstr.SplitAndTrim(m.tables, ",") {
			conditionArray.Append(m.getConditionOfTableStringForSoftDeleting(s))
		}
	}
	conditionArray.FilterEmpty()
	if conditionArray.Len() > 0 {
		return conditionArray.Join(" AND ")
	}
	// Only one table.
	if fieldName := m.getSoftFieldNameDeleted("", m.tablesInit); fieldName != "" {
		return fmt.Sprintf(`%s IS NULL`, m.db.GetCore().QuoteWord(fieldName))
	}
	return ""
}

// getConditionOfTableStringForSoftDeleting 函数的作用正如其名称所描述的那样。
// `s` 参数的例子：
// - `test`.`demo` as b （将`test`数据库中的`demo`表别名为b）
// - `test`.`demo` b （在`test`数据库中引用`demo`表，此处的 b 可能是别名或语法错误）
// - `demo` （假设是在当前默认数据库中引用`demo`表）
// - demo （与上例类似，直接引用`demo`表，未指定数据库）
func (m *Model) getConditionOfTableStringForSoftDeleting(s string) string {
	var (
		field  string
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
	field = m.getSoftFieldNameDeleted(schema, table)
	if field == "" {
		return ""
	}
	if len(array1) >= 3 {
		return fmt.Sprintf(`%s.%s IS NULL`, m.db.GetCore().QuoteWord(array1[2]), m.db.GetCore().QuoteWord(field))
	}
	if len(array1) >= 2 {
		return fmt.Sprintf(`%s.%s IS NULL`, m.db.GetCore().QuoteWord(array1[1]), m.db.GetCore().QuoteWord(field))
	}
	return fmt.Sprintf(`%s.%s IS NULL`, m.db.GetCore().QuoteWord(table), m.db.GetCore().QuoteWord(field))
}
