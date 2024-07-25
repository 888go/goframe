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

// RuleUrl 实现了 `url` 规则：
// URL。
//
// 格式：url md5:eb24acb21d876558
type RuleUrl struct{}

func init() {
	Register(RuleUrl{})
}

func (r RuleUrl) Name() string {
	return "url"
}

func (r RuleUrl) Message() string {
	return "The {field} value `{value}` is not a valid URL address"
}

func (r RuleUrl) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
