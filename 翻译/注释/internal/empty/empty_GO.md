
<原文开始>
// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame gf 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
# <翻译结束>


<原文开始>
// Package empty provides functions for checking empty/nil variables.
<原文结束>

# <翻译开始>
// 包 empty 提供了检查空/nil 变量的函数。
# <翻译结束>


<原文开始>
// iString is used for type assert api for String().
<原文结束>

# <翻译开始>
// iString 用于在进行类型断言时，配合 String() 方法使用。
# <翻译结束>


<原文开始>
// iInterfaces is used for type assert api for Interfaces.
<原文结束>

# <翻译开始>
// iInterfaces 用于对 Interfaces 进行类型断言的 API。
# <翻译结束>


<原文开始>
// iMapStrAny is the interface support for converting struct parameter to map.
<原文结束>

# <翻译开始>
// iMapStrAny 是支持将结构体参数转换为映射的接口。
# <翻译结束>


<原文开始>
// IsEmpty checks whether given `value` empty.
// It returns true if `value` is in: 0, nil, false, "", len(slice/map/chan) == 0,
// or else it returns false.
//
// The parameter `traceSource` is used for tracing to the source variable if given `value` is type of pointer
// that also points to a pointer. It returns true if the source is empty when `traceSource` is true.
// Note that it might use reflect feature which affects performance a little.
<原文结束>

# <翻译开始>
// IsEmpty 检查给定的 `value` 是否为空。
// 当 `value` 为以下情形之一时，返回 true：0, nil, false, "", 切片/映射/通道长度为0，
// 否则返回 false。
//
// 参数 `traceSource` 用于在 `value` 类型为指针且指向另一个指针时追踪到源变量。
// 如果 `traceSource` 为 true 并且源变量为空，则它会返回 true。
// 注意，这可能会使用 reflect 特性，对性能稍有影响。
# <翻译结束>


<原文开始>
	// It firstly checks the variable as common types using assertion to enhance the performance,
	// and then using reflection.
<原文结束>

# <翻译开始>
// 它首先通过断言检查变量作为常见类型以提升性能，
// 然后再使用反射进行处理。
# <翻译结束>







<原文开始>
			// =========================
			// Common interfaces checks.
			// =========================
<原文结束>

# <翻译开始>
// ========================================
// 常用接口检查。
// ========================================
# <翻译结束>


<原文开始>
// IsNil checks whether given `value` is nil, especially for interface{} type value.
// Parameter `traceSource` is used for tracing to the source variable if given `value` is type of pointer
// that also points to a pointer. It returns nil if the source is nil when `traceSource` is true.
// Note that it might use reflect feature which affects performance a little.
<原文结束>

# <翻译开始>
// IsNil 检查给定的 `value` 是否为 nil，特别是对 interface{} 类型的值。
// 参数 `traceSource` 用于在 `value` 是指针类型且指向另一个指针时，追踪到源变量。如果源变量为 nil 并且 `traceSource` 为真，则返回 nil。
// 注意，该函数可能会使用 reflect 特性，这会对性能造成一定的影响。
# <翻译结束>


<原文开始>
// Finally, using reflect.
<原文结束>

# <翻译开始>
// 最后，使用 reflect 包
# <翻译结束>

