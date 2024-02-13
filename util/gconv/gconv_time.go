// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类

import (
	"time"
	
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/gtime"
)

// Time将`any`转换为time.Time类型。
func X取时间(any interface{}, 格式 ...string) time.Time {
	// 它已经是这种类型了。
	if len(格式) == 0 {
		if v, ok := any.(time.Time); ok {
			return v
		}
	}
	if t := X取gtime时间类(any, 格式...); t != nil {
		return t.Time
	}
	return time.Time{}
}

// Duration 将 `any` 转换为 time.Duration 类型。
// 如果 `any` 是字符串，那么它会使用 time.ParseDuration 来进行转换。
// 如果 `any` 是数字类型，则将 `any` 视为纳秒进行转换。
func X取时长(值 interface{}) time.Duration {
	// 它已经是这种类型了。
	if v, ok := 值.(time.Duration); ok {
		return v
	}
	s := String(值)
	if !utils.IsNumeric(s) {
		d, _ := 时间类.X文本取时长(s)
		return d
	}
	return time.Duration(X取整数64位(值))
}

// GTime 将 `any` 类型转换为 *gtime.Time 类型。
// 参数 `format` 可用于指定 `any` 的格式。
// 它将返回与 formats 切片中第一个格式匹配的转换后的值。
// 如果未提供 `format`，当 `any` 为数值类型时，使用 gtime.NewFromTimeStamp 进行转换；
// 当 `any` 为字符串类型时，则使用 gtime.StrToTime 进行转换。
func X取gtime时间类(值 interface{}, 格式 ...string) *时间类.Time {
	if 值 == nil {
		return nil
	}
	if v, ok := 值.(iGTime); ok {
		return v.X取gtime时间类(格式...)
	}
	// 它已经是这种类型了。
	if len(格式) == 0 {
		if v, ok := 值.(*时间类.Time); ok {
			return v
		}
		if t, ok := 值.(time.Time); ok {
			return 时间类.X创建(t)
		}
		if t, ok := 值.(*time.Time); ok {
			return 时间类.X创建(t)
		}
	}
	s := String(值)
	if len(s) == 0 {
		return 时间类.X创建()
	}
	// 使用给定格式进行优先级转换。
	if len(格式) > 0 {
		for _, item := range 格式 {
			t, err := 时间类.StrToTimeFormat别名(s, item)
			if t != nil && err == nil {
				return t
			}
		}
		return nil
	}
	if utils.IsNumeric(s) {
		return 时间类.X创建并从时间戳(X取整数64位(s))
	} else {
		t, _ := 时间类.X转换文本(s)
		return t
	}
}
