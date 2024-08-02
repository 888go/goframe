// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"time"

	gtime "github.com/888go/goframe/os/gtime"
)

// RuleDatetime 实现了 `datetime` 规则：
// 标准日期时间格式，例如：2006-01-02 12:00:00。
//
// 格式：datetime
// md5:cd1da4eb292e6ea1
type RuleDatetime struct{}

func init() {
	Register(RuleDatetime{})
}

func (r RuleDatetime) Name() string {
	return "datetime"
}

func (r RuleDatetime) Message() string {
	return "The {field} value `{value}` is not a valid datetime"
}

func (r RuleDatetime) Run(in RunInput) error {
	type iTime interface {
		Date() (year int, month time.Month, day int)
		IsZero() bool
	}
	// 支持时间值，例如：gtime.Time/*gtime.Time, time.Time/*time.Time。 md5:fc74717f7b27de8d
	if obj, ok := in.Value.Val().(iTime); ok {
		if obj.IsZero() {
			return errors.New(in.Message)
		}
		return nil
	}
	if _, err := gtime.StrToTimeFormat(in.Value.String(), `Y-m-d H:i:s`); err != nil {
		return errors.New(in.Message)
	}
	return nil
}
