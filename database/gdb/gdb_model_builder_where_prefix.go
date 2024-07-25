// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gdb

// WherePrefix 的功能类似于 Where，但它会在 where 语句中的每个字段前添加一个前缀。
// 例如：
// WherePrefix("order", "status", "paid")                        => WHERE `order`.`status`='paid'
// WherePrefix("order", struct{Status:"paid", "channel":"bank"}) => WHERE `order`.`status`='paid' AND `order`.`channel`='bank'
//
// 这段注释的中文翻译为：
//
// WherePrefix 的行为与 Where 相似，但它会在 where 子句里的每个字段前加上一个前缀。
// 例如：
// WherePrefix("order", "status", "paid")                        => 生成 WHERE `order`.`status`='paid'
// WherePrefix("order", struct{Status:"paid", "channel":"bank"}) => 生成 WHERE `order`.`status`='paid' AND `order`.`channel`='bank' md5:062302edb484784b
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

// WherePrefixLT 构建 `prefix.column < value` 语句。 md5:de5cb5259ef84499
func (b *WhereBuilder) WherePrefixLT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s < ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixLTE构建`prefix.column <= value`语句。 md5:1c5d93f173a39b03
func (b *WhereBuilder) WherePrefixLTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s <= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixGT构建`prefix.column > value`语句。 md5:61d5cbbb9f5422fd
func (b *WhereBuilder) WherePrefixGT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s > ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixGTE 生成 `prefix.column >= value` 的语句。 md5:1b581ea600e215e7
func (b *WhereBuilder) WherePrefixGTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s >= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixBetween 构建 `prefix.column BETWEEN min AND max` 语句。 md5:e6176c55b8a31575
func (b *WhereBuilder) WherePrefixBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WherePrefixLike构建`prefix.column LIKE like`语句。 md5:baf08eac5c7dc2aa
func (b *WhereBuilder) WherePrefixLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WherePrefixIn 构建 `prefix.column IN (in)` 语句。 md5:fd691f634711ba7f
func (b *WhereBuilder) WherePrefixIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s.%s IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WherePrefixNull 构建 `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` 语句。 md5:ac08bde96048fdce
func (b *WhereBuilder) WherePrefixNull(prefix string, columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.Wheref(`%s.%s IS NULL`, b.model.QuoteWord(prefix), b.model.QuoteWord(column))
	}
	return builder
}

// WherePrefixNotBetween 构建 `prefix.column NOT BETWEEN min AND max` 语句。 md5:a16703b511af05c3
func (b *WhereBuilder) WherePrefixNotBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s NOT BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WherePrefixNotLike构建`prefix.column NOT LIKE like`语句。 md5:083bd1d45c368a83
func (b *WhereBuilder) WherePrefixNotLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s NOT LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WherePrefixNot构建`prefix.column != value`语句。 md5:c1366e00cd0da49e
func (b *WhereBuilder) WherePrefixNot(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s != ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixNotIn 用于构建 `prefix.column NOT IN (in)` 的SQL语句。 md5:3b790678c07a51fd
func (b *WhereBuilder) WherePrefixNotIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s.%s NOT IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WherePrefixNotNull 构建 `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` 语句。 md5:d5a307a7c3004dda
func (b *WhereBuilder) WherePrefixNotNull(prefix string, columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.Wheref(`%s.%s IS NOT NULL`, b.model.QuoteWord(prefix), b.model.QuoteWord(column))
	}
	return builder
}
