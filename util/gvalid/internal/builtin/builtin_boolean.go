// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	"strings"
	)
// RuleBoolean 实现了 `boolean` 规则：
// Boolean(1,true,on,yes:true | 0,false,off,no,"":false)
//
// 格式：boolean
// 注释翻译：
// RuleBoolean 实现了布尔值规则功能：
// 参数可以为以下布尔表示形式，转换为布尔值：
// - 1、true、on 或 yes 表示 true
// - 0、false、off 或 no 表示 false
// 空字符串（""）也表示 false
//
// 使用格式：boolean
type RuleBoolean struct{}

// boolMap 定义了布尔值。
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
