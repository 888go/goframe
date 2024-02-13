// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/text/gregex"
)

// RulePassword2 实现了 `password2` 规则：
// 通用密码格式规则2：
// 必须符合密码规则1，并且必须包含大小写字母和数字。
//
// 格式：password2
type RulePassword2 struct{}

func init() {
	Register(RulePassword2{})
}

func (r RulePassword2) Name() string {
	return "password2"
}

func (r RulePassword2) Message() string {
	return "The {field} value `{value}` is not a valid password2 format"
}

func (r RulePassword2) Run(in RunInput) error {
	var value = in.Value.String()
	if 正则类.X是否匹配文本(`^[\w\S]{6,18}$`, value) &&
		正则类.X是否匹配文本(`[a-z]+`, value) &&
		正则类.X是否匹配文本(`[A-Z]+`, value) &&
		正则类.X是否匹配文本(`\d+`, value) {
		return nil
	}
	return errors.New(in.Message)
}
