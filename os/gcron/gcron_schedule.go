// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gcron

import (
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
)

// cronSchedule 是定时任务的调度计划。 md5:4731e43288725f27
type cronSchedule struct {
	createTimestamp int64            // 创建时间的时间戳，以秒为单位。 md5:4a0001cda2177f41
	everySeconds    int64            // 运行间隔（以秒为单位）。 md5:a62fd57ffa9e26f4
	pattern         string           // 在创建cron作业时传递的原始cron模式字符串。 md5:18b07692590ddf66
	ignoreSeconds   bool             // 标记该模式是标准的5部分cron表达式模式，而不是6部分模式。 md5:89774325ba9632d2
	secondMap       map[int]struct{} // 该Job可以在这些秒数内运行。 md5:603e4f208dcc04bf
	minuteMap       map[int]struct{} // Job可以在这些分钟数运行。 md5:9cc64d9456bc318a
	hourMap         map[int]struct{} // Job可以在这些小时数运行。 md5:cf1a7bc2b7ada427
	dayMap          map[int]struct{} // Job 可以在这些天数中运行。 md5:9be6d3ae1549f6c8
	weekMap         map[int]struct{} // Job 可以在这些星期数中运行。 md5:e9d2ed887e372b17
	monthMap        map[int]struct{} // Job可以在这些月份运行。 md5:e58af4ea6da7e868

	// 这个字段存储满足计划的最新时间戳。 md5:df6f9fc73fbf03d6
	lastMeetTimestamp *gtype.Int64

	// 最后一个时间戳编号，用于在某些延迟情况下固定时间戳。 md5:6839316ecd982e4b
	lastCheckTimestamp *gtype.Int64
}

type patternItemType int

const (
	patternItemTypeSecond patternItemType = iota
	patternItemTypeMinute
	patternItemTypeHour
	patternItemTypeDay
	patternItemTypeWeek
	patternItemTypeMonth
)

const (
	// 正则表达式表示的cron模式，包含6个时间单位部分。 md5:75e472ef39ca5aab
	regexForCron = `^([\-/\d\*,#]+)\s+([\-/\d\*,]+)\s+([\-/\d\*,]+)\s+([\-/\d\*\?,]+)\s+([\-/\d\*,A-Za-z]+)\s+([\-/\d\*\?,A-Za-z]+)$`
)

var (
	// 预定义的模式映射。 md5:dc23a289b509e3b6
	predefinedPatternMap = map[string]string{
		"@yearly":   "# 0 0 1 1 *",
		"@annually": "# 0 0 1 1 *",
		"@monthly":  "# 0 0 1 * *",
		"@weekly":   "# 0 0 * * 0",
		"@daily":    "# 0 0 * * *",
		"@midnight": "# 0 0 * * *",
		"@hourly":   "# 0 * * * *",
	}
	// 短月名到其对应的数字。 md5:44f6938b62580af0
	monthShortNameMap = map[string]int{
		"jan": 1,
		"feb": 2,
		"mar": 3,
		"apr": 4,
		"may": 5,
		"jun": 6,
		"jul": 7,
		"aug": 8,
		"sep": 9,
		"oct": 10,
		"nov": 11,
		"dec": 12,
	}
	// 完整的月份名称转换为其对应的数字。 md5:e9b9f99b1f2191d0
	monthFullNameMap = map[string]int{
		"january":   1,
		"february":  2,
		"march":     3,
		"april":     4,
		"may":       5,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
	}
	// 短星期名转换为对应的数字。 md5:c8dde2776e296b0a
	weekShortNameMap = map[string]int{
		"sun": 0,
		"mon": 1,
		"tue": 2,
		"wed": 3,
		"thu": 4,
		"fri": 5,
		"sat": 6,
	}
	// 完整的星期名称到其数字。 md5:05d1a360fc5b25ee
	weekFullNameMap = map[string]int{
		"sunday":    0,
		"monday":    1,
		"tuesday":   2,
		"wednesday": 3,
		"thursday":  4,
		"friday":    5,
		"saturday":  6,
	}
)

