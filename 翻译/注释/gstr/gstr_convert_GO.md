
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
// octReg is the regular expression object for checks octal string.
<原文结束>

# <翻译开始>
// octReg 是用于检查八进制字符串的正则表达式对象。
# <翻译结束>


<原文开始>
// Chr return the ascii string of a number(0-255).
<原文结束>

# <翻译开始>
// Chr 返回一个数字（0-255）对应的ASCII字符串。
# <翻译结束>


<原文开始>
// Ord converts the first byte of a string to a value between 0 and 255.
<原文结束>

# <翻译开始>
// Ord将字符串的第一个字节转换为0到255之间的值。
# <翻译结束>


<原文开始>
// Reverse returns a string which is the reverse of `str`.
<原文结束>

# <翻译开始>
// Reverse 函数返回一个字符串，该字符串是 `str` 的反转字符串。
# <翻译结束>


<原文开始>
// NumberFormat formats a number with grouped thousands.
// `decimals`: Sets the number of decimal points.
// `decPoint`: Sets the separator for the decimal point.
// `thousandsSep`: Sets the thousands' separator.
// See http://php.net/manual/en/function.number-format.php.
<原文结束>

# <翻译开始>
// NumberFormat 对数字进行格式化，添加千位分隔符。
// `decimals`: 设置小数点后的位数。
// `decPoint`: 设置小数点的分隔符。
// `thousandsSep`: 设置千位之间的分隔符。
// 参考：http://php.net/manual/en/function.number-format.php
// 这段注释是Go语言代码中对NumberFormat函数功能和参数的中文解释。
# <翻译结束>







<原文开始>
// Shuffle randomly shuffles a string.
// It considers parameter `str` as unicode string.
<原文结束>

# <翻译开始>
// Shuffle 随机打乱一个字符串。
// 它将参数 `str` 视为unicode字符串。
# <翻译结束>


<原文开始>
// HideStr replaces part of the string `str` to `hide` by `percentage` from the `middle`.
// It considers parameter `str` as unicode string.
<原文结束>

# <翻译开始>
// HideStr 将字符串 `str` 中间部分按 `percentage` 百分比替换为 `hide`。
// 此函数视参数 `str` 为 unicode 字符串。
# <翻译结束>


<原文开始>
// Nl2Br inserts HTML line breaks(`br`|<br />) before all newlines in a string:
// \n\r, \r\n, \r, \n.
// It considers parameter `str` as unicode string.
<原文结束>

# <翻译开始>
// Nl2Br在字符串中的所有换行符前插入HTML换行标签(`br`|<br />)：
// 包括\n\r、\r\n、\r、\n。
// 此函数将参数`str`视为Unicode字符串处理。
# <翻译结束>


<原文开始>
// WordWrap wraps a string to a given number of characters.
// This function supports cut parameters of both english and chinese punctuations.
// TODO: Enable custom cut parameter, see http://php.net/manual/en/function.wordwrap.php.
<原文结束>

# <翻译开始>
// WordWrap 将字符串按照给定的字符数进行换行处理。
// 该函数支持对英文和中文标点符号进行截断处理。
// TODO: 实现自定义截断参数功能，参考 http://php.net/manual/en/function.wordwrap.php.
# <翻译结束>












<原文开始>
// English Punctuations.
<原文结束>

# <翻译开始>
// 英文标点符号
# <翻译结束>


<原文开始>
// Chinese Punctuations.
<原文结束>

# <翻译开始>
// 中文标点符号
# <翻译结束>

