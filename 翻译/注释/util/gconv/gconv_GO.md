
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
// Package gconv implements powerful and convenient converting functionality for any types of variables.
//
// This package should keep much less dependencies with other packages.
<原文结束>

# <翻译开始>
// 包gconv实现了对任何类型变量的强大而便捷的转换功能。
//
// 此包应尽量减少与其他包的依赖关系。
// md5:b18f07aca2be5125
# <翻译结束>


<原文开始>
	// StructTagPriority defines the default priority tags for Map*/Struct* functions.
	// Note that, the `gconv/param` tags are used by old version of package.
	// It is strongly recommended using short tag `c/p` instead in the future.
<原文结束>

# <翻译开始>
// StructTagPriority 定义了Map*/Struct*函数的默认优先级标签。
// 注意，`gconv/param` 标签由旧版本的包使用。强烈建议未来改用简短的标签 `c/p`。
// md5:c4b7d2fe8905ed52
# <翻译结束>


<原文开始>
// Byte converts `any` to byte.
<原文结束>

# <翻译开始>
// Byte将`any`转换为byte。 md5:aeef919e3fba4f95
# <翻译结束>


<原文开始>
// Bytes converts `any` to []byte.
<原文结束>

# <翻译开始>
// Bytes 将 `any` 转换为 []byte。 md5:06125d6ba5f449a5
# <翻译结束>


<原文开始>
// Rune converts `any` to rune.
<原文结束>

# <翻译开始>
// Rune 将 `any` 转换为 rune。 md5:3459f7528861cc23
# <翻译结束>


<原文开始>
// Runes converts `any` to []rune.
<原文结束>

# <翻译开始>
// Runes将`any`转换为[]rune。 md5:25552cd961d1d6bb
# <翻译结束>


<原文开始>
// String converts `any` to string.
// It's most commonly used converting function.
<原文结束>

# <翻译开始>
// String 将 `any` 转换为字符串。它是最常用的转换函数。
// md5:722d0704c061781b
# <翻译结束>


<原文开始>
			// If the variable implements the String() interface,
			// then use that interface to perform the conversion
<原文结束>

# <翻译开始>
// 如果变量实现了String()接口，
// 则使用该接口来进行转换
// md5:08e76021f60d81ed
# <翻译结束>


<原文开始>
			// If the variable implements the Error() interface,
			// then use that interface to perform the conversion
<原文结束>

# <翻译开始>
// /* 如果该变量实现了Error()接口，
//    则使用该接口进行转换 */
// md5:7c7c512864a0b034
# <翻译结束>


<原文开始>
// Finally, we use json.Marshal to convert.
<原文结束>

# <翻译开始>
// 最后，我们使用json.Marshal进行转换。 md5:57829b67798bbc93
# <翻译结束>


<原文开始>
// Bool converts `any` to bool.
// It returns false if `any` is: false, "", 0, "false", "off", "no", empty slice/map.
<原文结束>

# <翻译开始>
// Bool 将 `any` 转换为布尔值。
// 如果 `any` 是：false，""，0，"false"，"off"，"no"，空切片/映射，则返回 false。
// md5:b9d150a8798a274a
# <翻译结束>


<原文开始>
// checkJsonAndUnmarshalUseNumber checks if given `any` is JSON formatted string value and does converting using `json.UnmarshalUseNumber`.
<原文结束>

# <翻译开始>
// checkJsonAndUnmarshalUseNumber 检查给定的 `any` 是否为 JSON 格式的字符串值，并使用 `json.UnmarshalUseNumber` 进行转换。 md5:ce3edf33e8eea76c
# <翻译结束>

