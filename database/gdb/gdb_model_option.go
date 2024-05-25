// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

const (
	optionOmitNil             = optionOmitNilWhere | optionOmitNilData
	optionOmitEmpty           = optionOmitEmptyWhere | optionOmitEmptyData
	optionOmitNilDataInternal = optionOmitNilData | optionOmitNilDataList // 此选项仅用于内部的 ForDao 功能。. md5:0473e93966fb8e20
	optionOmitEmptyWhere      = 1 << iota                                 // 8
	optionOmitEmptyData                                                   // 16
	optionOmitNilWhere                                                    // 32
	optionOmitNilData                                                     // 64
	optionOmitNilDataList                                                 // 128
)

// OmitEmpty 为模型设置了 optionOmitEmpty 选项，该选项会自动过滤掉
// 数据和 where 参数中的 `空值`。
// md5:bf1dc800704b3324
func (m *Model) OmitEmpty() *Model {
	model := m.getModel()
	model.option = model.option | optionOmitEmpty
	return model
}

// OmitEmptyWhere 为模型设置 optionOmitEmptyWhere 选项，该选项会自动过滤掉 Where/Having 参数中的"空"值。
//
// 示例：
//
// 	Where("id", []int{}).All()             -> SELECT xxx FROM xxx WHERE 0=1
// 	Where("name", "").All()                -> SELECT xxx FROM xxx WHERE `name`=''
// 	OmitEmpty().Where("id", []int{}).All() -> SELECT xxx FROM xxx
// 	OmitEmpty().("name", "").All()         -> SELECT xxx FROM xxx.
// md5:df62f4199a96c566
func (m *Model) OmitEmptyWhere() *Model {
	model := m.getModel()
	model.option = model.option | optionOmitEmptyWhere
	return model
}

// OmitEmptyData 为模型设置 optionOmitEmptyData 选项，该选项会自动过滤掉 Data 参数中的空值。
// md5:56dab615155b1550
func (m *Model) OmitEmptyData() *Model {
	model := m.getModel()
	model.option = model.option | optionOmitEmptyData
	return model
}

// OmitNil为模型设置optionOmitNil选项，该选项会自动过滤掉`nil`值的数据和where参数。
// md5:d24d4fb4b4f59068
func (m *Model) OmitNil() *Model {
	model := m.getModel()
	model.option = model.option | optionOmitNil
	return model
}

// OmitNilWhere 为模型设置了 optionOmitNilWhere 选项，该选项会自动过滤掉
// Where/Having 参数中的 `nil` 值。
// md5:b5927ba5d926adaf
func (m *Model) OmitNilWhere() *Model {
	model := m.getModel()
	model.option = model.option | optionOmitNilWhere
	return model
}

// OmitNilData 为模型设置 optionOmitNilData 选项，该选项会自动过滤掉 Data 参数中的 `nil` 值。
// md5:e6503d524a0d8d31
func (m *Model) OmitNilData() *Model {
	model := m.getModel()
	model.option = model.option | optionOmitNilData
	return model
}
