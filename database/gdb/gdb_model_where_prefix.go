// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

// WherePrefix 的行为与 Where 相似，但它会为 where 语句中的每个字段添加前缀。
// 参考 WhereBuilder.WherePrefix。
// md5:857520a0e9f2f42c
// ff:条件带前缀
// m:
// prefix:字段前缀
// where:条件
// args:参数
func (m *Model) WherePrefix(prefix string, where interface{}, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefix(prefix, where, args...))
}

// WherePrefixLT构建`prefix.column < value`语句。
// 参考WhereBuilder.WherePrefixLT。
// md5:772fb94f7bcccb22
// ff:条件小于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:比较值
func (m *Model) WherePrefixLT(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixLT(prefix, column, value))
}

// WherePrefixLTE构建`prefix.column <= value`语句。
// 参见WhereBuilder.WherePrefixLTE。
// md5:09fe8d74131bca96
// ff:条件小于等于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:比较值
func (m *Model) WherePrefixLTE(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixLTE(prefix, column, value))
}

// WherePrefixGT 构建 `prefix.column > value` 的语句。
// 参见 WhereBuilder.WherePrefixGT 的用法。
// md5:7bd6fca29c275204
// ff:条件大于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:比较值
func (m *Model) WherePrefixGT(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixGT(prefix, column, value))
}

// WherePrefixGTE 构建 `prefix.column >= value` 语句。
// 参见 WhereBuilder.WherePrefixGTE。
// md5:f4256046c4ee0127
// ff:条件大于等于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:比较值
func (m *Model) WherePrefixGTE(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixGTE(prefix, column, value))
}

// WherePrefixBetween构建`prefix.column BETWEEN min AND max`语句。
// 参见WhereBuilder.WherePrefixBetween。
// md5:5bd292414bb58ab2
// ff:条件取范围并带前缀
// m:
// prefix:字段前缀
// column:字段名
// min:最小值
// max:最大值
func (m *Model) WherePrefixBetween(prefix string, column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixBetween(prefix, column, min, max))
}

// WherePrefixLike构建`prefix.column LIKE like`语句。
// 请参考WhereBuilder.WherePrefixLike。
// md5:52c7b5d6907728da
// ff:条件模糊匹配并带前缀
// m:
// prefix:字段前缀
// column:字段名
// like:通配符条件值
func (m *Model) WherePrefixLike(prefix string, column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixLike(prefix, column, like))
}

// WherePrefixIn 构建 `prefix.column IN (in)` 语句。
// 参见 WhereBuilder.WherePrefixIn。
// md5:21e5a1f77d32a6ce
// ff:条件包含并带前缀
// m:
// prefix:字段前缀
// column:字段名
// in:包含值
func (m *Model) WherePrefixIn(prefix string, column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixIn(prefix, column, in))
}

// WherePrefixNull 用于构建 `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` 的SQL语句。
// 参考 WhereBuilder.WherePrefixNull 方法。
// md5:e66f4e8ba9e64abe
// ff:条件NULL值并带前缀
// m:
// prefix:字段前缀
// columns:字段名
func (m *Model) WherePrefixNull(prefix string, columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixNull(prefix, columns...))
}

// WherePrefixNotBetween构建`prefix.column NOT BETWEEN min AND max`语句。
// 参见WhereBuilder.WherePrefixNotBetween。
// md5:2cfa2b60438f51a2
// ff:条件取范围以外并带前缀
// m:
// prefix:字段前缀
// column:字段名
// min:最小值
// max:最大值
func (m *Model) WherePrefixNotBetween(prefix string, column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixNotBetween(prefix, column, min, max))
}

// WherePrefixNotLike构建`prefix.column NOT LIKE like`语句。
// 参见WhereBuilder.WherePrefixNotLike。
// md5:a9a2f0b1ba94005b
// ff:条件模糊匹配以外并带前缀
// m:
// prefix:字段前缀
// column:字段名
// like:通配符条件值
func (m *Model) WherePrefixNotLike(prefix string, column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixNotLike(prefix, column, like))
}

// WherePrefixNot 构建 `prefix.column != value` 的语句。
// 参见 WhereBuilder.WherePrefixNot 的用法。
// md5:52ce1c05f4e382fb
// ff:条件不等于并带前缀
// m:
// prefix:字段前缀
// column:字段名
// value:值
func (m *Model) WherePrefixNot(prefix string, column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixNot(prefix, column, value))
}

// WherePrefixNotIn 构建 `prefix.column NOT IN (in)` 语句。
// 参考 WhereBuilder.WherePrefixNotIn。
// md5:ae3cd87814389feb
// ff:条件不包含并带前缀
// m:
// prefix:字段前缀
// column:字段名
// in:不包含值
func (m *Model) WherePrefixNotIn(prefix string, column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixNotIn(prefix, column, in))
}

				// WherePrefixNotNull 构建 `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` 语句。
				// 参见 WhereBuilder.WherePrefixNotNull。
				// md5:953102e755997338
// ff:条件非Null并带前缀
// m:
// prefix:字段前缀
// columns:字段名
func (m *Model) WherePrefixNotNull(prefix string, columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePrefixNotNull(prefix, columns...))
}
