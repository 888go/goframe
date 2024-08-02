// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"strconv"
	"strings"

	gstr "github.com/888go/goframe/text/gstr"
)

// RuleBetween 实现了 `between` 规则：
// 范围在 :min 和 :max 之间。它支持整数和浮点数。
//
// 格式：between:min,max
// md5:99eb04ab3fef89e8
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
