// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"fmt"

	gstr "github.com/888go/goframe/text/gstr"
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
		placeHolderCount = gstr.X统计次数(format, "?")
		conditionStr     = fmt.Sprintf(format, args[:len(args)-placeHolderCount]...)
	)
	return b.doWhereOrType(t, conditionStr, args[len(args)-placeHolderCount:]...)
}

// X条件或 向 WHERE 语句中添加“OR”条件。 md5:753c32f428b02541
func (b *WhereBuilder) X条件或(条件 interface{}, 参数 ...interface{}) *WhereBuilder {
	return b.doWhereOrType(``, 条件, 参数...)
}

// X条件或格式化 使用fmt.Sprintf和参数构建`OR`条件字符串。 md5:aa04236f081a2885
// Eg:
// X条件或格式化(`amount<? and status=%s`, "paid", 100)  => WHERE xxx OR `amount`<100 and status='paid'
// X条件或格式化(`amount<%d and status=%s`, 100, "paid") => WHERE xxx OR `amount`<100 and status='paid'
func (b *WhereBuilder) X条件或格式化(格式 string, 参数 ...interface{}) *WhereBuilder {
	return b.doWhereOrfType(``, 格式, 参数...)
}

// X条件或不等于在`OR`条件下构建`column != value`语句。 md5:adc6d63e61bf279f
func (b *WhereBuilder) X条件或不等于(字段名 string, 值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s != ?`, 字段名, 值)
}

// X条件或小于 在 `OR` 条件下构建 `column < value` 的语句。 md5:5517b3812e2c8e8b
func (b *WhereBuilder) X条件或小于(字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s < ?`, 字段名, 比较值)
}

// X条件或小于等于 在 OR 条件中构建 `column <= value` 语句。 md5:3b0287bd1f8030ce
func (b *WhereBuilder) X条件或小于等于(字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s <= ?`, 字段名, 比较值)
}

// X条件或大于在`OR`条件下构建`column > value`语句。 md5:2289d39bb82e521f
func (b *WhereBuilder) X条件或大于(字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s > ?`, 字段名, 比较值)
}

// X条件或大于等于在`OR`条件下构建`column >= value`语句。 md5:e178dd8cfc5661e5
func (b *WhereBuilder) X条件或大于等于(字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s >= ?`, 字段名, 比较值)
}

// X条件或取范围 用于构建 `column BETWEEN min AND max` 语句，并在 `OR` 条件下使用。 md5:90f98622a1fd5981
func (b *WhereBuilder) X条件或取范围(字段名 string, 最小值, 最大值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s BETWEEN ? AND ?`, b.model.X底层QuoteWord(字段名), 最小值, 最大值)
}

// X条件或模糊匹配 在 `OR` 条件中构建 `column LIKE 'like'` 语句。 md5:7a2d37411752fb51
func (b *WhereBuilder) X条件或模糊匹配(字段名 string, 通配符条件值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s LIKE ?`, b.model.X底层QuoteWord(字段名), 通配符条件值)
}

// X条件或包含 在`OR`条件下构建`column IN (in)`语句。 md5:4bb93b5ae9a5e887
func (b *WhereBuilder) X条件或包含(字段名 string, 包含值 interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s IN (?)`, b.model.X底层QuoteWord(字段名), 包含值)
}

// X条件或NULL值 在 `OR` 条件下构建 `columns[0] IS NULL OR columns[1] IS NULL ...` 语句。 md5:08d38a60dc594441
func (b *WhereBuilder) X条件或NULL值(字段名 ...string) *WhereBuilder {
	var builder *WhereBuilder
	for _, column := range 字段名 {
		builder = b.X条件或格式化(`%s IS NULL`, b.model.X底层QuoteWord(column))
	}
	return builder
}

// X条件或取范围以外 用于构建在 `OR` 条件下的 `column NOT BETWEEN min AND max` 语句。 md5:f20408e0126bbbab
func (b *WhereBuilder) X条件或取范围以外(字段名 string, 最小值, 最大值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s NOT BETWEEN ? AND ?`, b.model.X底层QuoteWord(字段名), 最小值, 最大值)
}

// X条件或模糊匹配以外 在 OR 条件中构建 `column NOT LIKE like` 语句。 md5:751e840816119632
func (b *WhereBuilder) X条件或模糊匹配以外(字段名 string, 通配符条件值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s NOT LIKE ?`, b.model.X底层QuoteWord(字段名), 通配符条件值)
}

// X条件或不包含构建`column NOT IN (in)`语句。 md5:433fd8a0f224fc24
func (b *WhereBuilder) X条件或不包含(字段名 string, 不包含值 interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s NOT IN (?)`, b.model.X底层QuoteWord(字段名), 不包含值)
}

// X条件或非Null 构建 `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` 的 `OR` 条件语句。 md5:e122f662846a4ba4
func (b *WhereBuilder) X条件或非Null(字段名 ...string) *WhereBuilder {
	builder := b
	for _, column := range 字段名 {
		builder = builder.X条件或格式化(`%s IS NOT NULL`, b.model.X底层QuoteWord(column))
	}
	return builder
}
