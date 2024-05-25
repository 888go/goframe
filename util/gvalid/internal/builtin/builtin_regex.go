// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"

	"github.com/gogf/gf/v2/text/gregex"
)

// RuleRegex 实现了 `regex` 规则：
// 值应该匹配自定义的正则表达式模式。
//
// 格式：regex:pattern
// md5:51f168fc18898499
type RuleRegex struct{}

func init() {
	Register(RuleRegex{})
}

func (r RuleRegex) Name() string {
	return "regex"
}

func (r RuleRegex) Message() string {
	return "The {field} value `{value}` must be in regex of: {pattern}"
}

func (r RuleRegex) Run(in RunInput) error {
	if !gregex.IsMatchString(in.RulePattern, in.Value.String()) {
		return errors.New(in.Message)
	}
	return nil
}
