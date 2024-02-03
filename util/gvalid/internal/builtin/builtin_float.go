// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	"strconv"
)

// RuleFloat 实现了 `float` 规则：
// 浮点数规则。请注意，整数实际上是一个浮点数。
//
// 格式：float
type RuleFloat struct{}

func init() {
	Register(RuleFloat{})
}

func (r RuleFloat) Name() string {
	return "float"
}

func (r RuleFloat) Message() string {
	return "The {field} value `{value}` is not of valid float type"
}

func (r RuleFloat) Run(in RunInput) error {
	if _, err := strconv.ParseFloat(in.Value.String(), 10); err == nil {
		return nil
	}
	return errors.New(in.Message)
}
