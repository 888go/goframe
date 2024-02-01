
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
// Split splits string `str` by a string `delimiter`, to an array.
<原文结束>

# <翻译开始>
// Split 函数通过一个分隔符字符串 `delimiter` 将字符串 `str` 分割成一个数组。
# <翻译结束>


<原文开始>
// SplitAndTrim splits string `str` by a string `delimiter` to an array,
// and calls Trim to every element of this array. It ignores the elements
// which are empty after Trim.
<原文结束>

# <翻译开始>
// SplitAndTrim通过字符串`delimiter`将字符串`str`分割成一个数组，
// 然后对数组中的每个元素调用Trim方法。它会忽略经过Trim处理后为空的元素。
# <翻译结束>


<原文开始>
// Join concatenates the elements of `array` to create a single string. The separator string
// `sep` is placed between elements in the resulting string.
<原文结束>

# <翻译开始>
// Join 函数将 `array` 中的元素连接起来生成一个单一的字符串。在生成的字符串中，分隔符字符串 `sep` 会被放置在各个元素之间。
# <翻译结束>


<原文开始>
// JoinAny concatenates the elements of `array` to create a single string. The separator string
// `sep` is placed between elements in the resulting string.
//
// The parameter `array` can be any type of slice, which be converted to string array.
<原文结束>

# <翻译开始>
// JoinAny 通过连接 `array` 中的元素来创建一个单一字符串。分隔符字符串
// `sep` 将会被放置在结果字符串中各元素之间。
//
// 参数 `array` 可以是任意类型的切片，只要它可以转换为字符串数组。
# <翻译结束>


<原文开始>
// Explode splits string `str` by a string `delimiter`, to an array.
// See http://php.net/manual/en/function.explode.php.
<原文结束>

# <翻译开始>
// Explode 将字符串 `str` 通过指定的分隔符 `delimiter` 进行拆分，结果存入数组中。
// 参考：http://php.net/manual/en/function.explode.php.
# <翻译结束>


<原文开始>
// Implode joins array elements `pieces` with a string `glue`.
// http://php.net/manual/en/function.implode.php
<原文结束>

# <翻译开始>
// Implode 函数通过字符串 `glue` 连接数组元素 `pieces`。
// 参考：http://php.net/manual/en/function.implode.php
# <翻译结束>


<原文开始>
// ChunkSplit splits a string into smaller chunks.
// Can be used to split a string into smaller chunks which is useful for
// e.g. converting BASE64 string output to match RFC 2045 semantics.
// It inserts end every chunkLen characters.
// It considers parameter `body` and `end` as unicode string.
<原文结束>

# <翻译开始>
// ChunkSplit 将一个字符串分割成更小的块。
// 可用于将字符串分割成较小的块，这对于例如将 BASE64 字符串输出转换为匹配 RFC 2045 语义非常有用。
// 它会在每 chunkLen 个字符后插入 end。
// 它将参数 `body` 和 `end` 视为 unicode 字符串。
# <翻译结束>


<原文开始>
// Fields returns the words used in a string as slice.
<原文结束>

# <翻译开始>
// Fields返回字符串中使用的单词作为一个切片。
# <翻译结束>

