
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
// Package gtime provides functionality for measuring and displaying time.
//
// This package should keep much less dependencies with other packages.
<原文结束>

# <翻译开始>
// Package gtime 提供了测量和显示时间的功能。
// 
// 本包应该尽量减少与其他包的依赖。
// md5:34aae194a36b5e34
# <翻译结束>


<原文开始>
// Short writes for common usage durations.
<原文结束>

# <翻译开始>
// 为常见的使用时长进行简短写入。. md5:368ca473af5d327a
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
// 正则表达式1（日期时间分隔符支持'-', '/' 和 '.'）。
// 例如：
// "2017-12-14 04:51:34 +0805 LMT",
// "2017/12/14 04:51:34 +0805 LMT",
// "2006-01-02T15:04:05Z07:00",
// "2014-01-17T01:19:15+08:00",
// "2018-02-09T20:46:17.897Z",
// "2018-02-09 20:46:17.897",
// "2018-02-09T20:46:17Z",
// "2018-02-09 20:46:17",
// "2018/10/31 - 16:38:46"
// "2018-02-09",
// "2018.02.09",
// md5:2b97a95934c21f54
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
// 正则表达式（日期时间分隔符支持'-'、'/'和'.'）。
// 例如：
// 01-Nov-2018 11:50:28
// 01/Nov/2018 11:50:28
// 01.Nov.2018 11:50:28
// 01.Nov.2018:11:50:28
// md5:1f5fc72c5b6eb4f8
# <翻译结束>


<原文开始>
	// Regular expression3(time).
	// Eg:
	// 11:50:28
	// 11:50:28.897
<原文结束>

# <翻译开始>
// 正则表达式3（时间）。
// 例如：
// 11:50:28
// 11:50:28.897
// md5:99204487b527a8dc
# <翻译结束>


<原文开始>
	// It's more high performance using regular expression
	// than time.ParseInLocation to parse the datetime string.
<原文结束>

# <翻译开始>
// 使用正则表达式解析日期时间字符串比time.ParseInLocation更高效。
// md5:08fb8b42e551caf6
# <翻译结束>


<原文开始>
// Month words to arabic numerals mapping.
<原文结束>

# <翻译开始>
// 月份英文到阿拉伯数字的映射。. md5:7730d474ee8e496c
# <翻译结束>


<原文开始>
// Timestamp retrieves and returns the timestamp in seconds.
<原文结束>

# <翻译开始>
// Timestamp 获取并返回秒为单位的时间戳。. md5:f859aadcb9d86dfc
# <翻译结束>


<原文开始>
// TimestampMilli retrieves and returns the timestamp in milliseconds.
<原文结束>

# <翻译开始>
// TimestampMilli 获取并返回毫秒级的时间戳。. md5:59f90b1b1bf0bd83
# <翻译结束>


<原文开始>
// TimestampMicro retrieves and returns the timestamp in microseconds.
<原文结束>

# <翻译开始>
// TimestampMicro 获取并返回以微秒为单位的时间戳。. md5:f773c1913603fb89
# <翻译结束>


<原文开始>
// TimestampNano retrieves and returns the timestamp in nanoseconds.
<原文结束>

# <翻译开始>
// TimestampNano 获取并以纳秒为单位返回时间戳。. md5:8b782ae92acea8e7
# <翻译结束>


<原文开始>
// TimestampStr is a convenience method which retrieves and returns
// the timestamp in seconds as string.
<原文结束>

# <翻译开始>
// TimestampStr 是一个方便的方法，它获取并返回时间戳（以秒为单位）的字符串形式。
// md5:f638769b91eb1dd5
# <翻译结束>


<原文开始>
// TimestampMilliStr is a convenience method which retrieves and returns
// the timestamp in milliseconds as string.
<原文结束>

# <翻译开始>
// TimestampMilliStr是一个方便的方法，它获取并返回毫秒级的时间戳作为字符串。
// md5:cf293e6d5c9383d0
# <翻译结束>


<原文开始>
// TimestampMicroStr is a convenience method which retrieves and returns
// the timestamp in microseconds as string.
<原文结束>

# <翻译开始>
// TimestampMicroStr是一个方便的方法，它获取并返回微秒级别的时间戳作为字符串。
// md5:2930c4dc2c5feaae
# <翻译结束>


<原文开始>
// TimestampNanoStr is a convenience method which retrieves and returns
// the timestamp in nanoseconds as string.
<原文结束>

# <翻译开始>
// TimestampNanoStr 是一个便捷方法，用于获取并以字符串形式返回纳秒级的时间戳。
// md5:ff842fbe274c5052
# <翻译结束>


<原文开始>
// Date returns current date in string like "2006-01-02".
<原文结束>

# <翻译开始>
// Date 返回当前日期的字符串，格式为 "2006-01-02"。. md5:e4ebaf573ffc4bd6
# <翻译结束>


<原文开始>
// Datetime returns current datetime in string like "2006-01-02 15:04:05".
<原文结束>

# <翻译开始>
// Datetime 返回当前日期时间的字符串格式，例如 "2006-01-02 15:04:05"。. md5:a1afd811808be0ca
# <翻译结束>


<原文开始>
// ISO8601 returns current datetime in ISO8601 format like "2006-01-02T15:04:05-07:00".
<原文结束>

