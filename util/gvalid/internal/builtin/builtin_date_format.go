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

	"github.com/gogf/gf/v2/os/gtime"
)

// RuleDateFormat 实现了 `date-format` 规则：
// 定制日期格式。
//
// 格式：date-format:format
// md5:e25dffe38dd1e8e8
type RuleDateFormat struct{}

func init() {
	Register(RuleDateFormat{})
}

func (r RuleDateFormat) Name() string {
	return "date-format"
}

func (r RuleDateFormat) Message() string {
	return "The {field} value `{value}` does not match the format: {pattern}"
}

func (r RuleDateFormat) Run(in RunInput) error {
	type iTime interface {
		Date() (year int, month time.Month, day int)
		IsZero() bool
	}
	// 支持时间值，例如：gtime.Time/*gtime.Time, time.Time/*time.Time。. md5:fc74717f7b27de8d
	if obj, ok := in.Value.Val().(iTime); ok {
		if obj.IsZero() {
			return errors.New(in.Message)
		}
		return nil
	}
	if _, err := gtime.StrToTimeFormat(in.Value.String(), in.RulePattern); err != nil {
		return errors.New(in.Message)
	}
	return nil
}
