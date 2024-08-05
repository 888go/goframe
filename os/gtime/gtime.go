// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package gtime 提供了测量和显示时间的功能。
// 
// 本包应该尽量减少与其他包的依赖。
// md5:34aae194a36b5e34
package gtime

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/text/gregex"
)

const (
		// 为常见的使用时长进行简短写入。 md5:368ca473af5d327a

	D  = 24 * time.Hour
	H  = time.Hour
	M  = time.Minute
	S  = time.Second
	MS = time.Millisecond
	US = time.Microsecond
	NS = time.Nanosecond

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
	timeRegexPattern1 = `(\d{4}[-/\.]\d{1,2}[-/\.]\d{1,2})[:\sT-]*(\d{0,2}:{0,1}\d{0,2}:{0,1}\d{0,2}){0,1}\.{0,1}(\d{0,9})([\sZ]{0,1})([\+-]{0,1})([:\d]*)`

	// 正则表达式（日期时间分隔符支持'-'、'/'和'.'）。
	// 例如：
	// 01-Nov-2018 11:50:28
	// 01/Nov/2018 11:50:28
	// 01.Nov.2018 11:50:28
	// 01.Nov.2018:11:50:28
	// md5:1f5fc72c5b6eb4f8
	timeRegexPattern2 = `(\d{1,2}[-/\.][A-Za-z]{3,}[-/\.]\d{4})[:\sT-]*(\d{0,2}:{0,1}\d{0,2}:{0,1}\d{0,2}){0,1}\.{0,1}(\d{0,9})([\sZ]{0,1})([\+-]{0,1})([:\d]*)`

	// 正则表达式3（时间）。
	// 例如：
	// 11:50:28
	// 11:50:28.897
	// md5:99204487b527a8dc
	timeRegexPattern3 = `(\d{2}):(\d{2}):(\d{2})\.{0,1}(\d{0,9})`
)

var (
	// 使用正则表达式解析日期时间字符串比time.ParseInLocation更高效。
	// md5:08fb8b42e551caf6
	timeRegex1 = regexp.MustCompile(timeRegexPattern1)
	timeRegex2 = regexp.MustCompile(timeRegexPattern2)
	timeRegex3 = regexp.MustCompile(timeRegexPattern3)

		// 月份英文到阿拉伯数字的映射。 md5:7730d474ee8e496c
	monthMap = map[string]int{
		"jan":       1,
		"feb":       2,
		"mar":       3,
		"apr":       4,
		"may":       5,
		"jun":       6,
		"jul":       7,
		"aug":       8,
		"sep":       9,
		"sept":      9,
		"oct":       10,
		"nov":       11,
		"dec":       12,
		"january":   1,
		"february":  2,
		"march":     3,
		"april":     4,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
	}
)

// Timestamp 获取并返回秒为单位的时间戳。 md5:f859aadcb9d86dfc
func Timestamp() int64 {
	return Now().Timestamp()
}

// TimestampMilli 获取并返回毫秒级的时间戳。 md5:59f90b1b1bf0bd83
func TimestampMilli() int64 {
	return Now().TimestampMilli()
}

// TimestampMicro 获取并返回以微秒为单位的时间戳。 md5:f773c1913603fb89
func TimestampMicro() int64 {
	return Now().TimestampMicro()
}

// TimestampNano 获取并以纳秒为单位返回时间戳。 md5:8b782ae92acea8e7
func TimestampNano() int64 {
	return Now().TimestampNano()
}

// TimestampStr 是一个方便的方法，它获取并返回时间戳（以秒为单位）的字符串形式。
// md5:f638769b91eb1dd5
func TimestampStr() string {
	return Now().TimestampStr()
}

// TimestampMilliStr是一个方便的方法，它获取并返回毫秒级的时间戳作为字符串。
// md5:cf293e6d5c9383d0
func TimestampMilliStr() string {
	return Now().TimestampMilliStr()
}

// TimestampMicroStr是一个方便的方法，它获取并返回微秒级别的时间戳作为字符串。
// md5:2930c4dc2c5feaae
func TimestampMicroStr() string {
	return Now().TimestampMicroStr()
}

// TimestampNanoStr 是一个便捷方法，用于获取并以字符串形式返回纳秒级的时间戳。
// md5:ff842fbe274c5052
func TimestampNanoStr() string {
	return Now().TimestampNanoStr()
}

// Date 返回当前日期的字符串，格式为 "2006-01-02"。 md5:e4ebaf573ffc4bd6
func Date() string {
	return time.Now().Format("2006-01-02")
}

// Datetime 返回当前日期时间的字符串格式，例如 "2006-01-02 15:04:05"。 md5:a1afd811808be0ca
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// ISO8601 返回当前日期时间，格式为 "2006-01-02T15:04:05-07:00"。 md5:ab2ef4b60100081e
func ISO8601() string {
	return time.Now().Format("2006-01-02T15:04:05-07:00")
}

