// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

import (
	"errors"

	"github.com/gogf/gf/v2/text/gregex"
)

// RuleEmail 实现了 `email` 规则：
// 邮箱地址。
//
// 格式：email md5:238fc353e79c531a
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
	ok := gregex.IsMatchString(
		`^[a-zA-Z0-9_\-\.]+@[a-zA-Z0-9_\-]+(\.[a-zA-Z0-9_\-]+)+$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
