
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
// LeEncode encodes one or multiple `values` into bytes using LittleEndian.
// It uses type asserting checking the type of each value of `values` and internally
// calls corresponding converting function do the bytes converting.
//
// It supports common variable type asserting, and finally it uses fmt.Sprintf converting
// value to string and then to bytes.
<原文结束>

# <翻译开始>
// LeEncode 使用小端字节序编码一个或多个 `values` 成字节。
// 它通过类型断言检查 `values` 中每个值的类型，并在内部调用相应的转换函数来进行字节转换。
//
// 它支持常见的变量类型断言，最终使用 fmt.Sprintf 将值转换为字符串，再转换为字节。
// md5:4210ce2214f05499
# <翻译结束>


<原文开始>
// LeFillUpSize fills up the bytes `b` to given length `l` using LittleEndian.
//
// Note that it creates a new bytes slice by copying the original one to avoid changing
// the original parameter bytes.
<原文结束>

# <翻译开始>
// LeFillUpSize 使用LittleEndian字节序将字节`b`填充到给定长度`l`。
//
// 请注意，它通过复制原始字节切片创建一个新的字节切片，以避免修改原始参数字节。
// md5:9ac3ba6f3d5c0177
# <翻译结束>

