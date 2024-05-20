
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31
# <翻译结束>


<原文开始>
// Package command provides console operations, like options/arguments reading.
<原文结束>

# <翻译开始>
// 包command提供控制台操作，如选项/参数读取。. md5:940e3926fff20c9a
# <翻译结束>


<原文开始>
// Init does custom initialization.
<原文结束>

# <翻译开始>
// Init 进行自定义初始化。. md5:08f8a2052942d9c8
# <翻译结束>


<原文开始>
// Parsing os.Args with default algorithm.
<原文结束>

# <翻译开始>
// 使用默认算法解析os.Args。. md5:460cde73efccc8f2
# <翻译结束>


<原文开始>
// ParseUsingDefaultAlgorithm parses arguments using default algorithm.
<原文结束>

# <翻译开始>
// 使用默认算法解析参数。. md5:d48f7f39a81379e1
# <翻译结束>


<原文开始>
// Example: gf gen -d -n 1
<原文结束>

# <翻译开始>
// 例子：gf gen -d -n 1
// 
// 这段Go语言的注释翻译成中文是：“这是一个示例命令行：使用gf（一个工具）生成代码，指定-d表示启用调试模式，-n 1表示生成1个示例。”. md5:a66be18aee4c44b9
# <翻译结束>


<原文开始>
// GetOpt returns the option value named `name`.
<原文结束>

# <翻译开始>
// GetOpt 返回名为 `name` 的选项值。. md5:5de4cb85c231ce6b
# <翻译结束>


<原文开始>
// GetOptAll returns all parsed options.
<原文结束>

# <翻译开始>
// GetOptAll 返回所有已解析的选项。. md5:6de4d266d8991786
# <翻译结束>


<原文开始>
// ContainsOpt checks whether option named `name` exist in the arguments.
<原文结束>

# <翻译开始>
// ContainsOpt 检查参数中是否存在名为 `name` 的选项。. md5:32ce4c1cf77651fb
# <翻译结束>


<原文开始>
// GetArg returns the argument at `index`.
<原文结束>

# <翻译开始>
// GetArg 返回索引为 `index` 的参数。. md5:822343a8734fe602
# <翻译结束>


<原文开始>
// GetArgAll returns all parsed arguments.
<原文结束>

# <翻译开始>
// GetArgAll 返回所有解析的参数。. md5:85cc0fd5995d4878
# <翻译结束>


<原文开始>
// GetOptWithEnv returns the command line argument of the specified `key`.
// If the argument does not exist, then it returns the environment variable with specified `key`.
// It returns the default value `def` if none of them exists.
//
// Fetching Rules:
// 1. Command line arguments are in lowercase format, eg: gf.package.variable;
// 2. Environment arguments are in uppercase format, eg: GF_PACKAGE_VARIABLE；
<原文结束>

# <翻译开始>
// GetOptWithEnv 返回指定 `key` 的命令行参数。
// 如果该参数不存在，则返回指定 `key` 的环境变量。
// 如果两者都不存在，它将返回默认值 `def`。
//
// 获取规则：
// 1. 命令行参数采用小写格式，例如：gf.package.variable；
// 2. 环境变量采用大写格式，例如：GF_PACKAGE_VARIABLE。
// md5:13bcb9c2795488a1
# <翻译结束>

