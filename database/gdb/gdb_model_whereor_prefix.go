// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

// WhereOrPrefix 的行为与 WhereOr 相似，但它会在where语句中的每个字段前添加前缀。
// 参考 WhereBuilder.WhereOrPrefix。
// md5:4ea5d18d5615ff17
// ff:条件或并带前缀
// m:
// prefix:字段前缀
// where:条件
// args:参数
func (m *Model) WhereOrPrefix(prefix string, where interface{}, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefix(prefix, where, args...))
}

// WhereOrPrefixLT在"OR"条件下构建`prefix.column < value`语句。参阅WhereBuilder.WhereOrPrefixLT。
// md5:0a8c07ff239fa7e7
// ff:条件或小于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:比较值
func (m *Model) WhereOrPrefixLT(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixLT(prefix, column, value))
}

// WhereOrPrefixLTE 在 "OR" 条件下构建 `prefix.column <= value` 语句。
// 参见 WhereBuilder.WhereOrPrefixLTE。
// md5:5c67992d7a2d9176
// ff:条件或小于等于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:比较值
func (m *Model) WhereOrPrefixLTE(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixLTE(prefix, column, value))
}

// WhereOrPrefixGT 构建在 `OR` 条件下的 `prefix.column > value` 语句。
// 参见 WhereBuilder.WhereOrPrefixGT 的用法。
// md5:c2ab1e36dc3d561b
// ff:条件或大于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:比较值
func (m *Model) WhereOrPrefixGT(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixGT(prefix, column, value))
}

// WhereOrPrefixGTE 在 OR 条件中构建 `prefix.column >= value` 语句。
// 参考 WhereBuilder.WhereOrPrefixGTE。
// md5:a3ee1fcd237d45d6
// ff:条件或大于等于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:比较值
func (m *Model) WhereOrPrefixGTE(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixGTE(prefix, column, value))
}

// WhereOrPrefixBetween在`OR`条件下构建`prefix.column BETWEEN min AND max`语句。请参考WhereBuilder.WhereOrPrefixBetween。
// md5:bcd63a0bb32b253d
// ff:条件或取范围并带前缀
// m:
// prefix:字段前缀
// column:字段名
// min:最小值
// max:最大值
func (m *Model) WhereOrPrefixBetween(prefix string, column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixBetween(prefix, column, min, max))
}

// WhereOrPrefixLike构建了`prefix.column LIKE like`条件下的`OR`语句。请参考WhereBuilder.WhereOrPrefixLike。
// md5:42d57b9f251b31f3
// ff:条件或模糊匹配并带前缀
// m:
// prefix:字段前缀
// column:字段名
// like:通配符条件值
func (m *Model) WhereOrPrefixLike(prefix string, column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixLike(prefix, column, like))
}

// WhereOrPrefixIn 用于构建 `OR` 条件下的 `prefix.column IN (in)` 语句。
// 参见 WhereBuilder.WhereOrPrefixIn 的用法。
// md5:16a0a007f82dbf8e
// ff:条件或包含并带前缀
// m:
// prefix:字段前缀
// column:字段名
// in:包含值
func (m *Model) WhereOrPrefixIn(prefix string, column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixIn(prefix, column, in))
}

// WhereOrPrefixNull 用于在`OR`条件中构建 `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` 的语句。
// 参考 WhereBuilder.WhereOrPrefixNull。
// md5:526f8d0f44781d5f
// ff:条件或NULL值并带前缀
// m:
// prefix:字段前缀
// columns:字段名
func (m *Model) WhereOrPrefixNull(prefix string, columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNull(prefix, columns...))
}

// WhereOrPrefixNotBetween 在 `OR` 条件下构建 `prefix.column NOT BETWEEN min AND max` 语句。
// 参考 WhereBuilder.WhereOrPrefixNotBetween。
// md5:a2c385cd5a8a13f7
// ff:条件或取范围以外并带前缀
// m:
// prefix:字段前缀
// column:字段名
// min:最小值
// max:最大值
func (m *Model) WhereOrPrefixNotBetween(prefix string, column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNotBetween(prefix, column, min, max))
}

// WhereOrPrefixNotLike 在 `OR` 条件下构建 `prefix.column NOT LIKE like` 语句。
// 参见 WhereBuilder.WhereOrPrefixNotLike。
// md5:2c4e846be65e70c1
// ff:条件或模糊匹配以外并带前缀
// m:
// prefix:字段前缀
// column:字段名
// like:通配符条件值
func (m *Model) WhereOrPrefixNotLike(prefix string, column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNotLike(prefix, column, like))
}

// WhereOrPrefixNotIn 用于构建 `prefix.column NOT IN (in)` 的SQL语句。
// 参见WhereBuilder中的WhereOrPrefixNotIn方法。
// md5:890322e319ab2ff8
// ff:条件或不包含并带前缀
// m:
// prefix:字段前缀
// column:字段名
// in:不包含值
func (m *Model) WhereOrPrefixNotIn(prefix string, column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNotIn(prefix, column, in))
}

// WhereOrPrefixNotNull 在 `OR` 条件中构建 `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` 语句。
// 参考 WhereBuilder.WhereOrPrefixNotNull。
// md5:537c634be4bd78f3
// ff:条件或非Null并带前缀
// m:
// prefix:字段前缀
// columns:字段名
func (m *Model) WhereOrPrefixNotNull(prefix string, columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNotNull(prefix, columns...))
}

// WhereOrPrefixNot在`OR`条件下构建`prefix.column != value`语句。请参阅WhereBuilder.WhereOrPrefixNot。
// md5:46f5833fb4aa8a66
// ff:条件或不等于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:值
func (m *Model) WhereOrPrefixNot(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrPrefixNot(prefix, column, value))
}
