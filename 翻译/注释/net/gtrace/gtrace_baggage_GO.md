
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
// Baggage holds the data through all tracing spans.
<原文结束>

# <翻译开始>
// Baggage在所有跟踪span中持有数据。. md5:0ad27152ec042f81
# <翻译结束>


<原文开始>
// NewBaggage creates and returns a new Baggage object from given tracing context.
<原文结束>

# <翻译开始>
// NewBaggage 从给定的追踪上下文中创建并返回一个新的Baggage对象。. md5:6c3e8093bd06a60a
# <翻译结束>


<原文开始>
// Ctx returns the context that Baggage holds.
<原文结束>

# <翻译开始>
// Ctx 返回Baggage持有的上下文。. md5:37268f528c617799
# <翻译结束>


<原文开始>
// SetValue is a convenient function for adding one key-value pair to baggage.
// Note that it uses attribute.Any to set the key-value pair.
<原文结束>

# <翻译开始>
// SetValue 是一个方便的函数，用于向 baggage 中添加一个键值对。
// 注意，它使用 attribute.Any 设置键值对。
// md5:830faae9a81721ce
# <翻译结束>


<原文开始>
// SetMap is a convenient function for adding map key-value pairs to baggage.
// Note that it uses attribute.Any to set the key-value pair.
<原文结束>

# <翻译开始>
// SetMap 是一个方便的函数，用于将映射键值对添加到行李中。
// 注意，它使用 attribute.Any 设置键值对。
// md5:a18951801562457c
# <翻译结束>


<原文开始>
// GetMap retrieves and returns the baggage values as map.
<原文结束>

# <翻译开始>
// GetMap 获取并以映射形式返回baggage值。. md5:d6024d765655a29e
# <翻译结束>


<原文开始>
// GetVar retrieves value and returns a *gvar.Var for specified key from baggage.
<原文结束>

# <翻译开始>
// GetVar 从行李中获取指定键的值，并返回一个 *gvar.Var。. md5:6cda7fcfb8ff1c6e
# <翻译结束>

