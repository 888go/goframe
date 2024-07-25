
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
// Now creates and returns a time object of now.
<原文结束>

# <翻译开始>
// 现在创建并返回一个表示当前时间的对象。 md5:1cfc3114797b1f98
# <翻译结束>


<原文开始>
	// May Output:
	// 2021-11-06 13:41:08
<原文结束>

# <翻译开始>
	// May Output:
	// 2021-11-06 13:41:08
# <翻译结束>


<原文开始>
// NewFromTime creates and returns a Time object with given time.Time object.
<原文结束>

# <翻译开始>
// NewFromTime 根据给定的time.Time对象创建并返回一个Time对象。 md5:e1cf178ea024f53b
# <翻译结束>


<原文开始>
// NewFromStr creates and returns a Time object with given string.
// Note that it returns nil if there's error occurs.
<原文结束>

# <翻译开始>
// NewFromStr 根据给定的字符串创建并返回一个 Time 对象。
// 注意，如果发生错误，它将返回 nil。
// md5:4687b38a27582a12
# <翻译结束>


<原文开始>
// NewFromStrFormat creates and returns a Time object with given string and
// custom format like: Y-m-d H:i:s.
// Note that it returns nil if there's error occurs.
<原文结束>

# <翻译开始>
// NewFromStrFormat 通过给定的字符串和自定义格式（如：Y-m-d H:i:s）创建并返回一个Time对象。
// 注意，如果发生错误，它将返回nil。
// md5:ed9966a0a8156f1d
# <翻译结束>


<原文开始>
// NewFromStrLayout creates and returns a Time object with given string and
// stdlib layout like: 2006-01-02 15:04:05.
// Note that it returns nil if there's error occurs.
<原文结束>

# <翻译开始>
// NewFromStrLayout 根据给定的字符串和标准库格式（如：2006-01-02 15:04:05）创建并返回一个Time对象。
// 注意，如果出现错误，它将返回nil。
// md5:027f4d0876baa1a8
# <翻译结束>


<原文开始>
// NewFromTimeStamp creates and returns a Time object with given timestamp,
// which can be in seconds to nanoseconds.
// Eg: 1600443866 and 1600443866199266000 are both considered as valid timestamp number.
<原文结束>

# <翻译开始>
// NewFromTimeStamp 根据给定的时间戳创建并返回一个 Time 对象，
// 该时间戳可以是秒到纳秒的精度。
// 例如：1600443866 和 1600443866199266000 都被视为有效的时间戳数值。
// md5:6a84edd691c97a4f
# <翻译结束>


<原文开始>
// Timestamp returns the timestamp in seconds.
<原文结束>

# <翻译开始>
// Timestamp 返回时间戳，以秒为单位。 md5:52f3b8b0088c2fab
# <翻译结束>


<原文开始>
	// May output:
	// 1533686888
<原文结束>

# <翻译开始>
	// May output:
	// 1533686888
# <翻译结束>


<原文开始>
// Timestamp returns the timestamp in milliseconds.
<原文结束>

# <翻译开始>
// Timestamp 返回以毫秒为单位的时间戳。 md5:b4836efd766d4f28
# <翻译结束>


<原文开始>
	// May output:
	// 1533686888000
<原文结束>

# <翻译开始>
	// May output:
	// 1533686888000
# <翻译结束>


<原文开始>
// Timestamp returns the timestamp in microseconds.
<原文结束>

# <翻译开始>
// Timestamp 返回时间戳，以微秒为单位。 md5:92d47303429ab4d0
# <翻译结束>


<原文开始>
	// May output:
	// 1533686888000000
<原文结束>

# <翻译开始>
	// May output:
	// 1533686888000000
# <翻译结束>


<原文开始>
// Timestamp returns the timestamp in nanoseconds.
<原文结束>

# <翻译开始>
// Timestamp 返回纳秒级的时间戳。 md5:5f8d54218fb362c4
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
// Month returns the month of the year specified by t.
<原文结束>

# <翻译开始>
// Month 返回指定时间t的月份。 md5:84f113a801a5eb29
# <翻译结束>


<原文开始>
// Second returns the second offset within the minute specified by t,
// in the range [0, 59].
<原文结束>

# <翻译开始>
// Second返回由t指定的分钟内的第二个偏移量，范围在[0, 59]之间。
// md5:5666ae5cbf21989d
# <翻译结束>


<原文开始>
// String returns current time object as string.
<原文结束>

