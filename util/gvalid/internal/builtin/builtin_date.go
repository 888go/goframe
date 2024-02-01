// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	"time"
	
	"github.com/888go/goframe/text/gregex"
	)
// RuleDate 实现了 `date` 规则：
// 标准日期格式，例如：2006-01-02、20060102、2006.01.02。
//
// 格式：date
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
	// 支持时间值，例如：gtime.Time（gtime.Time, time.Time）
	if obj, ok := in.Value.Val().(iTime); ok {
		if obj.IsZero() {
			return errors.New(in.Message)
		}
	}
	if !gregex.IsMatchString(
		`\d{4}[\.\-\_/]{0,1}\d{2}[\.\-\_/]{0,1}\d{2}`,
		in.Value.String(),
	) {
		return errors.New(in.Message)
	}
	return nil
}
