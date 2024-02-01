// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb
import (
	"fmt"
	
	"github.com/888go/goframe/text/gstr"
	)
// doWhereType 为模型设置条件语句。参数`where`可以是以下类型：
// string（字符串）、map（映射表）、gmap（Golang自定义的映射类型）、slice（切片）、struct（结构体）、*struct（指向结构体的指针）等。
// 注意，如果该方法被调用超过一次，则多次调用时传入的多个条件会通过"AND"运算符连接到一起形成最终的where条件语句。
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

// doWherefType 使用fmt.Sprintf和参数构建条件字符串。
// 注意，如果`args`的数量大于`format`中的占位符，
// 额外的`args`将被用作Model的where条件参数。
func (b *WhereBuilder) doWherefType(t string, format string, args ...interface{}) *WhereBuilder {
	var (
		placeHolderCount = gstr.Count(format, "?")
		conditionStr     = fmt.Sprintf(format, args[:len(args)-placeHolderCount]...)
	)
	return b.doWhereType(t, conditionStr, args[len(args)-placeHolderCount:]...)
}

// Where 设置构建器的条件语句。参数`where`可以是以下类型：
// string/map/gmap/slice/struct/*struct 等等。请注意，如果调用该方法超过一次，
// 多个条件将会通过使用 "AND" 连接符合并到 where 语句中。
// 示例：
// Where("uid=10000") // 设置条件：uid为10000
// Where("uid", 10000) // 设置条件：uid为10000
// Where("money>? AND name like ?", 99999, "vip_%"）// 设置条件：money大于99999且name字段以"vip_"开头
// Where("uid", 1).Where("name", "john") // 设置条件：uid为1且name为"john"
// Where("status IN (?)", g.Slice{1,2,3}) // 设置条件：status在[1,2,3]范围内
// Where("age IN(?,?)", 18, 50) // 设置条件：age在[18,50]范围内
// Where(User{ Id : 1, UserName : "john"}) // 设置条件：根据User结构体内容（Id为1，UserName为"john"）
func (b *WhereBuilder) Where(where interface{}, args ...interface{}) *WhereBuilder {
	return b.doWhereType(``, where, args...)
}

// Wheref 通过 fmt.Sprintf 和参数构建条件字符串。
// 注意，如果 `args` 的数量大于 `format` 中的占位符数量，
// 额外的 `args` 将作为 Model 的 where 条件参数使用。
// 示例：
// Wheref(`amount<? and status=%s`, "paid", 100)  => WHERE `amount`<100 and status='paid'
// Wheref(`amount<%d and status=%s`, 100, "paid") => WHERE `amount`<100 and status='paid'
func (b *WhereBuilder) Wheref(format string, args ...interface{}) *WhereBuilder {
	return b.doWherefType(``, format, args...)
}

// WherePri 执行的逻辑与 Model.Where 相同，但有个特殊情况：
// 如果参数 `where` 是单个条件，如 int、string、float 或 slice 类型，
// 那么它会将这个条件视为主键值。也就是说，如果主键是 "id"，给定的 `where` 参数为 "123"，
// 那么 WherePri 函数会将条件处理为 "id=123"；
// 但是 Model.Where 函数则会将条件处理为字符串 "123"。
func (b *WhereBuilder) WherePri(where interface{}, args ...interface{}) *WhereBuilder {
	if len(args) > 0 {
		return b.Where(where, args...)
	}
	newWhere := GetPrimaryKeyCondition(b.model.getPrimaryKey(), where)
	return b.Where(newWhere[0], newWhere[1:]...)
}

// WhereLT 用于构建 `column < value` 语句。
func (b *WhereBuilder) WhereLT(column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s < ?`, b.model.QuoteWord(column), value)
}

// WhereLTE 用于构建 `column <= value` 的表达式语句。
func (b *WhereBuilder) WhereLTE(column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s <= ?`, b.model.QuoteWord(column), value)
}

// WhereGT 构建 `column > value` 语句。
func (b *WhereBuilder) WhereGT(column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s > ?`, b.model.QuoteWord(column), value)
}

// WhereGTE 用于构建 `column >= value` 的表达式语句。
func (b *WhereBuilder) WhereGTE(column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s >= ?`, b.model.QuoteWord(column), value)
}

// WhereBetween 用于构建 `column BETWEEN min AND max` 语句。
func (b *WhereBuilder) WhereBetween(column string, min, max interface{}) *WhereBuilder {
	return b.Wheref(`%s BETWEEN ? AND ?`, b.model.QuoteWord(column), min, max)
}

// WhereLike 用于构建 `column LIKE like` 语句。
func (b *WhereBuilder) WhereLike(column string, like string) *WhereBuilder {
	return b.Wheref(`%s LIKE ?`, b.model.QuoteWord(column), like)
}

// WhereIn 构建 `column IN (in)` 语句。
func (b *WhereBuilder) WhereIn(column string, in interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s IN (?)`, b.model.QuoteWord(column), in)
}

// WhereNull 用于构建 `columns[0] IS NULL AND columns[1] IS NULL ...` 语句。
func (b *WhereBuilder) WhereNull(columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.Wheref(`%s IS NULL`, b.model.QuoteWord(column))
	}
	return builder
}

// WhereNotBetween 用于构建 `column NOT BETWEEN min AND max` 语句。
func (b *WhereBuilder) WhereNotBetween(column string, min, max interface{}) *WhereBuilder {
	return b.Wheref(`%s NOT BETWEEN ? AND ?`, b.model.QuoteWord(column), min, max)
}

// WhereNotLike 构建 `column NOT LIKE like` 语句。
func (b *WhereBuilder) WhereNotLike(column string, like interface{}) *WhereBuilder {
	return b.Wheref(`%s NOT LIKE ?`, b.model.QuoteWord(column), like)
}

// WhereNot 用于构建 `column != value` 的表达式语句。
func (b *WhereBuilder) WhereNot(column string, value interface{}) *WhereBuilder {
	return b.Wheref(`%s != ?`, b.model.QuoteWord(column), value)
}

// WhereNotIn 构造 `column NOT IN (in)` 语句。
func (b *WhereBuilder) WhereNotIn(column string, in interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s NOT IN (?)`, b.model.QuoteWord(column), in)
}

// WhereNotNull 用于构建 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。
func (b *WhereBuilder) WhereNotNull(columns ...string) *WhereBuilder {
	builder := b
	for _, column := range columns {
		builder = builder.Wheref(`%s IS NOT NULL`, b.model.QuoteWord(column))
	}
	return builder
}
