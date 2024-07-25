// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
)

// RuleSize 实现了 `size` 规则：
// 长度必须为 :size。
// 长度计算使用的是 unicode 字符串，这意味着一个中文字符或字母的长度都为 1。
//
// 格式：size:size md5:89f4bf85bbd610c6
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
