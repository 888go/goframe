// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gtime 提供了用于测量和展示时间的功能。
//
// 本包应尽量减少与其他包的依赖关系。
package 时间类

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/888go/goframe/gtime/internal/intlog"
	"github.com/888go/goframe/gtime/internal/utils"
	"github.com/gogf/gf/v2/text/gregex"
)

const (
	// 对于常见使用时长提供简写方式。

	D  = 24 * time.Hour
	H  = time.Hour
	M  = time.Minute
	S  = time.Second
	MS = time.Millisecond
	US = time.Microsecond
	NS = time.Nanosecond

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
	timeRegexPattern1 = `(\d{4}[-/\.]\d{1,2}[-/\.]\d{1,2})[:\sT-]*(\d{0,2}:{0,1}\d{0,2}:{0,1}\d{0,2}){0,1}\.{0,1}(\d{0,9})([\sZ]{0,1})([\+-]{0,1})([:\d]*)`

// 正则表达式2（日期时间分隔符支持'-'、'/'和'.'）。
// 示例：
// 01-Nov-2018 11:50:28
// 01/Nov/2018 11:50:28
// 01.Nov.2018 11:50:28
// 01.Nov.2018:11:50:28
// 此正则表达式用于匹配格式类似于上述样例的日期时间字符串，其中日期和月份之间的分隔符可以是短横线、斜线或点，并且在某些情况下，时间部分也可以使用冒号与日期部分连接。
	timeRegexPattern2 = `(\d{1,2}[-/\.][A-Za-z]{3,}[-/\.]\d{4})[:\sT-]*(\d{0,2}:{0,1}\d{0,2}:{0,1}\d{0,2}){0,1}\.{0,1}(\d{0,9})([\sZ]{0,1})([\+-]{0,1})([:\d]*)`

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
	timeRegexPattern3 = `(\d{2}):(\d{2}):(\d{2})\.{0,1}(\d{0,9})`
)

