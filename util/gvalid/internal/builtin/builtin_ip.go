// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

import (
	"errors"

	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/net/gipv6"
)

// RuleIp 实现了`ip`规则：
// IPv4/IPv6。
//
// 格式：ip md5:29f0c39a26475dfa
type RuleIp struct{}

func init() {
	Register(RuleIp{})
}

func (r RuleIp) Name() string {
	return "ip"
}

func (r RuleIp) Message() string {
	return "The {field} value `{value}` is not a valid IP address"
}

func (r RuleIp) Run(in RunInput) error {
	var (
		ok    bool
		value = in.Value.String()
	)
	if ok = gipv4.Validate(value) || gipv6.Validate(value); !ok {
		return errors.New(in.Message)
	}
	return nil
}
