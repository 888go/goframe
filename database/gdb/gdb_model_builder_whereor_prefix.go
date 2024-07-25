// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gdb

// WhereOrPrefix 的功能类似于 WhereOr，但它会在 WHERE 子句中的每个字段前添加一个前缀。
// 例如：
// WhereOrPrefix("order", "status", "paid")                        => WHERE xxx OR (`order`.`status`='paid')
// WhereOrPrefix("order", struct{Status:"paid", "channel":"bank"}) => WHERE xxx OR (`order`.`status`='paid' AND `order`.`channel`='bank')
//
// 这意味着 WhereOrPrefix 函数允许你在一个 WHERE 子句中指定多个条件，并且自动为这些条件的字段名加上一个指定的前缀，以便清晰地指向某个表或结构。它可以处理单个字段值的情况，也可以处理包含多个键值对的结构体，以生成更复杂的逻辑组合。 md5:2358d43541f472f5
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

// WhereOrPrefixNot 在 OR 条件中构建 `prefix.column != value` 语句。 md5:385a9f9fb58b8fc3
func (b *WhereBuilder) WhereOrPrefixNot(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s != ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixLT在"OR"条件下构建`prefix.column < value`语句。 md5:c1a6baf94f553043
func (b *WhereBuilder) WhereOrPrefixLT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s < ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixLTE 在“OR”条件下构建 `prefix.column <= value` 语句。 md5:77072877b38f04a8
func (b *WhereBuilder) WhereOrPrefixLTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s <= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixGT 在 `OR` 条件下构建 `prefix.column > value` 的语句。 md5:d34b5bdc0e6b2fa8
func (b *WhereBuilder) WhereOrPrefixGT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s > ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixGTE 在 OR 条件中构建 `prefix.column >= value` 语句。 md5:d652ca0304ac835e
func (b *WhereBuilder) WhereOrPrefixGTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s >= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixBetween在`OR`条件下构建`prefix.column BETWEEN min AND max`语句。 md5:d7adaf273fa5681b
func (b *WhereBuilder) WhereOrPrefixBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WhereOrPrefixLike在`OR`条件下构建`prefix.column LIKE 'like'`语句。 md5:c975b47e3a5cc2c1
func (b *WhereBuilder) WhereOrPrefixLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WhereOrPrefixIn 用于构建 `prefix.column IN (in)` 形式的 SQL 语句，在 `OR` 条件下。 md5:18e0cf5cc971267d
func (b *WhereBuilder) WhereOrPrefixIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s.%s IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WhereOrPrefixNull 在"OR"条件中构建`prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...`语句。 md5:facf88eb72f3d299
func (b *WhereBuilder) WhereOrPrefixNull(prefix string, columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.WhereOrf(`%s.%s IS NULL`, b.model.QuoteWord(prefix), b.model.QuoteWord(column))
	}
	return builder
}

// WhereOrPrefixNotBetween在`OR`条件下构建`prefix.column NOT BETWEEN min AND max`语句。 md5:15259f135308893b
func (b *WhereBuilder) WhereOrPrefixNotBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s NOT BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WhereOrPrefixNotLike 在 `OR` 条件下构建 `prefix.column NOT LIKE 'like'` 语句。 md5:2785cbc79e811104
func (b *WhereBuilder) WhereOrPrefixNotLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s NOT LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WhereOrPrefixNotIn 用于构建 `prefix.column NOT IN (in)` 的SQL语句。 md5:bd296110bb5635a1
func (b *WhereBuilder) WhereOrPrefixNotIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s.%s NOT IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WhereOrPrefixNotNull 在`OR`条件中构建`prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...`语句。 md5:9ecd3bffabf47cb7
func (b *WhereBuilder) WhereOrPrefixNotNull(prefix string, columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.WhereOrf(`%s.%s IS NOT NULL`, b.model.QuoteWord(prefix), b.model.QuoteWord(column))
	}
	return builder
}
