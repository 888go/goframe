// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

// X条件或并带前缀 的行为与 WhereOr 相似，但它会在where语句中的每个字段前添加前缀。
// 参考 WhereBuilder.X条件或并带前缀。
// md5:4ea5d18d5615ff17
func (m *Model) X条件或并带前缀(字段前缀 string, 条件 interface{}, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或并带前缀(字段前缀, 条件, 参数...))
}

// X条件或小于并带前缀在"OR"条件下构建`prefix.column < value`语句。参阅WhereBuilder.X条件或小于并带前缀。
// md5:0a8c07ff239fa7e7
func (m *Model) X条件或小于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或小于并带前缀(字段前缀, 字段名, 比较值))
}

// X条件或小于等于并带前缀 在 "OR" 条件下构建 `prefix.column <= value` 语句。
// 参见 WhereBuilder.X条件或小于等于并带前缀。
// md5:5c67992d7a2d9176
func (m *Model) X条件或小于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或小于等于并带前缀(字段前缀, 字段名, 比较值))
}

// X条件或大于并带前缀 构建在 `OR` 条件下的 `prefix.column > value` 语句。
// 参见 WhereBuilder.X条件或大于并带前缀 的用法。
// md5:c2ab1e36dc3d561b
func (m *Model) X条件或大于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或大于并带前缀(字段前缀, 字段名, 比较值))
}

// X条件或大于等于并带前缀 在 OR 条件中构建 `prefix.column >= value` 语句。
// 参考 WhereBuilder.X条件或大于等于并带前缀。
// md5:a3ee1fcd237d45d6
func (m *Model) X条件或大于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或大于等于并带前缀(字段前缀, 字段名, 比较值))
}

// X条件或取范围并带前缀在`OR`条件下构建`prefix.column BETWEEN min AND max`语句。请参考WhereBuilder.X条件或取范围并带前缀。
// md5:bcd63a0bb32b253d
func (m *Model) X条件或取范围并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或取范围并带前缀(字段前缀, 字段名, 最小值, 最大值))
}

// X条件或模糊匹配并带前缀构建了`prefix.column LIKE like`条件下的`OR`语句。请参考WhereBuilder.X条件或模糊匹配并带前缀。
// md5:42d57b9f251b31f3
func (m *Model) X条件或模糊匹配并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或模糊匹配并带前缀(字段前缀, 字段名, 通配符条件值))
}

// X条件或包含并带前缀 用于构建 `OR` 条件下的 `prefix.column IN (in)` 语句。
// 参见 WhereBuilder.X条件或包含并带前缀 的用法。
// md5:16a0a007f82dbf8e
func (m *Model) X条件或包含并带前缀(字段前缀 string, 字段名 string, 包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或包含并带前缀(字段前缀, 字段名, 包含值))
}

// X条件或NULL值并带前缀 用于在`OR`条件中构建 `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` 的语句。
// 参考 WhereBuilder.X条件或NULL值并带前缀。
// md5:526f8d0f44781d5f
func (m *Model) X条件或NULL值并带前缀(字段前缀 string, 字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或NULL值并带前缀(字段前缀, 字段名...))
}

// X条件或取范围以外并带前缀 在 `OR` 条件下构建 `prefix.column NOT BETWEEN min AND max` 语句。
// 参考 WhereBuilder.X条件或取范围以外并带前缀。
// md5:a2c385cd5a8a13f7
func (m *Model) X条件或取范围以外并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或取范围以外并带前缀(字段前缀, 字段名, 最小值, 最大值))
}

// X条件或模糊匹配以外并带前缀 在 `OR` 条件下构建 `prefix.column NOT LIKE like` 语句。
// 参见 WhereBuilder.X条件或模糊匹配以外并带前缀。
// md5:2c4e846be65e70c1
func (m *Model) X条件或模糊匹配以外并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或模糊匹配以外并带前缀(字段前缀, 字段名, 通配符条件值))
}

// X条件或不包含并带前缀 用于构建 `prefix.column NOT IN (in)` 的SQL语句。
// 参见WhereBuilder中的X条件或不包含并带前缀方法。
// md5:890322e319ab2ff8
func (m *Model) X条件或不包含并带前缀(字段前缀 string, 字段名 string, 不包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或不包含并带前缀(字段前缀, 字段名, 不包含值))
}

// X条件或非Null并带前缀 在 `OR` 条件中构建 `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` 语句。
// 参考 WhereBuilder.X条件或非Null并带前缀。
// md5:537c634be4bd78f3
func (m *Model) X条件或非Null并带前缀(字段前缀 string, 字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或非Null并带前缀(字段前缀, 字段名...))
}

// X条件或不等于并带前缀在`OR`条件下构建`prefix.column != value`语句。请参阅WhereBuilder.X条件或不等于并带前缀。
// md5:46f5833fb4aa8a66
func (m *Model) X条件或不等于并带前缀(字段前缀 string, 字段名 string, 值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或不等于并带前缀(字段前缀, 字段名, 值))
}
