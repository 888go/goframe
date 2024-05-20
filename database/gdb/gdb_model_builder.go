// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"fmt"
)

// WhereBuilder 在一个分组中持有多个where条件。. md5:184a682305e36f1a
type WhereBuilder struct {
	model       *Model        // WhereBuilder 应该绑定到特定的模型上。. md5:b2939842d84a957f
	whereHolder []WhereHolder // 条件字符串，用于where操作。. md5:081bfa2bf97e01c3
}

// WhereHolder是用于准备where条件的持有者。. md5:be3c4766444c1f3c
type WhereHolder struct {
	Type     string        // Type of this holder.
	Operator int           // 该持有者的运算符。. md5:52eab23e6ea6fbdb
	Where    interface{}   // Where 参数，通常可以是字符串、映射或结构体类型。. md5:dc1254ec6b4f31fa
	Args     []interface{} // where参数的arguments。. md5:3ae3eaf7373aa4ff
	Prefix   string        // 字段前缀，例如："user."，"order."。. md5:1fc1d00029c03395
}

// Builder 创建并返回一个 WhereBuilder。请注意，Builder 是线程安全的。. md5:d2708a694ae36dfa
func (m *Model) Builder() *WhereBuilder {
	b := &WhereBuilder{
		model:       m,
		whereHolder: make([]WhereHolder, 0),
	}
	return b
}

// getBuilder 创建并返回当前WhereBuilder的克隆体WhereBuilder. md5:b77a9fb73a67bab3
func (b *WhereBuilder) getBuilder() *WhereBuilder {
	return b.Clone()
}

// Clone 克隆并返回一个与当前 WhereBuilder 相同的副本。. md5:d6ddf2152377c7f0
func (b *WhereBuilder) Clone() *WhereBuilder {
	newBuilder := b.model.Builder()
	newBuilder.whereHolder = make([]WhereHolder, len(b.whereHolder))
	copy(newBuilder.whereHolder, b.whereHolder)
	return newBuilder
}

// Build构建当前的WhereBuilder，并返回条件字符串和参数。. md5:08aa1af8cbe06d71
func (b *WhereBuilder) Build() (conditionWhere string, conditionArgs []interface{}) {
	var (
		ctx                         = b.model.GetCtx()
		autoPrefix                  = b.model.getAutoPrefix()
		tableForMappingAndFiltering = b.model.tables
	)
	if len(b.whereHolder) > 0 {
		for _, holder := range b.whereHolder {
			if holder.Prefix == "" {
				holder.Prefix = autoPrefix
			}
			switch holder.Operator {
			case whereHolderOperatorWhere, whereHolderOperatorAnd:
				newWhere, newArgs := formatWhereHolder(ctx, b.model.db, formatWhereHolderInput{
					WhereHolder: holder,
					OmitNil:     b.model.option&optionOmitNilWhere > 0,
					OmitEmpty:   b.model.option&optionOmitEmptyWhere > 0,
					Schema:      b.model.schema,
					Table:       tableForMappingAndFiltering,
				})
				if len(newWhere) > 0 {
					if len(conditionWhere) == 0 {
						conditionWhere = newWhere
					} else if conditionWhere[0] == '(' {
						conditionWhere = fmt.Sprintf(`%s AND (%s)`, conditionWhere, newWhere)
					} else {
						conditionWhere = fmt.Sprintf(`(%s) AND (%s)`, conditionWhere, newWhere)
					}
					conditionArgs = append(conditionArgs, newArgs...)
				}

			case whereHolderOperatorOr:
				newWhere, newArgs := formatWhereHolder(ctx, b.model.db, formatWhereHolderInput{
					WhereHolder: holder,
					OmitNil:     b.model.option&optionOmitNilWhere > 0,
					OmitEmpty:   b.model.option&optionOmitEmptyWhere > 0,
					Schema:      b.model.schema,
					Table:       tableForMappingAndFiltering,
				})
				if len(newWhere) > 0 {
					if len(conditionWhere) == 0 {
						conditionWhere = newWhere
					} else if conditionWhere[0] == '(' {
						conditionWhere = fmt.Sprintf(`%s OR (%s)`, conditionWhere, newWhere)
					} else {
						conditionWhere = fmt.Sprintf(`(%s) OR (%s)`, conditionWhere, newWhere)
					}
					conditionArgs = append(conditionArgs, newArgs...)
				}
			}
		}
	}
	return
}

// convertWhereBuilder 将参数 `where` 转换为条件字符串和参数，如果 `where` 也是一个 WhereBuilder。. md5:a6141391e787f1ad
func (b *WhereBuilder) convertWhereBuilder(where interface{}, args []interface{}) (newWhere interface{}, newArgs []interface{}) {
	var builder *WhereBuilder
	switch v := where.(type) {
	case WhereBuilder:
		builder = &v

	case *WhereBuilder:
		builder = v
	}
	if builder != nil {
		conditionWhere, conditionArgs := builder.Build()
		if conditionWhere != "" && (len(b.whereHolder) == 0 || len(builder.whereHolder) > 1) {
			conditionWhere = "(" + conditionWhere + ")"
		}
		return conditionWhere, conditionArgs
	}
	return where, args
}
