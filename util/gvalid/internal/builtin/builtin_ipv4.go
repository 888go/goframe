// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	
	"github.com/888go/goframe/net/gipv4"
	)
// RuleIpv4 实现了 `ipv4` 规则：
// IPv4。
//
// 格式：ipv4
type RuleIpv4 struct{}

func init() {
	Register(RuleIpv4{})
}

func (r RuleIpv4) Name() string {
	return "ipv4"
}

func (r RuleIpv4) Message() string {
	return "The {field} value `{value}` is not a valid IPv4 address"
}

func (r RuleIpv4) Run(in RunInput) error {
	var (
		ok    bool
		value = in.Value.String()
	)
	if ok = gipv4.Validate(value); !ok {
		return errors.New(in.Message)
	}
	return nil
}
