
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
// Time is a wrapper for time.Time for additional features.
<原文结束>

# <翻译开始>
// Time 是一个包装了 time.Time 的结构，用于添加额外的功能。. md5:96d9b7cb3af14206
# <翻译结束>


<原文开始>
// iUnixNano is an interface definition commonly for custom time.Time wrapper.
<原文结束>

# <翻译开始>
// iUnixNano 是一个常用的自定义 time.Time 包装器的接口定义。. md5:5c0387efec09a99b
# <翻译结束>


<原文开始>
// New creates and returns a Time object with given parameter.
// The optional parameter is the time object which can be type of: time.Time/*time.Time, string or integer.
// Example:
// New("2024-10-29")
// New(1390876568)
// New(t) // The t is type of time.Time.
<原文结束>

# <翻译开始>
// New 函数创建并返回一个 Time 对象，使用给定的参数。可选参数是一个时间对象，可以是以下类型：time.Time/*time.Time、字符串或整数。
// 例子：
// New("2024-10-29")
// New(1390876568)
// New(t) // t 是 time.Time 类型。
// md5:6951100c014c4ba9
# <翻译结束>


<原文开始>
// Now creates and returns a time object of now.
<原文结束>

# <翻译开始>
// 现在创建并返回一个表示当前时间的对象。. md5:1cfc3114797b1f98
# <翻译结束>


<原文开始>
// NewFromTime creates and returns a Time object with given time.Time object.
<原文结束>

# <翻译开始>
// NewFromTime 根据给定的time.Time对象创建并返回一个Time对象。. md5:e1cf178ea024f53b
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
// Timestamp 返回时间戳，以秒为单位。. md5:52f3b8b0088c2fab
# <翻译结束>


<原文开始>
// TimestampMilli returns the timestamp in milliseconds.
<原文结束>

# <翻译开始>
// TimestampMilli 返回毫秒级的时间戳。. md5:945db1871b08c49f
# <翻译结束>


<原文开始>
// TimestampMicro returns the timestamp in microseconds.
<原文结束>

# <翻译开始>
// TimestampMicro 返回以微秒为单位的时间戳。. md5:20da1d303fcad848
# <翻译结束>


<原文开始>
// TimestampNano returns the timestamp in nanoseconds.
<原文结束>

# <翻译开始>
// TimestampNano 返回以纳秒为单位的时间戳。. md5:93016ce343f59007
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
// Month returns the month of the year specified by t.
<原文结束>

# <翻译开始>
// Month 返回指定时间t的月份。. md5:84f113a801a5eb29
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
// Millisecond returns the millisecond offset within the second specified by t,
// in the range [0, 999].
<原文结束>

# <翻译开始>
// Millisecond 返回给定时间 t 所在秒内的毫秒偏移，范围为 [0, 999]。
// md5:8bb4c372dc3ada79
# <翻译结束>


<原文开始>
// Microsecond returns the microsecond offset within the second specified by t,
// in the range [0, 999999].
<原文结束>

# <翻译开始>
// Microsecond 返回 t 指定的秒内微秒偏移量，范围为 [0, 999999]。
// md5:cb28fad241f60582
# <翻译结束>


<原文开始>
// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
<原文结束>

# <翻译开始>
// Nanosecond 返回 t 所指定秒内的纳秒偏移量，范围为 [0, 999999999]。
// md5:c1dcd3dd99062cf7
# <翻译结束>


<原文开始>
// String returns current time object as string.
<原文结束>

# <翻译开始>
// String 返回当前时间对象作为字符串。. md5:4f5a1f3896ca049d
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
// Clone returns a new Time object which is a clone of current time object.
<原文结束>

# <翻译开始>
// Clone 返回一个与当前时间对象相克隆的新Time对象。. md5:8a0848cce3c64ef5
# <翻译结束>


<原文开始>
// Add adds the duration to current time.
<原文结束>

# <翻译开始>
// Add 将持续时间添加到当前时间。. md5:8a845aeaaa064af4
# <翻译结束>


