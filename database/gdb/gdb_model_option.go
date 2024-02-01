// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

const (
	optionOmitNil             = optionOmitNilWhere | optionOmitNilData
	optionOmitEmpty           = optionOmitEmptyWhere | optionOmitEmptyData
	optionOmitNilDataInternal = optionOmitNilData | optionOmitNilDataList // 此选项仅供内部使用，仅用于 ForDao 功能。
	optionOmitEmptyWhere      = 1 << iota                                 // 8
	optionOmitEmptyData                                                   // 16
	optionOmitNilWhere                                                    // 32
	optionOmitNilData                                                     // 64
	optionOmitNilDataList                                                 // 128
)

// OmitEmpty 为模型设置 optionOmitEmpty 选项，该选项会自动过滤掉数据和 `empty` 值的参数。
// optionOmitEmpty 选项，该选项会自动过滤掉数据和 `empty` 值的
func (m *Model) OmitEmpty() *Model {
	model := m.getModel()
	model.option = model.option | optionOmitEmpty
	return model
}

// OmitEmptyWhere 为模型设置 optionOmitEmptyWhere 选项，该选项会自动过滤 Where/Having 子句中 `empty` 值的参数。
//
// 示例：
//
//	Where("id", []int{}).All()             -> SELECT xxx FROM xxx WHERE 0=1 （当id为空数组时，生成一个恒不成立的条件）
//	Where("name", "").All()                -> SELECT xxx FROM xxx WHERE `name`='' （当name为空字符串时，生成相应空值条件）
//	OmitEmpty().Where("id", []int{}).All() -> SELECT xxx FROM xxx （在OmitEmpty作用下，忽略空数组id条件，直接执行查询）
//	OmitEmpty().Where("name", "").All()         -> SELECT xxx FROM xxx （在OmitEmpty作用下，忽略空字符串name条件，直接执行查询）
func (m *Model) OmitEmptyWhere() *Model {
	model := m.getModel()
	model.option = model.option | optionOmitEmptyWhere
	return model
}

// OmitEmptyData 为模型设置 optionOmitEmptyData 选项，该选项会自动过滤`空`值的 Data 参数。
func (m *Model) OmitEmptyData() *Model {
	model := m.getModel()
	model.option = model.option | optionOmitEmptyData
	return model
}

// OmitNil 为模型设置 optionOmitNil 选项，该选项会自动过滤数据和参数中的 `nil` 值。
// 这意味着在处理数据或设置条件时，将忽略值为 nil 的字段。
func (m *Model) OmitNil() *Model {
	model := m.getModel()
	model.option = model.option | optionOmitNil
	return model
}

// OmitNilWhere 为模型设置 optionOmitNilWhere 选项，该选项会自动过滤 Where/Having 条件中的 `nil` 值。
func (m *Model) OmitNilWhere() *Model {
	model := m.getModel()
	model.option = model.option | optionOmitNilWhere
	return model
}

// OmitNilData 为模型设置 optionOmitNilData 选项，该选项会自动过滤 `nil` 值的 Data 参数。
func (m *Model) OmitNilData() *Model {
	model := m.getModel()
	model.option = model.option | optionOmitNilData
	return model
}
