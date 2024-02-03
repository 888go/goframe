// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/net/gipv4"
	"github.com/888go/goframe/net/gipv6"
)

// RuleIp 实现了 `ip` 规则：
// 适用于 IPv4/IPv6。
//
// 格式：ip
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
