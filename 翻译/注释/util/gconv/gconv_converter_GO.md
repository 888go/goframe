
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
// customConverters for internal converter storing.
<原文结束>

# <翻译开始>
// customConverters 用于内部转换器存储。
# <翻译结束>


<原文开始>
// RegisterConverter to register custom converter.
// It must be registered before you use this custom converting feature.
// It is suggested to do it in boot.
//
// Note:
//  1. The parameter `fn` must be defined as pattern `func(T1) (T2, error)`.
//     It will convert type `T1` to type `T2`.
//  2. The `T1` should not be type of pointer, but the `T2` should be type of pointer.
<原文结束>

# <翻译开始>
// RegisterConverter 用于注册自定义转换器。
// 在使用此自定义转换功能之前，必须先进行注册。
// 建议在初始化阶段完成此操作。
// 注意：
// 1. 参数 `fn` 必须定义为模式 `func(T1) (T2, error)`。
//    它将把类型 `T1` 转换为类型 `T2`。
// 2. `T1` 不应为指针类型，但 `T2` 应为指针类型。
# <翻译结束>


<原文开始>
// The Key and Value of the converter map should not be pointer.
<原文结束>

# <翻译开始>
// 转换器映射中的键和值不应为指针类型。
# <翻译结束>


<原文开始>
// callCustomConverter call the custom converter. It will try some possible type.
<原文结束>

# <翻译开始>
// callCustomConverter 调用自定义转换器。它会尝试一些可能的类型。
# <翻译结束>


<原文开始>
// firstly, it searches the map by input parameter type.
<原文结束>

# <翻译开始>
// 首先，通过输入参数类型搜索映射。
# <翻译结束>


<原文开始>
	// secondly, it searches the input parameter type map
	// and finds the result converter function by the output parameter type.
<原文结束>

# <翻译开始>
// 其次，它在输入参数类型映射中搜索
// 并通过输出参数类型找到结果转换函数。
# <翻译结束>


<原文开始>
// Converter function calling.
<原文结束>

# <翻译开始>
// 转换器函数调用。
# <翻译结束>


<原文开始>
// The `result[0]` is a pointer.
<原文结束>

# <翻译开始>
// `result[0]` 是一个指针。
# <翻译结束>

