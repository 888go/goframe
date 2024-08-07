// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

// X条件或并带前缀 的功能类似于 WhereOr，但它会在 WHERE 子句中的每个字段前添加一个前缀。
// 例如：
// X条件或并带前缀("order", "status", "paid")                        => WHERE xxx OR (`order`.`status`='paid')
// X条件或并带前缀("order", struct{Status:"paid", "channel":"bank"}) => WHERE xxx OR (`order`.`status`='paid' AND `order`.`channel`='bank')
// 
// 这意味着 X条件或并带前缀 函数允许你在一个 WHERE 子句中指定多个条件，并且自动为这些条件的字段名加上一个指定的前缀，以便清晰地指向某个表或结构。它可以处理单个字段值的情况，也可以处理包含多个键值对的结构体，以生成更复杂的逻辑组合。
// md5:2358d43541f472f5
func (b *WhereBuilder) X条件或并带前缀(字段前缀 string, 条件 interface{}, 参数 ...interface{}) *WhereBuilder {
	条件, 参数 = b.convertWhereBuilder(条件, 参数)

	builder := b.getBuilder()
	builder.whereHolder = append(builder.whereHolder, WhereHolder{
		Type:     whereHolderTypeDefault,
		Operator: whereHolderOperatorOr,
		Where:    条件,
		Args:     参数,
		Prefix:   字段前缀,
	})
	return builder
}

// X条件或不等于并带前缀 在 OR 条件中构建 `prefix.column != value` 语句。 md5:385a9f9fb58b8fc3
func (b *WhereBuilder) X条件或不等于并带前缀(字段前缀 string, 字段名 string, 值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s.%s != ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 值)
}

// X条件或小于并带前缀在"OR"条件下构建`prefix.column < value`语句。 md5:c1a6baf94f553043
func (b *WhereBuilder) X条件或小于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s.%s < ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 比较值)
}

// X条件或小于等于并带前缀 在“OR”条件下构建 `prefix.column <= value` 语句。 md5:77072877b38f04a8
func (b *WhereBuilder) X条件或小于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s.%s <= ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 比较值)
}

// X条件或大于并带前缀 在 `OR` 条件下构建 `prefix.column > value` 的语句。 md5:d34b5bdc0e6b2fa8
func (b *WhereBuilder) X条件或大于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s.%s > ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 比较值)
}

// X条件或大于等于并带前缀 在 OR 条件中构建 `prefix.column >= value` 语句。 md5:d652ca0304ac835e
func (b *WhereBuilder) X条件或大于等于并带前缀(字段前缀 string, 字段名 string, 比较值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s.%s >= ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 比较值)
}

// X条件或取范围并带前缀在`OR`条件下构建`prefix.column BETWEEN min AND max`语句。 md5:d7adaf273fa5681b
func (b *WhereBuilder) X条件或取范围并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s.%s BETWEEN ? AND ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 最小值, 最大值)
}

// X条件或模糊匹配并带前缀在`OR`条件下构建`prefix.column LIKE 'like'`语句。 md5:c975b47e3a5cc2c1
func (b *WhereBuilder) X条件或模糊匹配并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s.%s LIKE ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 通配符条件值)
}

// X条件或包含并带前缀 用于构建 `prefix.column IN (in)` 形式的 SQL 语句，在 `OR` 条件下。 md5:18e0cf5cc971267d
func (b *WhereBuilder) X条件或包含并带前缀(字段前缀 string, 字段名 string, 包含值 interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s.%s IN (?)`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 包含值)
}

// X条件或NULL值并带前缀 在"OR"条件中构建`prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...`语句。 md5:facf88eb72f3d299
func (b *WhereBuilder) X条件或NULL值并带前缀(字段前缀 string, 字段名 ...string) *WhereBuilder {
	builder := b
	for _, column := range 字段名 {
		builder = builder.X条件或格式化(`%s.%s IS NULL`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(column))
	}
	return builder
}

// X条件或取范围以外并带前缀在`OR`条件下构建`prefix.column NOT BETWEEN min AND max`语句。 md5:15259f135308893b
func (b *WhereBuilder) X条件或取范围以外并带前缀(字段前缀 string, 字段名 string, 最小值, 最大值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s.%s NOT BETWEEN ? AND ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 最小值, 最大值)
}

// X条件或模糊匹配以外并带前缀 在 `OR` 条件下构建 `prefix.column NOT LIKE 'like'` 语句。 md5:2785cbc79e811104
func (b *WhereBuilder) X条件或模糊匹配以外并带前缀(字段前缀 string, 字段名 string, 通配符条件值 interface{}) *WhereBuilder {
	return b.X条件或格式化(`%s.%s NOT LIKE ?`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 通配符条件值)
}

// X条件或不包含并带前缀 用于构建 `prefix.column NOT IN (in)` 的SQL语句。 md5:bd296110bb5635a1
func (b *WhereBuilder) X条件或不包含并带前缀(字段前缀 string, 字段名 string, 不包含值 interface{}) *WhereBuilder {
	return b.doWhereOrfType(whereHolderTypeIn, `%s.%s NOT IN (?)`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(字段名), 不包含值)
}

// X条件或非Null并带前缀 在`OR`条件中构建`prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...`语句。 md5:9ecd3bffabf47cb7
func (b *WhereBuilder) X条件或非Null并带前缀(字段前缀 string, 字段名 ...string) *WhereBuilder {
	builder := b
	for _, column := range 字段名 {
		builder = builder.X条件或格式化(`%s.%s IS NOT NULL`, b.model.X底层QuoteWord(字段前缀), b.model.X底层QuoteWord(column))
	}
	return builder
}
