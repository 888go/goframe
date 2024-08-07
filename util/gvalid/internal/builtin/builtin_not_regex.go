// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"

	gregex "github.com/888go/goframe/text/gregex"
)

// RuleNotRegex 实现了 `not-regex` 规则：
// 值不应匹配自定义正则表达式模式。
//
// 格式：not-regex:pattern
// md5:ae2c4d23553ee3de
type RuleNotRegex struct{}

func init() {
	Register(RuleNotRegex{})
}

func (r RuleNotRegex) Name() string {
	return "not-regex"
}

func (r RuleNotRegex) Message() string {
	return "The {field} value `{value}` should not be in regex of: {pattern}"
}

func (r RuleNotRegex) Run(in RunInput) error {
	if gregex.X是否匹配文本(in.RulePattern, in.Value.String()) {
		return errors.New(in.Message)
	}
	return nil
}
