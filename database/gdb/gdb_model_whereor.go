// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

// WhereOr 向 WHERE 语句添加“OR”条件。
// 请参阅 WhereBuilder.WhereOr。
func (m *Model) X条件或(条件 interface{}, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或(条件, 参数...))
}

// WhereOrf 通过 fmt.Sprintf 和参数构建 `OR` 条件字符串。 
// 参见 WhereBuilder.WhereOrf 。
func (m *Model) X条件或格式化(格式 string, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或格式化(格式, 参数...))
}

// WhereOrLT 用于构建在“OR”条件中的 `column < value` 语句。
// 请参阅 WhereBuilder.WhereOrLT。
func (m *Model) X条件或小于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或小于(字段名, 比较值))
}

// WhereOrLTE 用于构建 `column <= value` 条件语句，并以 `OR` 连接。 // 详情请参考 WhereBuilder.WhereOrLTE 。
func (m *Model) X条件或小于等于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或小于等于(字段名, 比较值))
}

// WhereOrGT 用于构建 `column > value` 条件语句，并以 `OR` 连接。 
// 详情参见 WhereBuilder.WhereOrGT。
func (m *Model) X条件或大于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或大于(字段名, 比较值))
}

// WhereOrGTE 用于构建在“OR”条件中的`column >= value`语句。
// 参见WhereBuilder.WhereOrGTE。
func (m *Model) X条件或大于等于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或大于等于(字段名, 比较值))
}

// WhereOrBetween 在“OR”条件下构建 `column BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereOrBetween。
func (m *Model) X条件或取范围(字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或取范围(字段名, 最小值, 最大值))
}

// WhereOrLike 用于构建 `column LIKE like` 语句并在 `OR` 条件中使用。
// 参考 WhereBuilder.WhereOrLike。
func (m *Model) X条件或模糊匹配(字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或模糊匹配(字段名, 通配符条件值))
}

// WhereOrIn 在`OR`条件下构建`column IN (in)`语句。 
// 参见WhereBuilder.WhereOrIn方法。
func (m *Model) X条件或包含(字段名 string, 包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或包含(字段名, 包含值))
}

// WhereOrNull 用于构建以 `OR` 条件连接的 `columns[0] IS NULL OR columns[1] IS NULL ...` 语句。
// 请参阅 WhereBuilder.WhereOrNull。
func (m *Model) X条件或NULL值(字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或NULL值(字段名...))
}

// WhereOrNotBetween 用于构建 `column NOT BETWEEN min AND max` 语句，并将其以 `OR` 条件形式加入到查询中。
// 参考 WhereBuilder.WhereOrNotBetween。
func (m *Model) X条件或取范围以外(字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或取范围以外(字段名, 最小值, 最大值))
}

// WhereOrNotLike 用于构建在“OR”条件中的`column NOT LIKE 'like'`语句。
// 参见 WhereBuilder.WhereOrNotLike。
func (m *Model) X条件或模糊匹配以外(字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或模糊匹配以外(字段名, 通配符条件值))
}

// WhereOrNot 用于构建 `column != value` 的语句。
// 参见 WhereBuilder.WhereOrNot。
func (m *Model) X条件或不等于(字段名 string, 值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或不等于(字段名, 值))
}

// WhereOrNotIn 用于构建 `column NOT IN (in)` 语句。
// 参见 WhereBuilder.WhereOrNotIn。
func (m *Model) X条件或不包含(字段名 string, 不包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或不包含(字段名, 不包含值))
}

// WhereOrNotNull 用于构建在`OR`条件中的`columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...`语句。
// 参见WhereBuilder.WhereOrNotNull方法。
func (m *Model) X条件或非Null(字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或非Null(字段名...))
}
