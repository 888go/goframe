
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
// Now creates and returns a time object of now.
<原文结束>

# <翻译开始>
// Now 创建并返回一个表示当前时间的时间对象。
# <翻译结束>


<原文开始>
// NewFromTime creates and returns a Time object with given time.Time object.
<原文结束>

# <翻译开始>
// NewFromTime 根据给定的 time.Time 对象创建并返回一个 Time 对象。
# <翻译结束>


<原文开始>
// NewFromStr creates and returns a Time object with given string.
// Note that it returns nil if there's error occurs.
<原文结束>

# <翻译开始>
// NewFromStr 通过给定的字符串创建并返回一个 Time 对象。
// 注意，如果发生错误，它将返回 nil。
# <翻译结束>


<原文开始>
// NewFromStrFormat creates and returns a Time object with given string and
// custom format like: Y-m-d H:i:s.
// Note that it returns nil if there's error occurs.
<原文结束>

# <翻译开始>
// NewFromStrFormat 根据给定的字符串和自定义格式（如：Y-m-d H:i:s）创建并返回一个Time对象。
// 需要注意，如果发生错误，它将返回nil。
# <翻译结束>


<原文开始>
// NewFromStrLayout creates and returns a Time object with given string and
// stdlib layout like: 2006-01-02 15:04:05.
// Note that it returns nil if there's error occurs.
<原文结束>

# <翻译开始>
// NewFromStrLayout 根据给定的字符串和标准库布局（如：2006-01-02 15:04:05）创建并返回一个Time对象。
// 需要注意的是，如果发生错误，它将返回nil。
# <翻译结束>


<原文开始>
// NewFromTimeStamp creates and returns a Time object with given timestamp,
// which can be in seconds to nanoseconds.
// Eg: 1600443866 and 1600443866199266000 are both considered as valid timestamp number.
<原文结束>

# <翻译开始>
// NewFromTimeStamp 根据给定的时间戳创建并返回一个 Time 对象，
// 时间戳可以是秒级到纳秒级精度。
// 例如：1600443866 和 1600443866199266000 都被认为是有效的时间戳数值。
# <翻译结束>


<原文开始>
// Timestamp returns the timestamp in seconds.
<原文结束>

# <翻译开始>
// Timestamp 返回以秒为单位的时间戳。
# <翻译结束>


<原文开始>
// Timestamp returns the timestamp in milliseconds.
<原文结束>

# <翻译开始>
// Timestamp 返回以毫秒为单位的时间戳。
# <翻译结束>


<原文开始>
// Timestamp returns the timestamp in microseconds.
<原文结束>

# <翻译开始>
// Timestamp 返回以微秒为单位的的时间戳。
# <翻译结束>


<原文开始>
// Timestamp returns the timestamp in nanoseconds.
<原文结束>

# <翻译开始>
// Timestamp 返回以纳秒为单位的时间戳。
# <翻译结束>


<原文开始>
// TimestampStr is a convenience method which retrieves and returns
// the timestamp in seconds as string.
<原文结束>

# <翻译开始>
// TimestampStr 是一个便捷方法，用于获取并返回以字符串形式表示的秒级时间戳。
# <翻译结束>


<原文开始>
// Month returns the month of the year specified by t.
<原文结束>

# <翻译开始>
// Month 返回由 t 指定的年份中的月份。
# <翻译结束>


<原文开始>
// Second returns the second offset within the minute specified by t,
// in the range [0, 59].
<原文结束>

# <翻译开始>
// Second 返回给定时间 t 的分钟内第二个偏移量，
// 范围在 [0, 59] 内。
# <翻译结束>


<原文开始>
// String returns current time object as string.
<原文结束>

# <翻译开始>
// String 将当前时间对象转换为字符串并返回。
# <翻译结束>


<原文开始>
// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
<原文结束>

# <翻译开始>
// IsZero 判断 t 是否代表零时间点，即公元1年1月1日 00:00:00 UTC。
# <翻译结束>


<原文开始>
// Add adds the duration to current time.
<原文结束>

# <翻译开始>
// Add 将持续时间添加到当前时间。
# <翻译结束>


<原文开始>
// AddDate adds year, month and day to the time.
<原文结束>

# <翻译开始>
// AddDate 向时间添加年、月和日。
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
// Round 函数将 t 舍入到最接近 d 的倍数（以零时间点为基准）。
// 对于刚好位于中间值的舍入行为是向上舍入。
// 如果 d 小于等于 0，Round 函数将返回剥离了单调时钟读数但其他部分保持不变的 t。
//
// Round 函数针对的是以零时间为基准的绝对持续时间上的时间；
// 它并不作用在时间的表现形式上。因此，即使调用 Round(Hour)，
// 返回的时间也可能存在非零分钟值，这取决于时间所处的 Location（时区）。
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
// Truncate 方法将 t 向下舍入至 d 的倍数（以零时间点为基准）。
// 若 d 小于等于 0，Truncate 方法会返回剥离了单调时钟读数但其他部分保持不变的 t。
//
// Truncate 对时间进行操作时将其视为从零时间点开始的绝对持续时间；
// 它并不直接作用于时间的展示形式。因此，调用 Truncate(Hour) 可能会返回一个分钟不为零的时间，
// 具体取决于该时间的位置（Location）。
# <翻译结束>


<原文开始>
// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 CEST and 4:00 UTC are Equal.
// See the documentation on the Time type for the pitfalls of using == with
// Time values; most code should use Equal instead.
<原文结束>

# <翻译开始>
// Equal 判断 t 和 u 是否表示相同的时刻。
// 即使两个时间位于不同的时区，它们也可能相等。
// 例如，6:00 +0200 CEST（中欧夏令时）和 4:00 UTC 是相等的。
// 查看 Time 类型的文档，了解使用 == 操作符比较 Time 值时的陷阱；
// 大多数代码应使用 Equal 方法代替。
# <翻译结束>


<原文开始>
// Before reports whether the time instant t is before u.
<原文结束>

# <翻译开始>
// Before 判断时间点 t 是否在时间点 u 之前。
# <翻译结束>


<原文开始>
// After reports whether the time instant t is after u.
<原文结束>

# <翻译开始>
// After 判断时间点 t 是否在时间点 u 之后。
# <翻译结束>


<原文开始>
// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
// To compute t-d for a duration d, use t.Add(-d).
<原文结束>

# <翻译开始>
// Sub 计算并返回时间段 t-u。如果结果超出了 Duration 类型能够存储的最大（或最小）值，
// 则会返回最大（或最小）的有效持续时间。
// 若要计算 t 与一个持续时间 d 的差值（t-d），请使用 t.Add(-d)。
# <翻译结束>


<原文开始>
// StartOfMinute clones and returns a new time of which the seconds is set to 0.
<原文结束>

# <翻译开始>
// StartOfMinute 复制并返回一个新的时间，其秒数设置为0。
# <翻译结束>

 
<原文开始>
	// Output
	// &{goframe 2018-08-08 08:08:08}
<原文结束>

# <翻译开始>
// 输出
// &{goframe 2018-08-08 08:08:08} 
// （这段注释表明该段代码的执行结果会输出一个包含"goframe"和"2018-08-08 08:08:08"信息的数据结构，并以地址引用形式显示。）
# <翻译结束>

