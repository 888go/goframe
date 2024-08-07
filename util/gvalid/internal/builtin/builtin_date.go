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

	gregex "github.com/888go/goframe/text/gregex"
)

// RuleDate 实现了 `date` 规则：
// 标准日期格式，例如：2006-01-02, 20060102, 2006.01.02。
//
// 格式：date
// md5:91044156fb254923
type RuleDate struct{}

func init() {
	Register(RuleDate{})
}

func (r RuleDate) Name() string {
	return "date"
}

func (r RuleDate) Message() string {
	return "The {field} value `{value}` is not a valid date"
}

func (r RuleDate) Run(in RunInput) error {
	type iTime interface {
		Date() (year int, month time.Month, day int)
		IsZero() bool
	}
		// 支持时间值，例如：gtime.Time/*gtime.Time, time.Time/*time.Time。 md5:fc74717f7b27de8d
	if obj, ok := in.Value.X取值().(iTime); ok {
		if obj.IsZero() {
			return errors.New(in.Message)
		}
	}
	if !gregex.X是否匹配文本(
		`\d{4}[\.\-\_/]{0,1}\d{2}[\.\-\_/]{0,1}\d{2}`,
		in.Value.String(),
	) {
		return errors.New(in.Message)
	}
	return nil
}
