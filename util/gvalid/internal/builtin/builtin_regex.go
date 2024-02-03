// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/text/gregex"
)

// RuleRegex 实现了 `regex` 规则：
// 值应当匹配自定义的正则表达式模式。
//
// 格式：regex:pattern
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
