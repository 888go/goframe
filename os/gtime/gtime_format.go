// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtime

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
	
	"github.com/888go/goframe/text/gregex"
)

var (
	// 参考：http://php.net/manual/en/function.date.php
	formats = map[byte]string{
		'd': "02",                        // Day: 月份中的日期，以零填充的两位数字表示。例如：01至31。
		'D': "Mon",                       // Day: 一个表示星期的文本表示形式，由三个字母组成。例如：Mon（周一）至Sun（周日）。
		'w': "Monday",                    // Day: 表示星期几的数字表示。例如：0（代表星期日）到6（代表星期六）。
		'N': "Monday",                    // Day: ISO-8601格式表示的星期几的数字表示。例如：1（代表星期一）到7（代表星期日）。
		'j': "=j=02",                     // Day: 月份中的日期，不带前导零。例如：1 至 31。
		'S': "02",                        // Day: 对于月份中某一天的英文序数后缀，2个字符。例如：st（用于1），nd（用于2），rd（用于3）或th（用于其它）。与j配合使用效果更佳。
		'l': "Monday",                    // Day: 完整的星期几的文字表示。例如：星期日到星期六。
		'z': "",                          // Day: 一年中的天数（从0开始计数）。例如：0至365。
		'W': "",                          // Week: 根据ISO-8601标准的年份中的周数，以星期一开始计算。例如：42（一年中的第42周）。
		'F': "January",                   // Month: 一个完整表示月份的文本，例如一月或三月。例如：从一月到十二月。
		'm': "01",                        // Month: 月份的数字表示，前面带有零。例如：01至12。
		'M': "Jan",                       // Month: 月份的简短文本表示，共三个字母。例如：Jan 至 Dec。
		'n': "1",                         // Month: 月份的数字表示，不带前导零。例如：1 至 12。
		't': "",                          // Month: 给定月份的天数。例如：28至31。
		'Y': "2006",                      // Year：完整表示年份的数字形式，共4位数字。例如：1999或2003。
		'y': "06",                        // Year: 一个两位数表示的年份。例如：99 或 03。
		'a': "pm",                        // 时间：小写的上午和下午。例如：am 或 pm。
		'A': "PM",                        // 时间：大写的上午和下午。例如：AM 或 PM。
		'g': "3",                         // Time: 12小时制格式的小时，无前导零。例如：1至12。
		'G': "=G=15",                     // Time: 24小时制的小时格式，无前导零。例如：0至23。
		'h': "03",                        // Time: 12小时制格式的小时，前面补零。例如：01至12。
		'H': "15",                        // Time: 24小时格式的小时，前面补零。例如：00至23。
		'i': "04",                        // 时间：带有前导零的分钟。例如：00至59。
		's': "05",                        // Time: 带前导零的秒数。例如：00至59。
		'u': "=u=.000",                   // 时间：毫秒。例如：234，678。
		'U': "",                          // 时间：自 Unix Epoch（1970年1月1日 00:00:00 GMT）以来的秒数。
		'O': "-0700",                     // Zone: 与格林尼治标准时间（GMT）的时差，以小时为单位。例如：+0200。
		'P': "-07:00",                    // Zone: 与格林尼治标准时间（GMT）的时差，小时和分钟之间用冒号分隔。例如：+02:00。
		'T': "MST",                       // Zone: 时区缩写。例如：UTC、EST、MDT ...
		'c': "2006-01-02T15:04:05-07:00", // 格式：ISO 8601日期。例如：2004-02-12T15:19:21+00:00。
		'r': "Mon, 02 Jan 06 15:04 MST",  // 格式：RFC 2822格式化日期。例如：Thu, 21 Dec 2000 16:01:07 +0200。
// （注：RFC 2822是互联网消息格式的标准，定义了电子邮件等协议中日期和时间的表示方式。）
	}

	// 星期到数字的映射。
	weekMap = map[string]string{
		"Sunday":    "0",
		"Monday":    "1",
		"Tuesday":   "2",
		"Wednesday": "3",
		"Thursday":  "4",
		"Friday":    "5",
		"Saturday":  "6",
	}

	// 下面是每年非闰年的每月天数。
	dayOfMonth = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
)

