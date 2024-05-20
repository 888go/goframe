// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gvalid

import (
	"context"
	"fmt"
	"reflect"
	"runtime"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/internal/intlog"
)

// RuleFunc 是用于数据验证的自定义函数。. md5:7988c41777832ac1
type RuleFunc func(ctx context.Context, in RuleFuncInput) error

// RuleFuncInput 是传递给自定义规则函数 RuleFunc 的输入参数。. md5:071da67c908f30a9
type RuleFuncInput struct {
	// Rule 定义了验证规则字符串，例如 "required"、"between:1,100" 等等。. md5:0903f4201c9e300d
	Rule string

	// Message 指定此规则的自定义错误消息或配置的 i18n 消息。. md5:407649d2c7943432
	Message string

	// Field 指定此规则要验证的字段。. md5:b21049696367d3c3
	Field string

	// ValueType 指定了值的类型，可能为 nil。. md5:b1ad5cfd9a152a1d
	ValueType reflect.Type

	// Value 指定此规则用于验证的值。. md5:29bdb57107181fe6
	Value *gvar.Var

// Data 指定了传递给Validator的数据，它可以是map/结构体类型或nil值。如果你的自定义验证规则不需要这个参数，可以忽略它。
// md5:fd9ebb5b1bdabe03
	Data *gvar.Var
}

var (
// customRuleFuncMap 存储自定义规则函数。
// map[Rule]RuleFunc
// md5:ddde03f9fa92aae7
	customRuleFuncMap = make(map[string]RuleFunc)
)

// RegisterRule 为包注册自定义验证规则和函数。. md5:bb0c3971adfb8935
func RegisterRule(rule string, f RuleFunc) {
	if customRuleFuncMap[rule] != nil {
		intlog.PrintFunc(context.TODO(), func() string {
			return fmt.Sprintf(
				`rule "%s" is overwritten by function "%s"`,
				rule, runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
			)
		})
	}
	customRuleFuncMap[rule] = f
}

// RegisterRuleByMap 通过映射为包注册自定义验证规则。. md5:6f3ae52bddfd4a24
func RegisterRuleByMap(m map[string]RuleFunc) {
	for k, v := range m {
		customRuleFuncMap[k] = v
	}
}

// GetRegisteredRuleMap 返回所有自定义注册的规则及其关联的函数。. md5:3abbd0fbfe9f3c51
func GetRegisteredRuleMap() map[string]RuleFunc {
	if len(customRuleFuncMap) == 0 {
		return nil
	}
	ruleMap := make(map[string]RuleFunc)
	for k, v := range customRuleFuncMap {
		ruleMap[k] = v
	}
	return ruleMap
}

// DeleteRule 从全局包中删除一个或多个自定义定义的验证规则及其关联函数。. md5:474d821f8f0b7fdc
func DeleteRule(rules ...string) {
	for _, rule := range rules {
		delete(customRuleFuncMap, rule)
	}
}
