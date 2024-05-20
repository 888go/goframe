
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// this option is used internally only for ForDao feature.
<原文结束>

# <翻译开始>
// 此选项仅用于内部的 ForDao 功能。. md5:0473e93966fb8e20
# <翻译结束>


<原文开始>
// OmitEmpty sets optionOmitEmpty option for the model, which automatically filers
// the data and where parameters for `empty` values.
<原文结束>

# <翻译开始>
// OmitEmpty 为模型设置了 optionOmitEmpty 选项，该选项会自动过滤掉
// 数据和 where 参数中的 `空值`。
// md5:bf1dc800704b3324
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
// OmitEmptyWhere 为模型设置 optionOmitEmptyWhere 选项，该选项会自动过滤掉 Where/Having 参数中的"空"值。
//
// 示例：
//
// 	Where("id", []int{}).All()             -> SELECT xxx FROM xxx WHERE 0=1
// 	Where("name", "").All()                -> SELECT xxx FROM xxx WHERE `name`=''
// 	OmitEmpty().Where("id", []int{}).All() -> SELECT xxx FROM xxx
// 	OmitEmpty().("name", "").All()         -> SELECT xxx FROM xxx.
// md5:df62f4199a96c566
# <翻译结束>


<原文开始>
// OmitEmptyData sets optionOmitEmptyData option for the model, which automatically filers
// the Data parameters for `empty` values.
<原文结束>

# <翻译开始>
// OmitEmptyData 为模型设置 optionOmitEmptyData 选项，该选项会自动过滤掉 Data 参数中的空值。
// md5:56dab615155b1550
# <翻译结束>


<原文开始>
// OmitNil sets optionOmitNil option for the model, which automatically filers
// the data and where parameters for `nil` values.
<原文结束>

# <翻译开始>
// OmitNil为模型设置optionOmitNil选项，该选项会自动过滤掉`nil`值的数据和where参数。
// md5:d24d4fb4b4f59068
# <翻译结束>


<原文开始>
// OmitNilWhere sets optionOmitNilWhere option for the model, which automatically filers
// the Where/Having parameters for `nil` values.
<原文结束>

# <翻译开始>
// OmitNilWhere 为模型设置了 optionOmitNilWhere 选项，该选项会自动过滤掉
// Where/Having 参数中的 `nil` 值。
// md5:b5927ba5d926adaf
# <翻译结束>


<原文开始>
// OmitNilData sets optionOmitNilData option for the model, which automatically filers
// the Data parameters for `nil` values.
<原文结束>

# <翻译开始>
// OmitNilData 为模型设置 optionOmitNilData 选项，该选项会自动过滤掉 Data 参数中的 `nil` 值。
// md5:e6503d524a0d8d31
# <翻译结束>

