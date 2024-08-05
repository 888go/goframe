// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtime

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/text/gregex"
)

var (
			//php.net/manual/en/function.date.php. md5:8a7c43207206ddc7
	formats = map[byte]string{
		'd': "02",                        // Day: 月份中的天数，两位数字，前面带零。例如：01 到 31。 md5:172194d0b1814362
		'D': "Mon",                       // Day：一个文本表示一天，三个字母。例如：周一到周日。 md5:675cfde285bf0c68
		'w': "Monday",                    // Day：表示星期几的数字表示。例如：0（代表周日）到 6（代表周六）。 md5:587051ffc46cb88f
		'N': "Monday",                    // Day: 一周中ISO-8601数字表示的日期。例如：1（代表星期一）到7（代表星期日）。 md5:1d2c4841413ee16b
		'j': "=j=02",                     // Day：月份中的天数，不带前导零。例如：1 到 31。 md5:1b22606749433e52
		'S': "02",                        // Day：用于表示月份中某一天的英文序数后缀，长度为2个字符。例如：st（第一），nd（第二），rd（第21, 31天）或th（其他）。与j配合使用效果良好。 md5:32738b7522d652ac
		'l': "Monday",                    // Day：一周中每一天的完整文本表示。例如：从周日到周六。 md5:b88c9e68e02dfc0f
		'z': "",                          // Day: 一年中的第几天（从0开始）。例如：0到365。 md5:f49ddc25d665aca9
		'W': "",                          // Week：ISO-8601年的周数，从星期一开始计数。例如：42（一年中的第42周）。 md5:b80d760d2e0ae49e
		'F': "January",                   // 月份：一个完整的文本表示的月份，如一月或三月。例如：一月至十二月。 md5:eeed261e583df1e4
		'm': "01",                        // Month：月份的数字表示，前导零。例如：01 到 12。 md5:dbd3bd5aeba7d2a6
		'M': "Jan",                       // Month: 月份的简写，三个字母表示。例如：Jan 到 Dec。 md5:7421ba4bd6a330f2
		'n': "1",                         // 月份：月份的数字表示，不带前导零。例如：1 到 12。 md5:a21d42dda9aa0303
		't': "",                          // 月份：给定月份的天数。例如：28到31。 md5:015a4430f81eb6af
		'Y': "2006",                      // Year: 一年的完整数字表示，共4位数。例如：1999或2003。 md5:75ba896183b8c9b6
		'y': "06",                        // Year: 以两位数字表示的年份。例如：99或03。 md5:8297f526cc8cf6f6
		'a': "pm",                        // 时间：小写的上午 (ante meridiem) 和下午 (post meridiem)。例如：am 或 pm。 md5:6a2ba15e8b7c7334
		'A': "PM",                        // 时间：大写的上午和下午。例如：AM或PM。 md5:1623916a72adc605
		'g': "3",                         // 时间：12小时制的小时表示，不带前导零。例如：1到12。 md5:e16c4eca54a59bc7
		'G': "=G=15",                     // 时间：24小时制格式的小时，不带前导零。例如：0 到 23。 md5:bded74e930df423a
		'h': "03",                        // 时间：24小时制的小时，带前导零。例如：01到12。 md5:34580e7a214880fe
		'H': "15",                        // 时间：24小时格式的小时，带前导零。例如：00到23。 md5:acc91e68d6dad4bb
		'i': "04",                        // 时间：带有前导零的分钟数。例如：00到59。 md5:b5517938304d47d3
		's': "05",                        // 时间：带有前导零的秒数。例如：00 到 59。 md5:dd437c5def368e39
		'u': "=u=.000",                   // 时间：毫秒。例如：234，678。 md5:c77bfea15fee7d89
		'U': "",                          // 时间：从Unix纪元（1970年1月1日GMT 00:00:00）开始的秒数。 md5:4fbca20050391fc7
		'O': "-0700",                     // 时区：与格林威治标准时间（GMT）的小时差。例如：+0200。 md5:27b1435bd4ed817e
		'P': "-07:00",                    // Zone: 与格林威治时间(GMT)的差值，小时和分钟之间用冒号分隔。例如：+02:00。 md5:ba4884ccf645f9f2
		'T': "MST",                       // 区域：时区缩写。例如：UTC，EST，MDT 等等。 md5:0fe0c9bfc89456fc
		'c': "2006-01-02T15:04:05-07:00", // 格式：ISO 8601日期。例如：2004-02-12T15:19:21+00:00。 md5:456ee7005cc751d8
		'r': "Mon, 02 Jan 06 15:04 MST",  // 格式：RFC 2822 格式的日期。例如：Thu, 21 Dec 2000 16:01:07 +0200。 md5:832bc55e456d82e4
	}

		// 周到数字的映射。 md5:a200a191f65df59c
	weekMap = map[string]string{
		"Sunday":    "0",
		"Monday":    "1",
		"Tuesday":   "2",
		"Wednesday": "3",
		"Thursday":  "4",
		"Friday":    "5",
		"Saturday":  "6",
	}

		// 每个非闰年的月份的天数。 md5:2519753cb5104326
	dayOfMonth = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
)

