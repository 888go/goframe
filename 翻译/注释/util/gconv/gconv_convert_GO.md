
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
// Convert converts the variable `fromValue` to the type `toTypeName`, the type `toTypeName` is specified by string.
//
// The optional parameter `extraParams` is used for additional necessary parameter for this conversion.
// It supports common basic types conversion as its conversion based on type name string.
<原文结束>

# <翻译开始>
// Convert 将变量 `fromValue` 转换为类型 `toTypeName`，其中 `toTypeName` 由字符串指定。
//
// 可选参数 `extraParams` 用于提供此转换所需的额外参数。
// 它支持基于类型名称字符串的常见基本类型转换。
// md5:e081c8fc6552be4c
# <翻译结束>


<原文开始>
// ConvertWithRefer converts the variable `fromValue` to the type referred by value `referValue`.
//
// The optional parameter `extraParams` is used for additional necessary parameter for this conversion.
// It supports common basic types conversion as its conversion based on type name string.
<原文结束>

# <翻译开始>
// ConvertWithRefer 将变量 `fromValue` 转换为由 `referValue` 指定的类型。
//
// 可选参数 `extraParams` 用于此转换所需的额外参数。它支持基于类型名称字符串的常见基本类型转换。
// md5:0badd37157c72db1
# <翻译结束>


<原文开始>
// Value that is converted from.
<原文结束>

# <翻译开始>
// 要转换的值。 md5:b9384f7def81e56d
# <翻译结束>


<原文开始>
// Target value type name in string.
<原文结束>

# <翻译开始>
// 目标值类型名称（字符串形式）。 md5:56863f5417d5b24f
# <翻译结束>


<原文开始>
// Referred value, a value in type `ToTypeName`. Note that its type might be reflect.Value.
<原文结束>

# <翻译开始>
// 引用的值，类型为`ToTypeName`。请注意，它的类型可能是`reflect.Value`。 md5:7e9c4375ec4d26f3
# <翻译结束>


<原文开始>
// Extra values for implementing the converting.
<原文结束>

# <翻译开始>
// 用于实现转换的额外值。 md5:c5e0f680118ba627
# <翻译结束>


<原文开始>
	// Marks that the value is already converted and set to `ReferValue`. Caller can ignore the returned result.
	// It is an attribute for internal usage purpose.
<原文结束>

# <翻译开始>
	// 标记该值已经转换并设置为`ReferValue`。调用者可以忽略返回的结果。
	// 这是一个用于内部使用的属性。
	// md5:91187d21c0d0ac16
# <翻译结束>


<原文开始>
// doConvert does commonly use types converting.
<原文结束>

# <翻译开始>
// doConvert 执行常用类型转换。 md5:a4f52e85ed63dbe3
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
				// 自定义类型指针的类型转换。
				// 例如：
				// type PayMode int
				// type Req struct{
				//     Mode *PayMode
				// }
				// 
				// Struct(`{"Mode": 1000}`, &req)
				// md5:d218e7f3f409c5f7
# <翻译结束>


<原文开始>
// Not support some kinds.
<原文结束>

# <翻译开始>
// 不支持某些类型。 md5:74a7c80d66154fb9
# <翻译结束>

