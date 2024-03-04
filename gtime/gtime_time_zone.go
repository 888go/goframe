// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

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

// SetTimeZone 设置当前整个进程的时区。
// 参数 `zone` 是一个区域字符串，用于指定相应时区，例如：Asia/Shanghai。
//
// **非常重要，请注意**：
// 1. 此函数应在导入 "time" 包之前调用。
// 2. 此函数应仅调用一次。
// 3. 参考相关问题：https://github.com/golang/go/issues/34814
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

	// 它已经被设置为 time.Local。
	if strings.EqualFold(zone, time.Local.String()) {
		return
	}

	// 从指定名称加载时区信息。
	location, err := time.LoadLocation(zone)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `time.LoadLocation failed for zone "%s"`, zone)
		return err
	}

	// 更新一次 time.Local。
	time.Local = location

	// 更新*nix系统的时间zone环境变量。
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

// ToLocation将当前时间转换为指定时区的时间。
func (t *Time) ToLocation(location *time.Location) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.In(location)
	return newTime
}

// ToZone 将当前时间转换为指定时区，如：Asia/Shanghai。
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

// Local将时间转换为本地时区。
func (t *Time) Local() *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Local()
	return newTime
}
