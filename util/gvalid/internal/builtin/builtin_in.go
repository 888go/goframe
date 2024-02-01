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
// RuleIn 实现了 `in` 规则：
// 值应该在: value1, value2, ... 之中
//
// 格式：in:value1,value2,...
type RuleIn struct{}

func init() {
	Register(RuleIn{})
}

func (r RuleIn) Name() string {
	return "in"
}

func (r RuleIn) Message() string {
	return "The {field} value `{value}` is not in acceptable range: {pattern}"
}

func (r RuleIn) Run(in RunInput) error {
	var ok bool
	for _, rulePattern := range gstr.SplitAndTrim(in.RulePattern, ",") {
		if in.Option.CaseInsensitive {
			ok = strings.EqualFold(in.Value.String(), strings.TrimSpace(rulePattern))
		} else {
			ok = strings.Compare(in.Value.String(), strings.TrimSpace(rulePattern)) == 0
		}
		if ok {
			return nil
		}
	}
	return errors.New(in.Message)
}
