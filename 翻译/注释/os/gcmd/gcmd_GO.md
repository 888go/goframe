
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
// Package gcmd provides console operations, like options/arguments reading and command running.
<原文结束>

# <翻译开始>
// 包gcmd提供控制台操作，如读取选项/参数和运行命令。. md5:bb72337a704c599f
# <翻译结束>


<原文开始>
// Init does custom initialization.
<原文结束>

# <翻译开始>
// Init 进行自定义初始化。. md5:08f8a2052942d9c8
# <翻译结束>


<原文开始>
// GetOpt returns the option value named `name` as gvar.Var.
<原文结束>

# <翻译开始>
// GetOpt 作为gvar.Var返回名为`name`的选项值。. md5:1859b868ee779be0
# <翻译结束>


<原文开始>
// GetOptAll returns all parsed options.
<原文结束>

# <翻译开始>
// GetOptAll 返回所有已解析的选项。. md5:6de4d266d8991786
# <翻译结束>


<原文开始>
// GetArg returns the argument at `index` as gvar.Var.
<原文结束>

# <翻译开始>
// GetArg 作为gvar.Var返回索引为`index`的参数。. md5:12ea2f8d74c6370d
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
// 1. Command line arguments are in lowercase format, eg: gf.`package name`.<variable name>;
// 2. Environment arguments are in uppercase format, eg: GF_`package name`_<variable name>；
<原文结束>

# <翻译开始>
// GetOptWithEnv 返回指定 `key` 的命令行参数。
// 如果该参数不存在，则返回指定 `key` 的环境变量。
// 如果两者都不存在，它将返回默认值 `def`。
//
// 获取规则：
// 1. 命令行参数采用小写格式，例如：gf.`包名`.<变量名>;
// 2. 环境变量采用大写格式，例如：GF_`包名`_<变量名>。
// md5:e3d5c0c773430740
# <翻译结束>


<原文开始>
// BuildOptions builds the options as string.
<原文结束>

# <翻译开始>
// BuildOptions 将选项构建为字符串。. md5:c722b017f3a50346
# <翻译结束>

