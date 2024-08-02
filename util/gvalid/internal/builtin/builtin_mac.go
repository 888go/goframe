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

// RuleMac 实现了 `mac` 规则：
// MAC 地址校验。
//
// 格式：mac
// md5:b1c58071a5d631fc
type RuleMac struct{}

func init() {
	Register(RuleMac{})
}

func (r RuleMac) Name() string {
	return "mac"
}

func (r RuleMac) Message() string {
	return "The {field} value `{value}` is not a valid MAC address"
}

func (r RuleMac) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^([0-9A-Fa-f]{2}[\-:]){5}[0-9A-Fa-f]{2}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
