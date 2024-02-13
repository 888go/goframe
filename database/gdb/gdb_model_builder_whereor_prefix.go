// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

// WhereOrPrefix 的行为类似于 WhereOr，但是它会在 where 语句中的每个字段前添加指定的前缀。
// 示例：
// WhereOrPrefix("order", "status", "paid")                        => WHERE xxx OR (`order`.`status`='已支付')
// WhereOrPrefix("order", struct{Status:"paid", Channel:"bank"}) => WHERE xxx OR (`order`.`status`='已支付' AND `order`.`channel`='银行')
func (b *X组合条件) X条件或并带前缀(字段前缀 string, 条件 interface{}, 参数 ...interface{}) *X组合条件 {
	条件, 参数 = b.convertWhereBuilder(条件, 参数)

	builder := b.getBuilder()
	builder.whereHolder = append(builder.whereHolder, WhereHolder{
		Type:     whereHolderTypeDefault,
		Operator: whereHolderOperatorOr,
		Where:    条件,
		Args:     参数,
		Prefix:   字段前缀,
	})
	return builder
}

// WhereOrPrefixNot 在“OR”条件下构建“prefix.column != value”语句。
func (b *X组合条件) X条件或不等于并带前缀(字段前缀 string, 字段名 string, 值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s.%s != ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 值)
}

// WhereOrPrefixLT 在`OR`条件下构建 `prefix.column < value` 语句。
func (b *X组合条件) X条件或小于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s.%s < ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 比较值)
}

// WhereOrPrefixLTE 在`OR`条件下构建 `prefix.column <= value` 语句。
func (b *X组合条件) X条件或小于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s.%s <= ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 比较值)
}

// WhereOrPrefixGT 在“OR”条件下构建 `prefix.column > value` 语句。
func (b *X组合条件) X条件或大于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s.%s > ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 比较值)
}

// WhereOrPrefixGTE 在“OR”条件下构建“prefix.column >= value”语句。
func (b *X组合条件) X条件或大于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s.%s >= ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 比较值)
}

// WhereOrPrefixBetween 在`OR`条件下构建 `prefix.column BETWEEN min AND max` 语句。
func (b *X组合条件) X条件或取范围并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s.%s BETWEEN ? AND ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 最小值, 最大值)
}

// WhereOrPrefixLike 在`OR`条件下构建`prefix.column LIKE 'like'`语句。
func (b *X组合条件) X条件或模糊匹配并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s.%s LIKE ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 通配符条件值)
}

// WhereOrPrefixIn 在“OR”条件下构建 `prefix.column IN (in)` 语句。
func (b *X组合条件) X条件或包含并带前缀(字段前缀 string, 字段名 string, 包含值 interface{}) *X组合条件 {
	return b.doWhereOrfType(whereHolderTypeIn, `%s.%s IN (?)`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 包含值)
}

// WhereOrPrefixNull 在`OR`条件下构建 `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` 语句。
func (b *X组合条件) X条件或NULL值并带前缀(字段前缀 string, 字段名 ...string) *X组合条件 {
	builder := b
	for _, column := range 字段名 {
		builder = builder.X条件或格式化(`%s.%s IS NULL`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(column))
	}
	return builder
}

// WhereOrPrefixNotBetween 在“OR”条件下构建 `prefix.column NOT BETWEEN min AND max` 语句。
func (b *X组合条件) X条件或取范围以外并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s.%s NOT BETWEEN ? AND ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 最小值, 最大值)
}

// WhereOrPrefixNotLike 在`OR`条件下构建 `prefix.column NOT LIKE 'like'` 语句。
func (b *X组合条件) X条件或模糊匹配以外并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s.%s NOT LIKE ?`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 通配符条件值)
}

// WhereOrPrefixNotIn 用于构建 `prefix.column NOT IN (in)` 语句。
func (b *X组合条件) X条件或不包含并带前缀(字段前缀 string, 字段名 string, 不包含值 interface{}) *X组合条件 {
	return b.doWhereOrfType(whereHolderTypeIn, `%s.%s NOT IN (?)`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(字段名), 不包含值)
}

// WhereOrPrefixNotNull 用于构建 `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` 形式的 OR 条件语句。
func (b *X组合条件) X条件或非Null并带前缀(字段前缀 string, 字段名 ...string) *X组合条件 {
	builder := b
	for _, column := range 字段名 {
		builder = builder.X条件或格式化(`%s.%s IS NOT NULL`, b.model.底层QuoteWord(字段前缀), b.model.底层QuoteWord(column))
	}
	return builder
}
