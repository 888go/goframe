
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
// customConverters for internal converter storing.
<原文结束>

# <翻译开始>
// customConverters 用于内部转换器的存储。 md5:eb816f1844daac79
# <翻译结束>


<原文开始>
// RegisterConverter to register custom converter.
// It must be registered before you use this custom converting feature.
// It is suggested to do it in boot procedure of the process.
//
// Note:
//  1. The parameter `fn` must be defined as pattern `func(T1) (T2, error)`.
//     It will convert type `T1` to type `T2`.
//  2. The `T1` should not be type of pointer, but the `T2` should be type of pointer.
<原文结束>

# <翻译开始>
// RegisterConverter 用于注册自定义转换器。
// 必须在使用此自定义转换功能之前进行注册。
// 建议在进程的启动程序中执行此操作。
//
// 注意：
//  1. 参数 `fn` 必须定义为模式 `func(T1) (T2, error)`。
//     它将类型 `T1` 转换为类型 `T2`。
//  2. `T1` 不应为指针类型，但 `T2` 应为指针类型。
// md5:8fbaa372837e6d8c
# <翻译结束>


<原文开始>
// The Key and Value of the converter map should not be pointer.
<原文结束>

# <翻译开始>
// 转换映射的键和值不应该是指针。 md5:79bb068f1985b81a
# <翻译结束>


<原文开始>
// firstly, it searches the map by input parameter type.
<原文结束>

# <翻译开始>
// 首先，它通过输入参数类型在映射中搜索。 md5:019f9d8418285668
# <翻译结束>


<原文开始>
// Might be **struct, which is support as designed.
<原文结束>

# <翻译开始>
// 可能是**struct，这是设计上支持的。 md5:cb1f21754e39c3a1
# <翻译结束>


<原文开始>
	// secondly, it searches the input parameter type map
	// and finds the result converter function by the output parameter type.
<原文结束>

# <翻译开始>
	// 其次，它会在输入参数类型映射中搜索
	// 并通过输出参数类型找到结果转换函数。
	// md5:3781290987232f09
# <翻译结束>


<原文开始>
// callCustomConverter call the custom converter. It will try some possible type.
<原文结束>

# <翻译开始>
// callCustomConverter 调用自定义转换器。它会尝试一些可能的类型。 md5:44d83ddc5510baed
# <翻译结束>


<原文开始>
// Converter function calling.
<原文结束>

# <翻译开始>
// 转换函数调用。 md5:1780fb4f627f751d
# <翻译结束>


<原文开始>
// The `result[0]` is a pointer.
<原文结束>

# <翻译开始>
// `result[0]`是一个指针。 md5:6505f86b6cd1e865
# <翻译结束>

