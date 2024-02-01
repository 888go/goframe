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
// RuleDatetime 实现了 `datetime` 规则：
// 标准日期时间格式，例如：2006-01-02 12:00:00。
//
// 格式：datetime
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
	// 支持时间值，例如：gtime.Time（gtime.Time, time.Time）
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
