// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"fmt"
)

// WhereBuilder 在一个组合中持有多个条件语句。
type X组合条件 struct {
	model       *Model        // WhereBuilder 应该绑定到特定的 Model。
	whereHolder []WhereHolder // 条件字符串，用于where操作。
}

// WhereHolder 是用于准备 where 条件的占位符。
type WhereHolder struct {
	Type     string        // 此持有者的类型。
	Operator int           // 此持有者的操作员。
	Where    interface{}   // Where 参数，通常可以是 string、map 或 struct 类型。
	Args     []interface{} // "where" 参数的对应条件。
	Prefix   string        // 字段前缀，例如："user.", "order."
}

// Builder 创建并返回一个 WhereBuilder。请注意，该构建器支持链式调用，即链式安全。
func (m *Model) X创建组合条件() *X组合条件 {
	b := &X组合条件{
		model:       m,
		whereHolder: make([]WhereHolder, 0),
	}
	return b
}

// getBuilder 创建并返回当前WhereBuilder的一个克隆副本的WhereBuilder
func (b *X组合条件) getBuilder() *X组合条件 {
	return b.X取副本()
}

// Clone 克隆并返回当前 WhereBuilder 的副本。
func (b *X组合条件) X取副本() *X组合条件 {
	newBuilder := b.model.X创建组合条件()
	newBuilder.whereHolder = make([]WhereHolder, len(b.whereHolder))
	copy(newBuilder.whereHolder, b.whereHolder)
	return newBuilder
}

// Build 函数用于构建当前的 WhereBuilder，并返回条件字符串及参数。
func (b *X组合条件) X生成条件字符串及参数() (条件字符串 string, 参数 []interface{}) {
	var (
		ctx                         = b.model.X取上下文对象()
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
					if len(条件字符串) == 0 {
						条件字符串 = newWhere
					} else if 条件字符串[0] == '(' {
						条件字符串 = fmt.Sprintf(`%s AND (%s)`, 条件字符串, newWhere)
					} else {
						条件字符串 = fmt.Sprintf(`(%s) AND (%s)`, 条件字符串, newWhere)
					}
					参数 = append(参数, newArgs...)
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
					if len(条件字符串) == 0 {
						条件字符串 = newWhere
					} else if 条件字符串[0] == '(' {
						条件字符串 = fmt.Sprintf(`%s OR (%s)`, 条件字符串, newWhere)
					} else {
						条件字符串 = fmt.Sprintf(`(%s) OR (%s)`, 条件字符串, newWhere)
					}
					参数 = append(参数, newArgs...)
				}
			}
		}
	}
	return
}

// convertWhereBuilder 将参数 `where` 转换为条件字符串及参数，如果 `where` 也是一个 WhereBuilder 对象。
func (b *X组合条件) convertWhereBuilder(where interface{}, args []interface{}) (newWhere interface{}, newArgs []interface{}) {
	var builder *X组合条件
	switch v := where.(type) {
	case X组合条件:
		builder = &v

	case *X组合条件:
		builder = v
	}
	if builder != nil {
		conditionWhere, conditionArgs := builder.X生成条件字符串及参数()
		if conditionWhere != "" && (len(b.whereHolder) == 0 || len(builder.whereHolder) > 1) {
			conditionWhere = "(" + conditionWhere + ")"
		}
		return conditionWhere, conditionArgs
	}
	return where, args
}
