
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
// Time is a wrapper for time.Time for additional features.
<原文结束>

# <翻译开始>
// Time 是对 time.Time 的一个封装，用于提供额外功能。
# <翻译结束>


<原文开始>
// iUnixNano is an interface definition commonly for custom time.Time wrapper.
<原文结束>

# <翻译开始>
// iUnixNano 是一个接口定义，通常用于自定义 time.Time 封装器。
# <翻译结束>


<原文开始>
// New creates and returns a Time object with given parameter.
// The optional parameter can be type of: time.Time/*time.Time, string or integer.
<原文结束>

# <翻译开始>
// New 函数根据给定的参数创建并返回一个 Time 对象。
// 可选参数可以是以下类型：time.Time、*time.Time、字符串或整数。
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
// TimestampMilli returns the timestamp in milliseconds.
<原文结束>

# <翻译开始>
// TimestampMilli 返回以毫秒为单位的当前时间戳。
# <翻译结束>


<原文开始>
// TimestampMicro returns the timestamp in microseconds.
<原文结束>

# <翻译开始>
// TimestampMicro 返回以微秒为单位的时间戳。
# <翻译结束>


<原文开始>
// TimestampNano returns the timestamp in nanoseconds.
<原文结束>

# <翻译开始>
// TimestampNano 返回以纳秒为单位的当前时间戳。
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
// Millisecond returns the millisecond offset within the second specified by t,
// in the range [0, 999].
<原文结束>

# <翻译开始>
// Millisecond 返回由 t 指定的秒内毫秒偏移量，范围在 [0, 999] 之间。
# <翻译结束>


<原文开始>
// Microsecond returns the microsecond offset within the second specified by t,
// in the range [0, 999999].
<原文结束>

# <翻译开始>
// Microsecond 返回由 t 指定的秒内微秒偏移量，范围在 [0, 999999] 之间。
# <翻译结束>


<原文开始>
// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
<原文结束>

# <翻译开始>
// Nanosecond 返回由t指定的秒内纳秒偏移量，
// 范围在 [0, 999999999] 之间。
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
// 2024-01-22 不能翻译方法名称.
// IsZero 判断 t 是否代表零时间点，即公元1年1月1日 00:00:00 UTC。
# <翻译结束>


<原文开始>
// Clone returns a new Time object which is a clone of current time object.
<原文结束>

# <翻译开始>
// Clone 返回一个新的 Time 对象，它是当前时间对象的克隆副本。
# <翻译结束>


<原文开始>
// Add adds the duration to current time.
<原文结束>

# <翻译开始>
// Add 将持续时间添加到当前时间。
# <翻译结束>


<原文开始>
// AddStr parses the given duration as string and adds it to current time.
<原文结束>

# <翻译开始>
// AddStr 将给定的以字符串形式表示的时间间隔解析，并将其添加到当前时间。
# <翻译结束>


<原文开始>
// UTC converts current time to UTC timezone.
<原文结束>

# <翻译开始>
// UTC 将当前时间转换为UTC时区。
# <翻译结束>


<原文开始>
// ISO8601 formats the time as ISO8601 and returns it as string.
<原文结束>

# <翻译开始>
// ISO8601将时间格式化为ISO8601标准格式并以字符串形式返回。
# <翻译结束>


<原文开始>
// RFC822 formats the time as RFC822 and returns it as string.
<原文结束>

# <翻译开始>
// RFC822格式化时间并按照RFC822标准返回其字符串表示形式。
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
// StartOfHour clones and returns a new time of which the hour, minutes and seconds are set to 0.
<原文结束>

# <翻译开始>
// StartOfHour 创建并返回一个新的时间，其中小时、分钟和秒被设置为0。
# <翻译结束>


<原文开始>
// StartOfDay clones and returns a new time which is the start of day, its time is set to 00:00:00.
<原文结束>

# <翻译开始>
// StartOfDay 复制并返回一个新的时间，该时间为一天的开始，其时间设置为 00:00:00。
# <翻译结束>


<原文开始>
// StartOfWeek clones and returns a new time which is the first day of week and its time is set to
// 00:00:00.
<原文结束>

# <翻译开始>
// StartOfWeek 创建并返回一个新的时间，该时间为所在周的第一天，并将其时间设置为00:00:00。
# <翻译结束>


<原文开始>
// StartOfMonth clones and returns a new time which is the first day of the month and its is set to
// 00:00:00
<原文结束>

# <翻译开始>
// StartOfMonth 克隆并返回一个新的时间，该时间设置为当月的第一天且其时间为00:00:00
// ```go
// StartOfMonth 创建并返回一个新时间对象的副本，该对象表示所在月份的月初，
// 即第一天，并将其小时、分钟和秒都设定为00:00:00。
# <翻译结束>


