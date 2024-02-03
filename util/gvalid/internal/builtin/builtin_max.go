// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	"strconv"
	
	"github.com/888go/goframe/text/gstr"
)

// RuleMax 实现了 `max` 规则：
// 等于或小于 :max。它同时支持整数和浮点数。
//
// 格式：max:max
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
