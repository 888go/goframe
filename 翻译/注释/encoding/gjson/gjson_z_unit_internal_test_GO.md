
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
		// fmt.Println(gregex.IsMatch(`^[\s\t\n\r]*[\w\-]+\s*:\s*".+"`, data))
		// fmt.Println(gregex.IsMatch(`^[\s\t\n\r]*[\w\-]+\s*:\s*\w+`, data))
		// fmt.Println(gregex.IsMatch(`[\s\t\n\r]+[\w\-]+\s*:\s*".+"`, data))
		// fmt.Println(gregex.IsMatch(`[\n\r]+[\w\-\s\t]+\s*:\s*\w+`, data))
		// fmt.Println(gregex.MatchString(`[\n\r]+[\w\-\s\t]+\s*:\s*\w+`, string(data)))
<原文结束>

# <翻译开始>
// 使用gregex函数检查数据是否匹配以空格、制表符、换行符开头，后跟一个或多个字母数字字符或减号，然后是空格，接着是冒号和双引号的字符串模式。
// fmt.Println(gregex.IsMatch(`^[\s\t\n\r]*[\w\-]+\s*:\s*".+"`, data))
// 
// 使用gregex函数检查数据是否匹配以空格、制表符、换行符开头，后跟一个或多个字母数字字符或减号，然后是空格，接着是冒号和一个或多个字母数字字符的模式。
// fmt.Println(gregex.IsMatch(`^[\s\t\n\r]*[\w\-]+\s*:\s*\w+`, data))
// 
// 使用gregex函数检查数据是否匹配任何数量的空格、制表符、换行符，后跟一个或多个字母数字字符或减号，然后是空格，接着是冒号和双引号的字符串模式。
// fmt.Println(gregex.IsMatch(`[\s\t\n\r]+[\w\-]+\s*:\s*".+"`, data))
// 
// 使用gregex函数检查数据是否匹配任何数量的换行符或回车符，后跟一个或多个字母数字字符、减号、空格、制表符，然后是空格，接着是冒号和一个或多个字母数字字符的模式。
// fmt.Println(gregex.IsMatch(`[\n\r]+[\w\-\s\t]+\s*:\s*\w+`, data))
// 
// 将数据转换为字符串并使用gregex函数检查是否匹配上述的换行符、字母数字字符、减号、空格、制表符模式，然后是冒号和一个或多个字母数字字符的模式。
// fmt.Println(gregex.MatchString(`[\n\r]+[\w\-\s\t]+\s*:\s*\w+`, string(data)))
// md5:dfba1419755c382a
# <翻译结束>

