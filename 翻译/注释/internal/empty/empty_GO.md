
<原文开始>
// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。 md5:c14c707c81272457
# <翻译结束>


<原文开始>
// Package empty provides functions for checking empty/nil variables.
<原文结束>

# <翻译开始>
// Package empty 提供检查空/空指针变量的函数。 md5:4be7b468c813f750
# <翻译结束>


<原文开始>
// iString is used for type assert api for String().
<原文结束>

# <翻译开始>
// iString 用于类型断言API，用于String()。 md5:8ec0af717c4f530e
# <翻译结束>


<原文开始>
// iInterfaces is used for type assert api for Interfaces.
<原文结束>

# <翻译开始>
// iInterfaces 用于接口类型的断言API。 md5:9162512bdb64ee64
# <翻译结束>


<原文开始>
// iMapStrAny is the interface support for converting struct parameter to map.
<原文结束>

# <翻译开始>
// iMapStrAny 是一个接口，支持将结构体参数转换为映射。 md5:cfd4642c77fca6ec
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
// 如果 `value` 为以下情况，函数返回 true：0, nil, false, "", slice/映射/通道的长度为0，
// 否则返回 false。
//
// 参数 `traceSource` 用于在 `value` 是指向指针的指针类型时追踪源变量。
// 当 `traceSource` 为 true 且源变量为空时，返回 true。
// 注意，这可能使用反射功能，可能会稍微影响性能。 md5:343856f448e80aef
# <翻译结束>


<原文开始>
	// It firstly checks the variable as common types using assertion to enhance the performance,
	// and then using reflection.
<原文结束>

# <翻译开始>
	// 它首先使用断言检查变量为常见类型，以提高性能，然后使用反射。 md5:9722a28f813b5ddb
# <翻译结束>


<原文开始>
// Finally, using reflect.
<原文结束>

# <翻译开始>
		// 最后，使用反射。 md5:e4ce8ad5b39b80cd
# <翻译结束>


<原文开始>
			// =========================
			// Common interfaces checks.
			// =========================
<原文结束>

# <翻译开始>
			// =========================
			// 公共接口检查。
			// ========================= md5:e561bbb4afe04dee
# <翻译结束>


<原文开始>
// IsNil checks whether given `value` is nil, especially for interface{} type value.
// Parameter `traceSource` is used for tracing to the source variable if given `value` is type of pointer
// that also points to a pointer. It returns nil if the source is nil when `traceSource` is true.
// Note that it might use reflect feature which affects performance a little.
<原文结束>

# <翻译开始>
// IsNil 检查给定的 `value` 是否为 nil，特别是对于 interface{} 类型的值。
// 参数 `traceSource` 用于在给定的 `value` 是指向指针的指针类型时，追踪到源变量。
// 当 `traceSource` 为真且源为 nil 时，它会返回 nil。
// 注意，该函数可能使用反射功能，这可能稍微影响性能。 md5:c12efd8c176fc73a
# <翻译结束>

