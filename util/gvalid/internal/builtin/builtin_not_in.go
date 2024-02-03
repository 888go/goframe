// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	"strings"
	
	"github.com/888go/goframe/text/gstr"
)

// RuleNotIn 实现了“not-in”规则：
// 值不应包含在：value1, value2, ...
//
// 格式：not-in:value1,value2,...
type RuleNotIn struct{}

func init() {
	Register(RuleNotIn{})
}

func (r RuleNotIn) Name() string {
	return "not-in"
}

func (r RuleNotIn) Message() string {
	return "The {field} value `{value}` must not be in range: {pattern}"
}

func (r RuleNotIn) Run(in RunInput) error {
	var (
		ok    = true
		value = in.Value.String()
	)
	for _, rulePattern := range gstr.SplitAndTrim(in.RulePattern, ",") {
		if in.Option.CaseInsensitive {
			ok = !strings.EqualFold(value, strings.TrimSpace(rulePattern))
		} else {
			ok = strings.Compare(value, strings.TrimSpace(rulePattern)) != 0
		}
		if !ok {
			return errors.New(in.Message)
		}
	}
	return nil
}
