// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/text/gregex"
)

// RulePassport 实现了 `passport` 规则：
// 通用护照格式规则：
// 以字母开头，仅包含数字或下划线，长度在6到18之间
//
// 格式：passport
type RulePassport struct{}

func init() {
	Register(RulePassport{})
}

func (r RulePassport) Name() string {
	return "passport"
}

func (r RulePassport) Message() string {
	return "The {field} value `{value}` is not a valid passport format"
}

func (r RulePassport) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^[a-zA-Z]{1}\w{5,17}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
