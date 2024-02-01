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
// RuleMinLength 实现了 `min-length` 规则:
// 字符串长度等于或大于 :min。
// 长度计算采用的是 Unicode 字符串方式，这意味着一个中文字符或英文字母的长度都为 1。
//
// 格式: min-length:min
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
		valueRunes = gconv.Runes(in.Value.String())
		valueLen   = len(valueRunes)
	)
	min, err := strconv.Atoi(in.RulePattern)
	if valueLen < min || err != nil {
		return errors.New(gstr.Replace(in.Message, "{min}", strconv.Itoa(min)))
	}
	return nil
}
