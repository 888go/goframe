// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcron

import (
	"context"
	"strconv"
	"strings"
	"time"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/text/gregex"
)

// cronSchedule 是cron作业的计划安排。
type cronSchedule struct {
	createTimestamp int64            // 创建时间戳（以秒为单位）
	everySeconds    int64            // 运行间隔时间（单位：秒）
	pattern         string           // 原始的cron模式字符串。
	secondMap       map[int]struct{} // Job可以在以下秒数运行。
	minuteMap       map[int]struct{} // Job可以在以下分钟数运行。
	hourMap         map[int]struct{} // Job可以在以下小时数运行。
	dayMap          map[int]struct{} // Job可以在这些天数中运行。
	weekMap         map[int]struct{} // Job可以在这些周数中运行。
	monthMap        map[int]struct{} // Job 可以在以下月份数字中运行。
	lastTimestamp   *gtype.Int64     // 上次时间戳编号，用于修正某些延迟情况下的时间戳。
}

const (
	// 正则表达式用于cron模式，该模式包含6部分时间单元。
	regexForCron           = `^([\-/\d\*\?,]+)\s+([\-/\d\*\?,]+)\s+([\-/\d\*\?,]+)\s+([\-/\d\*\?,]+)\s+([\-/\d\*\?,A-Za-z]+)\s+([\-/\d\*\?,A-Za-z]+)$`
	patternItemTypeUnknown = iota
	patternItemTypeWeek
	patternItemTypeMonth
)

var (
	// 预定义模式映射
	predefinedPatternMap = map[string]string{
		"@yearly":   "0 0 0 1 1 *",
		"@annually": "0 0 0 1 1 *",
		"@monthly":  "0 0 0 1 * *",
		"@weekly":   "0 0 0 * * 0",
		"@daily":    "0 0 0 * * *",
		"@midnight": "0 0 0 * * *",
		"@hourly":   "0 0 * * * *",
	}
	// 短月份名称及其对应的数字。
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
	// 完整的月份名称转为对应的数字。
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
	// 短星期名称及其对应的数字。
	weekShortNameMap = map[string]int{
		"sun": 0,
		"mon": 1,
		"tue": 2,
		"wed": 3,
		"thu": 4,
		"fri": 5,
		"sat": 6,
	}
	// 完整周名称及其对应的周数。
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

// newSchedule 根据给定的cron模式创建并返回一个schedule对象。
func newSchedule(pattern string) (*cronSchedule, error) {
	var currentTimestamp = time.Now().Unix()
	// 检查预定义的模式是否存在
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
				createTimestamp: currentTimestamp,
				everySeconds:    int64(d.Seconds()),
				pattern:         pattern,
				lastTimestamp:   gtype.NewInt64(currentTimestamp),
			}, nil
		} else {
			return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid pattern: "%s"`, pattern)
		}
	}
// 处理常见的cron模式，例如：
// 0 0 0 1 1 2
// （注：该段代码注释省略了对cron模式的具体解释，以下是补充说明）
// 上述代码注释提到的"常见的cron模式"在Unix/Linux系统中用于表示定时任务的时间配置，
// 其格式为：分 时 天(月) 月 星期 周
// 示例 "0 0 0 1 1 2" 的含义是：
// 在每月的第一天（1号）的第一个星期二（2）的凌晨0点0分执行定时任务。
	if match, _ := gregex.MatchString(regexForCron, pattern); len(match) == 7 {
		schedule := &cronSchedule{
			createTimestamp: currentTimestamp,
			everySeconds:    0,
			pattern:         pattern,
			lastTimestamp:   gtype.NewInt64(currentTimestamp),
		}
		// Second.
		if m, err := parsePatternItem(match[1], 0, 59, false); err != nil {
			return nil, err
		} else {
			schedule.secondMap = m
		}
		// Minute.
		if m, err := parsePatternItem(match[2], 0, 59, false); err != nil {
			return nil, err
		} else {
			schedule.minuteMap = m
		}
		// Hour.
		if m, err := parsePatternItem(match[3], 0, 23, false); err != nil {
			return nil, err
		} else {
			schedule.hourMap = m
		}
		// Day.
		if m, err := parsePatternItem(match[4], 1, 31, true); err != nil {
			return nil, err
		} else {
			schedule.dayMap = m
		}
		// Month.
		if m, err := parsePatternItem(match[5], 1, 12, false); err != nil {
			return nil, err
		} else {
			schedule.monthMap = m
		}
		// Week.
		if m, err := parsePatternItem(match[6], 0, 6, true); err != nil {
			return nil, err
		} else {
			schedule.weekMap = m
		}
		return schedule, nil
	}
	return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid pattern: "%s"`, pattern)
}

