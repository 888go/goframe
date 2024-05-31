// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

// WhereOr adds "OR" condition to the where statement.
// See WhereBuilder.WhereOr.

// ff:条件或
// args:参数
// where:条件
func (m *Model) WhereOr(where interface{}, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOr(where, args...))
}

// WhereOrf builds `OR` condition string using fmt.Sprintf and arguments.
// See WhereBuilder.WhereOrf.

// ff:条件或格式化
// args:参数
// format:格式
func (m *Model) WhereOrf(format string, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrf(format, args...))
}

// WhereOrLT builds `column < value` statement in `OR` conditions.
// See WhereBuilder.WhereOrLT.

// ff:条件或小于
// value:比较值
// column:字段名
func (m *Model) WhereOrLT(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrLT(column, value))
}

// WhereOrLTE builds `column <= value` statement in `OR` conditions.
// See WhereBuilder.WhereOrLTE.

// ff:条件或小于等于
// value:比较值
// column:字段名
func (m *Model) WhereOrLTE(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrLTE(column, value))
}

// WhereOrGT builds `column > value` statement in `OR` conditions.
// See WhereBuilder.WhereOrGT.

// ff:条件或大于
// value:比较值
// column:字段名
func (m *Model) WhereOrGT(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrGT(column, value))
}

// WhereOrGTE builds `column >= value` statement in `OR` conditions.
// See WhereBuilder.WhereOrGTE.

// ff:条件或大于等于
// value:比较值
// column:字段名
func (m *Model) WhereOrGTE(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrGTE(column, value))
}

// WhereOrBetween builds `column BETWEEN min AND max` statement in `OR` conditions.
// See WhereBuilder.WhereOrBetween.

// ff:条件或取范围
// max:最大值
// min:最小值
// column:字段名
func (m *Model) WhereOrBetween(column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrBetween(column, min, max))
}

// WhereOrLike builds `column LIKE like` statement in `OR` conditions.
// See WhereBuilder.WhereOrLike.

// ff:条件或模糊匹配
// like:通配符条件值
// column:字段名
func (m *Model) WhereOrLike(column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrLike(column, like))
}

// WhereOrIn builds `column IN (in)` statement in `OR` conditions.
// See WhereBuilder.WhereOrIn.

// ff:条件或包含
// in:包含值
// column:字段名
func (m *Model) WhereOrIn(column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrIn(column, in))
}

// WhereOrNull builds `columns[0] IS NULL OR columns[1] IS NULL ...` statement in `OR` conditions.
// See WhereBuilder.WhereOrNull.

// ff:条件或NULL值
// columns:字段名
func (m *Model) WhereOrNull(columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNull(columns...))
}

// WhereOrNotBetween builds `column NOT BETWEEN min AND max` statement in `OR` conditions.
// See WhereBuilder.WhereOrNotBetween.

// ff:条件或取范围以外
// max:最大值
// min:最小值
// column:字段名
func (m *Model) WhereOrNotBetween(column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNotBetween(column, min, max))
}

// WhereOrNotLike builds `column NOT LIKE 'like'` statement in `OR` conditions.
// See WhereBuilder.WhereOrNotLike.

// ff:条件或模糊匹配以外
// like:通配符条件值
// column:字段名
func (m *Model) WhereOrNotLike(column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNotLike(column, like))
}

// WhereOrNot builds `column != value` statement.
// See WhereBuilder.WhereOrNot.

// ff:条件或不等于
// value:值
// column:字段名
func (m *Model) WhereOrNot(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNot(column, value))
}

// WhereOrNotIn builds `column NOT IN (in)` statement.
// See WhereBuilder.WhereOrNotIn.

// ff:条件或不包含
// in:不包含值
// column:字段名
func (m *Model) WhereOrNotIn(column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNotIn(column, in))
}

// WhereOrNotNull builds `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` statement in `OR` conditions.
// See WhereBuilder.WhereOrNotNull.

// ff:条件或非Null
// columns:字段名
func (m *Model) WhereOrNotNull(columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNotNull(columns...))
}
