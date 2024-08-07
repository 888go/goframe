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

// RuleDomain 实现了 `domain` 规则：
// 域名。
//
// 格式：domain
// md5:c411d3a07b620a12
type RuleDomain struct{}

func init() {
	Register(RuleDomain{})
}

func (r RuleDomain) Name() string {
	return "domain"
}

func (r RuleDomain) Message() string {
	return "The {field} value `{value}` is not a valid domain format"
}

func (r RuleDomain) Run(in RunInput) error {
	ok := gregex.X是否匹配文本(
		`^([0-9a-zA-Z][0-9a-zA-Z\-]{0,62}\.)+([a-zA-Z]{0,62})$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
