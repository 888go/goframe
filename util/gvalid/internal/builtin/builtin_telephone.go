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

// RuleTelephone 实现了 `telephone` 规则：
// "XXXX-XXXXXXX"
// "XXXX-XXXXXXXX"
// "XXX-XXXXXXX"
// "XXX-XXXXXXXX"
// "XXXXXXX"
// "XXXXXXXX"
//
// 格式：电话号码
// md5:993bc2476599c2b2
type RuleTelephone struct{}

func init() {
	Register(RuleTelephone{})
}

func (r RuleTelephone) Name() string {
	return "telephone"
}

func (r RuleTelephone) Message() string {
	return "The {field} value `{value}` is not a valid telephone number"
}

func (r RuleTelephone) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^((\d{3,4})|\d{3,4}-)?\d{7,8}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
