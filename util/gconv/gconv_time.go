// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gconv

import (
	"time"

	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/os/gtime"
)

// Time 将 `any` 转换为 time.Time 类型。 md5:2e2c448d3d063180
func Time(any interface{}, format ...string) time.Time {
	// 它已经是这种类型了。 md5:9b7d52c77ca28e7b
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

// Duration 将 `any` 转换为 time.Duration。
// 如果 `any` 是字符串，那么它使用 time.ParseDuration 进行转换。
// 如果 `any` 是数字，那么它将 `any` 作为纳秒来转换。
// md5:4328f63b0561b4f4
func Duration(any interface{}) time.Duration {
	// 它已经是这种类型了。 md5:9b7d52c77ca28e7b
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

// GTime 将 `any` 转换为 *gtime.Time。
// 参数 `format` 可用于指定 `any` 的格式。
// 它返回与格式切片中第一个格式匹配的转换值。
// 如果没有提供 `format`，则当 `any` 为数字时使用 gtime.NewFromTimeStamp 进行转换，
// 或者当 `any` 为字符串时使用 gtime.StrToTime 进行转换。
// md5:a1c7656c4b134443
func GTime(any interface{}, format ...string) *gtime.Time {
	if any == nil {
		return nil
	}
	if v, ok := any.(iGTime); ok {
		return v.GTime(format...)
	}
	// 它已经是这种类型了。 md5:9b7d52c77ca28e7b
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
	// 使用给定格式进行优先级转换。 md5:df8606795bd8c76b
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