<原文开始>
// StartOfQuarter clones and returns a new time which is the first day of the quarter and its time is set
// to 00:00:00.
<原文结束>

# <翻译开始>
// StartOfQuarter 创建并返回一个新的时间，该时间为所在季度的第一天，并将其时间设置为 00:00:00。
# <翻译结束>


<原文开始>
// StartOfHalf clones and returns a new time which is the first day of the half year and its time is set
// to 00:00:00.
<原文结束>

# <翻译开始>
// StartOfHalf 创建并返回一个新的时间副本，该时间是当年上半年的第一天，并将其时间设置为00:00:00。
# <翻译结束>


<原文开始>
// StartOfYear clones and returns a new time which is the first day of the year and its time is set to
// 00:00:00.
<原文结束>

# <翻译开始>
// StartOfYear 创建并返回一个新的时间，该时间设置为当年的第一天，并且其具体时间设置为00:00:00。
# <翻译结束>


<原文开始>
// getPrecisionDelta returns the precision parameter for time calculation depending on `withNanoPrecision` option.
<原文结束>

# <翻译开始>
// getPrecisionDelta 返回一个用于时间计算的精度参数，该参数取决于 `withNanoPrecision` 选项。
# <翻译结束>


<原文开始>
// EndOfMinute clones and returns a new time of which the seconds is set to 59.
<原文结束>

# <翻译开始>
// EndOfMinute复制并返回一个新的时间，其秒数设置为59。
# <翻译结束>


<原文开始>
// EndOfHour clones and returns a new time of which the minutes and seconds are both set to 59.
<原文结束>

# <翻译开始>
// EndOfHour 克隆并返回一个新的时间，其分钟和秒数都被设置为59。
# <翻译结束>


<原文开始>
// EndOfDay clones and returns a new time which is the end of day the and its time is set to 23:59:59.
<原文结束>

# <翻译开始>
// EndOfDay 克隆并返回一个新的时间，该时间设置为原时间所在日期的结束时刻，即 23:59:59。
# <翻译结束>


<原文开始>
// EndOfWeek clones and returns a new time which is the end of week and its time is set to 23:59:59.
<原文结束>

# <翻译开始>
// EndOfWeek 创建并返回一个新的时间对象，该对象为所在周的结束时间，并将其时间设置为 23:59:59。
# <翻译结束>


<原文开始>
// EndOfMonth clones and returns a new time which is the end of the month and its time is set to 23:59:59.
<原文结束>

# <翻译开始>
// EndOfMonth 克隆并返回一个新的时间，该时间为所在月份的月末，并将其时间设置为 23:59:59。
# <翻译结束>


<原文开始>
// EndOfQuarter clones and returns a new time which is end of the quarter and its time is set to 23:59:59.
<原文结束>

# <翻译开始>
// EndOfQuarter 克隆并返回一个新的时间，该时间为季度末，并将其时间设置为 23:59:59。
# <翻译结束>


<原文开始>
// EndOfHalf clones and returns a new time which is the end of the half year and its time is set to 23:59:59.
<原文结束>

# <翻译开始>
// EndOfHalf 克隆并返回一个新的时间，这个时间被设定为半年的结束时刻，并且其具体时间为 23:59:59。
# <翻译结束>


<原文开始>
// EndOfYear clones and returns a new time which is the end of the year and its time is set to 23:59:59.
<原文结束>

# <翻译开始>
// EndOfYear 克隆并返回一个新的时间，该时间为当年的年末，并将其时间设置为 23:59:59。
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// Note that, DO NOT use `(t *Time) MarshalJSON() ([]byte, error)` as it looses interface
// implement of `MarshalJSON` for struct of Time.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
// 注意，切勿使用 `(t *Time) MarshalJSON() ([]byte, error)` 这种形式，因为它会导致 Time 结构体丢失 MarshalJSON 接口的实现。
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
<原文结束>

# <翻译开始>
// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
# <翻译结束>


<原文开始>
// UnmarshalText implements the encoding.TextUnmarshaler interface.
// Note that it overwrites the same implementer of `time.Time`.
<原文结束>

# <翻译开始>
// UnmarshalText 实现了 encoding.TextUnmarshaler 接口。
// 注意，它覆盖了 `time.Time` 同样的实现。
# <翻译结束>


<原文开始>
// NoValidation marks this struct object will not be validated by package gvalid.
<原文结束>

# <翻译开始>
// NoValidation 表示该结构体对象将不会被 gvalid 包进行验证。
# <翻译结束>


<原文开始>
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy 实现接口，用于当前类型的深度复制。
# <翻译结束>

