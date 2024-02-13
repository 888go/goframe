// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/text/gregex"
)

// RuleEmail 实现了 `email` 规则：
// 电子邮件地址。
//
// 格式：email
type RuleEmail struct{}

func init() {
	Register(RuleEmail{})
}

func (r RuleEmail) Name() string {
	return "email"
}

func (r RuleEmail) Message() string {
	return "The {field} value `{value}` is not a valid email address"
}

func (r RuleEmail) Run(in RunInput) error {
	ok := 正则类.X是否匹配文本(
		`^[a-zA-Z0-9_\-\.]+@[a-zA-Z0-9_\-]+(\.[a-zA-Z0-9_\-]+)+$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
