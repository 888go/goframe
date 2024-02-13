// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 时间类

import (
	"os"
	"strings"
	"sync"
	"time"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
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
func X设置时区(时区 string) (错误 error) {
	setTimeZoneMu.Lock()
	defer setTimeZoneMu.Unlock()
	if setTimeZoneName != "" && !strings.EqualFold(时区, setTimeZoneName) {
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidOperation,
			`process timezone already set using "%s"`,
			setTimeZoneName,
		)
	}
	defer func() {
		if 错误 == nil {
			setTimeZoneName = 时区
		}
	}()

	// 它已经被设置为 time.Local。
	if strings.EqualFold(时区, time.Local.String()) {
		return
	}

	// 从指定名称加载时区信息。
	location, 错误 := time.LoadLocation(时区)
	if 错误 != nil {
		错误 = 错误类.X多层错误码并格式化(错误码类.CodeInvalidParameter, 错误, `time.LoadLocation failed for zone "%s"`, 时区)
		return 错误
	}

	// 更新一次 time.Local。
	time.Local = location

	// 更新*nix系统的时间zone环境变量。
	var (
		envKey   = "TZ"
		envValue = location.String()
	)
	if 错误 = os.Setenv(envKey, envValue); 错误 != nil {
		错误 = 错误类.X多层错误码并格式化(
			错误码类.CodeUnknown,
			错误,
			`set environment failed with key "%s", value "%s"`,
			envKey, envValue,
		)
	}
	return
}

// ToLocation将当前时间转换为指定时区的时间。
func (t *Time) X转换时区Location(时区 *time.Location) *Time {
	newTime := t.X取副本()
	newTime.Time = newTime.Time.In(时区)
	return newTime
}

// ToZone 将当前时间转换为指定时区，如：Asia/Shanghai。
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
			err = 错误类.X多层错误并格式化(err, `time.LoadLocation failed for name "%s"`, name)
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
func (t *Time) X取本地时区() *Time {
	newTime := t.X取副本()
	newTime.Time = newTime.Time.Local()
	return newTime
}
