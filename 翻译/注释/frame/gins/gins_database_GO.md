
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
// Database returns an instance of database ORM object with specified configuration group name.
// Note that it panics if any error occurs duration instance creating.
<原文结束>

# <翻译开始>
// Database 返回一个根据指定配置组名实例化的数据库ORM对象。
// 注意，如果在实例创建过程中发生任何错误，它将会直接 panic。
// md5:c8c0e8142b2f24af
# <翻译结束>


<原文开始>
// It ignores returned error to avoid file no found error while it's not necessary.
<原文结束>

# <翻译开始>
		// 忽略返回的错误，以避免在不需要时出现“文件未找到”错误。 md5:47e693921809cd8c
# <翻译结束>


<原文开始>
// It firstly searches the configuration of the instance name.
<原文结束>

# <翻译开始>
		// 它首先搜索实例名称的配置。 md5:0b825658b318a2f7
# <翻译结束>


<原文开始>
// No configuration found, it formats and panics error.
<原文结束>

# <翻译开始>
		// 没有找到配置，它会格式化并引发 panic 错误。 md5:8716646cceaee999
# <翻译结束>


<原文开始>
// File configuration object checks.
<原文结束>

# <翻译开始>
			// 文件配置对象检查。 md5:fdae1c62b2593d55
# <翻译结束>


<原文开始>
// Panic if nothing found in Config object or in gdb configuration.
<原文结束>

# <翻译开始>
			// 如果在Config对象或gdb配置中找不到任何内容，则引发恐慌。 md5:2c3aa642bbae15da
# <翻译结束>


<原文开始>
// Parse `m` as map-slice and adds it to global configurations for package gdb.
<原文结束>

# <翻译开始>
		// 将 `m` 解析为映射切片，并将其添加到gdb包的全局配置中。 md5:8970d506724c2880
# <翻译结束>


<原文开始>
		// Parse `m` as a single node configuration,
		// which is the default group configuration.
<原文结束>

# <翻译开始>
		// 将 `m` 解析为单个节点配置，
		// 这是默认的组配置。
		// md5:8f62d1ad0b43783e
# <翻译结束>


<原文开始>
// Create a new ORM object with given configurations.
<原文结束>

# <翻译开始>
		// 使用给定的配置创建一个新的ORM对象。 md5:8114aaedeed4c350
# <翻译结束>


<原文开始>
// Initialize logger for ORM.
<原文结束>

# <翻译开始>
			// 初始化ORM的日志记录器。 md5:5fbf0eb7ce9402d0
# <翻译结束>


<原文开始>
// If panics, often because it does not find its configuration for given group.
<原文结束>

# <翻译开始>
			// 如果出现恐慌，通常是由于它没有找到给定组的配置。 md5:461786d647ecc99d
# <翻译结束>


<原文开始>
// Find possible `Link` configuration content.
<原文结束>

# <翻译开始>
	// 查找可能的 `Link` 配置内容。 md5:c3acedff678206f1
# <翻译结束>

