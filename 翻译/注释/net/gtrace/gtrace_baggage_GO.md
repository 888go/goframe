
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
// Baggage holds the data through all tracing spans.
<原文结束>

# <翻译开始>
// Baggage 在所有追踪跨度中保存数据。
# <翻译结束>


<原文开始>
// NewBaggage creates and returns a new Baggage object from given tracing context.
<原文结束>

# <翻译开始>
// NewBaggage 从给定的追踪上下文中创建并返回一个新的 Baggage 对象。
# <翻译结束>


<原文开始>
// Ctx returns the context that Baggage holds.
<原文结束>

# <翻译开始>
// Ctx 返回 Baggage 持有的上下文。
# <翻译结束>


<原文开始>
// SetValue is a convenient function for adding one key-value pair to baggage.
// Note that it uses attribute.Any to set the key-value pair.
<原文结束>

# <翻译开始>
// SetValue 是一个方便的函数，用于向 baggage 中添加一对键值对。
// 注意，它使用 attribute.Any 来设置键值对。
# <翻译结束>


<原文开始>
// SetMap is a convenient function for adding map key-value pairs to baggage.
// Note that it uses attribute.Any to set the key-value pair.
<原文结束>

# <翻译开始>
// SetMap 是一个方便的函数，用于向 baggage 添加映射键值对。
// 注意，它使用 attribute.Any 来设置键值对。
# <翻译结束>


<原文开始>
// GetMap retrieves and returns the baggage values as map.
<原文结束>

# <翻译开始>
// GetMap 获取并以map形式返回 baggage 的值。
# <翻译结束>


<原文开始>
// GetVar retrieves value and returns a *gvar.Var for specified key from baggage.
<原文结束>

# <翻译开始>
// GetVar 从 baggage 中根据指定的键获取值，并返回一个指向该值的*gvar.Var指针。
# <翻译结束>

