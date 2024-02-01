// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb
import (
	"fmt"
	)
// WhereBuilder 在一个组合中持有多个条件语句。
type WhereBuilder struct {
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
func (m *Model) Builder() *WhereBuilder {
	b := &WhereBuilder{
		model:       m,
		whereHolder: make([]WhereHolder, 0),
	}
	return b
}

// getBuilder 创建并返回当前WhereBuilder的一个克隆副本的WhereBuilder
func (b *WhereBuilder) getBuilder() *WhereBuilder {
	return b.Clone()
}

// Clone 克隆并返回当前 WhereBuilder 的副本。
func (b *WhereBuilder) Clone() *WhereBuilder {
	newBuilder := b.model.Builder()
	newBuilder.whereHolder = make([]WhereHolder, len(b.whereHolder))
	copy(newBuilder.whereHolder, b.whereHolder)
	return newBuilder
}

// Build 函数用于构建当前的 WhereBuilder，并返回条件字符串及参数。
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

// convertWhereBuilder 将参数 `where` 转换为条件字符串及参数，如果 `where` 也是一个 WhereBuilder 对象。
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