# <翻译开始>
// ISO8601 返回当前日期时间，格式为 "2006-01-02T15:04:05-07:00"。. md5:ab2ef4b60100081e
# <翻译结束>


<原文开始>
// RFC822 returns current datetime in RFC822 format like "Mon, 02 Jan 06 15:04 MST".
<原文结束>

# <翻译开始>
// RFC822 返回当前日期时间的 RFC822 格式，例如 "Mon, 02 Jan 06 15:04 MST"。. md5:a6362395180caeda
# <翻译结束>


<原文开始>
// parseDateStr parses the string to year, month and day numbers.
<原文结束>

# <翻译开始>
// parseDateStr 将字符串解析为年份、月份和日期数字。. md5:697ca0661d5cecd9
# <翻译结束>


<原文开始>
// Checking the year in head or tail.
<原文结束>

# <翻译开始>
// 检查年份在头部还是尾部。. md5:33266655c259b475
# <翻译结束>


<原文开始>
// StrToTime converts string to *Time object. It also supports timestamp string.
// The parameter `format` is unnecessary, which specifies the format for converting like "Y-m-d H:i:s".
// If `format` is given, it acts as same as function StrToTimeFormat.
// If `format` is not given, it converts string as a "standard" datetime string.
// Note that, it fails and returns error if there's no date string in `str`.
<原文结束>

# <翻译开始>
// StrToTime 将字符串转换为 *Time 对象。它也支持时间戳字符串。
// 参数 `format` 是不必要的，用于指定转换格式，如 "Y-m-d H:i:s"。
// 如果提供了 `format`，它的行为与 StrToTimeFormat 函数相同。
// 如果没有提供 `format`，它将把字符串作为 "标准" 日期时间字符串进行转换。
// 注意，如果 `str` 中没有日期字符串，它将失败并返回错误。
// md5:5e4dd2ec67cb758d
# <翻译结束>


<原文开始>
// Nanoseconds, check and perform bits filling
<原文结束>

# <翻译开始>
// 纳秒，检查并执行位填充. md5:c54cea495b82ade6
# <翻译结束>


<原文开始>
	// If there's zone information in the string,
	// it then performs time zone conversion, which converts the time zone to UTC.
<原文结束>

# <翻译开始>
// 如果字符串中包含时区信息，
// 然后进行时区转换，将时区转换为UTC。
// md5:57a54806130bc3f5
# <翻译结束>


<原文开始>
// If there's offset in the string, it then firstly processes the offset.
<原文结束>

# <翻译开始>
// 如果字符串中有偏移量，那么它会首先处理这个偏移量。. md5:5a183f25f01ee951
# <翻译结束>


<原文开始>
// Comparing the given time zone whether equals to current time zone,
<原文结束>

# <翻译开始>
// 比较给定的时区是否等于当前时区，. md5:7b9947fc7651a35e
# <翻译结束>


<原文开始>
// ConvertZone converts time in string `strTime` from `fromZone` to `toZone`.
// The parameter `fromZone` is unnecessary, it is current time zone in default.
<原文结束>

# <翻译开始>
// ConvertZone 将字符串格式的时间`strTime`从`fromZone`时区转换为`toZone`时区。
// 参数`fromZone`是可选的，默认情况下它代表当前所在的时区。
// md5:9c73950cf06cb368
# <翻译结束>


<原文开始>
// StrToTimeFormat parses string `str` to *Time object with given format `format`.
// The parameter `format` is like "Y-m-d H:i:s".
<原文结束>

# <翻译开始>
// StrToTimeFormat 函数将字符串 `str` 根据给定的格式 `format` 解析为 *Time 对象。
// 参数 `format` 的格式类似于 "Y-m-d H:i:s"。
// md5:0eb1a22261a21da1
# <翻译结束>


<原文开始>
// StrToTimeLayout parses string `str` to *Time object with given format `layout`.
// The parameter `layout` is in stdlib format like "2006-01-02 15:04:05".
<原文结束>

# <翻译开始>
// StrToTimeLayout 将字符串 `str` 解析为具有给定格式 `layout` 的 *Time 对象。参数 `layout` 应使用标准库中的格式，如 "2006-01-02 15:04:05"。
// md5:54702732831e3f2e
# <翻译结束>


<原文开始>
// ParseTimeFromContent retrieves time information for content string, it then parses and returns it
// as *Time object.
// It returns the first time information if there are more than one time string in the content.
// It only retrieves and parses the time information with given first matched `format` if it's passed.
<原文结束>

# <翻译开始>
// ParseTimeFromContent 从内容字符串中提取时间信息，然后解析并返回一个 *Time 类型的对象。
// 如果内容中有多个时间字符串，它将返回第一个时间信息。
// 如果提供了 `format`，它只会检索并解析与之匹配的第一个时间信息。
// md5:37e6a9bec5011038
# <翻译结束>


<原文开始>
// FuncCost calculates the cost time of function `f` in nanoseconds.
<原文结束>

# <翻译开始>
// FuncCost 计算函数 `f` 的执行时间成本，以纳秒为单位。. md5:f6d4e0146ba246f1
# <翻译结束>


<原文开始>
// isTimestampStr checks and returns whether given string a timestamp string.
<原文结束>

# <翻译开始>
// isTimestampStr 检查并返回给定的字符串是否为时间戳字符串。. md5:cb1a63922b0758a1
# <翻译结束>

