
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
// Command holds the info about an argument that can handle custom logic.
<原文结束>

# <翻译开始>
// Command 包含有关可以处理自定义逻辑的参数的信息。. md5:b0e0f23cc6e868c5
# <翻译结束>


<原文开始>
// Command name(case-sensitive).
<原文结束>

# <翻译开始>
// 命令名称（大小写敏感）。. md5:44e7c13c9c0eced2
# <翻译结束>


<原文开始>
// A brief line description about its usage, eg: gf build main.go [OPTION]
<原文结束>

# <翻译开始>
// 关于其用途的简短描述，例如：gf build main.go [选项]. md5:e2660484a0edfee8
# <翻译结束>


<原文开始>
// A brief info that describes what this command will do.
<原文结束>

# <翻译开始>
// 一个简短的描述，说明此命令将执行的操作。. md5:4a0304d2ac452238
# <翻译结束>


<原文开始>
// A detailed description.
<原文结束>

# <翻译开始>
// 一个详细的描述。. md5:b83b3d2318b54bce
# <翻译结束>


<原文开始>
// Argument array, configuring how this command act.
<原文结束>

# <翻译开始>
// 参数数组，配置此命令的行为。. md5:9c82b0f6e377e648
# <翻译结束>


<原文开始>
// Custom function with output parameters that can interact with command caller.
<原文结束>

# <翻译开始>
// 自定义函数，带有输出参数，能够与命令调用者进行交互。. md5:586037addaa736f6
# <翻译结束>


<原文开始>
// Additional info about this command, which will be appended to the end of help info.
<原文结束>

# <翻译开始>
// 关于此命令的附加信息，将追加到帮助信息的末尾。. md5:328b9830bf970895
# <翻译结束>


<原文开始>
// Strict parsing options, which means it returns error if invalid option given.
<原文结束>

# <翻译开始>
// 严格的解析选项，这意味着如果提供无效的选项，它会返回错误。. md5:5e8eec207aef7c7a
# <翻译结束>


<原文开始>
// CaseSensitive parsing options, which means it parses input options in case-sensitive way.
<原文结束>

# <翻译开始>
// 区分大小写的解析选项，这意味着它以区分大小写的方式解析输入选项。. md5:b18eddb5f60c7176
# <翻译结束>


<原文开始>
// Config node name, which also retrieves the values from config component along with command line.
<原文结束>

# <翻译开始>
// 配置节点名称，它也会从配置组件中获取值，以及从命令行中获取。. md5:0f67ea7288e8e541
# <翻译结束>


<原文开始>
// Parent command for internal usage.
<原文结束>

# <翻译开始>
// 用于内部使用的父命令。. md5:6572369b29bb2e3e
# <翻译结束>


<原文开始>
// Sub commands of this command.
<原文结束>

# <翻译开始>
// 该命令的子命令。. md5:579fb8699f4ff8e3
# <翻译结束>


<原文开始>
// Function is a custom command callback function that is bound to a certain argument.
<原文结束>

# <翻译开始>
// Function 是一个绑定到特定参数的自定义命令回调函数。. md5:74f820cae660a1b5
# <翻译结束>


<原文开始>
// FuncWithValue is similar like Func but with output parameters that can interact with command caller.
<原文结束>

# <翻译开始>
// FuncWithValue 类似于 Func，但它带有输出参数，这些参数可以与命令调用者进行交互。. md5:e8459756fad8cbb9
# <翻译结束>


<原文开始>
// Argument is the command value that are used by certain command.
<原文结束>

# <翻译开始>
// Argument 是某些命令使用的命令值。. md5:e5c110dcf519025a
# <翻译结束>


<原文开始>
// Brief info about this Option, which is used in help info.
<原文结束>

# <翻译开始>
// 这个选项的简要信息，用于帮助信息中。. md5:b913553040a0d889
# <翻译结束>


<原文开始>
// IsArg marks this argument taking value from command line argument instead of option.
<原文结束>

# <翻译开始>
// IsArg 标记这个参数从命令行参数而不是选项中获取值。. md5:24e6cc6cb658557a
# <翻译结束>


<原文开始>
// Whether this Option having or having no value bound to it.
<原文结束>

# <翻译开始>
// 此选项是否有值与之绑定。. md5:bc1b6ee078e2683c
# <翻译结束>


<原文开始>
// defaultHelpOption is the default help option that will be automatically added to each command.
<原文结束>

# <翻译开始>
// defaultHelpOption 是默认的帮助选项，将会自动添加到每个命令中。. md5:3593428e8c7dfe0a
# <翻译结束>


<原文开始>
// CommandFromCtx retrieves and returns Command from context.
<原文结束>

# <翻译开始>
// CommandFromCtx从上下文检索并返回Command。. md5:81a6b36fc029401b
# <翻译结束>


<原文开始>
// AddCommand adds one or more sub-commands to current command.
<原文结束>

# <翻译开始>
// AddCommand向当前命令添加一个或多个子命令。. md5:f1582e4eafa78dd7
# <翻译结束>


<原文开始>
// doAddCommand adds one sub-command to current command.
<原文结束>

# <翻译开始>
// doAddCommand 向当前命令添加一个子命令。. md5:bd1d8d447805aafd
# <翻译结束>


<原文开始>
// Add the given command to its sub-commands array.
<原文结束>

# <翻译开始>
// 将给定的命令添加到其子命令数组中。. md5:ddd450893c5e1fcc
# <翻译结束>


<原文开始>
// AddObject adds one or more sub-commands to current command using struct object.
<原文结束>

# <翻译开始>
// AddObject 通过struct对象向当前命令添加一个或多个子命令。. md5:8de76f64f667f83d
# <翻译结束>

