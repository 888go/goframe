// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 时间类

import (
	"bytes"
	"strconv"
	"time"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

// Time 是一个包装了 time.Time 的结构，用于添加额外的功能。 md5:96d9b7cb3af14206
type Time struct {
	wrapper
}

// iUnixNano 是一个常用的自定义 time.Time 包装器的接口定义。 md5:5c0387efec09a99b
type iUnixNano interface {
	UnixNano() int64
}

// X创建 函数创建并返回一个 Time 对象，使用给定的参数。可选参数是一个时间对象，可以是以下类型：time.Time/*time.Time、字符串或整数。
// 例子：
// X创建("2024-10-29")
// X创建(1390876568)
// X创建(t) // t 是 time.Time 类型。
// md5:6951100c014c4ba9
func X创建(参数 ...interface{}) *Time {
	if len(参数) > 0 {
		switch r := 参数[0].(type) {
		case time.Time:
			return X创建并按Time(r)
		case *time.Time:
			return X创建并按Time(*r)

		case Time:
			return &r

		case *Time:
			return r

		case string:
			if len(参数) > 1 {
				switch t := 参数[1].(type) {
				case string:
					return X创建并按给定格式文本(r, t)
				case []byte:
					return X创建并按给定格式文本(r, string(t))
				}
			}
			return X创建并从文本(r)

		case []byte:
			if len(参数) > 1 {
				switch t := 参数[1].(type) {
				case string:
					return X创建并按给定格式文本(string(r), t)
				case []byte:
					return X创建并按给定格式文本(string(r), string(t))
				}
			}
			return X创建并从文本(string(r))

		case int:
			return X创建并从时间戳(int64(r))

		case int64:
			return X创建并从时间戳(r)

		default:
			if v, ok := r.(iUnixNano); ok {
				return X创建并从时间戳(v.UnixNano())
			}
		}
	}
	return &Time{
		wrapper{time.Time{}},
	}
}

// 现在创建并返回一个表示当前时间的对象。 md5:1cfc3114797b1f98
func X创建并按当前时间() *Time {
	return &Time{
		wrapper{time.Now()},
	}
}

// X创建并按Time 根据给定的time.Time对象创建并返回一个Time对象。 md5:e1cf178ea024f53b
func X创建并按Time(t time.Time) *Time {
	return &Time{
		wrapper{t},
	}
}

// X创建并从文本 根据给定的字符串创建并返回一个 Time 对象。
// 注意，如果发生错误，它将返回 nil。
// md5:4687b38a27582a12
func X创建并从文本(文本时间 string) *Time {
	if t, err := X转换文本(文本时间); err == nil {
		return t
	}
	return nil
}

// X创建并按给定格式文本 通过给定的字符串和自定义格式（如：Y-m-d H:i:s）创建并返回一个Time对象。
// 注意，如果发生错误，它将返回nil。
// md5:ed9966a0a8156f1d
func X创建并按给定格式文本(文本时间 string, 格式 string) *Time {
	if t, err := StrToTimeFormat别名(文本时间, 格式); err == nil {
		return t
	}
	return nil
}

// X创建并按Layout格式文本 根据给定的字符串和标准库格式（如：2006-01-02 15:04:05）创建并返回一个Time对象。
// 注意，如果出现错误，它将返回nil。
// md5:027f4d0876baa1a8
func X创建并按Layout格式文本(文本时间 string, 格式 string) *Time {
	if t, err := X转换文本Layout(文本时间, 格式); err == nil {
		return t
	}
	return nil
}

// X创建并从时间戳 根据给定的时间戳创建并返回一个 Time 对象，
// 该时间戳可以是秒到纳秒的精度。
// 例如：1600443866 和 1600443866199266000 都被视为有效的时间戳数值。
// md5:6a84edd691c97a4f
func X创建并从时间戳(时间戳 int64) *Time {
	if 时间戳 == 0 {
		return &Time{}
	}
	var sec, nano int64
	if 时间戳 > 1e9 {
		for 时间戳 < 1e18 {
			时间戳 *= 10
		}
		sec = 时间戳 / 1e9
		nano = 时间戳 % 1e9
	} else {
		sec = 时间戳
	}
	return &Time{
		wrapper{time.Unix(sec, nano)},
	}
}

// X取时间戳秒 返回时间戳，以秒为单位。 md5:52f3b8b0088c2fab
func (t *Time) X取时间戳秒() int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixNano() / 1e9
}

