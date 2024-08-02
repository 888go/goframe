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

// RulePassword2 实现了 `password2` 规则：
// 通用密码格式规则2：
// 必须满足密码规则1，且必须包含大小写字母和数字。
//
// 格式：password2
// md5:97a3c93771b5b020
type RulePassword2 struct{}

func init() {
	Register(RulePassword2{})
}

func (r RulePassword2) Name() string {
	return "password2"
}

func (r RulePassword2) Message() string {
	return "The {field} value `{value}` is not a valid password2 format"
}

func (r RulePassword2) Run(in RunInput) error {
	var value = in.Value.String()
	if gregex.IsMatchString(`^[\w\S]{6,18}$`, value) &&
		gregex.IsMatchString(`[a-z]+`, value) &&
		gregex.IsMatchString(`[A-Z]+`, value) &&
		gregex.IsMatchString(`\d+`, value) {
		return nil
	}
	return errors.New(in.Message)
}
