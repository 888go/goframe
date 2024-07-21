
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
// Run calls custom function in os.Args that bound to this command.
// It exits this process with exit code 1 if any error occurs.
<原文结束>

# <翻译开始>
// Run 调用与该命令绑定的自定义函数，根据os.Args执行。
// 如果发生任何错误，它将使进程退出并返回退出代码1。
// md5:f6512536eb3555fe
# <翻译结束>


<原文开始>
// RunWithValue calls custom function in os.Args that bound to this command with value output.
// It exits this process with exit code 1 if any error occurs.
<原文结束>

# <翻译开始>
// RunWithValue 调用与该命令绑定的 os.Args 中的自定义函数，传入值作为输出。如果发生任何错误，它将退出进程并返回退出码 1。
// md5:4d204c2503673c10
# <翻译结束>


<原文开始>
// RunWithError calls custom function in os.Args that bound to this command with error output.
<原文结束>

# <翻译开始>
// RunWithError 调用与该命令关联的 os.Args 中的自定义函数，同时输出错误信息。 md5:59f4632a1aab9342
# <翻译结束>


<原文开始>
// RunWithValueError calls custom function in os.Args that bound to this command with value and error output.
<原文结束>

# <翻译开始>
// RunWithValueError 使用os.Args中的值调用与此命令关联的自定义函数，并带有值和错误输出。 md5:007ad372fee78f96
# <翻译结束>


<原文开始>
// RunWithSpecificArgs calls custom function in specific args that bound to this command with value and error output.
<原文结束>

# <翻译开始>
// RunWithSpecificArgs 使用绑定到该命令的特定参数调用自定义函数，并将值和错误输出传递给它。 md5:48c98cbef4733851
# <翻译结束>


<原文开始>
// Exclude the root binary name.
<原文结束>

# <翻译开始>
// 排除根二进制文件名。 md5:74be69a4a70a25b8
# <翻译结束>


<原文开始>
// Find the matched command and run it.
<原文结束>

# <翻译开始>
// 找到匹配的命令并运行它。 md5:7cabb95b952de688
# <翻译结束>


<原文开始>
// Print error and help command if no command found.
<原文结束>

# <翻译开始>
// 如果未找到命令，则打印错误和帮助信息。 md5:e8829411cb2fb3df
# <翻译结束>


<原文开始>
// Check built-in help command.
<原文结束>

# <翻译开始>
// 检查内置的帮助命令。 md5:80aa5adefafed66d
# <翻译结束>


<原文开始>
// OpenTelemetry for command.
<原文结束>

# <翻译开始>
// 为命令提供OpenTelemetry。 md5:46407dd5b38f692f
# <翻译结束>


<原文开始>
// Reparse the original arguments for current command configuration.
<原文结束>

# <翻译开始>
// 为当前命令配置重新解析原始参数。 md5:6dfe6c6434a27ec5
# <翻译结束>


<原文开始>
// Registered command function calling.
<原文结束>

# <翻译开始>
// 注册的命令函数调用。 md5:9e5739c9c6b28f0f
# <翻译结束>


<原文开始>
// If no function defined in current command, it then prints help.
<原文结束>

# <翻译开始>
// 如果当前命令中没有定义函数，那么它会打印帮助信息。 md5:35f280d9901715f5
# <翻译结束>


<原文开始>
// reParse parses the original arguments using option configuration of current command.
<原文结束>

# <翻译开始>
// reParse使用当前命令的选项配置重新解析原始参数。 md5:c23561243bbefff3
# <翻译结束>


<原文开始>
// Retrieve option values from config component if it has "config" tag.
<原文结束>

# <翻译开始>
// 如果config组件有"config"标签，从其中获取选项值。 md5:25fb126ffe7890dc
# <翻译结束>


<原文开始>
// The command line has the high priority.
<原文结束>

# <翻译开始>
// 命令行具有较高优先级。 md5:8326234bd7de1eaa
# <翻译结束>


<原文开始>
// Merge the config value into parser.
<原文结束>

# <翻译开始>
// 将配置值合并到解析器中。 md5:82c508be2619b799
# <翻译结束>


<原文开始>
// searchCommand recursively searches the command according given arguments.
<原文结束>

# <翻译开始>
// searchCommand 递归地根据给定的参数搜索命令。 md5:5a28ecf7bd849fd7
# <翻译结束>


<原文开始>
		// Recursively searching the command.
		// String comparison case-sensitive.
<原文结束>

# <翻译开始>
		// 递归搜索命令。
		// 字符串比较区分大小写。
		// md5:801cc6b5c74b2a82
# <翻译结束>


<原文开始>
			// If this command needs argument,
			// it then gives all its left arguments to it using arg index marks.
			//
			// Note that the args here (using default args parsing) could be different with the args
			// that are parsed in command.
<原文结束>

# <翻译开始>
			// 如果此命令需要参数，
			// 则使用参数索引标记将其剩余的所有参数传递给它。
			//
			// 注意，这里使用的args（采用默认的参数解析方式）可能与在命令中解析到的args有所不同。
			// md5:6f65480aaaabf1f3
# <翻译结束>

