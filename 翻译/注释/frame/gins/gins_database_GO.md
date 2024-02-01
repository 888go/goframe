
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
// Database returns an instance of database ORM object with specified configuration group name.
// Note that it panics if any error occurs duration instance creating.
<原文结束>

# <翻译开始>
// Database 返回指定配置组名称的数据库 ORM 对象实例。
// 注意，如果实例创建过程中发生任何错误，它会触发 panic。
# <翻译结束>


<原文开始>
// It ignores returned error to avoid file no found error while it's not necessary.
<原文结束>

# <翻译开始>
// 它忽略返回的错误，以防止在不必要的时候出现文件未找到的错误。
# <翻译结束>


<原文开始>
// It firstly searches the configuration of the instance name.
<原文结束>

# <翻译开始>
// 首先，它会搜索实例名称的配置。
# <翻译结束>


<原文开始>
// No configuration found, it formats and panics error.
<原文结束>

# <翻译开始>
// 未找到配置，它将格式化并引发错误。
# <翻译结束>


<原文开始>
// File configuration object checks.
<原文结束>

# <翻译开始>
// 文件配置对象检查。
# <翻译结束>


<原文开始>
// Panic if nothing found in Config object or in gdb configuration.
<原文结束>

# <翻译开始>
// 如果在Config对象或gdb配置中未找到任何内容，则引发panic。
# <翻译结束>


<原文开始>
// Parse `m` as map-slice and adds it to global configurations for package gdb.
<原文结束>

# <翻译开始>
// 将`m`解析为映射切片并将其添加到gdb包的全局配置中。
# <翻译结束>


<原文开始>
		// Parse `m` as a single node configuration,
		// which is the default group configuration.
<原文结束>

# <翻译开始>
// 将`m`解析为单节点配置，
// 这是默认的组配置。
# <翻译结束>


<原文开始>
// Create a new ORM object with given configurations.
<原文结束>

# <翻译开始>
// 使用给定的配置创建一个新的ORM对象。
# <翻译结束>







<原文开始>
// If panics, often because it does not find its configuration for given group.
<原文结束>

# <翻译开始>
// 如果出现 panic，通常是因为在给定的组中没有找到其配置。
# <翻译结束>


<原文开始>
// Find possible `Link` configuration content.
<原文结束>

# <翻译开始>
// 查找可能的`Link`配置内容。
# <翻译结束>


<原文开始>
// Initialize logger for ORM.
<原文结束>

# <翻译开始>
// 初始化ORM的logger（日志器）。
# <翻译结束>

