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

// RulePassport 实现了 `passport` 规则：
// 通用护照格式规则：
// 以字母开头，只包含数字或下划线，长度在6到18个字符之间
//
// 格式：passport md5:198c8f6806d80344
type RulePassport struct{}

func init() {
	Register(RulePassport{})
}

func (r RulePassport) Name() string {
	return "passport"
}

func (r RulePassport) Message() string {
	return "The {field} value `{value}` is not a valid passport format"
}

func (r RulePassport) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^[a-zA-Z]{1}\w{5,17}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
