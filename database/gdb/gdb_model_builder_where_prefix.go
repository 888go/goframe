// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

// WherePrefix performs as Where, but it adds prefix to each field in where statement.
// WherePrefix("order", "status", "paid")                        => WHERE `order`.`status`='paid'
// WherePrefix("order", struct{Status:"paid", "channel":"bank"}) => WHERE `order`.`status`='paid' AND `order`.`channel`='bank'
// ff:条件带前缀
// b:
// prefix:字段前缀
// where:条件
// args:参数
func (b *WhereBuilder) WherePrefix(prefix string, where interface{}, args ...interface{}) *WhereBuilder {
	where, args = b.convertWhereBuilder(where, args)

	builder := b.getBuilder()
	if builder.whereHolder == nil {
		builder.whereHolder = make([]WhereHolder, 0)
	}
	builder.whereHolder = append(builder.whereHolder, WhereHolder{
		Type:     whereHolderTypeDefault,
		Operator: whereHolderOperatorWhere,
		Where:    where,
		Args:     args,
		Prefix:   prefix,
	})
	return builder
}

// WherePrefixLT builds `prefix.column < value` statement.
// ff:条件小于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:比较值
func (b *WhereBuilder) WherePrefixLT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s < ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixLTE builds `prefix.column <= value` statement.
// ff:条件小于等于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:比较值
func (b *WhereBuilder) WherePrefixLTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s <= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixGT builds `prefix.column > value` statement.
// ff:条件大于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:比较值
func (b *WhereBuilder) WherePrefixGT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s > ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixGTE builds `prefix.column >= value` statement.
// ff:条件大于等于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:比较值
func (b *WhereBuilder) WherePrefixGTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s >= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixBetween builds `prefix.column BETWEEN min AND max` statement.
// ff:条件取范围并带前缀
// b:
// prefix:字段前缀
// column:字段名
// min:最小值
// max:最大值
func (b *WhereBuilder) WherePrefixBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WherePrefixLike builds `prefix.column LIKE like` statement.
// ff:条件模糊匹配并带前缀
// b:
// prefix:字段前缀
// column:字段名
// like:通配符条件值
func (b *WhereBuilder) WherePrefixLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WherePrefixIn builds `prefix.column IN (in)` statement.
// ff:条件包含并带前缀
// b:
// prefix:字段前缀
// column:字段名
// in:包含值
func (b *WhereBuilder) WherePrefixIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s.%s IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WherePrefixNull builds `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` statement.
// ff:条件NULL值并带前缀
// b:
// prefix:字段前缀
// columns:字段名
func (b *WhereBuilder) WherePrefixNull(prefix string, columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.Wheref(`%s.%s IS NULL`, b.model.QuoteWord(prefix), b.model.QuoteWord(column))
	}
	return builder
}

// WherePrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement.
// ff:条件取范围以外并带前缀
// b:
// prefix:字段前缀
// column:字段名
// min:最小值
// max:最大值
func (b *WhereBuilder) WherePrefixNotBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s NOT BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WherePrefixNotLike builds `prefix.column NOT LIKE like` statement.
// ff:条件模糊匹配以外并带前缀
// b:
// prefix:字段前缀
// column:字段名
// like:通配符条件值
func (b *WhereBuilder) WherePrefixNotLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s NOT LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WherePrefixNot builds `prefix.column != value` statement.
// ff:条件不等于并带前缀
// b:
// prefix:字段前缀
// column:字段名
// value:值
func (b *WhereBuilder) WherePrefixNot(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s != ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixNotIn builds `prefix.column NOT IN (in)` statement.
// ff:条件不包含并带前缀
// b:
// prefix:字段前缀
// column:字段名
// in:不包含值
func (b *WhereBuilder) WherePrefixNotIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s.%s NOT IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WherePrefixNotNull builds `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` statement.
// ff:条件非Null并带前缀
// b:
// prefix:字段前缀
// columns:字段名
func (b *WhereBuilder) WherePrefixNotNull(prefix string, columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.Wheref(`%s.%s IS NOT NULL`, b.model.QuoteWord(prefix), b.model.QuoteWord(column))
	}
	return builder
}
