// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/text/gregex"
)

// RuleMac 实现了 `mac` 规则：
// MAC（介质访问控制）地址规则。
//
// 格式：mac
type RuleMac struct{}

func init() {
	Register(RuleMac{})
}

func (r RuleMac) Name() string {
	return "mac"
}

func (r RuleMac) Message() string {
	return "The {field} value `{value}` is not a valid MAC address"
}

func (r RuleMac) Run(in RunInput) error {
	ok := 正则类.X是否匹配文本(
		`^([0-9A-Fa-f]{2}[\-:]){5}[0-9A-Fa-f]{2}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