// X取时间戳毫秒 返回毫秒级的时间戳。 md5:945db1871b08c49f
func (t *Time) X取时间戳毫秒() int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixNano() / 1e6
}

// X取时间戳微秒 返回以微秒为单位的时间戳。 md5:20da1d303fcad848
func (t *Time) X取时间戳微秒() int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixNano() / 1e3
}

// X取时间戳纳秒 返回以纳秒为单位的时间戳。 md5:93016ce343f59007
func (t *Time) X取时间戳纳秒() int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixNano()
}

// X取文本时间戳秒 是一个方便的方法，它获取并返回时间戳（以秒为单位）的字符串形式。
// md5:f638769b91eb1dd5
func (t *Time) X取文本时间戳秒() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.X取时间戳秒(), 10)
}

// X取文本时间戳毫秒是一个方便的方法，它获取并返回毫秒级的时间戳作为字符串。
// md5:cf293e6d5c9383d0
func (t *Time) X取文本时间戳毫秒() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.X取时间戳毫秒(), 10)
}

// X取文本时间戳微秒是一个方便的方法，它获取并返回微秒级别的时间戳作为字符串。
// md5:2930c4dc2c5feaae
func (t *Time) X取文本时间戳微秒() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.X取时间戳微秒(), 10)
}

// X取文本时间戳纳秒 是一个便捷方法，用于获取并以字符串形式返回纳秒级的时间戳。
// md5:ff842fbe274c5052
func (t *Time) X取文本时间戳纳秒() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.X取时间戳纳秒(), 10)
}

// X取月份 返回指定时间t的月份。 md5:84f113a801a5eb29
func (t *Time) X取月份() int {
	if t.IsZero() {
		return 0
	}
	return int(t.Time.Month())
}

// X取秒返回由t指定的分钟内的第二个偏移量，范围在[0, 59]之间。
// md5:5666ae5cbf21989d
func (t *Time) X取秒() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Second()
}

// X取毫秒 返回给定时间 t 所在秒内的毫秒偏移，范围为 [0, 999]。
// md5:8bb4c372dc3ada79
func (t *Time) X取毫秒() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Nanosecond() / 1e6
}

// X取微秒 返回 t 指定的秒内微秒偏移量，范围为 [0, 999999]。
// md5:cb28fad241f60582
func (t *Time) X取微秒() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Nanosecond() / 1e3
}

// X取纳秒 返回 t 所指定秒内的纳秒偏移量，范围为 [0, 999999999]。
// md5:c1dcd3dd99062cf7
func (t *Time) X取纳秒() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Nanosecond()
}

// String 返回当前时间对象作为字符串。 md5:4f5a1f3896ca049d
func (t *Time) String() string {
	if t.IsZero() {
		return ""
	}
	return t.wrapper.String()
}

// IsZero报告是否`t`表示零时间点，即UTC时间的1970年1月1日00:00:00。
// md5:4e2b46d4fa63a878
func (t *Time) IsZero() bool {
	if t == nil {
		return true
	}
	return t.Time.IsZero()
}

// X取副本 返回一个与当前时间对象相克隆的新Time对象。 md5:8a0848cce3c64ef5
func (t *Time) X取副本() *Time {
	return X创建(t.Time)
}

// X增加时长 将持续时间添加到当前时间。 md5:8a845aeaaa064af4
func (t *Time) X增加时长(时长 time.Duration) *Time {
	newTime := t.X取副本()
	newTime.Time = newTime.Time.Add(时长)
	return newTime
}

