// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gdb

import (
	"fmt"

	"github.com/gogf/gf/v2/text/gstr"
)

// WhereOr 向 WHERE 语句中添加“OR”条件。 md5:753c32f428b02541
func (b *WhereBuilder) doWhereOrType(t string, where interface{}, args ...interface{}) *WhereBuilder {
	where, args = b.convertWhereBuilder(where, args)

	builder := b.getBuilder()
	if builder.whereHolder == nil {
		builder.whereHolder = make([]WhereHolder, 0)
	}
	builder.whereHolder = append(builder.whereHolder, WhereHolder{
		Type:     t,
		Operator: whereHolderOperatorOr,
		Where:    where,
		Args:     args,
	})
	return builder
}

// WhereOrf 使用fmt.Sprintf和参数构建`OR`条件字符串。 md5:aa04236f081a2885
func (b *WhereBuilder) doWhereOrfType(t string, format string, args ...interface{}) *WhereBuilder {
	var (
		placeHolderCount = gstr.Count(format, "?")
		conditionStr     = fmt.Sprintf(format, args[:len(args)-placeHolderCount]...)
	)
	return b.doWhereOrType(t, conditionStr, args[len(args)-placeHolderCount:]...)
}

// WhereOr 向 WHERE 语句中添加“OR”条件。 md5:753c32f428b02541
func (b *WhereBuilder) WhereOr(where interface{}, args ...interface{}) *WhereBuilder {
	return b.doWhereOrType(``, where, args...)
}

// WhereOrf 使用fmt.Sprintf和参数构建`OR`条件字符串。 md5:aa04236f081a2885
// Eg:
// WhereOrf(`amount<? and status=%s`, "paid", 100)  => WHERE xxx OR `amount`<100 and status='paid'
// WhereOrf(`amount<%d and status=%s`, 100, "paid") => WHERE xxx OR `amount`<100 and status='paid'
func (b *WhereBuilder) WhereOrf(format string, args ...interface{}) *WhereBuilder {
	return b.doWhereOrfType(``, format, args...)
}

// WhereOrNot在`OR`条件下构建`column != value`语句。 md5:adc6d63e61bf279f
func (b *WhereBuilder) WhereOrNot(column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s != ?`, column, value)
}

// WhereOrLT 在 `OR` 条件下构建 `column < value` 的语句。 md5:5517b3812e2c8e8b
func (b *WhereBuilder) WhereOrLT(column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s < ?`, column, value)
}

// WhereOrLTE 在 OR 条件中构建 `column <= value` 语句。 md5:3b0287bd1f8030ce
func (b *WhereBuilder) WhereOrLTE(column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s <= ?`, column, value)
}

// WhereOrGT在`OR`条件下构建`column > value`语句。 md5:2289d39bb82e521f
func (b *WhereBuilder) WhereOrGT(column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s > ?`, column, value)
}

// WhereOrGTE在`OR`条件下构建`column >= value`语句。 md5:e178dd8cfc5661e5
func (b *WhereBuilder) WhereOrGTE(column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s >= ?`, column, value)
}

// WhereOrBetween 用于构建 `column BETWEEN min AND max` 语句，并在 `OR` 条件下使用。 md5:90f98622a1fd5981
func (b *WhereBuilder) WhereOrBetween(column string, min, max interface{}) *WhereBuilder {
	return b.WhereOrf(`%s BETWEEN ? AND ?`, b.model.QuoteWord(column), min, max)
}

// WhereOrLike 在 `OR` 条件中构建 `column LIKE 'like'` 语句。 md5:7a2d37411752fb51
func (b *WhereBuilder) WhereOrLike(column string, like interface{}) *WhereBuilder {
	return b.WhereOrf(`%s LIKE ?`, b.model.QuoteWord(column), like)
}

// WhereOrIn 在`OR`条件下构建`column IN (in)`语句。 md5:4bb93b5ae9a5e887
func (b *WhereBuilder) WhereOrIn(column string, in interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s IN (?)`, b.model.QuoteWord(column), in)
}

// WhereOrNull 在 `OR` 条件下构建 `columns[0] IS NULL OR columns[1] IS NULL ...` 语句。 md5:08d38a60dc594441
func (b *WhereBuilder) WhereOrNull(columns ...string) *WhereBuilder {
	var builder *WhereBuilder
	for _, column := range columns {
		builder = b.WhereOrf(`%s IS NULL`, b.model.QuoteWord(column))
	}
	return builder
}

// WhereOrNotBetween 用于构建在 `OR` 条件下的 `column NOT BETWEEN min AND max` 语句。 md5:f20408e0126bbbab
func (b *WhereBuilder) WhereOrNotBetween(column string, min, max interface{}) *WhereBuilder {
	return b.WhereOrf(`%s NOT BETWEEN ? AND ?`, b.model.QuoteWord(column), min, max)
}

// WhereOrNotLike 在 OR 条件中构建 `column NOT LIKE like` 语句。 md5:751e840816119632
func (b *WhereBuilder) WhereOrNotLike(column string, like interface{}) *WhereBuilder {
	return b.WhereOrf(`%s NOT LIKE ?`, b.model.QuoteWord(column), like)
}

// WhereOrNotIn构建`column NOT IN (in)`语句。 md5:433fd8a0f224fc24
func (b *WhereBuilder) WhereOrNotIn(column string, in interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s NOT IN (?)`, b.model.QuoteWord(column), in)
}

// WhereOrNotNull 构建 `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` 的 `OR` 条件语句。 md5:e122f662846a4ba4
func (b *WhereBuilder) WhereOrNotNull(columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.WhereOrf(`%s IS NOT NULL`, b.model.QuoteWord(column))
	}
	return builder
}
