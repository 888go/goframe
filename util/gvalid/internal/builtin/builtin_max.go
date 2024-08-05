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

// RuleMax 实现了 `max` 规则：
// 小于或等于 :max。它支持整数和浮点数。
//
// 格式：max:max
// md5:178b1e5315ab61af
type RuleMax struct{}

func init() {
	Register(RuleMax{})
}

func (r RuleMax) Name() string {
	return "max"
}

func (r RuleMax) Message() string {
	return "The {field} value `{value}` must be equal or lesser than {max}"
}

func (r RuleMax) Run(in RunInput) error {
	var (
		max, err1    = strconv.ParseFloat(in.RulePattern, 10)
		valueN, err2 = strconv.ParseFloat(in.Value.String(), 10)
	)
	if valueN > max || err1 != nil || err2 != nil {
		return errors.New(gstr.Replace(in.Message, "{max}", strconv.FormatFloat(max, 'f', -1, 64)))
	}
	return nil
}