// X增加文本时长解析给定的字符串持续时间，并将其添加到当前时间。 md5:3c2278027933d90f
func (t *Time) X增加文本时长(时长 string) (*Time, error) {
	if d, err := time.ParseDuration(时长); err != nil {
		err = gerror.X多层错误并格式化(err, `time.ParseDuration failed for string "%s"`, 时长)
		return nil, err
	} else {
		return t.X增加时长(d), nil
	}
}

// X取UTC时区 将当前时间转换为X取UTC时区时区。 md5:5067cfa0c7c94f95
func (t *Time) X取UTC时区() *Time {
	newTime := t.X取副本()
	newTime.Time = newTime.Time.UTC()
	return newTime
}

// X取文本时间ISO8601将时间格式化为X取文本时间ISO8601标准格式，并以字符串形式返回。 md5:6ddd62f8570c26f4
func (t *Time) X取文本时间ISO8601() string {
	return t.X取Layout格式文本("2006-01-02T15:04:05-07:00")
}

// X取文本时间RFC822 根据 X取文本时间RFC822 格式将时间转换为字符串并返回。 md5:1b6d66ac42df19de
func (t *Time) X取文本时间RFC822() string {
	return t.X取Layout格式文本("Mon, 02 Jan 06 15:04 MST")
}

// X增加时间 向时间添加年、月和日。 md5:643cfbc24c5bd938
func (t *Time) X增加时间(年 int, 月 int, 日 int) *Time {
	newTime := t.X取副本()
	newTime.Time = newTime.Time.AddDate(年, 月, 日)
	return newTime
}

// X向上舍入 返回将 t 四舍五入到 d 的倍数的结果（从零时间开始）。对于半等值，四舍五入行为向上取整。
// 如果 d 小于等于 0，X向上舍入 会返回 t 并移除任何单调时钟读数，但保持不变。
//
// X向上舍入 以绝对的自零时间以来的时间段进行操作；它不处理时间的呈现形式。因此，X向上舍入(Hour) 可能返回一个非零分钟的时间，具体取决于时间的 Location。
// md5:b2557220790fc058
func (t *Time) X向上舍入(时长 time.Duration) *Time {
	newTime := t.X取副本()
	newTime.Time = newTime.Time.Round(时长)
	return newTime
}

// X向下舍入 返回将时间t向下舍入到d的倍数的结果（从零时间开始）。
// 如果d<=0，X向下舍入会返回t，但去除任何单调时钟读数，否则保持不变。
//
// X向下舍入是基于时间从零时间点起的绝对持续时间来进行操作的；
// 它并不作用于时间的展示形式。因此，X向下舍入(Hour)可能返回一个分钟数非零的时间，
// 这取决于该时间的位置信息（Location）。
// md5:f72e0e00b245e691
func (t *Time) X向下舍入(时长 time.Duration) *Time {
	newTime := t.X取副本()
	newTime.Time = newTime.Time.Truncate(时长)
	return newTime
}

