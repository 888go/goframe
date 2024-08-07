// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

// X条件带前缀 的功能类似于 Where，但它会在 where 语句中的每个字段前添加一个前缀。
// 例如：
// X条件带前缀("order", "status", "paid")                        => WHERE `order`.`status`='paid'
// X条件带前缀("order", struct{Status:"paid", "channel":"bank"}) => WHERE `order`.`status`='paid' AND `order`.`channel`='bank'
// 
// 这段注释的中文翻译为：
// 
// X条件带前缀 的行为与 Where 相似，但它会在 where 子句里的每个字段前加上一个前缀。
// 例如：
// X条件带前缀("order", "status", "paid")                        => 生成 WHERE `order`.`status`='paid'
// X条件带前缀("order", struct{Status:"paid", "channel":"bank"}) => 生成 WHERE `order`.`status`='paid' AND `order`.`channel`='bank'
// md5:062302edb484784b
func (b *WhereBuilder) X条件带前缀(字段前缀 string, 条件 interface{}, 参数 ...interface{}) *WhereBuilder {
	条件, 参数 = b.convertWhereBuilder(条件, 参数)

	builder := b.getBuilder()
	if builder.whereHolder == nil {
		builder.whereHolder = make([]WhereHolder, 0)
	}
	builder.whereHolder = append(builder.whereHolder, WhereHolder{
		Type:     whereHolderTypeDefault,
		Operator: whereHolderOperatorWhere,
		Where:    条件,
		Args:     参数,
		Prefix:   字段前缀,
	})
	return builder
}

// X条件小于并带前缀 构建 `prefix.column < value` 语句。 md5:de5cb5259ef84499
func (b *WhereBuilder) X条件小于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s.%s < ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 比较值)
}

// X条件小于等于并带前缀构建`prefix.column <= value`语句。 md5:1c5d93f173a39b03
func (b *WhereBuilder) X条件小于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s.%s <= ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 比较值)
}

// X条件大于并带前缀构建`prefix.column > value`语句。 md5:61d5cbbb9f5422fd
func (b *WhereBuilder) X条件大于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s.%s > ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 比较值)
}

// X条件大于等于并带前缀 生成 `prefix.column >= value` 的语句。 md5:1b581ea600e215e7
func (b *WhereBuilder) X条件大于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s.%s >= ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 比较值)
}

// X条件取范围并带前缀 构建 `prefix.column BETWEEN min AND max` 语句。 md5:e6176c55b8a31575
func (b *WhereBuilder) X条件取范围并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s.%s BETWEEN ? AND ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 最小值, 最大值)
}

// X条件模糊匹配并带前缀构建`prefix.column LIKE like`语句。 md5:baf08eac5c7dc2aa
func (b *WhereBuilder) X条件模糊匹配并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s.%s LIKE ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 通配符条件值)
}

// X条件包含并带前缀 构建 `prefix.column IN (in)` 语句。 md5:fd691f634711ba7f
func (b *WhereBuilder) X条件包含并带前缀(字段前缀 string, 字段名 string, 包含值 interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s.%s IN (?)`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 包含值)
}

// X条件NULL值并带前缀 构建 `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` 语句。 md5:ac08bde96048fdce
func (b *WhereBuilder) X条件NULL值并带前缀(字段前缀 string, 字段名 ...string) *WhereBuilder {
	builder := b
	for _, column := range 字段名 {
		builder = builder.X条件格式化(`%s.%s IS NULL`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(column))
	}
	return builder
}

// X条件取范围以外并带前缀 构建 `prefix.column NOT BETWEEN min AND max` 语句。 md5:a16703b511af05c3
func (b *WhereBuilder) X条件取范围以外并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s.%s NOT BETWEEN ? AND ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 最小值, 最大值)
}

// X条件模糊匹配以外并带前缀构建`prefix.column NOT LIKE like`语句。 md5:083bd1d45c368a83
func (b *WhereBuilder) X条件模糊匹配以外并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s.%s NOT LIKE ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 通配符条件值)
}

// X条件不等于并带前缀构建`prefix.column != value`语句。 md5:c1366e00cd0da49e
func (b *WhereBuilder) X条件不等于并带前缀(字段前缀 string, 字段名 string, 值 interface{}) *WhereBuilder {
	return b.X条件格式化(`%s.%s != ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 值)
}

// X条件不包含并带前缀 用于构建 `prefix.column NOT IN (in)` 的SQL语句。 md5:3b790678c07a51fd
func (b *WhereBuilder) X条件不包含并带前缀(字段前缀 string, 字段名 string, 不包含值 interface{}) *WhereBuilder {
	return b.doWherefType(whereHolderTypeIn, `%s.%s NOT IN (?)`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 不包含值)
}

// X条件非Null并带前缀 构建 `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` 语句。 md5:d5a307a7c3004dda
func (b *WhereBuilder) X条件非Null并带前缀(字段前缀 string, 字段名 ...string) *WhereBuilder {
	builder := b
	for _, column := range 字段名 {
		builder = builder.X条件格式化(`%s.%s IS NOT NULL`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(column))
	}
	return builder
}
