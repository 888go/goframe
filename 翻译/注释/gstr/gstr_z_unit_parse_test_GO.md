
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
// v[][a]=m&v[][a]=b => map["v"]:[{"a":"m"},{"a":"b"}]
<原文结束>

# <翻译开始>
// 原始注释：v[][a]=m&v[][a]=b => map["v"]:[{"a":"m"},{"a":"b"}]
// 中文注释：当表达式为 v[][a]=m 与 v[][a]=b 时，可以转换为映射形式，
// 即 map 中键为 "v" 的值是一个切片，切片包含两个结构体元素，每个结构体中字段 a 的值分别为 "m" 和 "b"：
// map["v"] := [{"a": "m"}, {"a": "b"}]
# <翻译结束>







<原文开始>
// v[][a]=m&v[][b]=b => map["v"]:[{"a":"m","b":"b"}]
<原文结束>

# <翻译开始>
// 原始注释：v[][a]=m&v[][b]=b => map["v"]:[{"a":"m","b":"b"}]
// 翻译后的中文注释：
// 当表达式为 v[][a]=m 且 v[][b]=b 时，转换为映射格式为 map["v"]:[{"a":"m","b":"b"}]
// 其中，v 是一个二维切片或切片，"a" 和 "m"、"b" 和 "b" 分别表示键值对，
// 将这些键值对组合在内嵌的 JSON 对象中，并以切片形式存储在 map 的 "v" 键下。
# <翻译结束>


<原文开始>
// go test *.go -bench=".*"
<原文结束>

# <翻译开始>
// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试
# <翻译结束>

