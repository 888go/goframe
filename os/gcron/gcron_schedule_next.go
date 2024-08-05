// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gcron

import (
	"time"
)

// Next 返回此调度的下次激活时间，大于给定的时间。如果找不到满足调度的时间，则返回零时间。
// md5:a7867a51955c4fd0
func (s *cronSchedule) Next(lastMeetTime time.Time) time.Time {
	if s.everySeconds != 0 {
		var (
			diff  = lastMeetTime.Unix() - s.createTimestamp
			count = diff/s.everySeconds + 1
		)
		return lastMeetTime.Add(time.Duration(count*s.everySeconds) * time.Second)
	}

	var currentTime = lastMeetTime
	if s.ignoreSeconds {
				// 从最早的时间开始（即将来临的分钟）。 md5:8677cc0d5c129643
		currentTime = currentTime.Add(1*time.Minute - time.Duration(currentTime.Nanosecond())*time.Nanosecond)
	} else {
				// 从最早可能的时间开始（即即将到来的下一秒）。 md5:ea5d8844c8e2b464
		currentTime = currentTime.Add(1*time.Second - time.Duration(currentTime.Nanosecond())*time.Nanosecond)
	}

	var (
		loc       = currentTime.Location()
		yearLimit = currentTime.Year() + 5
	)

WRAP:
	if currentTime.Year() > yearLimit {
		return currentTime // 谁会在五年后关心这个任务呢？. md5:b515d5d4b0e4c598
	}

	for !s.checkMeetMonth(currentTime) {
		currentTime = currentTime.AddDate(0, 1, 0)
		currentTime = time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, loc)
		if currentTime.Month() == time.January {
			goto WRAP
		}
	}
	for !s.checkMeetWeek(currentTime) || !s.checkMeetDay(currentTime) {
		currentTime = currentTime.AddDate(0, 0, 1)
		currentTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, loc)
		if currentTime.Day() == 1 {
			goto WRAP
		}
	}
	for !s.checkMeetHour(currentTime) {
		currentTime = currentTime.Add(time.Hour)
		currentTime = currentTime.Truncate(time.Hour)
		if currentTime.Hour() == 0 {
			goto WRAP
		}
	}
	for !s.checkMeetMinute(currentTime) {
		currentTime = currentTime.Add(1 * time.Minute)
		currentTime = currentTime.Truncate(time.Minute)
		if currentTime.Minute() == 0 {
			goto WRAP
		}
	}

	for !s.checkMeetSecond(lastMeetTime, currentTime) {
		currentTime = currentTime.Add(1 * time.Second)
		if currentTime.Second() == 0 {
			goto WRAP
		}
	}
	return currentTime.In(loc)
}
