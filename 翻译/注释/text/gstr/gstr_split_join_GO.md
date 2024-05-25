
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
// Split splits string `str` by a string `delimiter`, to an array.
<原文结束>

# <翻译开始>
// Split将字符串`str`按照字符串`delimiter`进行分割，返回一个数组。 md5:905c146c396933a0
# <翻译结束>


<原文开始>
// SplitAndTrim splits string `str` by a string `delimiter` to an array,
// and calls Trim to every element of this array. It ignores the elements
// which are empty after Trim.
<原文结束>

# <翻译开始>
// SplitAndTrim 通过字符串 `delimiter` 将字符串 `str` 分割成一个数组，然后对数组中的每个元素调用 Trim 函数。它会忽略 Trim 后为空的元素。
// md5:20d7e1d120928c19
# <翻译结束>


<原文开始>
// Join concatenates the elements of `array` to create a single string. The separator string
// `sep` is placed between elements in the resulting string.
<原文结束>

# <翻译开始>
// Join 函数将 `array` 中的元素连接起来以创建一个单一的字符串。分隔符字符串 `sep` 会被放置在结果字符串中各元素之间。
// md5:c796a29e291a2864
# <翻译结束>


<原文开始>
// JoinAny concatenates the elements of `array` to create a single string. The separator string
// `sep` is placed between elements in the resulting string.
//
// The parameter `array` can be any type of slice, which be converted to string array.
<原文结束>

# <翻译开始>
// JoinAny 将 `array` 中的元素连接成一个单一的字符串。在结果字符串中，元素之间由分隔符字符串 `sep` 分隔。
//
// 参数 `array` 可以是任何类型的切片，它将被转换为字符串数组。
// md5:fc531415278a603b
# <翻译结束>


<原文开始>
// Explode splits string `str` by a string `delimiter`, to an array.
// See http://php.net/manual/en/function.explode.php.
<原文结束>

# <翻译开始>
// Explode 函数将字符串 `str` 按照字符串 `delimiter` 进行分割，返回一个数组。
// 参考：http://php.net/manual/zh/function.explode.php。
// md5:28fb7a55d9ec56dc
# <翻译结束>


<原文开始>
// Implode joins array elements `pieces` with a string `glue`.
// http://php.net/manual/en/function.implode.php
<原文结束>

# <翻译开始>
// Implode 将数组元素 `pieces` 用字符串 `glue` 进行连接。// 参考：http://php.net/manual/zh/function.implode.php
// md5:e20b4f4a627b156b
# <翻译结束>


<原文开始>
// ChunkSplit splits a string into smaller chunks.
// Can be used to split a string into smaller chunks which is useful for
// e.g. converting BASE64 string output to match RFC 2045 semantics.
// It inserts end every chunkLen characters.
// It considers parameter `body` and `end` as unicode string.
<原文结束>

# <翻译开始>
// ChunkSplit 将字符串分割成更小的块。
// 可用于将字符串分割成更小的块，这对于
// 例如，将BASE64字符串输出转换为符合RFC 2045语义的情况非常有用。
// 它会在每chunkLen个字符后插入结束标记`end`。
// 它将参数`body`和`end`视为Unicode字符串处理。
// md5:94b7c0d7df7ca2e3
# <翻译结束>


<原文开始>
// Fields returns the words used in a string as slice.
<原文结束>

# <翻译开始>
// Fields 将字符串中的单词返回为一个切片。 md5:b66b97aa739d583c
# <翻译结束>

