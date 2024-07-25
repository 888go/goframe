
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Adapter is the interface for configuration retrieving.
<原文结束>

# <翻译开始>
// Adapter是用于获取配置的接口。 md5:5c3d613bea87d056
# <翻译结束>


<原文开始>
	// Available checks and returns the backend configuration service is available.
	// The optional parameter `resource` specifies certain configuration resource.
	//
	// Note that this function does not return error as it just does simply check for
	// backend configuration service.
<原文结束>

# <翻译开始>
	// 可用性检查并返回后端配置服务是否可用。
	// 可选参数 `resource` 指定特定的配置资源。
	//
	// 请注意，此函数不会返回错误，因为它只是简单地检查后端配置服务。 md5:8c240c72c0849cd7
# <翻译结束>


<原文开始>
	// Get retrieves and returns value by specified `pattern` in current resource.
	// Pattern like:
	// "x.y.z" for map item.
	// "x.0.y" for slice item.
<原文结束>

# <翻译开始>
	// Get 通过在当前资源中指定的`pattern`获取并返回值。
	// 模式示例：
	// 对于映射项，使用 "x.y.z"。
	// 对于切片项，使用 "x.0.y"。 md5:821429a92b84150c
# <翻译结束>


<原文开始>
	// Data retrieves and returns all configuration data in current resource as map.
	// Note that this function may lead lots of memory usage if configuration data is too large,
	// you can implement this function if necessary.
<原文结束>

# <翻译开始>
	// Data 获取并返回当前资源中的所有配置数据作为映射。
	// 请注意，如果配置数据过大，此函数可能导致大量内存使用。
	// 如果需要，你可以自行实现这个函数。 md5:7eaedd1a7f099a23
# <翻译结束>

