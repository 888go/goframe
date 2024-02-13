// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"fmt"
	
	"github.com/888go/goframe/text/gstr"
)

// WhereOr 向 where 语句添加“OR”条件。
func (b *X组合条件) doWhereOrType(t string, where interface{}, args ...interface{}) *X组合条件 {
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
func (b *X组合条件) doWhereOrfType(t string, format string, args ...interface{}) *X组合条件 {
	var (
		placeHolderCount = 文本类.X统计次数(format, "?")
		conditionStr     = fmt.Sprintf(format, args[:len(args)-placeHolderCount]...)
	)
	return b.doWhereOrType(t, conditionStr, args[len(args)-placeHolderCount:]...)
}

// WhereOr 向 where 语句添加“OR”条件。
func (b *X组合条件) X条件或(条件 interface{}, 参数 ...interface{}) *X组合条件 {
	return b.doWhereOrType(``, 条件, 参数...)
}

// WhereOrf 使用 fmt.Sprintf 和参数构建 `OR` 条件字符串。
// Eg:
// WhereOrf(`amount<? and status=%s`, "paid", 100)  => WHERE xxx OR `amount`<100 and status='paid'
// WhereOrf(`amount<%d and status=%s`, 100, "paid") => WHERE xxx OR `amount`<100 and status='paid'
func (b *X组合条件) X条件或格式化(格式 string, 参数 ...interface{}) *X组合条件 {
	return b.doWhereOrfType(``, 格式, 参数...)
}

// WhereOrNot 用于构建在“OR”条件中的“column != value”语句。
func (b *X组合条件) X条件或不等于(字段名 string, 值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s != ?`, 字段名, 值)
}

// WhereOrLT 在“OR”条件中构建“column < value”语句。
func (b *X组合条件) X条件或小于(字段名 string, 比较值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s < ?`, 字段名, 比较值)
}

// WhereOrLTE 用于构建在“OR”条件中的`column <= value`语句。
func (b *X组合条件) X条件或小于等于(字段名 string, 比较值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s <= ?`, 字段名, 比较值)
}

// WhereOrGT 在“OR”条件下构建“column > value”语句。
func (b *X组合条件) X条件或大于(字段名 string, 比较值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s > ?`, 字段名, 比较值)
}

// WhereOrGTE 在“OR”条件下构建“column >= value”语句。
func (b *X组合条件) X条件或大于等于(字段名 string, 比较值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s >= ?`, 字段名, 比较值)
}

// WhereOrBetween 在`OR`条件下构建 `column BETWEEN min AND max` 语句。
func (b *X组合条件) X条件或取范围(字段名 string, 最小值, 最大值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s BETWEEN ? AND ?`, b.model.底层QuoteWord(字段名), 最小值, 最大值)
}

// WhereOrLike 在`OR`条件下构建`column LIKE 'like'`语句。
func (b *X组合条件) X条件或模糊匹配(字段名 string, 通配符条件值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s LIKE ?`, b.model.底层QuoteWord(字段名), 通配符条件值)
}

// WhereOrIn 在“OR”条件中构建“column IN (in)”语句。
func (b *X组合条件) X条件或包含(字段名 string, 包含值 interface{}) *X组合条件 {
	return b.doWhereOrfType(whereHolderTypeIn, `%s IN (?)`, b.model.底层QuoteWord(字段名), 包含值)
}

// WhereOrNull 根据“或”条件构建 `columns[0] IS NULL OR columns[1] IS NULL ...` 语句。
func (b *X组合条件) X条件或NULL值(字段名 ...string) *X组合条件 {
	var builder *X组合条件
	for _, column := range 字段名 {
		builder = b.X条件或格式化(`%s IS NULL`, b.model.底层QuoteWord(column))
	}
	return builder
}

// WhereOrNotBetween 在“OR”条件中构建`column NOT BETWEEN min AND max`语句。
func (b *X组合条件) X条件或取范围以外(字段名 string, 最小值, 最大值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s NOT BETWEEN ? AND ?`, b.model.底层QuoteWord(字段名), 最小值, 最大值)
}

// WhereOrNotLike 在“OR”条件下构建`column NOT LIKE like`语句。
func (b *X组合条件) X条件或模糊匹配以外(字段名 string, 通配符条件值 interface{}) *X组合条件 {
	return b.X条件或格式化(`%s NOT LIKE ?`, b.model.底层QuoteWord(字段名), 通配符条件值)
}

// WhereOrNotIn 用于构建 `column NOT IN (in)` 语句。
func (b *X组合条件) X条件或不包含(字段名 string, 不包含值 interface{}) *X组合条件 {
	return b.doWhereOrfType(whereHolderTypeIn, `%s NOT IN (?)`, b.model.底层QuoteWord(字段名), 不包含值)
}

// WhereOrNotNull 在`OR`条件下构建 `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` 语句。
func (b *X组合条件) X条件或非Null(字段名 ...string) *X组合条件 {
	builder := b
	for _, column := range 字段名 {
		builder = builder.X条件或格式化(`%s IS NOT NULL`, b.model.底层QuoteWord(column))
	}
	return builder
}
