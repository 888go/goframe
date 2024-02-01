// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

// WherePrefix 的行为类似于 Where，但它会在 where 语句中的每个字段前添加前缀。
// 示例：
// WherePrefix("order", "status", "paid")                        => WHERE `order`.`status`='paid'
// WherePrefix("order", struct{Status:"paid", Channel:"bank"}) => WHERE `order`.`status`='paid' AND `order`.`channel`='bank'
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

// WherePrefixLT 用于构建 `prefix.column < value` 的语句。
func (b *WhereBuilder) WherePrefixLT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s < ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixLTE 用于构建 `prefix.column <= value` 的语句。
func (b *WhereBuilder) WherePrefixLTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s <= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixGT 用于构建 `prefix.column > value` 的表达式语句。
func (b *WhereBuilder) WherePrefixGT(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s > ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixGTE 生成 `prefix.column >= value` 语句。
func (b *WhereBuilder) WherePrefixGTE(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s >= ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixBetween 用于构建 `prefix.column BETWEEN min AND max` 的语句。
func (b *WhereBuilder) WherePrefixBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WherePrefixLike 用于构建 `prefix.column LIKE like` 语句。
func (b *WhereBuilder) WherePrefixLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WherePrefixIn 用于构建 `prefix.column IN (in)` 语句。
func (b *WhereBuilder) WherePrefixIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s.%s IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WherePrefixNull 用于构建如 `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` 形式的语句。
func (b *WhereBuilder) WherePrefixNull(prefix string, columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.Wheref(`%s.%s IS NULL`, b.model.QuoteWord(prefix), b.model.QuoteWord(column))
	}
	return builder
}

// WherePrefixNotBetween 用于构建 `prefix.column NOT BETWEEN min AND max` 的表达式语句。
func (b *WhereBuilder) WherePrefixNotBetween(prefix string, column string, min, max interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s NOT BETWEEN ? AND ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), min, max)
}

// WherePrefixNotLike 用于构建 `prefix.column NOT LIKE like` 语句。
func (b *WhereBuilder) WherePrefixNotLike(prefix string, column string, like interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s NOT LIKE ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), like)
}

// WherePrefixNot 用于构建 `prefix.column != value` 的表达式语句。
func (b *WhereBuilder) WherePrefixNot(prefix string, column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s.%s != ?`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), value)
}

// WherePrefixNotIn 用于构建 `prefix.column NOT IN (in)` 语句。
func (b *WhereBuilder) WherePrefixNotIn(prefix string, column string, in interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s.%s NOT IN (?)`, b.model.QuoteWord(prefix), b.model.QuoteWord(column), in)
}

// WherePrefixNotNull 用于构建 `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` 语句。
func (b *WhereBuilder) WherePrefixNotNull(prefix string, columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.Wheref(`%s.%s IS NOT NULL`, b.model.QuoteWord(prefix), b.model.QuoteWord(column))
	}
	return builder
}
