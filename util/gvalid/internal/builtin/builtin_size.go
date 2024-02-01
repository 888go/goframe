// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	"strconv"
	"strings"
	
	"github.com/888go/goframe/util/gconv"
	)
// RuleSize 实现 `size` 规则:
// 长度必须为 :size。
// 长度计算采用 unicode 字符串方式，这意味着一个中文字符或一个字母的长度都为 1。
//
// 格式: size:size
type RuleSize struct{}

func init() {
	Register(RuleSize{})
}

func (r RuleSize) Name() string {
	return "size"
}

func (r RuleSize) Message() string {
	return "The {field} value `{value}` length must be {size}"
}

func (r RuleSize) Run(in RunInput) error {
	var (
		valueRunes = gconv.Runes(in.Value.String())
		valueLen   = len(valueRunes)
	)
	size, err := strconv.Atoi(in.RulePattern)
	if valueLen != size || err != nil {
		return errors.New(strings.Replace(in.Message, "{size}", strconv.Itoa(size), -1))
	}
	return nil
}
