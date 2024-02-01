// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	"strconv"
	"strings"
	
	"github.com/888go/goframe/text/gstr"
	)
// RuleBetween 实现了 `between` 规则：
// 范围在 :min 和 :max 之间。它同时支持整数和浮点数。
//
// 格式：between:min,max
type RuleBetween struct{}

func init() {
	Register(RuleBetween{})
}

func (r RuleBetween) Name() string {
	return "between"
}

func (r RuleBetween) Message() string {
	return "The {field} value `{value}` must be between {min} and {max}"
}

func (r RuleBetween) Run(in RunInput) error {
	var (
		array = strings.Split(in.RulePattern, ",")
		min   = float64(0)
		max   = float64(0)
	)
	if len(array) > 0 {
		if v, err := strconv.ParseFloat(strings.TrimSpace(array[0]), 10); err == nil {
			min = v
		}
	}
	if len(array) > 1 {
		if v, err := strconv.ParseFloat(strings.TrimSpace(array[1]), 10); err == nil {
			max = v
		}
	}
	valueF, err := strconv.ParseFloat(in.Value.String(), 10)
	if valueF < min || valueF > max || err != nil {
		return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
			"{min}": strconv.FormatFloat(min, 'f', -1, 64),
			"{max}": strconv.FormatFloat(max, 'f', -1, 64),
		}))
	}
	return nil
}
