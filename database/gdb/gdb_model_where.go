// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://githum.com/gogf/gf 获取一份。

package db类

// callWhereBuilder 函数根据当前 Model 是否安全，创建并返回一个新的 Model。
// 如果当前 Model 是安全的，则设置其 WhereBuilder；如果不安全，则直接设置 WhereBuilder 并返回当前 Model。
func (m *Model) callWhereBuilder(builder *X组合条件) *Model {
	model := m.getModel()
	model.whereBuilder = builder
	return model
}

// Where 为生成器设置条件语句。参数`where`可以是以下类型：
// string/map/gmap/slice/struct/*struct 等等。请注意，如果多次调用，
// 多个条件将会通过 "AND" 连接符合并到 where 语句中。
// 参见 WhereBuilder.Where 。
func (m *Model) X条件(条件 interface{}, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件(条件, 参数...))
}

// Wheref 通过 fmt.Sprintf 和参数构建条件字符串。
// 注意，如果 `args` 的数量大于 `format` 中的占位符数量，
// 额外的 `args` 将作为 Model 的 where 条件参数使用。
// 参见 WhereBuilder.Wheref。
func (m *Model) X条件格式化(格式 string, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件格式化(格式, 参数...))
}

// WherePri 执行的逻辑与 Model.Where 相同，但区别在于：如果参数 `where` 是一个单一条件，如 int、string、float 或 slice 类型，
// 那么它会将这个条件视为主键值。也就是说，如果主键是 "id"，给定的 `where` 参数为 "123"，那么 WherePri 函数会把条件当作 "id=123" 处理，
// 但是 Model.Where 函数会将该条件当作字符串 "123" 处理。
// 请参阅 WhereBuilder.WherePri。
func (m *Model) X条件并识别主键(条件 interface{}, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件并识别主键(条件, 参数...))
}

// WhereLT构建`column < value`表达式语句。
// 参见WhereBuilder.WhereLT。
func (m *Model) X条件小于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件小于(字段名, 比较值))
}

// WhereLTE 用于构建 `column <= value` 条件语句。
// 参见 WhereBuilder.WhereLTE。
func (m *Model) X条件小于等于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件小于等于(字段名, 比较值))
}

// WhereGT构建`column > value`语句。
// 请参阅WhereBuilder.WhereGT。
func (m *Model) X条件大于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件大于(字段名, 比较值))
}

// WhereGTE 用于构建 `column >= value` 条件语句。
// 参见 WhereBuilder.WhereGTE。
func (m *Model) X条件大于等于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件大于等于(字段名, 比较值))
}

// WhereBetween 用于构建 `column BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereBetween。
func (m *Model) X条件取范围(字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件取范围(字段名, 最小值, 最大值))
}

// WhereLike 用于构建 `column LIKE like` 语句。
// 参见 WhereBuilder.WhereLike。
func (m *Model) X条件模糊匹配(字段名 string, 通配符条件值 string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件模糊匹配(字段名, 通配符条件值))
}

// WhereIn 构建 `column IN (in)` 语句。
// 参见 WhereBuilder.WhereIn。
func (m *Model) X条件包含(字段名 string, 包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件包含(字段名, 包含值))
}

// WhereNull 构建 `columns[0] IS NULL AND columns[1] IS NULL ...` 语句。
// 参见 WhereBuilder.WhereNull 方法。
func (m *Model) X条件NULL值(字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件NULL值(字段名...))
}

// WhereNotBetween 用于构建 `column NOT BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereNotBetween。
func (m *Model) X条件取范围以外(字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件取范围以外(字段名, 最小值, 最大值))
}

// WhereNotLike 用于构建 `column NOT LIKE like` 语句。
// 参见 WhereBuilder.WhereNotLike。
func (m *Model) X条件模糊匹配以外(字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件模糊匹配以外(字段名, 通配符条件值))
}

// WhereNot 用于构建 `column != value` 条件语句。
// 请参阅 WhereBuilder.WhereNot。
func (m *Model) X条件不等于(字段名 string, 值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件不等于(字段名, 值))
}

// WhereNotIn 用于构建 `column NOT IN (in)` 语句。
// 请参阅 WhereBuilder.WhereNotIn。
func (m *Model) X条件不包含(字段名 string, 不包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件不包含(字段名, 不包含值))
}

// WhereNotNull 用于构建 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。
// 请参考 WhereBuilder.WhereNotNull。
func (m *Model) X条件非Null(字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件非Null(字段名...))
}
