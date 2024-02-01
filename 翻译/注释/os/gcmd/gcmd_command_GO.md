
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//
# <翻译结束>


<原文开始>
// Command holds the info about an argument that can handle custom logic.
<原文结束>

# <翻译开始>
// Command 结构体保存了关于一个可以处理自定义逻辑的参数的信息。
# <翻译结束>


<原文开始>
// Command name(case-sensitive).
<原文结束>

# <翻译开始>
// 命令名称（区分大小写）。
# <翻译结束>


<原文开始>
// A brief line description about its usage, eg: gf build main.go [OPTION]
<原文结束>

# <翻译开始>
// 一行简短的描述，关于其使用方式，例如：gf build main.go [选项]
# <翻译结束>


<原文开始>
// A brief info that describes what this command will do.
<原文结束>

# <翻译开始>
// 这个命令将要执行的简短描述
# <翻译结束>







<原文开始>
// Argument array, configuring how this command act.
<原文结束>

# <翻译开始>
// 参数数组，用于配置此命令的行为。
# <翻译结束>


<原文开始>
// Custom function with output parameters that can interact with command caller.
<原文结束>

# <翻译开始>
// 自定义函数，带有输出参数，可以与命令调用者进行交互。
# <翻译结束>







<原文开始>
// Additional info about this command, which will be appended to the end of help info.
<原文结束>

# <翻译开始>
// 这个命令的附加信息，将会被添加到帮助信息的末尾。
# <翻译结束>


<原文开始>
// Strict parsing options, which means it returns error if invalid option given.
<原文结束>

# <翻译开始>
// 严格解析选项，这意味着如果给出无效选项将返回错误。
# <翻译结束>


<原文开始>
// CaseSensitive parsing options, which means it parses input options in case-sensitive way.
<原文结束>

# <翻译开始>
// CaseSensitive 解析选项，表示以区分大小写的方式解析输入选项。
# <翻译结束>


<原文开始>
// Config node name, which also retrieves the values from config component along with command line.
<原文结束>

# <翻译开始>
// 配置节点名称，该名称同时从配置组件和命令行中获取值。
# <翻译结束>


<原文开始>
// Parent command for internal usage.
<原文结束>

# <翻译开始>
// 用于内部使用的父命令。
# <翻译结束>


<原文开始>
// Sub commands of this command.
<原文结束>

# <翻译开始>
// 此命令的子命令。
# <翻译结束>


<原文开始>
// Function is a custom command callback function that is bound to a certain argument.
<原文结束>

# <翻译开始>
// Function 是一个自定义命令回调函数，它绑定到某个特定参数。
# <翻译结束>


<原文开始>
// FuncWithValue is similar like Func but with output parameters that can interact with command caller.
<原文结束>

# <翻译开始>
// FuncWithValue 类似于 Func，但是带有输出参数，可以与命令调用者进行交互。
# <翻译结束>


<原文开始>
// Argument is the command value that are used by certain command.
<原文结束>

# <翻译开始>
// Argument 是某些命令所使用的命令值。
# <翻译结束>


<原文开始>
// Brief info about this Option, which is used in help info.
<原文结束>

# <翻译开始>
// 该Option的简要信息，用于帮助信息中。
# <翻译结束>


<原文开始>
// IsArg marks this argument taking value from command line argument instead of option.
<原文结束>

# <翻译开始>
// IsArg 标记这个参数从命令行参数而非选项中获取值。
# <翻译结束>


<原文开始>
// Whether this Option having or having no value bound to it.
<原文结束>

# <翻译开始>
// 是否此Option已绑定或未绑定值。
# <翻译结束>


<原文开始>
// defaultHelpOption is the default help option that will be automatically added to each command.
<原文结束>

# <翻译开始>
// defaultHelpOption 是默认的帮助选项，它将自动添加到每个命令中。
# <翻译结束>


<原文开始>
// CommandFromCtx retrieves and returns Command from context.
<原文结束>

# <翻译开始>
// CommandFromCtx 从上下文中检索并返回 Command。
# <翻译结束>


<原文开始>
// AddCommand adds one or more sub-commands to current command.
<原文结束>

# <翻译开始>
// AddCommand 向当前命令添加一个或多个子命令。
# <翻译结束>


<原文开始>
// doAddCommand adds one sub-command to current command.
<原文结束>

# <翻译开始>
// doAddCommand 向当前命令添加一个子命令。
# <翻译结束>


<原文开始>
// Add the given command to its sub-commands array.
<原文结束>

# <翻译开始>
// 将给定的命令添加到其子命令数组中。
# <翻译结束>


<原文开始>
// AddObject adds one or more sub-commands to current command using struct object.
<原文结束>

# <翻译开始>
// AddObject 通过结构体对象向当前命令添加一个或多个子命令。
# <翻译结束>

















<原文开始>
// A detailed description.
<原文结束>

# <翻译开始>
// 一段详细描述
# <翻译结束>


<原文开始>
// Custom help function
<原文结束>

# <翻译开始>
// 自定义帮助函数
# <翻译结束>

