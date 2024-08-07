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

// RulePostcode 实现了 `postcode` 规则：
// 邮政编码数字。
//
// 格式：postcode
// md5:3a14b940da5b3d7b
type RulePostcode struct{}

func init() {
	Register(RulePostcode{})
}

func (r RulePostcode) Name() string {
	return "postcode"
}

func (r RulePostcode) Message() string {
	return "The {field} value `{value}` is not a valid postcode format"
}

func (r RulePostcode) Run(in RunInput) error {
	ok := gregex.X是否匹配文本(
		`^\d{6}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