// RFC822 返回当前日期时间的 RFC822 格式，例如 "Mon, 02 Jan 06 15:04 MST"。 md5:a6362395180caeda
func RFC822() string {
	return time.Now().Format("Mon, 02 Jan 06 15:04 MST")
}

// parseDateStr 将字符串解析为年份、月份和日期数字。 md5:697ca0661d5cecd9
func parseDateStr(s string) (year, month, day int) {
	array := strings.Split(s, "-")
	if len(array) < 3 {
		array = strings.Split(s, "/")
	}
	if len(array) < 3 {
		array = strings.Split(s, ".")
	}
	// Parsing failed.
	if len(array) < 3 {
		return
	}
		// 检查年份在头部还是尾部。 md5:33266655c259b475
	if utils.IsNumeric(array[1]) {
		year, _ = strconv.Atoi(array[0])
		month, _ = strconv.Atoi(array[1])
		day, _ = strconv.Atoi(array[2])
	} else {
		if v, ok := monthMap[strings.ToLower(array[1])]; ok {
			month = v
		} else {
			return
		}
		year, _ = strconv.Atoi(array[2])
		day, _ = strconv.Atoi(array[0])
	}
	return
}

// StrToTime 将字符串转换为 *Time 对象。它也支持时间戳字符串。
// 参数 `format` 是不必要的，用于指定转换格式，如 "Y-m-d H:i:s"。
// 如果提供了 `format`，它的行为与 StrToTimeFormat 函数相同。
// 如果没有提供 `format`，它将把字符串作为 "标准" 日期时间字符串进行转换。
// 注意，如果 `str` 中没有日期字符串，它将失败并返回错误。
// md5:5e4dd2ec67cb758d
func StrToTime(str string, format ...string) (*Time, error) {
	if str == "" {
		return &Time{wrapper{time.Time{}}}, nil
	}
	if len(format) > 0 {
		return StrToTimeFormat(str, format[0])
	}
	if isTimestampStr(str) {
		timestamp, _ := strconv.ParseInt(str, 10, 64)
		return NewFromTimeStamp(timestamp), nil
	}
	var (
		year, month, day     int
		hour, min, sec, nsec int
		match                []string
		local                = time.Local
	)
	if match = timeRegex1.FindStringSubmatch(str); len(match) > 0 && match[1] != "" {
		year, month, day = parseDateStr(match[1])
	} else if match = timeRegex2.FindStringSubmatch(str); len(match) > 0 && match[1] != "" {
		year, month, day = parseDateStr(match[1])
	} else if match = timeRegex3.FindStringSubmatch(str); len(match) > 0 && match[1] != "" {
		s := strings.ReplaceAll(match[2], ":", "")
		if len(s) < 6 {
			s += strings.Repeat("0", 6-len(s))
		}
		hour, _ = strconv.Atoi(match[1])
		min, _ = strconv.Atoi(match[2])
		sec, _ = strconv.Atoi(match[3])
		nsec, _ = strconv.Atoi(match[4])
		for i := 0; i < 9-len(match[4]); i++ {
			nsec *= 10
		}
		return NewFromTime(time.Date(0, time.Month(1), 1, hour, min, sec, nsec, local)), nil
	} else {
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `unsupported time converting for string "%s"`, str)
	}

	// Time
	if len(match[2]) > 0 {
		s := strings.ReplaceAll(match[2], ":", "")
		if len(s) < 6 {
			s += strings.Repeat("0", 6-len(s))
		}
		hour, _ = strconv.Atoi(s[0:2])
		min, _ = strconv.Atoi(s[2:4])
		sec, _ = strconv.Atoi(s[4:6])
	}
		// 纳秒，检查并执行位填充. md5:c54cea495b82ade6
	if len(match[3]) > 0 {
		nsec, _ = strconv.Atoi(match[3])
		for i := 0; i < 9-len(match[3]); i++ {
			nsec *= 10
		}
	}
	// 如果字符串中包含时区信息，
	// 然后进行时区转换，将时区转换为UTC。
	// md5:57a54806130bc3f5
	if match[4] != "" && match[6] == "" {
		match[6] = "000000"
	}
		// 如果字符串中有偏移量，那么它会首先处理这个偏移量。 md5:5a183f25f01ee951
	if match[6] != "" {
		zone := strings.ReplaceAll(match[6], ":", "")
		zone = strings.TrimLeft(zone, "+-")
		if len(zone) <= 6 {
			zone += strings.Repeat("0", 6-len(zone))
			h, _ := strconv.Atoi(zone[0:2])
			m, _ := strconv.Atoi(zone[2:4])
			s, _ := strconv.Atoi(zone[4:6])
			if h > 24 || m > 59 || s > 59 {
				return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid zone string "%s"`, match[6])
			}
			operation := match[5]
			if operation != "+" && operation != "-" {
				operation = "-"
			}
									// 比较给定的时区是否等于当前时区，. md5:7b9947fc7651a35e
			_, localOffset := time.Now().Zone()
			zoneOffset := h*3600 + m*60 + s
			if operation == "-" {
				zoneOffset = -zoneOffset
			}
			// Comparing in seconds.
			if localOffset != zoneOffset {
				local = time.FixedZone("", zoneOffset)
			}
		}
	}
	if month <= 0 || day <= 0 {
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid time string "%s"`, str)
	}
	return NewFromTime(time.Date(year, time.Month(month), day, hour, min, sec, nsec, local)), nil
}

