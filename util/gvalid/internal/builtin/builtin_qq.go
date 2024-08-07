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

// RuleQQ实现了`qq`规则：
// 腾讯QQ号码。
//
// 格式：qq
// md5:0d96642ad3034935
type RuleQQ struct{}

func init() {
	Register(RuleQQ{})
}

func (r RuleQQ) Name() string {
	return "qq"
}

func (r RuleQQ) Message() string {
	return "The {field} value `{value}` is not a valid QQ number"
}

func (r RuleQQ) Run(in RunInput) error {
	ok := gregex.X是否匹配文本(
		`^[1-9][0-9]{4,}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