var (
// 使用正则表达式解析日期时间字符串比使用time.ParseInLocation性能更高。
	timeRegex1 = regexp.MustCompile(timeRegexPattern1)
	timeRegex2 = regexp.MustCompile(timeRegexPattern2)
	timeRegex3 = regexp.MustCompile(timeRegexPattern3)

	// 月份单词到阿拉伯数字的映射。
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

// Timestamp 获取并返回以秒为单位的时间戳。
func X取时间戳秒() int64 {
	return X创建并按当前时间().X取时间戳秒()
}

// TimestampMilli 获取并返回毫秒级的时间戳。
func X取时间戳毫秒() int64 {
	return X创建并按当前时间().X取时间戳毫秒()
}

// TimestampMicro 获取并返回微秒级的时间戳。
func X取时间戳微秒() int64 {
	return X创建并按当前时间().X取时间戳微秒()
}

// TimestampNano 获取并返回纳秒级别的时间戳。
func X取时间戳纳秒() int64 {
	return X创建并按当前时间().X取时间戳纳秒()
}

// TimestampStr 是一个便捷方法，用于获取并返回以字符串形式表示的秒级时间戳。
func X取文本时间戳秒() string {
	return X创建并按当前时间().X取文本时间戳秒()
}

// TimestampMilliStr 是一个便捷方法，用于获取并返回以字符串形式表示的毫秒级时间戳。
func X取文本时间戳毫秒() string {
	return X创建并按当前时间().X取文本时间戳毫秒()
}

// TimestampMicroStr 是一个便捷方法，用于获取并返回微秒级的时间戳字符串。
func X取文本时间戳微秒() string {
	return X创建并按当前时间().X取文本时间戳微秒()
}

// TimestampNanoStr 是一个便捷方法，用于获取并返回纳秒级时间戳的字符串表示。
func X取文本时间戳纳秒() string {
	return X创建并按当前时间().X取文本时间戳纳秒()
}

// Date 返回当前日期的字符串，格式如 "2006-01-02"。
func Date() string {
	return time.Now().Format("2006-01-02")
}

// Datetime 返回当前日期时间的字符串，格式如 "2006-01-02 15:04:05"。
func X取当前日期时间() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// ISO8601 返回当前日期时间，格式为符合ISO 8601标准的字符串，如 "2006-01-02T15:04:05-07:00"。
func X取当前日期时间ISO8601() string {
	return time.Now().Format("2006-01-02T15:04:05-07:00")
}

// RFC822 返回当前日期时间的 RFC822 格式，例如 "Mon, 02 Jan 06 15:04 MST"。
func X取当前日期时间RFC822() string {
	return time.Now().Format("Mon, 02 Jan 06 15:04 MST")
}

// parseDateStr 将字符串解析为年、月和日的数字表示。
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
	// 检查年份在头部或尾部。
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

// StrToTime 将字符串转换为 *Time 类型对象。同时支持时间戳字符串。
// 参数 `format` 不是必需的，用于指定类似 "Y-m-d H:i:s" 的转换格式。
// 如果提供了 `format`，它的行为与函数 StrToTimeFormat 相同。
// 如果未提供 `format`，它将字符串按“标准”日期时间格式进行转换。
// 注意，如果 `str` 中不包含日期字符串，则转换失败并返回错误。
func X转换文本(文本时间 string, 格式 ...string) (*Time, error) {
	if 文本时间 == "" {
		return &Time{wrapper{time.Time{}}}, nil
	}
	if len(格式) > 0 {
		return StrToTimeFormat别名(文本时间, 格式[0])
	}
	if isTimestampStr(文本时间) {
		timestamp, _ := strconv.ParseInt(文本时间, 10, 64)
		return X创建并从时间戳(timestamp), nil
	}
	var (
		year, month, day     int
		hour, min, sec, nsec int
		match                []string
		local                = time.Local
	)
	if match = timeRegex1.FindStringSubmatch(文本时间); len(match) > 0 && match[1] != "" {
		year, month, day = parseDateStr(match[1])
	} else if match = timeRegex2.FindStringSubmatch(文本时间); len(match) > 0 && match[1] != "" {
		year, month, day = parseDateStr(match[1])
	} else if match = timeRegex3.FindStringSubmatch(文本时间); len(match) > 0 && match[1] != "" {
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
		return X创建并按Time(time.Date(0, time.Month(1), 1, hour, min, sec, nsec, local)), nil
	} else {
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `unsupported time converting for string "%s"`, 文本时间)
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
	// 纳秒，检查并执行位填充
	if len(match[3]) > 0 {
		nsec, _ = strconv.Atoi(match[3])
		for i := 0; i < 9-len(match[3]); i++ {
			nsec *= 10
		}
	}
// 如果字符串中包含时区信息，
// 则进行时区转换操作，将时间转换为 UTC 格式。
	if match[4] != "" && match[6] == "" {
		match[6] = "000000"
	}
	// 如果字符串中存在偏移量，则首先处理该偏移量。
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
// 检查给定时区是否等于当前时区，
// 如果不相等，则将其转换为UTC。
			_, localOffset := time.Now().Zone()
			// 按秒进行比较。
			if (h*3600+m*60+s) != localOffset ||
				(localOffset > 0 && operation == "-") ||
				(localOffset < 0 && operation == "+") {
				local = time.UTC
				// UTC conversion.
				switch operation {
				case "+":
					if h > 0 {
						hour -= h
					}
					if m > 0 {
						min -= m
					}
					if s > 0 {
						sec -= s
					}
				case "-":
					if h > 0 {
						hour += h
					}
					if m > 0 {
						min += m
					}
					if s > 0 {
						sec += s
					}
				}
			}
		}
	}
	if month <= 0 || day <= 0 {
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid time string "%s"`, 文本时间)
	}
	return X创建并按Time(time.Date(year, time.Month(month), day, hour, min, sec, nsec, local)), nil
}

// ConvertZone 将字符串 `strTime` 中的时间从 `fromZone` 转换为 `toZone`。
// 参数 `fromZone` 在默认情况下是不必要的，它表示当前时区。
func X转换时区(文本时间 string, 新时区 string, 旧时区 ...string) (*Time, error) {
	t, err := X转换文本(文本时间)
	if err != nil {
		return nil, err
	}
	var l *time.Location
	if len(旧时区) > 0 {
		if l, err = time.LoadLocation(旧时区[0]); err != nil {
			err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `time.LoadLocation failed for name "%s"`, 旧时区[0])
			return nil, err
		} else {
			t.Time = time.Date(t.Year(), time.Month(t.X取月份()), t.Day(), t.Hour(), t.Minute(), t.Time.Second(), t.Time.Nanosecond(), l)
		}
	}
	if l, err = time.LoadLocation(新时区); err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `time.LoadLocation failed for name "%s"`, 新时区)
		return nil, err
	} else {
		return t.X转换时区Location(l), nil
	}
}

// StrToTimeFormat将字符串`str`按照给定的格式`format`解析为*time.Time对象。
// 参数`format`如"Y-m-d H:i:s"。
func StrToTimeFormat别名(str string, format string) (*Time, error) {
	return X转换文本Layout(str, formatToStdLayout(format))
}

// StrToTimeLayout将字符串`str`按照给定的格式`layout`解析为*time.Time对象。
// 参数`layout`遵循标准库格式，如"2006-01-02 15:04:05"。
func X转换文本Layout(文本时间 string, 格式 string) (*Time, error) {
	if t, err := time.ParseInLocation(格式, 文本时间, time.Local); err == nil {
		return X创建并按Time(t), nil
	} else {
		return nil, gerror.WrapCodef(
			gcode.CodeInvalidParameter, err,
			`time.ParseInLocation failed for layout "%s" and value "%s"`,
			格式, 文本时间,
		)
	}
}

// ParseTimeFromContent 从内容字符串中检索时间信息，然后将其解析并作为 *Time 对象返回。
// 如果内容中存在多个时间字符串，则返回第一个时间信息。
// 如果传递了匹配的 `format` 参数，那么它将只检索并解析与第一个匹配格式对应的时间信息。
func X解析文本(文本 string, 格式 ...string) *Time {
	var (
		err   error
		match []string
	)
	if len(格式) > 0 {
		for _, item := range 格式 {
			match, err = gregex.MatchString(formatToRegexPattern(item), 文本)
			if err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
			if len(match) > 0 {
				return X创建并按给定格式文本(match[0], item)
			}
		}
	} else {
		if match = timeRegex1.FindStringSubmatch(文本); len(match) >= 1 {
			return X创建并从文本(strings.Trim(match[0], "./_- \n\r"))
		} else if match = timeRegex2.FindStringSubmatch(文本); len(match) >= 1 {
			return X创建并从文本(strings.Trim(match[0], "./_- \n\r"))
		} else if match = timeRegex3.FindStringSubmatch(文本); len(match) >= 1 {
			return X创建并从文本(strings.Trim(match[0], "./_- \n\r"))
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
func X文本取时长(文本 string) (纳秒 time.Duration, 错误 error) {
	var (
		num int64
	)
	if utils.IsNumeric(文本) {
		num, 错误 = strconv.ParseInt(文本, 10, 64)
		if 错误 != nil {
			错误 = gerror.WrapCodef(gcode.CodeInvalidParameter, 错误, `strconv.ParseInt failed for string "%s"`, 文本)
			return 0, 错误
		}
		return time.Duration(num), nil
	}
	match, 错误 := gregex.MatchString(`^([\-\d]+)[dD](.*)$`, 文本)
	if 错误 != nil {
		return 0, 错误
	}
	if len(match) == 3 {
		num, 错误 = strconv.ParseInt(match[1], 10, 64)
		if 错误 != nil {
			错误 = gerror.WrapCodef(gcode.CodeInvalidParameter, 错误, `strconv.ParseInt failed for string "%s"`, match[1])
			return 0, 错误
		}
		文本 = fmt.Sprintf(`%dh%s`, num*24, match[2])
		纳秒, 错误 = time.ParseDuration(文本)
		if 错误 != nil {
			错误 = gerror.WrapCodef(gcode.CodeInvalidParameter, 错误, `time.ParseDuration failed for string "%s"`, 文本)
		}
		return
	}
	纳秒, 错误 = time.ParseDuration(文本)
	错误 = gerror.WrapCodef(gcode.CodeInvalidParameter, 错误, `time.ParseDuration failed for string "%s"`, 文本)
	return
}

// FuncCost 计算函数 `f` 的执行耗时，单位为纳秒。
func X取函数执行时长(执行函数 func()) time.Duration {
	t := time.Now()
	执行函数()
	return time.Since(t)
}

// isTimestampStr 检查并返回给定字符串是否为时间戳字符串。
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
