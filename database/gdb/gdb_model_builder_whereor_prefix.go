// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

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

// WhereOrPrefixNot builds `prefix.column != value` statement in `OR` conditions.
// ff:条件或不等于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:值
func (b *WhereBuilder) WhereOrPrefixNot(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s != ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixLT builds `prefix.column < value` statement in `OR` conditions.
// ff:条件或小于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:比较值
func (b *WhereBuilder) WhereOrPrefixLT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s < ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixLTE builds `prefix.column <= value` statement in `OR` conditions.
// ff:条件或小于等于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:比较值
func (b *WhereBuilder) WhereOrPrefixLTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s <= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixGT builds `prefix.column > value` statement in `OR` conditions.
// ff:条件或大于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:比较值
func (b *WhereBuilder) WhereOrPrefixGT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s > ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixGTE builds `prefix.column >= value` statement in `OR` conditions.
// ff:条件或大于等于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:比较值
func (b *WhereBuilder) WhereOrPrefixGTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s >= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WhereOrPrefixBetween builds `prefix.column BETWEEN min AND max` statement in `OR` conditions.
// ff:条件或取范围并带前缀
// b:
// prefix:字段前缀
// column:字段名
// min:最小值
// max:最大值
func (b *WhereBuilder) WhereOrPrefixBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WhereOrPrefixLike builds `prefix.column LIKE 'like'` statement in `OR` conditions.
// ff:条件或模糊匹配并带前缀
// b:
// prefix:字段前缀
// column:字段名
// like:通配符条件值
func (b *WhereBuilder) WhereOrPrefixLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WhereOrPrefixIn builds `prefix.column IN (in)` statement in `OR` conditions.
// ff:条件或包含并带前缀
// b:
// prefix:字段前缀
// column:字段名
// in:包含值
func (b *WhereBuilder) WhereOrPrefixIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s.%s IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WhereOrPrefixNull builds `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` statement in `OR` conditions.
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

// WhereOrPrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement in `OR` conditions.
// ff:条件或取范围以外并带前缀
// b:
// prefix:字段前缀
// column:字段名
// min:最小值
// max:最大值
func (b *WhereBuilder) WhereOrPrefixNotBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s NOT BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WhereOrPrefixNotLike builds `prefix.column NOT LIKE 'like'` statement in `OR` conditions.
// ff:条件或模糊匹配以外并带前缀
// b:
// prefix:字段前缀
// column:字段名
// like:通配符条件值
func (b *WhereBuilder) WhereOrPrefixNotLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.WhereOrf(`%s.%s NOT LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WhereOrPrefixNotIn builds `prefix.column NOT IN (in)` statement.
// ff:条件或不包含并带前缀
// b:
// prefix:字段前缀
// column:字段名
// in:不包含值
func (b *WhereBuilder) WhereOrPrefixNotIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s.%s NOT IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WhereOrPrefixNotNull builds `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` statement in `OR` conditions.
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
