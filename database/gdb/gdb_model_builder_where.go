// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"fmt"

	"github.com/gogf/gf/v2/text/gstr"
)

// doWhereType 设置模型的条件语句。参数 `where` 可以是类型为 string、map、gmap、slice、struct 或其派生结构等。需要注意的是，如果多次调用此函数，多个条件将使用 "AND" 连接到 where 语句中。
// md5:92ee322b44569cba
func (b *WhereBuilder) doWhereType(whereType string, where interface{}, args ...interface{}) *WhereBuilder {
	where, args = b.convertWhereBuilder(where, args)

	builder := b.getBuilder()
	if builder.whereHolder == nil {
		builder.whereHolder = make([]WhereHolder, 0)
	}
	if whereType == "" {
		if len(args) == 0 {
			whereType = whereHolderTypeNoArgs
		} else {
			whereType = whereHolderTypeDefault
		}
	}
	builder.whereHolder = append(builder.whereHolder, WhereHolder{
		Type:     whereType,
		Operator: whereHolderOperatorWhere,
		Where:    where,
		Args:     args,
	})
	return builder
}

// doWherefType 使用 fmt.Sprintf 和参数构建条件字符串。
// 注意，如果 `args` 的数量多于 `format` 中的占位符，多余的 `args` 将作为 Model 的 where 条件参数使用。
// md5:67cfb01201c57037
func (b *WhereBuilder) doWherefType(t string, format string, args ...interface{}) *WhereBuilder {
	var (
		placeHolderCount = gstr.Count(format, "?")
		conditionStr     = fmt.Sprintf(format, args[:len(args)-placeHolderCount]...)
	)
	return b.doWhereType(t, conditionStr, args[len(args)-placeHolderCount:]...)
}

// Where sets the condition statement for the builder. The parameter `where` can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"}).
// ff:条件
// b:
// where:条件
// args:参数
func (b *WhereBuilder) Where(where interface{}, args ...interface{}) *WhereBuilder {
	return b.doWhereType(``, where, args...)
}

// Wheref builds condition string using fmt.Sprintf and arguments.
// Note that if the number of `args` is more than the placeholder in `format`,
// the extra `args` will be used as the where condition arguments of the Model.
// Wheref(`amount<? and status=%s`, "paid", 100)  => WHERE `amount`<100 and status='paid'
// Wheref(`amount<%d and status=%s`, 100, "paid") => WHERE `amount`<100 and status='paid'
// ff:条件格式化
// b:
// format:格式
// args:参数
func (b *WhereBuilder) Wheref(format string, args ...interface{}) *WhereBuilder {
	return b.doWherefType(``, format, args...)
}

// WherePri 的逻辑与 Model.Where 相同，但当参数 `where` 是单个条件，如 int、string、float 或 slice 时，它将该条件视为主键值。也就是说，如果主键是 "id" 并且给定的 `where` 参数为 "123"，WherePri 函数会将条件解析为 "id=123"，而 Model.Where 则会将条件视为字符串 "123"。
// md5:2545fa57bcbd235c
// ff:条件并识别主键
// b:
// where:条件
// args:参数
func (b *WhereBuilder) WherePri(where interface{}, args ...interface{}) *WhereBuilder {
	if len(args) > 0 {
		return b.Where(where, args...)
	}
	newWhere := GetPrimaryKeyCondition(b.model.getPrimaryKey(), where)
	return b.Where(newWhere[0], newWhere[1:]...)
}

// WhereLT构建`column < value`语句。 md5:438e43e951037408
// ff:条件小于
// b:
// column:字段名
// value:比较值
func (b *WhereBuilder) WhereLT(column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s < ?`, b.model.QuoteWord(column), value)
}

// WhereLTE 构建 `column <= value` 的语句。 md5:149d7bc478d211fd
// ff:条件小于等于
// b:
// column:字段名
// value:比较值
func (b *WhereBuilder) WhereLTE(column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s <= ?`, b.model.QuoteWord(column), value)
}

// WhereGT 构建 `column > value` 语句。 md5:41527fa039c8a299
// ff:条件大于
// b:
// column:字段名
// value:比较值
func (b *WhereBuilder) WhereGT(column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s > ?`, b.model.QuoteWord(column), value)
}

// WhereGTE构建`column >= value`语句。 md5:fff159ae64237621
// ff:条件大于等于
// b:
// column:字段名
// value:比较值
func (b *WhereBuilder) WhereGTE(column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s >= ?`, b.model.QuoteWord(column), value)
}

// WhereBetween构建`column BETWEEN min AND max`语句。 md5:cdb9b4a2942f3b60
// ff:条件取范围
// b:
// column:字段名
// min:最小值
// max:最大值
func (b *WhereBuilder) WhereBetween(column string, min, max interface{}) *WhereBuilder {
	return b.Wheref(`%s BETWEEN ? AND ?`, b.model.QuoteWord(column), min, max)
}

// WhereLike 构建 `column LIKE like` 语句。 md5:5cf0790f9754307f
// ff:条件模糊匹配
// b:
// column:字段名
// like:通配符条件值
func (b *WhereBuilder) WhereLike(column string, like string) *WhereBuilder {
	return b.Wheref(`%s LIKE ?`, b.model.QuoteWord(column), like)
}

// WhereIn 构建 `column IN (in)` 语句。 md5:08648a50bb84e2ee
// ff:条件包含
// b:
// column:字段名
// in:包含值
func (b *WhereBuilder) WhereIn(column string, in interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s IN (?)`, b.model.QuoteWord(column), in)
}

// WhereNull 构建 `columns[0] IS NULL AND columns[1] IS NULL ...` 语句。 md5:9341218ae0c32357
// ff:条件NULL值
// b:
// columns:字段名
func (b *WhereBuilder) WhereNull(columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.Wheref(`%s IS NULL`, b.model.QuoteWord(column))
	}
	return builder
}

// WhereNotBetween构建`column NOT BETWEEN min AND max`语句。 md5:ac5d20d314a9fa0c
// ff:条件取范围以外
// b:
// column:字段名
// min:最小值
// max:最大值
func (b *WhereBuilder) WhereNotBetween(column string, min, max interface{}) *WhereBuilder {
	return b.Wheref(`%s NOT BETWEEN ? AND ?`, b.model.QuoteWord(column), min, max)
}

// WhereNotLike 构建 `column NOT LIKE like` 的 SQL 语句。 md5:683105cb42e27e3b
// ff:条件模糊匹配以外
// b:
// column:字段名
// like:通配符条件值
func (b *WhereBuilder) WhereNotLike(column string, like interface{}) *WhereBuilder {
	return b.Wheref(`%s NOT LIKE ?`, b.model.QuoteWord(column), like)
}

// WhereNot 构建 `column != value` 语句。 md5:d409867c3e8a9641
// ff:条件不等于
// b:
// column:字段名
// value:值
func (b *WhereBuilder) WhereNot(column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s != ?`, b.model.QuoteWord(column), value)
}

// WhereNotIn构建`column NOT IN (in)`语句。 md5:658ffbae4d294fa4
// ff:条件不包含
// b:
// column:字段名
// in:不包含值
func (b *WhereBuilder) WhereNotIn(column string, in interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s NOT IN (?)`, b.model.QuoteWord(column), in)
}

// WhereNotNull 构建 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。 md5:2444d6e2f6bcbf2d
// ff:条件非Null
// b:
// columns:字段名
func (b *WhereBuilder) WhereNotNull(columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.Wheref(`%s IS NOT NULL`, b.model.QuoteWord(column))
	}
	return builder
}
