
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
// Convert converts the variable `fromValue` to the type `toTypeName`, the type `toTypeName` is specified by string.
//
// The optional parameter `extraParams` is used for additional necessary parameter for this conversion.
// It supports common basic types conversion as its conversion based on type name string.
<原文结束>

# <翻译开始>
// Convert 将变量 `fromValue` 转换为类型 `toTypeName`，其中 `toTypeName` 由字符串指定。
//
// 可选参数 `extraParams` 用于提供此次转换所需的额外必要参数。
// 它支持基于类型名称字符串的基本常见类型的转换。
# <翻译结束>


<原文开始>
// ConvertWithRefer converts the variable `fromValue` to the type referred by value `referValue`.
//
// The optional parameter `extraParams` is used for additional necessary parameter for this conversion.
// It supports common basic types conversion as its conversion based on type name string.
<原文结束>

# <翻译开始>
// ConvertWithRefer 将变量 `fromValue` 转换为由值 `referValue` 所引用的类型。
//
// 可选参数 `extraParams` 用于提供本次转换所需的额外必要参数。
// 它支持基于类型名称字符串的基本常见类型的转换。
# <翻译结束>


<原文开始>
// Value that is converted from.
<原文结束>

# <翻译开始>
// 需要转换的原始值。
# <翻译结束>


<原文开始>
// Target value type name in string.
<原文结束>

# <翻译开始>
// 字符串形式的目标值类型名称
# <翻译结束>


<原文开始>
// Extra values for implementing the converting.
<原文结束>

# <翻译开始>
// 用于实现转换功能的额外值。
# <翻译结束>


<原文开始>
	// Marks that the value is already converted and set to `ReferValue`. Caller can ignore the returned result.
	// It is an attribute for internal usage purpose.
<原文结束>

# <翻译开始>
// 标记该值已转换并设置为`ReferValue`。调用者可以忽略返回的结果。
// 这是一个用于内部使用的属性。
# <翻译结束>


<原文开始>
// doConvert does commonly use types converting.
<原文结束>

# <翻译开始>
// doConvert 执行常用类型的转换。
# <翻译结束>


<原文开始>
				// Type converting for custom type pointers.
				// Eg:
				// type PayMode int
				// type Req struct{
				//     Mode *PayMode
				// }
				//
				// Struct(`{"Mode": 1000}`, &req)
<原文结束>

# <翻译开始>
// 自定义类型指针的类型转换
// 示例：
// 定义自定义类型 PayMode，其为 int 类型的别名
// 定义结构体 Req，其中包含一个指向 PayMode 类型的指针 Mode
//
// type PayMode int
// type Req struct{
//     Mode *PayMode
// }
//
// 通过 Struct 函数将 `{"Mode": 1000}` 转换并解析到 req 指针所指向的结构体中
// Struct(`{"Mode": 1000}`, &req)
# <翻译结束>







<原文开始>
// Referred value, a value in type `ToTypeName`. Note that its type might be reflect.Value.
<原文结束>

# <翻译开始>
// 指针引用的值，类型为 `ToTypeName` 的值。注意，其实际类型可能为 reflect.Value。
# <翻译结束>


<原文开始>
// Not support some kinds.
<原文结束>

# <翻译开始>
// 不支持某些类型。
# <翻译结束>

