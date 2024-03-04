
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
// Copied from encoding/json/tables.go.
//
// safeSet holds the value true if the ASCII character with the given array
// position can be represented inside a JSON string without any further
// escaping.
//
// All values are true except for the ASCII control characters (0-31), the
// double quote ("), and the backslash character ("\").
<原文结束>

# <翻译开始>
// 从encoding/json/tables.go复制而来。
//
// safeSet 记录了如果ASCII字符在给定数组位置的值能在JSON字符串中无须额外转义即可表示，则其值为真（true）。
//
// 除了ASCII控制字符（0-31）、双引号（"）和反斜杠字符（\）之外，其余所有字符的值都为真。
# <翻译结束>


<原文开始>
// HandlerStructure is a handler for output logging content as a structured string.
<原文结束>

# <翻译开始>
// HandlerStructure 是一个处理器，用于将输出的日志内容以结构化字符串形式记录。
# <翻译结束>


<原文开始>
// If the values cannot be the pair, move the first one to content.
<原文结束>

# <翻译开始>
// 如果这些值不能构成一对，则将第一个值移动到content中。
# <翻译结束>


<原文开始>
			// Quote anything except a backslash that would need quoting in a
			// JSON string, as well as space and '='
<原文结束>

# <翻译开始>
// 将JSON字符串中需要转义的除反斜杠以外的任何字符，以及空格和'='进行引用
# <翻译结束>

