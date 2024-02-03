// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

// WhereOr 向 WHERE 语句添加“OR”条件。
// 请参阅 WhereBuilder.WhereOr。
func (m *Model) WhereOr(where interface{}, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOr(where, args...))
}

// WhereOrf 通过 fmt.Sprintf 和参数构建 `OR` 条件字符串。 
// 参见 WhereBuilder.WhereOrf 。
func (m *Model) WhereOrf(format string, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrf(format, args...))
}

// WhereOrLT 用于构建在“OR”条件中的 `column < value` 语句。
// 请参阅 WhereBuilder.WhereOrLT。
func (m *Model) WhereOrLT(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrLT(column, value))
}

// WhereOrLTE 用于构建 `column <= value` 条件语句，并以 `OR` 连接。 // 详情请参考 WhereBuilder.WhereOrLTE 。
func (m *Model) WhereOrLTE(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrLTE(column, value))
}

// WhereOrGT 用于构建 `column > value` 条件语句，并以 `OR` 连接。 
// 详情参见 WhereBuilder.WhereOrGT。
func (m *Model) WhereOrGT(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrGT(column, value))
}

// WhereOrGTE 用于构建在“OR”条件中的`column >= value`语句。
// 参见WhereBuilder.WhereOrGTE。
func (m *Model) WhereOrGTE(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrGTE(column, value))
}

// WhereOrBetween 在“OR”条件下构建 `column BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereOrBetween。
func (m *Model) WhereOrBetween(column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrBetween(column, min, max))
}

// WhereOrLike 用于构建 `column LIKE like` 语句并在 `OR` 条件中使用。
// 参考 WhereBuilder.WhereOrLike。
func (m *Model) WhereOrLike(column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrLike(column, like))
}

// WhereOrIn 在`OR`条件下构建`column IN (in)`语句。 
// 参见WhereBuilder.WhereOrIn方法。
func (m *Model) WhereOrIn(column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrIn(column, in))
}

// WhereOrNull 用于构建以 `OR` 条件连接的 `columns[0] IS NULL OR columns[1] IS NULL ...` 语句。
// 请参阅 WhereBuilder.WhereOrNull。
func (m *Model) WhereOrNull(columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNull(columns...))
}

// WhereOrNotBetween 用于构建 `column NOT BETWEEN min AND max` 语句，并将其以 `OR` 条件形式加入到查询中。
// 参考 WhereBuilder.WhereOrNotBetween。
func (m *Model) WhereOrNotBetween(column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNotBetween(column, min, max))
}

// WhereOrNotLike 用于构建在“OR”条件中的`column NOT LIKE 'like'`语句。
// 参见 WhereBuilder.WhereOrNotLike。
func (m *Model) WhereOrNotLike(column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNotLike(column, like))
}

// WhereOrNot 用于构建 `column != value` 的语句。
// 参见 WhereBuilder.WhereOrNot。
func (m *Model) WhereOrNot(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNot(column, value))
}

// WhereOrNotIn 用于构建 `column NOT IN (in)` 语句。
// 参见 WhereBuilder.WhereOrNotIn。
func (m *Model) WhereOrNotIn(column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNotIn(column, in))
}

// WhereOrNotNull 用于构建在`OR`条件中的`columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...`语句。
// 参见WhereBuilder.WhereOrNotNull方法。
func (m *Model) WhereOrNotNull(columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNotNull(columns...))
}
