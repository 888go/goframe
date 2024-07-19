// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

// WhereOrPrefix performs as WhereOr, but it adds prefix to each field in where statement.
// See WhereBuilder.WhereOrPrefix.
// ff:条件或并带前缀
// m:
// prefix:字段前缀
// where:条件
// args:参数
func (m *Model) WhereOrPrefix(prefix string, where interface{}, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefix(prefix, where, args...))
}

// WhereOrPrefixLT builds `prefix.column < value` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixLT.
// ff:条件或小于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:比较值
func (m *Model) WhereOrPrefixLT(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixLT(prefix, column, value))
}

// WhereOrPrefixLTE builds `prefix.column <= value` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixLTE.
// ff:条件或小于等于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:比较值
func (m *Model) WhereOrPrefixLTE(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixLTE(prefix, column, value))
}

// WhereOrPrefixGT builds `prefix.column > value` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixGT.
// ff:条件或大于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:比较值
func (m *Model) WhereOrPrefixGT(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixGT(prefix, column, value))
}

// WhereOrPrefixGTE builds `prefix.column >= value` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixGTE.
// ff:条件或大于等于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:比较值
func (m *Model) WhereOrPrefixGTE(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixGTE(prefix, column, value))
}

// WhereOrPrefixBetween builds `prefix.column BETWEEN min AND max` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixBetween.
// ff:条件或取范围并带前缀
// m:
// prefix:字段前缀
// column:字段名
// min:最小值
// max:最大值
func (m *Model) WhereOrPrefixBetween(prefix string, column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixBetween(prefix, column, min, max))
}

// WhereOrPrefixLike builds `prefix.column LIKE like` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixLike.
// ff:条件或模糊匹配并带前缀
// m:
// prefix:字段前缀
// column:字段名
// like:通配符条件值
func (m *Model) WhereOrPrefixLike(prefix string, column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixLike(prefix, column, like))
}

// WhereOrPrefixIn builds `prefix.column IN (in)` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixIn.
// ff:条件或包含并带前缀
// m:
// prefix:字段前缀
// column:字段名
// in:包含值
func (m *Model) WhereOrPrefixIn(prefix string, column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixIn(prefix, column, in))
}

// WhereOrPrefixNull builds `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixNull.
// ff:条件或NULL值并带前缀
// m:
// prefix:字段前缀
// columns:字段名
func (m *Model) WhereOrPrefixNull(prefix string, columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNull(prefix, columns...))
}

// WhereOrPrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixNotBetween.
// ff:条件或取范围以外并带前缀
// m:
// prefix:字段前缀
// column:字段名
// min:最小值
// max:最大值
func (m *Model) WhereOrPrefixNotBetween(prefix string, column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNotBetween(prefix, column, min, max))
}

// WhereOrPrefixNotLike builds `prefix.column NOT LIKE like` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixNotLike.
// ff:条件或模糊匹配以外并带前缀
// m:
// prefix:字段前缀
// column:字段名
// like:通配符条件值
func (m *Model) WhereOrPrefixNotLike(prefix string, column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNotLike(prefix, column, like))
}

// WhereOrPrefixNotIn builds `prefix.column NOT IN (in)` statement.
// See WhereBuilder.WhereOrPrefixNotIn.
// ff:条件或不包含并带前缀
// m:
// prefix:字段前缀
// column:字段名
// in:不包含值
func (m *Model) WhereOrPrefixNotIn(prefix string, column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNotIn(prefix, column, in))
}

// WhereOrPrefixNotNull builds `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixNotNull.
// ff:条件或非Null并带前缀
// m:
// prefix:字段前缀
// columns:字段名
func (m *Model) WhereOrPrefixNotNull(prefix string, columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNotNull(prefix, columns...))
}

// WhereOrPrefixNot builds `prefix.column != value` statement in `OR` conditions.
// See WhereBuilder.WhereOrPrefixNot.
// ff:条件或不等于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:值
func (m *Model) WhereOrPrefixNot(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNot(prefix, column, value))
}
