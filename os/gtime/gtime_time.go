// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtime

import (
	"bytes"
	"strconv"
	"time"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
)

// Time 是对 time.Time 的一个封装，用于提供额外功能。
type Time struct {
	wrapper
}

// iUnixNano 是一个接口定义，通常用于自定义 time.Time 封装器。
type iUnixNano interface {
	UnixNano() int64
}

// New 函数根据给定的参数创建并返回一个 Time 对象。
// 可选参数可以是以下类型：time.Time、*time.Time、字符串或整数。
func New(param ...interface{}) *Time {
	if len(param) > 0 {
		switch r := param[0].(type) {
		case time.Time:
			return NewFromTime(r)
		case *time.Time:
			return NewFromTime(*r)

		case Time:
			return &r

		case *Time:
			return r

		case string:
			if len(param) > 1 {
				switch t := param[1].(type) {
				case string:
					return NewFromStrFormat(r, t)
				case []byte:
					return NewFromStrFormat(r, string(t))
				}
			}
			return NewFromStr(r)

		case []byte:
			if len(param) > 1 {
				switch t := param[1].(type) {
				case string:
					return NewFromStrFormat(string(r), t)
				case []byte:
					return NewFromStrFormat(string(r), string(t))
				}
			}
			return NewFromStr(string(r))

		case int:
			return NewFromTimeStamp(int64(r))

		case int64:
			return NewFromTimeStamp(r)

		default:
			if v, ok := r.(iUnixNano); ok {
				return NewFromTimeStamp(v.UnixNano())
			}
		}
	}
	return &Time{
		wrapper{time.Time{}},
	}
}

// Now 创建并返回一个表示当前时间的时间对象。
func Now() *Time {
	return &Time{
		wrapper{time.Now()},
	}
}

// NewFromTime 根据给定的 time.Time 对象创建并返回一个 Time 对象。
func NewFromTime(t time.Time) *Time {
	return &Time{
		wrapper{t},
	}
}

// NewFromStr 通过给定的字符串创建并返回一个 Time 对象。
// 注意，如果发生错误，它将返回 nil。
func NewFromStr(str string) *Time {
	if t, err := StrToTime(str); err == nil {
		return t
	}
	return nil
}

// NewFromStrFormat 根据给定的字符串和自定义格式（如：Y-m-d H:i:s）创建并返回一个Time对象。
// 需要注意，如果发生错误，它将返回nil。
func NewFromStrFormat(str string, format string) *Time {
	if t, err := StrToTimeFormat(str, format); err == nil {
		return t
	}
	return nil
}

// NewFromStrLayout 根据给定的字符串和标准库布局（如：2006-01-02 15:04:05）创建并返回一个Time对象。
// 需要注意的是，如果发生错误，它将返回nil。
func NewFromStrLayout(str string, layout string) *Time {
	if t, err := StrToTimeLayout(str, layout); err == nil {
		return t
	}
	return nil
}

// NewFromTimeStamp 根据给定的时间戳创建并返回一个 Time 对象，
// 时间戳可以是秒级到纳秒级精度。
// 例如：1600443866 和 1600443866199266000 都被认为是有效的时间戳数值。
func NewFromTimeStamp(timestamp int64) *Time {
	if timestamp == 0 {
		return &Time{}
	}
	var sec, nano int64
	if timestamp > 1e9 {
		for timestamp < 1e18 {
			timestamp *= 10
		}
		sec = timestamp / 1e9
		nano = timestamp % 1e9
	} else {
		sec = timestamp
	}
	return &Time{
		wrapper{time.Unix(sec, nano)},
	}
}

// Timestamp 返回以秒为单位的时间戳。
func (t *Time) Timestamp() int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixNano() / 1e9
}

// TimestampMilli 返回以毫秒为单位的当前时间戳。
func (t *Time) TimestampMilli() int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixNano() / 1e6
}

// TimestampMicro 返回以微秒为单位的时间戳。
func (t *Time) TimestampMicro() int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixNano() / 1e3
}

// TimestampNano 返回以纳秒为单位的当前时间戳。
func (t *Time) TimestampNano() int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixNano()
}