// parsePatternItem 解析模式中的每一项，并将结果以映射形式返回，该映射用于索引。
func parsePatternItem(item string, min int, max int, allowQuestionMark bool) (map[int]struct{}, error) {
	m := make(map[int]struct{}, max-min+1)
	if item == "*" || (allowQuestionMark && item == "?") {
		for i := min; i <= max; i++ {
			m[i] = struct{}{}
		}
		return m, nil
	}
	// Like: MON,FRI
	for _, itemElem := range strings.Split(item, ",") {
		var (
			interval      = 1
			intervalArray = strings.Split(itemElem, "/")
		)
		if len(intervalArray) == 2 {
			if number, err := strconv.Atoi(intervalArray[1]); err != nil {
				return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid pattern item: "%s"`, itemElem)
			} else {
				interval = number
			}
		}
		var (
			rangeMin   = min
			rangeMax   = max
			itemType   = patternItemTypeUnknown
			rangeArray = strings.Split(intervalArray[0], "-") // 类似于：1-30，JAN-DEC
		)
		switch max {
		case 6:
			// 正在检查周字段。
			itemType = patternItemTypeWeek

		case 12:
			// 正在检查月份字段。
			itemType = patternItemTypeMonth
		}
		// Eg: */5
		if rangeArray[0] != "*" {
			if number, err := parsePatternItemValue(rangeArray[0], itemType); err != nil {
				return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid pattern item: "%s"`, itemElem)
			} else {
				rangeMin = number
				if len(intervalArray) == 1 {
					rangeMax = number
				}
			}
		}
		if len(rangeArray) == 2 {
			if number, err := parsePatternItemValue(rangeArray[1], itemType); err != nil {
				return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid pattern item: "%s"`, itemElem)
			} else {
				rangeMax = number
			}
		}
		for i := rangeMin; i <= rangeMax; i += interval {
			m[i] = struct{}{}
		}
	}
	return m, nil
}

// parsePatternItemValue 根据字段类型将字段值解析为数字。
func parsePatternItemValue(value string, itemType int) (int, error) {
	if gregex.IsMatchString(`^\d+$`, value) {
		// 这是一个纯数字。
		if number, err := strconv.Atoi(value); err == nil {
			return number, nil
		}
	} else {
// 检查其中是否包含字母，
// 根据预定义的映射将其值转换为数字。
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

// checkMeetAndUpdateLastSeconds 检查给定的时间 `t` 是否满足作业的可执行时间点，并更新最后执行秒数。
func (s *cronSchedule) checkMeetAndUpdateLastSeconds(ctx context.Context, t time.Time) bool {
	var (
		lastTimestamp = s.getAndUpdateLastTimestamp(ctx, t)
		lastTime      = gtime.NewFromTimeStamp(lastTimestamp)
	)

	if s.everySeconds != 0 {
		// 它使用间隔进行检查。
		secondsAfterCreated := lastTime.Timestamp() - s.createTimestamp
		if secondsAfterCreated > 0 {
			return secondsAfterCreated%s.everySeconds == 0
		}
		return false
	}

	// 它使用标准cron模式进行检查。
	if _, ok := s.secondMap[lastTime.Second()]; !ok {
		return false
	}
	if _, ok := s.minuteMap[lastTime.Minute()]; !ok {
		return false
	}
	if _, ok := s.hourMap[lastTime.Hour()]; !ok {
		return false
	}
	if _, ok := s.dayMap[lastTime.Day()]; !ok {
		return false
	}
	if _, ok := s.monthMap[lastTime.Month()]; !ok {
		return false
	}
	if _, ok := s.weekMap[int(lastTime.Weekday())]; !ok {
		return false
	}
	return true
}

// Next 函数返回该计划下一次激活的时间，该时间大于给定的时间。
// 如果找不到满足计划要求的时间，则返回零时间（即时间的零值，表示无效时间）。
func (s *cronSchedule) Next(t time.Time) time.Time {
	if s.everySeconds != 0 {
		var (
			diff  = t.Unix() - s.createTimestamp
			count = diff/s.everySeconds + 1
		)
		return t.Add(time.Duration(count*s.everySeconds) * time.Second)
	}

	// 从最早可能的时间开始（即即将到来的下一秒）。
	t = t.Add(1*time.Second - time.Duration(t.Nanosecond())*time.Nanosecond)
	var (
		loc       = t.Location()
		added     = false
		yearLimit = t.Year() + 5
	)

WRAP:
	if t.Year() > yearLimit {
		return t // 谁会在意五年后运行的那份工作
	}

	for !s.match(s.monthMap, int(t.Month())) {
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, loc)
		}
		t = t.AddDate(0, 1, 0)
		// need recheck
		if t.Month() == time.January {
			goto WRAP
		}
	}

	for !s.dayMatches(t) {
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
		}
		t = t.AddDate(0, 0, 1)

// 注意由于DST（夏令时）导致的小时数是否不再为午夜。
// 如果是23点则加1小时，如果是1点则减1小时。
		if t.Hour() != 0 {
			if t.Hour() > 12 {
				t = t.Add(time.Duration(24-t.Hour()) * time.Hour)
			} else {
				t = t.Add(time.Duration(-t.Hour()) * time.Hour)
			}
		}
		if t.Day() == 1 {
			goto WRAP
		}
	}
	for !s.match(s.hourMap, t.Hour()) {
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, loc)
		}
		t = t.Add(time.Hour)
		// need recheck
		if t.Hour() == 0 {
			goto WRAP
		}
	}
	for !s.match(s.minuteMap, t.Minute()) {
		if !added {
			added = true
			t = t.Truncate(time.Minute)
		}
		t = t.Add(1 * time.Minute)

		if t.Minute() == 0 {
			goto WRAP
		}
	}
	for !s.match(s.secondMap, t.Second()) {
		if !added {
			added = true
			t = t.Truncate(time.Second)
		}
		t = t.Add(1 * time.Second)
		if t.Second() == 0 {
			goto WRAP
		}
	}
	return t.In(loc)
}

// dayMatches 函数返回一个布尔值，如果给定时间满足该计划的周几和每月几日的限制条件，则返回true。
func (s *cronSchedule) dayMatches(t time.Time) bool {
	_, ok1 := s.dayMap[t.Day()]
	_, ok2 := s.weekMap[int(t.Weekday())]
	return ok1 && ok2
}

func (s *cronSchedule) match(m map[int]struct{}, key int) bool {
	_, ok := m[key]
	return ok
}
