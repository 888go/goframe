
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
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457
# <翻译结束>


<原文开始>
// SetGlobalAttributesOption binds the global attributes to certain instrument.
<原文结束>

# <翻译开始>
// SetGlobalAttributesOption 将全局属性绑定到特定的仪表。 md5:5cba96bea01d2134
# <翻译结束>


<原文开始>
// Instrument specifies the instrument name.
<原文结束>

# <翻译开始>
// Instrument 指定乐器名称。 md5:9d3f75d7c5acf5ce
# <翻译结束>


<原文开始>
// Instrument specifies the instrument version.
<原文结束>

# <翻译开始>
// Instrument 指定仪器版本。 md5:08f1ad86326ce5c0
# <翻译结束>


<原文开始>
	// InstrumentPattern specifies instrument by regular expression on Instrument name.
	// Example:
	// 1. given `.+` will match all instruments.
	// 2. given `github.com/gogf/gf.+` will match all goframe instruments.
<原文结束>

# <翻译开始>
	// InstrumentPattern 通过正则表达式指定要操作的Instrument名称。
	// 示例：
	// 1. 如果设置为`.+`，将匹配所有Instrument。
	// 2. 如果设置为`github.com/gogf/gf.+`，将匹配所有goframe相关的Instrument。
	// md5:a3225129ad31cbb0
# <翻译结束>


<原文开始>
// GetGlobalAttributesOption binds the global attributes to certain instrument.
<原文结束>

# <翻译开始>
// GetGlobalAttributesOption 将全局属性绑定到特定的仪表。 md5:b534095d7e4c28c6
# <翻译结束>


<原文开始>
// globalAttributes stores the global attributes to a map.
<原文结束>

# <翻译开始>
	// globalAttributes 将全局属性存储到一个映射中。 md5:e8b73fe60d039913
# <翻译结束>


<原文开始>
// SetGlobalAttributes appends global attributes according `SetGlobalAttributesOption`.
// It appends global attributes to all metrics if given `SetGlobalAttributesOption` is empty.
// It appends global attributes to certain instrument by given `SetGlobalAttributesOption`.
<原文结束>

# <翻译开始>
// SetGlobalAttributes 根据 `SetGlobalAttributesOption` 添加全局属性。如果给定的 `SetGlobalAttributesOption` 为空，它将向所有指标添加全局属性。如果提供了特定的 `SetGlobalAttributesOption`，它将向指定的度量添加全局属性。
// md5:5ba03a1e3d761b95
# <翻译结束>


<原文开始>
// GetGlobalAttributes retrieves and returns the global attributes by `GetGlobalAttributesOption`.
// It returns the global attributes if given `GetGlobalAttributesOption` is empty.
// It returns global attributes of certain instrument if `GetGlobalAttributesOption` is not empty.
<原文结束>

# <翻译开始>
// GetGlobalAttributes 通过 `GetGlobalAttributesOption` 获取并返回全局属性。
// 如果给定的 `GetGlobalAttributesOption` 为空，它将返回所有全局属性。
// 如果 `GetGlobalAttributesOption` 不为空，它将返回特定仪器的全局属性。
// md5:8327524dc9d44419
# <翻译结束>