// TimestampStr 是一个便捷方法，用于获取并返回以字符串形式表示的秒级时间戳。
func (t *Time) TimestampStr() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.Timestamp(), 10)
}

// TimestampMilliStr 是一个便捷方法，用于获取并返回以字符串形式表示的毫秒级时间戳。
func (t *Time) TimestampMilliStr() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.TimestampMilli(), 10)
}

// TimestampMicroStr 是一个便捷方法，用于获取并返回微秒级的时间戳字符串。
func (t *Time) TimestampMicroStr() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.TimestampMicro(), 10)
}

// TimestampNanoStr 是一个便捷方法，用于获取并返回纳秒级时间戳的字符串表示。
func (t *Time) TimestampNanoStr() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.TimestampNano(), 10)
}

// Month 返回由 t 指定的年份中的月份。
func (t *Time) Month() int {
	if t.IsZero() {
		return 0
	}
	return int(t.Time.Month())
}

// Second 返回给定时间 t 的分钟内第二个偏移量，
// 范围在 [0, 59] 内。
func (t *Time) Second() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Second()
}

// Millisecond 返回由 t 指定的秒内毫秒偏移量，范围在 [0, 999] 之间。
func (t *Time) Millisecond() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Nanosecond() / 1e6
}

// Microsecond 返回由 t 指定的秒内微秒偏移量，范围在 [0, 999999] 之间。
func (t *Time) Microsecond() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Nanosecond() / 1e3
}

// Nanosecond 返回由t指定的秒内纳秒偏移量，
// 范围在 [0, 999999999] 之间。
func (t *Time) Nanosecond() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Nanosecond()
}

// String 将当前时间对象转换为字符串并返回。
func (t *Time) String() string {
	if t.IsZero() {
		return ""
	}
	return t.wrapper.String()
}

// 2024-01-22 不能翻译方法名称.
// IsZero 判断 t 是否代表零时间点，即公元1年1月1日 00:00:00 UTC。
func (t *Time) IsZero() bool {
	if t == nil {
		return true
	}
	return t.Time.IsZero()
}

// Clone 返回一个新的 Time 对象，它是当前时间对象的克隆副本。
func (t *Time) Clone() *Time {
	return New(t.Time)
}

// Add 将持续时间添加到当前时间。
func (t *Time) Add(d time.Duration) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Add(d)
	return newTime
}

// AddStr 将给定的以字符串形式表示的时间间隔解析，并将其添加到当前时间。
func (t *Time) AddStr(duration string) (*Time, error) {
	if d, err := time.ParseDuration(duration); err != nil {
		err = gerror.Wrapf(err, `time.ParseDuration failed for string "%s"`, duration)
		return nil, err
	} else {
		return t.Add(d), nil
	}
}

// UTC 将当前时间转换为UTC时区。
func (t *Time) UTC() *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.UTC()
	return newTime
}

// ISO8601将时间格式化为ISO8601标准格式并以字符串形式返回。
func (t *Time) ISO8601() string {
	return t.Layout("2006-01-02T15:04:05-07:00")
}

// RFC822格式化时间并按照RFC822标准返回其字符串表示形式。
func (t *Time) RFC822() string {
	return t.Layout("Mon, 02 Jan 06 15:04 MST")
}

// AddDate 向时间添加年、月和日。
func (t *Time) AddDate(years int, months int, days int) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.AddDate(years, months, days)
	return newTime
}

// Round 函数将 t 舍入到最接近 d 的倍数（以零时间点为基准）。
// 对于刚好位于中间值的舍入行为是向上舍入。
// 如果 d 小于等于 0，Round 函数将返回剥离了单调时钟读数但其他部分保持不变的 t。
//
// Round 函数针对的是以零时间为基准的绝对持续时间上的时间；
// 它并不作用在时间的表现形式上。因此，即使调用 Round(Hour)，
// 返回的时间也可能存在非零分钟值，这取决于时间所处的 Location（时区）。
func (t *Time) Round(d time.Duration) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Round(d)
	return newTime
}