// ConvertZone 将字符串格式的时间`strTime`从`fromZone`时区转换为`toZone`时区。
// 参数`fromZone`是可选的，默认情况下它代表当前所在的时区。
// md5:9c73950cf06cb368
func ConvertZone(strTime string, toZone string, fromZone ...string) (*Time, error) {
	t, err := StrToTime(strTime)
	if err != nil {
		return nil, err
	}
	var l *time.Location
	if len(fromZone) > 0 {
		if l, err = time.LoadLocation(fromZone[0]); err != nil {
			err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `time.LoadLocation failed for name "%s"`, fromZone[0])
			return nil, err
		} else {
			t.Time = time.Date(t.Year(), time.Month(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Time.Second(), t.Time.Nanosecond(), l)
		}
	}
	if l, err = time.LoadLocation(toZone); err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `time.LoadLocation failed for name "%s"`, toZone)
		return nil, err
	} else {
		return t.ToLocation(l), nil
	}
}

// StrToTimeFormat 函数将字符串 `str` 根据给定的格式 `format` 解析为 *Time 对象。
// 参数 `format` 的格式类似于 "Y-m-d H:i:s"。
// md5:0eb1a22261a21da1
func StrToTimeFormat(str string, format string) (*Time, error) {
	return StrToTimeLayout(str, formatToStdLayout(format))
}

// StrToTimeLayout 将字符串 `str` 解析为具有给定格式 `layout` 的 *Time 对象。参数 `layout` 应使用标准库中的格式，如 "2006-01-02 15:04:05"。
// md5:54702732831e3f2e
func StrToTimeLayout(str string, layout string) (*Time, error) {
	if t, err := time.ParseInLocation(layout, str, time.Local); err == nil {
		return NewFromTime(t), nil
	} else {
		return nil, gerror.WrapCodef(
			gcode.CodeInvalidParameter, err,
			`time.ParseInLocation failed for layout "%s" and value "%s"`,
			layout, str,
		)
	}
}

// ParseTimeFromContent 从内容字符串中提取时间信息，然后解析并返回一个 *Time 类型的对象。
// 如果内容中有多个时间字符串，它将返回第一个时间信息。
// 如果提供了 `format`，它只会检索并解析与之匹配的第一个时间信息。
// md5:37e6a9bec5011038
func ParseTimeFromContent(content string, format ...string) *Time {
	var (
		err   error
		match []string
	)
	if len(format) > 0 {
		for _, item := range format {
			match, err = gregex.MatchString(formatToRegexPattern(item), content)
			if err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
			if len(match) > 0 {
				return NewFromStrFormat(match[0], item)
			}
		}
	} else {
		if match = timeRegex1.FindStringSubmatch(content); len(match) >= 1 {
			return NewFromStr(strings.Trim(match[0], "./_- \n\r"))
		} else if match = timeRegex2.FindStringSubmatch(content); len(match) >= 1 {
			return NewFromStr(strings.Trim(match[0], "./_- \n\r"))
		} else if match = timeRegex3.FindStringSubmatch(content); len(match) >= 1 {
			return NewFromStr(strings.Trim(match[0], "./_- \n\r"))
		}
	}
	return nil
}

// ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h", "1d" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h", "d".
//
// Very note that it supports unit "d" more than function time.ParseDuration.
func ParseDuration(s string) (duration time.Duration, err error) {
	var (
		num int64
	)
	if utils.IsNumeric(s) {
		num, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `strconv.ParseInt failed for string "%s"`, s)
			return 0, err
		}
		return time.Duration(num), nil
	}
	match, err := gregex.MatchString(`^([\-\d]+)[dD](.*)$`, s)
	if err != nil {
		return 0, err
	}
	if len(match) == 3 {
		num, err = strconv.ParseInt(match[1], 10, 64)
		if err != nil {
			err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `strconv.ParseInt failed for string "%s"`, match[1])
			return 0, err
		}
		s = fmt.Sprintf(`%dh%s`, num*24, match[2])
		duration, err = time.ParseDuration(s)
		if err != nil {
			err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `time.ParseDuration failed for string "%s"`, s)
		}
		return
	}
	duration, err = time.ParseDuration(s)
	err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `time.ParseDuration failed for string "%s"`, s)
	return
}

// FuncCost 计算函数 `f` 的执行时间成本，以纳秒为单位。 md5:f6d4e0146ba246f1
func FuncCost(f func()) time.Duration {
	t := time.Now()
	f()
	return time.Since(t)
}

// isTimestampStr 检查并返回给定的字符串是否为时间戳字符串。 md5:cb1a63922b0758a1
func isTimestampStr(s string) bool {
	length := len(s)
	if length == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