<原文开始>
// AddStr parses the given duration as string and adds it to current time.
<原文结束>

# <翻译开始>
// AddStr解析给定的字符串持续时间，并将其添加到当前时间。. md5:3c2278027933d90f
# <翻译结束>


<原文开始>
// UTC converts current time to UTC timezone.
<原文结束>

# <翻译开始>
// UTC 将当前时间转换为UTC时区。. md5:5067cfa0c7c94f95
# <翻译结束>


<原文开始>
// ISO8601 formats the time as ISO8601 and returns it as string.
<原文结束>

# <翻译开始>
// ISO8601将时间格式化为ISO8601标准格式，并以字符串形式返回。. md5:6ddd62f8570c26f4
# <翻译结束>


<原文开始>
// RFC822 formats the time as RFC822 and returns it as string.
<原文结束>

# <翻译开始>
// RFC822 根据 RFC822 格式将时间转换为字符串并返回。. md5:1b6d66ac42df19de
# <翻译结束>


<原文开始>
// AddDate adds year, month and day to the time.
<原文结束>

# <翻译开始>
// AddDate 向时间添加年、月和日。. md5:643cfbc24c5bd938
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
// Before 返回时间点 t 是否在 u 之前。. md5:36690a50c1e8d9d4
# <翻译结束>


<原文开始>
// After reports whether the time instant t is after u.
<原文结束>

# <翻译开始>
// After 判断时间点t是否在u之后。. md5:750eca8bb04e1a25
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
// StartOfMinute 克隆并返回一个新的时间，其中秒数被设置为0。. md5:dc10ea1284a17280
# <翻译结束>


<原文开始>
// StartOfHour clones and returns a new time of which the hour, minutes and seconds are set to 0.
<原文结束>

# <翻译开始>
// StartOfHour克隆并返回一个新的时间，其中小时、分钟和秒设置为0。. md5:d52e77457a157871
# <翻译结束>


<原文开始>
// StartOfDay clones and returns a new time which is the start of day, its time is set to 00:00:00.
<原文结束>

# <翻译开始>
// StartOfDay克隆并返回一个新的时间，它是新的一天的开始，其时间被设置为00:00:00。. md5:a9262cc6eafed6da
# <翻译结束>


<原文开始>
// StartOfWeek clones and returns a new time which is the first day of week and its time is set to
// 00:00:00.
<原文结束>

# <翻译开始>
// StartOfWeek 克隆并返回一个新的时间，该时间为一周的第一天，其时间设置为00:00:00。
// md5:46c7f050c7f59e0a
# <翻译结束>


<原文开始>
// StartOfMonth clones and returns a new time which is the first day of the month and its is set to
// 00:00:00
<原文结束>

# <翻译开始>
// StartOfMonth 创建并返回一个新的时间，该时间是月份的第一天，并且时间设置为 00:00:00
// md5:3de8c28f482566bb
# <翻译结束>


<原文开始>
// StartOfQuarter clones and returns a new time which is the first day of the quarter and its time is set
// to 00:00:00.
<原文结束>

# <翻译开始>
// StartOfQuarter克隆并返回一个新的时间，它是季度的第一天，时间被设置为00:00:00。
// md5:814969ee5c648fb0
# <翻译结束>


<原文开始>
// StartOfHalf clones and returns a new time which is the first day of the half year and its time is set
// to 00:00:00.
<原文结束>

# <翻译开始>
// StartOfHalf克隆并返回一个新的时间，它是半年的第一天，时间被设置为00:00:00。
// md5:5b53c4e328da312e
# <翻译结束>


<原文开始>
// StartOfYear clones and returns a new time which is the first day of the year and its time is set to
// 00:00:00.
<原文结束>

# <翻译开始>
// StartOfYear 克隆并返回一个新的时间，该时间为一年中的第一天，其时间设置为00:00:00。
// md5:7bfbc3ec2e634ff2
# <翻译结束>


<原文开始>
// getPrecisionDelta returns the precision parameter for time calculation depending on `withNanoPrecision` option.
<原文结束>

