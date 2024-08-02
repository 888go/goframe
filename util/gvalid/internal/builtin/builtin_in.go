// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"strings"

	gstr "github.com/888go/goframe/text/gstr"
)

// RuleIn 实现了 `in` 规则：
// 值应包含在：value1, value2, ...
//
// 格式：in:value1,value2,...
// md5:b2270542c36429af
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
	var (
		ok               bool
		inputValueString = in.Value.String()
	)

	for _, rulePattern := range gstr.SplitAndTrim(in.RulePattern, ",") {
		if in.Option.CaseInsensitive {
			ok = strings.EqualFold(inputValueString, strings.TrimSpace(rulePattern))
		} else {
			ok = strings.Compare(inputValueString, strings.TrimSpace(rulePattern)) == 0
		}
		if ok {
			return nil
		}
	}
	return errors.New(in.Message)
}
