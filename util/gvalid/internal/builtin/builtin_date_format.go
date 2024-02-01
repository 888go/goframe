// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	"time"
	
	"github.com/888go/goframe/os/gtime"
	)
// RuleDateFormat 实现了 `date-format` 规则：
// 自定义日期格式。
//
// 格式：date-format:format
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
	// 支持时间值，例如：gtime.Time（gtime.Time, time.Time）
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