# <翻译开始>
// getPrecisionDelta 根据`withNanoPrecision`选项返回时间计算的精度参数。. md5:8bcdeaaf0e87d398
# <翻译结束>


<原文开始>
// EndOfMinute clones and returns a new time of which the seconds is set to 59.
<原文结束>

# <翻译开始>
// EndOfMinute克隆并返回一个新的时间，其中秒设置为59。. md5:f1cc1512e831d5fa
# <翻译结束>


<原文开始>
// EndOfHour clones and returns a new time of which the minutes and seconds are both set to 59.
<原文结束>

# <翻译开始>
// EndOfHour克隆并返回一个新的时间，其中分钟和秒都设置为59。. md5:ea49434e1e5b1bbb
# <翻译结束>


<原文开始>
// EndOfDay clones and returns a new time which is the end of day the and its time is set to 23:59:59.
<原文结束>

# <翻译开始>
// EndOfDay 克隆并返回一个新的时间，该时间设置为当天的结束，即时间部分被设置为 23:59:59。. md5:77a284f48ab6cac4
# <翻译结束>


<原文开始>
// EndOfWeek clones and returns a new time which is the end of week and its time is set to 23:59:59.
<原文结束>

# <翻译开始>
// EndOfWeek 创建并返回一个新的时间，该时间表示一周的结束，并将其时间设置为23:59:59。. md5:eb899f421cfb25b4
# <翻译结束>


<原文开始>
// EndOfMonth clones and returns a new time which is the end of the month and its time is set to 23:59:59.
<原文结束>

# <翻译开始>
// EndOfMonth克隆并返回一个新的时间，它是当月的结束，时间设置为23:59:59。. md5:6c2259b48332a891
# <翻译结束>


<原文开始>
// EndOfQuarter clones and returns a new time which is end of the quarter and its time is set to 23:59:59.
<原文结束>

# <翻译开始>
// EndOfQuarter克隆并返回一个新的时间，它是季度结束，其时间设置为23:59:59。. md5:c2e7dca6753c6e99
# <翻译结束>


<原文开始>
// EndOfHalf clones and returns a new time which is the end of the half year and its time is set to 23:59:59.
<原文结束>

# <翻译开始>
// EndOfHalf 克隆并返回一个新的时间，该时间设置为半年的结束时刻，具体时间为 23:59:59。. md5:2f3662f357ee5f6d
# <翻译结束>


<原文开始>
// EndOfYear clones and returns a new time which is the end of the year and its time is set to 23:59:59.
<原文结束>

# <翻译开始>
// EndOfYear 克隆并返回一个新的时间，该时间是当年的年末，时间设置为23:59:59。. md5:33b38d1d0badf6ad
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// Note that, DO NOT use `(t *Time) MarshalJSON() ([]byte, error)` as it looses interface
// implement of `MarshalJSON` for struct of Time.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口的 MarshalJSON 方法。注意，不要使用 `(t *Time) MarshalJSON() ([]byte, error)`，因为它会丢失 Time 结构体的 MarshalJSON 接口实现。
// md5:daef718235a856ce
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
<原文结束>

# <翻译开始>
// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。. md5:f6766b88cf3d63c2
# <翻译结束>


<原文开始>
// UnmarshalText implements the encoding.TextUnmarshaler interface.
// Note that it overwrites the same implementer of `time.Time`.
<原文结束>

# <翻译开始>
// UnmarshalText实现了encoding.TextUnmarshaler接口。
// 注意，它会覆盖与`time.Time`相同的实现者。
// md5:8aa957653e42443a
# <翻译结束>


<原文开始>
// NoValidation marks this struct object will not be validated by package gvalid.
<原文结束>

# <翻译开始>
// NoValidation 标记这个结构体对象将不会被 gvalid 包进行验证。. md5:5241ee7a51fb1912
# <翻译结束>


<原文开始>
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy实现当前类型的深拷贝接口。. md5:9cfbcb08109f6ce1
# <翻译结束>

