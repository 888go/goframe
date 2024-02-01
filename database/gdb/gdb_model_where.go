// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://githum.com/gogf/gf 获取一份。

package gdb

// callWhereBuilder 函数根据当前 Model 是否安全，创建并返回一个新的 Model。
// 如果当前 Model 是安全的，则设置其 WhereBuilder；如果不安全，则直接设置 WhereBuilder 并返回当前 Model。
func (m *Model) callWhereBuilder(builder *WhereBuilder) *Model {
	model := m.getModel()
	model.whereBuilder = builder
	return model
}

// Where 为生成器设置条件语句。参数`where`可以是以下类型：
// string/map/gmap/slice/struct/*struct 等等。请注意，如果多次调用，
// 多个条件将会通过 "AND" 连接符合并到 where 语句中。
// 参见 WhereBuilder.Where 。
func (m *Model) Where(where interface{}, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.Where(where, args...))
}

// Wheref 通过 fmt.Sprintf 和参数构建条件字符串。
// 注意，如果 `args` 的数量大于 `format` 中的占位符数量，
// 额外的 `args` 将作为 Model 的 where 条件参数使用。
// 参见 WhereBuilder.Wheref。
func (m *Model) Wheref(format string, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.Wheref(format, args...))
}

// WherePri 执行的逻辑与 Model.Where 相同，但区别在于：如果参数 `where` 是一个单一条件，如 int、string、float 或 slice 类型，
// 那么它会将这个条件视为主键值。也就是说，如果主键是 "id"，给定的 `where` 参数为 "123"，那么 WherePri 函数会把条件当作 "id=123" 处理，
// 但是 Model.Where 函数会将该条件当作字符串 "123" 处理。
// 请参阅 WhereBuilder.WherePri。
func (m *Model) WherePri(where interface{}, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WherePri(where, args...))
}

// WhereLT构建`column < value`表达式语句。
// 参见WhereBuilder.WhereLT。
func (m *Model) WhereLT(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereLT(column, value))
}

// WhereLTE 用于构建 `column <= value` 条件语句。
// 参见 WhereBuilder.WhereLTE。
func (m *Model) WhereLTE(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereLTE(column, value))
}

// WhereGT构建`column > value`语句。
// 请参阅WhereBuilder.WhereGT。
func (m *Model) WhereGT(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereGT(column, value))
}

// WhereGTE 用于构建 `column >= value` 条件语句。
// 参见 WhereBuilder.WhereGTE。
func (m *Model) WhereGTE(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereGTE(column, value))
}

// WhereBetween 用于构建 `column BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereBetween。
func (m *Model) WhereBetween(column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereBetween(column, min, max))
}

// WhereLike 用于构建 `column LIKE like` 语句。
// 参见 WhereBuilder.WhereLike。
func (m *Model) WhereLike(column string, like string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereLike(column, like))
}

// WhereIn 构建 `column IN (in)` 语句。
// 参见 WhereBuilder.WhereIn。
func (m *Model) WhereIn(column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereIn(column, in))
}

// WhereNull 构建 `columns[0] IS NULL AND columns[1] IS NULL ...` 语句。
// 参见 WhereBuilder.WhereNull 方法。
func (m *Model) WhereNull(columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereNull(columns...))
}

// WhereNotBetween 用于构建 `column NOT BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereNotBetween。
func (m *Model) WhereNotBetween(column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereNotBetween(column, min, max))
}

// WhereNotLike 用于构建 `column NOT LIKE like` 语句。
// 参见 WhereBuilder.WhereNotLike。
func (m *Model) WhereNotLike(column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereNotLike(column, like))
}

// WhereNot 用于构建 `column != value` 条件语句。
// 请参阅 WhereBuilder.WhereNot。
func (m *Model) WhereNot(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereNot(column, value))
}

// WhereNotIn 用于构建 `column NOT IN (in)` 语句。
// 请参阅 WhereBuilder.WhereNotIn。
func (m *Model) WhereNotIn(column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereNotIn(column, in))
}

// WhereNotNull 用于构建 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。
// 请参考 WhereBuilder.WhereNotNull。
func (m *Model) WhereNotNull(columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereNotNull(columns...))
}
