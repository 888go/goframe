// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"strconv"
)

// RuleFloat 实现了 `float` 规则：
// 浮点数。需要注意的是，整数实际上也是浮点数。
//
// 格式：float
// md5:1de569690e034998
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
