// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	"strconv"
	)
// RuleInteger 实现了 `integer` 规则：
// 整数。
//
// 格式：整数
type RuleInteger struct{}

func init() {
	Register(RuleInteger{})
}

func (r RuleInteger) Name() string {
	return "integer"
}

func (r RuleInteger) Message() string {
	return "The {field} value `{value}` is not an integer"
}

func (r RuleInteger) Run(in RunInput) error {
	if _, err := strconv.Atoi(in.Value.String()); err == nil {
		return nil
	}
	return errors.New(in.Message)
}
