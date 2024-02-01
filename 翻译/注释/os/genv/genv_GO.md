
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
// Package genv provides operations for environment variables of system.
<原文结束>

# <翻译开始>
// Package genv 提供了对系统环境变量的操作。
# <翻译结束>


<原文开始>
// All returns a copy of strings representing the environment,
// in the form "key=value".
<原文结束>

# <翻译开始>
// All 函数返回一个表示环境变量的字符串副本，其形式为 "key=value"。
# <翻译结束>


<原文开始>
// Map returns a copy of strings representing the environment as a map.
<原文结束>

# <翻译开始>
// Map 返回一个字符串环境表示的映射副本。
# <翻译结束>


<原文开始>
// Get creates and returns a Var with the value of the environment variable
// named by the `key`. It uses the given `def` if the variable does not exist
// in the environment.
<原文结束>

# <翻译开始>
// Get 函数创建并返回一个 Var，其值为环境变量中名为 `key` 的变量的值。
// 如果该变量在环境中不存在，则使用给定的 `def` 作为默认值。
# <翻译结束>


<原文开始>
// Set sets the value of the environment variable named by the `key`.
// It returns an error, if any.
<原文结束>

# <翻译开始>
// Set 函数用于设置名为 `key` 的环境变量的值。
// 如果出现错误，该函数会返回一个错误。
# <翻译结束>


<原文开始>
// SetMap sets the environment variables using map.
<原文结束>

# <翻译开始>
// SetMap 通过 map 设置环境变量。
# <翻译结束>


<原文开始>
// Contains checks whether the environment variable named `key` exists.
<原文结束>

# <翻译开始>
// Contains 检查名为 `key` 的环境变量是否存在。
# <翻译结束>


<原文开始>
// Remove deletes one or more environment variables.
<原文结束>

# <翻译开始>
// Remove 删除一个或多个环境变量。
# <翻译结束>


<原文开始>
// Build builds a map to an environment variable slice.
<原文结束>

# <翻译开始>
// Build 函数用于构建一个映射到环境变量切片的映射。
# <翻译结束>


<原文开始>
// MapFromEnv converts environment variables from slice to map.
<原文结束>

# <翻译开始>
// MapFromEnv 将环境变量从切片转换为映射（map）。
# <翻译结束>


<原文开始>
// MapToEnv converts environment variables from map to slice.
<原文结束>

# <翻译开始>
// MapToEnv 将环境变量从映射转换为切片。
# <翻译结束>


<原文开始>
// Filter filters repeated items from given environment variables.
<原文结束>

# <翻译开始>
// Filter 从给定的环境变量中过滤掉重复项。
# <翻译结束>

