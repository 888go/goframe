// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gconv

import (
	"time"
	
	"github.com/888go/goframe/gconv/internal/utils"
	"github.com/gogf/gf/v2/os/gtime"
)

// Time将`any`转换为time.Time类型。
func Time(any interface{}, format ...string) time.Time {
	// 它已经是这种类型了。
	if len(format) == 0 {
		if v, ok := any.(time.Time); ok {
			return v
		}
	}
	if t := GTime(any, format...); t != nil {
		return t.Time
	}
	return time.Time{}
}

// Duration 将 `any` 转换为 time.Duration 类型。
// 如果 `any` 是字符串，那么它会使用 time.ParseDuration 来进行转换。
// 如果 `any` 是数字类型，则将 `any` 视为纳秒进行转换。
func Duration(any interface{}) time.Duration {
	// 它已经是这种类型了。
	if v, ok := any.(time.Duration); ok {
		return v
	}
	s := String(any)
	if !utils.IsNumeric(s) {
		d, _ := gtime.ParseDuration(s)
		return d
	}
	return time.Duration(Int64(any))
}

// GTime 将 `any` 类型转换为 *gtime.Time 类型。
// 参数 `format` 可用于指定 `any` 的格式。
// 它将返回与 formats 切片中第一个格式匹配的转换后的值。
// 如果未提供 `format`，当 `any` 为数值类型时，使用 gtime.NewFromTimeStamp 进行转换；
// 当 `any` 为字符串类型时，则使用 gtime.StrToTime 进行转换。
func GTime(any interface{}, format ...string) *gtime.Time {
	if any == nil {
		return nil
	}
	if v, ok := any.(iGTime); ok {
		return v.GTime(format...)
	}
	// 它已经是这种类型了。
	if len(format) == 0 {
		if v, ok := any.(*gtime.Time); ok {
			return v
		}
		if t, ok := any.(time.Time); ok {
			return gtime.New(t)
		}
		if t, ok := any.(*time.Time); ok {
			return gtime.New(t)
		}
	}
	s := String(any)
	if len(s) == 0 {
		return gtime.New()
	}
	// 使用给定格式进行优先级转换。
	if len(format) > 0 {
		for _, item := range format {
			t, err := gtime.StrToTimeFormat(s, item)
			if t != nil && err == nil {
				return t
			}
		}
		return nil
	}
	if utils.IsNumeric(s) {
		return gtime.NewFromTimeStamp(Int64(s))
	} else {
		t, _ := gtime.StrToTime(s)
		return t
	}
}
