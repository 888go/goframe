// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类

import (
	"time"

	"github.com/888go/goframe/internal/utils"
	gtime "github.com/888go/goframe/os/gtime"
)

// X取时间 将 `any` 转换为 time.X取时间 类型。 md5:2e2c448d3d063180
func X取时间(any interface{}, 格式 ...string) time.Time {
		// 它已经是这种类型了。 md5:9b7d52c77ca28e7b
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

// X取时长 将 `any` 转换为 time.X取时长。
// 如果 `any` 是字符串，那么它使用 time.ParseDuration 进行转换。
// 如果 `any` 是数字，那么它将 `any` 作为纳秒来转换。
// md5:4328f63b0561b4f4
func X取时长(值 interface{}) time.Duration {
		// 它已经是这种类型了。 md5:9b7d52c77ca28e7b
	if v, ok := 值.(time.Duration); ok {
		return v
	}
	s := String(值)
	if !utils.IsNumeric(s) {
		d, _ := gtime.X文本取时长(s)
		return d
	}
	return time.Duration(X取整数64位(值))
}

// X取gtime时间类 将 `any` 转换为 *gtime.Time。
// 参数 `format` 可用于指定 `any` 的格式。
// 它返回与格式切片中第一个格式匹配的转换值。
// 如果没有提供 `format`，则当 `any` 为数字时使用 gtime.NewFromTimeStamp 进行转换，
// 或者当 `any` 为字符串时使用 gtime.StrToTime 进行转换。
// md5:a1c7656c4b134443
func X取gtime时间类(any interface{}, format ...string) *gtime.Time {
	if any == nil {
		return nil
	}
	if v, ok := any.(iGTime); ok {
		return v.X取gtime时间类(format...)
	}
		// 它已经是这种类型了。 md5:9b7d52c77ca28e7b
	if len(format) == 0 {
		if v, ok := any.(*gtime.Time); ok {
			return v
		}
		if t, ok := any.(time.Time); ok {
			return gtime.X创建(t)
		}
		if t, ok := any.(*time.Time); ok {
			return gtime.X创建(t)
		}
	}
	s := String(any)
	if len(s) == 0 {
		return gtime.X创建()
	}
		// 使用给定格式进行优先级转换。 md5:df8606795bd8c76b
	if len(format) > 0 {
		for _, item := range format {
			t, err := gtime.StrToTimeFormat别名(s, item)
			if t != nil && err == nil {
				return t
			}
		}
		return nil
	}
	if utils.IsNumeric(s) {
		return gtime.X创建并从时间戳(X取整数64位(s))
	} else {
		t, _ := gtime.X转换文本(s)
		return t
	}
}
