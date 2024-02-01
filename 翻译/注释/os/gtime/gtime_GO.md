
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
// Package gtime provides functionality for measuring and displaying time.
//
// This package should keep much less dependencies with other packages.
<原文结束>

# <翻译开始>
// Package gtime 提供了用于测量和展示时间的功能。
//
// 本包应尽量减少与其他包的依赖关系。
# <翻译结束>


<原文开始>
// Short writes for common usage durations.
<原文结束>

# <翻译开始>
// 对于常见使用时长提供简写方式。
# <翻译结束>


<原文开始>
	// Regular expression1(datetime separator supports '-', '/', '.').
	// Eg:
	// "2017-12-14 04:51:34 +0805 LMT",
	// "2017-12-14 04:51:34 +0805 LMT",
	// "2006-01-02T15:04:05Z07:00",
	// "2014-01-17T01:19:15+08:00",
	// "2018-02-09T20:46:17.897Z",
	// "2018-02-09 20:46:17.897",
	// "2018-02-09T20:46:17Z",
	// "2018-02-09 20:46:17",
	// "2018/10/31 - 16:38:46"
	// "2018-02-09",
	// "2018.02.09",
<原文结束>

# <翻译开始>
// 正则表达式1（日期时间分隔符支持 '-', '/', '.'）。
// 示例：
// "2017-12-14 04:51:34 +0805 LMT",
// "2017/12/14 04:51:34 +0805 LMT",
// "2006-01-02T15:04:05Z07:00" （ISO 8601格式，带时区偏移量）
// "2014-01-17T01:19:15+08:00" （ISO 8601格式，带时区偏移量）
// "2018-02-09T20:46:17.897Z" （ISO 8601格式，带毫秒精度和UTC时区）
// "2018-02-09 20:46:17.897" （带毫秒精度）
// "2018-02-09T20:46:17Z" （ISO 8601格式，带UTC时区）
// "2018-02-09 20:46:17" （标准日期时间格式）
// "2018/10/31 - 16:38:46" （使用'/'作为分隔符的日期时间格式）
// "2018-02-09" （仅日期格式）
// "2018.02.09" （使用'.'作为分隔符的日期格式）
// 此段代码注释描述了该正则表达式用于匹配多种日期时间格式，这些格式中使用的分隔符可以是短横线`-`、斜线`/`或点`.`。同时列举了一系列符合此正则表达式的示例字符串，涵盖了不同的日期时间表示方式以及时区信息。
# <翻译结束>


<原文开始>
	// Regular expression2(datetime separator supports '-', '/', '.').
	// Eg:
	// 01-Nov-2018 11:50:28
	// 01/Nov/2018 11:50:28
	// 01.Nov.2018 11:50:28
	// 01.Nov.2018:11:50:28
<原文结束>

# <翻译开始>
// 正则表达式2（日期时间分隔符支持'-'、'/'和'.'）。
// 示例：
// 01-Nov-2018 11:50:28
// 01/Nov/2018 11:50:28
// 01.Nov.2018 11:50:28
// 01.Nov.2018:11:50:28
// 此正则表达式用于匹配格式类似于上述样例的日期时间字符串，其中日期和月份之间的分隔符可以是短横线、斜线或点，并且在某些情况下，时间部分也可以使用冒号与日期部分连接。
# <翻译结束>


<原文开始>
	// Regular expression3(time).
	// Eg:
	// 11:50:28
	// 11:50:28.897
<原文结束>

# <翻译开始>
// 正则表达式3（时间）
// 示例：
// 11:50:28
// 11:50:28.897
// ```go
// 正则表达式用于匹配时间格式，例如：
// 格式1：小时:分钟:秒
// 格式2：小时:分钟:秒.毫秒
// 例如：
// 示例1：11:50:28
// 示例2：11:50:28.897
# <翻译结束>


<原文开始>
	// It's more high performance using regular expression
	// than time.ParseInLocation to parse the datetime string.
<原文结束>

