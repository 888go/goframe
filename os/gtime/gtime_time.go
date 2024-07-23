// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gtime

import (
	"bytes"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// Time is a wrapper for time.Time for additional features.
type Time struct {
	wrapper
}

// iUnixNano is an interface definition commonly for custom time.Time wrapper.
type iUnixNano interface {
	UnixNano() int64
}

// New creates and returns a Time object with given parameter.
// The optional parameter is the time object which can be type of: time.Time/*time.Time, string or integer.
// New("2024-10-29")
// New(1390876568)
// New(t) // The t is type of time.Time.
// ff:创建
// param:参数
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

// Now creates and returns a time object of now.
// ff:创建并按当前时间
func Now() *Time {
	return &Time{
		wrapper{time.Now()},
	}
}

// NewFromTime creates and returns a Time object with given time.Time object.
// ff:创建并按Time
// t:
func NewFromTime(t time.Time) *Time {
	return &Time{
		wrapper{t},
	}
}

// NewFromStr creates and returns a Time object with given string.
// Note that it returns nil if there's error occurs.
// ff:创建并从文本
// str:文本时间
func NewFromStr(str string) *Time {
	if t, err := StrToTime(str); err == nil {
		return t
	}
	return nil
}

// NewFromStrFormat creates and returns a Time object with given string and
// custom format like: Y-m-d H:i:s.
// Note that it returns nil if there's error occurs.
// ff:创建并按给定格式文本
// str:文本时间
// format:格式
func NewFromStrFormat(str string, format string) *Time {
	if t, err := StrToTimeFormat(str, format); err == nil {
		return t
	}
	return nil
}

// NewFromStrLayout creates and returns a Time object with given string and
// stdlib layout like: 2006-01-02 15:04:05.
// Note that it returns nil if there's error occurs.
// ff:创建并按Layout格式文本
// str:文本时间
// layout:格式
func NewFromStrLayout(str string, layout string) *Time {
	if t, err := StrToTimeLayout(str, layout); err == nil {
		return t
	}
	return nil
}

// NewFromTimeStamp creates and returns a Time object with given timestamp,
// which can be in seconds to nanoseconds.
// ff:创建并从时间戳
// timestamp:时间戳
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

// Timestamp returns the timestamp in seconds.
// ff:取时间戳秒
// t:
func (t *Time) Timestamp() int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixNano() / 1e9
}

// TimestampMilli returns the timestamp in milliseconds.
// ff:取时间戳毫秒
// t:
func (t *Time) TimestampMilli() int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixNano() / 1e6
}

// TimestampMicro returns the timestamp in microseconds.
// ff:取时间戳微秒
// t:
func (t *Time) TimestampMicro() int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixNano() / 1e3
}

// TimestampNano returns the timestamp in nanoseconds.
// ff:取时间戳纳秒
// t:
func (t *Time) TimestampNano() int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixNano()
}

// TimestampStr is a convenience method which retrieves and returns
// the timestamp in seconds as string.
// ff:取文本时间戳秒
// t:
func (t *Time) TimestampStr() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.Timestamp(), 10)
}

// TimestampMilliStr is a convenience method which retrieves and returns
// the timestamp in milliseconds as string.
// ff:取文本时间戳毫秒
// t:
func (t *Time) TimestampMilliStr() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.TimestampMilli(), 10)
}

// TimestampMicroStr is a convenience method which retrieves and returns
// the timestamp in microseconds as string.
// ff:取文本时间戳微秒
// t:
func (t *Time) TimestampMicroStr() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.TimestampMicro(), 10)
}

// TimestampNanoStr is a convenience method which retrieves and returns
// the timestamp in nanoseconds as string.
// ff:取文本时间戳纳秒
// t:
func (t *Time) TimestampNanoStr() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.TimestampNano(), 10)
}

// Month returns the month of the year specified by t.
// ff:取月份
// t:
func (t *Time) Month() int {
	if t.IsZero() {
		return 0
	}
	return int(t.Time.Month())
}

// Second returns the second offset within the minute specified by t,
// in the range [0, 59].
// ff:取秒
// t:
func (t *Time) Second() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Second()
}

// Millisecond returns the millisecond offset within the second specified by t,
// in the range [0, 999].
// ff:取毫秒
// t:
func (t *Time) Millisecond() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Nanosecond() / 1e6
}

// Microsecond returns the microsecond offset within the second specified by t,
// in the range [0, 999999].
// ff:取微秒
// t:
func (t *Time) Microsecond() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Nanosecond() / 1e3
}

// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
// ff:取纳秒
// t:
func (t *Time) Nanosecond() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Nanosecond()
}

// String returns current time object as string.
// ff:
// t:
func (t *Time) String() string {
	if t.IsZero() {
		return ""
	}
	return t.wrapper.String()
}

// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
// ff:
// t:
func (t *Time) IsZero() bool {
	if t == nil {
		return true
	}
	return t.Time.IsZero()
}

// Clone returns a new Time object which is a clone of current time object.
// ff:取副本
// t:
func (t *Time) Clone() *Time {
	return New(t.Time)
}

// Add adds the duration to current time.
// ff:增加时长
// t:
// d:时长
func (t *Time) Add(d time.Duration) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Add(d)
	return newTime
}

// AddStr parses the given duration as string and adds it to current time.
// ff:增加文本时长
// t:
// duration:时长
func (t *Time) AddStr(duration string) (*Time, error) {
	if d, err := time.ParseDuration(duration); err != nil {
		err = gerror.Wrapf(err, `time.ParseDuration failed for string "%s"`, duration)
		return nil, err
	} else {
		return t.Add(d), nil
	}
}

// UTC converts current time to UTC timezone.
// ff:取UTC时区
// t:
func (t *Time) UTC() *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.UTC()
	return newTime
}

// ISO8601 formats the time as ISO8601 and returns it as string.
// ff:取文本时间ISO8601
// t:
func (t *Time) ISO8601() string {
	return t.Layout("2006-01-02T15:04:05-07:00")
}

// RFC822 formats the time as RFC822 and returns it as string.
// ff:取文本时间RFC822
// t:
func (t *Time) RFC822() string {
	return t.Layout("Mon, 02 Jan 06 15:04 MST")
}

// AddDate adds year, month and day to the time.
// ff:增加时间
// t:
// years:年
// months:月
// days:日
func (t *Time) AddDate(years int, months int, days int) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.AddDate(years, months, days)
	return newTime
}

// Round returns the result of rounding t to the nearest multiple of d (since the zero time).
// The rounding behavior for halfway values is to round up.
// If d <= 0, Round returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Round operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Round(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
// ff:向上舍入
// t:
// d:时长
func (t *Time) Round(d time.Duration) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Round(d)
	return newTime
}

// Truncate returns the result of rounding t down to a multiple of d (since the zero time).
// If d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Truncate operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Truncate(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
// ff:向下舍入
// t:
// d:时长
func (t *Time) Truncate(d time.Duration) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Truncate(d)
	return newTime
}

// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 CEST and 4:00 UTC are Equal.
// See the documentation on the Time type for the pitfalls of using == with
// Time values; most code should use Equal instead.
// ff:是否相等
// t:
// u:
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

// Before reports whether the time instant t is before u.
// ff:是否之前
// t:
// u:
func (t *Time) Before(u *Time) bool {
	return t.Time.Before(u.Time)
}

// After reports whether the time instant t is after u.
// ff:是否之后
// t:
// u:
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

// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
// To compute t-d for a duration d, use t.Add(-d).
// ff:取纳秒时长
// t:
// u:
func (t *Time) Sub(u *Time) time.Duration {
	if t == nil || u == nil {
		return 0
	}
	return t.Time.Sub(u.Time)
}

// StartOfMinute clones and returns a new time of which the seconds is set to 0.
// ff:取副本忽略秒
// t:
func (t *Time) StartOfMinute() *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Truncate(time.Minute)
	return newTime
}

// StartOfHour clones and returns a new time of which the hour, minutes and seconds are set to 0.
// ff:取副本忽略分钟秒
// t:
func (t *Time) StartOfHour() *Time {
	y, m, d := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(y, m, d, newTime.Time.Hour(), 0, 0, 0, newTime.Time.Location())
	return newTime
}

// StartOfDay clones and returns a new time which is the start of day, its time is set to 00:00:00.
// ff:取副本忽略小时分钟秒
// t:
func (t *Time) StartOfDay() *Time {
	y, m, d := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(y, m, d, 0, 0, 0, 0, newTime.Time.Location())
	return newTime
}

// StartOfWeek clones and returns a new time which is the first day of week and its time is set to
// ff:取副本周第一天
// t:
func (t *Time) StartOfWeek() *Time {
	weekday := int(t.Weekday())
	return t.StartOfDay().AddDate(0, 0, -weekday)
}

// StartOfMonth clones and returns a new time which is the first day of the month and its is set to
// ff:取副本月第一天
// t:
func (t *Time) StartOfMonth() *Time {
	y, m, _ := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(y, m, 1, 0, 0, 0, 0, newTime.Time.Location())
	return newTime
}

