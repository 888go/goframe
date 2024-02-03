// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/text/gregex"
)

// RulePostcode 实现了 `postcode` 规则：
// 邮政编码号码。
//
// 格式：postcode
type RulePostcode struct{}

func init() {
	Register(RulePostcode{})
}

func (r RulePostcode) Name() string {
	return "postcode"
}

func (r RulePostcode) Message() string {
	return "The {field} value `{value}` is not a valid postcode format"
}

func (r RulePostcode) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^\d{6}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