// Format 使用自定义`format`格式化并返回格式化结果。如果你想要遵循stdlib（标准库）的布局，可以参考 Layout 方法。
// md5:8f91fb876a2c8a6d
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
								// 特定的字符应该在这里被处理。 md5:e72d802cbb002d25
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

// FormatNew 根据给定的自定义`format`格式化并返回一个新的 Time 对象。 md5:651ea5fc95b95b2e
func (t *Time) FormatNew(format string) *Time {
	if t == nil {
		return nil
	}
	return NewFromStr(t.Format(format))
}

// FormatTo使用给定的自定义`format`格式化`t`。 md5:d34597383793dd06
func (t *Time) FormatTo(format string) *Time {
	if t == nil {
		return nil
	}
	t.Time = NewFromStr(t.Format(format)).Time
	return t
}

// Layout使用stdlib布局格式化时间并返回格式化后的结果。 md5:bf29a9bede753c3a
func (t *Time) Layout(layout string) string {
	if t == nil {
		return ""
	}
	return t.Time.Format(layout)
}

// LayoutNew 使用stdlib布局格式化时间，并返回一个新的Time对象。 md5:6849149696989dbb
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

// LayoutTo 使用stdlib布局格式化`t`。 md5:fb1407c2e7429179
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

// IsLeapYear 检查给定的时间是否为闰年。 md5:cc71272fbb6cec2b
func (t *Time) IsLeapYear() bool {
	year := t.Year()
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

// DayOfYear 检查并返回该年中的某一天的位置。 md5:c518ddcab14a7a55
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

// DaysInMonth 返回当前月份的天数。 md5:0cd1f14a8bb1f8fc
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

// WeeksOfYear 返回当前年份中的当前周数。 md5:a51898ffdc6f00df
func (t *Time) WeeksOfYear() int {
	_, week := t.ISOWeek()
	return week
}

// formatToStdLayout 将自定义格式转换为标准库布局。 md5:41bf62abc4e12b34
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
								// 处理特定字符。 md5:65410637e4dc8fe5
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

// formatToRegexPattern 将自定义格式转换为其对应的正则表达式。 md5:39433c1bc246e82c
func formatToRegexPattern(format string) string {
	s := regexp.QuoteMeta(formatToStdLayout(format))
	s, _ = gregex.ReplaceString(`[0-9]`, `[0-9]`, s)
	s, _ = gregex.ReplaceString(`[A-Za-z]`, `[A-Za-z]`, s)
	s, _ = gregex.ReplaceString(`\s+`, `\s+`, s)
	return s
}

// formatMonthDaySuffixMap 返回当前日期的简短英文后缀词。 md5:faf4d1234f7e3ab7
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
