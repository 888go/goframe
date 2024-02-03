// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/text/gregex"
)

// RulePassword3 实现了 `password3` 规则：
// 通用密码格式规则3：
// 必须满足密码规则1，且必须包含小写字母、大写字母、数字和特殊字符。
//
// 格式：password3
type RulePassword3 struct{}

func init() {
	Register(RulePassword3{})
}

func (r RulePassword3) Name() string {
	return "password3"
}

func (r RulePassword3) Message() string {
	return "The {field} value `{value}` is not a valid password3 format"
}

func (r RulePassword3) Run(in RunInput) error {
	var value = in.Value.String()
	if gregex.IsMatchString(`^[\w\S]{6,18}$`, value) &&
		gregex.IsMatchString(`[a-z]+`, value) &&
		gregex.IsMatchString(`[A-Z]+`, value) &&
		gregex.IsMatchString(`\d+`, value) &&
		gregex.IsMatchString(`[^a-zA-Z0-9]+`, value) {
		return nil
	}
	return errors.New(in.Message)
}
