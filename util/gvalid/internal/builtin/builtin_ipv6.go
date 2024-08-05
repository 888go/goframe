// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"

	"github.com/gogf/gf/v2/net/gipv6"
)

// RuleIpv6 实现了 `ipv6` 规则：
// IPv6。
//
// 格式：ipv6
// md5:409b6971df3398d7
type RuleIpv6 struct{}

func init() {
	Register(RuleIpv6{})
}

func (r RuleIpv6) Name() string {
	return "ipv6"
}

func (r RuleIpv6) Message() string {
	return "The {field} value `{value}` is not a valid IPv6 address"
}

func (r RuleIpv6) Run(in RunInput) error {
	var (
		ok    bool
		value = in.Value.String()
	)
	if ok = gipv6.Validate(value); !ok {
		return errors.New(in.Message)
	}
	return nil
}
