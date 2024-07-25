
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
// octReg is the regular expression object for checks octal string.
<原文结束>

# <翻译开始>
	// octReg 是用于检查八进制字符串的正则表达式对象。 md5:5c5c93db5da71e18
# <翻译结束>


<原文开始>
// Chr return the ascii string of a number(0-255).
//
// Example:
// Chr(65) -> "A"
<原文结束>

# <翻译开始>
// Chr 函数返回一个数字（0-255）的ASCII字符串。
//
// 示例：
// Chr(65) -> "A"
// md5:1eeda35a229d907f
# <翻译结束>


<原文开始>
// Ord converts the first byte of a string to a value between 0 and 255.
//
// Example:
// Chr("A") -> 65
<原文结束>

# <翻译开始>
// Ord 将字符串的第一个字节转换为0到255之间的值。
//
// 示例：
// Ord("A") -> 65
// md5:4b57c924e8be0a49
# <翻译结束>


<原文开始>
// Reverse returns a string which is the reverse of `str`.
//
// Example:
// Reverse("123456") -> "654321"
<原文结束>

# <翻译开始>
// Reverse 函数返回一个字符串，它是 `str` 的反向字符串。
//
// 示例：
// Reverse("123456") -> "654321"
// md5:7106270467ce887e
# <翻译结束>


<原文开始>
// NumberFormat formats a number with grouped thousands.
// Parameter `decimals`: Sets the number of decimal points.
// Parameter `decPoint`: Sets the separator for the decimal point.
// Parameter `thousandsSep`: Sets the thousands' separator.
// See http://php.net/manual/en/function.number-format.php.
//
// Example:
// NumberFormat(1234.56, 2, ".", "")  -> 1234,56
// NumberFormat(1234.56, 2, ",", " ") -> 1 234,56
<原文结束>

# <翻译开始>
// NumberFormat 用千位分隔符格式化数字。
// 参数 `decimals`：设置小数点后的位数。
// 参数 `decPoint`：设置小数点的分隔符。
// 参数 `thousandsSep`：设置千位分隔符。
// 参考：http://php.net/manual/en/function.number-format.php。
// 
// 示例：
// NumberFormat(1234.56, 2, ".", "") -> 1234,56
// NumberFormat(1234.56, 2, ",", " ") -> 1 234,56
// md5:c4f419bbc874acfc
# <翻译结束>


<原文开始>
// Shuffle randomly shuffles a string.
// It considers parameter `str` as unicode string.
//
// Example:
// Shuffle("123456") -> "325164"
// Shuffle("123456") -> "231546"
// ...
<原文结束>

# <翻译开始>
// `Shuffle` 随机打乱一个字符串。
// 它将参数 `str` 视为 Unicode 字符串。
// 
// 示例：
// Shuffle("123456") -> "325164"
// Shuffle("123456") -> "231546"
// ...
// md5:2e7f0ae98e6b5210
# <翻译结束>


<原文开始>
// HideStr replaces part of the string `str` to `hide` by `percentage` from the `middle`.
// It considers parameter `str` as unicode string.
<原文结束>

# <翻译开始>
// HideStr 函数将字符串 `str` 的从中间开始按 `percentage` 比例部分内容替换为 `hide`。
// 此函数将参数 `str` 视为Unicode字符串处理。
// md5:f9986962939bb788
# <翻译结束>


<原文开始>
// Nl2Br inserts HTML line breaks(`br`|<br />) before all newlines in a string:
// \n\r, \r\n, \r, \n.
// It considers parameter `str` as unicode string.
<原文结束>

# <翻译开始>
// Nl2Br 在字符串中的所有换行符(\n\r, \r\n, \r, \n)前插入HTML换行标签(`br`|<br />)。
// 它将参数`str`视为Unicode字符串。
// md5:6cad5f70848065d0
# <翻译结束>


<原文开始>
// WordWrap wraps a string to a given number of characters.
// This function supports cut parameters of both english and chinese punctuations.
// TODO: Enable custom cut parameter, see http://php.net/manual/en/function.wordwrap.php.
<原文结束>

# <翻译开始>
// WordWrap 将一个字符串按照给定的字符数进行换行。
// 这个函数支持英文和中文标点符号的截断参数。
// TODO: 开启自定义截断参数，参考 http://php.net/manual/en/function.wordwrap.php。
// md5:389c5474efb0a8e8
# <翻译结束>

