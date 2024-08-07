// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

// X条件或在where语句中添加"OR"条件。请参考WhereBuilder.X条件或。
// md5:3dc9824669296cea
func (m *Model) X条件或(条件 interface{}, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或(条件, 参数...))
}

// X条件或格式化 使用 fmt.Sprintf 和参数构建 `OR` 条件字符串。
// 参见 WhereBuilder.X条件或格式化。
// md5:a94c42c383ac4960
func (m *Model) X条件或格式化(格式 string, 参数 ...interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或格式化(格式, 参数...))
}

// X条件或小于 在 `OR` 条件下构建 `column < value` 语句。
// 参见 WhereBuilder.X条件或小于。
// md5:d9d4ee2080c8c040
func (m *Model) X条件或小于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或小于(字段名, 比较值))
}

// X条件或小于等于 在`OR`条件中构建 `column <= value` 的语句。
// 参考 WhereBuilder.X条件或小于等于。
// md5:36414de9c787b690
func (m *Model) X条件或小于等于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或小于等于(字段名, 比较值))
}

// X条件或大于在`OR`条件下构建`column > value`语句。请参阅WhereBuilder.X条件或大于。
// md5:5b5f0de728017e9e
func (m *Model) X条件或大于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或大于(字段名, 比较值))
}

// X条件或大于等于在`OR`条件下构建`column >= value`语句。
// 参见WhereBuilder.X条件或大于等于。
// md5:5e6ab2d7c60899f4
func (m *Model) X条件或大于等于(字段名 string, 比较值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或大于等于(字段名, 比较值))
}

// X条件或取范围 构建在 `OR` 条件下的 `column BETWEEN min AND max` 语句。
// 参见 WhereBuilder.X条件或取范围 的用法。
// md5:c9b005a18a1fb87e
func (m *Model) X条件或取范围(字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或取范围(字段名, 最小值, 最大值))
}

// X条件或模糊匹配 在`OR`条件中构建 `column LIKE like` 语句。
// 参考 WhereBuilder.X条件或模糊匹配。
// md5:70c0895d15fd2cc9
func (m *Model) X条件或模糊匹配(字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或模糊匹配(字段名, 通配符条件值))
}

// X条件或包含在`OR`条件下构建`column IN (in)`语句。参见WhereBuilder.X条件或包含。
// md5:fac500879081e3cc
func (m *Model) X条件或包含(字段名 string, 包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或包含(字段名, 包含值))
}

// X条件或NULL值 构建 `columns[0] IS NULL OR columns[1] IS NULL ...` 的 `OR` 条件语句。
// 参考 WhereBuilder.X条件或NULL值。
// md5:66907cf860f22eff
func (m *Model) X条件或NULL值(字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或NULL值(字段名...))
}

// X条件或取范围以外 用于构建 `column NOT BETWEEN min AND max` 的 SQL 语句，在 `OR` 条件下。
// 参见 WhereBuilder.X条件或取范围以外 的用法。
// md5:e040a9f04b492725
func (m *Model) X条件或取范围以外(字段名 string, 最小值, 最大值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或取范围以外(字段名, 最小值, 最大值))
}

// X条件或模糊匹配以外 用于构建在`OR`条件下的 `column NOT LIKE 'like'` 语句。
// 参考 WhereBuilder.X条件或模糊匹配以外。
// md5:588ea3675853468b
func (m *Model) X条件或模糊匹配以外(字段名 string, 通配符条件值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或模糊匹配以外(字段名, 通配符条件值))
}

// X条件或不等于构建`column != value`语句。
// 参见WhereBuilder.X条件或不等于。
// md5:076a864671142e49
func (m *Model) X条件或不等于(字段名 string, 值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或不等于(字段名, 值))
}

// X条件或不包含构建`column NOT IN (in)`语句。
// 参见WhereBuilder.X条件或不包含。
// md5:22915d1ba70db001
func (m *Model) X条件或不包含(字段名 string, 不包含值 interface{}) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或不包含(字段名, 不包含值))
}

// X条件或非Null 构建在 `OR` 条件下的 `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` 语句。
// 参见 WhereBuilder.X条件或非Null 的用法。
// md5:5645594c534afc8e
func (m *Model) X条件或非Null(字段名 ...string) *Model {
	return m.callWhereBuilder(m.whereBuilder.X条件或非Null(字段名...))
}
