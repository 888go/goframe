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

	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// RuleMinLength 实现了 `min-length` 规则：
// 长度等于或大于 :min。
// 长度计算使用的是 unicode 字符串，这意味着一个中文字符或字母的长度都为 1。
//
// 格式：min-length:min
// md5:565104f674e70ec7
type RuleMinLength struct{}

func init() {
	Register(RuleMinLength{})
}

func (r RuleMinLength) Name() string {
	return "min-length"
}

func (r RuleMinLength) Message() string {
	return "The {field} value `{value}` length must be equal or greater than {min}"
}

func (r RuleMinLength) Run(in RunInput) error {
	var (
		valueRunes = gconv.X取字符切片(in.Value.String())
		valueLen   = len(valueRunes)
	)
	min, err := strconv.Atoi(in.RulePattern)
	if valueLen < min || err != nil {
		return errors.New(gstr.X替换(in.Message, "{min}", strconv.Itoa(min)))
	}
	return nil
}