# <翻译开始>
// 使用正则表达式解析日期时间字符串比使用time.ParseInLocation性能更高。
# <翻译结束>


<原文开始>
// Month words to arabic numerals mapping.
<原文结束>

# <翻译开始>
// 月份单词到阿拉伯数字的映射。
# <翻译结束>


<原文开始>
// Timestamp retrieves and returns the timestamp in seconds.
<原文结束>

# <翻译开始>
// Timestamp 获取并返回以秒为单位的时间戳。
# <翻译结束>


<原文开始>
// TimestampMilli retrieves and returns the timestamp in milliseconds.
<原文结束>

# <翻译开始>
// TimestampMilli 获取并返回毫秒级的时间戳。
# <翻译结束>


<原文开始>
// TimestampMicro retrieves and returns the timestamp in microseconds.
<原文结束>

# <翻译开始>
// TimestampMicro 获取并返回微秒级的时间戳。
# <翻译结束>


<原文开始>
// TimestampNano retrieves and returns the timestamp in nanoseconds.
<原文结束>

# <翻译开始>
// TimestampNano 获取并返回纳秒级别的时间戳。
# <翻译结束>


<原文开始>
// TimestampStr is a convenience method which retrieves and returns
// the timestamp in seconds as string.
<原文结束>

# <翻译开始>
// TimestampStr 是一个便捷方法，用于获取并返回以字符串形式表示的秒级时间戳。
# <翻译结束>


<原文开始>
// TimestampMilliStr is a convenience method which retrieves and returns
// the timestamp in milliseconds as string.
<原文结束>

# <翻译开始>
// TimestampMilliStr 是一个便捷方法，用于获取并返回以字符串形式表示的毫秒级时间戳。
# <翻译结束>


<原文开始>
// TimestampMicroStr is a convenience method which retrieves and returns
// the timestamp in microseconds as string.
<原文结束>

# <翻译开始>
// TimestampMicroStr 是一个便捷方法，用于获取并返回微秒级的时间戳字符串。
# <翻译结束>


<原文开始>
// TimestampNanoStr is a convenience method which retrieves and returns
// the timestamp in nanoseconds as string.
<原文结束>

# <翻译开始>
// TimestampNanoStr 是一个便捷方法，用于获取并返回纳秒级时间戳的字符串表示。
# <翻译结束>


<原文开始>
// Date returns current date in string like "2006-01-02".
<原文结束>

# <翻译开始>
// Date 返回当前日期的字符串，格式如 "2006-01-02"。
# <翻译结束>


<原文开始>
// Datetime returns current datetime in string like "2006-01-02 15:04:05".
<原文结束>

# <翻译开始>
// Datetime 返回当前日期时间的字符串，格式如 "2006-01-02 15:04:05"。
# <翻译结束>


<原文开始>
// ISO8601 returns current datetime in ISO8601 format like "2006-01-02T15:04:05-07:00".
<原文结束>

# <翻译开始>
// ISO8601 返回当前日期时间，格式为符合ISO 8601标准的字符串，如 "2006-01-02T15:04:05-07:00"。
# <翻译结束>


<原文开始>
// RFC822 returns current datetime in RFC822 format like "Mon, 02 Jan 06 15:04 MST".
<原文结束>

# <翻译开始>
// RFC822 返回当前日期时间的 RFC822 格式，例如 "Mon, 02 Jan 06 15:04 MST"。
# <翻译结束>


<原文开始>
// parseDateStr parses the string to year, month and day numbers.
<原文结束>

# <翻译开始>
// parseDateStr 将字符串解析为年、月和日的数字表示。
# <翻译结束>


<原文开始>
// Checking the year in head or tail.
<原文结束>

# <翻译开始>
// 检查年份在头部或尾部。
# <翻译结束>


<原文开始>
// StrToTime converts string to *Time object. It also supports timestamp string.
// The parameter `format` is unnecessary, which specifies the format for converting like "Y-m-d H:i:s".
// If `format` is given, it acts as same as function StrToTimeFormat.
// If `format` is not given, it converts string as a "standard" datetime string.
// Note that, it fails and returns error if there's no date string in `str`.
<原文结束>

