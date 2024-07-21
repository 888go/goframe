// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码文件受MIT许可协议条款的约束。
// 如果您没有随此文件分发MIT许可证的副本，
// 您可以从https://gitee.com/gogf/gf获取。
// md5:358e3ba76264232a

package gdb

// callWhereBuilder 创建并返回一个新的Model实例，如果当前Model是安全的，则设置其WhereBuilder。
// 如果当前Model不是安全的，它将直接设置WhereBuilder并返回当前Model。
// md5:f847d7aad8140312
func (m *Model) callWhereBuilder(builder *WhereBuilder) *Model {
	model := m.getModel()
	model.whereBuilder = builder
	return model
}

// Where 为构建器设置条件语句。参数 `where` 可以是
// 字符串/映射/gmap/切片/结构体/结构体指针等类型。注意，如果调用多次，
// 多个条件将使用 "AND" 连接成 WHERE 语句。
// 参见 WhereBuilder.Where。
// md5:d18e25934a430281
// ff:条件
// m:
// where:条件
// args:参数
func (m *Model) Where(where interface{}, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.Where(where, args...))
}

// Wheref 使用 fmt.Sprintf 和参数构建条件字符串。
// 注意，如果 `args` 的数量多于 `format` 中的占位符，
// 多余的 `args` 将用作 Model 的 WHERE 条件参数。
// 参见 WhereBuilder.Wheref。
// md5:ecb69c9051fee97d
// ff:条件格式化
// m:
// format:格式
// args:参数
func (m *Model) Wheref(format string, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.Wheref(format, args...))
}

// WherePri 的逻辑与 Model.Where 相同，但当参数 `where` 是一个单一的条件，如 int、string、float 或 slice 时，它会将条件视为主键值。也就是说，如果主键是 "id"，并且给定的 `where` 参数为 "123"，WherePri 函数会将条件解析为 "id=123"，而 Model.Where 则会将条件视为字符串 "123"。
// 参阅 WhereBuilder.WherePri。
// md5:13dc66b105f841e9
// ff:条件并识别主键
// m:
// where:条件
// args:参数
func (m *Model) WherePri(where interface{}, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePri(where, args...))
}

// WhereLT 用于构建 `column < value` 的语句。
// 参见 WhereBuilder.WhereLT。
// md5:16aa57cc63797f44
// ff:条件小于
// m:
// column:字段名
// value:比较值
func (m *Model) WhereLT(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereLT(column, value))
}

// WhereLTE 构建 `column <= value` 条件语句。
// 参考 WhereBuilder.WhereLTE。
// md5:85440b32a52d1c2c
// ff:条件小于等于
// m:
// column:字段名
// value:比较值
func (m *Model) WhereLTE(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereLTE(column, value))
}

// WhereGT构建`column > value`语句。
// 参见WhereBuilder.WhereGT。
// md5:900e4d7769d650c2
// ff:条件大于
// m:
// column:字段名
// value:比较值
func (m *Model) WhereGT(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereGT(column, value))
}

// WhereGTE构建`column >= value`语句。
// 参见WhereBuilder.WhereGTE。
// md5:cd051232a4d3707c
// ff:条件大于等于
// m:
// column:字段名
// value:比较值
func (m *Model) WhereGTE(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereGTE(column, value))
}

// WhereBetween 用于构建 `column BETWEEN min AND max` 的语句。
// 参见 WhereBuilder.WhereBetween 的使用。
// md5:88a499f60e180ae2
// ff:条件取范围
// m:
// column:字段名
// min:最小值
// max:最大值
func (m *Model) WhereBetween(column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereBetween(column, min, max))
}

// WhereLike 构建 `column LIKE like` 语句。
// 参考 WhereBuilder.WhereLike。
// md5:0d0b14277dfc3be8
// ff:条件模糊匹配
// m:
// column:字段名
// like:通配符条件值
func (m *Model) WhereLike(column string, like string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereLike(column, like))
}

// WhereIn 构建 `column IN (in)` 语句。
// 参考 WhereBuilder.WhereIn。
// md5:b4b77eb17cf9b671
// ff:条件包含
// m:
// column:字段名
// in:包含值
func (m *Model) WhereIn(column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereIn(column, in))
}

// WhereNull构建`columns[0] IS NULL AND columns[1] IS NULL ...`语句。
// 参见WhereBuilder.WhereNull。
// md5:af598d8379efcab6
// ff:条件NULL值
// m:
// columns:字段名
func (m *Model) WhereNull(columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereNull(columns...))
}

// WhereNotBetween 用于构建 `column NOT BETWEEN min AND max` 的SQL语句。
// 参见WhereBuilder.WhereNotBetween的用法。
// md5:be0e739db028ee79
// ff:条件取范围以外
// m:
// column:字段名
// min:最小值
// max:最大值
func (m *Model) WhereNotBetween(column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereNotBetween(column, min, max))
}

// WhereNotLike 构建 `column NOT LIKE like` 语句。
// 参考 WhereBuilder.WhereNotLike。
// md5:1be6fd9ae98cc213
// ff:条件模糊匹配以外
// m:
// column:字段名
// like:通配符条件值
func (m *Model) WhereNotLike(column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereNotLike(column, like))
}

// WhereNot构建`column != value`语句。
// 参见WhereBuilder.WhereNot。
// md5:973b26be7332c7b2
// ff:条件不等于
// m:
// column:字段名
// value:值
func (m *Model) WhereNot(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereNot(column, value))
}

// WhereNotIn构建`column NOT IN (in)`语句。
// 请参阅WhereBuilder.WhereNotIn。
// md5:4fbfcfa2b85a83d2
// ff:条件不包含
// m:
// column:字段名
// in:不包含值
func (m *Model) WhereNotIn(column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereNotIn(column, in))
}

				// WhereNotNull 构建 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。
				// 参见 WhereBuilder.WhereNotNull 的用法。
				// md5:05b9b4179d41a28b
// ff:条件非Null
// m:
// columns:字段名
func (m *Model) WhereNotNull(columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereNotNull(columns...))
}
