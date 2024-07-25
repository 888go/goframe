
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
// Package genv provides operations for environment variables of system.
<原文结束>

# <翻译开始>
// 包genv提供了对系统环境变量的操作。 md5:9605f9d2a2186f5b
# <翻译结束>


<原文开始>
// All returns a copy of strings representing the environment,
// in the form "key=value".
<原文结束>

# <翻译开始>
// All返回一个字符串切片的副本，表示环境，形式为"key=value"。 md5:723df5605f199f2b
# <翻译结束>


<原文开始>
// Map returns a copy of strings representing the environment as a map.
<原文结束>

# <翻译开始>
// Map 返回一个副本，该副本将环境表示为映射（map）形式的字符串。 md5:9477b6d266100b3d
# <翻译结束>


<原文开始>
// Get creates and returns a Var with the value of the environment variable
// named by the `key`. It uses the given `def` if the variable does not exist
// in the environment.
<原文结束>

# <翻译开始>
// Get 根据给定的`key`创建并返回一个具有环境变量值的Var。如果环境中不存在该变量，则使用给定的`def`作为默认值。 md5:1c5c61ffd2aa5106
# <翻译结束>


<原文开始>
// Set sets the value of the environment variable named by the `key`.
// It returns an error, if any.
<原文结束>

# <翻译开始>
// Set 设置环境变量的值，该变量由 `key` 指定。如果发生任何错误，它将返回一个错误。 md5:3d9ca695de9bb4ad
# <翻译结束>


<原文开始>
// SetMap sets the environment variables using map.
<原文结束>

# <翻译开始>
// SetMap 使用映射设置环境变量。 md5:78d0cfffe3bc8311
# <翻译结束>


<原文开始>
// Contains checks whether the environment variable named `key` exists.
<原文结束>

# <翻译开始>
// Contains 检查名为 `key` 的环境变量是否存在。 md5:76124e3be6d217ff
# <翻译结束>


<原文开始>
// Remove deletes one or more environment variables.
<原文结束>

# <翻译开始>
// Remove 删除一个或多个环境变量。 md5:546a01a7df799055
# <翻译结束>


<原文开始>
// GetWithCmd returns the environment value specified `key`.
// If the environment value does not exist, then it retrieves and returns the value from command line options.
// It returns the default value `def` if none of them exists.
//
// Fetching Rules:
// 1. Environment arguments are in uppercase format, eg: GF_<package name>_<variable name>；
// 2. Command line arguments are in lowercase format, eg: gf.<package name>.<variable name>;
<原文结束>

# <翻译开始>
// GetWithCmd 返回指定的环境变量值 `key`。
// 如果环境变量不存在，它将从命令行选项中检索并返回值。如果两者都不存在，它将返回默认值 `def`。
//
// 获取规则：
// 1. 环境变量参数使用大写格式，例如：GF_<包名>_<变量名>；
// 2. 命令行参数使用小写格式，例如：gf.<包名>.<变量名>； md5:1bba2e845d6ee0d6
# <翻译结束>


<原文开始>
// Build builds a map to an environment variable slice.
<原文结束>

# <翻译开始>
// Build 构建一个映射到环境变量切片的map。 md5:f58dc9490f9468a7
# <翻译结束>


<原文开始>
// MapFromEnv converts environment variables from slice to map.
<原文结束>

# <翻译开始>
// MapFromEnv 将环境变量从切片转换为映射。 md5:1c7b8b3cbc6a6d0d
# <翻译结束>


<原文开始>
// MapToEnv converts environment variables from map to slice.
<原文结束>

# <翻译开始>
// MapToEnv 将环境变量从映射转换为切片。 md5:3cef9db0baccea9f
# <翻译结束>


<原文开始>
// Filter filters repeated items from given environment variables.
<原文结束>

# <翻译开始>
// Filter 从给定的环境变量中过滤重复项。 md5:7b495d60bfff573e
# <翻译结束>

