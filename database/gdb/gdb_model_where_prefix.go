// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

// WherePrefix 的行为类似于 Where，但它会在 where 语句中的每个字段前添加前缀。
// 请参阅 WhereBuilder.WherePrefix。
func (m *Model) X条件带前缀(字段前缀 string, 条件 interface{}, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件带前缀(字段前缀, 条件, 参数...))
}

// WherePrefixLT 用于构建 `prefix.column < value` 语句。
// 请参阅 WhereBuilder.WherePrefixLT。
func (m *Model) X条件小于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件小于并带前缀(字段前缀, 字段名, 比较值))
}

// WherePrefixLTE 用于构建 `prefix.column <= value` 条件语句。
// 参见 WhereBuilder.WherePrefixLTE。
func (m *Model) X条件小于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件小于等于并带前缀(字段前缀, 字段名, 比较值))
}

// WherePrefixGT 用于构建 `prefix.column > value` 语句。
// 参见 WhereBuilder.WherePrefixGT。
func (m *Model) X条件大于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件大于并带前缀(字段前缀, 字段名, 比较值))
}

// WherePrefixGTE 用于构建 `prefix.column >= value` 条件语句。
// 请参阅 WhereBuilder.WherePrefixGTE。
func (m *Model) X条件大于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件大于等于并带前缀(字段前缀, 字段名, 比较值))
}

// WherePrefixBetween 用于构建 `prefix.column BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WherePrefixBetween。
func (m *Model) X条件取范围并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件取范围并带前缀(字段前缀, 字段名, 最小值, 最大值))
}

// WherePrefixLike 用于构建 `prefix.column LIKE like` 语句。
// 参见 WhereBuilder.WherePrefixLike。
func (m *Model) X条件模糊匹配并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件模糊匹配并带前缀(字段前缀, 字段名, 通配符条件值))
}

// WherePrefixIn 用于构建 `prefix.column IN (in)` 语句。
// 参见 WhereBuilder.WherePrefixIn。
func (m *Model) X条件包含并带前缀(字段前缀 string, 字段名 string, 包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件包含并带前缀(字段前缀, 字段名, 包含值))
}

// WherePrefixNull 用于构建 `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` 语句。
// 参见 WhereBuilder.WherePrefixNull。
func (m *Model) X条件NULL值并带前缀(字段前缀 string, 字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件NULL值并带前缀(字段前缀, 字段名...))
}

// WherePrefixNotBetween 用于构建 `prefix.column NOT BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WherePrefixNotBetween。
func (m *Model) X条件取范围以外并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件取范围以外并带前缀(字段前缀, 字段名, 最小值, 最大值))
}

// WherePrefixNotLike 用于构建 `prefix.column NOT LIKE like` 语句。
// 参见 WhereBuilder.WherePrefixNotLike。
func (m *Model) X条件模糊匹配以外并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件模糊匹配以外并带前缀(字段前缀, 字段名, 通配符条件值))
}

// WherePrefixNot 用于构建 `prefix.column != value` 的语句。
// 参见 WhereBuilder.WherePrefixNot。
func (m *Model) X条件不等于并带前缀(字段前缀 string, 字段名 string, 值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件不等于并带前缀(字段前缀, 字段名, 值))
}

// WherePrefixNotIn 用于构建 `prefix.column NOT IN (in)` 语句。
// 请参阅 WhereBuilder.WherePrefixNotIn。
func (m *Model) X条件不包含并带前缀(字段前缀 string, 字段名 string, 不包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件不包含并带前缀(字段前缀, 字段名, 不包含值))
}

// WherePrefixNotNull 用于构建 `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` 形式的语句。
// 请参阅 WhereBuilder.WherePrefixNotNull。
func (m *Model) X条件非Null并带前缀(字段前缀 string, 字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件非Null并带前缀(字段前缀, 字段名...))
}
