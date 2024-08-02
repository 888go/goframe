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
	"strings"

	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// RuleLength 实现了 `length` 规则：
// 字符长度在 :min 和 :max 之间。
// 长度是通过Unicode字符串计算的，这意味着一个中文字符或字母的长度都为1。
//
// 格式：length:min,max
// md5:a3bd03423a5a8b7e
type RuleLength struct{}

func init() {
	Register(RuleLength{})
}

func (r RuleLength) Name() string {
	return "length"
}

func (r RuleLength) Message() string {
	return "The {field} value `{value}` length must be between {min} and {max}"
}

func (r RuleLength) Run(in RunInput) error {
	var (
		valueRunes = gconv.Runes(in.Value.String())
		valueLen   = len(valueRunes)
	)
	var (
		min   = 0
		max   = 0
		array = strings.Split(in.RulePattern, ",")
	)
	if len(array) > 0 {
		if v, err := strconv.Atoi(strings.TrimSpace(array[0])); err == nil {
			min = v
		}
	}
	if len(array) > 1 {
		if v, err := strconv.Atoi(strings.TrimSpace(array[1])); err == nil {
			max = v
		}
	}
	if valueLen < min || valueLen > max {
		return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
			"{min}": strconv.Itoa(min),
			"{max}": strconv.Itoa(max),
		}))
	}
	return nil
}
