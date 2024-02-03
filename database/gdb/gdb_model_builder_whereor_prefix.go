// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

// WhereOrPrefix 的行为类似于 WhereOr，但是它会在 where 语句中的每个字段前添加指定的前缀。
// 示例：
// WhereOrPrefix("order", "status", "paid")                        => WHERE xxx OR (`order`.`status`='已支付')
// WhereOrPrefix("order", struct{Status:"paid", Channel:"bank"}) => WHERE xxx OR (`order`.`status`='已支付' AND `order`.`channel`='银行')
func (b *WhereBuilder) WhereOrPrefix(prefix string, where interface{}, args ...interface{}) *WhereBuilder {
	where, args = b.convertWhereBuilder(where, args)

	builder := b.getBuilder()
	builder.whereHolder = append(builder.whereHolder, WhereHolder{
		Type:     whereHolderTypeDefault,
		Operator: whereHolderOperatorOr,
		Where:    where,
		Args:     args,
		Prefix:   prefix,
	})
	return builder
}

// WhereOrPrefixNot 在“OR”条件下构建“prefix.column != value”语句。
func (b *WhereBuilder) WhereOrPrefixNot(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s != ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixLT 在`OR`条件下构建 `prefix.column < value` 语句。
func (b *WhereBuilder) WhereOrPrefixLT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s < ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixLTE 在`OR`条件下构建 `prefix.column <= value` 语句。
func (b *WhereBuilder) WhereOrPrefixLTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s <= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixGT 在“OR”条件下构建 `prefix.column > value` 语句。
func (b *WhereBuilder) WhereOrPrefixGT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s > ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixGTE 在“OR”条件下构建“prefix.column >= value”语句。
func (b *WhereBuilder) WhereOrPrefixGTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s >= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixBetween 在`OR`条件下构建 `prefix.column BETWEEN min AND max` 语句。
func (b *WhereBuilder) WhereOrPrefixBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WhereOrPrefixLike 在`OR`条件下构建`prefix.column LIKE 'like'`语句。
func (b *WhereBuilder) WhereOrPrefixLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WhereOrPrefixIn 在“OR”条件下构建 `prefix.column IN (in)` 语句。
func (b *WhereBuilder) WhereOrPrefixIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s.%s IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WhereOrPrefixNull 在`OR`条件下构建 `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` 语句。
func (b *WhereBuilder) WhereOrPrefixNull(prefix string, columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.WhereOrf(`%s.%s IS NULL`, b.model.QuoteWord(prefix), b.model.QuoteWord(column))
	}
	return builder
}

// WhereOrPrefixNotBetween 在“OR”条件下构建 `prefix.column NOT BETWEEN min AND max` 语句。
func (b *WhereBuilder) WhereOrPrefixNotBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s NOT BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WhereOrPrefixNotLike 在`OR`条件下构建 `prefix.column NOT LIKE 'like'` 语句。
func (b *WhereBuilder) WhereOrPrefixNotLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s NOT LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WhereOrPrefixNotIn 用于构建 `prefix.column NOT IN (in)` 语句。
func (b *WhereBuilder) WhereOrPrefixNotIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s.%s NOT IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WhereOrPrefixNotNull 用于构建 `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` 形式的 OR 条件语句。
func (b *WhereBuilder) WhereOrPrefixNotNull(prefix string, columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.WhereOrf(`%s.%s IS NOT NULL`, b.model.QuoteWord(prefix), b.model.QuoteWord(column))
	}
	return builder
}
