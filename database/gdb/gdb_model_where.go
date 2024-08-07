// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码文件受MIT许可协议条款的约束。
// 如果您没有随此文件分发MIT许可证的副本，
// 您可以从https://gitee.com/gogf/gf获取。
// md5:358e3ba76264232a

package db类

// callWhereBuilder 创建并返回一个新的Model实例，如果当前Model是安全的，则设置其WhereBuilder。
// 如果当前Model不是安全的，它将直接设置WhereBuilder并返回当前Model。
// md5:f847d7aad8140312
func (m *Model) callWhereBuilder(builder *WhereBuilder) *Model {
	model := m.getModel()
	model.whereBuilder = builder
	return model
}

// X条件 为构建器设置条件语句。参数 `where` 可以是
// 字符串/映射/gmap/切片/结构体/结构体指针等类型。注意，如果调用多次，
// 多个条件将使用 "AND" 连接成 WHERE 语句。
// 参见 WhereBuilder.X条件。
// md5:d18e25934a430281
func (m *Model) X条件(条件 interface{}, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件(条件, 参数...))
}

// X条件格式化 使用 fmt.Sprintf 和参数构建条件字符串。
// 注意，如果 `args` 的数量多于 `format` 中的占位符，
// 多余的 `args` 将用作 Model 的 WHERE 条件参数。
// 参见 WhereBuilder.X条件格式化。
// md5:ecb69c9051fee97d
func (m *Model) X条件格式化(格式 string, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件格式化(格式, 参数...))
}

// X条件并识别主键 的逻辑与 Model.Where 相同，但当参数 `where` 是一个单一的条件，如 int、string、float 或 slice 时，它会将条件视为主键值。也就是说，如果主键是 "id"，并且给定的 `where` 参数为 "123"，X条件并识别主键 函数会将条件解析为 "id=123"，而 Model.Where 则会将条件视为字符串 "123"。
// 参阅 WhereBuilder.X条件并识别主键。
// md5:13dc66b105f841e9
func (m *Model) X条件并识别主键(条件 interface{}, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件并识别主键(条件, 参数...))
}

// X条件小于 用于构建 `column < value` 的语句。
// 参见 WhereBuilder.X条件小于。
// md5:16aa57cc63797f44
func (m *Model) X条件小于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件小于(字段名, 比较值))
}

// X条件小于等于 构建 `column <= value` 条件语句。
// 参考 WhereBuilder.X条件小于等于。
// md5:85440b32a52d1c2c
func (m *Model) X条件小于等于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件小于等于(字段名, 比较值))
}

// X条件大于构建`column > value`语句。
// 参见WhereBuilder.X条件大于。
// md5:900e4d7769d650c2
func (m *Model) X条件大于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件大于(字段名, 比较值))
}

// X条件大于等于构建`column >= value`语句。
// 参见WhereBuilder.X条件大于等于。
// md5:cd051232a4d3707c
func (m *Model) X条件大于等于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件大于等于(字段名, 比较值))
}

// X条件取范围 用于构建 `column BETWEEN min AND max` 的语句。
// 参见 WhereBuilder.X条件取范围 的使用。
// md5:88a499f60e180ae2
func (m *Model) X条件取范围(字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件取范围(字段名, 最小值, 最大值))
}

// X条件模糊匹配 构建 `column LIKE like` 语句。
// 参考 WhereBuilder.X条件模糊匹配。
// md5:0d0b14277dfc3be8
func (m *Model) X条件模糊匹配(字段名 string, 通配符条件值 string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件模糊匹配(字段名, 通配符条件值))
}

// X条件包含 构建 `column IN (in)` 语句。
// 参考 WhereBuilder.X条件包含。
// md5:b4b77eb17cf9b671
func (m *Model) X条件包含(字段名 string, 包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件包含(字段名, 包含值))
}

// X条件NULL值构建`columns[0] IS NULL AND columns[1] IS NULL ...`语句。
// 参见WhereBuilder.X条件NULL值。
// md5:af598d8379efcab6
func (m *Model) X条件NULL值(字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件NULL值(字段名...))
}

// X条件取范围以外 用于构建 `column NOT BETWEEN min AND max` 的SQL语句。
// 参见WhereBuilder.X条件取范围以外的用法。
// md5:be0e739db028ee79
func (m *Model) X条件取范围以外(字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件取范围以外(字段名, 最小值, 最大值))
}

// X条件模糊匹配以外 构建 `column NOT LIKE like` 语句。
// 参考 WhereBuilder.X条件模糊匹配以外。
// md5:1be6fd9ae98cc213
func (m *Model) X条件模糊匹配以外(字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件模糊匹配以外(字段名, 通配符条件值))
}

// X条件不等于构建`column != value`语句。
// 参见WhereBuilder.X条件不等于。
// md5:973b26be7332c7b2
func (m *Model) X条件不等于(字段名 string, 值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件不等于(字段名, 值))
}

// X条件不包含构建`column NOT IN (in)`语句。
// 请参阅WhereBuilder.X条件不包含。
// md5:4fbfcfa2b85a83d2
func (m *Model) X条件不包含(字段名 string, 不包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件不包含(字段名, 不包含值))
}

// X条件非Null 构建 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。
// 参见 WhereBuilder.X条件非Null 的用法。
// md5:05b9b4179d41a28b
func (m *Model) X条件非Null(字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件非Null(字段名...))
}
