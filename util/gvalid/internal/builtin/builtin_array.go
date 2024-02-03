// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/internal/json"
)

// RuleArray 实现了 `array` 规则：
// 值的类型应为数组。
//
// 格式：array
type RuleArray struct{}

func init() {
	Register(RuleArray{})
}

func (r RuleArray) Name() string {
	return "array"
}

func (r RuleArray) Message() string {
	return "The {field} value `{value}` is not of valid array type"
}

func (r RuleArray) Run(in RunInput) error {
	if in.Value.IsSlice() {
		return nil
	}
	if json.Valid(in.Value.Bytes()) {
		value := in.Value.String()
		if len(value) > 1 && value[0] == '[' && value[len(value)-1] == ']' {
			return nil
		}
	}
	return errors.New(in.Message)
}