// Truncate 方法将 t 向下舍入至 d 的倍数（以零时间点为基准）。
// 若 d 小于等于 0，Truncate 方法会返回剥离了单调时钟读数但其他部分保持不变的 t。
//
// Truncate 对时间进行操作时将其视为从零时间点开始的绝对持续时间；
// 它并不直接作用于时间的展示形式。因此，调用 Truncate(Hour) 可能会返回一个分钟不为零的时间，
// 具体取决于该时间的位置（Location）。
func (t *Time) Truncate(d time.Duration) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Truncate(d)
	return newTime
}

// Equal 判断 t 和 u 是否表示相同的时刻。
// 即使两个时间位于不同的时区，它们也可能相等。
// 例如，6:00 +0200 CEST（中欧夏令时）和 4:00 UTC 是相等的。
// 查看 Time 类型的文档，了解使用 == 操作符比较 Time 值时的陷阱；
// 大多数代码应使用 Equal 方法代替。
func (t *Time) Equal(u *Time) bool {
	switch {
	case t == nil && u != nil:
		return false
	case t == nil && u == nil:
		return true
	case t != nil && u == nil:
		return false
	default:
		return t.Time.Equal(u.Time)
	}
}

// Before 判断时间点 t 是否在时间点 u 之前。
func (t *Time) Before(u *Time) bool {
	return t.Time.Before(u.Time)
}

// After 判断时间点 t 是否在时间点 u 之后。
func (t *Time) After(u *Time) bool {
	switch {
	case t == nil:
		return false
	case t != nil && u == nil:
		return true
	default:
		return t.Time.After(u.Time)
	}
}

// Sub 计算并返回时间段 t-u。如果结果超出了 Duration 类型能够存储的最大（或最小）值，
// 则会返回最大（或最小）的有效持续时间。
// 若要计算 t 与一个持续时间 d 的差值（t-d），请使用 t.Add(-d)。
func (t *Time) Sub(u *Time) time.Duration {
	if t == nil || u == nil {
		return 0
	}
	return t.Time.Sub(u.Time)
}

// StartOfMinute 复制并返回一个新的时间，其秒数设置为0。
func (t *Time) StartOfMinute() *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Truncate(time.Minute)
	return newTime
}

// StartOfHour 创建并返回一个新的时间，其中小时、分钟和秒被设置为0。
func (t *Time) StartOfHour() *Time {
	y, m, d := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(y, m, d, newTime.Time.Hour(), 0, 0, 0, newTime.Time.Location())
	return newTime
}

// StartOfDay 复制并返回一个新的时间，该时间为一天的开始，其时间设置为 00:00:00。
func (t *Time) StartOfDay() *Time {
	y, m, d := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(y, m, d, 0, 0, 0, 0, newTime.Time.Location())
	return newTime
}

// StartOfWeek 创建并返回一个新的时间，该时间为所在周的第一天，并将其时间设置为00:00:00。
func (t *Time) StartOfWeek() *Time {
	weekday := int(t.Weekday())
	return t.StartOfDay().AddDate(0, 0, -weekday)
}

// StartOfMonth 克隆并返回一个新的时间，该时间设置为当月的第一天且其时间为00:00:00
// ```go
// StartOfMonth 创建并返回一个新时间对象的副本，该对象表示所在月份的月初，
// 即第一天，并将其小时、分钟和秒都设定为00:00:00。
func (t *Time) StartOfMonth() *Time {
	y, m, _ := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(y, m, 1, 0, 0, 0, 0, newTime.Time.Location())
	return newTime
}

