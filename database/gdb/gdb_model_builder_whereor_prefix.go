// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

// WhereOrPrefix performs as WhereOr, but it adds prefix to each field in where statement.
// WhereOrPrefix("order", "status", "paid")                        => WHERE xxx OR (`order`.`status`='paid')
// WhereOrPrefix("order", struct{Status:"paid", "channel":"bank"}) => WHERE xxx OR (`order`.`status`='paid' AND `order`.`channel`='bank')
// ff:条件或并带前缀
// b:
// prefix:字段前缀
// where:条件
// args:参数
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
// ff:条件或不等于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:值
func (b *WhereBuilder) WhereOrPrefixNot(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s != ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixLT在"OR"条件下构建`prefix.column < value`语句。 md5:c1a6baf94f553043
// ff:条件或小于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:比较值
func (b *WhereBuilder) WhereOrPrefixLT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s < ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixLTE 在“OR”条件下构建 `prefix.column <= value` 语句。 md5:77072877b38f04a8
// ff:条件或小于等于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:比较值
func (b *WhereBuilder) WhereOrPrefixLTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s <= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixGT 在 `OR` 条件下构建 `prefix.column > value` 的语句。 md5:d34b5bdc0e6b2fa8
// ff:条件或大于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:比较值
func (b *WhereBuilder) WhereOrPrefixGT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s > ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixGTE 在 OR 条件中构建 `prefix.column >= value` 语句。 md5:d652ca0304ac835e
// ff:条件或大于等于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:比较值
func (b *WhereBuilder) WhereOrPrefixGTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s >= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixBetween在`OR`条件下构建`prefix.column BETWEEN min AND max`语句。 md5:d7adaf273fa5681b
// ff:条件或取范围并带前缀
// b:
// prefix:字段前缀
// column:字段名
// min:最小值
// max:最大值
func (b *WhereBuilder) WhereOrPrefixBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WhereOrPrefixLike在`OR`条件下构建`prefix.column LIKE 'like'`语句。 md5:c975b47e3a5cc2c1
// ff:条件或模糊匹配并带前缀
// b:
// prefix:字段前缀
// column:字段名
// like:通配符条件值
func (b *WhereBuilder) WhereOrPrefixLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WhereOrPrefixIn 用于构建 `prefix.column IN (in)` 形式的 SQL 语句，在 `OR` 条件下。 md5:18e0cf5cc971267d
// ff:条件或包含并带前缀
// b:
// prefix:字段前缀
// column:字段名
// in:包含值
func (b *WhereBuilder) WhereOrPrefixIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s.%s IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WhereOrPrefixNull 在"OR"条件中构建`prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...`语句。 md5:facf88eb72f3d299
// ff:条件或NULL值并带前缀
// b:
// prefix:字段前缀
// columns:字段名
func (b *WhereBuilder) WhereOrPrefixNull(prefix string, columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.WhereOrf(`%s.%s IS NULL`, b.model.QuoteWord(prefix), b.model.QuoteWord(column))
	}
	return builder
}

// WhereOrPrefixNotBetween在`OR`条件下构建`prefix.column NOT BETWEEN min AND max`语句。 md5:15259f135308893b
// ff:条件或取范围以外并带前缀
// b:
// prefix:字段前缀
// column:字段名
// min:最小值
// max:最大值
func (b *WhereBuilder) WhereOrPrefixNotBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s NOT BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WhereOrPrefixNotLike 在 `OR` 条件下构建 `prefix.column NOT LIKE 'like'` 语句。 md5:2785cbc79e811104
// ff:条件或模糊匹配以外并带前缀
// b:
// prefix:字段前缀
// column:字段名
// like:通配符条件值
func (b *WhereBuilder) WhereOrPrefixNotLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s NOT LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WhereOrPrefixNotIn 用于构建 `prefix.column NOT IN (in)` 的SQL语句。 md5:bd296110bb5635a1
// ff:条件或不包含并带前缀
// b:
// prefix:字段前缀
// column:字段名
// in:不包含值
func (b *WhereBuilder) WhereOrPrefixNotIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s.%s NOT IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WhereOrPrefixNotNull 在`OR`条件中构建`prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...`语句。 md5:9ecd3bffabf47cb7
// ff:条件或非Null并带前缀
// b:
// prefix:字段前缀
// columns:字段名
func (b *WhereBuilder) WhereOrPrefixNotNull(prefix string, columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.WhereOrf(`%s.%s IS NOT NULL`, b.model.QuoteWord(prefix), b.model.QuoteWord(column))
	}
	return builder
}
