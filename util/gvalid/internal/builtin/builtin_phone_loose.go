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

// RulePhoneLoose implements `phone-loose` rule:
// Loose mobile phone number verification(宽松的手机号验证)
// As long as the 11 digits numbers beginning with
// 13, 14, 15, 16, 17, 18, 19 can pass the verification
// (只要满足 13、14、15、16、17、18、19开头的11位数字都可以通过验证).
//
// Format: phone-loose
type RulePhoneLoose struct{}

func init() {
	Register(RulePhoneLoose{})
}

func (r RulePhoneLoose) Name() string {
	return "phone-loose"
}

func (r RulePhoneLoose) Message() string {
	return "The {field} value `{value}` is not a valid phone number"
}

func (r RulePhoneLoose) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^1(3|4|5|6|7|8|9)\d{9}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
