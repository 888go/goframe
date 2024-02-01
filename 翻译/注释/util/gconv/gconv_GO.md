
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
// Package gconv implements powerful and convenient converting functionality for any types of variables.
//
// This package should keep much less dependencies with other packages.
<原文结束>

# <翻译开始>
// Package gconv 提供了强大且方便的任意类型变量转换功能。
//
// 本包应尽量减少对其他包的依赖。
# <翻译结束>


<原文开始>
	// StructTagPriority defines the default priority tags for Map*/Struct* functions.
	// Note that, the `gconv/param` tags are used by old version of package.
	// It is strongly recommended using short tag `c/p` instead in the future.
<原文结束>

# <翻译开始>
// StructTagPriority 定义了 Map*/Struct* 函数的默认优先级标签。
// 注意，`gconv/param` 标签在旧版本包中使用。
// 强烈建议在未来使用简写标签 `c/p` 代替。
# <翻译结束>


<原文开始>
// Byte converts `any` to byte.
<原文结束>

# <翻译开始>
// Byte将`any`转换为字节。
# <翻译结束>


<原文开始>
// Bytes converts `any` to []byte.
<原文结束>

# <翻译开始>
// Bytes 将 `any` 类型转换为 []byte 类型。
# <翻译结束>


<原文开始>
// Rune converts `any` to rune.
<原文结束>

# <翻译开始>
// Rune将`any`转换为rune类型。
# <翻译结束>


<原文开始>
// Runes converts `any` to []rune.
<原文结束>

# <翻译开始>
// Runes 将 `any` 转换为 []rune 类型。
# <翻译结束>


<原文开始>
// String converts `any` to string.
// It's most commonly used converting function.
<原文结束>

# <翻译开始>
// String 将`any`转换为字符串。
// 这是最常用的转换函数。
# <翻译结束>


<原文开始>
			// If the variable implements the String() interface,
			// then use that interface to perform the conversion
<原文结束>

# <翻译开始>
// 如果变量实现了String()接口，
// 那么就使用该接口来执行转换
# <翻译结束>


<原文开始>
			// If the variable implements the Error() interface,
			// then use that interface to perform the conversion
<原文结束>

# <翻译开始>
// 如果变量实现了Error()接口，
// 那么使用该接口进行转换
# <翻译结束>


<原文开始>
// Finally, we use json.Marshal to convert.
<原文结束>

# <翻译开始>
// 最后，我们使用json.Marshal将数据转换。
# <翻译结束>


<原文开始>
// Bool converts `any` to bool.
// It returns false if `any` is: false, "", 0, "false", "off", "no", empty slice/map.
<原文结束>

# <翻译开始>
// Bool将`any`转换为布尔值。
// 当`any`为：false、空字符串、0、"false"、"off"、"no"或空切片/映射时，返回false。
# <翻译结束>


<原文开始>
// checkJsonAndUnmarshalUseNumber checks if given `any` is JSON formatted string value and does converting using `json.UnmarshalUseNumber`.
<原文结束>

# <翻译开始>
// checkJsonAndUnmarshalUseNumber 检查给定的 `any` 是否为格式化的 JSON 字符串值，并使用 `json.UnmarshalUseNumber` 进行转换。
# <翻译结束>






