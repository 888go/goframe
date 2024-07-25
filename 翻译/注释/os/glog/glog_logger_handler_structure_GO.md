
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
// 从encoding/json/tables.go复制。
//
// safeSet用于存储一个布尔值，表示具有给定数组位置的ASCII字符可以在JSON字符串中表示，而无需进一步转义。
//
// 除了ASCII控制字符（0-31）、双引号（"）和反斜杠字符（\）之外，所有值都为true。
// md5:2df5305c3a107923
# <翻译结束>


<原文开始>
// HandlerStructure is a handler for output logging content as a structured string.
<原文结束>

# <翻译开始>
// HandlerStructure 是一个处理器，用于将输出的日志内容以结构化字符串的形式处理。 md5:392f74b46dcdd1eb
# <翻译结束>


<原文开始>
// If the values cannot be the pair, move the first one to content.
<原文结束>

# <翻译开始>
	// 如果这些值不能组成一对，将第一个移到content中。 md5:2bc1ae2ae5605225
# <翻译结束>


<原文开始>
			// Quote anything except a backslash that would need quoting in a
			// JSON string, as well as space and '='
<原文结束>

# <翻译开始>
			// 在JSON字符串中，除了需要转义的反斜杠、空格和'='之外，对任何内容进行引号包裹
			// md5:0202f0293260c21e
# <翻译结束>

