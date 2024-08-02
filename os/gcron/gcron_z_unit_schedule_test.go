// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 定时cron类

import (
	"strings"
	"testing"
	"time"

	gtest "github.com/888go/goframe/test/gtest"
)

func TestSlash(t *testing.T) {
	runs := []struct {
		spec     string
		expected map[int]struct{}
	}{
		{"0 0 0 * Feb Mon/2", map[int]struct{}{1: {}, 3: {}, 5: {}}},
		{"0 0 0 * Feb *", map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 0: {}}},
	}
	gtest.C(t, func(t *gtest.T) {
		for _, c := range runs {
			s, err := newSchedule(c.spec)
			if err != nil {
				t.Fatal(err)

			}
			t.AssertEQ(s.weekMap, c.expected)
		}
	})
}

func TestNext(t *testing.T) {
	runs := []struct {
		time, spec string
		expected   string
	}{
		// Simple cases
		{"Mon Jul 9 14:45 2012", "0 0/15 * * * *", "Mon Jul 9 15:00 2012"},
		{"Mon Jul 9 14:59 2012", "0 0/15 * * * *", "Mon Jul 9 15:00 2012"},
		{"Mon Jul 9 14:59:59 2012", "0 0/15 * * * *", "Mon Jul 9 15:00 2012"},

		// Wrap around hours
		{"Mon Jul 9 15:45 2012", "0 20-35/15 * * * *", "Mon Jul 9 16:20 2012"},

		// Wrap around days
		{"Mon Jul 9 23:46 2012", "0 */15 * * * *", "Tue Jul 10 00:00 2012"},
		{"Mon Jul 9 23:45 2012", "0 20-35/15 * * * *", "Tue Jul 10 00:20 2012"},
		{"Mon Jul 9 23:35:51 2012", "15/35 20-35/15 * * * *", "Tue Jul 10 00:20:15 2012"},
		{"Mon Jul 9 23:35:51 2012", "15/35 20-35/15 1/2 * * *", "Tue Jul 10 01:20:15 2012"},
		{"Mon Jul 9 23:35:51 2012", "15/35 20-35/15 10-12 * * *", "Tue Jul 10 10:20:15 2012"},

		{"Mon Jul 9 23:35:51 2012", "15/35 20-35/15 1/2 */2 * *", "Thu Jul 11 01:20:15 2012"},
		{"Mon Jul 9 23:35:51 2012", "15/35 20-35/15 * 9-20 * *", "Wed Jul 10 00:20:15 2012"},
		{"Mon Jul 9 23:35:51 2012", "15/35 20-35/15 * 9-20 Jul *", "Wed Jul 10 00:20:15 2012"},

		// Wrap around months
		{"Mon Jul 9 23:35 2012", "0 0 0 9 Apr-Oct ?", "Thu Aug 9 00:00 2012"},
		{"Mon Jul 9 23:35 2012", "0 0 0 */5 Apr,Aug,Oct Mon", "Mon Aug 6 00:00 2012"},
		{"Mon Jul 9 23:35 2012", "0 0 0 */5 Oct Mon", "Mon Oct 1 00:00 2012"},

		// Wrap around years
		{"Mon Jul 9 23:35 2012", "0 0 0 * Feb Mon", "Mon Feb 4 00:00 2013"},
		{"Mon Jul 9 23:35 2012", "0 0 0 * Feb Mon/2", "Fri Feb 1 00:00 2013"},

						// 对分钟、小时、天、月和年进行循环. md5:5e7fbfff239d3004
		{"Mon Dec 31 23:59:45 2012", "0 * * * * *", "Tue Jan 1 00:00:00 2013"},

		// Leap year
		{"Mon Jul 9 23:35 2012", "0 0 0 29 Feb ?", "Mon Feb 29 00:00 2016"},

				// 预定义的模式映射。 md5:dc23a289b509e3b6
		{"Mon Jul 9 23:35 2012", "@yearly", "Sun Jan 1 00:00:00 2013"},
		{"Mon Jul 9 23:35 2012", "@annually", "Sun Jan 1 00:00:00 2013"},
		{"Mon Jul 9 23:35 2012", "@monthly", "Mon Aug 1 00:00:00 2012"},
		{"Mon Jul 9 23:35 2012", "@weekly", "Sun Jul 15 00:00:00 2012"},
		{"Mon Jul 9 23:35 2012", "@daily", "Tue Jul 10 00:00:00 2012"},
		{"Mon Jul 9 23:35 2012", "@midnight", "Tue Jul 10 00:00:00 2012"},
		{"Mon Jul 9 23:35 2012", "@hourly", "Tue Jul 10 00:00:00 2012"},

		// Ignore seconds.
		{"Mon Jul 9 23:35 2012", "# * * * * *", "Mon Jul 9 23:36 2012"},
		{"Mon Jul 9 23:35 2012", "# */2 * * * *", "Mon Jul 9 23:36 2012"},
	}

	for _, c := range runs {
		s, err := newSchedule(c.spec)
		if err != nil {
			t.Error(err)
			continue
		}
				// 使用fmt.Printf函数打印sched的值，其中"%+v"是一个格式化字符串，表示输出 sched 的值，并且包括其类型信息。 md5:00923f8d723abc44
		actual := s.Next(getTime(c.time))
		expected := getTime(c.expected)
		if !(actual.Unix() == expected.Unix()) {
			t.Errorf(
				"%s, \"%s\": (expected) %v != %v (actual)",
				c.time, c.spec, expected, actual,
			)
		}
	}
}

func getTime(value string) time.Time {
	if value == "" {
		return time.Time{}
	}

	var location = time.Local
	if strings.HasPrefix(value, "TZ=") {
		parts := strings.Fields(value)
		loc, err := time.LoadLocation(parts[0][len("TZ="):])
		if err != nil {
			panic("could not parse location:" + err.Error())
		}
		location = loc
		value = parts[1]
	}

	var layouts = []string{
		"Mon Jan 2 15:04 2006",
		"Mon Jan 2 15:04:05 2006",
	}
	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, value, location); err == nil {
			return t
		}
	}
	if t, err := time.ParseInLocation("2006-01-02T15:04:05-0700", value, location); err == nil {
		return t
	}
	panic("could not parse time value " + value)
}
