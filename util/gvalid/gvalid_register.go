// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvalid

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/intlog"
)

// RuleFunc 是用于数据验证的自定义函数。
type RuleFunc func(ctx context.Context, in RuleFuncInput) error

// RuleFuncInput 用于存储传递给自定义规则函数 RuleFunc 的输入参数。
type RuleFuncInput struct {
	// Rule 指定验证规则字符串，如 "必填", "范围:1,100" 等等。
	Rule string

	// Message 指定此规则的自定义错误消息或配置的 i18n 消息。
	Message string

	// Field 指定该规则进行验证的字段。
	Field string

	// ValueType 指定值的类型，该值可能为 nil。
	ValueType reflect.Type

	// Value 指定此规则验证的值。
	Value *gvar.Var

// Data 指定传递给 Validator 的 `data`，它可以是 map 或 struct 类型，也可以是 nil 值。
// 如果在自定义验证规则中并不真正需要这个参数，你可以忽略 `Data`。
	Data *gvar.Var
}

var (
// customRuleFuncMap 用于存储自定义规则函数。
// map[Rule]RuleFunc 表示键为 Rule 类型，值为 RuleFunc 类型的映射表。
	customRuleFuncMap = make(map[string]RuleFunc)
)

// RegisterRule 注册自定义验证规则及其相关函数，供包内使用。
func RegisterRule(rule string, f RuleFunc) {
	if customRuleFuncMap[rule] != nil {
		intlog.PrintFunc(context.TODO(), func() string {
			return fmt.Sprintf(
				`rule "%s" is overwrotten by function "%s"`,
				rule, runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
			)
		})
	}
	customRuleFuncMap[rule] = f
}

// RegisterRuleByMap 通过映射表为包注册自定义验证规则。
func RegisterRuleByMap(m map[string]RuleFunc) {
	for k, v := range m {
		customRuleFuncMap[k] = v
	}
}

// GetRegisteredRuleMap 返回所有已注册的自定义规则及其关联的函数。
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

// DeleteRule 从全局包中删除一个或多个自定义验证规则及其关联函数。
func DeleteRule(rules ...string) {
	for _, rule := range rules {
		delete(customRuleFuncMap, rule)
	}
}
