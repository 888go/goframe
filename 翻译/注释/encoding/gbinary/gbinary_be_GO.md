
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
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// BeEncode encodes one or multiple `values` into bytes using BigEndian.
// It uses type asserting checking the type of each value of `values` and internally
// calls corresponding converting function do the bytes converting.
//
// It supports common variable type asserting, and finally it uses fmt.Sprintf converting
// value to string and then to bytes.
<原文结束>

# <翻译开始>
// BeEncode 采用大端序（BigEndian）将一个或多个 `values` 值编码为字节。
// 它通过类型断言检查 `values` 中每个值的类型，并在内部调用相应的转换函数进行字节转换。
//
// 它支持常见的变量类型断言，最终使用 fmt.Sprintf 将值转换为字符串，然后再转为字节。
# <翻译结束>


<原文开始>
// BeFillUpSize fills up the bytes `b` to given length `l` using big BigEndian.
//
// Note that it creates a new bytes slice by copying the original one to avoid changing
// the original parameter bytes.
<原文结束>

# <翻译开始>
// BeFillUpSize 使用大端序将字节切片`b`填充到给定长度`l`。
//
// 注意：为了防止改变原始参数字节，它会通过复制原始字节切片创建一个新的字节切片。
# <翻译结束>

