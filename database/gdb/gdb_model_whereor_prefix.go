// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

// WhereOrPrefix 执行的功能与 WhereOr 相同，但会在 where 语句中的每个字段前添加指定的前缀。
// 请参阅 WhereBuilder.WhereOrPrefix。
func (m *Model) X条件或并带前缀(字段前缀 string, 条件 interface{}, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或并带前缀(字段前缀, 条件, 参数...))
}

// WhereOrPrefixLT 用于构建 `prefix.column < value` 形式的表达式，并将其以 `OR` 条件方式组合。 
// 详情请参阅 WhereBuilder.WhereOrPrefixLT 方法。
func (m *Model) X条件或小于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或小于并带前缀(字段前缀, 字段名, 比较值))
}

// WhereOrPrefixLTE 用于构建 `prefix.column <= value` 形式的 OR 条件语句。
// 请参阅 WhereBuilder.WhereOrPrefixLTE 方法。
func (m *Model) X条件或小于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或小于等于并带前缀(字段前缀, 字段名, 比较值))
}

// WhereOrPrefixGT 用于构建 `prefix.column > value` 形式的表达式，并将其以 `OR` 条件的方式加入到语句中。
// 详情请参阅 WhereBuilder.WhereOrPrefixGT。
func (m *Model) X条件或大于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或大于并带前缀(字段前缀, 字段名, 比较值))
}

// WhereOrPrefixGTE 用于构建 `prefix.column >= value` 形式的 OR 条件语句。
// 请参考 WhereBuilder.WhereOrPrefixGTE。
func (m *Model) X条件或大于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或大于等于并带前缀(字段前缀, 字段名, 比较值))
}

// WhereOrPrefixBetween 用于构建在“OR”条件中的 `prefix.column BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereOrPrefixBetween。
func (m *Model) X条件或取范围并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或取范围并带前缀(字段前缀, 字段名, 最小值, 最大值))
}

// WhereOrPrefixLike 在“OR”条件下构建 `prefix.column LIKE like` 语句。
// 参见 WhereBuilder.WhereOrPrefixLike。
func (m *Model) X条件或模糊匹配并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或模糊匹配并带前缀(字段前缀, 字段名, 通配符条件值))
}

// WhereOrPrefixIn 在`OR`条件中构建 `prefix.column IN (in)` 语句。
// 参见 WhereBuilder.WhereOrPrefixIn。
func (m *Model) X条件或包含并带前缀(字段前缀 string, 字段名 string, 包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或包含并带前缀(字段前缀, 字段名, 包含值))
}

// WhereOrPrefixNull 用于构建在`OR`条件中的 `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` 语句。
// 参见 WhereBuilder.WhereOrPrefixNull。
func (m *Model) X条件或NULL值并带前缀(字段前缀 string, 字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或NULL值并带前缀(字段前缀, 字段名...))
}

// WhereOrPrefixNotBetween 在“OR”条件下构建 `prefix.column NOT BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereOrPrefixNotBetween。
func (m *Model) X条件或取范围以外并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或取范围以外并带前缀(字段前缀, 字段名, 最小值, 最大值))
}

// WhereOrPrefixNotLike 用于构建 `prefix.column NOT LIKE like` 语句并在 `OR` 条件中使用。
// 参见 WhereBuilder.WhereOrPrefixNotLike。
func (m *Model) X条件或模糊匹配以外并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或模糊匹配以外并带前缀(字段前缀, 字段名, 通配符条件值))
}

// WhereOrPrefixNotIn 用于构建 `prefix.column NOT IN (in)` 语句。
// 请参阅 WhereBuilder.WhereOrPrefixNotIn。
func (m *Model) X条件或不包含并带前缀(字段前缀 string, 字段名 string, 不包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或不包含并带前缀(字段前缀, 字段名, 不包含值))
}

// WhereOrPrefixNotNull 用于构建在`OR`条件中的 `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` 语句。
// 参见 WhereBuilder.WhereOrPrefixNotNull。
func (m *Model) X条件或非Null并带前缀(字段前缀 string, 字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或非Null并带前缀(字段前缀, 字段名...))
}

// WhereOrPrefixNot 在`OR`条件中构建 `prefix.column != value` 语句。
// 参见 WhereBuilder.WhereOrPrefixNot。
func (m *Model) X条件或不等于并带前缀(字段前缀 string, 字段名 string, 值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或不等于并带前缀(字段前缀, 字段名, 值))
}
