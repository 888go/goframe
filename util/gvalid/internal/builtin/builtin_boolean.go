// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"strings"
)

// RuleBoolean 实现了 `boolean` 规则：
// Boolean(1, true, on, yes: true | 0, false, off, no, "": false)
// 
// 格式：boolean
// md5:b799fb342217ec2b
type RuleBoolean struct{}

// boolMap 定义布尔值。 md5:daf0be92a3ad3752
var boolMap = map[string]struct{}{
	"1":     {},
	"true":  {},
	"on":    {},
	"yes":   {},
	"":      {},
	"0":     {},
	"false": {},
	"off":   {},
	"no":    {},
}

func init() {
	Register(RuleBoolean{})
}

func (r RuleBoolean) Name() string {
	return "boolean"
}

func (r RuleBoolean) Message() string {
	return "The {field} value `{value}` field must be true or false"
}

func (r RuleBoolean) Run(in RunInput) error {
	if _, ok := boolMap[strings.ToLower(in.Value.String())]; ok {
		return nil
	}
	return errors.New(in.Message)
}