// newSchedule根据给定的cron模式创建并返回一个调度对象。 md5:14dff188c64f1e56
func newSchedule(pattern string) (*cronSchedule, error) {
	var currentTimestamp = time.Now().Unix()
	// 检查给定的`pattern`是否在预定义的模式中。 md5:31badfbc0ed60d2b
	if match, _ := gregex.MatchString(`(@\w+)\s*(\w*)\s*`, pattern); len(match) > 0 {
		key := strings.ToLower(match[1])
		if v, ok := predefinedPatternMap[key]; ok {
			pattern = v
		} else if strings.Compare(key, "@every") == 0 {
			d, err := gtime.ParseDuration(match[2])
			if err != nil {
				return nil, err
			}
			return &cronSchedule{
				createTimestamp:    currentTimestamp,
				everySeconds:       int64(d.Seconds()),
				pattern:            pattern,
				lastMeetTimestamp:  gtype.NewInt64(currentTimestamp),
				lastCheckTimestamp: gtype.NewInt64(currentTimestamp),
			}, nil
		} else {
			return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid pattern: "%s"`, pattern)
		}
	}
	// 处理给定的`pattern`作为常见的6部分模式。 md5:224ce220d8873fe0
	match, _ := gregex.MatchString(regexForCron, pattern)
	if len(match) != 7 {
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid pattern: "%s"`, pattern)
	}
	var (
		err error
		cs  = &cronSchedule{
			createTimestamp:    currentTimestamp,
			everySeconds:       0,
			pattern:            pattern,
			lastMeetTimestamp:  gtype.NewInt64(currentTimestamp),
			lastCheckTimestamp: gtype.NewInt64(currentTimestamp),
		}
	)

	// Second.
	if match[1] == "#" {
		cs.ignoreSeconds = true
	} else {
		cs.secondMap, err = parsePatternItem(match[1], 0, 59, false, patternItemTypeSecond)
		if err != nil {
			return nil, err
		}
	}
	// Minute.
	cs.minuteMap, err = parsePatternItem(match[2], 0, 59, false, patternItemTypeMinute)
	if err != nil {
		return nil, err
	}
	// Hour.
	cs.hourMap, err = parsePatternItem(match[3], 0, 23, false, patternItemTypeHour)
	if err != nil {
		return nil, err
	}
	// Day.
	cs.dayMap, err = parsePatternItem(match[4], 1, 31, true, patternItemTypeDay)
	if err != nil {
		return nil, err
	}
	// Month.
	cs.monthMap, err = parsePatternItem(match[5], 1, 12, false, patternItemTypeMonth)
	if err != nil {
		return nil, err
	}
	// Week.
	cs.weekMap, err = parsePatternItem(match[6], 0, 6, true, patternItemTypeWeek)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

// parsePatternItem 解析模式中的每个项目，并将结果作为映射返回，该映射用于索引。 md5:66716855d8c0f694
func parsePatternItem(
	item string, min int, max int,
	allowQuestionMark bool, itemType patternItemType,
) (itemMap map[int]struct{}, err error) {
	itemMap = make(map[int]struct{}, max-min+1)
	if item == "*" || (allowQuestionMark && item == "?") {
		for i := min; i <= max; i++ {
			itemMap[i] = struct{}{}
		}
		return itemMap, nil
	}
	// 例子：1-10/2，11-30/3
// 
// 这个注释表示一个范围的分组示例。"1-10/2" 表示从1开始到10，每2个数一组；"11-30/3" 表示从11开始到30，每3个数一组。 md5:7074496c7eb487df
	var number int
	for _, itemElem := range strings.Split(item, ",") {
		var (
			interval      = 1
			intervalArray = strings.Split(itemElem, "/")
		)
		if len(intervalArray) == 2 {
			if number, err = strconv.Atoi(intervalArray[1]); err != nil {
				return nil, gerror.NewCodef(
					gcode.CodeInvalidParameter, `invalid pattern item: "%s"`, itemElem,
				)
			} else {
				interval = number
			}
		}
		var (
			rangeMin   = min
			rangeMax   = max
			rangeArray = strings.Split(intervalArray[0], "-") // Example: 1-30, JAN-DEC
		)
		// Example: 1-30/2
		if rangeArray[0] != "*" {
			if number, err = parseWeekAndMonthNameToInt(rangeArray[0], itemType); err != nil {
				return nil, gerror.NewCodef(
					gcode.CodeInvalidParameter, `invalid pattern item: "%s"`, itemElem,
				)
			} else {
				rangeMin = number
				if len(intervalArray) == 1 {
					rangeMax = number
				}
			}
		}
		// Example: 1-30/2
		if len(rangeArray) == 2 {
			if number, err = parseWeekAndMonthNameToInt(rangeArray[1], itemType); err != nil {
				return nil, gerror.NewCodef(
					gcode.CodeInvalidParameter, `invalid pattern item: "%s"`, itemElem,
				)
			} else {
				rangeMax = number
			}
		}
		for i := rangeMin; i <= rangeMax; i += interval {
			itemMap[i] = struct{}{}
		}
	}
	return
}

// parseWeekAndMonthNameToInt 根据字段类型将字段值解析为数字。 md5:10e98c83dca57c49
func parseWeekAndMonthNameToInt(value string, itemType patternItemType) (int, error) {
	if gregex.IsMatchString(`^\d+$`, value) {
		// It is pure number.
		if number, err := strconv.Atoi(value); err == nil {
			return number, nil
		}
	} else {
// 检查是否包含字母，
// 根据预定义的映射将值转换为数字。
// md5:d6cf713cc1230de9
		switch itemType {
		case patternItemTypeWeek:
			if number, ok := weekShortNameMap[strings.ToLower(value)]; ok {
				return number, nil
			}
			if number, ok := weekFullNameMap[strings.ToLower(value)]; ok {
				return number, nil
			}
		case patternItemTypeMonth:
			if number, ok := monthShortNameMap[strings.ToLower(value)]; ok {
				return number, nil
			}
			if number, ok := monthFullNameMap[strings.ToLower(value)]; ok {
				return number, nil
			}
		}
	}
	return 0, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid pattern value: "%s"`, value)
}
