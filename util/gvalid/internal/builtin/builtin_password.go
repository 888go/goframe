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

// RulePassword 实现了 `password` 规则：
// 全局密码格式规则1：
// 包含任何可见字符，长度在6到18之间。
//
// 格式：password
// md5:174006e615e50650
type RulePassword struct{}

func init() {
	Register(RulePassword{})
}

func (r RulePassword) Name() string {
	return "password"
}

func (r RulePassword) Message() string {
	return "The {field} value `{value}` is not a valid password format"
}

func (r RulePassword) Run(in RunInput) error {
	if !gregex.IsMatchString(`^[\w\S]{6,18}$`, in.Value.String()) {
		return errors.New(in.Message)
	}
	return nil
}
