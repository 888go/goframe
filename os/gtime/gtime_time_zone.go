// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 时间类

import (
	"os"
	"strings"
	"sync"
	"time"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

var (
	setTimeZoneMu   sync.Mutex
	setTimeZoneName string
	zoneMap         = make(map[string]*time.Location)
	zoneMu          sync.RWMutex
)

// X设置时区 设置当前整个进程的时间区域。
// 参数 `zone` 是一个指定对应时区的区域字符串，例如：Asia/Shanghai。
// 
// 请注意：
// 1. 应在导入 "time" 包之前调用此函数。
// 2. 此函数仅需调用一次。
// 3. 请参阅问题：https://github.com/golang/go/issues/34814
// md5:4d2c0d7e82a0e0f8
func X设置时区(时区 string) (错误 error) {
	setTimeZoneMu.Lock()
	defer setTimeZoneMu.Unlock()
	if setTimeZoneName != "" && !strings.EqualFold(时区, setTimeZoneName) {
		return gerror.X创建错误码并格式化(
			gcode.CodeInvalidOperation,
			`process timezone already set using "%s"`,
			setTimeZoneName,
		)
	}
	defer func() {
		if 错误 == nil {
			setTimeZoneName = 时区
		}
	}()

		// 它已经被设置为time.Local。 md5:1fa5641c118746d6
	if strings.EqualFold(时区, time.Local.String()) {
		return
	}

		// 从指定的名称加载区域信息。 md5:dada678d8dfb8df3
	location, 错误 := time.LoadLocation(时区)
	if 错误 != nil {
		错误 = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, 错误, `time.LoadLocation failed for zone "%s"`, 时区)
		return 错误
	}

		// 为了保险起见，更新一次time.Local。 md5:b8b9f7daf1488924
	time.Local = location

		// 更新*nix系统中的时区环境。 md5:e9774c5a2d209c8d
	var (
		envKey   = "TZ"
		envValue = location.String()
	)
	if 错误 = os.Setenv(envKey, envValue); 错误 != nil {
		错误 = gerror.X多层错误码并格式化(
			gcode.CodeUnknown,
			错误,
			`set environment failed with key "%s", value "%s"`,
			envKey, envValue,
		)
	}
	return
}

// X转换时区Location 将当前时间转换为指定时区的时间。 md5:ee4fd7d4de93340a
func (t *Time) X转换时区Location(时区 *time.Location) *Time {
	newTime := t.X取副本()
	newTime.Time = newTime.Time.In(时区)
	return newTime
}

// X转换时区 将当前时间转换为指定的时区，如 Asia/Shanghai。 md5:1226213d40f57eb2
func (t *Time) X转换时区(时区 string) (*Time, error) {
	if location, err := t.getLocationByZoneName(时区); err == nil {
		return t.X转换时区Location(location), nil
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
			err = gerror.X多层错误并格式化(err, `time.LoadLocation failed for name "%s"`, name)
		}
		if location != nil {
			zoneMu.Lock()
			zoneMap[name] = location
			zoneMu.Unlock()
		}
	}
	return
}

// X取本地时区将时间转换为本地时区。 md5:8eaacff0234ddea5
func (t *Time) X取本地时区() *Time {
	newTime := t.X取副本()
	newTime.Time = newTime.Time.Local()
	return newTime
}
