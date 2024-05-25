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

	"github.com/gogf/gf/v2/text/gstr"
)

// RuleMin 实现了`min`规则：
// 大于或等于`:min`。它支持整数和浮点数。
//
// 格式：min:min
// md5:3eacce55c94e5617
type RuleMin struct{}

func init() {
	Register(RuleMin{})
}

func (r RuleMin) Name() string {
	return "min"
}

func (r RuleMin) Message() string {
	return "The {field} value `{value}` must be equal or greater than {min}"
}

func (r RuleMin) Run(in RunInput) error {
	var (
		min, err1    = strconv.ParseFloat(in.RulePattern, 10)
		valueN, err2 = strconv.ParseFloat(in.Value.String(), 10)
	)
	if valueN < min || err1 != nil || err2 != nil {
		return errors.New(gstr.Replace(in.Message, "{min}", strconv.FormatFloat(min, 'f', -1, 64)))
	}
	return nil
}
