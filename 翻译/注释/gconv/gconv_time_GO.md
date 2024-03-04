
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
// Time converts `any` to time.Time.
<原文结束>

# <翻译开始>
// Time将`any`转换为time.Time类型。
# <翻译结束>







<原文开始>
// Duration converts `any` to time.Duration.
// If `any` is string, then it uses time.ParseDuration to convert it.
// If `any` is numeric, then it converts `any` as nanoseconds.
<原文结束>

# <翻译开始>
// Duration 将 `any` 转换为 time.Duration 类型。
// 如果 `any` 是字符串，那么它会使用 time.ParseDuration 来进行转换。
// 如果 `any` 是数字类型，则将 `any` 视为纳秒进行转换。
# <翻译结束>


<原文开始>
// GTime converts `any` to *gtime.Time.
// The parameter `format` can be used to specify the format of `any`.
// It returns the converted value that matched the first format of the formats slice.
// If no `format` given, it converts `any` using gtime.NewFromTimeStamp if `any` is numeric,
// or using gtime.StrToTime if `any` is string.
<原文结束>

# <翻译开始>
// GTime 将 `any` 类型转换为 *gtime.Time 类型。
// 参数 `format` 可用于指定 `any` 的格式。
// 它将返回与 formats 切片中第一个格式匹配的转换后的值。
// 如果未提供 `format`，当 `any` 为数值类型时，使用 gtime.NewFromTimeStamp 进行转换；
// 当 `any` 为字符串类型时，则使用 gtime.StrToTime 进行转换。
# <翻译结束>


<原文开始>
// Priority conversion using given format.
<原文结束>

# <翻译开始>
// 使用给定格式进行优先级转换。
# <翻译结束>


<原文开始>
// It's already this type.
<原文结束>

# <翻译开始>
// 它已经是这种类型了。
# <翻译结束>

