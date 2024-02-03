// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	"strconv"
	
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

// RuleMaxLength 实现了 `max-length` 规则：
// 长度等于或小于 :max。
// 长度计算采用的是 unicode 字符串，这意味着一个中文字符或字母的长度都计为 1。
//
// 格式：max-length:max
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