// StartOfQuarter 创建并返回一个新的时间，该时间为所在季度的第一天，并将其时间设置为 00:00:00。
func (t *Time) StartOfQuarter() *Time {
	month := t.StartOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

// StartOfHalf 创建并返回一个新的时间副本，该时间是当年上半年的第一天，并将其时间设置为00:00:00。
func (t *Time) StartOfHalf() *Time {
	month := t.StartOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return month.AddDate(0, -offset, 0)
}

// StartOfYear 创建并返回一个新的时间，该时间设置为当年的第一天，并且其具体时间设置为00:00:00。
func (t *Time) StartOfYear() *Time {
	y, _, _ := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(y, time.January, 1, 0, 0, 0, 0, newTime.Time.Location())
	return newTime
}

// getPrecisionDelta 返回一个用于时间计算的精度参数，该参数取决于 `withNanoPrecision` 选项。
func getPrecisionDelta(withNanoPrecision ...bool) time.Duration {
	if len(withNanoPrecision) > 0 && withNanoPrecision[0] {
		return time.Nanosecond
	}
	return time.Second
}

// EndOfMinute复制并返回一个新的时间，其秒数设置为59。
func (t *Time) EndOfMinute(withNanoPrecision ...bool) *Time {
	return t.StartOfMinute().Add(time.Minute - getPrecisionDelta(withNanoPrecision...))
}

// EndOfHour 克隆并返回一个新的时间，其分钟和秒数都被设置为59。
func (t *Time) EndOfHour(withNanoPrecision ...bool) *Time {
	return t.StartOfHour().Add(time.Hour - getPrecisionDelta(withNanoPrecision...))
}

// EndOfDay 克隆并返回一个新的时间，该时间设置为原时间所在日期的结束时刻，即 23:59:59。
func (t *Time) EndOfDay(withNanoPrecision ...bool) *Time {
	y, m, d := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(
		y, m, d, 23, 59, 59, int(time.Second-getPrecisionDelta(withNanoPrecision...)), newTime.Time.Location(),
	)
	return newTime
}

// EndOfWeek 创建并返回一个新的时间对象，该对象为所在周的结束时间，并将其时间设置为 23:59:59。
func (t *Time) EndOfWeek(withNanoPrecision ...bool) *Time {
	return t.StartOfWeek().AddDate(0, 0, 7).Add(-getPrecisionDelta(withNanoPrecision...))
}

// EndOfMonth 克隆并返回一个新的时间，该时间为所在月份的月末，并将其时间设置为 23:59:59。
func (t *Time) EndOfMonth(withNanoPrecision ...bool) *Time {
	return t.StartOfMonth().AddDate(0, 1, 0).Add(-getPrecisionDelta(withNanoPrecision...))
}

// EndOfQuarter 克隆并返回一个新的时间，该时间为季度末，并将其时间设置为 23:59:59。
func (t *Time) EndOfQuarter(withNanoPrecision ...bool) *Time {
	return t.StartOfQuarter().AddDate(0, 3, 0).Add(-getPrecisionDelta(withNanoPrecision...))
}

// EndOfHalf 克隆并返回一个新的时间，这个时间被设定为半年的结束时刻，并且其具体时间为 23:59:59。
func (t *Time) EndOfHalf(withNanoPrecision ...bool) *Time {
	return t.StartOfHalf().AddDate(0, 6, 0).Add(-getPrecisionDelta(withNanoPrecision...))
}

// EndOfYear 克隆并返回一个新的时间，该时间为当年的年末，并将其时间设置为 23:59:59。
func (t *Time) EndOfYear(withNanoPrecision ...bool) *Time {
	return t.StartOfYear().AddDate(1, 0, 0).Add(-getPrecisionDelta(withNanoPrecision...))
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
// 注意，切勿使用 `(t *Time) MarshalJSON() ([]byte, error)` 这种形式，因为它会导致 Time 结构体丢失 MarshalJSON 接口的实现。
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (t *Time) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		t.Time = time.Time{}
		return nil
	}
	newTime, err := StrToTime(string(bytes.Trim(b, `"`)))
	if err != nil {
		return err
	}
	t.Time = newTime.Time
	return nil
}

// UnmarshalText 实现了 encoding.TextUnmarshaler 接口。
// 注意，它覆盖了 `time.Time` 同样的实现。
func (t *Time) UnmarshalText(data []byte) error {
	vTime := New(data)
	if vTime != nil {
		*t = *vTime
		return nil
	}
	return gerror.NewCodef(gcode.CodeInvalidParameter, `invalid time value: %s`, data)
}

// NoValidation 表示该结构体对象将不会被 gvalid 包进行验证。
func (t *Time) NoValidation() {}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (t *Time) DeepCopy() interface{} {
	if t == nil {
		return nil
	}
	return New(t.Time)
}
