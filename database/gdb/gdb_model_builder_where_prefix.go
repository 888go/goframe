// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

// WherePrefix 的行为类似于 Where，但它会在 where 语句中的每个字段前添加前缀。
// 示例：
// WherePrefix("order", "status", "paid")                        => WHERE `order`.`status`='paid'
// WherePrefix("order", struct{Status:"paid", Channel:"bank"}) => WHERE `order`.`status`='paid' AND `order`.`channel`='bank'
func (b *X组合条件) X条件带前缀(字段前缀 string, 条件 interface{}, 参数 ...interface{}) *X组合条件 {
	条件, 参数 = b.convertWhereBuilder(条件, 参数)

	builder := b.getBuilder()
	if builder.whereHolder == nil {
		builder.whereHolder = make([]WhereHolder, 0)
	}
	builder.whereHolder = append(builder.whereHolder, WhereHolder{
		Type:     whereHolderTypeDefault,
		Operator: whereHolderOperatorWhere,
		Where:    条件,
		Args:     参数,
		Prefix:   字段前缀,
	})
	return builder
}

// WherePrefixLT 用于构建 `prefix.column < value` 的语句。
func (b *X组合条件) X条件小于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *X组合条件 {
	return b.X条件格式化(`%s.%s < ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 比较值)
}

// WherePrefixLTE 用于构建 `prefix.column <= value` 的语句。
func (b *X组合条件) X条件小于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *X组合条件 {
	return b.X条件格式化(`%s.%s <= ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 比较值)
}

// WherePrefixGT 用于构建 `prefix.column > value` 的表达式语句。
func (b *X组合条件) X条件大于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *X组合条件 {
	return b.X条件格式化(`%s.%s > ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 比较值)
}

// WherePrefixGTE 生成 `prefix.column >= value` 语句。
func (b *X组合条件) X条件大于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *X组合条件 {
	return b.X条件格式化(`%s.%s >= ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 比较值)
}

// WherePrefixBetween 用于构建 `prefix.column BETWEEN min AND max` 的语句。
func (b *X组合条件) X条件取范围并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *X组合条件 {
	return b.X条件格式化(`%s.%s BETWEEN ? AND ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 最小值, 最大值)
}

// WherePrefixLike 用于构建 `prefix.column LIKE like` 语句。
func (b *X组合条件) X条件模糊匹配并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *X组合条件 {
	return b.X条件格式化(`%s.%s LIKE ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 通配符条件值)
}

// WherePrefixIn 用于构建 `prefix.column IN (in)` 语句。
func (b *X组合条件) X条件包含并带前缀(字段前缀 string, 字段名 string, 包含值 interface{}) *X组合条件 {
	return b.doWherefType(whereHolderTypeIn, `%s.%s IN (?)`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 包含值)
}

// WherePrefixNull 用于构建如 `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` 形式的语句。
func (b *X组合条件) X条件NULL值并带前缀(字段前缀 string, 字段名 ...string) *X组合条件 {
	builder := b
	for _, column := range 字段名 {
		builder = builder.X条件格式化(`%s.%s IS NULL`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(column))
	}
	return builder
}

// WherePrefixNotBetween 用于构建 `prefix.column NOT BETWEEN min AND max` 的表达式语句。
func (b *X组合条件) X条件取范围以外并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *X组合条件 {
	return b.X条件格式化(`%s.%s NOT BETWEEN ? AND ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 最小值, 最大值)
}

// WherePrefixNotLike 用于构建 `prefix.column NOT LIKE like` 语句。
func (b *X组合条件) X条件模糊匹配以外并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *X组合条件 {
	return b.X条件格式化(`%s.%s NOT LIKE ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 通配符条件值)
}

// WherePrefixNot 用于构建 `prefix.column != value` 的表达式语句。
func (b *X组合条件) X条件不等于并带前缀(字段前缀 string, 字段名 string, 值 interface{}) *X组合条件 {
	return b.X条件格式化(`%s.%s != ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 值)
}

// WherePrefixNotIn 用于构建 `prefix.column NOT IN (in)` 语句。
func (b *X组合条件) X条件不包含并带前缀(字段前缀 string, 字段名 string, 不包含值 interface{}) *X组合条件 {
	return b.doWherefType(whereHolderTypeIn, `%s.%s NOT IN (?)`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 不包含值)
}

// WherePrefixNotNull 用于构建 `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` 语句。
func (b *X组合条件) X条件非Null并带前缀(字段前缀 string, 字段名 ...string) *X组合条件 {
	builder := b
	for _, column := range 字段名 {
		builder = builder.X条件格式化(`%s.%s IS NOT NULL`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(column))
	}
	return builder
}
