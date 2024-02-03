// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/text/gregex"
)

// RuleNotRegex 实现了 `not-regex` 规则：
// 值不应匹配自定义正则表达式模式。
//
// 格式：not-regex:pattern
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
	if gregex.IsMatchString(in.RulePattern, in.Value.String()) {
		return errors.New(in.Message)
	}
	return nil
}
