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
		placeHolderCount = gstr.X统计次数(format, "?")
		conditionStr     = fmt.Sprintf(format, args[:len(args)-placeHolderCount]...)
	)
	return b.doWhereType(t, conditionStr, args[len(args)-placeHolderCount:]...)
}

// X条件 方法为构建器设置条件语句。参数 `where` 可以是多种类型，包括
// 字符串、映射、gmap（可能是自定义的映射类型）、切片、结构体、指针到结构体等。需要注意的是，
// 如果该方法被调用了多次，多个条件将会使用 "AND" 连接起来组成 WHERE 子句。
// 例如：
// X条件("uid=10000")
// X条件("uid", 10000)
// X条件("money>? AND name like ?", 99999, "vip_%")
// X条件("uid", 1).X条件("name", "john")
// X条件("status IN (?)", g.Slice{1,2,3})
// X条件("age IN(?,?)", 18, 50)
// X条件(User{ Id : 1, UserName : "john"}) // 使用结构体作为查询条件
// md5:38a2f7ff889346c5
func (b *WhereBuilder) X条件(条件 interface{}, 参数 ...interface{}) *WhereBuilder {
	return b.doWhereType(``, 条件, 参数...)
}

// X条件格式化 使用 fmt.Sprintf 和参数构建条件字符串。
// 注意，如果`args`的数目超过`format`中的占位符，
// 多余的`args`将作为Model的WHERE条件参数。
// 例如：
// X条件格式化(`amount<? and status=%s`, "paid", 100)  => WHERE `amount`<100 and status='paid'
// X条件格式化(`amount<%d and status=%s`, 100, "paid") => WHERE `amount`<100 and status='paid'
// md5:e4748efd7332202a
func (b *WhereBuilder) X条件格式化(格式 string, 参数 ...interface{}) *WhereBuilder {
	return b.doWherefType(``, 格式, 参数...)
}

// X条件并识别主键 的逻辑与 Model.Where 相同，但当参数 `where` 是单个条件，如 int、string、float 或 slice 时，它将该条件视为主键值。也就是说，如果主键是 "id" 并且给定的 `where` 参数为 "123"，X条件并识别主键 函数会将条件解析为 "id=123"，而 Model.Where 则会将条件视为字符串 "123"。
// md5:2545fa57bcbd235c
func (b *WhereBuilder) X条件并识别主键(条件 interface{}, 参数 ...interface{}) *WhereBuilder {
	if len(参数) > 0 {
		return b.X条件(条件, 参数...)
	}
	newWhere := X底层GetPrimaryKeyCondition(b.model.getPrimaryKey(), 条件)
	return b.X条件(newWhere[0], newWhere[1:]...)
}

// X条件小于构建`column < value`语句。 md5:438e43e951037408
func (b *WhereBuilder) X条件小于(字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s < ?`, b.model.X底层QuoteWord(字段名), 比较值)
}

// X条件小于等于 构建 `column <= value` 的语句。 md5:149d7bc478d211fd
func (b *WhereBuilder) X条件小于等于(字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s <= ?`, b.model.X底层QuoteWord(字段名), 比较值)
}

// X条件大于 构建 `column > value` 语句。 md5:41527fa039c8a299
func (b *WhereBuilder) X条件大于(字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s > ?`, b.model.X底层QuoteWord(字段名), 比较值)
}

// X条件大于等于构建`column >= value`语句。 md5:fff159ae64237621
func (b *WhereBuilder) X条件大于等于(字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s >= ?`, b.model.X底层QuoteWord(字段名), 比较值)
}

// X条件取范围构建`column BETWEEN min AND max`语句。 md5:cdb9b4a2942f3b60
func (b *WhereBuilder) X条件取范围(字段名 string, 最小值, 最大值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s BETWEEN ? AND ?`, b.model.X底层QuoteWord(字段名), 最小值, 最大值)
}

// X条件模糊匹配 构建 `column LIKE like` 语句。 md5:5cf0790f9754307f
func (b *WhereBuilder) X条件模糊匹配(字段名 string, 通配符条件值 string) *WhereBuilder {
	return b.X条件格式化(`%s LIKE ?`, b.model.X底层QuoteWord(字段名), 通配符条件值)
}

// X条件包含 构建 `column IN (in)` 语句。 md5:08648a50bb84e2ee
func (b *WhereBuilder) X条件包含(字段名 string, 包含值 interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s IN (?)`, b.model.X底层QuoteWord(字段名), 包含值)
}

// X条件NULL值 构建 `columns[0] IS NULL AND columns[1] IS NULL ...` 语句。 md5:9341218ae0c32357
func (b *WhereBuilder) X条件NULL值(字段名 ...string) *WhereBuilder {
	builder := b
	for _, column := range 字段名 {
		builder = builder.X条件格式化(`%s IS NULL`, b.model.X底层QuoteWord(column))
	}
	return builder
}

// X条件取范围以外构建`column NOT BETWEEN min AND max`语句。 md5:ac5d20d314a9fa0c
func (b *WhereBuilder) X条件取范围以外(字段名 string, 最小值, 最大值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s NOT BETWEEN ? AND ?`, b.model.X底层QuoteWord(字段名), 最小值, 最大值)
}

// X条件模糊匹配以外 构建 `column NOT LIKE like` 的 SQL 语句。 md5:683105cb42e27e3b
func (b *WhereBuilder) X条件模糊匹配以外(字段名 string, 通配符条件值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s NOT LIKE ?`, b.model.X底层QuoteWord(字段名), 通配符条件值)
}

// X条件不等于 构建 `column != value` 语句。 md5:d409867c3e8a9641
func (b *WhereBuilder) X条件不等于(字段名 string, 值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s != ?`, b.model.X底层QuoteWord(字段名), 值)
}

// X条件不包含构建`column NOT IN (in)`语句。 md5:658ffbae4d294fa4
func (b *WhereBuilder) X条件不包含(字段名 string, 不包含值 interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s NOT IN (?)`, b.model.X底层QuoteWord(字段名), 不包含值)
}

// X条件非Null 构建 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。 md5:2444d6e2f6bcbf2d
func (b *WhereBuilder) X条件非Null(字段名 ...string) *WhereBuilder {
	builder := b
	for _, column := range 字段名 {
		builder = builder.X条件格式化(`%s IS NOT NULL`, b.model.X底层QuoteWord(column))
	}
	return builder
}
