// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/text/gregex"
)

// RulePassword 实现了 `password` 规则：
// 通用密码格式规则1：
// 包含任何可见字符，长度在6到18之间。
//
// 格式：password
type RulePassword struct{}

func init() {
	Register(RulePassword{})
}

func (r RulePassword) Name() string {
	return "password"
}

func (r RulePassword) Message() string {
	return "The {field} value `{value}` is not a valid password format"
}

func (r RulePassword) Run(in RunInput) error {
	if !正则类.X是否匹配文本(`^[\w\S]{6,18}$`, in.Value.String()) {
		return errors.New(in.Message)
	}
	return nil
}
