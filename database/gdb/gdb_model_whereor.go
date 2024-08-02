// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

// WhereOr在where语句中添加"OR"条件。请参考WhereBuilder.WhereOr。
// md5:3dc9824669296cea
func (m *Model) WhereOr(where interface{}, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOr(where, args...))
}

// WhereOrf 使用 fmt.Sprintf 和参数构建 `OR` 条件字符串。
// 参见 WhereBuilder.WhereOrf。
// md5:a94c42c383ac4960
func (m *Model) WhereOrf(format string, args ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrf(format, args...))
}

// WhereOrLT 在 `OR` 条件下构建 `column < value` 语句。
// 参见 WhereBuilder.WhereOrLT。
// md5:d9d4ee2080c8c040
func (m *Model) WhereOrLT(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrLT(column, value))
}

// WhereOrLTE 在`OR`条件中构建 `column <= value` 的语句。
// 参考 WhereBuilder.WhereOrLTE。
// md5:36414de9c787b690
func (m *Model) WhereOrLTE(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrLTE(column, value))
}

// WhereOrGT在`OR`条件下构建`column > value`语句。请参阅WhereBuilder.WhereOrGT。
// md5:5b5f0de728017e9e
func (m *Model) WhereOrGT(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrGT(column, value))
}

// WhereOrGTE在`OR`条件下构建`column >= value`语句。
// 参见WhereBuilder.WhereOrGTE。
// md5:5e6ab2d7c60899f4
func (m *Model) WhereOrGTE(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrGTE(column, value))
}

// WhereOrBetween 构建在 `OR` 条件下的 `column BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereOrBetween 的用法。
// md5:c9b005a18a1fb87e
func (m *Model) WhereOrBetween(column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrBetween(column, min, max))
}

// WhereOrLike 在`OR`条件中构建 `column LIKE like` 语句。
// 参考 WhereBuilder.WhereOrLike。
// md5:70c0895d15fd2cc9
func (m *Model) WhereOrLike(column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrLike(column, like))
}

// WhereOrIn在`OR`条件下构建`column IN (in)`语句。参见WhereBuilder.WhereOrIn。
// md5:fac500879081e3cc
func (m *Model) WhereOrIn(column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrIn(column, in))
}

// WhereOrNull 构建 `columns[0] IS NULL OR columns[1] IS NULL ...` 的 `OR` 条件语句。
// 参考 WhereBuilder.WhereOrNull。
// md5:66907cf860f22eff
func (m *Model) WhereOrNull(columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNull(columns...))
}

// WhereOrNotBetween 用于构建 `column NOT BETWEEN min AND max` 的 SQL 语句，在 `OR` 条件下。
// 参见 WhereBuilder.WhereOrNotBetween 的用法。
// md5:e040a9f04b492725
func (m *Model) WhereOrNotBetween(column string, min, max interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNotBetween(column, min, max))
}

// WhereOrNotLike 用于构建在`OR`条件下的 `column NOT LIKE 'like'` 语句。
// 参考 WhereBuilder.WhereOrNotLike。
// md5:588ea3675853468b
func (m *Model) WhereOrNotLike(column string, like interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNotLike(column, like))
}

// WhereOrNot构建`column != value`语句。
// 参见WhereBuilder.WhereOrNot。
// md5:076a864671142e49
func (m *Model) WhereOrNot(column string, value interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNot(column, value))
}

// WhereOrNotIn构建`column NOT IN (in)`语句。
// 参见WhereBuilder.WhereOrNotIn。
// md5:22915d1ba70db001
func (m *Model) WhereOrNotIn(column string, in interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNotIn(column, in))
}

// WhereOrNotNull 构建在 `OR` 条件下的 `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` 语句。
// 参见 WhereBuilder.WhereOrNotNull 的用法。
// md5:5645594c534afc8e
func (m *Model) WhereOrNotNull(columns ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.WhereOrNotNull(columns...))
}
