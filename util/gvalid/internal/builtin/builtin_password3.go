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

// RulePassword3 实现了 `password3` 规则：
// 全局密码格式规则3：
// 必须符合密码规则1，必须包含小写字母、大写字母、数字和特殊字符。
//
// 格式：password3
// md5:b62e78e30236a94a
type RulePassword3 struct{}

func init() {
	Register(RulePassword3{})
}

func (r RulePassword3) Name() string {
	return "password3"
}

func (r RulePassword3) Message() string {
	return "The {field} value `{value}` is not a valid password3 format"
}

func (r RulePassword3) Run(in RunInput) error {
	var value = in.Value.String()
	if gregex.X是否匹配文本(`^[\w\S]{6,18}$`, value) &&
		gregex.X是否匹配文本(`[a-z]+`, value) &&
		gregex.X是否匹配文本(`[A-Z]+`, value) &&
		gregex.X是否匹配文本(`\d+`, value) &&
		gregex.X是否匹配文本(`[^a-zA-Z0-9]+`, value) {
		return nil
	}
	return errors.New(in.Message)
}
