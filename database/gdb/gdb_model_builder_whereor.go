// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

import (
	"fmt"
	
	"github.com/888go/goframe/text/gstr"
)

// WhereOr 向 where 语句添加“OR”条件。
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

// WhereOrf 使用 fmt.Sprintf 和参数构建 `OR` 条件字符串。
func (b *WhereBuilder) doWhereOrfType(t string, format string, args ...interface{}) *WhereBuilder {
	var (
		placeHolderCount = gstr.Count(format, "?")
		conditionStr     = fmt.Sprintf(format, args[:len(args)-placeHolderCount]...)
	)
	return b.doWhereOrType(t, conditionStr, args[len(args)-placeHolderCount:]...)
}

// WhereOr 向 where 语句添加“OR”条件。
func (b *WhereBuilder) WhereOr(where interface{}, args ...interface{}) *WhereBuilder {
	return b.doWhereOrType(``, where, args...)
}

// WhereOrf 使用 fmt.Sprintf 和参数构建 `OR` 条件字符串。
// Eg:
// WhereOrf(`amount<? and status=%s`, "paid", 100)  => WHERE xxx OR `amount`<100 and status='paid'
// WhereOrf(`amount<%d and status=%s`, 100, "paid") => WHERE xxx OR `amount`<100 and status='paid'
func (b *WhereBuilder) WhereOrf(format string, args ...interface{}) *WhereBuilder {
	return b.doWhereOrfType(``, format, args...)
}

// WhereOrNot 用于构建在“OR”条件中的“column != value”语句。
func (b *WhereBuilder) WhereOrNot(column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s != ?`, column, value)
}

// WhereOrLT 在“OR”条件中构建“column < value”语句。
func (b *WhereBuilder) WhereOrLT(column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s < ?`, column, value)
}

// WhereOrLTE 用于构建在“OR”条件中的`column <= value`语句。
func (b *WhereBuilder) WhereOrLTE(column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s <= ?`, column, value)
}

// WhereOrGT 在“OR”条件下构建“column > value”语句。
func (b *WhereBuilder) WhereOrGT(column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s > ?`, column, value)
}

// WhereOrGTE 在“OR”条件下构建“column >= value”语句。
func (b *WhereBuilder) WhereOrGTE(column string, value interface{}) *WhereBuilder {
	return b.WhereOrf(`%s >= ?`, column, value)
}

// WhereOrBetween 在`OR`条件下构建 `column BETWEEN min AND max` 语句。
func (b *WhereBuilder) WhereOrBetween(column string, min, max interface{}) *WhereBuilder {
	return b.WhereOrf(`%s BETWEEN ? AND ?`, b.model.QuoteWord(column), min, max)
}

// WhereOrLike 在`OR`条件下构建`column LIKE 'like'`语句。
func (b *WhereBuilder) WhereOrLike(column string, like interface{}) *WhereBuilder {
	return b.WhereOrf(`%s LIKE ?`, b.model.QuoteWord(column), like)
}

// WhereOrIn 在“OR”条件中构建“column IN (in)”语句。
func (b *WhereBuilder) WhereOrIn(column string, in interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s IN (?)`, b.model.QuoteWord(column), in)
}

// WhereOrNull 根据“或”条件构建 `columns[0] IS NULL OR columns[1] IS NULL ...` 语句。
func (b *WhereBuilder) WhereOrNull(columns ...string) *WhereBuilder {
	var builder *WhereBuilder
	for _, column := range columns {
		builder = b.WhereOrf(`%s IS NULL`, b.model.QuoteWord(column))
	}
	return builder
}

// WhereOrNotBetween 在“OR”条件中构建`column NOT BETWEEN min AND max`语句。
func (b *WhereBuilder) WhereOrNotBetween(column string, min, max interface{}) *WhereBuilder {
	return b.WhereOrf(`%s NOT BETWEEN ? AND ?`, b.model.QuoteWord(column), min, max)
}

// WhereOrNotLike 在“OR”条件下构建`column NOT LIKE like`语句。
func (b *WhereBuilder) WhereOrNotLike(column string, like interface{}) *WhereBuilder {
	return b.WhereOrf(`%s NOT LIKE ?`, b.model.QuoteWord(column), like)
}

// WhereOrNotIn 用于构建 `column NOT IN (in)` 语句。
func (b *WhereBuilder) WhereOrNotIn(column string, in interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s NOT IN (?)`, b.model.QuoteWord(column), in)
}

// WhereOrNotNull 在`OR`条件下构建 `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` 语句。
func (b *WhereBuilder) WhereOrNotNull(columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.WhereOrf(`%s IS NOT NULL`, b.model.QuoteWord(column))
	}
	return builder
}
