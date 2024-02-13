// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	"strconv"
	"strings"
	
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

// RuleLength 实现了 `length` 规则：
// 长度在 :min 和 :max 之间。
// 长度计算采用的是 unicode 字符串，这意味着一个中文字符或字母的长度都为 1。
//
// 格式：length:min,max
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
		valueRunes = 转换类.X取字符数组(in.Value.String())
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
		return errors.New(文本类.Map替换(in.Message, map[string]string{
			"{min}": strconv.Itoa(min),
			"{max}": strconv.Itoa(max),
		}))
	}
	return nil
}
