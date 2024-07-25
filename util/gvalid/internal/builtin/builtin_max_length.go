// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"strconv"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// RuleMaxLength 实现了 `max-length` 规则：
// 长度小于或等于 :max。长度是根据 Unicode 字符串计算的，这意味着一个中文字符或字母的长度均为 1。
//
// 格式：max-length:max md5:79ec85f367534ee2
type RuleMaxLength struct{}

func init() {
	Register(RuleMaxLength{})
}

func (r RuleMaxLength) Name() string {
	return "max-length"
}

func (r RuleMaxLength) Message() string {
	return "The {field} value `{value}` length must be equal or lesser than {max}"
}

func (r RuleMaxLength) Run(in RunInput) error {
	var (
		valueRunes = gconv.Runes(in.Value.String())
		valueLen   = len(valueRunes)
	)
	max, err := strconv.Atoi(in.RulePattern)
	if valueLen > max || err != nil {
		return errors.New(gstr.Replace(in.Message, "{max}", strconv.Itoa(max)))
	}
	return nil
}
