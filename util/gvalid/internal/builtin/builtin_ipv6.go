// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/net/gipv6"
)

// RuleIpv6 实现了 `ipv6` 规则：
// IPv6。
//
// 格式：ipv6
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
