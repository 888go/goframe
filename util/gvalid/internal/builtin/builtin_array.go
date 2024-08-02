// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"

	"github.com/888go/goframe/internal/json"
)

// RuleArray 实现了 `array` 规则：
// 值应为数组类型。
//
// 格式：array
// md5:242ac4bf271dd8e1
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
