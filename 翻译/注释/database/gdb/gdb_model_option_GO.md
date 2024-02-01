
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// this option is used internally only for ForDao feature.
<原文结束>

# <翻译开始>
// 此选项仅供内部使用，仅用于 ForDao 功能。
# <翻译结束>


<原文开始>
// OmitEmpty sets optionOmitEmpty option for the model, which automatically filers
// the data and where parameters for `empty` values.
<原文结束>

# <翻译开始>
// OmitEmpty 为模型设置 optionOmitEmpty 选项，该选项会自动过滤掉数据和 `empty` 值的参数。
// optionOmitEmpty 选项，该选项会自动过滤掉数据和 `empty` 值的
# <翻译结束>


<原文开始>
// OmitEmptyWhere sets optionOmitEmptyWhere option for the model, which automatically filers
// the Where/Having parameters for `empty` values.
//
// Eg:
//
//	Where("id", []int{}).All()             -> SELECT xxx FROM xxx WHERE 0=1
//	Where("name", "").All()                -> SELECT xxx FROM xxx WHERE `name`=''
//	OmitEmpty().Where("id", []int{}).All() -> SELECT xxx FROM xxx
//	OmitEmpty().("name", "").All()         -> SELECT xxx FROM xxx.
<原文结束>

# <翻译开始>
// OmitEmptyWhere 为模型设置 optionOmitEmptyWhere 选项，该选项会自动过滤 Where/Having 子句中 `empty` 值的参数。
//
// 示例：
//
//	Where("id", []int{}).All()             -> SELECT xxx FROM xxx WHERE 0=1 （当id为空数组时，生成一个恒不成立的条件）
//	Where("name", "").All()                -> SELECT xxx FROM xxx WHERE `name`='' （当name为空字符串时，生成相应空值条件）
//	OmitEmpty().Where("id", []int{}).All() -> SELECT xxx FROM xxx （在OmitEmpty作用下，忽略空数组id条件，直接执行查询）
//	OmitEmpty().Where("name", "").All()         -> SELECT xxx FROM xxx （在OmitEmpty作用下，忽略空字符串name条件，直接执行查询）
# <翻译结束>


<原文开始>
// OmitEmptyData sets optionOmitEmptyData option for the model, which automatically filers
// the Data parameters for `empty` values.
<原文结束>

# <翻译开始>
// OmitEmptyData 为模型设置 optionOmitEmptyData 选项，该选项会自动过滤`空`值的 Data 参数。
# <翻译结束>


<原文开始>
// OmitNil sets optionOmitNil option for the model, which automatically filers
// the data and where parameters for `nil` values.
<原文结束>

# <翻译开始>
// OmitNil 为模型设置 optionOmitNil 选项，该选项会自动过滤数据和参数中的 `nil` 值。
// 这意味着在处理数据或设置条件时，将忽略值为 nil 的字段。
# <翻译结束>


<原文开始>
// OmitNilWhere sets optionOmitNilWhere option for the model, which automatically filers
// the Where/Having parameters for `nil` values.
<原文结束>

# <翻译开始>
// OmitNilWhere 为模型设置 optionOmitNilWhere 选项，该选项会自动过滤 Where/Having 条件中的 `nil` 值。
# <翻译结束>


<原文开始>
// OmitNilData sets optionOmitNilData option for the model, which automatically filers
// the Data parameters for `nil` values.
<原文结束>

# <翻译开始>
// OmitNilData 为模型设置 optionOmitNilData 选项，该选项会自动过滤 `nil` 值的 Data 参数。
# <翻译结束>

