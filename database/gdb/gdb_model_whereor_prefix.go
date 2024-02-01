// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

// WhereOrPrefix 执行的功能与 WhereOr 相同，但会在 where 语句中的每个字段前添加指定的前缀。
// 请参阅 WhereBuilder.WhereOrPrefix。
func (m *Model) WhereOrPrefix(prefix string, where interface{}, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefix(prefix, where, args...))
}

// WhereOrPrefixLT 用于构建 `prefix.column < value` 形式的表达式，并将其以 `OR` 条件方式组合。 
// 详情请参阅 WhereBuilder.WhereOrPrefixLT 方法。
func (m *Model) WhereOrPrefixLT(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixLT(prefix, column, value))
}

// WhereOrPrefixLTE 用于构建 `prefix.column <= value` 形式的 OR 条件语句。
// 请参阅 WhereBuilder.WhereOrPrefixLTE 方法。
func (m *Model) WhereOrPrefixLTE(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixLTE(prefix, column, value))
}

// WhereOrPrefixGT 用于构建 `prefix.column > value` 形式的表达式，并将其以 `OR` 条件的方式加入到语句中。
// 详情请参阅 WhereBuilder.WhereOrPrefixGT。
func (m *Model) WhereOrPrefixGT(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixGT(prefix, column, value))
}

// WhereOrPrefixGTE 用于构建 `prefix.column >= value` 形式的 OR 条件语句。
// 请参考 WhereBuilder.WhereOrPrefixGTE。
func (m *Model) WhereOrPrefixGTE(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixGTE(prefix, column, value))
}

// WhereOrPrefixBetween 用于构建在“OR”条件中的 `prefix.column BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereOrPrefixBetween。
func (m *Model) WhereOrPrefixBetween(prefix string, column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixBetween(prefix, column, min, max))
}

// WhereOrPrefixLike 在“OR”条件下构建 `prefix.column LIKE like` 语句。
// 参见 WhereBuilder.WhereOrPrefixLike。
func (m *Model) WhereOrPrefixLike(prefix string, column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixLike(prefix, column, like))
}

// WhereOrPrefixIn 在`OR`条件中构建 `prefix.column IN (in)` 语句。
// 参见 WhereBuilder.WhereOrPrefixIn。
func (m *Model) WhereOrPrefixIn(prefix string, column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixIn(prefix, column, in))
}

// WhereOrPrefixNull 用于构建在`OR`条件中的 `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` 语句。
// 参见 WhereBuilder.WhereOrPrefixNull。
func (m *Model) WhereOrPrefixNull(prefix string, columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNull(prefix, columns...))
}

// WhereOrPrefixNotBetween 在“OR”条件下构建 `prefix.column NOT BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereOrPrefixNotBetween。
func (m *Model) WhereOrPrefixNotBetween(prefix string, column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNotBetween(prefix, column, min, max))
}

// WhereOrPrefixNotLike 用于构建 `prefix.column NOT LIKE like` 语句并在 `OR` 条件中使用。
// 参见 WhereBuilder.WhereOrPrefixNotLike。
func (m *Model) WhereOrPrefixNotLike(prefix string, column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNotLike(prefix, column, like))
}

// WhereOrPrefixNotIn 用于构建 `prefix.column NOT IN (in)` 语句。
// 请参阅 WhereBuilder.WhereOrPrefixNotIn。
func (m *Model) WhereOrPrefixNotIn(prefix string, column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNotIn(prefix, column, in))
}

// WhereOrPrefixNotNull 用于构建在`OR`条件中的 `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` 语句。
// 参见 WhereBuilder.WhereOrPrefixNotNull。
func (m *Model) WhereOrPrefixNotNull(prefix string, columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNotNull(prefix, columns...))
}

// WhereOrPrefixNot 在`OR`条件中构建 `prefix.column != value` 语句。
// 参见 WhereBuilder.WhereOrPrefixNot。
func (m *Model) WhereOrPrefixNot(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNot(prefix, column, value))
}
