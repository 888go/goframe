
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
// Adapter is the interface for configuration retrieving.
<原文结束>

# <翻译开始>
// Adapter 是用于获取配置的接口。
# <翻译结束>


<原文开始>
	// Available checks and returns the backend configuration service is available.
	// The optional parameter `resource` specifies certain configuration resource.
	//
	// Note that this function does not return error as it just does simply check for
	// backend configuration service.
<原文结束>

# <翻译开始>
// Available 检查并返回配置服务是否可用。
// 可选参数 `resource` 用于指定特定的配置资源。
//
// 注意，此函数不会返回错误，因为它只是简单地检查后端配置服务。
# <翻译结束>


<原文开始>
	// Get retrieves and returns value by specified `pattern` in current resource.
	// Pattern like:
	// "x.y.z" for map item.
	// "x.0.y" for slice item.
<原文结束>

# <翻译开始>
// Get 通过指定的 `pattern` 从当前资源中获取并返回值。
// 模式示例：
// "x.y.z" 用于获取映射项。
// "x.0.y" 用于获取切片项。
# <翻译结束>


<原文开始>
	// Data retrieves and returns all configuration data in current resource as map.
	// Note that this function may lead lots of memory usage if configuration data is too large,
	// you can implement this function if necessary.
<原文结束>

# <翻译开始>
// Data 函数从当前资源中获取并返回所有的配置数据，以 map 的形式。
// 注意，如果配置数据过大，此函数可能会导致大量内存使用，
// 如有必要，您可以自行实现这个函数。
# <翻译结束>