// StartOfQuarter clones and returns a new time which is the first day of the quarter and its time is set
// to 00:00:00.
// ff:取副本季度第一天
// t:
func (t *Time) StartOfQuarter() *Time {
	month := t.StartOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

// StartOfHalf clones and returns a new time which is the first day of the half year and its time is set
// to 00:00:00.
// ff:取副本半年第一天
// t:
func (t *Time) StartOfHalf() *Time {
	month := t.StartOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return month.AddDate(0, -offset, 0)
}

// StartOfYear clones and returns a new time which is the first day of the year and its time is set to
// ff:取副本年第一天
// t:
func (t *Time) StartOfYear() *Time {
	y, _, _ := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(y, time.January, 1, 0, 0, 0, 0, newTime.Time.Location())
	return newTime
}

// getPrecisionDelta returns the precision parameter for time calculation depending on `withNanoPrecision` option.
func getPrecisionDelta(withNanoPrecision ...bool) time.Duration {
	if len(withNanoPrecision) > 0 && withNanoPrecision[0] {
		return time.Nanosecond
	}
	return time.Second
}

// EndOfMinute clones and returns a new time of which the seconds is set to 59.
// ff:取副本59秒
// t:
// withNanoPrecision:纳秒精度
func (t *Time) EndOfMinute(withNanoPrecision ...bool) *Time {
	return t.StartOfMinute().Add(time.Minute - getPrecisionDelta(withNanoPrecision...))
}

// EndOfHour clones and returns a new time of which the minutes and seconds are both set to 59.
// ff:取副本59分59秒
// t:
// withNanoPrecision:纳秒精度
func (t *Time) EndOfHour(withNanoPrecision ...bool) *Time {
	return t.StartOfHour().Add(time.Hour - getPrecisionDelta(withNanoPrecision...))
}

// EndOfDay clones and returns a new time which is the end of day the and its time is set to 23:59:59.
// ff:取副本23点59分59秒
// t:
// withNanoPrecision:纳秒精度
func (t *Time) EndOfDay(withNanoPrecision ...bool) *Time {
	y, m, d := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(
		y, m, d, 23, 59, 59, int(time.Second-getPrecisionDelta(withNanoPrecision...)), newTime.Time.Location(),
	)
	return newTime
}

// EndOfWeek clones and returns a new time which is the end of week and its time is set to 23:59:59.
// ff:取副本周末23点59分59秒
// t:
// withNanoPrecision:纳秒精度
func (t *Time) EndOfWeek(withNanoPrecision ...bool) *Time {
	return t.StartOfWeek().AddDate(0, 0, 7).Add(-getPrecisionDelta(withNanoPrecision...))
}

// EndOfMonth clones and returns a new time which is the end of the month and its time is set to 23:59:59.
// ff:取副本月末23点59分59秒
// t:
// withNanoPrecision:纳秒精度
func (t *Time) EndOfMonth(withNanoPrecision ...bool) *Time {
	return t.StartOfMonth().AddDate(0, 1, 0).Add(-getPrecisionDelta(withNanoPrecision...))
}

// EndOfQuarter clones and returns a new time which is end of the quarter and its time is set to 23:59:59.
// ff:取副本季末23点59分59秒
// t:
// withNanoPrecision:纳秒精度
func (t *Time) EndOfQuarter(withNanoPrecision ...bool) *Time {
	return t.StartOfQuarter().AddDate(0, 3, 0).Add(-getPrecisionDelta(withNanoPrecision...))
}

// EndOfHalf clones and returns a new time which is the end of the half year and its time is set to 23:59:59.
// ff:取副本半年末23点59分59秒
// t:
// withNanoPrecision:纳秒精度
func (t *Time) EndOfHalf(withNanoPrecision ...bool) *Time {
	return t.StartOfHalf().AddDate(0, 6, 0).Add(-getPrecisionDelta(withNanoPrecision...))
}

// EndOfYear clones and returns a new time which is the end of the year and its time is set to 23:59:59.
// ff:取副本年末23点59分59秒
// t:
// withNanoPrecision:纳秒精度
func (t *Time) EndOfYear(withNanoPrecision ...bool) *Time {
	return t.StartOfYear().AddDate(1, 0, 0).Add(-getPrecisionDelta(withNanoPrecision...))
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// Note that, DO NOT use `(t *Time) MarshalJSON() ([]byte, error)` as it looses interface
// implement of `MarshalJSON` for struct of Time.
// ff:
// t:
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
// ff:
// t:
// b:
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

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// Note that it overwrites the same implementer of `time.Time`.
// ff:
// t:
// data:
func (t *Time) UnmarshalText(data []byte) error {
	vTime := New(data)
	if vTime != nil {
		*t = *vTime
		return nil
	}
	return gerror.NewCodef(gcode.CodeInvalidParameter, `invalid time value: %s`, data)
}

// NoValidation marks this struct object will not be validated by package gvalid.
// ff:
// t:
func (t *Time) NoValidation() {}

// DeepCopy implements interface for deep copy of current type.
// ff:
// t:
func (t *Time) DeepCopy() interface{} {
	if t == nil {
		return nil
	}
	return New(t.Time)
}
