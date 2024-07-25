// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gcron

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

// checkMeetAndUpdateLastSeconds 检查给定的时间 `t` 是否满足作业的可运行点。
// 这个函数每秒被调用一次。 md5:7e9ffb92b302c297
func (s *cronSchedule) checkMeetAndUpdateLastSeconds(ctx context.Context, currentTime time.Time) (ok bool) {
	var (
		lastCheckTimestamp = s.getAndUpdateLastCheckTimestamp(ctx, currentTime)
		lastCheckTime      = gtime.NewFromTimeStamp(lastCheckTimestamp)
		lastMeetTime       = gtime.NewFromTimeStamp(s.lastMeetTimestamp.Val())
	)
	defer func() {
		if ok {
			s.lastMeetTimestamp.Set(currentTime.Unix())
		}
	}()
	if !s.checkMinIntervalAndItemMapMeet(lastMeetTime.Time, lastCheckTime.Time, currentTime) {
		return false
	}
	return true
}

func (s *cronSchedule) checkMinIntervalAndItemMapMeet(
	lastMeetTime, lastCheckTime, currentTime time.Time,
) (ok bool) {
	if s.everySeconds != 0 {
		// 它使用间隔进行检查。 md5:59665c64bd4530b9
		secondsAfterCreated := lastCheckTime.UnixNano()/1e9 - s.createTimestamp
		if secondsAfterCreated > 0 {
			return secondsAfterCreated%s.everySeconds == 0
		}
		return false
	}
	if !s.checkMeetSecond(lastMeetTime, currentTime) {
		return false
	}
	if !s.checkMeetMinute(currentTime) {
		return false
	}
	if !s.checkMeetHour(currentTime) {
		return false
	}
	if !s.checkMeetDay(currentTime) {
		return false
	}
	if !s.checkMeetMonth(currentTime) {
		return false
	}
	if !s.checkMeetWeek(currentTime) {
		return false
	}
	return true
}

func (s *cronSchedule) checkMeetSecond(lastMeetTime, currentTime time.Time) (ok bool) {
	if s.ignoreSeconds {
		if currentTime.Unix()-lastMeetTime.Unix() < 60 {
			return false
		}
	} else {
		// 如果此模式在精确到秒的时间设置，
		// 则不允许在同一时间执行。 md5:b3ec1446bf507768
		if len(s.secondMap) == 1 && lastMeetTime.Format(time.RFC3339) == currentTime.Format(time.RFC3339) {
			return false
		}
		if !s.keyMatch(s.secondMap, currentTime.Second()) {
			return false
		}
	}
	return true
}

func (s *cronSchedule) checkMeetMinute(currentTime time.Time) (ok bool) {
	if !s.keyMatch(s.minuteMap, currentTime.Minute()) {
		return false
	}
	return true
}

func (s *cronSchedule) checkMeetHour(currentTime time.Time) (ok bool) {
	if !s.keyMatch(s.hourMap, currentTime.Hour()) {
		return false
	}
	return true
}

func (s *cronSchedule) checkMeetDay(currentTime time.Time) (ok bool) {
	if !s.keyMatch(s.dayMap, currentTime.Day()) {
		return false
	}
	return true
}

func (s *cronSchedule) checkMeetMonth(currentTime time.Time) (ok bool) {
	if !s.keyMatch(s.monthMap, int(currentTime.Month())) {
		return false
	}
	return true
}

func (s *cronSchedule) checkMeetWeek(currentTime time.Time) (ok bool) {
	if !s.keyMatch(s.weekMap, int(currentTime.Weekday())) {
		return false
	}
	return true
}

func (s *cronSchedule) keyMatch(m map[int]struct{}, key int) bool {
	_, ok := m[key]
	return ok
}

func (s *cronSchedule) checkItemMapMeet(lastMeetTime, currentTime time.Time) (ok bool) {
	// second.
	if s.ignoreSeconds {
		if currentTime.Unix()-lastMeetTime.Unix() < 60 {
			return false
		}
	} else {
		if !s.keyMatch(s.secondMap, currentTime.Second()) {
			return false
		}
	}
	// minute.
	if !s.keyMatch(s.minuteMap, currentTime.Minute()) {
		return false
	}
	// hour.
	if !s.keyMatch(s.hourMap, currentTime.Hour()) {
		return false
	}
	// day.
	if !s.keyMatch(s.dayMap, currentTime.Day()) {
		return false
	}
	// month.
	if !s.keyMatch(s.monthMap, int(currentTime.Month())) {
		return false
	}
	// week.
	if !s.keyMatch(s.weekMap, int(currentTime.Weekday())) {
		return false
	}
	return true
}
