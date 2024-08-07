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

// RuleNotIn 实现了 "not-in" 规则：
// 值不应该在：value1, value2, ...
//
// 格式：not-in:value1,value2,...
// md5:de8746fa53d5a5b4
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
	for _, rulePattern := range gstr.X分割并忽略空值(in.RulePattern, ",") {
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