# <翻译开始>
// StrToTime 将字符串转换为 *Time 类型对象。同时支持时间戳字符串。
// 参数 `format` 不是必需的，用于指定类似 "Y-m-d H:i:s" 的转换格式。
// 如果提供了 `format`，它的行为与函数 StrToTimeFormat 相同。
// 如果未提供 `format`，它将字符串按“标准”日期时间格式进行转换。
// 注意，如果 `str` 中不包含日期字符串，则转换失败并返回错误。
# <翻译结束>


<原文开始>
// Nanoseconds, check and perform bits filling
<原文结束>

# <翻译开始>
// 纳秒，检查并执行位填充
# <翻译结束>


<原文开始>
	// If there's zone information in the string,
	// it then performs time zone conversion, which converts the time zone to UTC.
<原文结束>

# <翻译开始>
// 如果字符串中包含时区信息，
// 则进行时区转换操作，将时间转换为 UTC 格式。
# <翻译结束>


<原文开始>
// If there's offset in the string, it then firstly processes the offset.
<原文结束>

# <翻译开始>
// 如果字符串中存在偏移量，则首先处理该偏移量。
# <翻译结束>


<原文开始>
			// Comparing the given time zone whether equals to current time zone,
			// it converts it to UTC if they do not equal.
<原文结束>

# <翻译开始>
// 检查给定时区是否等于当前时区，
// 如果不相等，则将其转换为UTC。
# <翻译结束>












<原文开始>
// ConvertZone converts time in string `strTime` from `fromZone` to `toZone`.
// The parameter `fromZone` is unnecessary, it is current time zone in default.
<原文结束>

# <翻译开始>
// ConvertZone 将字符串 `strTime` 中的时间从 `fromZone` 转换为 `toZone`。
// 参数 `fromZone` 在默认情况下是不必要的，它表示当前时区。
# <翻译结束>


<原文开始>
// StrToTimeFormat parses string `str` to *Time object with given format `format`.
// The parameter `format` is like "Y-m-d H:i:s".
<原文结束>

# <翻译开始>
// StrToTimeFormat将字符串`str`按照给定的格式`format`解析为*time.Time对象。
// 参数`format`如"Y-m-d H:i:s"。
# <翻译结束>


<原文开始>
// StrToTimeLayout parses string `str` to *Time object with given format `layout`.
// The parameter `layout` is in stdlib format like "2006-01-02 15:04:05".
<原文结束>

# <翻译开始>
// StrToTimeLayout将字符串`str`按照给定的格式`layout`解析为*time.Time对象。
// 参数`layout`遵循标准库格式，如"2006-01-02 15:04:05"。
# <翻译结束>


<原文开始>
// ParseTimeFromContent retrieves time information for content string, it then parses and returns it
// as *Time object.
// It returns the first time information if there are more than one time string in the content.
// It only retrieves and parses the time information with given first matched `format` if it's passed.
<原文结束>

# <翻译开始>
// ParseTimeFromContent 从内容字符串中检索时间信息，然后将其解析并作为 *Time 对象返回。
// 如果内容中存在多个时间字符串，则返回第一个时间信息。
// 如果传递了匹配的 `format` 参数，那么它将只检索并解析与第一个匹配格式对应的时间信息。
# <翻译结束>


<原文开始>
// FuncCost calculates the cost time of function `f` in nanoseconds.
<原文结束>

# <翻译开始>
// FuncCost 计算函数 `f` 的执行耗时，单位为纳秒。
# <翻译结束>


<原文开始>
// isTimestampStr checks and returns whether given string a timestamp string.
<原文结束>

# <翻译开始>
// isTimestampStr 检查并返回给定字符串是否为时间戳字符串。
# <翻译结束>







<原文开始>
// Comparing in seconds.
<原文结束>

# <翻译开始>
// 按秒进行比较。
# <翻译结束>

