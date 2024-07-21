// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtime

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	setTimeZoneMu   sync.Mutex
	setTimeZoneName string
	zoneMap         = make(map[string]*time.Location)
	zoneMu          sync.RWMutex
)

// SetTimeZone sets the time zone for current whole process.
// The parameter `zone` is an area string specifying corresponding time zone,
//
// PLEASE VERY NOTE THAT:
// 1. This should be called before package "time" import.
// 2. This function should be called once.
// 3. Please refer to issue: https://github.com/golang/go/issues/34814
// ff:设置时区
// zone:时区
// err:错误
func SetTimeZone(zone string) (err error) {
	setTimeZoneMu.Lock()
	defer setTimeZoneMu.Unlock()
	if setTimeZoneName != "" && !strings.EqualFold(zone, setTimeZoneName) {
		return gerror.NewCodef(
			gcode.CodeInvalidOperation,
			`process timezone already set using "%s"`,
			setTimeZoneName,
		)
	}
	defer func() {
		if err == nil {
			setTimeZoneName = zone
		}
	}()

	// 它已经被设置为time.Local。 md5:1fa5641c118746d6
	if strings.EqualFold(zone, time.Local.String()) {
		return
	}

	// 从指定的名称加载区域信息。 md5:dada678d8dfb8df3
	location, err := time.LoadLocation(zone)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `time.LoadLocation failed for zone "%s"`, zone)
		return err
	}

	// 为了保险起见，更新一次time.Local。 md5:b8b9f7daf1488924
	time.Local = location

	// 更新*nix系统中的时区环境。 md5:e9774c5a2d209c8d
	var (
		envKey   = "TZ"
		envValue = location.String()
	)
	if err = os.Setenv(envKey, envValue); err != nil {
		err = gerror.WrapCodef(
			gcode.CodeUnknown,
			err,
			`set environment failed with key "%s", value "%s"`,
			envKey, envValue,
		)
	}
	return
}

// ToLocation 将当前时间转换为指定时区的时间。 md5:ee4fd7d4de93340a
// ff:转换时区Location
// t:
// location:时区
func (t *Time) ToLocation(location *time.Location) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.In(location)
	return newTime
}

// ToZone 将当前时间转换为指定的时区，如 Asia/Shanghai。 md5:1226213d40f57eb2
// ff:转换时区
// t:
// zone:时区
func (t *Time) ToZone(zone string) (*Time, error) {
	if location, err := t.getLocationByZoneName(zone); err == nil {
		return t.ToLocation(location), nil
	} else {
		return nil, err
	}
}

func (t *Time) getLocationByZoneName(name string) (location *time.Location, err error) {
	zoneMu.RLock()
	location = zoneMap[name]
	zoneMu.RUnlock()
	if location == nil {
		location, err = time.LoadLocation(name)
		if err != nil {
			err = gerror.Wrapf(err, `time.LoadLocation failed for name "%s"`, name)
		}
		if location != nil {
			zoneMu.Lock()
			zoneMap[name] = location
			zoneMu.Unlock()
		}
	}
	return
}

// Local将时间转换为本地时区。 md5:8eaacff0234ddea5
// ff:取本地时区
// t:
func (t *Time) Local() *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Local()
	return newTime
}
