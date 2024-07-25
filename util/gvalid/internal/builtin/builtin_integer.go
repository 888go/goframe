// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"strconv"
)

// RuleInteger 实现了 `integer` 规则：
// 整数。
//
// 格式：integer md5:2d9eab0afd545045
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