# <翻译开始>
// String 返回当前时间对象作为字符串。 md5:4f5a1f3896ca049d
# <翻译结束>


<原文开始>
// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
<原文结束>

# <翻译开始>
// IsZero报告是否`t`表示零时间点，即UTC时间的1970年1月1日00:00:00。
// md5:4e2b46d4fa63a878
# <翻译结束>


<原文开始>
// Add adds the duration to current time.
<原文结束>

# <翻译开始>
// Add 将持续时间添加到当前时间。 md5:8a845aeaaa064af4
# <翻译结束>


<原文开始>
// AddDate adds year, month and day to the time.
<原文结束>

# <翻译开始>
// AddDate 向时间添加年、月和日。 md5:643cfbc24c5bd938
# <翻译结束>


<原文开始>
// Round returns the result of rounding t to the nearest multiple of d (since the zero time).
// The rounding behavior for halfway values is to round up.
// If d <= 0, Round returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Round operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Round(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
<原文结束>

# <翻译开始>
// Round 返回将 t 四舍五入到 d 的倍数的结果（从零时间开始）。对于半等值，四舍五入行为向上取整。
// 如果 d 小于等于 0，Round 会返回 t 并移除任何单调时钟读数，但保持不变。
//
// Round 以绝对的自零时间以来的时间段进行操作；它不处理时间的呈现形式。因此，Round(Hour) 可能返回一个非零分钟的时间，具体取决于时间的 Location。
// md5:b2557220790fc058
# <翻译结束>


<原文开始>
// Truncate returns the result of rounding t down to a multiple of d (since the zero time).
// If d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Truncate operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Truncate(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
<原文结束>

# <翻译开始>
// Truncate 返回将时间t向下舍入到d的倍数的结果（从零时间开始）。
// 如果d<=0，Truncate会返回t，但去除任何单调时钟读数，否则保持不变。
//
// Truncate是基于时间从零时间点起的绝对持续时间来进行操作的；
// 它并不作用于时间的展示形式。因此，Truncate(Hour)可能返回一个分钟数非零的时间，
// 这取决于该时间的位置信息（Location）。
// md5:f72e0e00b245e691
# <翻译结束>


<原文开始>
// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 CEST and 4:00 UTC are Equal.
// See the documentation on the Time type for the pitfalls of using == with
// Time values; most code should use Equal instead.
<原文结束>

# <翻译开始>
// Equal 函数报告 t 和 u 是否表示相同的时刻。
// 即使两个时间在不同的时区，它们也可以相等。
// 例如，CEST 的 6:00 +0200 和 UTC 的 4:00 是相等的。
// 查看 Time 类型的文档，了解使用 == 操作符比较时间值时可能遇到的问题；
// 大多数代码应使用 Equal 而非 ==。
// md5:a28e147d11d5fe0f
# <翻译结束>


<原文开始>
// Before reports whether the time instant t is before u.
<原文结束>

# <翻译开始>
// Before 返回时间点 t 是否在 u 之前。 md5:36690a50c1e8d9d4
# <翻译结束>


<原文开始>
// After reports whether the time instant t is after u.
<原文结束>

# <翻译开始>
// After 判断时间点t是否在u之后。 md5:750eca8bb04e1a25
# <翻译结束>


<原文开始>
// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
// To compute t-d for a duration d, use t.Add(-d).
<原文结束>

# <翻译开始>
// Sub 返回持续时间 t-u。如果结果超过了能存储在 Duration 类型中的最大（或最小）
// 值，那么将返回最大（或最小）的持续时间。
// 要计算 t-d（其中 d 为一个持续时间），请使用 t.Add(-d)。
// md5:c975e5087c03d3b9
# <翻译结束>


<原文开始>
// StartOfMinute clones and returns a new time of which the seconds is set to 0.
<原文结束>

# <翻译开始>
// StartOfMinute 克隆并返回一个新的时间，其中秒数被设置为0。 md5:dc10ea1284a17280
# <翻译结束>


<原文开始>
	// Output
	// &{goframe 2018-08-08 08:08:08}
<原文结束>

# <翻译开始>
	// 输出
	// &{goframe 2018-08-08 08:08:08} 
	// 
	// 这段Go代码的注释表示这是一个输出（Output），内容是关于一个结构体（&{...}）的引用，该结构体名为goframe，包含了日期和时间信息（2018-08-08 08:08:08）。
	// md5:a93ddd4a9e34a1af
# <翻译结束>