// X是否相等 函数报告 t 和 u 是否表示相同的时刻。
// 即使两个时间在不同的时区，它们也可以相等。
// 例如，CEST 的 6:00 +0200 和 UTC 的 4:00 是相等的。
// 查看 Time 类型的文档，了解使用 == 操作符比较时间值时可能遇到的问题；
// 大多数代码应使用 X是否相等 而非 ==。
// md5:a28e147d11d5fe0f
func (t *Time) X是否相等(u *Time) bool {
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

// X是否之前 返回时间点 t 是否在 u 之前。 md5:36690a50c1e8d9d4
func (t *Time) X是否之前(u *Time) bool {
	return t.Time.Before(u.Time)
}

// X是否之后 判断时间点t是否在u之后。 md5:750eca8bb04e1a25
func (t *Time) X是否之后(u *Time) bool {
	switch {
	case t == nil:
		return false
	case t != nil && u == nil:
		return true
	default:
		return t.Time.After(u.Time)
	}
}

// X取纳秒时长 返回持续时间 t-u。如果结果超过了能存储在 Duration 类型中的最大（或最小）
// 值，那么将返回最大（或最小）的持续时间。
// 要计算 t-d（其中 d 为一个持续时间），请使用 t.Add(-d)。
// md5:c975e5087c03d3b9
func (t *Time) X取纳秒时长(u *Time) time.Duration {
	if t == nil || u == nil {
		return 0
	}
	return t.Time.Sub(u.Time)
}

// X取副本忽略秒 克隆并返回一个新的时间，其中秒数被设置为0。 md5:dc10ea1284a17280
func (t *Time) X取副本忽略秒() *Time {
	newTime := t.X取副本()
	newTime.Time = newTime.Time.Truncate(time.Minute)
	return newTime
}

// X取副本忽略分钟秒克隆并返回一个新的时间，其中小时、分钟和秒设置为0。 md5:d52e77457a157871
func (t *Time) X取副本忽略分钟秒() *Time {
	y, m, d := t.Date()
	newTime := t.X取副本()
	newTime.Time = time.Date(y, m, d, newTime.Time.Hour(), 0, 0, 0, newTime.Time.Location())
	return newTime
}

// X取副本忽略小时分钟秒克隆并返回一个新的时间，它是新的一天的开始，其时间被设置为00:00:00。 md5:a9262cc6eafed6da
func (t *Time) X取副本忽略小时分钟秒() *Time {
	y, m, d := t.Date()
	newTime := t.X取副本()
	newTime.Time = time.Date(y, m, d, 0, 0, 0, 0, newTime.Time.Location())
	return newTime
}

// X取副本周第一天 克隆并返回一个新的时间，该时间为一周的第一天，其时间设置为00:00:00。
// md5:46c7f050c7f59e0a
func (t *Time) X取副本周第一天() *Time {
	weekday := int(t.Weekday())
	return t.X取副本忽略小时分钟秒().X增加时间(0, 0, -weekday)
}

// X取副本月第一天 创建并返回一个新的时间，该时间是月份的第一天，并且时间设置为 00:00:00
// md5:3de8c28f482566bb
func (t *Time) X取副本月第一天() *Time {
	y, m, _ := t.Date()
	newTime := t.X取副本()
	newTime.Time = time.Date(y, m, 1, 0, 0, 0, 0, newTime.Time.Location())
	return newTime
}

// X取副本季度第一天克隆并返回一个新的时间，它是季度的第一天，时间被设置为00:00:00。
// md5:814969ee5c648fb0
func (t *Time) X取副本季度第一天() *Time {
	month := t.X取副本月第一天()
	offset := (int(month.X取月份()) - 1) % 3
	return month.X增加时间(0, -offset, 0)
}

// X取副本半年第一天克隆并返回一个新的时间，它是半年的第一天，时间被设置为00:00:00。
// md5:5b53c4e328da312e
func (t *Time) X取副本半年第一天() *Time {
	month := t.X取副本月第一天()
	offset := (int(month.X取月份()) - 1) % 6
	return month.X增加时间(0, -offset, 0)
}

// X取副本年第一天 克隆并返回一个新的时间，该时间为一年中的第一天，其时间设置为00:00:00。
// md5:7bfbc3ec2e634ff2
func (t *Time) X取副本年第一天() *Time {
	y, _, _ := t.Date()
	newTime := t.X取副本()
	newTime.Time = time.Date(y, time.January, 1, 0, 0, 0, 0, newTime.Time.Location())
	return newTime
}

// getPrecisionDelta 根据`withNanoPrecision`选项返回时间计算的精度参数。 md5:8bcdeaaf0e87d398
func getPrecisionDelta(withNanoPrecision ...bool) time.Duration {
	if len(withNanoPrecision) > 0 && withNanoPrecision[0] {
		return time.Nanosecond
	}
	return time.Second
}

// X取副本59秒克隆并返回一个新的时间，其中秒设置为59。 md5:f1cc1512e831d5fa
func (t *Time) X取副本59秒(纳秒精度 ...bool) *Time {
	return t.X取副本忽略秒().X增加时长(time.Minute - getPrecisionDelta(纳秒精度...))
}

// X取副本59分59秒克隆并返回一个新的时间，其中分钟和秒都设置为59。 md5:ea49434e1e5b1bbb
func (t *Time) X取副本59分59秒(纳秒精度 ...bool) *Time {
	return t.X取副本忽略分钟秒().X增加时长(time.Hour - getPrecisionDelta(纳秒精度...))
}

// X取副本23点59分59秒 克隆并返回一个新的时间，该时间设置为当天的结束，即时间部分被设置为 23:59:59。 md5:77a284f48ab6cac4
func (t *Time) X取副本23点59分59秒(纳秒精度 ...bool) *Time {
	y, m, d := t.Date()
	newTime := t.X取副本()
	newTime.Time = time.Date(
		y, m, d, 23, 59, 59, int(time.Second-getPrecisionDelta(纳秒精度...)), newTime.Time.Location(),
	)
	return newTime
}

// X取副本周末23点59分59秒 创建并返回一个新的时间，该时间表示一周的结束，并将其时间设置为23:59:59。 md5:eb899f421cfb25b4
func (t *Time) X取副本周末23点59分59秒(纳秒精度 ...bool) *Time {
	return t.X取副本周第一天().X增加时间(0, 0, 7).X增加时长(-getPrecisionDelta(纳秒精度...))
}

// X取副本月末23点59分59秒克隆并返回一个新的时间，它是当月的结束，时间设置为23:59:59。 md5:6c2259b48332a891
func (t *Time) X取副本月末23点59分59秒(纳秒精度 ...bool) *Time {
	return t.X取副本月第一天().X增加时间(0, 1, 0).X增加时长(-getPrecisionDelta(纳秒精度...))
}

// X取副本季末23点59分59秒克隆并返回一个新的时间，它是季度结束，其时间设置为23:59:59。 md5:c2e7dca6753c6e99
func (t *Time) X取副本季末23点59分59秒(纳秒精度 ...bool) *Time {
	return t.X取副本季度第一天().X增加时间(0, 3, 0).X增加时长(-getPrecisionDelta(纳秒精度...))
}

// X取副本半年末23点59分59秒 克隆并返回一个新的时间，该时间设置为半年的结束时刻，具体时间为 23:59:59。 md5:2f3662f357ee5f6d
func (t *Time) X取副本半年末23点59分59秒(纳秒精度 ...bool) *Time {
	return t.X取副本半年第一天().X增加时间(0, 6, 0).X增加时长(-getPrecisionDelta(纳秒精度...))
}

// X取副本年末23点59分59秒 克隆并返回一个新的时间，该时间是当年的年末，时间设置为23:59:59。 md5:33b38d1d0badf6ad
func (t *Time) X取副本年末23点59分59秒(纳秒精度 ...bool) *Time {
	return t.X取副本年第一天().X增加时间(1, 0, 0).X增加时长(-getPrecisionDelta(纳秒精度...))
}

// MarshalJSON 实现了 json.Marshal 接口的 MarshalJSON 方法。注意，不要使用 `(t *Time) MarshalJSON() ([]byte, error)`，因为它会丢失 Time 结构体的 MarshalJSON 接口实现。
// md5:daef718235a856ce
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (t *Time) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		t.Time = time.Time{}
		return nil
	}
	newTime, err := X转换文本(string(bytes.Trim(b, `"`)))
	if err != nil {
		return err
	}
	t.Time = newTime.Time
	return nil
}

// UnmarshalText实现了encoding.TextUnmarshaler接口。
// 注意，它会覆盖与`time.Time`相同的实现者。
// md5:8aa957653e42443a
func (t *Time) UnmarshalText(data []byte) error {
	vTime := X创建(data)
	if vTime != nil {
		*t = *vTime
		return nil
	}
	return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `invalid time value: %s`, data)
}

// NoValidation 标记这个结构体对象将不会被 gvalid 包进行验证。 md5:5241ee7a51fb1912
func (t *Time) NoValidation() {}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (t *Time) DeepCopy() interface{} {
	if t == nil {
		return nil
	}
	return X创建(t.Time)
}