// Format 函数使用自定义的 `format` 进行格式化，并返回格式化后的结果。
func (t *Time) Format(format string) string {
	if t == nil {
		return ""
	}
	runes := []rune(format)
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < len(runes); {
		switch runes[i] {
		case '\\':
			if i < len(runes)-1 {
				buffer.WriteRune(runes[i+1])
				i += 2
				continue
			} else {
				return buffer.String()
			}
		case 'W':
			buffer.WriteString(strconv.Itoa(t.WeeksOfYear()))
		case 'z':
			buffer.WriteString(strconv.Itoa(t.DayOfYear()))
		case 't':
			buffer.WriteString(strconv.Itoa(t.DaysInMonth()))
		case 'U':
			buffer.WriteString(strconv.FormatInt(t.Unix(), 10))
		default:
			if runes[i] > 255 {
				buffer.WriteRune(runes[i])
				break
			}
			if f, ok := formats[byte(runes[i])]; ok {
				result := t.Time.Format(f)
				// 特殊字符应该在这里处理。
				switch runes[i] {
				case 'j':
					for _, s := range []string{"=j=0", "=j="} {
						result = strings.ReplaceAll(result, s, "")
					}
					buffer.WriteString(result)
				case 'G':
					for _, s := range []string{"=G=0", "=G="} {
						result = strings.ReplaceAll(result, s, "")
					}
					buffer.WriteString(result)
				case 'u':
					buffer.WriteString(strings.ReplaceAll(result, "=u=.", ""))
				case 'w':
					buffer.WriteString(weekMap[result])
				case 'N':
					buffer.WriteString(strings.ReplaceAll(weekMap[result], "0", "7"))
				case 'S':
					buffer.WriteString(formatMonthDaySuffixMap(result))
				default:
					buffer.WriteString(result)
				}
			} else {
				buffer.WriteRune(runes[i])
			}
		}
		i++
	}
	return buffer.String()
}

// FormatNew 根据给定的自定义`format`格式化并返回一个新的Time对象。
func (t *Time) FormatNew(format string) *Time {
	if t == nil {
		return nil
	}
	return NewFromStr(t.Format(format))
}

// FormatTo 根据给定的自定义 `format` 格式化 `t`。
func (t *Time) FormatTo(format string) *Time {
	if t == nil {
		return nil
	}
	t.Time = NewFromStr(t.Format(format)).Time
	return t
}

// Layout 使用标准库的布局格式化时间，并返回格式化后的结果。
func (t *Time) Layout(layout string) string {
	if t == nil {
		return ""
	}
	return t.Time.Format(layout)
}

// LayoutNew 根据stdlib布局格式化时间，并返回一个新的Time对象。
func (t *Time) LayoutNew(layout string) *Time {
	if t == nil {
		return nil
	}
	newTime, err := StrToTimeLayout(t.Layout(layout), layout)
	if err != nil {
		panic(err)
	}
	return newTime
}

// LayoutTo 根据stdlib布局格式化`t`。
func (t *Time) LayoutTo(layout string) *Time {
	if t == nil {
		return nil
	}
	newTime, err := StrToTimeLayout(t.Layout(layout), layout)
	if err != nil {
		panic(err)
	}
	t.Time = newTime.Time
	return t
}

// IsLeapYear 检查给定的时间是否为闰年。
func (t *Time) IsLeapYear() bool {
	year := t.Year()
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

// DayOfYear 检查并返回一年中指定日期所在的天数位置。
func (t *Time) DayOfYear() int {
	var (
		day   = t.Day()
		month = t.Month()
	)
	if t.IsLeapYear() {
		if month > 2 {
			return dayOfMonth[month-1] + day
		}
		return dayOfMonth[month-1] + day - 1
	}
	return dayOfMonth[month-1] + day - 1
}

// DaysInMonth 返回当前月份的天数。
func (t *Time) DaysInMonth() int {
	switch t.Month() {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	}
	if t.IsLeapYear() {
		return 29
	}
	return 28
}

// WeeksOfYear 返回当前年份中的周点。
func (t *Time) WeeksOfYear() int {
	_, week := t.ISOWeek()
	return week
}

// formatToStdLayout 将自定义格式转换为stdlib布局。
func formatToStdLayout(format string) string {
	b := bytes.NewBuffer(nil)
	for i := 0; i < len(format); {
		switch format[i] {
		case '\\':
			if i < len(format)-1 {
				b.WriteByte(format[i+1])
				i += 2
				continue
			} else {
				return b.String()
			}

		default:
			if f, ok := formats[format[i]]; ok {
				// 处理特定字符。
				switch format[i] {
				case 'j':
					b.WriteString("2")
				case 'G':
					b.WriteString("15")
				case 'u':
					if i > 0 && format[i-1] == '.' {
						b.WriteString("000")
					} else {
						b.WriteString(".000")
					}

				default:
					b.WriteString(f)
				}
			} else {
				b.WriteByte(format[i])
			}
			i++
		}
	}
	return b.String()
}

// formatToRegexPattern 将自定义格式转换为其对应的正则表达式。
func formatToRegexPattern(format string) string {
	s := regexp.QuoteMeta(formatToStdLayout(format))
	s, _ = gregex.ReplaceString(`[0-9]`, `[0-9]`, s)
	s, _ = gregex.ReplaceString(`[A-Za-z]`, `[A-Za-z]`, s)
	s, _ = gregex.ReplaceString(`\s+`, `\s+`, s)
	return s
}

// formatMonthDaySuffixMap 返回当前日期的英文短词表示（如“日”，“月”）。
func formatMonthDaySuffixMap(day string) string {
	switch day {
	case "01", "21", "31":
		return "st"
	case "02", "22":
		return "nd"
	case "03", "23":
		return "rd"
	default:
		return "th"
	}
}
