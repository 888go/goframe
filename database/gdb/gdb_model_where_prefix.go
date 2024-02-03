// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

// WherePrefix 的行为类似于 Where，但它会在 where 语句中的每个字段前添加前缀。
// 请参阅 WhereBuilder.WherePrefix。
func (m *Model) WherePrefix(prefix string, where interface{}, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefix(prefix, where, args...))
}

// WherePrefixLT 用于构建 `prefix.column < value` 语句。
// 请参阅 WhereBuilder.WherePrefixLT。
func (m *Model) WherePrefixLT(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixLT(prefix, column, value))
}

// WherePrefixLTE 用于构建 `prefix.column <= value` 条件语句。
// 参见 WhereBuilder.WherePrefixLTE。
func (m *Model) WherePrefixLTE(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixLTE(prefix, column, value))
}

// WherePrefixGT 用于构建 `prefix.column > value` 语句。
// 参见 WhereBuilder.WherePrefixGT。
func (m *Model) WherePrefixGT(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixGT(prefix, column, value))
}

// WherePrefixGTE 用于构建 `prefix.column >= value` 条件语句。
// 请参阅 WhereBuilder.WherePrefixGTE。
func (m *Model) WherePrefixGTE(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixGTE(prefix, column, value))
}

// WherePrefixBetween 用于构建 `prefix.column BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WherePrefixBetween。
func (m *Model) WherePrefixBetween(prefix string, column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixBetween(prefix, column, min, max))
}

// WherePrefixLike 用于构建 `prefix.column LIKE like` 语句。
// 参见 WhereBuilder.WherePrefixLike。
func (m *Model) WherePrefixLike(prefix string, column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixLike(prefix, column, like))
}

// WherePrefixIn 用于构建 `prefix.column IN (in)` 语句。
// 参见 WhereBuilder.WherePrefixIn。
func (m *Model) WherePrefixIn(prefix string, column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixIn(prefix, column, in))
}

// WherePrefixNull 用于构建 `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` 语句。
// 参见 WhereBuilder.WherePrefixNull。
func (m *Model) WherePrefixNull(prefix string, columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixNull(prefix, columns...))
}

// WherePrefixNotBetween 用于构建 `prefix.column NOT BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WherePrefixNotBetween。
func (m *Model) WherePrefixNotBetween(prefix string, column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixNotBetween(prefix, column, min, max))
}

// WherePrefixNotLike 用于构建 `prefix.column NOT LIKE like` 语句。
// 参见 WhereBuilder.WherePrefixNotLike。
func (m *Model) WherePrefixNotLike(prefix string, column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixNotLike(prefix, column, like))
}

// WherePrefixNot 用于构建 `prefix.column != value` 的语句。
// 参见 WhereBuilder.WherePrefixNot。
func (m *Model) WherePrefixNot(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixNot(prefix, column, value))
}

// WherePrefixNotIn 用于构建 `prefix.column NOT IN (in)` 语句。
// 请参阅 WhereBuilder.WherePrefixNotIn。
func (m *Model) WherePrefixNotIn(prefix string, column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixNotIn(prefix, column, in))
}

// WherePrefixNotNull 用于构建 `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` 形式的语句。
// 请参阅 WhereBuilder.WherePrefixNotNull。
func (m *Model) WherePrefixNotNull(prefix string, columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixNotNull(prefix, columns...))
}
