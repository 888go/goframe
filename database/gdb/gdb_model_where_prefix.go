// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

// X条件带前缀 的行为与 Where 相似，但它会为 where 语句中的每个字段添加前缀。
// 参考 WhereBuilder.X条件带前缀。
// md5:857520a0e9f2f42c
func (m *Model) X条件带前缀(字段前缀 string, 条件 interface{}, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件带前缀(字段前缀, 条件, 参数...))
}

// X条件小于并带前缀构建`prefix.column < value`语句。
// 参考WhereBuilder.X条件小于并带前缀。
// md5:772fb94f7bcccb22
func (m *Model) X条件小于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件小于并带前缀(字段前缀, 字段名, 比较值))
}

// X条件小于等于并带前缀构建`prefix.column <= value`语句。
// 参见WhereBuilder.X条件小于等于并带前缀。
// md5:09fe8d74131bca96
func (m *Model) X条件小于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件小于等于并带前缀(字段前缀, 字段名, 比较值))
}

// X条件大于并带前缀 构建 `prefix.column > value` 的语句。
// 参见 WhereBuilder.X条件大于并带前缀 的用法。
// md5:7bd6fca29c275204
func (m *Model) X条件大于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件大于并带前缀(字段前缀, 字段名, 比较值))
}

// X条件大于等于并带前缀 构建 `prefix.column >= value` 语句。
// 参见 WhereBuilder.X条件大于等于并带前缀。
// md5:f4256046c4ee0127
func (m *Model) X条件大于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件大于等于并带前缀(字段前缀, 字段名, 比较值))
}

// X条件取范围并带前缀构建`prefix.column BETWEEN min AND max`语句。
// 参见WhereBuilder.X条件取范围并带前缀。
// md5:5bd292414bb58ab2
func (m *Model) X条件取范围并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件取范围并带前缀(字段前缀, 字段名, 最小值, 最大值))
}

// X条件模糊匹配并带前缀构建`prefix.column LIKE like`语句。
// 请参考WhereBuilder.X条件模糊匹配并带前缀。
// md5:52c7b5d6907728da
func (m *Model) X条件模糊匹配并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件模糊匹配并带前缀(字段前缀, 字段名, 通配符条件值))
}

// X条件包含并带前缀 构建 `prefix.column IN (in)` 语句。
// 参见 WhereBuilder.X条件包含并带前缀。
// md5:21e5a1f77d32a6ce
func (m *Model) X条件包含并带前缀(字段前缀 string, 字段名 string, 包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件包含并带前缀(字段前缀, 字段名, 包含值))
}

// X条件NULL值并带前缀 用于构建 `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` 的SQL语句。
// 参考 WhereBuilder.X条件NULL值并带前缀 方法。
// md5:e66f4e8ba9e64abe
func (m *Model) X条件NULL值并带前缀(字段前缀 string, 字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件NULL值并带前缀(字段前缀, 字段名...))
}

// X条件取范围以外并带前缀构建`prefix.column NOT BETWEEN min AND max`语句。
// 参见WhereBuilder.X条件取范围以外并带前缀。
// md5:2cfa2b60438f51a2
func (m *Model) X条件取范围以外并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件取范围以外并带前缀(字段前缀, 字段名, 最小值, 最大值))
}

// X条件模糊匹配以外并带前缀构建`prefix.column NOT LIKE like`语句。
// 参见WhereBuilder.X条件模糊匹配以外并带前缀。
// md5:a9a2f0b1ba94005b
func (m *Model) X条件模糊匹配以外并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件模糊匹配以外并带前缀(字段前缀, 字段名, 通配符条件值))
}

// X条件不等于并带前缀 构建 `prefix.column != value` 的语句。
// 参见 WhereBuilder.X条件不等于并带前缀 的用法。
// md5:52ce1c05f4e382fb
func (m *Model) X条件不等于并带前缀(字段前缀 string, 字段名 string, 值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件不等于并带前缀(字段前缀, 字段名, 值))
}

// X条件不包含并带前缀 构建 `prefix.column NOT IN (in)` 语句。
// 参考 WhereBuilder.X条件不包含并带前缀。
// md5:ae3cd87814389feb
func (m *Model) X条件不包含并带前缀(字段前缀 string, 字段名 string, 不包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件不包含并带前缀(字段前缀, 字段名, 不包含值))
}

// X条件非Null并带前缀 构建 `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` 语句。
// 参见 WhereBuilder.X条件非Null并带前缀。
// md5:953102e755997338
func (m *Model) X条件非Null并带前缀(字段前缀 string, 字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件非Null并带前缀(字段前缀, 字段名...))
}
